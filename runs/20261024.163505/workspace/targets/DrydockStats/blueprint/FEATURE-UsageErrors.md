# Feature: Usage Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Unknown reports, unknown options, missing values, and bad log-path counts exit 2 with empty stdout. |
| Depends On  | FEATURE-CommandParse.md |
| Provides    | exit 2 usage-error path |
| Consumes    | — |

## Purpose

An unknown report, unknown option, missing option value, repeated option, invalid date, or missing or extra log path exits 2 with nothing on stdout and one diagnostic line on stderr.

## Programmatic Acceptance

=== AC usage-unknown-report-exit-2 ===
Intent: An unknown report name exits 2 with empty stdout.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
build = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(build.stderr)
assert build.returncode == 0
log = os.path.join(tempfile.mkdtemp(), "log.jsonl")
open(log, "w", encoding="utf-8").close()
result = subprocess.run([out, "frobnicate", log], capture_output=True, text=True)
print(result.stderr)
assert result.returncode == 2
assert result.stdout == ""
=== END AC usage-unknown-report-exit-2 ===

=== AC usage-no-log-path-exit-2 ===
Intent: A report with no log path exits 2 with empty stdout.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
build = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(build.stderr)
assert build.returncode == 0
result = subprocess.run([out, "summary"], capture_output=True, text=True, stdin=subprocess.DEVNULL)
print(result.stderr)
assert result.returncode == 2
assert result.stdout == ""
=== END AC usage-no-log-path-exit-2 ===

## User Acceptance

- None.

## Guardrails

- Usage errors write nothing to stdout.
