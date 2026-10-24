# FEATURE: CLI Success on Partial Data

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Exit 0 and write the report when lines are skipped or no record matches. |
| Depends On  | FEATURE-CliUnreadableLog.md, FEATURE-CliParse.md, FEATURE-ReportEmptyMatch.md, FEATURE-LogDuplicates.md |
| Provides    | success exit 0 on partial data |
| Consumes    | ddkstat command line, unreadable log exit 1 |

## Purpose

A log with skipped lines, or a filter that matches no records, is a successful run. The reporter exits 0 and writes its report to standard output. Skip warnings go to standard error only.

## Behavior

- A log containing blank, malformed, required-field-missing, or duplicate lines exits 0.
- A filter that matches no records exits 0 and writes the header plus the zero TOTAL row where the report has one.
- Skip warnings appear on standard error only; standard output contains only the report CSV.
- No code path returns a non-zero status for these cases.

## Programmatic Acceptance

=== AC cli-partial-skipped-exit-zero ===
Intent: A log with malformed and duplicate lines exits 0 and its summary reports skipped lines.
Suite: scoped

import csv
import io
import json
import os
import subprocess
import sys
import tempfile

from pathlib import Path

app = Path(os.getcwd())
build = app / "ddkstat_build_probe"
subprocess.run(["go", "build", "-o", str(build), "./cmd/ddkstat"], cwd=app, check=True)

lines = [
    json.dumps({"schema_version": 1, "execution_id": "a", "started_at": "2026-10-14T01:00:00.000Z",
                "status": "succeeded", "job": {"command_name": "build", "target": "T", "llm": "claude", "model": "haiku"},
                "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "cost_usd": 0.5}}}),
    "not json at all",
    json.dumps({"schema_version": 1, "execution_id": "a", "started_at": "2026-10-14T01:00:00.000Z",
                "status": "succeeded", "job": {"command_name": "build", "target": "T", "llm": "claude", "model": "haiku"},
                "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "cost_usd": 0.5}}}),
]
with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "partial.jsonl"
    log.write_text("\n".join(lines) + "\n", encoding="utf-8")
    result = subprocess.run([str(build), "summary", str(log)], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0, (result.returncode, result.stderr)
    rows = list(csv.DictReader(io.StringIO(result.stdout)))
    assert len(rows) == 1
    assert rows[0]["records"] == "1"
    assert rows[0]["skipped"] == "2"
=== END AC cli-partial-skipped-exit-zero ===

=== AC cli-partial-no-match-exit-zero ===
Intent: A filter that matches no records exits 0 and the commands report writes its header with a zero TOTAL row.
Suite: scoped

import csv
import io
import json
import os
import subprocess
import sys
import tempfile

from pathlib import Path

app = Path(os.getcwd())
build = app / "ddkstat_build_probe"
subprocess.run(["go", "build", "-o", str(build), "./cmd/ddkstat"], cwd=app, check=True)

record = json.dumps({"schema_version": 1, "execution_id": "a", "started_at": "2026-10-14T01:00:00.000Z",
                     "status": "succeeded", "job": {"command_name": "build", "target": "T", "llm": "claude", "model": "haiku"},
                     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "cost_usd": 0.5}}})
with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "one.jsonl"
    log.write_text(record + "\n", encoding="utf-8")
    result = subprocess.run([str(build), "commands", "--target", "NoSuchTarget", str(log)],
                            capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0, (result.returncode, result.stderr)
    rows = list(csv.reader(io.StringIO(result.stdout)))
    assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"]
    assert rows[-1][0] == "TOTAL"
    assert rows[-1][1] == "0"
=== END AC cli-partial-no-match-exit-zero ===

## User Acceptance

- None.

## Guardrails

- Skip warnings never go to standard output.
- Exit code is never non-zero because of skipped lines or an empty match.
