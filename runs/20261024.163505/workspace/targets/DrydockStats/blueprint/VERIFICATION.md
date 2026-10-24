# FEATURE: Release Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Terminal release gate that runs the supplied DrydockStats grader against the completed application. |

## Programmatic Acceptance

=== AC release-grader-passes ===
Intent: The completed ddkstat application passes the supplied grader, whose exit status is the sole acceptance verdict.
Suite: full
Requires: executable=sh; scope=test

import subprocess
import sys

result = subprocess.run(["sh", "sources/run.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC release-grader-passes ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- The grader's exit status is the verdict; standard output is not asserted here.
- The application must build with `go build ./cmd/ddkstat` from the application root using only the standard library.
