# FEATURE: Summary Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the one-row summary report over the filtered log, with a whole-file skipped count. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportFormat.md, FEATURE-LogSkip.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliReadErrors.md, FEATURE-CliParse.md |
| Provides    | summary report, dispatch entry `summary` |
| Consumes    | internal/log Normalize, internal/log Record, internal/report FormatMinutes, internal/report FormatCost, internal/report CSV writer |

## Purpose

Emit the `summary` report: a single data row describing every record that survives the active filters, plus the count of skipped lines for the whole file.

## Behavior

- Header: `records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`.
- `records` is the count of kept records after filters.
- `skipped` is the whole-file skipped count from log-skip and ignores `--target` and date filters.
- `succeeded` counts kept records with `status == "succeeded"`; `failed` counts the rest.
- `first` and `last` are the earliest and latest start dates (first 10 characters of `started_at`) among kept records, empty when no record matches.
- `input_tokens`, `cached_tokens`, `output_tokens` are sums over kept records, absent values counted as zero.
- `cost_usd` is the summed cost formatted with `FormatCost` (`%.4f`).
- `minutes` is the summed `elapsed_ms` formatted with `FormatMinutes`.
- The row ends with `\n`; no TOTAL row is written.
- Dispatch: `summary` is registered in the `cli-parse` dispatch table.

## Programmatic Acceptance

=== AC summary-header-and-row ===
Intent: The summary report over the sample log writes the declared header and exactly one data row whose fields are read back by a CSV parser.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "summary", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
header = ",".join(rows[0])
assert header == "records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes", header
assert len(rows) == 2, len(rows)
=== END AC summary-header-and-row ===

=== AC summary-empty-match-dates ===
Intent: When no record matches, the summary row has empty first and last fields and zero counts, while skipped still reports the whole file.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "summary", "--target", "NoSuchTarget", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
row = rows[1]
assert row[0] == "0", row
assert row[1] == "3", row
assert row[4] == "", row
assert row[5] == "", row
=== END AC summary-empty-match-dates ===

## User Acceptance

- None.

## Guardrails

- The summary row never filters `skipped`.
- No TOTAL row is written for the summary report.
