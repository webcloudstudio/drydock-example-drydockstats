# FEATURE: Unreadable Log

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | A log that cannot be opened or read exits 1 with nothing on standard output. |
| Depends On  | FEATURE-LogReading.md, FEATURE-CommandLine.md |
| Provides    | exit code 1 |
| Consumes    | logread.Read, ddkstat <report> [--target] [--since] [--until] <log> |

## Purpose

An operational failure reading the log is reported differently from a usage error so that callers
can tell a bad invocation from a bad environment.

## Behavior

- A log path that does not exist, cannot be opened, or cannot be read exits 1.
- A directory named as the log path is unreadable and exits 1 (decision Q-003, `exit-1`).
- One diagnostic line is written to standard error; standard output stays empty.
- Report output is buffered until the whole log has been read, so a read error after some lines
  were consumed cannot leave partial standard output.
- This applies to every report name.
- The log is only ever opened for reading.

## Required Go Tests

The build writes this test in package `./internal/cli`, using a reader that fails partway through:

- `TestReadErrorLeavesStdoutEmpty`

## Programmatic Acceptance

=== AC unreadable-log-missing-file ===
Intent: A missing log exits 1 with empty standard output for every report.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

missing = os.path.join(tempfile.gettempdir(), "ddkstat-no-such-log-unreadable.jsonl")
assert not os.path.exists(missing)

for report in ("summary", "commands", "targets", "models", "daily", "failures"):
    code, out, err = cli_runner(report, missing)
    print(report, code)
    print(err)
    assert code == 1, (report, code, out, err)
    assert out == "", (report, out)
=== END AC unreadable-log-missing-file ===

=== AC unreadable-log-directory ===
Intent: A directory named as the log exits 1 with empty standard output.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

directory = tempfile.mkdtemp()
try:
    code, out, err = cli_runner("summary", directory)
    print(code)
    print(err)
    assert code == 1, (code, out, err)
    assert out == "", out
finally:
    os.rmdir(directory)
=== END AC unreadable-log-directory ===

=== AC unreadable-log-no-partial-output ===
Intent: A read error partway through the log leaves standard output empty.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestReadErrorLeavesStdoutEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC unreadable-log-no-partial-output ===

## User Acceptance

- None.

## Guardrails

- The log is never opened for writing and is never modified.
- An unreadable log never writes to standard output.
- An unreadable log exits 1, never 0 or 2.
