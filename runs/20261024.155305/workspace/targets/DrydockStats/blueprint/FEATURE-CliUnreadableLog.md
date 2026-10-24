# FEATURE: Unreadable Log Exit Code

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Maps an unopenable or unreadable log to exit code 1 with no standard output. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliParse.md, FEATURE-LogDecode.md |
| Provides    | unreadable log exit 1 |
| Consumes    | Run, log decoder |

## Purpose

Tell the operator that the log itself could not be used, distinct from a bad invocation.

## Behavior

- When the log cannot be opened, or a read fails, `Run` returns exit code `1`.
- Exactly one diagnostic line is written to standard error.
- Nothing is written to standard output.
- The full report is built in memory and written to stdout only after the whole log has been read successfully.

## Programmatic Acceptance

=== AC unreadable-missing-log ===
Intent: A log path that does not exist exits 1 with empty stdout and one stderr line.
Suite: scoped
Requires: executable=python3; scope=test

import subprocess
import os
import sys
import tempfile

app = os.getcwd()
with tempfile.TemporaryDirectory() as d:
    bin_path = os.path.join(d, "ddkstat")
    missing = os.path.join(d, "no-such-log.jsonl")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run([bin_path, "summary", missing], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 1
    assert result.stdout == ""
    assert len([line for line in result.stderr.splitlines() if line.strip()]) == 1
=== END AC unreadable-missing-log ===

=== AC unreadable-directory-log ===
Intent: A log path that names a directory exits 1 with empty stdout.
Suite: scoped
Requires: executable=python3; scope=test

import subprocess
import os
import sys
import tempfile

app = os.getcwd()
with tempfile.TemporaryDirectory() as d:
    bin_path = os.path.join(d, "ddkstat")
    subprocess.run(["go", "build", "-o", bin_path, "./cmd/ddkstat"], cwd=app, check=True)
    result = subprocess.run([bin_path, "summary", d], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 1
    assert result.stdout == ""
=== END AC unreadable-directory-log ===
