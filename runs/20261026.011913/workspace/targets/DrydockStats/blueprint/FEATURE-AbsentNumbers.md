# FEATURE: Absent Numbers

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Absent, null, or wrongly typed token counts, cost, and durations count as zero without dropping the record. |
| Depends On  | FEATURE-DuplicateIds.md |
| Provides    | logread.Record numeric fields |
| Consumes    | logread.Read |

## Purpose

The reporter sums token counts, cost, and elapsed time. Real Drydock logs omit `result.stats` for
some calls, write `cost_usd: null` for providers that report no cost, and could carry a field of an
unexpected JSON type. None of these may drop a record or fail the run.

## Behavior

- `result.stats.input_tokens`, `cached_input_tokens`, `output_tokens`, `cost_usd`, and `elapsed_ms`
  are each read independently.
- A field that is absent or JSON `null` counts as zero.
- A field present with a JSON type other than number or `null` (string, boolean, array, object)
  counts as zero and the record is kept (decision Q-004, `zero-keep`).
- A record with no `result.stats` object at all is still a record; every numeric field is zero.
- `result.returncode` that is JSON `null` or absent is recorded as absent, not as zero.
- `result.timed_out` that is absent or not a boolean is recorded as `false`.
- Numbers are decoded with `encoding/json`; integer fields tolerate values written like `12.0`
  only if they are whole numbers, otherwise they count as zero.
- Nothing here writes to the log being read.

## Required Go Tests

The build writes these tests in package `./internal/logread`, each asserting the behavior its name
states against a log or line written by the test itself:

- `TestAbsentStatsIsStillRecord`
- `TestNullCostCountsZero`
- `TestAbsentNumbersCountZero`
- `TestWrongTypedNumbersCountZero`
- `TestReturncodeNullOrAbsentIsAbsent`
- `TestTimedOutDefaultsFalse`

## Programmatic Acceptance

=== AC absent-numbers-no-stats ===
Intent: A record with no result.stats object is kept as a record with zero numeric fields.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestAbsentStatsIsStillRecord")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC absent-numbers-no-stats ===

=== AC absent-numbers-null-cost ===
Intent: A null cost_usd counts as zero while the other numeric fields of the record are still summed.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestNullCostCountsZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC absent-numbers-null-cost ===

=== AC absent-numbers-absent-fields ===
Intent: Each absent token count, cost, and duration counts as zero independently of the others.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestAbsentNumbersCountZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC absent-numbers-absent-fields ===

=== AC absent-numbers-wrong-type ===
Intent: A numeric field with a JSON type other than number or null counts as zero and keeps the record.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestWrongTypedNumbersCountZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC absent-numbers-wrong-type ===

=== AC absent-numbers-returncode ===
Intent: A null or absent returncode is recorded as absent rather than zero.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestReturncodeNullOrAbsentIsAbsent")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC absent-numbers-returncode ===

=== AC absent-numbers-timed-out ===
Intent: An absent timed_out defaults to false.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestTimedOutDefaultsFalse")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC absent-numbers-timed-out ===

## User Acceptance

- None.

## Guardrails

- A record is never dropped because a numeric field is missing, null, or of the wrong type.
- Never write to the log being read.
- No value is hardcoded; every number comes from the log being read.
