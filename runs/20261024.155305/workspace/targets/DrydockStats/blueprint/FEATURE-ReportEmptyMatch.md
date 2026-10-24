# FEATURE: Empty-Match Report Output

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Ensures every report writes its header and, where applicable, a zero TOTAL row when no record matches. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReportSummary.md, FEATURE-ReportCommands.md, FEATURE-ReportTargets.md, FEATURE-ReportModels.md, FEATURE-ReportDaily.md, FEATURE-ReportFailures.md |
| Provides    | empty-match report output |
| Consumes    | summary report, commands report, targets report, models report, daily report, failures report |

## Purpose

When a filter matches no record, every report still produces its header, and the reports that carry a TOTAL row produce a zero TOTAL row, so the output shape never depends on whether data matched.

## Behavior

- `summary`: one data row with `records=0`, `succeeded=0`, `failed=0`, all token totals `0`, `cost_usd` `0.0000`, `minutes` `0.0`, and empty `first` and `last`. `skipped` still counts the whole file.
- `commands`, `targets`, `models`, `daily`: header followed by a TOTAL row of zeros (`0.0` minutes, `0.0000` cost where present). Models writes an empty model column in its TOTAL row.
- `failures`: header only.
- Exit code remains 0 for an empty match.

## Programmatic Acceptance

=== AC empty-match-commands ===
Intent: The commands report writes header and zero TOTAL row when a filter matches nothing.
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
record = {"schema_version": 1, "execution_id": "x1", "started_at": "2026-10-14T02:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"model": "haiku", "input_tokens": 1,
                     "cached_input_tokens": 1, "output_tokens": 1, "cost_usd": 0.1, "elapsed_ms": 1000}}}
with tempfile.TemporaryDirectory() as d:
    log = os.path.join(d, "one.jsonl")
    with open(log, "w") as f:
        f.write(json.dumps(record) + "\n")
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run(
        [bin_path, "commands", "--target", "NoSuchTarget", log], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    rows = list(csv.reader(io.StringIO(result.stdout)))
    assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"]
    assert rows[-1][0] == "TOTAL"
    assert rows[-1][1:] == ["0", "0", "0", "0", "0", "0.0"]
=== END AC empty-match-commands ===

=== AC empty-match-summary ===
Intent: The summary report on an empty match writes empty first and last and zero totals.
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
record = {"schema_version": 1, "execution_id": "x2", "started_at": "2026-10-14T02:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False}}
with tempfile.TemporaryDirectory() as d:
    log = os.path.join(d, "one.jsonl")
    with open(log, "w") as f:
        f.write(json.dumps(record) + "\n")
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run(
        [bin_path, "summary", "--target", "NoSuchTarget", log], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    rows = list(csv.DictReader(io.StringIO(result.stdout)))
    assert len(rows) == 1
    assert rows[0]["records"] == "0"
    assert rows[0]["first"] == ""
    assert rows[0]["last"] == ""
    assert rows[0]["skipped"] == "0"
=== END AC empty-match-summary ===

## User Acceptance

- None.

## Guardrails

- An empty match never exits non-zero.
- Empty-match output never invents a row for a record that did not match.
