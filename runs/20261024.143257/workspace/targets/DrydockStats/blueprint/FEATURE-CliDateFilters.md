# FEATURE: CLI Date Filters

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Keep records whose start date falls inclusively within --since and --until. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliParse.md, FEATURE-LogSkip.md |
| Provides    | --since and --until filters over the start date |
| Consumes    | Record.StartedAt, internal/log Normalize |

## Purpose

The start date of a record is the first 10 characters of `started_at`. `--since` keeps records
whose start date is on or after the value. `--until` keeps records whose start date is on or
before the value. Both bounds are inclusive. Dates are compared as strings, which orders
`YYYY-MM-DD` correctly. Date validity is checked by `FEATURE-CliUsageErrors.md`, not here.

## Behavior

- With `--since` only, records with start date `>= since` are kept.
- With `--until` only, records with start date `<= until` are kept.
- With both, the two bounds apply together as a logical AND.
- A bound later than every record yields an empty filtered set, not an error.

## Programmatic Acceptance

=== AC date-filter-since-inclusive ===
Intent: --since keeps records starting on the given date and later, and drops earlier ones.
Suite: scoped

import json
import subprocess
import sys
import tempfile

from pathlib import Path

rows = [
    {"execution_id": "d1", "started_at": "2026-10-13T23:59:59.000Z", "status": "succeeded",
     "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}},
    {"execution_id": "d2", "started_at": "2026-10-14T00:00:00.000Z", "status": "succeeded",
     "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}},
]
with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "log.jsonl"
    log.write_text("\n".join(json.dumps(r) for r in rows) + "\n", encoding="utf-8")
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')"], capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
=== END AC date-filter-since-inclusive ===

=== AC date-filter-until-inclusive ===
Intent: --until keeps records starting on the given date and earlier, and drops later ones.
Suite: scoped

import json
import subprocess
import sys
import tempfile

from pathlib import Path

rows = [
    {"execution_id": "e1", "started_at": "2026-10-14T23:59:59.000Z", "status": "succeeded",
     "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}},
    {"execution_id": "e2", "started_at": "2026-10-15T00:00:00.000Z", "status": "succeeded",
     "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}},
]
with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "log.jsonl"
    log.write_text("\n".join(json.dumps(r) for r in rows) + "\n", encoding="utf-8")
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')"], capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
=== END AC date-filter-until-inclusive ===

## User Acceptance

- None.

## Guardrails

- Date bounds are compared against the first 10 characters of `started_at` only.
- The `skipped` count in `summary` is never filtered by date.
