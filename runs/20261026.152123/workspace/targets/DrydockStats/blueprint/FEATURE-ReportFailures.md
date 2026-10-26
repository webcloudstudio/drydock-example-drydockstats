# FEATURE: Report Failures

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Writes the failures CSV report listing each failed record in start order. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliExit.md, FEATURE-ValueFormat.md |
| Provides    | ddkstat failures |
| Consumes    | cli.Register, logread.Record |

## Purpose

The `failures` report lists every failed call so the operator can see which command, Target and model failed, and how.

## Behavior

- Header: `started_at,target,command,model,returncode,timed_out`.
- One row per failed record (status is not `succeeded`) among the filtered records.
- Rows are ordered by `started_at`, then `execution_id`, ascending by byte value.
- `started_at` is written exactly as logged.
- `model` is `Record.Model`: the executed model when present and non-empty, otherwise the requested model.
- `returncode` is empty when null or absent, otherwise the integer.
- `timed_out` is `true` or `false`.
- There is no `TOTAL` row. When no failed record matches, only the header is written.
- Selection and ordering (`SelectFailures`) and writing (`WriteFailures`) are separate functions in `internal/report`.
- The writer is registered with `cli.Register("failures", ...)`.

## Programmatic Acceptance

=== AC report-failures-unit ===
Intent: The native unit tests for failure selection, ordering, null returncode, timed_out text and empty output all pass.
Requires: executable=go; scope=test

from drydock_go import run_test

for name in (
    "TestFailuresOrdering",
    "TestFailuresNullReturncodeEmpty",
    "TestFailuresTimedOutText",
    "TestFailuresOnlyFailed",
    "TestFailuresEmpty",
):
    result = run_test("./internal/report", name)
    print(name, result.status)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC report-failures-unit ===

=== AC report-failures-rows ===
Intent: The failures command lists only failed records ordered by started_at then execution_id, with model fallback, empty null returncode and true/false timed_out.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile

from drydock_go import cli_runner


def rec(eid, started, target, cmd, status, rc, timed_out, stats_model):
    stats = {"input_tokens": 1, "elapsed_ms": 1000}
    if stats_model is not None:
        stats["model"] = stats_model
    return {
        "schema_version": 1,
        "execution_id": eid,
        "started_at": started,
        "status": status,
        "job": {"command_name": cmd, "target": target, "llm": "claude", "model": "requested-model"},
        "result": {"returncode": rc, "timed_out": timed_out, "stats": stats},
    }


recs = [
    rec("zz-9", "2026-10-15T06:00:00.000Z", "Beta", "diagnose", "failed", None, True, "executed-model"),
    rec("aa-1", "2026-10-14T02:00:00.000Z", "Alpha", "build", "succeeded", 0, False, "executed-model"),
    rec("bb-2", "2026-10-14T02:00:00.000Z", "Alpha", "build", "failed", 1, False, None),
    rec("aa-3", "2026-10-14T02:00:00.000Z", "Alpha", "plan", "failed", 2, False, "executed-model"),
    rec("cc-4", "2026-10-13T23:00:00.000Z", "Beta", "build", "timeout", 124, True, "executed-model"),
]

work = tempfile.mkdtemp()
path = os.path.join(work, "log.jsonl")
with open(path, "w", encoding="utf-8") as fh:
    for r in recs:
        fh.write(json.dumps(r) + "\n")

code, out, err = cli_runner("failures", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["started_at", "target", "command", "model", "returncode", "timed_out"]

failed = sorted(
    [r for r in recs if r["status"] != "succeeded"],
    key=lambda r: (r["started_at"], r["execution_id"]),
)
body = rows[1:]
assert len(body) == len(failed)
for row, r in zip(body, failed):
    model = r["result"]["stats"].get("model") or r["job"]["model"]
    rc = r["result"]["returncode"]
    assert row[0] == r["started_at"]
    assert row[1] == r["job"]["target"]
    assert row[2] == r["job"]["command_name"]
    assert row[3] == model
    assert row[4] == ("" if rc is None else str(rc))
    assert row[5] == str(r["result"]["timed_out"]).lower()
=== END AC report-failures-rows ===

=== AC report-failures-filters ===
Intent: The failures command honours --target and the inclusive --since/--until date range.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile

from drydock_go import cli_runner


def rec(eid, started, target):
    return {
        "schema_version": 1,
        "execution_id": eid,
        "started_at": started,
        "status": "failed",
        "job": {"command_name": "build", "target": target, "llm": "claude", "model": "m"},
        "result": {"returncode": 1, "timed_out": False, "stats": {"elapsed_ms": 1}},
    }


recs = [
    rec("id-1", "2026-10-14T01:00:00.000Z", "Alpha"),
    rec("id-2", "2026-10-15T01:00:00.000Z", "Alpha"),
    rec("id-3", "2026-10-15T02:00:00.000Z", "Beta"),
    rec("id-4", "2026-10-16T01:00:00.000Z", "Alpha"),
]
work = tempfile.mkdtemp()
path = os.path.join(work, "log.jsonl")
with open(path, "w", encoding="utf-8") as fh:
    for r in recs:
        fh.write(json.dumps(r) + "\n")

target = "Alpha"
since = "2026-10-15"
until = "2026-10-15"
kept = [
    r for r in recs
    if r["job"]["target"] == target and since <= r["started_at"][:10] <= until
]

code, out, err = cli_runner("failures", "--target", target, "--since", since, "--until", until, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert [row[0] for row in rows[1:]] == [r["started_at"] for r in kept]
=== END AC report-failures-filters ===

=== AC report-failures-empty ===
Intent: When no failed record matches, the failures command exits 0 and writes only the header row.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile

from drydock_go import cli_runner

record = {
    "schema_version": 1,
    "execution_id": "id-1",
    "started_at": "2026-10-14T01:00:00.000Z",
    "status": "succeeded",
    "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "m"},
    "result": {"returncode": 0, "timed_out": False},
}
work = tempfile.mkdtemp()
path = os.path.join(work, "log.jsonl")
with open(path, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

code, out, err = cli_runner("failures", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 1
assert rows[0] == ["started_at", "target", "command", "model", "returncode", "timed_out"]
=== END AC report-failures-empty ===

## User Acceptance

- None.

## Guardrails

- The report has no `TOTAL` row.
- Every row comes from the log being read; none is hardcoded.
- The report never writes to the log.
- Text values are written without quoting or escaping.
