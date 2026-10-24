# Feature: Entrypoint

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | The ddkstat process entry point delegates to cli.Run and exits with its code. |
| Depends On  | ARCHITECTURE.md |
| Provides    | ddkstat <report> [options] <log> |
| Consumes    | — |

## Purpose

Provide `cmd/ddkstat/main.go`, the only place that calls `os.Exit`. It passes `os.Args[1:]`, `os.Stdout`, and `os.Stderr` to `cli.Run` and exits with the returned code.

## Programmatic Acceptance

=== AC entrypoint-builds ===
Intent: The ddkstat main package builds with go build ./cmd/ddkstat.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
result = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC entrypoint-builds ===

=== AC entrypoint-no-args-exit-2 ===
Intent: Invoking ddkstat with no arguments exits 2 with empty stdout.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
build = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(build.stderr)
assert build.returncode == 0
result = subprocess.run([out], capture_output=True, text=True, stdin=subprocess.DEVNULL)
print(result.stderr)
assert result.returncode == 2
assert result.stdout == ""
=== END AC entrypoint-no-args-exit-2 ===

## User Acceptance

- None.

## Guardrails

- Only `cmd/ddkstat/main.go` calls `os.Exit`.
