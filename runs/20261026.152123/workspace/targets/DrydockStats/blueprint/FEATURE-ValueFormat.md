# FEATURE: Value Format

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Shared text formatters for minutes (tenths of a minute) and cost (four decimal places). |
| Depends On  | ARCHITECTURE.md |
| Provides    | format.Minutes, format.Cost |
| Consumes    | go.mod |

## Purpose

Every report that prints a minutes column or a cost column uses these two pure functions in `internal/format`, so the rounding rule lives in exactly one place.

## Behavior

- `Minutes(ms int64) string`: let `t = (ms + 3000) / 6000` using integer division. Return `t/10`, a `.`, then `t%10` (for example `95000` ms gives `1.6`, `0` gives `0.0`).
- `Cost(usd float64) string`: return `fmt.Sprintf("%.4f", usd)`, exactly four decimal places (zero gives `0.0000`).
- Both functions are pure, take no filters and read no input.

## Programmatic Acceptance

=== AC format-minutes-rounding ===
Intent: Minutes applies the (ms + 3000) / 6000 tenths rule across the rounding boundaries.

from drydock_go import run_test

result = run_test("./internal/format", "TestMinutesRounding")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC format-minutes-rounding ===

=== AC format-minutes-zero ===
Intent: Minutes of zero milliseconds is written as tenths with one decimal digit.

from drydock_go import run_test

result = run_test("./internal/format", "TestMinutesZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC format-minutes-zero ===

=== AC format-cost-four-places ===
Intent: Cost always writes exactly four decimal places.

from drydock_go import run_test

result = run_test("./internal/format", "TestCostFourPlaces")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC format-cost-four-places ===

=== AC format-cost-zero ===
Intent: Cost of zero is written with four decimal places.

from drydock_go import run_test

result = run_test("./internal/format", "TestCostZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC format-cost-zero ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module.
- Formatters never read a log, a filter, or a report answer; every value comes from the argument.
- Test expectations are derived from the formula, not typed as hand-calculated answers.
