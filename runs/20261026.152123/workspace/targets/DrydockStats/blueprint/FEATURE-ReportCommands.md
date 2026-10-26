# FEATURE: Report Commands

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The commands report: one CSV row per Drydock command with a TOTAL row. |
| Depends On  | FEATURE-CliExit.md, FEATURE-ValueFormat.md |
| Provides    | ddkstat commands |
| Consumes    | cli.Register, format.Minutes, logread.Record |

## Purpose

`ddkstat commands [options] <log>` shows how often each Drydock command called an LLM and what it consumed.

## Behavior

- Header: `command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes`.
- One row per `job.command_name`, ordered by calls descending, then command ascending by byte value.
- Final row `TOTAL,...` over all filtered records.
- Minutes use `format.Minutes` of the group's summed `elapsed_ms`.
- When no record matches, the output is the header and `TOTAL,0,0,0,0,0,0.0`.
- Text values are written unquoted and unescaped.
- `internal/report` holds the aggregate and write functions; `internal/cli` registers the writer as `commands`.

## Programmatic Acceptance

=== AC commands-ordering ===
Intent: Rows are ordered by calls descending then command ascending, including ties.

from drydock_go import run_test

result = run_test("./internal/report", "TestCommandsOrdering")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC commands-ordering ===

=== AC commands-total ===
Intent: The TOTAL row sums calls, failures, tokens and minutes over all filtered records.

from drydock_go import run_test

result = run_test("./internal/report", "TestCommandsTotal")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC commands-total ===

=== AC commands-empty ===
Intent: With no matching record the report is the header plus a zero TOTAL row.

from drydock_go import run_test

result = run_test("./internal/report", "TestCommandsEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC commands-empty ===

=== AC commands-write ===
Intent: The writer emits the declared header and columns.

from drydock_go import run_test

result = run_test("./internal/report", "TestCommandsWrite")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC commands-write ===

=== AC commands-run-via-cli ===
Intent: cli.Run dispatches the commands report over a test-written log, exits zero and writes parseable CSV.

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunCommandsReport")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC commands-run-via-cli ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no database; the log is never written.
- No row, count or total is hardcoded; all values come from the log being read.
- Text values are never quoted or escaped.
