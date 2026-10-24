# Feature: Test Suite

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Provide bin/test.sh, a POSIX sh launcher that runs go vet and go test over the application root. |
| Depends On  | ARCHITECTURE.md, FEATURE-Entrypoint.md, FEATURE-CommandParse.md, FEATURE-UsageErrors.md, FEATURE-ReadErrors.md, FEATURE-LogRead.md, FEATURE-LogSkip.md, FEATURE-ReportPipeline.md, FEATURE-SummaryReport.md, FEATURE-CommandsReport.md, FEATURE-TargetsReport.md, FEATURE-ModelsReport.md, FEATURE-DailyReport.md, FEATURE-FailuresReport.md |
| Provides    | bin/test.sh |
| Consumes    | all implementation packages |

## Purpose

Provide `bin/test.sh`, a POSIX `sh` script that runs `go vet ./...` and `go test ./...` from the application root and exits zero only when both pass. The suite covers each report, each filter, each skip rule, and each exit code using logs written by the tests.

## Programmatic Acceptance

=== AC test-script-exits-zero ===
Intent: bin/test.sh runs go vet and go test and exits 0 when both pass.
Suite: scoped
Requires: executable=sh; scope=test

import subprocess

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-script-exits-zero ===

=== AC test-script-is-posix-sh ===
Intent: bin/test.sh begins with a POSIX sh shebang so it runs under any sh.
Suite: scoped
Requires: executable=python3; scope=test

import os

with open(os.path.join("bin", "test.sh"), "rb") as f:
    first = f.readline()
assert first.startswith(b"#!/bin/sh"), first
=== END AC test-script-is-posix-sh ===

## User Acceptance

- None.

## Guardrails

- bin/test.sh is a launcher only; it runs the Go toolchain and never a vendored copy.
- The suite writes its own logs under temporary directories and never writes to `sources/`.
