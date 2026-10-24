# FEATURE: Models Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the per-provider and model report with calls, failures, tokens, and cost, ending in an empty-provider TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportSummary.md, FEATURE-ReportFormat.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliReadErrors.md |
| Provides    | models report, dispatch entry `models` |
| Consumes    | internal/log Normalize, internal/log Record, internal/report FormatCost, internal/report CSV writer |

## Purpose

Emit the `models` report: one row per provider and model pair among kept records, plus a TOTAL row.

## Behavior

- Header: `provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd`.
- Groups by `(job.llm, model)` where `model` is the Record model from log-model; ordered by provider then model ascending by byte value.
- `calls` is the group record count; `failed` counts records with `status != "succeeded"`.
- Token columns are sums; `cost_usd` is `FormatCost` of summed cost.
- The TOTAL row begins `TOTAL,,` (empty provider column) and covers all kept records.
- Empty match writes the header and a zero TOTAL row (`TOTAL,,0,0,0,0,0,0.0000`).
- Dispatch: `models` is registered in the `cli-parse` dispatch table.

## Programmatic Acceptance

=== AC models-total-empty-provider ===
Intent: The models report TOTAL row begins with an empty provider column, read back as parsed fields.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "models", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
total = rows[-1]
assert total[0] == "TOTAL", total
assert total[1] == "", total
=== END AC models-total-empty-provider ===

=== AC models-total-calls ===
Intent: The models report TOTAL calls equals the sum of the per-pair call counts, read back as parsed fields.
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess

sample = os.path.join("sources", "llm_sample.jsonl")
result = subprocess.run(
    ["./ddkstat", "models", sample],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
rows = list(csv.reader(io.StringIO(result.stdout)))
total = rows[-1]
group_calls = sum(int(r[2]) for r in rows[1:-1])
assert int(total[2]) == group_calls, (total, group_calls)
=== END AC models-total-calls ===

## User Acceptance

- None.

## Guardrails

- The model column for a record comes from the log-model derivation, never from a hardcoded alias.
- The TOTAL provider column is always empty.
