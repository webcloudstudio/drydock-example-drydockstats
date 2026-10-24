# Feature: CLI Parse

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Parse the report name, options, and the single log path in the required order. |
| Depends On  | ARCHITECTURE.md, FEATURE-Executable.md |
| Provides    | internal/cli Run dispatch, report name table, option grammar |
| Consumes    | internal/log ReadLines |

## Purpose

Turn the command line into a validated report request, or a usage error, before any log is read.

## Behavior

- The first argument is the report name. Options follow it. Exactly one log path is last.
- The six report names are `summary`, `commands`, `targets`, `models`, `daily`, and `failures`.
- A local `flag.FlagSet` with `flag.ContinueOnError` parses the options. Tests use `bytes.Buffer` for output.
- `--target`, `--since`, and `--until` are declared so they parse. Their semantics belong to the filter stories.
- This stage owns the report dispatch table. Each report story registers its name here.

## Programmatic Acceptance

=== AC cli-parse-report-names ===
Intent: Each of the six report names is accepted as the first argument.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseAcceptsSixReports")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-parse-report-names ===

=== AC cli-parse-order ===
Intent: The report name must come first, options follow, and the single log path comes last.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseArgumentOrder")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-parse-order ===

=== AC cli-parse-options ===
Intent: The --target, --since, and --until options parse with their values.
Requires: executable=python3; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseOptionValues")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-parse-options ===

## User Acceptance

- None.

## Guardrails

- Parsing never reads the log file.
- Only `cmd/ddkstat` calls `os.Exit`.
