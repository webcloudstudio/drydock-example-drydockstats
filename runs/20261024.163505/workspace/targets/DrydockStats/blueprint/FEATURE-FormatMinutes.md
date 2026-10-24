# FEATURE: Format Minutes in Tenths

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Render a summed elapsed_ms as minutes to one decimal place using the specified integer rounding. |
| Depends On  | ARCHITECTURE.md |
| Provides    | FormatMinutes(ms) string |
| Consumes    | none |

## Purpose

Every minutes column in every report is rendered by this single formatter so the rounding rule is
defined in one place.

## Behavior

- Input: the summed `elapsed_ms` of a group, as a non-negative integer.
- Compute `t = (ms + 3000) / 6000` in integer division.
- Output is `t / 10`, a period, then `t % 10`. Example: 95000 ms gives `t = 16`, rendered `1.6`.
- Zero input gives `t = 0`, rendered `0.0`.

## Programmatic Acceptance

=== AC fmt-minutes-example ===
Intent: The minutes formatter renders 95000 ms as 1.6 by observing a summary report that sums one record of that duration.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "j-1", "started_at": "2026-10-14T09:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 95000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    header = rows[0]
    assert header[-1] == "minutes", header
    assert rows[1][-1] == "1.6", rows
finally:
    os.remove(path)
=== END AC fmt-minutes-example ===

=== AC fmt-minutes-zero ===
Intent: A group with zero elapsed time renders minutes as 0.0 in the summary report.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "k-1", "started_at": "2026-10-14T10:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 0, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][-1] == "0.0", rows
finally:
    os.remove(path)
=== END AC fmt-minutes-zero ===

## User Acceptance

- None.

## Guardrails

- The formula is integer division exactly as stated; no floating-point rounding is used.
- Minutes never render in any other format.
