# Feature: Log Failed

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Classifies a kept record as failed when its status is any value other than succeeded, and exposes the fields the failures report needs. |
| Depends On  | FEATURE-LogSkip.md |
| Provides    | Record.Failed, nullable returncode, timed_out |
| Consumes    | record parser |

## Purpose

Identify failed calls for the summary, command, target, model, daily, and failures reports.

## Behavior

- A record is failed when `status` is any value other than `succeeded`.
- The record exposes `started_at`, `job.target`, `job.command_name`, its model (per FEATURE-LogModel), `result.returncode` as nullable, and `result.timed_out` for the failures report.
- `returncode` is absent or `null` when the record has no return code; `timed_out` is treated as false when absent.

## Programmatic Acceptance

=== AC log-failed-status-classification ===
Intent: A record whose status is not succeeded is counted as failed in the summary, and a succeeded record is not.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

ok = '{"schema_version":1,"execution_id":"f-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
bad = '{"schema_version":1,"execution_id":"f-2","started_at":"2026-10-14T06:38:54.176Z","status":"timeout","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":null,"timed_out":true,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(ok + "\n" + bad + "\n")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[0] == "2", out
    assert fields[2] == "1", out
    assert fields[3] == "1", out
=== END AC log-failed-status-classification ===

=== AC log-failed-returncode-and-timeout ===
Intent: The failures report shows an empty returncode for a null return code and true for timed_out.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

bad = '{"schema_version":1,"execution_id":"f-3","started_at":"2026-10-14T06:38:54.176Z","status":"failed","job":{"command_name":"diagnose","target":"T","llm":"claude","model":"sonnet"},"result":{"returncode":null,"timed_out":true,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(bad + "\n")
    code, out, err = cli_runner("failures", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[4] == "", out
    assert fields[5] == "true", out
=== END AC log-failed-returncode-and-timeout ===

## User Acceptance

- None.

## Guardrails

- A failed record is any status other than `succeeded` (Rule 5).
