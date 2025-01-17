use anyhow::Result;
use clap::Parser;
use op_succinct_host_utils::{
    block_range::get_validated_block_range,
    fetcher::{CacheMode, OPSuccinctDataFetcher, RunContext},
    get_proof_stdin,
    stats::ExecutionStats,
    ProgramType,
};
use op_succinct_prove::{execute_multi, generate_witness, DEFAULT_RANGE, RANGE_ELF};
use op_succinct_scripts::HostExecutorArgs;
use sp1_sdk::{utils, ProverClient};
use std::{fs, time::Duration};

/// Execute the OP Succinct program for multiple blocks.
#[tokio::main]
async fn main() -> Result<()> {
    // let args = HostExecutorArgs::parse();

    // dotenv::from_path(&args.env_file)?;
    // utils::setup_logger();

    // let data_fetcher = OPSuccinctDataFetcher::new_with_rollup_config(RunContext::Dev).await?;

    // let cache_mode = if args.use_cache {
    //     CacheMode::KeepCache
    // } else {
    //     CacheMode::DeleteCache
    // };

    // // If the end block is provided, check that it is less than the latest finalized block. If the end block is not provided, use the latest finalized block.
    // let (l2_start_block, l2_end_block) =
    //     get_validated_block_range(&data_fetcher, args.start, args.end, DEFAULT_RANGE).await?;

    // let host_cli = data_fetcher
    //     .get_host_cli_args(l2_start_block, l2_end_block, ProgramType::Multi, cache_mode)
    //     .await?;

    // // By default, re-run the native execution unless the user passes `--use-cache`.
    // let witness_generation_time_sec = if !args.use_cache {
    //     generate_witness(&host_cli).await?
    // } else {
    //     Duration::ZERO
    // };

    // Get the stdin for the block.
    let sp1_stdin = {
        let stdin_path = "stdin.bin";
        let stdin_bytes = std::fs::read(stdin_path)?;
        bincode::deserialize(&stdin_bytes)?
    };

    let prover = ProverClient::from_env();

    // If the prove flag is set, generate a proof.
    let (pk, _) = prover.setup(RANGE_ELF);

    // Generate proofs in compressed mode for aggregation verification.
    let proof = prover.prove(&pk, &sp1_stdin).compressed().run().unwrap();

    // Create a proof directory for the chain ID if it doesn't exist.
    let proof_dir = format!("data/proofs");
    if !std::path::Path::new(&proof_dir).exists() {
        fs::create_dir_all(&proof_dir).unwrap();
    }
    // Save the proof to the proof directory corresponding to the chain ID.
    proof
        .save(format!("{}/proof.bin", proof_dir))
        .expect("saving proof failed");

    Ok(())
}
