# FEATURE: Filter Records by Target

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Keep only log records whose job.target equals the --target value exactly. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogSkip.md, FEATURE-CommandParse.md |
| Provides    | target filter predicate |
| Consumes    | record parser, parsed options |

## Purpose

Every report applies the same target predicate. When `--target <name>` is given, only records whose
`job.target` string equals `<name>` byte for byte are kept. Without the option, every kept record
passes.

## Behavior

- The predicate is `record.Target == targetOption` using exact byte comparison. No trimming, case
  folding, or Unicode normalization.
- When `--target` is absent, the predicate accepts every record.
- The predicate is exposed as a function value that every report aggregation applies before
  aggregating. It is never applied by individual reports on their own.
- The `skipped` count is not affected by this predicate; it is computed by the parser over the whole
  file (see FEATURE-SkippedCount.md).
- `--target` with an empty string value matches only records whose target is the empty string. Such
  records cannot exist because `job.target` must be a non-empty string, so the result is zero matches
  and exit 0.

## Programmatic Acceptance

=== AC filter-target-exact-match ===
Intent: A record whose job.target equals the --target value is kept and a record differing only in case or trailing space is not.
Requires: executable=python3; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

lines = [
    {"schema_version": 1, "execution_id": "a-1", "started_at": "2026-10-14T01:00:00.000Z",
     "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}},
    {"schema_version": 1, "execution_id": "a-2", "started_at": "2026-10-14T01:00:01.000Z",
     "status": "succeeded", "job": {"command_name": "build", "target": "alpha", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}},
    {"schema_version": 1, "execution_id": "a-3", "started_at": "2026-10-14T01:00:02.000Z",
     "status": "succeeded", "job": {"command_name": "build", "target": "Alpha ", "llm": "claude", "model": "haiku"},
     "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}},
]

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    for obj in lines:
        fh.write(json.dumps(obj) + "\n")

try:
    code, out, err = cli_runner("summary", "--target", "Alpha", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = list(__import__("csv").reader(out.splitlines()))
    header = rows[0]
    data = rows[1]
    assert header[0] == "records", header
    assert data[0] == "1", data
finally:
    os.remove(path)
=== END AC filter-target-exact-match ===

=== AC filter-target-absent-keeps-all ===
Intent: Without --target every valid record is kept by the summary report.
Requires: executable=python3; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

records = []
for i, target in enumerate(["Alpha", "Beta", "Gamma"]):
    records.append({"schema_version": 1, "execution_id": f"b-{i}", "started_at": "2026-10-14T02:00:00.000Z",
                    "status": "succeeded", "job": {"command_name": "build", "target": target, "llm": "claude", "model": "haiku"},
                    "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                               "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}})

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    for obj in records:
        fh.write(json.dumps(obj) + "\n")

try:
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    import csv
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][0] == "3", rows
finally:
    os.remove(path)
=== END AC filter-target-absent-keeps-all ===

=== AC filter-target-no-match-exit-zero ===
Intent: A --target value matching no record yields a zero-count summary with exit 0.
Requires: executable=python3; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

record = {"schema_version": 1, "execution_id": "c-1", "started_at": "2026-10-14T03:00:00.000Z",
          "status": "succeeded", "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
          "result": {"returncode": 0, "timed_out": False, "stats": {"elapsed_ms": 1000, "input_tokens": 1,
                     "cached_input_tokens": 0, "output_tokens": 1, "cost_usd": 0.1}}}

fd, path = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

try:
    code, out, err = cli_runner("summary", "--target", "NoSuchTarget", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    import csv
    rows = list(csv.reader(out.splitlines()))
    assert rows[1][0] == "0", rows
finally:
    os.remove(path)
=== END AC filter-target-no-match-exit-zero ===

## User Acceptance

- None.

## Guardrails

- The filter compares job.target exactly; it never matches on case, whitespace, or substring.
- The filter never reads or alters the log file.
- `--target` is the only option this predicate consumes.
