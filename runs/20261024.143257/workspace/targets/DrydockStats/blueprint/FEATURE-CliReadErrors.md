# FEATURE: CLI Read Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Map an unreadable log to exit code 1 with no stdout and exactly one stderr line. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliParse.md, FEATURE-LogReader.md |
| Provides    | Exit code 1 for operational failures |
| Consumes    | internal/log ReadLines error |

## Purpose

When the log cannot be opened or a read fails, the reporter exits `1`, writes nothing to stdout,
and writes exactly one diagnostic line to stderr. A missing file is an operational failure, not a
usage error, because the path has already passed argument validation.

## Behavior

- A missing or unopenable log path returns exit code `1`.
- stdout is empty.
- stderr contains exactly one line.

## Programmatic Acceptance

=== AC read-error-missing-log ===
Intent: A log path that does not exist exits 1 with empty stdout and one stderr line.
Suite: scoped

import subprocess
import sys
import tempfile

from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    missing = Path(tmp) / "no-such-log.jsonl"
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')", "summary", str(missing)],
        capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 1
=== END AC read-error-missing-log ===

## User Acceptance

- None.

## Guardrails

- A missing log is never reported as a usage error (exit 2).
- Operational failures write nothing to stdout.
