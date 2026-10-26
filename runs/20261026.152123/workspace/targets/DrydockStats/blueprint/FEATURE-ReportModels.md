# FEATURE: Report Models

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The models report: one CSV row per provider and executed-model pair with a TOTAL row. |
| Depends On  | FEATURE-CliExit.md, FEATURE-ValueFormat.md, FEATURE-LogModel.md |
| Provides    | ddkstat models |
| Consumes    | cli.Register, format.Cost, logread.Record |

## Purpose

`ddkstat models [options] <log>` shows which providers and models ran the calls and what they consumed and cost.

## Behavior

- Header: `provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd`.
- One row per pair of provider (`job.llm`) and `Record.Model` (the executed model when present and non-empty, otherwise the requested model), ordered by provider then model ascending by byte value.
- Final row is `TOTAL,,...` with an empty model column, over all filtered records.
- Cost uses `format.Cost`. There is no minutes column.
- When no record matches, the output is the header and a TOTAL row of zeros with `0.0000` cost.
- `internal/report` holds the aggregate and write functions; `internal/cli` registers the writer as `models`.

## Programmatic Acceptance

=== AC models-grouping-executed-model ===
Intent: Records group under the executed model, falling back to the requested model.

from drydock_go import run_test

result = run_test("./internal/report", "TestModelsGroupingUsesExecutedModel")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC models-grouping-executed-model ===

=== AC models-ordering ===
Intent: Rows are ordered by provider then model ascending in byte order.

from drydock_go import run_test

result = run_test("./internal/report", "TestModelsOrdering")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC models-ordering ===

=== AC models-total-empty-model-column ===
Intent: The TOTAL row leaves the model column empty and sums every numeric column.

from drydock_go import run_test

result = run_test("./internal/report", "TestModelsTotalEmptyModelColumn")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC models-total-empty-model-column ===

=== AC models-empty ===
Intent: With no matching record the report is the header plus a zero TOTAL row.

from drydock_go import run_test

result = run_test("./internal/report", "TestModelsEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC models-empty ===

=== AC models-write ===
Intent: The writer emits the declared header and columns.

from drydock_go import run_test

result = run_test("./internal/report", "TestModelsWrite")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC models-write ===

=== AC models-run-via-cli ===
Intent: cli.Run dispatches the models report over a test-written log, exits zero and writes parseable CSV.

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunModelsReport")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC models-run-via-cli ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no database; the log is never written.
- No row, count, cost or total is hardcoded; all values come from the log being read.
- Provider and model names are written unquoted and unescaped.
