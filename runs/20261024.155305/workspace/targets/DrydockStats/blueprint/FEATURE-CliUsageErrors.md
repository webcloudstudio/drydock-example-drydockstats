# FEATURE: Usage Error Exit Codes

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Maps every command-line usage error to exit code 2 with no standard output. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliParse.md |
| Provides    | usage error exit 2 |
| Consumes    | Run |

## Purpose

Let a caller distinguish a malformed invocation from an operational failure, so a wrong command is never mistaken for an unreadable log.

## Behavior

- Usage errors are: unknown report name, unknown option, missing option value, invalid `--since` or `--until` date, and a missing or extra log path.
- Each usage error returns exit code `2`, writes one diagnostic line to standard error, and writes nothing to standard output.
- A date is valid only when `time.Parse` accepts it with layout `2006-01-02`; impossible dates such as `2026-02-30` are rejected.
- Usage checks run before the log is opened.

## Programmatic Acceptance

=== AC usage-unknown-report ===
Intent: An unknown report name exits 2 with empty stdout.
Suite: scoped
Requires: executable=python3; scope=test

import subprocess
import os
import sys
import tempfile

app = os.getcwd()
with tempfile.TemporaryDirectory() as d:
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run([bin_path, "frobnicate", os.devnull], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 2
    assert result.stdout == ""
=== END AC usage-unknown-report ===

=== AC usage-impossible-date ===
Intent: A calendar-impossible --since date exits 2 with empty stdout.
Suite: scoped
Requires: executable=python3; scope=test

import subprocess
import os
import sys
import tempfile

app = os.getcwd()
with tempfile.TemporaryDirectory() as d:
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run(
        [bin_path, "summary", "--since", "2026-02-30", os.devnull], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 2
    assert result.stdout == ""
=== END AC usage-impossible-date ===

## User Acceptance

- None.

## Guardrails

- A usage error never opens the log.
- Usage errors never write to standard output.
