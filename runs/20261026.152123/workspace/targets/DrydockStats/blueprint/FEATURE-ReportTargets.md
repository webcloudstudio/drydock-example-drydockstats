# FEATURE: Report Targets

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The targets report: one CSV row per Target with cost and a TOTAL row. |
| Depends On  | FEATURE-CliExit.md, FEATURE-ValueFormat.md |
| Provides    | ddkstat targets |
| Consumes    | cli.Register, format.Minutes, format.Cost, logread.Record |

## Purpose

`ddkstat targets [options] <log>` shows what each Target consumed and cost.

## Behavior

- Header: `target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`.
- One row per `job.target`, ordered by target ascending by byte value.
- Final row `TOTAL,...` over all filtered records, with total cost.
- Cost uses `format.Cost`; minutes use `format.Minutes`.
- When no record matches, the output is the header and a TOTAL row of zeros with `0.0000` cost and `0.0` minutes.
- `internal/report` holds the aggregate and write functions; `internal/cli` registers the writer as `targets`.

## Programmatic Acceptance

=== AC targets-ordering ===
Intent: Rows are ordered by target ascending in byte order.

from drydock_go import run_test

result = run_test("./internal/report", "TestTargetsOrdering")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC targets-ordering ===

=== AC targets-total-cost ===
Intent: The TOTAL row carries the summed cost formatted to four places.

from drydock_go import run_test

result = run_test("./internal/report", "TestTargetsTotalCost")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC targets-total-cost ===

=== AC targets-empty ===
Intent: With no matching record the report is the header plus a zero TOTAL row.

from drydock_go import run_test

result = run_test("./internal/report", "TestTargetsEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC targets-empty ===

=== AC targets-write ===
Intent: The writer emits the declared header and columns.

from drydock_go import run_test

result = run_test("./internal/report", "TestTargetsWrite")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC targets-write ===

=== AC targets-run-via-cli ===
Intent: cli.Run dispatches the targets report over a test-written log, exits zero and writes parseable CSV.

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunTargetsReport")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC targets-run-via-cli ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no database; the log is never written.
- No row, count, cost or total is hardcoded; all values come from the log being read.
- Target names are written unquoted and unescaped.
