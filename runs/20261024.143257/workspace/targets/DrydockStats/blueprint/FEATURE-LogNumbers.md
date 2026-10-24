# Feature: Log Numbers

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Zero-fill absent numeric fields and derive the failed flag on each normalized record. |
| Depends On  | FEATURE-LogReader.md |
| Provides    | Record numeric fields, Record.Failed, Record.ReturnCode, Record.TimedOut |
| Consumes    | internal/log Normalize, internal/log Record |

## Purpose

Make every numeric field safe to sum, and make failure status explicit, so reports never special-case missing data.

## Behavior

- An absent or null `input_tokens`, `cached_input_tokens`, `output_tokens`, `cost_usd`, or `elapsed_ms` counts as zero.
- A record with no `result.stats` object is still a record, with zero tokens, cost, and duration.
- `Failed` is set when `status` is not `succeeded`.
- `ReturnCode` is kept as nil when `returncode` is null or absent, and as its integer value otherwise.
- `TimedOut` is kept from `result.timed_out`.
- Cost is kept as `float64`. Durations are kept as integer milliseconds so the formatter can apply the minutes and cost rounding rules.

## Programmatic Acceptance

=== AC log-numbers-zero-fill ===
Intent: Absent or null numeric fields, and a record with no result.stats object, yield zero values.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestNumbersZeroFill")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-numbers-zero-fill ===

=== AC log-numbers-failed ===
Intent: A record whose status is not succeeded has Failed set, and a succeeded record does not.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestNumbersFailedFlag")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-numbers-failed ===

=== AC log-numbers-returncode ===
Intent: A null or absent returncode yields a nil ReturnCode, and a present one yields its integer value.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestNumbersReturnCode")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-numbers-returncode ===

## User Acceptance

- None.

## Guardrails

- Numeric fields are never left nil; absent values become zero.
