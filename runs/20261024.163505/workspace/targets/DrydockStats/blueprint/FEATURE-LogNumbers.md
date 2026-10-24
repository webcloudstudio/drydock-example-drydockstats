# Feature: Log Numbers

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Treats absent or null token counts, cost, and elapsed time as zero, and keeps a record without result.stats as a record. |
| Depends On  | FEATURE-LogSkip.md |
| Provides    | Record token, cost, and elapsed accessors defaulting to zero |
| Consumes    | record parser |

## Purpose

Let every aggregate sum numeric stats safely, so that partial or absent statistics never fail a run or drop a record.

## Behavior

- `input_tokens`, `cached_input_tokens`, `output_tokens`, `cost_usd`, and `elapsed_ms` each default to zero when absent or `null`.
- A record with no `result.stats` object is still a kept record and contributes zero to every numeric aggregate.
- Accessors return these values to the aggregation code; no numeric field causes a skip.

## Programmatic Acceptance

=== AC log-numbers-null-and-absent-are-zero ===
Intent: Null and absent numeric stats count as zero in the summary totals and the record is still counted.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

line = '{"schema_version":1,"execution_id":"n-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":null,"elapsed_ms":null,"input_tokens":5,"output_tokens":7}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(line + "\n")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[0] == "1", out
    assert fields[6] == "5", out
    assert fields[7] == "0", out
    assert fields[8] == "7", out
    assert fields[9] == "0.0000", out
    assert fields[10] == "0.0", out
=== END AC log-numbers-null-and-absent-are-zero ===

=== AC log-numbers-no-stats-object ===
Intent: A record with no result.stats object is still counted as a record with zero numeric totals.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

line = '{"schema_version":1,"execution_id":"n-2","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(line + "\n")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[0] == "1", out
    assert fields[6] == "0", out
    assert fields[9] == "0.0000", out
=== END AC log-numbers-no-stats-object ===

## User Acceptance

- None.

## Guardrails

- Missing or null numeric stats are zero, never a skip reason (Rule 3).
