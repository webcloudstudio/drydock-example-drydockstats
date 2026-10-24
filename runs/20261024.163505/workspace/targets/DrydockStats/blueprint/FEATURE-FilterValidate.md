# FEATURE: Validate Date Option Values

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Accept --since and --until only as real calendar dates written YYYY-MM-DD; any other value is a usage error with exit 2. |
| Depends On  | ARCHITECTURE.md, FEATURE-UsageErrors.md |
| Provides    | validated --since and --until values |
| Consumes    | usage-error path |

## Purpose

The date options feed the date filter. Only real calendar dates reach it, so an impossible date such
as `2026-13-01` or `2026-02-30` cannot silently filter the log. Validation runs over the full argument
list, including an invalid date that appears after the log path.

## Behavior

- A date value must match the form `YYYY-MM-DD` exactly: four digits, a hyphen, two digits, a hyphen,
  two digits.
- The value must name a real calendar date. Month is 01 through 12. Day must be valid for that month
  and year, including leap years in the Gregorian calendar.
- Any other value is a usage error: exit 2, nothing on stdout, one diagnostic line on stderr.
- Validation applies to every `--since` and `--until` occurrence, even when the log path precedes it.

## Programmatic Acceptance

=== AC filter-validate-rejects-impossible-date ===
Intent: An impossible calendar date in --since exits 2 with nothing on stdout.
Requires: executable=python3; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "g-1", "started_at": "2026-10-14T06:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", "--since", "2026-13-01", path)
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == "", out
finally:
    os.remove(path)
=== END AC filter-validate-rejects-impossible-date ===

=== AC filter-validate-rejects-invalid-day ===
Intent: A day that does not exist in its month (February 30) is a usage error with exit 2.
Requires: executable=python3; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "h-1", "started_at": "2026-10-14T07:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", "--until", "2026-02-30", path)
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == "", out
finally:
    os.remove(path)
=== END AC filter-validate-rejects-invalid-day ===

=== AC filter-validate-rejects-date-after-log ===
Intent: An invalid date that appears after the log path is still a usage error with exit 2.
Requires: executable=python3; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "i-1", "started_at": "2026-10-14T08:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", path, "--since", "2026-13-01")
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == "", out
finally:
    os.remove(path)
=== END AC filter-validate-rejects-date-after-log ===

## User Acceptance

- None.

## Guardrails

- A date is never corrected, rolled over, or reinterpreted to make it valid (for example, `2026-02-30`
  is never treated as `2026-03-01`).
- Only the form `YYYY-MM-DD` is accepted; other date spellings are rejected even when they parse.
