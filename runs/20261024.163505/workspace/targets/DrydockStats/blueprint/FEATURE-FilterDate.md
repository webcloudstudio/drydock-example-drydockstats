# FEATURE: Filter Records by Start Date

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Keep only log records whose start date falls within the inclusive --since and --until range. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogSkip.md, FEATURE-CommandParse.md, FEATURE-FilterValidate.md |
| Provides    | date range filter predicate |
| Consumes    | record parser, parsed options |

## Purpose

Every report applies the same start-date predicate. A record's start date is the first 10 characters
of its `started_at` string. When `--since` or `--until` is given, only records whose start date lies
within the inclusive range are kept.

## Behavior

- Start date = `started_at[0:10]`, compared as a `YYYY-MM-DD` string. Byte-order comparison of
  validated dates equals chronological order.
- `--since D` keeps records with start date `>= D`. `--until D` keeps records with start date `<= D`.
  Both bounds are inclusive and independently optional.
- Inverted ranges (`--since` later than `--until`) match no record and exit 0; the report is
  header-only or zero-TOTAL. This is the narrowest consistent behavior under the AC-004 precedent.
- Option values are validated as real calendar dates by FEATURE-FilterValidate.md before this
  predicate is built; this feature receives only valid dates.
- The `skipped` count is not affected by this predicate.

## Programmatic Acceptance

=== AC filter-date-inclusive-bounds ===
Intent: --since and --until keep records whose start date equals either bound and exclude dates outside the range.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

dates = ["2026-10-13", "2026-10-14", "2026-10-15", "2026-10-16", "2026-10-17"]
records = []
for i, d in enumerate(dates):
    records.append({"schema_version": 1, "execution_id": f"d-{i}", "started_at": f"{d}T04:00:00.000Z",
                    "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
                    "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                               "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}})

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    for obj in records:
        fh.write(json.dumps(obj) + "\n")

try:
    code, out, err = cli_runner("summary", "--since", "2026-10-14", "--until", "2026-10-16", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][0] == "3", rows
    assert rows[1][4] == "2026-10-14", rows
    assert rows[1][5] == "2026-10-16", rows
finally:
    os.remove(path)
=== END AC filter-date-inclusive-bounds ===

=== AC filter-date-since-only ===
Intent: --since alone keeps records on or after the bound and nothing earlier.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

dates = ["2026-10-13", "2026-10-14", "2026-10-16"]
records = []
for i, d in enumerate(dates):
    records.append({"schema_version": 1, "execution_id": f"e-{i}", "started_at": f"{d}T04:00:00.000Z",
                    "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
                    "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                               "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}})

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    for obj in records:
        fh.write(json.dumps(obj) + "\n")

try:
    code, out, err = cli_runner("summary", "--since", "2026-10-14", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][0] == "2", rows
    assert rows[1][4] == "2026-10-14", rows
finally:
    os.remove(path)
=== END AC filter-date-since-only ===

=== AC filter-date-inverted-range-exit-zero ===
Intent: An inverted range (--since later than --until) matches nothing and exits 0.
Requires: executable=python3; scope=test

import csv
import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "f-1", "started_at": "2026-10-14T05:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", "--since", "2026-10-16", "--until", "2026-10-14", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][0] == "0", rows
finally:
    os.remove(path)
=== END AC filter-date-inverted-range-exit-zero ===

## User Acceptance

- None.

## Guardrails

- The date predicate compares only the first 10 characters of `started_at`; it never parses times
  or time zones.
- Filter bounds are never adjusted or widened to make a range match.
- The predicate never reads or alters the log file.
