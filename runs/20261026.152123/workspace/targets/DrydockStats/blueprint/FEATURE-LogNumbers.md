# FEATURE: Log Numbers

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Treat absent, null or non-numeric token, cost and duration values as zero. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogSkip.md |
| Provides    | logread.Record numeric normalization |
| Consumes    | logread.Read |

## Purpose

Aggregation sums these fields, so each record must carry a usable number even when the log omits it.

## Behavior

For `result.stats.input_tokens`, `cached_input_tokens`, `output_tokens`, `cost_usd` and `elapsed_ms`:

- absent counts as zero;
- JSON `null` counts as zero;
- a value of the wrong JSON type (for example a string or boolean) counts as zero;
- a record with no `result.stats` object at all, or no `result` object, is still a record with all five values zero.

`result.returncode` that is `null` or absent is represented as no return code (`ReturnCode` is nil). `result.timed_out` absent counts as false.

## Required Go tests

TestNumbersAbsentAreZero, TestNumbersNullAreZero, TestNumbersNonNumericAreZero, TestNoStatsObjectStillRecord, TestReturnCodeNullIsNil.

## Programmatic Acceptance

=== AC numbers-absent ===
Intent: Absent numeric fields count as zero.

from drydock_go import run_test

result = run_test("./internal/logread", "TestNumbersAbsentAreZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numbers-absent ===

=== AC numbers-null ===
Intent: Null numeric fields, including a null cost_usd, count as zero.

from drydock_go import run_test

result = run_test("./internal/logread", "TestNumbersNullAreZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numbers-null ===

=== AC numbers-non-numeric ===
Intent: A non-numeric value in a numeric stats field counts as zero.

from drydock_go import run_test

result = run_test("./internal/logread", "TestNumbersNonNumericAreZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numbers-non-numeric ===

=== AC numbers-no-stats ===
Intent: A record with no result.stats object is still a kept record with zero numbers.

from drydock_go import run_test

result = run_test("./internal/logread", "TestNoStatsObjectStillRecord")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numbers-no-stats ===

=== AC numbers-returncode-null ===
Intent: A null or absent returncode is represented as no return code.

from drydock_go import run_test

result = run_test("./internal/logread", "TestReturnCodeNullIsNil")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numbers-returncode-null ===

## User Acceptance

- None.

## Guardrails

- A bad numeric value never causes a line to be skipped; only the required string fields do.
- Standard library only.
