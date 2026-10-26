# FEATURE: Test Entry Point

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | A POSIX-compatible bin/test.sh runs go vet and go test from the application root. |
| Depends On  | FEATURE-CommandsReport.md, FEATURE-TargetsReport.md, FEATURE-ModelsReport.md, FEATURE-DailyReport.md, FEATURE-FailuresReport.md, FEATURE-UsageErrors.md, FEATURE-UnreadableLog.md |
| Provides    | bin/test.sh |
| Consumes    | ddkstat summary, ddkstat commands, ddkstat targets, ddkstat models, ddkstat daily, ddkstat failures |

## Purpose

`bin/test.sh` is the project's single test entry point. It lets an operator or the build prove the application is healthy.

## Behavior

- POSIX `sh` only; no bashisms.
- Changes to the application root (the parent of `bin/`) regardless of the caller's working directory.
- Runs `go vet ./...` and then `go test ./...`, and exits zero only when both pass.
- A failure of either command makes the script exit non-zero.
- The Go test suite covers each report, each filter, each skip rule, and each exit code, using logs the tests write themselves. It never reads `sources/run.sh`.

## Programmatic Acceptance

=== AC test-script-runs-clean ===
Intent: bin/test.sh run from an unrelated directory exits zero, proving it locates the application root and that vet and tests pass.
Requires: executable=go; scope=test

import os
import subprocess
import sys
import tempfile

script = os.path.abspath(os.path.join("bin", "test.sh"))
assert os.path.isfile(script)
elsewhere = tempfile.mkdtemp()
result = subprocess.run(["sh", script], cwd=elsewhere, capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC test-script-runs-clean ===

=== AC test-script-posix-syntax ===
Intent: bin/test.sh parses under POSIX sh.

import os
import subprocess
import sys

script = os.path.abspath(os.path.join("bin", "test.sh"))
result = subprocess.run(["sh", "-n", script], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC test-script-posix-syntax ===

## User Acceptance

- None.

## Guardrails

- The script never uses a vendored toolchain; it runs the installed `go`.
- The script never modifies anything under `sources/`.
