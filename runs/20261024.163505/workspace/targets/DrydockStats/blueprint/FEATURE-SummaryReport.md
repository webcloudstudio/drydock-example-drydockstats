# Feature: Summary Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the one-row summary report over the filtered records. |
| Depends On  | ARCHITECTURE.md, FEATURE-SkippedCount.md, FEATURE-LogFailed.md, FEATURE-LogNumbers.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-FormatCSV.md, FEATURE-FormatMinutes.md, FEATURE-FormatCost.md |
| Provides    | summary report |
| Consumes    | skipped count, record aggregates, CSV and format helpers |

## Purpose

The summary report answers the operator's first question: how many records ran, how many succeeded or failed, over what date span, and what tokens, cost, and minutes they used.

## Behavior

- Writes one data row under the header `records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`.
- `records` is the count of filtered records; `skipped` is the whole-file skipped count.
- `succeeded` and `failed` partition the filtered records by status.
- `first` and `last` are the earliest and latest start dates (first 10 characters of `started_at`) among filtered records, empty when none match.
- Token columns are the sums of the filtered records' values; absent values count as zero.
- `cost_usd` is written with `%.4f`; `minutes` is written in tenths of a minute.
- A zero-match run writes the header and one zero row (`0,<skipped>,0,0,,,0,0,0,0.0000,0.0`).
- Exit code is 0 whenever the log can be read, including when no record matches.

## Programmatic Acceptance

=== AC summary-header-and-row-shape ===
Intent: The summary report writes the exact header and one data row with eleven fields.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-summary-shape.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":10,"elapsed_ms":60000,"input_tokens":5,"output_tokens":7,"cost_usd":0.25,"model":"haiku"}}}\n')

code, out, err = cli_runner("summary", log)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"], rows[0]
assert len(rows) == 2, rows
assert len(rows[1]) == 11, rows[1]
=== END AC summary-header-and-row-shape ===

=== AC summary-values-from-fixture ===
Intent: The summary report aggregates a known fixture into the expected counts, dates, sums, cost, and minutes.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-summary-values.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":10,"elapsed_ms":60000,"input_tokens":5,"output_tokens":7,"cost_usd":0.25,"model":"haiku"}}}\n')
    f.write('{"schema_version":1,"execution_id":"e2","started_at":"2026-10-16T06:37:54.176Z","status":"failed","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":1,"timed_out":false,"stats":{"cached_input_tokens":20,"elapsed_ms":35000,"input_tokens":6,"output_tokens":8,"cost_usd":0.5,"model":"haiku"}}}\n')

expected_records = "2"
expected_skipped = "0"
expected_succeeded = "1"
expected_failed = "1"
expected_first = "2026-10-14"
expected_last = "2026-10-16"

code, out, err = cli_runner("summary", log)
print(err)
assert code == 0, (code, out, err)
row = next(csv.reader(io.StringIO(out)))
rows = list(csv.reader(io.StringIO(out)))
assert rows[1][0] == expected_records, rows[1]
assert rows[1][1] == expected_skipped, rows[1]
assert rows[1][2] == expected_succeeded, rows[1]
assert rows[1][3] == expected_failed, rows[1]
assert rows[1][4] == expected_first, rows[1]
assert rows[1][5] == expected_last, rows[1]
assert rows[1][6] == "11", rows[1]
assert rows[1][7] == "30", rows[1]
assert rows[1][8] == "15", rows[1]
assert rows[1][9] == "0.7500", rows[1]
assert rows[1][10] == "1.6", rows[1]
=== END AC summary-values-from-fixture ===

## User Acceptance

- None.

## Guardrails

- The summary report is one row only; it never carries a TOTAL row.
- Filters narrow `records`, `succeeded`, `failed`, tokens, cost, and minutes, but never `skipped`.
