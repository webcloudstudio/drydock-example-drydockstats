# FEATURE: Daily Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the per-start-date report with calls, failures, combined tokens, and minutes, ending in a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportSummary.md, FEATURE-ReportFormat.md, FEATURE-LogNumbers.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliReadErrors.md |
| Provides    | daily report, dispatch entry `daily` |
| Consumes    | internal/log Normalize, internal/log Record, internal/report FormatMinutes, internal/report CSV writer |

## Purpose

Emit the `daily` report: one row per start date among kept records, plus a TOTAL row.

## Behavior

- Header: `date,calls,failed,tokens,minutes`.
- Groups by the first 10 characters of `started_at`, ordered ascending by byte value.
- `calls` is the group record count; `failed` counts records with `status != "succeeded"`.
- `tokens` is the group sum of input plus cached plus output tokens.
- `minutes` is `FormatMinutes` of summed `elapsed_ms`.
- A final `TOTAL` row covers all kept records in the same columns.
- Empty match writes the header and a zero TOTAL row (`TOTAL,0,0,0,0.0`).
- Dispatch: `daily` is registered in the `cli-parse` dispatch table.

## Programmatic Acceptance

=== AC daily-total-row ===
Intent: The daily report ends with a TOTAL row whose calls column equals the sum of the per-date call counts, read back as parsed fields.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "daily", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
assert rows[0][0] == "date", rows[0]
total = rows[-1]
assert total[0] == "TOTAL", total
group_calls = sum(int(r[1]) for r in rows[1:-1])
assert int(total[1]) == group_calls, (total, group_calls)
=== END AC daily-total-row ===

=== AC daily-order ===
Intent: Daily rows are ordered by start date ascending, read back as parsed fields.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "daily", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))[1:-1]
dates = [r[0] for r in rows]
assert dates == sorted(dates), dates
=== END AC daily-order ===

## User Acceptance

- None.

## Guardrails

- The date key is always the first 10 characters of `started_at`; no timezone conversion is applied.
- The TOTAL row covers filtered records only.
