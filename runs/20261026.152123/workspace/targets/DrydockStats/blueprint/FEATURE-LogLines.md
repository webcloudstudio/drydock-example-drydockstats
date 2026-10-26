# FEATURE: Log Lines

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Read the JSON Lines log into records, ignoring blank lines and classifying failed calls. |
| Depends On  | ARCHITECTURE.md |
| Provides    | logread.Read, logread.Record |
| Consumes    | go.mod |

## Purpose

Turn a JSON Lines stream into `logread.Record` values. Each non-blank UTF-8 line is one JSON object. This story establishes the reader, the `Record` type and the failed-call classification; later stories add skip rules, numeric normalization and model attribution.

## Behavior

- Read with `bufio.Reader` so no line-length limit applies.
- A line that is empty after trimming whitespace is blank: ignored, not counted, but it still advances the 1-based physical line number.
- Each object is decoded into a `Record` with `ExecutionID`, `StartedAt`, `Status`, `Command` (`job.command_name`), `Target` (`job.target`), `Provider` (`job.llm`), token counts, cost, `ElapsedMS`, `ReturnCode` (nullable) and `TimedOut`.
- `Failed` is true when `Status` is not `succeeded`.
- `Date` is the first 10 characters of `StartedAt`; `started_at` is accepted as any non-empty string with no further timestamp validation. A string shorter than 10 characters yields the whole string.

## Required Go tests

TestReadBlankLines, TestReadFailedStatus, TestReadStartDate, TestReadStartedAtNotValidated. Tests build their input as strings and call `logread.Read` with a `bytes.Buffer` warning writer.

## Programmatic Acceptance

=== AC lines-blank ===
Intent: Blank lines are ignored and not counted as records or skips.

from drydock_go import run_test

result = run_test("./internal/logread", "TestReadBlankLines")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC lines-blank ===

=== AC lines-failed-status ===
Intent: A record whose status is not succeeded is classified as failed.

from drydock_go import run_test

result = run_test("./internal/logread", "TestReadFailedStatus")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC lines-failed-status ===

=== AC lines-start-date ===
Intent: The record date is the first 10 characters of started_at.

from drydock_go import run_test

result = run_test("./internal/logread", "TestReadStartDate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC lines-start-date ===

=== AC lines-started-at-lenient ===
Intent: started_at is accepted as any non-empty string without timestamp validation.

from drydock_go import run_test

result = run_test("./internal/logread", "TestReadStartedAtNotValidated")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC lines-started-at-lenient ===

## User Acceptance

- None.

## Guardrails

- Never write to the log being read.
- Blank lines are not counted anywhere but still count toward physical line numbers.
- Standard library only.
