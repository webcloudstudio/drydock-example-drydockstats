# FEATURE: Failures Report

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The failures report lists each filtered record whose status is not succeeded, in start order. |
| Depends On  | FEATURE-SummaryReport.md, FEATURE-ModelAttribution.md, FEATURE-DateFilter.md |
| Provides    | ddkstat failures |
| Consumes    | logread.Record.Model, report.FilterTarget, report.FilterDates |

## Purpose

`ddkstat failures <log>` shows which calls failed, when, and how.

## Report Contract

- Header: `started_at,target,command,model,returncode,timed_out`.
- One row per filtered record whose `status` is not `succeeded`, whatever the status value.
- Rows are ordered by `started_at` ascending, then `execution_id` ascending, by byte value.
- `target` is `job.target`, `command` is `job.command_name`, and `model` is the attributed model (`result.stats.model` when non-empty, otherwise `job.model`).
- `returncode` is the integer, or empty when `result.returncode` is null or absent.
- `timed_out` is `true` or `false`.
- There is no `TOTAL` row. With no failed record the output is the header only.
- Output is CSV with `\n` line endings; `--target`, `--since`, and `--until` narrow the records first.

## Programmatic Acceptance

=== AC failures-rows-order-and-fields ===
Intent: Failures lists only non-succeeded records, ordered by start then id, with attributed model and blank null return codes.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

ABSENT = object()

def rec(eid, started, status, cmd, target, jm, rc, timed, stats=None):
    result = {"timed_out": timed}
    if rc is not ABSENT:
        result["returncode"] = rc
    if stats is not None:
        result["stats"] = stats
    return {"execution_id": eid, "started_at": started, "status": status,
            "job": {"command_name": cmd, "target": target, "llm": "claude", "model": jm},
            "result": result}

recs = [
    rec("e2", "2026-10-14T02:00:00.000Z", "failed", "build", "Alpha", "mA", 1, False),
    rec("ok", "2026-10-14T01:00:00.000Z", "succeeded", "build", "Alpha", "mA", 0, False),
    rec("e1", "2026-10-14T02:00:00.000Z", "failed", "plan", "Beta", "mB", 2, True, {"model": "executed-model"}),
    rec("e3", "2026-10-13T09:00:00.000Z", "failed", "diagnose", "Alpha", "mC", None, True),
    rec("e4", "2026-10-15T09:00:00.000Z", "error", "analyze", "Beta", "mD", ABSENT, False, {"model": ""}),
    rec("e5", "2026-10-16T09:00:00.000Z", "timeout", "build", "Alpha", "mE", 137, False),
]
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write("\n".join(json.dumps(r) for r in recs) + "\n")

code, out, err = cli_runner("failures", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["started_at", "target", "command", "model", "returncode", "timed_out"]
failed = sorted((r for r in recs if r["status"] != "succeeded"),
                key=lambda r: (r["started_at"], r["execution_id"]))
expected = []
for r in failed:
    s = r["result"].get("stats") or {}
    rc = r["result"].get("returncode")
    expected.append([r["started_at"], r["job"]["target"], r["job"]["command_name"],
                     s.get("model") or r["job"]["model"],
                     "" if rc is None else str(rc),
                     json.dumps(bool(r["result"]["timed_out"]))])
assert rows[1:] == expected
=== END AC failures-rows-order-and-fields ===

=== AC failures-filters ===
Intent: Target and date filters narrow the failures report.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

def rec(eid, started, target):
    return {"execution_id": eid, "started_at": started, "status": "failed",
            "job": {"command_name": "build", "target": target, "llm": "claude", "model": "m"},
            "result": {"returncode": 1, "timed_out": False}}

recs = [
    rec("a1", "2026-10-14T01:00:00.000Z", "Alpha"),
    rec("a2", "2026-10-15T01:00:00.000Z", "Alpha"),
    rec("a3", "2026-10-15T02:00:00.000Z", "Beta"),
    rec("a4", "2026-10-16T01:00:00.000Z", "Alpha"),
]
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write("\n".join(json.dumps(r) for r in recs) + "\n")

target, since, until = "Alpha", "2026-10-15", "2026-10-15"
code, out, err = cli_runner("failures", "--target", target, "--since", since, "--until", until, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
kept = [r for r in recs if r["job"]["target"] == target and since <= r["started_at"][:10] <= until]
assert [row[0] for row in rows[1:]] == [r["started_at"] for r in kept]
=== END AC failures-filters ===

=== AC failures-none ===
Intent: With no failed record the failures report is only the header and exits 0.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

rec = {"execution_id": "s1", "started_at": "2026-10-14T01:00:00.000Z", "status": "succeeded",
       "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "m"},
       "result": {"returncode": 0, "timed_out": False}}
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write(json.dumps(rec) + "\n")

code, out, err = cli_runner("failures", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 1
assert rows[0] == ["started_at", "target", "command", "model", "returncode", "timed_out"]
=== END AC failures-none ===

## User Acceptance

- None.

## Guardrails

- Every value is computed from the log being read; no row is hardcoded.
- The log is never written to, and `sources/run.sh` is never read by application code.
- Only the Go standard library is used.
