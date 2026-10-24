# Feature: Log Model

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Attribute each record to the executed model, falling back to the requested model. |
| Depends On  | FEATURE-LogReader.md |
| Provides    | Record.Model derivation in internal/log |
| Consumes    | internal/log Normalize, internal/log Record |

## Purpose

Give every report one consistent model attribution per record, preferring the model that actually executed.

## Behavior

- A record's Model is `result.stats.model` when that field is present and non-empty.
- Otherwise the Model is `job.model`.
- The derivation lives inside the record normalizer, so every report reads Model from the same field.

## Programmatic Acceptance

=== AC log-model-executed ===
Intent: A record with a non-empty result.stats.model is attributed to that model.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestModelPrefersExecuted")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-model-executed ===

=== AC log-model-fallback ===
Intent: A record with no result.stats.model, or an empty one, is attributed to job.model.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/log", "TestModelFallsBackToRequested")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC log-model-fallback ===

## User Acceptance

- None.

## Guardrails

- Model is derived only in the normalizer, never in a report.
