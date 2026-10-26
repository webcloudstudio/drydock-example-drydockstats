# FEATURE: Cli Args

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Parse the report name, options, and log path of a ddkstat invocation into a validated Options value. |
| Depends On  | ARCHITECTURE.md |
| Provides    | cli.ParseArgs, cli.Options, cli.UsageError |
| Consumes    | go.mod |

## Purpose

`internal/cli` turns the raw argument list into an `Options` value or a usage error. It performs no I/O and
never exits the process. Date value validation is owned by FEATURE-CliDates.md and layered onto this parser.

## Interface

```text
ddkstat <report> [--target <name>] [--since <YYYY-MM-DD>] [--until <YYYY-MM-DD>] <log>
```

- `type Options struct { Report, Target, Since, Until, LogPath string }`
- `func ParseArgs(args []string) (Options, error)`; every failure is a `*UsageError` (`errors.As`-detectable).
- `args` excludes the program name.

## Rules

1. The first argument is the report and must be one of `summary`, `commands`, `targets`, `models`, `daily`, `failures`.
2. After the report come zero or more options, each of `--target`, `--since`, `--until` followed by exactly one value.
3. Exactly one log path follows; it is the last argument.
4. A repeated option takes its last value.
5. An option placed before the report name is a usage error (it is not a known report).
6. An unknown report, an unknown option, a missing option value, a missing log path, or an extra positional
   argument is a usage error.
7. Parsing uses a hand-rolled loop or a local `flag.FlagSet` with `flag.ContinueOnError`; it never touches the
   global flag set and never calls `os.Exit`.

## Programmatic Acceptance

=== AC cli-args-report-and-log ===
Intent: A report name followed by a log path parses into Options with that report and path.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseReportAndLog")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-report-and-log ===

=== AC cli-args-options ===
Intent: The --target, --since and --until options populate Options.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseOptions")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-options ===

=== AC cli-args-repeated-option ===
Intent: A repeated option takes its last value.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseRepeatedOptionLastWins")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-repeated-option ===

=== AC cli-args-option-before-report ===
Intent: An option placed before the report name is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseOptionBeforeReportIsUsageError")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-option-before-report ===

=== AC cli-args-unknown-report ===
Intent: An unknown report name is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseUnknownReport")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-unknown-report ===

=== AC cli-args-unknown-option ===
Intent: An unknown option is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseUnknownOption")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-unknown-option ===

=== AC cli-args-missing-value ===
Intent: An option with no value is a usage error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseMissingValue")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-missing-value ===

=== AC cli-args-log-path-count ===
Intent: A missing log path and an extra log path are both usage errors.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseLogPathCount")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cli-args-log-path-count ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module.
- ParseArgs performs no I/O and never calls `os.Exit`.
- Option order is fixed: report first, then options, then the single log path.
