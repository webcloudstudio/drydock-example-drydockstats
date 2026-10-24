# FEATURE: Report Empty Match

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Assemble the empty-match behavior across all six reports: each writes its header and, where it has one, a zero TOTAL row. |
| Depends On  | FEATURE-ReportSummary.md, FEATURE-ReportCommands.md, FEATURE-ReportTargets.md, FEATURE-ReportModels.md, FEATURE-ReportDaily.md, FEATURE-ReportFailures.md |
| Provides    | empty-match behavior across all six reports |

## Purpose

When a filter or an empty log leaves no matching record, every report still writes its header. Reports with a TOTAL row write that row with zeros. The summary row has empty `first` and `last`. The failures report writes only its header.

## Behavior

- A report run with a filter that matches nothing exits `0`.
- `commands`, `targets`, `models`, and `daily` write a TOTAL row of zeros: `0.0` minutes and `0.0000` cost where those columns exist.
- `summary` writes one row with `records` zero and empty `first` and `last`. `skipped` still counts the whole file.
- `failures` writes only its header.

## Programmatic Acceptance

=== AC REM-001 ===
Intent: The summary report with a filter that matches nothing writes one row with zero records, empty first and last, and a skipped count over the whole file.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


def record(execution_id, started_at, status, command, target, model, returncode, timed_out):
    return {
        "schema_version": 1,
        "execution_id": execution_id,
        "started_at": started_at,
        "status": status,
        "job": {"command_name": command, "target": target, "llm": "claude", "model": model},
        "result": {
            "returncode": returncode,
            "timed_out": timed_out,
            "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                      "output_tokens": 1, "cost_usd": 0.01, "model": model},
        },
    }


NO_TARGET = "NoSuchTarget"
MALFORMED = ["not json"]
records = [record("e-run", "2026-10-14T02:38:41.904Z", "succeeded", "build", "Toml", "gemini-3-flash", 0, False)]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
        for line in MALFORMED:
            print(line, file=handle)
    code, out, err = cli_runner("summary", "--target", NO_TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert len(rows[0]) == 11, rows[0]
row = rows[1]
assert int(row[0]) == 0, row
assert int(row[1]) == len(MALFORMED), row
assert int(row[2]) == 0 and int(row[3]) == 0, row
assert row[4] == "" and row[5] == "", row
assert int(row[6]) == 0 and int(row[7]) == 0 and int(row[8]) == 0, row
assert float(row[9]) == 0 and float(row[10]) == 0, row
=== END AC REM-001 ===

=== AC REM-002 ===
Intent: The commands report with a filter that matches nothing writes its header and a zero TOTAL row.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


def record(execution_id, started_at, status, command, target, model, returncode, timed_out):
    return {
        "schema_version": 1,
        "execution_id": execution_id,
        "started_at": started_at,
        "status": status,
        "job": {"command_name": command, "target": target, "llm": "claude", "model": model},
        "result": {
            "returncode": returncode,
            "timed_out": timed_out,
            "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                      "output_tokens": 1, "cost_usd": 0.01, "model": model},
        },
    }


NO_TARGET = "NoSuchTarget"
records = [record("e-run", "2026-10-14T02:38:41.904Z", "succeeded", "build", "Toml", "gemini-3-flash", 0, False)]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("commands", "--target", NO_TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert len(rows[0]) == 7, rows[0]
total = rows[1]
assert total[0] == "TOTAL", total
assert all(int(v) == 0 for v in total[1:6]), total
assert float(total[6]) == 0, total
=== END AC REM-002 ===

=== AC REM-003 ===
Intent: The targets report with a filter that matches nothing writes its header and a zero TOTAL row.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


def record(execution_id, started_at, status, command, target, model, returncode, timed_out):
    return {
        "schema_version": 1,
        "execution_id": execution_id,
        "started_at": started_at,
        "status": status,
        "job": {"command_name": command, "target": target, "llm": "claude", "model": model},
        "result": {
            "returncode": returncode,
            "timed_out": timed_out,
            "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                      "output_tokens": 1, "cost_usd": 0.01, "model": model},
        },
    }


NO_TARGET = "NoSuchTarget"
records = [record("e-run", "2026-10-14T02:38:41.904Z", "succeeded", "build", "Toml", "gemini-3-flash", 0, False)]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("targets", "--target", NO_TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert len(rows[0]) == 8, rows[0]
total = rows[1]
assert total[0] == "TOTAL", total
assert all(int(v) == 0 for v in total[1:6]), total
assert float(total[6]) == 0 and float(total[7]) == 0, total
=== END AC REM-003 ===

=== AC REM-004 ===
Intent: The models report with a filter that matches nothing writes its header and a zero TOTAL row whose provider column is empty.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


def record(execution_id, started_at, status, command, target, model, returncode, timed_out):
    return {
        "schema_version": 1,
        "execution_id": execution_id,
        "started_at": started_at,
        "status": status,
        "job": {"command_name": command, "target": target, "llm": "claude", "model": model},
        "result": {
            "returncode": returncode,
            "timed_out": timed_out,
            "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                      "output_tokens": 1, "cost_usd": 0.01, "model": model},
        },
    }


NO_TARGET = "NoSuchTarget"
records = [record("e-run", "2026-10-14T02:38:41.904Z", "succeeded", "build", "Toml", "gemini-3-flash", 0, False)]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("models", "--target", NO_TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert len(rows[0]) == 8, rows[0]
total = rows[1]
assert total[0] == "TOTAL" and total[1] == "", total
assert all(int(v) == 0 for v in total[2:7]), total
assert float(total[7]) == 0, total
=== END AC REM-004 ===

=== AC REM-005 ===
Intent: The daily report with a filter that matches nothing writes its header and a zero TOTAL row.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


def record(execution_id, started_at, status, command, target, model, returncode, timed_out):
    return {
        "schema_version": 1,
        "execution_id": execution_id,
        "started_at": started_at,
        "status": status,
        "job": {"command_name": command, "target": target, "llm": "claude", "model": model},
        "result": {
            "returncode": returncode,
            "timed_out": timed_out,
            "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                      "output_tokens": 1, "cost_usd": 0.01, "model": model},
        },
    }


NO_TARGET = "NoSuchTarget"
records = [record("e-run", "2026-10-14T02:38:41.904Z", "succeeded", "build", "Toml", "gemini-3-flash", 0, False)]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("daily", "--target", NO_TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert len(rows[0]) == 5, rows[0]
total = rows[1]
assert total[0] == "TOTAL", total
assert all(int(v) == 0 for v in total[1:4]), total
assert float(total[4]) == 0, total
=== END AC REM-005 ===

=== AC REM-006 ===
Intent: The failures report with a filter that matches nothing writes its header and no data rows.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


def record(execution_id, started_at, status, command, target, model, returncode, timed_out):
    return {
        "schema_version": 1,
        "execution_id": execution_id,
        "started_at": started_at,
        "status": status,
        "job": {"command_name": command, "target": target, "llm": "claude", "model": model},
        "result": {
            "returncode": returncode,
            "timed_out": timed_out,
            "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                      "output_tokens": 1, "cost_usd": 0.01, "model": model},
        },
    }


NO_TARGET = "NoSuchTarget"
records = [record("e-run", "2026-10-14T02:38:41.904Z", "failed", "build", "Toml", "gemini-3-flash", 1, False)]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("failures", "--target", NO_TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 1, rows
assert len(rows[0]) == 6, rows[0]
=== END AC REM-006 ===

=== AC REM-007 ===
Intent: Every report over an empty log exits zero, writes its header, and writes a TOTAL row where the report has one.

import csv
import io
import os
import sys
import tempfile

from drydock_go import cli_runner


COLUMNS = {"summary": 11, "commands": 7, "targets": 8, "models": 8, "daily": 5, "failures": 6}
TOTAL_REPORTS = ["commands", "targets", "models", "daily"]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        handle.write("")
    results = {}
    for report in COLUMNS:
        code, out, err = cli_runner(report, path)
        assert code == 0, (report, code, out, err)
        results[report] = (code, out, err)

for report, (code, out, err) in results.items():
    print(report, code)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (report, code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert len(rows[0]) == COLUMNS[report], (report, rows[0])
    if report in TOTAL_REPORTS:
        assert rows[-1][0] == "TOTAL", (report, rows)
=== END AC REM-007 ===

## User Acceptance

- None.

## Guardrails

- Do not emit a placeholder or hard-coded row; every value is read back from the report output.
- Do not compare report output against a hand-typed known answer.
- Do not read or modify any file under `sources/` from application code.
