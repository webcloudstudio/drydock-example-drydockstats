# FEATURE: Report Summary

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The summary report: one CSV row of totals over the filtered records plus the whole-file skipped count. |
| Depends On  | FEATURE-CliExit.md, FEATURE-ValueFormat.md |
| Provides    | ddkstat summary |
| Consumes    | cli.Register, format.Minutes, format.Cost, logread.Record |

## Purpose

`ddkstat summary [options] <log>` answers how many calls ran, how many were skipped, and what they used and cost.

## Behavior

- Header: `records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`.
- One data row over the filtered records, `\n` line endings.
- `records` is the filtered record count; `succeeded` and `failed` split it by the failed rule; token columns are sums; `cost_usd` uses `format.Cost`; `minutes` uses `format.Minutes` of the summed `elapsed_ms`.
- `skipped` is the whole-file skipped-line count and ignores every filter.
- `first` and `last` are the earliest and latest start dates of the filtered records, empty when none match.
- Aggregation (`SummaryAggregate`) and writing (`WriteSummary(w io.Writer, records []logread.Record, skipped int) error`) are separate functions in `internal/report`; `internal/cli` registers the writer as `summary` with `cli.Register`.

## Programmatic Acceptance

=== AC summary-aggregate ===
Intent: SummaryAggregate sums counts, tokens, cost and duration over records built by the test.

from drydock_go import run_test

result = run_test("./internal/report", "TestSummaryAggregate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summary-aggregate ===

=== AC summary-empty-first-last ===
Intent: With no matching record, first and last are empty and counters are zero.

from drydock_go import run_test

result = run_test("./internal/report", "TestSummaryEmptyFirstLast")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summary-empty-first-last ===

=== AC summary-skipped-ignores-filters ===
Intent: The skipped count stays the whole-file count regardless of target and date filters.

from drydock_go import run_test

result = run_test("./internal/report", "TestSummarySkippedIgnoresFilters")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summary-skipped-ignores-filters ===

=== AC summary-write ===
Intent: WriteSummary writes the declared header and one row with the declared column order.

from drydock_go import run_test

result = run_test("./internal/report", "TestSummaryWrite")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summary-write ===

=== AC summary-run-via-cli ===
Intent: cli.Run dispatches the summary report over a test-written log and exits zero with a parseable CSV.

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunSummaryReport")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summary-run-via-cli ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no database; the log is never written.
- No report value is hardcoded; every number comes from the log being read.
- Skipped-line warnings go to standard error, never into the CSV.
