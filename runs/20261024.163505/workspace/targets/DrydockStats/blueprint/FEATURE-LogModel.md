# Feature: Log Model

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Attributes each kept record to its executed model when present and non-empty, otherwise to its requested model. |
| Depends On  | FEATURE-LogSkip.md |
| Provides    | Record.Model |
| Consumes    | record parser |

## Purpose

Name the model that ran each call, so the models and failures reports can group and display it.

## Behavior

- A record's model is `result.stats.model` when that field is present and a non-empty string.
- Otherwise the model is `job.model`.
- The model is exposed on each kept record as `Record.Model`.

## Programmatic Acceptance

=== AC log-model-prefers-executed ===
Intent: A record with result.stats.model reports that executed model in the models report rather than job.model.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

line = '{"schema_version":1,"execution_id":"m-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"requested-model"},"result":{"returncode":0,"timed_out":false,"stats":{"model":"executed-model","cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(line + "\n")
    code, out, err = cli_runner("models", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = out.splitlines()
    fields = rows[1].split(",")
    assert fields[0] == "claude", out
    assert fields[1] == "executed-model", out
=== END AC log-model-prefers-executed ===

=== AC log-model-falls-back-to-requested ===
Intent: A record without result.stats.model reports its job.model in the models report.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

line = '{"schema_version":1,"execution_id":"m-2","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"requested-model"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(line + "\n")
    code, out, err = cli_runner("models", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[1] == "requested-model", out
=== END AC log-model-falls-back-to-requested ===

## User Acceptance

- None.

## Guardrails

- The executed model is preferred whenever it is present and non-empty (Rule 4).
