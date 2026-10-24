# FEATURE: Summary Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the one-row summary CSV report over filtered records, with a whole-file skipped count. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md, FEATURE-LogRequiredFields.md, FEATURE-LogDuplicates.md, FEATURE-LogModel.md, FEATURE-LogNumeric.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-ReportFormat.md |
| Provides    | summary report |
| Consumes    | records, skipped, FormatMinutes, FormatCost, WriteCSV |

## Purpose

Write the `summary` report: one header row followed by exactly one data row.

## Report Contract

Header:

```
records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
```

- `records`: count of filtered records.
- `skipped`: count of skipped lines in the whole file, regardless of `--target` or date filters.
- `succeeded`: filtered records whose `status` is `succeeded`.
- `failed`: filtered records whose `status` is not `succeeded`.
- `first` / `last`: earliest and latest start date (first 10 characters of `started_at`) among filtered records; empty when no record matches.
- `input_tokens`, `cached_tokens`, `output_tokens`: sums over filtered records, absent values counted as zero.
- `cost_usd`: sum over filtered records, written by `FormatCost`.
- `minutes`: sum of `elapsed_ms` over filtered records, written by `FormatMinutes`.

Output is written with `WriteCSV`. Diagnostics go to standard error only.

## Programmatic Acceptance

=== AC summary-empty-file ===
Intent: A log with no records yields one data row with zero counts and empty first and last.
Suite: scoped

import json
import os
import subprocess
import sys
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "empty.jsonl")
    open(path, "w", encoding="utf-8").close()
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (code, out, err)
    import csv, io
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[0] == ["records", "skipped", "succeeded", "failed", "first", "last",
                       "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"]
    assert len(rows) == 2
    assert rows[1][0] == "0"
    assert rows[1][4] == ""
    assert rows[1][5] == ""
=== END AC summary-empty-file ===

=== AC summary-filters-and-skipped ===
Intent: Filtered counts exclude non-matching records while skipped counts the whole file; the first and last dates are the earliest and latest filtered start dates.
Suite: scoped

import csv
import io
import os
import sys
import tempfile

from drydock_go import cli_runner

lines = [
    '{"schema_version":1,"execution_id":"a","started_at":"2026-10-14T01:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Alpha","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":10,"cached_input_tokens":5,"output_tokens":2,"cost_usd":0.5,"elapsed_ms":60000}}}',
    '{"schema_version":1,"execution_id":"b","started_at":"2026-10-16T01:00:00.000Z","status":"failed","job":{"command_name":"build","target":"Beta","llm":"claude","model":"haiku"},"result":{"returncode":1,"timed_out":false,"stats":{"input_tokens":3,"cached_input_tokens":0,"output_tokens":1,"cost_usd":0.25,"elapsed_ms":0}}}',
    "not json",
]

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "log.jsonl")
    with open(path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")
    code, out, err = cli_runner("summary", "--target", "Alpha", path)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    row = rows[1]
    assert row[0] == "1"
    assert row[1] == "1"
    assert row[2] == "1"
    assert row[3] == "0"
    assert row[4] == "2026-10-14"
    assert row[5] == "2026-10-14"
    assert row[6] == "10"
    assert row[7] == "5"
    assert row[8] == "2"
=== END AC summary-filters-and-skipped ===

## User Acceptance

- None.

## Guardrails

- The summary report must not read `sources/run.sh` or embed any expected answer.
- `skipped` is never affected by `--target`, `--since`, or `--until`.
