use std::fmt;

use crate::fetcher::{ChainMode, SP1KonaDataFetcher};
use num_format::{Locale, ToFormattedString};
use sp1_sdk::ExecutionReport;

#[derive(Debug)]
pub struct BnStats {
    pub bn_pair_cycles: u64,
    pub bn_add_cycles: u64,
    pub bn_mul_cycles: u64,
}

/// Statistics for the multi-block execution.
#[derive(Debug)]
pub struct ExecutionStats {
    pub total_instruction_count: u64,
    pub block_execution_instruction_count: u64,
    pub nb_blocks: u64,
    pub nb_transactions: u64,
    pub total_gas_used: u64,
    pub bn_stats: BnStats,
}

/// Write a statistic to the formatter.
fn write_stat(f: &mut fmt::Formatter<'_>, label: &str, value: u64) -> fmt::Result {
    writeln!(
        f,
        "| {:<30} | {:>25} |",
        label,
        value.to_formatted_string(&Locale::en)
    )
}

impl fmt::Display for ExecutionStats {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let cycles_per_block = self.block_execution_instruction_count / self.nb_blocks;
        let cycles_per_transaction = self.block_execution_instruction_count / self.nb_transactions;
        let transactions_per_block = self.nb_transactions / self.nb_blocks;
        let gas_used_per_block = self.total_gas_used / self.nb_blocks;
        let gas_used_per_transaction = self.total_gas_used / self.nb_transactions;

        writeln!(
            f,
            "+--------------------------------+---------------------------+"
        )?;
        writeln!(f, "| {:<30} | {:<25} |", "Metric", "Value")?;
        writeln!(
            f,
            "+--------------------------------+---------------------------+"
        )?;
        write_stat(f, "Total Cycles", self.total_instruction_count)?;
        write_stat(
            f,
            "Block Execution Cycles",
            self.block_execution_instruction_count,
        )?;
        // Only write the BN stats if they're non-zero.
        if self.bn_stats.bn_pair_cycles > 0 {
            write_stat(f, "Bn Pair Cycles", self.bn_stats.bn_pair_cycles)?;
        }
        if self.bn_stats.bn_add_cycles > 0 {
            write_stat(f, "Bn Add Cycles", self.bn_stats.bn_add_cycles)?;
        }
        if self.bn_stats.bn_mul_cycles > 0 {
            write_stat(f, "Bn Mul Cycles", self.bn_stats.bn_mul_cycles)?;
        }
        write_stat(f, "Total Blocks", self.nb_blocks)?;
        write_stat(f, "Total Transactions", self.nb_transactions)?;
        write_stat(f, "Cycles per Block", cycles_per_block)?;
        write_stat(f, "Cycles per Transaction", cycles_per_transaction)?;
        write_stat(f, "Transactions per Block", transactions_per_block)?;
        write_stat(f, "Total Gas Used", self.total_gas_used)?;
        write_stat(f, "Gas Used per Block", gas_used_per_block)?;
        write_stat(f, "Gas Used per Transaction", gas_used_per_transaction)?;
        writeln!(
            f,
            "+--------------------------------+---------------------------+"
        )
    }
}

/// Get the execution stats for a given report.
pub async fn get_execution_stats(
    data_fetcher: &SP1KonaDataFetcher,
    start: u64,
    end: u64,
    report: &ExecutionReport,
) -> ExecutionStats {
    // Get the total instruction count for execution across all blocks.
    let block_execution_instruction_count: u64 =
        *report.cycle_tracker.get("block-execution").unwrap();

    let nb_blocks = end - start + 1;

    // Fetch the number of transactions in the blocks from the L2 RPC.
    let block_data_range = data_fetcher
        .get_block_data_range(ChainMode::L2, start, end)
        .await
        .expect("Failed to fetch block data range.");

    let nb_transactions = block_data_range.iter().map(|b| b.transaction_count).sum();
    let total_gas_used = block_data_range.iter().map(|b| b.gas_used).sum();

    let bn_stats = BnStats {
        bn_add_cycles: *report.cycle_tracker.get("precompile-bn-add").unwrap_or(&0),
        bn_mul_cycles: *report.cycle_tracker.get("precompile-bn-mul").unwrap_or(&0),
        bn_pair_cycles: *report.cycle_tracker.get("precompile-bn-pair").unwrap_or(&0),
    };

    ExecutionStats {
        total_instruction_count: report.total_instruction_count(),
        block_execution_instruction_count,
        nb_blocks,
        nb_transactions,
        total_gas_used,
        bn_stats,
    }
}

#[derive(Debug, Clone)]
pub struct SpanBatchStats {
    pub span_start: u64,
    pub span_end: u64,
    pub total_blocks: u64,
    pub total_transactions: u64,
    pub total_gas_used: u64,
    pub total_cycles: u64,
    pub total_sp1_gas: u64,
    pub cycles_per_block: u64,
    pub cycles_per_transaction: u64,
    pub gas_used_per_block: u64,
    pub gas_used_per_transaction: u64,
    pub total_derivation_cycles: u64,
    pub total_execution_cycles: u64,
    pub total_blob_verification_cycles: u64,
    pub bn_add_cycles: u64,
    pub bn_mul_cycles: u64,
    pub bn_pair_cycles: u64,
    pub kzg_eval_cycles: u64,
    pub ec_recover_cycles: u64,
}

impl fmt::Display for SpanBatchStats {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        writeln!(
            f,
            "+-------------------------------+---------------------------+"
        )?;
        writeln!(f, "| {:<30} | {:<25} |", "Metric", "Value")?;
        writeln!(
            f,
            "+-------------------------------+---------------------------+"
        )?;
        write_stat(f, "Span Start", self.span_start)?;
        write_stat(f, "Span End", self.span_end)?;
        write_stat(f, "Total Blocks", self.total_blocks)?;
        write_stat(f, "Total Transactions", self.total_transactions)?;
        write_stat(f, "Total Gas Used", self.total_gas_used)?;
        write_stat(f, "Total Cycles", self.total_cycles)?;
        write_stat(f, "Total SP1 Gas", self.total_sp1_gas)?;
        write_stat(f, "Cycles per Block", self.cycles_per_block)?;
        write_stat(f, "Cycles per Transaction", self.cycles_per_transaction)?;
        write_stat(f, "Gas Used per Block", self.gas_used_per_block)?;
        write_stat(f, "Gas Used per Transaction", self.gas_used_per_transaction)?;
        write_stat(f, "Total Derivation Cycles", self.total_derivation_cycles)?;
        write_stat(f, "Total Execution Cycles", self.total_execution_cycles)?;
        write_stat(
            f,
            "Total Blob Verification Cycles",
            self.total_blob_verification_cycles,
        )?;
        write_stat(f, "BN Add Cycles", self.bn_add_cycles)?;
        write_stat(f, "BN Mul Cycles", self.bn_mul_cycles)?;
        write_stat(f, "BN Pair Cycles", self.bn_pair_cycles)?;
        write_stat(f, "KZG Eval Cycles", self.kzg_eval_cycles)?;
        write_stat(f, "EC Recover Cycles", self.ec_recover_cycles)?;
        writeln!(
            f,
            "+-------------------------------+---------------------------+"
        )
    }
}
