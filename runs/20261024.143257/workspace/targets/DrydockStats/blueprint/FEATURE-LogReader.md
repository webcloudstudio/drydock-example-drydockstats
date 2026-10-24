# Feature: Log Reader

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Read the log path as JSON Lines, yielding each non-blank line with its 1-based line number. |
| Depends On  | ARCHITECTURE.md |
| Provides    | internal/log ReadLines, internal/log Line |
| Consumes    | — |

## Purpose

Turn the log file into an ordered sequence of non-blank lines, each carrying its 1-based line number, so that later stages can classify and count them.

## Behavior

- The file is read as UTF-8.
- Each non-blank line is yielded with its 1-based line number. A line is blank when it contains only whitespace.
- Blank lines are ignored and are not counted as records or as skips.
- The core reader accepts an `io.Reader` so tests can supply a `bytes.Buffer`. Opening the file is a thin wrapper around it.
- No JSON parsing or line classification happens here. Classification belongs to the skip stage.
- When the file cannot be opened or a read fails, an error is returned so the caller can map it to exit code 1.

## Programmatic Acceptance

=== AC log-reader-blank-lines ===
Intent: Blank lines are ignored and do not appear in the yielded lines.
Requires: executable=python3; scope=test

import subprocess
import sys

from drydock_go import run_test

result = run_test("./internal/log", "TestReadLinesSkipsBlank")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-reader-blank-lines ===

=== AC log-reader-line-numbers ===
Intent: Yielded lines carry 1-based line numbers that count blank lines.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestReadLinesLineNumbers")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-reader-line-numbers ===

=== AC log-reader-open-error ===
Intent: A log path that cannot be opened returns an error rather than an empty sequence.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestReadLinesOpenError")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-reader-open-error ===

## User Acceptance

- None.

## Guardrails

- The reader never writes to the log it reads.
- The reader does not parse JSON.
