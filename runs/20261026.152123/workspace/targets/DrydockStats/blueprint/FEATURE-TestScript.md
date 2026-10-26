# FEATURE: Test Script

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Provides a POSIX bin/test.sh that runs go vet and go test and exits zero only when both pass. |
| Depends On  | ARCHITECTURE.md |
| Provides    | bin/test.sh |
| Consumes    | go.mod |

## Purpose

One command, `sh bin/test.sh`, gates every behavioral story and the Go test suite.

## Behavior

- `bin/test.sh` is POSIX `sh`: it uses `set -eu` and no bashisms.
- It changes to the application root, derived from the script's own path, so it works from any current directory.
- It runs `go vet ./...` and then `go test ./...`.
- It exits zero only when both commands pass, and non-zero otherwise.
- It runs the toolchain installed on the machine and never a vendored copy.

## Programmatic Acceptance

=== AC test-script-syntax ===
Intent: bin/test.sh exists and parses as POSIX sh.

import subprocess
from pathlib import Path

script = Path("bin/test.sh")
assert script.is_file()
result = subprocess.run(["sh", "-n", str(script)], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-script-syntax ===

=== AC test-script-passes ===
Intent: sh bin/test.sh exits zero when go vet and every automated test pass.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True, timeout=1200)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-script-passes ===

=== AC test-script-any-cwd ===
Intent: bin/test.sh finds the application root regardless of the caller's working directory.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

script = os.path.abspath("bin/test.sh")
elsewhere = tempfile.mkdtemp()
result = subprocess.run(["sh", script], capture_output=True, text=True, cwd=elsewhere, timeout=1200)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-script-any-cwd ===

## User Acceptance

- None.

## Guardrails

- The script contains no bashisms.
- The script never edits source files or the log.
