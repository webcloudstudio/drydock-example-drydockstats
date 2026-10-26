# FEATURE: Command Line

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Parse the ddkstat command line into a report name, the three options, and exactly one trailing log path. |
| Depends On  | ARCHITECTURE.md |
| Provides    | ddkstat <report> [--target <name>] [--since <YYYY-MM-DD>] [--until <YYYY-MM-DD>] <log> |

## Purpose

The command is `ddkstat <report> [--target <name>] [--since <date>] [--until <date>] <log>`. The
report comes first, options follow it, and the log path is last. This feature delivers the
parser and the executable entry point at `./cmd/ddkstat`; date validation, usage-error exit codes,
and unreadable-log handling belong to their own features.

## Behavior

- `<report>` is one of `summary`, `commands`, `targets`, `models`, `daily`, `failures`.
- The accepted options are `--target <name>`, `--since <date>`, and `--until <date>`, each taking
  one value, in any order, between the report and the log path.
- Exactly one log path is the final argument.
- A repeated option keeps its last occurrence.
- `cli.Parse(args []string) (Options, error)` returns the typed `Options` from `ARCHITECTURE.md`,
  and returns a non-nil error for any argument list it cannot accept. `cli.Run` maps that error to
  the usage-error exit code in the usage-errors feature.
- `go build ./cmd/ddkstat` produces the executable using only the standard library.

## Required Unit Tests

The build writes these tests in `internal/cli`, each against real parser output:

| Test | Behavior |
|---|---|
| `TestParseAcceptsEveryReport` | Each of the six report names parses with a log path. |
| `TestParseOptionsFollowReport` | The three options parse in any order and populate `Options`. |
| `TestParseLogPathIsLast` | The final argument becomes `Options.Log`. |
| `TestParseRepeatedOptionLastWins` | A repeated option keeps its last value. |

## Programmatic Acceptance

=== AC cli-parse-every-report ===
Intent: Every one of the six report names with a log path is accepted as a valid command line, so none exits with the usage-error code.
Suite: scoped
Requires: executable=go; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

record = {
    "execution_id": "20261014.020414.582Z-db1fe4ea",
    "started_at": "2026-10-14T02:04:14.582Z",
    "status": "succeeded",
    "job": {"command_name": "analyze", "target": "Toml", "llm": "gemini", "model": "gemini-3-flash"},
}
usage_error = 2
reports = ["summary", "commands", "targets", "models", "daily", "failures"]

with tempfile.TemporaryDirectory() as workdir:
    log = os.path.join(workdir, "llm.jsonl")
    with open(log, "w", encoding="utf-8") as handle:
        handle.write(json.dumps(record) + "\n")
    for report in reports:
        code, out, err = cli_runner(report, log)
        print(report, code, err)
        assert code != usage_error, (report, code, out, err)
=== END AC cli-parse-every-report ===

=== AC cli-parse-options ===
Intent: The three options between the report and the log path, in either order, are accepted and do not exit with the usage-error code.
Suite: scoped
Requires: executable=go; scope=test

import json
import os
import tempfile

from drydock_go import cli_runner

record = {
    "execution_id": "20261015.061132.155Z-a5a1d0d4",
    "started_at": "2026-10-15T06:11:32.155Z",
    "status": "succeeded",
    "job": {"command_name": "analyze", "target": "CommonMark", "llm": "claude", "model": "sonnet"},
}
usage_error = 2

with tempfile.TemporaryDirectory() as workdir:
    log = os.path.join(workdir, "llm.jsonl")
    with open(log, "w", encoding="utf-8") as handle:
        handle.write(json.dumps(record) + "\n")
    orders = [
        ["--target", "CommonMark", "--since", "2026-10-01", "--until", "2026-10-31"],
        ["--until", "2026-10-31", "--since", "2026-10-01", "--target", "CommonMark"],
        ["--since", "2026-10-01"],
        ["--until", "2026-10-31"],
        ["--target", "CommonMark"],
    ]
    for options in orders:
        code, out, err = cli_runner("summary", *options, log)
        print(options, code, err)
        assert code != usage_error, (options, code, out, err)
=== END AC cli-parse-options ===

=== AC cli-parse-unit-tests ===
Intent: The parser unit tests required by this feature exist and pass.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

names = [
    "TestParseAcceptsEveryReport",
    "TestParseOptionsFollowReport",
    "TestParseLogPathIsLast",
    "TestParseRepeatedOptionLastWins",
]
for name in names:
    result = run_test("./internal/cli", name)
    print(name, result.status)
    print(result.stdout)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC cli-parse-unit-tests ===

## User Acceptance

- None.

## Guardrails

- The report name is positional and comes first; the log path is positional and comes last.
- `--help` is not special: it is an unknown option.
- The parser uses a local flag set or hand parser with `ContinueOnError` semantics and never the global flag set; only `cmd/ddkstat` calls `os.Exit`.
- Standard library only.
