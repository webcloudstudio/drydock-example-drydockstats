# Feature: Whole-File Skipped Count

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Counts skipped log lines over the whole file, independent of any filter. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogSkip.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md |
| Provides    | whole-file skipped count |
| Consumes    | record parser, target and date filters |

## Purpose

The `skipped` column of the summary report counts skipped lines in the entire log, regardless of `--target`, `--since`, or `--until`. Filters narrow the records a report aggregates; they never change the skipped count.

## Behavior

- The skipped count is the number of non-blank lines that the record parser rejects, computed over every line of the file.
- The count is computed before filters are applied, so `--target` and the date options have no effect on it.
- Blank lines are not counted (see FEATURE-LogRead.md).
- Running the same log with and without filters yields the same skipped count.

## Programmatic Acceptance

=== AC skipped-ignores-target ===
Intent: The skipped count is identical with and without a --target filter.
Requires: executable=python3; scope=test

import subprocess

from drydock_go import cli_runner

log = "/tmp/ddkstat-skip-fixture.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')
    f.write("not json\n")
    f.write('{"schema_version":1,"execution_id":"e2","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"B","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')

code_all, out_all, err_all = cli_runner("summary", log)
code_f, out_f, err_f = cli_runner("summary", "--target", "A", log)
print(err_all)
print(err_f)
assert code_all == 0, (code_all, out_all, err_all)
assert code_f == 0, (code_f, out_f, err_f)

skipped_all = out_all.split("\n")[1].split(",")[1]
skipped_f = out_f.split("\n")[1].split(",")[1]
assert skipped_all == "1", repr(skipped_all)
assert skipped_f == skipped_all, (skipped_all, skipped_f)
=== END AC skipped-ignores-target ===

=== AC skipped-ignores-date ===
Intent: The skipped count is identical with and without date filters.
Requires: executable=python3; scope=test

import subprocess

from drydock_go import cli_runner

log = "/tmp/ddkstat-skip-date-fixture.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')
    f.write("not json\n")
    f.write('{"schema_version":1,"execution_id":"e2","started_at":"2026-10-15T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')

code_all, out_all, err_all = cli_runner("summary", log)
code_f, out_f, err_f = cli_runner("summary", "--since", "2026-10-15", "--until", "2026-10-15", log)
print(err_all)
print(err_f)
assert code_all == 0, (code_all, out_all, err_all)
assert code_f == 0, (code_f, out_f, err_f)

skipped_all = out_all.split("\n")[1].split(",")[1]
skipped_f = out_f.split("\n")[1].split(",")[1]
assert skipped_all == "1", repr(skipped_all)
assert skipped_f == skipped_all, (skipped_all, skipped_f)
=== END AC skipped-ignores-date ===

## User Acceptance

- None.

## Guardrails

- The skipped count never depends on --target, --since, or --until.
