# FEATURE: Model Attribution

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Each record is attributed to the executed model, falling back to the requested model, under its provider. |
| Depends On  | FEATURE-AbsentNumbers.md |
| Provides    | logread.Record.Model |
| Consumes    | logread.Read |

## Purpose

`job.model` is the requested model alias (for example `sonnet`), while `result.stats.model` is the
model that actually ran (for example `claude-sonnet-5-5`). Reports group by what ran, and fall back
to the request only when the log does not say.

## Behavior

- A record's model is `result.stats.model` when it is present and a non-empty string.
- Otherwise the model is `job.model`.
- A `result.stats` object without `model`, a missing `result.stats`, a `null` model, or an empty
  string model all fall back to `job.model`.
- The provider of a record is `job.llm`.
- The model and provider are exposed on the record for the models and failures reports.

## Required Go Tests

The build writes these tests in package `./internal/logread`, each using lines written by the test:

- `TestModelPrefersExecuted`
- `TestModelFallsBackWhenStatsAbsent`
- `TestModelFallsBackWhenEmpty`
- `TestProviderIsJobLLM`

## Programmatic Acceptance

=== AC model-attribution-executed ===
Intent: A record with a non-empty result.stats.model is attributed to that model rather than job.model.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelPrefersExecuted")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-attribution-executed ===

=== AC model-attribution-no-stats ===
Intent: A record without result.stats is attributed to job.model.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelFallsBackWhenStatsAbsent")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-attribution-no-stats ===

=== AC model-attribution-empty ===
Intent: A null, absent, or empty result.stats.model falls back to job.model.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestModelFallsBackWhenEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-attribution-empty ===

=== AC model-attribution-provider ===
Intent: The provider of a record is read from job.llm.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/logread", "TestProviderIsJobLLM")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC model-attribution-provider ===

## User Acceptance

- None.

## Guardrails

- The executed model is never replaced by the requested model when the executed model is present and non-empty.
- Attribution reads only the log record; no model name is hardcoded.
