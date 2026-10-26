# FEATURE: Cli Dates

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Validate --since and --until values as real YYYY-MM-DD calendar dates during argument parsing. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliArgs.md |
| Provides    | cli.ParseArgs date validation |
| Consumes    | cli.ParseArgs |

## Purpose

A date option must be a real calendar date written `YYYY-MM-DD`. Anything else is a usage error that maps to
exit code 2 in FEATURE-CliExit.md.

## Rules

1. A value is valid only when `time.Parse("2006-01-02", v)` succeeds and the parsed time formatted with the same
   layout equals `v` exactly (rejecting shapes such as `2026-1-5`).
2. `2026-13-01` (bad month), `2026-02-30` (bad day), `2026-1-5` (wrong shape), and an empty value are usage errors.
3. A valid date passes through to `Options.Since` or `Options.Until` unchanged.
4. A missing value after `--since` or `--until` is a usage error.
5. Validation is applied while parsing, so `ParseArgs` returns the usage error before any log is opened.

## Programmatic Acceptance

=== AC cli-dates-valid ===
Intent: A real calendar date is accepted and stored unchanged.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestDateValid")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-dates-valid ===

=== AC cli-dates-invalid-month ===
Intent: A date with an impossible month is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestDateInvalidMonth")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-dates-invalid-month ===

=== AC cli-dates-invalid-day ===
Intent: A date with an impossible day of month is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestDateInvalidDay")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-dates-invalid-day ===

=== AC cli-dates-wrong-shape ===
Intent: A date not written as zero-padded YYYY-MM-DD is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestDateWrongShape")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-dates-wrong-shape ===

=== AC cli-dates-missing-value ===
Intent: A date option with no value is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestDateMissingValue")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-dates-missing-value ===

## User Acceptance

- None.

## Guardrails

- Standard library only.
- A `--since` later than `--until` is not a validation error here; see FEATURE-FilterDates.md.
