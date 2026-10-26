# FEATURE: Log Model

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Attribute each record to the executed model, falling back to the requested model. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogSkip.md |
| Provides    | logread.Record.Model |
| Consumes    | logread.Read |

## Purpose

Reports group by model, and the model that actually ran is more truthful than the one requested.

## Behavior

`Record.Model` is `result.stats.model` when that value is present and a non-empty string; otherwise it is `job.model`. This covers an absent value, an empty string, a non-string value, and a record with no `result.stats` object. `job.model` that is absent yields an empty `Model`; the record is still kept.

## Required Go tests

TestModelPrefersExecuted, TestModelFallsBackWhenAbsent, TestModelFallsBackWhenEmpty, TestModelFallsBackWithoutStats.

## Programmatic Acceptance

=== AC model-prefers-executed ===
Intent: result.stats.model is used when present and non-empty.

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelPrefersExecuted")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-prefers-executed ===

=== AC model-fallback-absent ===
Intent: job.model is used when result.stats.model is absent.

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelFallsBackWhenAbsent")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-fallback-absent ===

=== AC model-fallback-empty ===
Intent: job.model is used when result.stats.model is an empty string.

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelFallsBackWhenEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-fallback-empty ===

=== AC model-fallback-no-stats ===
Intent: job.model is used when the record has no result.stats object.

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelFallsBackWithoutStats")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-fallback-no-stats ===

## User Acceptance

- None.

## Guardrails

- The model is never taken from any field other than `result.stats.model` and `job.model`.
- Standard library only.
