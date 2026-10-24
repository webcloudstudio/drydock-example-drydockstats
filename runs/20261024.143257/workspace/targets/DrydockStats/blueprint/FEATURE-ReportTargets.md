# FEATURE: Targets Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the per-Target report with calls, failures, tokens, cost, and minutes, ending in a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportSummary.md, FEATURE-ReportFormat.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliReadErrors.md |
| Provides    | targets report, dispatch entry `targets` |
| Consumes    | internal/log Normalize, internal/log Record, internal/report FormatMinutes, internal/report FormatCost, internal/report CSV writer |

## Purpose

Emit the `targets` report: one row per `job.target` among kept records, plus a TOTAL row.

## Behavior

- Header: `target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`.
- Groups by `job.target`, ordered ascending by byte value.
- `calls` is the group record count; `failed` counts records with `status != "succeeded"`.
- Token columns are sums; `cost_usd` is `FormatCost` of summed cost; `minutes` is `FormatMinutes` of summed `elapsed_ms`.
- A final `TOTAL` row covers all kept records in the same columns.
- Empty match writes the header and a zero TOTAL row (`0.0000` cost, `0.0` minutes).
- Dispatch: `targets` is registered in the `cli-parse` dispatch table.

## Programmatic Acceptance

=== AC targets-total-row ===
Intent: The targets report ends with a TOTAL row whose calls column equals the sum of the per-target call counts, read back through a CSV parser.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "targets", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
assert rows[0][0] == "target", rows[0]
total = rows[-1]
assert total[0] == "TOTAL", total
group_calls = sum(int(r[1]) for r in rows[1:-1])
assert int(total[1]) == group_calls, (total, group_calls)
=== END AC targets-total-row ===

=== AC targets-order ===
Intent: Target rows are ordered ascending by byte value, read back as parsed fields.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "targets", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))[1:-1]
names = [r[0] for r in rows]
assert names == sorted(names), names
=== END AC targets-order ===

## User Acceptance

- None.

## Guardrails

- Target names are compared exactly as they appear in `job.target`.
- The TOTAL row covers filtered records only.
