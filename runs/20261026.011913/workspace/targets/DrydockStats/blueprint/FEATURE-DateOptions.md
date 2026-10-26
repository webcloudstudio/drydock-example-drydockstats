# FEATURE: Date Options

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The --since and --until option values are accepted only as real calendar dates written YYYY-MM-DD. |
| Depends On  | FEATURE-CommandLine.md |
| Provides    | cli date validation |
| Consumes    | ddkstat <report> [--target] [--since] [--until] <log> |

## Purpose

Date filters compare against the first 10 characters of `started_at`, so an option value must be
exactly that shape and must name a day that exists.

## Behavior

- A valid value is exactly ten characters: four-digit year, `-`, two-digit month, `-`, two-digit day.
- The value must also be a real calendar date, checked with `time.Parse` and the `2006-01-02`
  layout; leap days are honored (`02-29` is valid only in a leap year).
- Rejected: month 13, day 30 in February, non-padded forms such as `2026-1-5`, compact forms such
  as `20261014`, trailing or leading characters, an empty value, and any non-date text.
- The validator is a pure function in `internal/cli` returning the date text and an error; the
  parser maps its error to a usage error (exit handling is owned by FEATURE-UsageErrors.md).

## Required Go Tests

The build writes these tests in package `./internal/cli`:

- `TestParseDateAcceptsRealDates`
- `TestParseDateRejectsMalformed`
- `TestParseDateRejectsImpossibleDates`

## Programmatic Acceptance

=== AC date-options-accepts-real ===
Intent: Real YYYY-MM-DD dates, including a leap day, are accepted by the date validator.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseDateAcceptsRealDates")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-options-accepts-real ===

=== AC date-options-rejects-malformed ===
Intent: Non-padded, compact, padded, empty, and non-date values are rejected by the date validator.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseDateRejectsMalformed")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-options-rejects-malformed ===

=== AC date-options-rejects-impossible ===
Intent: Well-formed but impossible dates such as month 13 or 30 February are rejected.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseDateRejectsImpossibleDates")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-options-rejects-impossible ===

## User Acceptance

- None.

## Guardrails

- Only the standard library is used for date validation.
- A date is never silently normalized; an invalid value is rejected, not corrected.
