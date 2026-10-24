# FEATURE: Acceptance Gate

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Assembly gate: the supplied grader must exit zero as the release verdict. |
| Depends On  | FEATURE-TestRunner.md, FEATURE-CliUsageErrors.md, FEATURE-CliSuccessPartial.md |
| Provides    | sources/run.sh verdict |
| Consumes    | sources/run.sh |

## Purpose

Terminal assembly story. After every member story is verified, run `sh sources/run.sh` from the application root. Its exit status is the sole acceptance verdict.

## Behavior

- Runs once every member story is closed and verified.
- Executes `sh sources/run.sh` from the application root.
- The grader's exit status is the verdict; its output is not parsed.
- Does not edit anything under `sources/`.

## Programmatic Acceptance

=== AC acceptance-grader-passes ===
Intent: The supplied grader exits 0 against the completed application.
Suite: full

import subprocess
import sys
import os

from pathlib import Path

app = Path(os.getcwd())
result = subprocess.run(["sh", "sources/run.sh"], cwd=app, capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC acceptance-grader-passes ===

## User Acceptance

- None.

## Guardrails

- Never modify, rewrite, regenerate, or substitute any file under `sources/`.
- Never read `sources/run.sh` from application code.
