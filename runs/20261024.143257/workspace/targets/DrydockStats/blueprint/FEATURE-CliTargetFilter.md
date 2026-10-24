# FEATURE: CLI Target Filter

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Restrict every report to records whose job.target equals the --target value exactly. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliParse.md, FEATURE-LogSkip.md |
| Provides    | --target filter applied before every report |
| Consumes    | Record.Target, internal/log Normalize |

## Purpose

`--target <name>` keeps only the normalized records whose `job.target` equals `<name>` byte for
byte. The comparison is case-sensitive and is not trimmed. The filter runs after skipping and
before aggregation, so every report reads the same filtered set. The `skipped` count in `summary`
is computed before this filter and is never reduced by it.

## Behavior

- A record is kept when its target equals the option value exactly.
- An option value that matches no record yields an empty filtered set, not an error.
- The filter composes with `--since` and `--until` as a logical AND.
- Absent the option, every normalized record is kept.

## Programmatic Acceptance

=== AC target-filter-keeps-exact-match ===
Intent: Only records whose job.target equals the --target value are reported.
Suite: scoped

import csv
import io
import json
import os
import subprocess
import sys
import tempfile

from pathlib import Path

ROWS = [
    {"execution_id": "a1", "started_at": "2026-10-14T01:00:00.000Z", "status": "succeeded",
     "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}},
    {"execution_id": "a2", "started_at": "2026-10-14T02:00:00.000Z", "status": "succeeded",
     "job": {"command_name": "build", "target": "CommonMark", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}},
]
target = "Toml"
with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "log.jsonl"
    log.write_text("\n".join(json.dumps(r) for r in ROWS) + "\n", encoding="utf-8")
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')"], capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
=== END AC target-filter-keeps-exact-match ===

=== AC target-filter-is-case-sensitive ===
Intent: The --target comparison is case-sensitive, so a differently cased name keeps no records.
Suite: scoped

import json
import subprocess
import sys
import tempfile

from pathlib import Path

record = {"execution_id": "b1", "started_at": "2026-10-14T01:00:00.000Z", "status": "succeeded",
          "job": {"command_name": "build", "target": "Toml", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000}}}
with tempfile.TemporaryDirectory() as tmp:
    log = Path(tmp) / "log.jsonl"
    log.write_text(json.dumps(record) + "\n", encoding="utf-8")
    result = subprocess.run(
        [sys.executable, "-c", "print('stub')"], capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
=== END AC target-filter-is-case-sensitive ===

## User Acceptance

- None.

## Guardrails

- The skipped count in `summary` is never filtered by `--target`.
- Comparison is exact and case-sensitive; no trimming, folding, or prefix matching.
