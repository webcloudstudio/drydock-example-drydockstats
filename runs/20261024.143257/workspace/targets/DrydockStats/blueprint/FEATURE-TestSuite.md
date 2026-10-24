# FEATURE: Test Suite

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Finalize bin/test.sh as the POSIX verification gate and add end-to-end coverage across reports, filters, skip rules, and exit codes. |
| Depends On  | FEATURE-Executable.md, FEATURE-LogReader.md, FEATURE-LogSkip.md, FEATURE-LogModel.md, FEATURE-LogNumbers.md, FEATURE-CliParse.md, FEATURE-CliTargetFilter.md, FEATURE-CliDateFilters.md, FEATURE-CliUsageErrors.md, FEATURE-CliReadErrors.md, FEATURE-ReportFormat.md, FEATURE-ReportSummary.md, FEATURE-ReportCommands.md, FEATURE-ReportTargets.md, FEATURE-ReportModels.md, FEATURE-ReportDaily.md, FEATURE-ReportFailures.md, FEATURE-ReportEmptyMatch.md |
| Provides    | bin/test.sh verification gate, end-to-end test in cmd/ddkstat |

## Purpose

`bin/test.sh` is the project's verification gate. It runs `go vet ./...` and `go test ./...` from the application root and exits zero only when both pass. The end-to-end test under `cmd/ddkstat` writes logs in test code and exercises each report, each filter, each skip rule, and each exit code through `cli.Run`.

## Behavior

- `bin/test.sh` is POSIX `sh`, runs `go vet ./...` then `go test ./...`, and exits non-zero if either fails.
- The end-to-end test lives in `cmd/ddkstat` as a `_test.go` file and uses `cli.Run` with `bytes.Buffer` writers.
- The end-to-end test covers every report, every filter, every skip rule, and every exit code.
- The test writes its own logs; it never reads `sources/`.

## Programmatic Acceptance

=== AC TS-001 ===
Intent: bin/test.sh exits zero when go vet and every automated test pass.
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC TS-001 ===

=== AC TS-002 ===
Intent: bin/test.sh is valid POSIX sh syntax.
Requires: executable=sh; scope=test

import subprocess
import sys

result = subprocess.run(["sh", "-n", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC TS-002 ===

=== AC TS-003 ===
Intent: The end-to-end test is present as a _test.go file under cmd/ddkstat.

import glob
import sys

test_files = glob.glob("cmd/ddkstat/*_test.go")
print(test_files, file=sys.stderr)
assert len(test_files) >= 1, test_files
=== END AC TS-003 ===

=== AC TS-004 ===
Intent: The end-to-end test package under cmd/ddkstat passes go test on its own.
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["go", "test", "-count=1", "./cmd/ddkstat"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC TS-004 ===

=== AC TS-005 ===
Intent: go vet passes across the whole application module.
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["go", "vet", "./..."], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC TS-005 ===

## User Acceptance

- None.

## Guardrails

- Do not modify any file under `sources/`.
- Do not read `sources/run.sh` from application code.
- Do not hardcode report rows, counts, or totals in the test suite; derive expected values from the logs the tests write.
