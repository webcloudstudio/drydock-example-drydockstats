# FEATURE: Report Failures

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the failures report: one CSV row per failed record, ordered by start time then execution id, with no TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportSummary.md, FEATURE-ReportFormat.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliReadErrors.md |
| Provides    | failures report, dispatch entry for failures |

## Purpose

The failures report lists each failed record so the operator can see which calls failed, where, and how. It is the only report with no TOTAL row.

## Behavior

- Header: `started_at,target,command,model,returncode,timed_out`.
- A record is failed when `status` is not `succeeded`, per the shared log rules.
- One row per failed record, ordered by `started_at` then `execution_id`, both ascending by byte value.
- `returncode` is empty when `result.returncode` is `null` or absent. Otherwise it is the integer.
- `timed_out` is written as `true` or `false`.
- `model` is the attributed model from the log-model rule.
- Filters `--target`, `--since`, and `--until` apply before rows are selected.
- The `failures` name is registered in the `cli-parse` dispatch table.

## Programmatic Acceptance

=== AC RF-001 ===
Intent: The failures report writes its header and one row per failed record in started_at then execution_id order, with returncode empty when null and timed_out as true or false.

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


def row_of(rec):
    rc = rec["result"].get("returncode")
    return [
        rec["started_at"],
        rec["job"]["target"],
        rec["job"]["command_name"],
        rec["result"]["stats"]["model"],
        "" if rc is None else str(rc),
        "true" if rec["result"]["timed_out"] else "false",
    ]


HEADER = ["started_at", "target", "command", "model", "returncode", "timed_out"]
records = [
    record("b-run", "2026-10-14T02:38:41.904Z", "failed", "diagnose", "Toml", "gemini-3-flash", 1, False),
    record("a-run", "2026-10-14T02:38:41.904Z", "failed", "build", "CommonMark", "claude-sonnet-5-5", None, True),
    record("c-run", "2026-10-15T06:50:39.621Z", "failed", "diagnose", "CommonMark", "claude-sonnet-5-5", None, True),
    record("d-run", "2026-10-16T00:00:00.000Z", "succeeded", "build", "Toml", "gemini-3-flash", 0, False),
]
failed = sorted(
    [r for r in records if r["status"] != "succeeded"],
    key=lambda r: (r["started_at"], r["execution_id"]),
)
expected = [row_of(r) for r in failed]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("failures", path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == HEADER, rows[0]
assert rows[1:] == expected, (rows[1:], expected)
=== END AC RF-001 ===

=== AC RF-002 ===
Intent: The failures report has no TOTAL row, and --target keeps only the failed records of that Target.

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


def row_of(rec):
    rc = rec["result"].get("returncode")
    return [
        rec["started_at"],
        rec["job"]["target"],
        rec["job"]["command_name"],
        rec["result"]["stats"]["model"],
        "" if rc is None else str(rc),
        "true" if rec["result"]["timed_out"] else "false",
    ]


TARGET = "CommonMark"
records = [
    record("b-run", "2026-10-14T02:38:41.904Z", "failed", "diagnose", "Toml", "gemini-3-flash", 1, False),
    record("a-run", "2026-10-14T02:38:41.904Z", "failed", "build", "CommonMark", "claude-sonnet-5-5", None, True),
    record("c-run", "2026-10-15T06:50:39.621Z", "failed", "diagnose", "CommonMark", "claude-sonnet-5-5", None, True),
    record("d-run", "2026-10-16T00:00:00.000Z", "succeeded", "build", "CommonMark", "claude-sonnet-5-5", 0, False),
]
failed = sorted(
    [r for r in records if r["status"] != "succeeded" and r["job"]["target"] == TARGET],
    key=lambda r: (r["started_at"], r["execution_id"]),
)
expected = [row_of(r) for r in failed]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        for rec in records:
            print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("failures", "--target", TARGET, path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[1:] == expected, (rows[1:], expected)
assert not any(row and row[0] == "TOTAL" for row in rows), rows
=== END AC RF-002 ===

=== AC RF-003 ===
Intent: A failed record with returncode absent from the log writes an empty returncode field.

import csv
import io
import json
import os
import sys
import tempfile

from drydock_go import cli_runner


rec = {
    "schema_version": 1,
    "execution_id": "absent-run",
    "started_at": "2026-10-17T01:02:03.000Z",
    "status": "failed",
    "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
    "result": {
        "timed_out": False,
        "stats": {"elapsed_ms": 1000, "input_tokens": 1, "cached_input_tokens": 0,
                  "output_tokens": 1, "cost_usd": 0.01, "model": "claude-haiku-5-5"},
    },
}

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "llm.jsonl")
    with open(path, "w", encoding="utf-8") as handle:
        print(json.dumps(rec), file=handle)
    code, out, err = cli_runner("failures", path)

print(out)
print(err, file=sys.stderr)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1][4] == "", rows[1]
assert rows[1][5] == "false", rows[1]
=== END AC RF-003 ===

## User Acceptance

- None.

## Guardrails

- Do not emit a TOTAL row in the failures report.
- Do not compare any report row against a hand-typed known answer; expected rows derive from the fixture records.
- Do not modify or read any file under `sources/` from application code.
