# Feature: Command Parse

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Parse the report name, options, and the single log path from the argument list. |
| Depends On  | FEATURE-Entrypoint.md |
| Provides    | report, options, log path |
| Consumes    | — |

## Purpose

Take the report name as the first argument, parse the following options with a local `flag.FlagSet` (ContinueOnError), then require exactly one log path. Options follow the report name; the log path is last.

## Programmatic Acceptance

=== AC parse-summary-with-target ===
Intent: `summary --target X <log>` parses the report, the option value, and the log path.
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
result = subprocess.run([out, "summary", "--target", "X", log], capture_output=True, text=True)
print(result.stderr)
assert result.returncode == 0
=== END AC parse-summary-with-target ===

=== AC parse-repeated-option-exit-2 ===
Intent: A repeated --target option is a usage error with exit 2 and empty stdout.
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
result = subprocess.run(
    [out, "summary", "--target", "A", "--target", "B", log],
    capture_output=True, text=True,
)
print(result.stderr)
assert result.returncode == 2
assert result.stdout == ""
=== END AC parse-repeated-option-exit-2 ===

## User Acceptance

- None.

## Guardrails

- Options are parsed with a local `flag.FlagSet` using `flag.ContinueOnError`.
