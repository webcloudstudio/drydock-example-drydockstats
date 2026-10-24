# FEATURE: Failures Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces the failures report, one CSV row per failed filtered record. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportFormat.md, FEATURE-LogModel.md, FEATURE-FilterDate.md |
| Provides    | failures report |
| Consumes    | FormatCost, WriteCSV, filtered records, model attribution |

## Purpose

Write one row for each failed record that survives the --target, --since, and --until filters, so an operator can see which calls failed and why.

## Behavior

- Header: `started_at,target,command,model,returncode,timed_out`.
- A record is failed when `status` is not `succeeded`.
- Each row carries the record's `started_at`, `job.target`, `job.command_name`, the model from LOG-004 attribution, `result.returncode`, and `result.timed_out`.
- `returncode` is written empty when `result.returncode` is `null` or absent; otherwise its integer value.
- `timed_out` is written as the literal `true` or `false`.
- Rows are ordered by `started_at`, then `execution_id`, both ascending by byte-wise comparison.
- There is no TOTAL row.
- A filter that matches no failed record writes the header only.
- The skipped count is not part of this report.

## Programmatic Acceptance

=== AC report-failures-header ===
Intent: The failures report writes the declared header on an empty match.
Suite: scoped
Requires: executable=python3; scope=test

import subprocess
import csv
import io
import json
import os
import sys
import tempfile

app = os.getcwd()
with tempfile.TemporaryDirectory() as d:
    log = os.path.join(d, "empty.jsonl")
    open(log, "w").close()
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run(
        [bin_path, "failures", log], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    rows = list(csv.reader(io.StringIO(result.stdout)))
    expected_header = ["started_at", "target", "command", "model", "returncode", "timed_out"]
    assert rows[0] == expected_header
    assert len(rows) == 1
=== END AC report-failures-header ===

=== AC report-failures-rows ===
Intent: Each failed record yields one row with empty returncode for null and literal timed_out.
Suite: scoped
Requires: executable=python3; scope=test

import subprocess
import csv
import io
import json
import os
import sys
import tempfile

app = os.getcwd()
records = [
    {"schema_version": 1, "execution_id": "b", "started_at": "2026-10-14T02:00:00.000Z",
     "status": "failed", "job": {"command_name": "diagnose", "target": "Toml", "llm": "gemini", "model": "gemini-3-flash"},
     "result": {"returncode": 1, "timed_out": False, "stats": {"model": "gemini-3-flash"}}},
    {"schema_version": 1, "execution_id": "a", "started_at": "2026-10-15T06:50:39.621Z",
     "status": "failed", "job": {"command_name": "diagnose", "target": "CommonMark", "llm": "claude", "model": "sonnet"},
     "result": {"returncode": None, "timed_out": True, "stats": {"model": "claude-sonnet-5-5"}}},
    {"schema_version": 1, "execution_id": "c", "started_at": "2026-10-16T00:00:00.000Z",
     "status": "succeeded", "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "sonnet"},
     "result": {"returncode": 0, "timed_out": False}},
]
with tempfile.TemporaryDirectory() as d:
    log = os.path.join(d, "failures.jsonl")
    with open(log, "w") as f:
        for r in records:
            f.write(json.dumps(r) + "\n")
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run(
        [bin_path, "failures", log], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    rows = list(csv.DictReader(io.StringIO(result.stdout)))
    assert [r["execution_id"] if "execution_id" in r else r["started_at"] for r in rows] == [
        "2026-10-14T02:00:00.000Z", "2026-10-15T06:50:39.621Z"
    ]
    assert rows[0]["returncode"] == "1"
    assert rows[0]["timed_out"] == "false"
    assert rows[1]["returncode"] == ""
    assert rows[1]["timed_out"] == "true"
    assert rows[1]["model"] == "claude-sonnet-5-5"
=== END AC report-failures-rows ===

## User Acceptance

- None.

## Guardrails

- The failures report never writes a TOTAL row.
- Failed status is decided by `status` only; `returncode` never decides failure.
