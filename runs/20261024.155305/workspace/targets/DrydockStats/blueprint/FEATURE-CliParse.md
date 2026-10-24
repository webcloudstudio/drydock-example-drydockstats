# FEATURE: Command Line Parsing

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Parses the report name, options, and the single log path, then dispatches to the matching report. |
| Depends On  | ARCHITECTURE.md, FEATURE-BuildEntry.md, FEATURE-FilterDate.md, FEATURE-ReportEmptyMatch.md |
| Provides    | ddkstat command line |
| Consumes    | Run, FormatCost, summary report, commands report, targets report, models report, daily report, failures report |

## Purpose

Give the operator one invocation shape: `ddkstat <report> [--target <name>] [--since <date>] [--until <date>] <log>`.

## Behavior

- `Run(args []string, stdout, stderr io.Writer) int` uses a local `flag.FlagSet` configured with `flag.ContinueOnError`.
- The first argument is the report name and must be one of `summary`, `commands`, `targets`, `models`, `daily`, `failures`.
- Options follow the report name. Exactly one log path is the final argument.
- `--target`, `--since`, and `--until` are passed to the filters; `--since` and `--until` are validated by the CLI before filtering.
- On a valid invocation, the matching report function is called and its exit code is returned.
- Usage-error classification is owned by FEATURE-CliUsageErrors.md; this feature only parses and dispatches.

## Programmatic Acceptance

=== AC cli-parse-dispatch-summary ===
Intent: A valid summary invocation with one log path exits 0 and writes the summary header.
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
record = {"schema_version": 1, "execution_id": "p1", "started_at": "2026-10-14T02:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False}}
with tempfile.TemporaryDirectory() as d:
    log = os.path.join(d, "one.jsonl")
    with open(log, "w") as f:
        f.write(json.dumps(record) + "\n")
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run([bin_path, "summary", log], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    rows = list(csv.reader(io.StringIO(result.stdout)))
    assert rows[0][0] == "records"
    assert rows[1][0] == "1"
=== END AC cli-parse-dispatch-summary ===

=== AC cli-parse-target-option ===
Intent: The --target option is passed to the filter so only the named target's records are counted.
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
toml = {"schema_version": 1, "execution_id": "p2", "started_at": "2026-10-14T02:00:00.000Z",
        "status": "succeeded", "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
        "result": {"returncode": 0, "timed_out": False}}
other = {"schema_version": 1, "execution_id": "p3", "started_at": "2026-10-14T03:00:00.000Z",
         "status": "succeeded", "job": {"command_name": "build", "target": "Other", "llm": "claude", "model": "haiku"},
         "result": {"returncode": 0, "timed_out": False}}
with tempfile.TemporaryDirectory() as d:
    log = os.path.join(d, "two.jsonl")
    with open(log, "w") as f:
        f.write(json.dumps(toml) + "\n")
        f.write(json.dumps(other) + "\n")
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run(
        [bin_path, "summary", "--target", "Toml", log], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    rows = list(csv.DictReader(io.StringIO(result.stdout)))
    assert rows[0]["records"] == "1"
=== END AC cli-parse-target-option ===

## User Acceptance

- None.

## Guardrails

- The log path is always the final argument; options never precede the report name.
- No report is produced when argument parsing fails.
