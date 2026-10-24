# FEATURE: Format Summed Cost

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Render a summed cost_usd with exactly four decimal places using %.4f. |
| Depends On  | ARCHITECTURE.md |
| Provides    | FormatCost(usd) string |
| Consumes    | none |

## Purpose

Every cost column in every report is rendered by this single formatter so the precision is defined in
one place.

## Behavior

- Input: the summed `cost_usd` of a group as a `float64`. Null or absent costs were already counted as
  zero by FEATURE-LogNumbers.md.
- Output is the value formatted with `%.4f`: exactly four digits after the decimal point.
- A zero total renders `0.0000`.

## Programmatic Acceptance

=== AC fmt-cost-four-places ===
Intent: The summed cost renders with exactly four decimal places in the summary report.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "l-1", "started_at": "2026-10-14T11:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.5}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    cost = rows[1][9]
    assert len(cost.split(".")[1]) == 4, cost
    assert float(cost) == 0.5, cost
finally:
    os.remove(path)
=== END AC fmt-cost-four-places ===

=== AC fmt-cost-zero ===
Intent: A zero cost total renders as 0.0000 in the summary report.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "m-1", "started_at": "2026-10-14T12:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][9] == "0.0000", rows
finally:
    os.remove(path)
=== END AC fmt-cost-zero ===

## User Acceptance

- None.

## Guardrails

- The cost is formatted with `%.4f` only; it is never rounded by any other method or truncated.
