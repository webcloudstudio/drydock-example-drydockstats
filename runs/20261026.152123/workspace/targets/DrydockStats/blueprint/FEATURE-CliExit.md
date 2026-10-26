# FEATURE: Cli Exit

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Assemble the ddkstat pipeline in cli.Run and map usage and read failures to exit codes 2 and 1. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliDates.md, FEATURE-LogNumbers.md, FEATURE-LogModel.md, FEATURE-FilterTarget.md, FEATURE-FilterDates.md |
| Provides    | cli.Run, cli.Register |
| Consumes    | cli.ParseArgs, logread.Read, filter.ByTarget, filter.ByDateRange |

## Purpose

`cli.Run` is the single entry point behind `cmd/ddkstat/main.go`. It runs the pipeline: parse arguments, read the
log, apply filters, dispatch to the registered report writer. Report stories plug in through a registry so `Run`
never imports them.

## Interface

- `func Run(args []string, stdout, stderr io.Writer) int`
- `func Register(name string, fn func(w io.Writer, records []logread.Record, skipped int) error)`

## Rules

1. Usage error: write one usage-summary line to stderr, nothing to stdout, return 2.
2. Log cannot be opened or read: write one diagnostic line to stderr, nothing to stdout, return 1. The log is only
   ever opened for reading.
3. Otherwise read records (skip warnings go to stderr), apply `ByTarget` then `ByDateRange`, and call the writer
   registered under the report name. `skipped` is the whole-file count, unaffected by filters.
4. Report output is buffered and flushed to stdout only when the writer succeeds. A writer error writes one line to
   stderr, nothing to stdout, and returns 1.
5. A report name with no registered writer writes nothing and returns 0 until its report story lands.
6. Skipped lines never change the exit code.
7. `Run` never calls `os.Exit`; only `cmd/ddkstat/main.go` does.

## Programmatic Acceptance

=== AC cli-exit-usage-code ===
Intent: A usage error returns exit code 2.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunUsageErrorExit2")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-exit-usage-code ===

=== AC cli-exit-usage-stdout-empty ===
Intent: A usage error writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunUsageErrorWritesNothingToStdout")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-exit-usage-stdout-empty ===

=== AC cli-exit-missing-log-code ===
Intent: An unreadable log returns exit code 1.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunMissingLogExit1")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-exit-missing-log-code ===

=== AC cli-exit-missing-log-stdout-empty ===
Intent: An unreadable log writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunMissingLogWritesNothingToStdout")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-exit-missing-log-stdout-empty ===

=== AC cli-exit-skipped-lines-ok ===
Intent: A log containing skipped lines still exits 0.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunSkippedLinesStillExit0")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-exit-skipped-lines-ok ===

=== AC cli-exit-filter-order ===
Intent: The registered writer receives records narrowed by the target and date filters and the whole-file skipped count.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestRunFilterOrder")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-exit-filter-order ===

## User Acceptance

- None.

## Guardrails

- Never write to the log being read.
- A failing command writes nothing to stdout and exactly one diagnostic line to stderr.
- Diagnostics go to stderr and are never compared by the grader.
- `Run` does not import any report package; reports register themselves.
