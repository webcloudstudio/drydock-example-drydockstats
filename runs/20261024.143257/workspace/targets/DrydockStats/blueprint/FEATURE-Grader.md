# FEATURE: Grader Run

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Run the supplied grader `sh sources/run.sh` against the completed application as the terminal gate. |
| Depends On  | ARCHITECTURE.md, FEATURE-Executable.md, FEATURE-TestSuite.md, FEATURE-ReportEmptyMatch.md, FEATURE-ReportFailures.md |
| Provides    | terminal verification gate |

## Purpose

This is the terminal story and runs last. Its acceptance runs the supplied conformance harness `sources/run.sh` from the application root and requires exit status zero. The harness exit status is the sole verdict; its output is printed for diagnosis only.

## Behavior

- `sh sources/run.sh` is invoked from the application root with no arguments.
- The harness builds `./cmd/ddkstat` and runs its own known-answer checks.
- The acceptance asserts `returncode == 0`. It does not re-type the check total or any known answer.
- If checks fail, the captured output is the diagnosis. The application is fixed, not the grader.

## Programmatic Acceptance

=== AC GR-001 ===
Intent: The terminal gate: the supplied grader sh sources/run.sh builds ddkstat and exits zero, and its exit status is the sole verdict.
Suite: full
Requires: executable=sh; scope=test
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["sh", "sources/run.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC GR-001 ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute `sources/run.sh` or `sources/llm_sample.jsonl`.
- Do not re-type the 20-check total or any known answer in this acceptance.
- Application code never reads `sources/run.sh`.
