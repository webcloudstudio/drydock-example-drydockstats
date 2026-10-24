# Feature: Failures Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the failed-record CSV report over the filtered log records. |
| Depends On  | ARCHITECTURE.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-LogFailed.md, FEATURE-LogModel.md, FEATURE-FormatCSV.md |
| Provides    | failures report |
| Consumes    | failed-record accessors, model attribution, CSV helpers |

## Purpose

Write one CSV row per failed record over the records that pass the active `--target`, `--since`, and `--until` filters. A record is failed when its `status` is any value other than `succeeded`.

## Output

Header row, in this exact order:

```
started_at,target,command,model,returncode,timed_out
```

- One data row per failed record, ordered by `started_at` then `execution_id`, ascending by byte value.
- `model` follows the model attribution rule: `result.stats.model` when present and non-empty, otherwise `job.model`.
- `returncode` is empty when `result.returncode` is null or absent; otherwise the integer.
- `timed_out` is `true` or `false`.
- No TOTAL row.
- A report with no matching failed records writes the header only.

## Programmatic Acceptance

=== AC failures-report-whole-sample ===
Intent: The failures report over the full sample matches the two known failed records, parsed as CSV.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("failures", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["started_at", "target", "command", "model", "returncode", "timed_out"], rows[0]
assert len(rows) == 3, rows
assert rows[1][1:3] == ["Toml", "diagnose"], rows[1]
assert rows[2][1:3] == ["CommonMark", "diagnose"], rows[2]
assert all(r[0] != "TOTAL" for r in rows), rows
=== END AC failures-report-whole-sample ===

=== AC failures-report-target-filter ===
Intent: With --target CommonMark the failures report contains only the CommonMark failed record, with an empty returncode and timed_out true, parsed as CSV.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("failures", "--target", "CommonMark", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1][1] == "CommonMark", rows[1]
assert rows[1][4] == "", rows[1]
assert rows[1][5] == "true", rows[1]
=== END AC failures-report-target-filter ===

=== AC failures-report-zero-match ===
Intent: A filter matching no failed record writes only the header with exit 0.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("failures", "--target", "NoSuchTarget", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 1, rows
=== END AC failures-report-zero-match ===

## User Acceptance

- None.

## Guardrails

- The failures report never writes a TOTAL row.
- Values are computed from the log read, never from a hardcoded row or count.
