# Feature: Log Skip

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Classify each non-blank line and skip malformed or duplicate records, warning on stderr. |
| Depends On  | FEATURE-LogReader.md |
| Provides    | internal/log Normalize, internal/log Record, skipped count |
| Consumes    | internal/log ReadLines, internal/log Line |

## Purpose

Separate usable records from unusable lines so every report aggregates only well-formed, unique records, while counting what was dropped.

## Behavior

- A line is skipped when it is not a JSON object.
- A line is skipped when any of `execution_id`, `started_at`, `status`, `job.command_name`, or `job.target` is absent or not a non-empty string.
- A line is skipped when its `execution_id` already appeared on an earlier line. The first occurrence is kept.
- Each skipped line writes one warning to stderr naming its line number. Stdout is never written.
- The stage returns the kept records and the total skipped count for the whole file.
- Skipping never changes the exit code.
- Filters are applied after this stage, so the skipped count is unaffected by filters.

## Programmatic Acceptance

=== AC log-skip-malformed ===
Intent: A non-object line and a line missing a required field are both counted as skipped.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestNormalizeSkipsMalformed")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-skip-malformed ===

=== AC log-skip-duplicate ===
Intent: The first occurrence of an execution_id is kept and later duplicates are skipped.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestNormalizeKeepsFirstDuplicate")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-skip-duplicate ===

=== AC log-skip-warning ===
Intent: Each skipped line writes exactly one stderr warning naming its line number and nothing to stdout.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestNormalizeWarnsOncePerSkip")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-skip-warning ===

## User Acceptance

- None.

## Guardrails

- Stdout is never written by this stage.
- Skipping never changes the exit code.
