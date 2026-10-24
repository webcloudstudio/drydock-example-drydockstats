# Feature: Log Skip

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Parses each non-blank log line as a JSON object, skips malformed, incomplete, and duplicate records with a stderr warning, and exposes the whole-file skipped count. |
| Depends On  | FEATURE-LogRead.md |
| Provides    | record parser with whole-file skipped count |
| Consumes    | logfile line reader |

## Purpose

Turn raw log lines into kept records and a whole-file skipped count. Malformed and duplicate lines are never fatal; they are skipped, counted, and reported on stderr.

## Behavior

- Each non-blank line is parsed as a JSON object. A line that is not a JSON object is skipped.
- A line is skipped when any of `execution_id`, `started_at`, `status`, `job.command_name`, or `job.target` is absent or not a non-empty string.
- A line whose `execution_id` already appeared on an earlier kept record is skipped. The first occurrence is kept.
- Each skipped line writes exactly one warning to standard error naming its one-based line number. Nothing is written to standard output.
- Skipping never changes the exit code.
- The skipped count covers the whole file and is independent of any filter.

## Programmatic Acceptance

=== AC log-skip-counts-and-exit ===
Intent: Non-object, incomplete, and malformed lines are skipped, counted in the summary, and do not change exit status.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

good = '{"schema_version":1,"execution_id":"k-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
lines = [good, "not json", "[1,2,3]", '{"execution_id":"k-2","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan"}}']
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write("\n".join(lines) + "\n")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[0] == "1", out
    assert fields[1] == "3", out
=== END AC log-skip-counts-and-exit ===

=== AC log-skip-duplicate-first-kept ===
Intent: A duplicate execution_id is skipped on its later occurrence and the first occurrence is kept.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

first = '{"schema_version":1,"execution_id":"d-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
second = '{"schema_version":1,"execution_id":"d-1","started_at":"2026-10-15T06:37:54.176Z","status":"failed","job":{"command_name":"build","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":1,"timed_out":false,"stats":{"cost_usd":0.9,"elapsed_ms":9000,"input_tokens":9,"cached_input_tokens":0,"output_tokens":9}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(first + "\n" + second + "\n")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    fields = out.splitlines()[1].split(",")
    assert fields[0] == "1", out
    assert fields[1] == "1", out
    assert fields[2] == "1", out
    assert fields[3] == "0", out
=== END AC log-skip-duplicate-first-kept ===

=== AC log-skip-warning-per-line ===
Intent: Each skipped line writes exactly one stderr warning naming its line number, and nothing is written to stdout for skips.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

good = '{"schema_version":1,"execution_id":"w-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":0.1,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write("garbage\n" + good + "\n")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    warnings = [line for line in err.splitlines() if line.strip()]
    assert len(warnings) == 1, err
    assert "1" in warnings[0], err
    assert "garbage" not in out, out
=== END AC log-skip-warning-per-line ===

## User Acceptance

- None.

## Guardrails

- Skipping never changes the exit code (Rule 2).
- Skipped-line warnings go to standard error only and are never compared by the grader.
- The skipped count is computed over the whole file regardless of filters (Rule 6).
