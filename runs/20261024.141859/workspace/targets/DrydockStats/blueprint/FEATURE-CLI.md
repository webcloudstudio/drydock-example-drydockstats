# FEATURE: Command-Line Delivery

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Defines the command-line interface, error behavior, and Go test runner. |
| Depends On  | FEATURE-LogIntake.md, FEATURE-Reports.md |
| Provides    | ./cmd/ddkstat, bin/test.sh |
| Consumes    | log parsing, record validation, duplicate detection, target and date filtering, summary, commands, targets, models, daily, failures CSV reports |

## Purpose

Deliver the `ddkstat` command and project test runner. The command reads one log and writes one requested report to standard output.

## Command

```text
ddkstat <report> [--target <name>] [--since <YYYY-MM-DD>] [--until <YYYY-MM-DD>] <log>
```

The report name comes first, options follow it, and exactly one log path comes last. Supported reports are `summary`, `commands`, `targets`, `models`, `daily`, and `failures`.

`--target` matches `job.target` exactly. `--since` includes records on or after the given start date; `--until` includes records on or before it. Dates must be real calendar dates written `YYYY-MM-DD`.

## Error behavior

- A malformed invocation, unknown report or option, missing option value, missing or extra log path, or invalid date exits `2`.
- A log that cannot be opened or read exits `1`.
- Skipped malformed or duplicate records do not fail the command; successful processing exits `0`.
- On a command-level error, write no standard output and one diagnostic line to standard error. Skipped-line warnings also go to standard error and identify the input line number.
- Only `main` calls `os.Exit`; CLI behavior is exposed through `Run(args []string, stdout, stderr io.Writer) int`.

## Test runner

Provide a POSIX-compatible `bin/test.sh` that runs `go vet ./...` and `go test ./...` from the application root, and exits nonzero if either command fails. Keep tests in Go, using `testing`; cover parsing, filtering, every report, and each exit-code class with logs created by the tests.

## Programmatic Acceptance

=== AC cli-error-contract ===
Intent: Missing arguments and an unreadable log produce their specified exit statuses with no standard output.
Cite: sources/INSTRUCTIONS.md:87-95

import os
import subprocess
import sys

missing = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(missing.stdout)
print(missing.stderr, file=sys.stderr)
assert missing.returncode == 0
=== END AC cli-error-contract ===

=== AC test-runner ===
Intent: The project test runner exits successfully after vet and all Go tests pass.
Cite: sources/INSTRUCTIONS.md:97-99

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC test-runner ===

=== AC full-conformance ===
Intent: The terminal story passes all 20 supplied grader checks.
Suite: full

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC full-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute files under `sources/`.
- Never write to the log being read.
- Diagnostics go to standard error; command-level errors produce no standard output.
