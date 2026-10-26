# FEATURE: Filter Dates

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Keep only records whose start date falls within an inclusive since/until range. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogNumbers.md, FEATURE-LogModel.md |
| Provides    | filter.ByDateRange |
| Consumes    | logread.Record |

## Purpose

A pure function in `internal/filter` that narrows a record slice by start date. Dates are already validated
`YYYY-MM-DD` text, so string comparison orders them correctly.

## Interface

`func ByDateRange(records []logread.Record, since, until string) []logread.Record`

## Rules

1. A record is kept when `Record.Date >= since` and `Record.Date <= until`; both bounds are inclusive.
2. An empty `since` or `until` is an unset bound and does not restrict.
3. When `since` is later than `until` the result is empty; this is not an error and the exit code stays 0.
4. Input order is preserved and the input slice is not modified.

## Programmatic Acceptance

=== AC filter-dates-since-inclusive ===
Intent: A record dated exactly on the since bound is kept and earlier ones are dropped.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestSinceInclusive")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-dates-since-inclusive ===

=== AC filter-dates-until-inclusive ===
Intent: A record dated exactly on the until bound is kept and later ones are dropped.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestUntilInclusive")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-dates-until-inclusive ===

=== AC filter-dates-both-bounds ===
Intent: With both bounds set only records inside the inclusive range remain.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestRangeBothBounds")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-dates-both-bounds ===

=== AC filter-dates-since-after-until ===
Intent: A since later than until yields an empty result without error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestSinceAfterUntilEmpty")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-dates-since-after-until ===

=== AC filter-dates-bounds-unset ===
Intent: Empty bounds leave every record in place.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestBoundsUnset")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-dates-bounds-unset ===

## User Acceptance

- None.

## Guardrails

- The filter is a pure function: no I/O, no global state.
- It does not depend on any report implementation; reports verify it later through CLI output.
