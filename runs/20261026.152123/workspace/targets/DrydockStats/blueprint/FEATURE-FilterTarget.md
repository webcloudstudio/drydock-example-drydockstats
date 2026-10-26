# FEATURE: Filter Target

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Keep only records whose job.target equals a given name exactly. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogNumbers.md, FEATURE-LogModel.md |
| Provides    | filter.ByTarget |
| Consumes    | logread.Record |

## Purpose

A pure function in `internal/filter` that narrows a record slice by Target. It performs no I/O and knows nothing
about reports; every report receives its output through the CLI pipeline.

## Interface

`func ByTarget(records []logread.Record, name string) []logread.Record`

## Rules

1. A record is kept when `Record.Target == name` exactly; comparison is case-sensitive.
2. An empty `name` means no `--target` was given and all records are returned.
3. No match returns an empty slice, never an error.
4. Input order is preserved and the input slice is not modified.

## Programmatic Acceptance

=== AC filter-target-exact ===
Intent: Only records whose Target equals the name are kept.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestByTargetExact")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-target-exact ===

=== AC filter-target-case-sensitive ===
Intent: Target matching is case-sensitive.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestByTargetCaseSensitive")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-target-case-sensitive ===

=== AC filter-target-no-match ===
Intent: A name matching no record yields an empty result.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestByTargetNoMatch")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-target-no-match ===

=== AC filter-target-unset ===
Intent: An empty name leaves every record in place.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/filter", "TestByTargetUnset")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC filter-target-unset ===

## User Acceptance

- None.

## Guardrails

- The filter is a pure function: no I/O, no global state.
- It does not depend on any report implementation; reports verify it later through CLI output.
