# Feature: Models Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the per-provider and per-model CSV report over the filtered log records. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-LogFailed.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-FormatCSV.md, FEATURE-FormatCost.md |
| Provides    | models report |
| Consumes    | model attribution, record aggregates, CSV and format helpers |

## Purpose

Write one CSV row per provider and model pair over the records that pass the active `--target`, `--since`, and `--until` filters. Each row reports calls, failed calls, summed input, cached, and output tokens, and summed cost. The model of a record comes from `result.stats.model` when present and non-empty, otherwise `job.model`, and the provider comes from `job.llm`.

## Output

Header row, in this exact order:

```
provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
```

- One data row per distinct `(provider, model)` pair, ordered by provider ascending, then model ascending, by byte value.
- `calls` counts filtered records in the pair; `failed` counts those with a status other than `succeeded`.
- Token columns are integer sums; absent or null values count as zero.
- `cost_usd` is the summed cost written with `%.4f`.
- A final `TOTAL` row over all filtered records leaves the provider and model cells empty: `TOTAL,,…`.
- A report with no matching records writes the header and a `TOTAL,,0,0,0,0,0,0.0000` row.

## Acceptance Source

- `sources/run.sh` A4, B5, and C1-style zero-match behavior for this report.
- `sources/INSTRUCTIONS.md` Output table row `models`.

## Programmatic Acceptance

=== AC models-report-whole-sample ===
Intent: The models report over the full sample matches the supplied known answer for every pair and the TOTAL row, parsed as CSV.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os
import subprocess
import tempfile

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("models", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"], rows[0]
assert rows[-1][0] == "TOTAL" and rows[-1][1] == "", rows[-1]
assert rows[-1][2] == "33" and rows[-1][3] == "2", rows[-1]
=== END AC models-report-whole-sample ===

=== AC models-report-target-filter ===
Intent: With --target ReadingList the models report contains only the ReadingList claude-haiku-5-5 pair and its TOTAL row, parsed as CSV.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("models", "--target", "ReadingList", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0][0:2] == ["provider", "model"], rows[0]
assert len(rows) == 3, rows
assert rows[1][0:2] == ["claude", "claude-haiku-5-5"], rows[1]
assert rows[2][0] == "TOTAL" and rows[2][2] == rows[1][2], rows
=== END AC models-report-target-filter ===

=== AC models-report-zero-match ===
Intent: A filter matching no record writes the header and a zero TOTAL row with exit 0.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("models", "--target", "NoSuchTarget", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1] == ["TOTAL", "", "0", "0", "0", "0", "0", "0.0000"], rows[1]
=== END AC models-report-zero-match ===

## User Acceptance

- None.

## Guardrails

- Provider and model cells are never reconstructed from another field; provider comes from `job.llm` only.
- The TOTAL row is the only row whose provider and model cells are empty.
- Values are computed from the log read, never from a hardcoded row or count.
