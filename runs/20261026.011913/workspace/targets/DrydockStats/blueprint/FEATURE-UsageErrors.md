# FEATURE: Usage Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Every usage error exits 2 with nothing on standard output and a usage line on standard error. |
| Depends On  | FEATURE-DateOptions.md |
| Provides    | exit code 2 |
| Consumes    | cli date validation |

## Purpose

Operators and scripts distinguish a mistyped command (exit 2) from an unreadable log (exit 1).

## Behavior

- Exit 2, with empty standard output, for: no arguments; an unknown report; an unknown option
  (including `--help`, decision Q-001 `unknown-exit-2`); an option missing its value; an invalid
  date value; a missing log path; more than one log path.
- Usage errors are decided before the log is opened, so a usage error never becomes exit 1 even if
  the named log does not exist.
- One usage line is written to standard error naming each of the six reports (`summary`,
  `commands`, `targets`, `models`, `daily`, `failures`) and each of the three options (`--target`,
  `--since`, `--until`). Standard error is never compared by acceptance.
- A repeated option is not a usage error: the last occurrence wins (decision Q-002).

## Required Go Tests

The build writes this test in package `./internal/cli`, writing the usage line to a `bytes.Buffer`
and checking it names every report and every option:

- `TestUsageLineNamesReportsAndOptions`

## Programmatic Acceptance

=== AC usage-errors-exit-2 ===
Intent: Each kind of usage error exits 2 and writes nothing to standard output, with an existing log path supplied.
Requires: executable=go; scope=test

import json
import os
import tempfile
from drydock_go import cli_runner

fd, log = tempfile.mkstemp(suffix=".jsonl")
with os.fdopen(fd, "w", encoding="utf-8") as f:
    f.write(json.dumps({"execution_id": "a", "started_at": "2026-10-14T00:00:00.000Z",
                        "status": "succeeded",
                        "job": {"command_name": "build", "target": "T"}}) + "\n")

cases = {
    "no-args": [],
    "unknown-report": ["frobnicate", log],
    "missing-log": ["summary"],
    "unknown-option": ["summary", "--colour", "red", log],
    "help-option": ["summary", "--help", log],
    "extra-log": ["summary", log, log],
    "missing-value-target": ["summary", "--target"],
    "missing-value-since": ["summary", "--since"],
    "missing-value-until": ["summary", "--until"],
    "bad-since-month": ["summary", "--since", "2026-13-01", log],
    "bad-until-day": ["summary", "--until", "2026-02-30", log],
    "bad-since-form": ["summary", "--since", "2026-1-5", log],
}
try:
    for name, args in cases.items():
        code, out, err = cli_runner(*args)
        print(name, code)
        print(err)
        assert code == 2, (name, code, out, err)
        assert out == "", (name, out)
finally:
    os.unlink(log)
=== END AC usage-errors-exit-2 ===

=== AC usage-errors-before-open ===
Intent: A usage error takes precedence over an unreadable log, so it exits 2 rather than 1.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

missing = os.path.join(tempfile.gettempdir(), "ddkstat-no-such-log-usage.jsonl")
assert not os.path.exists(missing)

for args in (["frobnicate", missing],
             ["summary", "--since", "2026-13-01", missing],
             ["summary", "--colour", "red", missing]):
    code, out, err = cli_runner(*args)
    print(args, code)
    print(err)
    assert code == 2, (args, code, out, err)
    assert out == "", (args, out)
=== END AC usage-errors-before-open ===

=== AC usage-errors-usage-line ===
Intent: The usage line names every report and every option.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestUsageLineNamesReportsAndOptions")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC usage-errors-usage-line ===

## User Acceptance

- None.

## Guardrails

- A usage error never writes to standard output.
- Diagnostics go to standard error only and are never part of the report contract.
- `--help` is an ordinary unknown option; no help mode exists.
