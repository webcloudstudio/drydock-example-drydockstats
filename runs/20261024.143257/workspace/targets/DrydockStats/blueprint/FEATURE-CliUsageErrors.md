# FEATURE: CLI Usage Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Map every usage error to exit code 2 with no stdout and exactly one stderr line. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliParse.md |
| Provides    | Exit code 2 for usage errors |
| Consumes    | cli.Run argument handling |

## Purpose

A usage error is an unknown report name, an unknown option, a missing option value, a date that is
not a real `YYYY-MM-DD` calendar date, or a missing or extra log path. Each exits `2`, writes
nothing to stdout, and writes exactly one diagnostic line to stderr. Dates are validated with
`time.Parse` using layout `2006-01-02`, rejecting any string that does not round-trip exactly.

## Behavior

- Unknown report, unknown option, missing option value, invalid date, and missing or extra log path
  each return exit code `2`.
- stdout is empty for every usage error.
- stderr contains exactly one line.

## Programmatic Acceptance

=== AC usage-error-unknown-report ===
Intent: An unknown report name exits 2 with empty stdout and one stderr line.
Suite: scoped

import subprocess
import sys
import tempfile

from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "log.jsonl"
    log.write_text("", encoding="utf-8")
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')", "frobnicate", str(log)],
        capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 2
=== END AC usage-error-unknown-report ===

=== AC usage-error-invalid-date ===
Intent: A date that is not a real YYYY-MM-DD calendar date exits 2 with empty stdout.
Suite: scoped

import subprocess
import sys
import tempfile

from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "log.jsonl"
    log.write_text("", encoding="utf-8")
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')", "summary", "--since", "2026-13-01", str(log)],
        capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 2
=== END AC usage-error-invalid-date ===

## User Acceptance

- None.

## Guardrails

- Usage errors never write to stdout.
- Exactly one diagnostic line is written to stderr per usage error.
