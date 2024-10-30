use anyhow::Result;
use clap::Parser;
use futures::StreamExt;
use kona_host::HostCli;
use log::info;
use op_succinct_host_utils::{
    fetcher::{BlockInfo, CacheMode, OPSuccinctDataFetcher},
    get_proof_stdin,
    stats::ExecutionStats,
    witnessgen::WitnessGenExecutor,
    ProgramType,
};
use rayon::iter::{IndexedParallelIterator, IntoParallelRefIterator, ParallelIterator};
use serde::{Deserialize, Serialize};
use sp1_sdk::{utils, ProverClient};
use std::{
    cmp::{max, min},
    collections::HashMap,
    fs::{self, OpenOptions},
    future::Future,
    io::Seek,
    path::PathBuf,
    time::Instant,
};
use tokio::task::block_in_place;

pub const MULTI_BLOCK_ELF: &[u8] = include_bytes!("../../../elf/range-elf");

/// The arguments for the host executable.
#[derive(Debug, Clone, Parser)]
struct HostArgs {
    /// The start block of the range to execute.
    #[clap(long)]
    start: u64,
    /// The end block of the range to execute.
    #[clap(long)]
    end: u64,
    /// The number of blocks to execute in a single batch.
    #[clap(long)]
    batch_size: Option<u64>,
    /// Whether to generate a proof or just execute the block.
    #[clap(long)]
    prove: bool,
    /// The path to the CSV file containing the execution data.
    #[clap(long, default_value = "report.csv")]
    report_path: PathBuf,
    /// Use cached witness generation.
    #[clap(long)]
    use_cache: bool,
    /// The environment file to use.
    #[clap(long, default_value = ".env")]
    env_file: PathBuf,
    /// Configuration flag for whether to just grab the block data and not execute the blocks.
    #[clap(long)]
    fast: bool,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
struct SpanBatchRange {
    start: u64,
    end: u64,
}

fn get_max_span_batch_range_size(l2_chain_id: u64, supplied_range_size: Option<u64>) -> u64 {
    // TODO: The default size/batch size should be dynamic based on the L2 chain. Specifically, look at the gas used across the block range (should be fast to compute) and then set the batch size accordingly.
    if let Some(supplied_range_size) = supplied_range_size {
        return supplied_range_size;
    }

    const DEFAULT_SIZE: u64 = 300;
    match l2_chain_id {
        8453 => 5,      // Base
        11155420 => 30, // OP Sepolia
        10 => 10,       // OP Mainnet
        _ => DEFAULT_SIZE,
    }
}

/// Split a range of blocks into a list of span batch ranges.
fn split_range(
    start: u64,
    end: u64,
    l2_chain_id: u64,
    supplied_range_size: Option<u64>,
) -> Vec<SpanBatchRange> {
    let mut ranges = Vec::new();
    let mut current_start = start;
    let max_size = get_max_span_batch_range_size(l2_chain_id, supplied_range_size);

    while current_start < end {
        let current_end = min(current_start + max_size, end);
        ranges.push(SpanBatchRange {
            start: current_start,
            end: current_end,
        });
        current_start = current_end + 1;
    }

    ranges
}

/// Concurrently run the native data generation process for each split range.
async fn run_native_data_generation(host_clis: &[HostCli]) {
    const CONCURRENT_NATIVE_HOST_RUNNERS: usize = 5;

    // Split the entire range into chunks of size CONCURRENT_NATIVE_HOST_RUNNERS and process chunks
    // serially. Generate witnesses within each chunk in parallel. This prevents the RPC from
    // being overloaded with too many concurrent requests, while also improving witness generation
    // throughput.
    for chunk in host_clis.chunks(CONCURRENT_NATIVE_HOST_RUNNERS) {
        let mut witnessgen_executor = WitnessGenExecutor::default();

        for host_cli in chunk {
            block_on(witnessgen_executor.spawn_witnessgen(host_cli))
                .expect("Failed to spawn witness generation process");
        }

        block_on(witnessgen_executor.flush()).expect("Failed to generate witnesses");
    }
}

/// Utility method for blocking on an async function.
///
/// If we're already in a tokio runtime, we'll block in place. Otherwise, we'll create a new
/// runtime.
pub fn block_on<T>(fut: impl Future<Output = T>) -> T {
    // Handle case if we're already in an tokio runtime.
    if let Ok(handle) = tokio::runtime::Handle::try_current() {
        block_in_place(|| handle.block_on(fut))
    } else {
        // Otherwise create a new runtime.
        let rt = tokio::runtime::Runtime::new().expect("Failed to create a new runtime");
        rt.block_on(fut)
    }
}

/// Run the zkVM execution process for each split range in parallel.
async fn execute_blocks_parallel(
    host_clis: &[HostCli],
    ranges: Vec<SpanBatchRange>,
    prover: &ProverClient,
    l2_chain_id: u64,
    args: &HostArgs,
) -> Vec<ExecutionStats> {
    // Create a new execution stats map between the start and end block and the default ExecutionStats.
    let mut execution_stats_map = HashMap::new();

    // Fetch all of the execution stats block ranges in parallel.
    let exec_stats = futures::stream::iter(ranges.clone())
        .map(|range| async move {
            // Create a new data fetcher. This avoids the runtime dropping the provider dispatch task.
            let data_fetcher = OPSuccinctDataFetcher::default();
            let mut exec_stats = ExecutionStats::default();
            exec_stats
                .add_block_data(&data_fetcher, range.start, range.end)
                .await;
            ((range.start, range.end), exec_stats)
        })
        .buffered(15)
        .collect::<Vec<_>>()
        .await;

    for (range, stats) in exec_stats {
        execution_stats_map.insert(range, stats);
    }

    let cargo_metadata = cargo_metadata::MetadataCommand::new().exec().unwrap();
    let root_dir = PathBuf::from(cargo_metadata.workspace_root);
    let report_path = root_dir.join(format!(
        "execution-reports/{}/{}-{}-report.csv",
        l2_chain_id, args.start, args.end
    ));
    // Create the parent directory if it doesn't exist
    if let Some(parent) = report_path.parent() {
        if !parent.exists() {
            fs::create_dir_all(parent).unwrap();
        }
    }

    // Create an empty file since canonicalize requires the path to exist
    fs::File::create(&report_path).unwrap();
    let report_path = report_path.canonicalize().unwrap();

    // Run the zkVM execution process for each split range in parallel and fill in the execution stats.
    host_clis
        .par_iter()
        .zip(ranges.par_iter())
        .for_each(|(host_cli, range)| {
            let sp1_stdin = get_proof_stdin(&host_cli).unwrap();

            // TODO: Implement retries with a smaller block range if this fails.
            let (_, report) = prover
                .execute(MULTI_BLOCK_ELF, sp1_stdin)
                .run()
                .unwrap_or_else(|e| {
                    panic!(
                        "Failed to execute blocks {:?} - {:?}: {:?}",
                        range.start, range.end, e
                    )
                });

            // Get the existing execution stats and modify it in place.
            let mut exec_stats = execution_stats_map
                .get(&(range.start, range.end))
                .unwrap()
                .clone();
            exec_stats.add_report_data(&report);
            exec_stats.add_aggregate_data();

            let mut file = OpenOptions::new()
                .read(true)
                .write(true)
                .append(true)
                .open(&report_path)
                .unwrap();

            // Writes the headers only if the file is empty.
            let needs_header = file.seek(std::io::SeekFrom::End(0)).unwrap() == 0;

            let mut csv_writer = csv::WriterBuilder::new()
                .has_headers(needs_header)
                .from_writer(file);

            csv_writer
                .serialize(exec_stats.clone())
                .expect("Failed to write execution stats to CSV.");
            csv_writer.flush().expect("Failed to flush CSV writer.");
        });

    info!("Execution is complete.");

    let execution_stats = execution_stats_map.iter().map(|(_, v)| v.clone()).collect();
    drop(execution_stats_map);
    execution_stats
}

/// Write the block data to a CSV file.
fn write_block_data_to_csv(
    block_data: &[BlockInfo],
    l2_chain_id: u64,
    args: &HostArgs,
) -> Result<()> {
    let report_path = PathBuf::from(format!(
        "block-data/{}/{}-{}-block-data.csv",
        l2_chain_id, args.start, args.end
    ));
    if let Some(parent) = report_path.parent() {
        fs::create_dir_all(parent)?;
    }

    let mut csv_writer = csv::Writer::from_path(report_path)?;

    for block in block_data {
        csv_writer
            .serialize(block)
            .expect("Failed to write execution stats to CSV.");
    }
    csv_writer.flush().expect("Failed to flush CSV writer.");

    Ok(())
}

/// Write the execution stats to a CSV file.
fn write_execution_stats_to_csv(
    execution_stats: &[ExecutionStats],
    l2_chain_id: u64,
    args: &HostArgs,
) -> Result<()> {
    let report_path = PathBuf::from(format!(
        "execution-reports/{}/{}-{}-report.csv",
        l2_chain_id, args.start, args.end
    ));
    if let Some(parent) = report_path.parent() {
        fs::create_dir_all(parent)?;
    }

    let mut csv_writer = csv::Writer::from_path(report_path)?;

    for stats in execution_stats {
        csv_writer
            .serialize(stats)
            .expect("Failed to write execution stats to CSV.");
    }
    csv_writer.flush().expect("Failed to flush CSV writer.");

    Ok(())
}

/// Aggregate the execution statistics for an array of execution stats objects.
fn aggregate_execution_stats(
    execution_stats: &[ExecutionStats],
    total_execution_time_sec: u64,
    witness_generation_time_sec: u64,
) -> ExecutionStats {
    let mut aggregate_stats = ExecutionStats::default();
    let mut batch_start = u64::MAX;
    let mut batch_end = u64::MIN;
    for stats in execution_stats {
        batch_start = min(batch_start, stats.batch_start);
        batch_end = max(batch_end, stats.batch_end);

        // Accumulate most statistics across all blocks.
        aggregate_stats.total_instruction_count += stats.total_instruction_count;
        aggregate_stats.oracle_verify_instruction_count += stats.oracle_verify_instruction_count;
        aggregate_stats.derivation_instruction_count += stats.derivation_instruction_count;
        aggregate_stats.block_execution_instruction_count +=
            stats.block_execution_instruction_count;
        aggregate_stats.blob_verification_instruction_count +=
            stats.blob_verification_instruction_count;
        aggregate_stats.total_sp1_gas += stats.total_sp1_gas;
        aggregate_stats.nb_blocks += stats.nb_blocks;
        aggregate_stats.nb_transactions += stats.nb_transactions;
        aggregate_stats.eth_gas_used += stats.eth_gas_used;
        aggregate_stats.bn_pair_cycles += stats.bn_pair_cycles;
        aggregate_stats.bn_add_cycles += stats.bn_add_cycles;
        aggregate_stats.bn_mul_cycles += stats.bn_mul_cycles;
        aggregate_stats.kzg_eval_cycles += stats.kzg_eval_cycles;
        aggregate_stats.ec_recover_cycles += stats.ec_recover_cycles;
    }

    // For statistics that are per-block or per-transaction, we take the average over the entire
    // range.
    aggregate_stats.cycles_per_block =
        aggregate_stats.total_instruction_count / aggregate_stats.nb_blocks;
    aggregate_stats.cycles_per_transaction =
        aggregate_stats.total_instruction_count / aggregate_stats.nb_transactions;
    aggregate_stats.transactions_per_block =
        aggregate_stats.nb_transactions / aggregate_stats.nb_blocks;
    aggregate_stats.gas_used_per_block = aggregate_stats.eth_gas_used / aggregate_stats.nb_blocks;
    aggregate_stats.gas_used_per_transaction =
        aggregate_stats.eth_gas_used / aggregate_stats.nb_transactions;

    // Use the earliest start and latest end across all blocks.
    aggregate_stats.batch_start = batch_start;
    aggregate_stats.batch_end = batch_end;

    // Set the total execution time to the total execution time of the entire range.
    aggregate_stats.total_execution_time_sec = total_execution_time_sec;
    aggregate_stats.witness_generation_time_sec = witness_generation_time_sec;

    aggregate_stats
}

#[tokio::main]
async fn main() -> Result<()> {
    let args = HostArgs::parse();

    dotenv::from_path(&args.env_file).ok();
    utils::setup_logger();

    let data_fetcher = OPSuccinctDataFetcher::default();

    let l2_chain_id = data_fetcher.get_l2_chain_id().await?;

    // If we want to execute fast, just get the block data and return.
    if args.fast {
        let l2_block_data = data_fetcher
            .get_l2_block_data_range(args.start, args.end)
            .await?;
        write_block_data_to_csv(&l2_block_data, l2_chain_id, &args)?;
        return Ok(());
    }

    let split_ranges = split_range(args.start, args.end, l2_chain_id, args.batch_size);

    info!(
        "The span batch ranges which will be executed: {:?}",
        split_ranges
    );

    let prover = ProverClient::new();

    let cache_mode = if args.use_cache {
        CacheMode::KeepCache
    } else {
        CacheMode::DeleteCache
    };

    // Get the host CLIs in order, in parallel.
    let host_clis = futures::stream::iter(split_ranges.iter())
        .map(|range| async {
            data_fetcher
                .get_host_cli_args(range.start, range.end, ProgramType::Multi, cache_mode)
                .await
                .expect("Failed to get host CLI args")
        })
        .buffered(15)
        .collect::<Vec<_>>()
        .await;

    let start_time = Instant::now();
    if !args.use_cache {
        // Get the host CLI args
        run_native_data_generation(&host_clis).await;
    }
    let witness_generation_time_sec = start_time.elapsed().as_secs();

    let start_time = Instant::now();
    let execution_stats =
        execute_blocks_parallel(&host_clis, split_ranges, &prover, l2_chain_id, &args).await;
    let total_execution_time_sec = start_time.elapsed().as_secs();

    let aggregate_execution_stats = aggregate_execution_stats(
        &execution_stats,
        total_execution_time_sec,
        witness_generation_time_sec,
    );

    // Read the execution stats from the CSV file.
    let cargo_metadata = cargo_metadata::MetadataCommand::new().exec().unwrap();
    let root_dir = PathBuf::from(cargo_metadata.workspace_root);
    let report_path = root_dir.join(format!(
        "execution-reports/{}/{}-{}-report.csv",
        l2_chain_id, args.start, args.end
    ));

    let mut final_execution_stats = Vec::new();
    let mut csv_reader = csv::Reader::from_path(report_path)?;
    for result in csv_reader.deserialize() {
        let stats: ExecutionStats = result?;
        final_execution_stats.push(stats);
    }

    println!(
        "Aggregate Execution Stats for Chain {}: \n {}",
        l2_chain_id, aggregate_execution_stats
    );

    Ok(())
}
