# Feature: Read Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | An unreadable log exits 1 with empty stdout and one stderr diagnostic line. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogRead.md |
| Provides    | exit 1 for an unreadable log |
| Consumes    | — |

## Purpose

When the log file cannot be opened or read, exit 1 with nothing on stdout and one diagnostic line on stderr.

## Programmatic Acceptance

=== AC read-missing-log-exit-1 ===
Intent: A log path that does not exist exits 1 with empty stdout.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
build = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(build.stderr)
assert build.returncode == 0
missing = os.path.join(tempfile.mkdtemp(), "no-such-log.jsonl")
result = subprocess.run([out, "summary", missing], capture_output=True, text=True)
print(result.stderr)
assert result.returncode == 1
assert result.stdout == ""
=== END AC read-missing-log-exit-1 ===

## User Acceptance

- None.

## Guardrails

- A read failure writes nothing to stdout.
