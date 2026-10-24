# FEATURE: Commands Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the per-command report with calls, failures, tokens, and minutes, ending in a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportSummary.md, FEATURE-ReportFormat.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliReadErrors.md |
| Provides    | commands report, dispatch entry `commands` |
| Consumes    | internal/log Normalize, internal/log Record, internal/report FormatMinutes, internal/report CSV writer |

## Purpose

Emit the `commands` report: one row per `job.command_name` among kept records, plus a TOTAL row.

## Behavior

- Header: `command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes`.
- Groups by `job.command_name`, ordered by `calls` descending, then command ascending by byte value.
- `calls` is the group record count; `failed` counts records with `status != "succeeded"`.
- Token columns are sums; `minutes` is `FormatMinutes` of summed `elapsed_ms`.
- A final `TOTAL` row covers all kept records, written with the same columns.
- Empty match writes the header and a `TOTAL,0,0,0,0,0,0.0` row.
- Dispatch: `commands` is registered in the `cli-parse` dispatch table.

## Programmatic Acceptance

=== AC commands-total-row ===
Intent: The commands report ends with a TOTAL row whose calls column equals the sum of the per-command call counts, read back through a CSV parser.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "commands", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
assert rows[0][0] == "command", rows[0]
total = rows[-1]
assert total[0] == "TOTAL", total
group_calls = sum(int(r[1]) for r in rows[1:-1])
assert int(total[1]) == group_calls, (total, group_calls)
=== END AC commands-total-row ===

=== AC commands-order ===
Intent: Command rows are ordered by calls descending, then command ascending, read back as parsed fields.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "commands", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))[1:-1]
keys = [(-int(r[1]), r[0]) for r in rows]
assert keys == sorted(keys), keys
=== END AC commands-order ===

## User Acceptance

- None.

## Guardrails

- Groups use `job.command_name` exactly; no command names are normalized or hardcoded.
- The TOTAL row covers filtered records only.
