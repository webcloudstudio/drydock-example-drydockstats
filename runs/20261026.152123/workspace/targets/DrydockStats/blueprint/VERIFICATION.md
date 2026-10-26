# FEATURE: Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Terminal release gate that runs the supplied grader against the finished ddkstat tool. |
| Depends On  | FEATURE-ReportSummary.md, FEATURE-ReportCommands.md, FEATURE-ReportTargets.md, FEATURE-ReportModels.md, FEATURE-ReportDaily.md, FEATURE-ReportFailures.md, FEATURE-BuildExecutable.md, FEATURE-TestScript.md |
| Consumes    | ddkstat summary, ddkstat commands, ddkstat targets, ddkstat models, ddkstat daily, ddkstat failures |

## Purpose

This is the last story in the build. Every report, filter, skip rule and exit code is built by earlier stories. This story proves that they work together by running the supplied grader.

## Behavior

- `sources/run.sh` and `sources/llm_sample.jsonl` are staged in the build directory, unmodified.
- `sh sources/run.sh` runs from the application root.
- It builds `./cmd/ddkstat`, runs its 20 checks over the sample log, and exits zero only when the build succeeds and every check passes.
- A failure is repaired in the member behavior. The staged assets and the known answers are never edited, and no report answer is hardcoded in application code.

## Programmatic Acceptance

=== AC verification-grader ===
Intent: The supplied grader sources/run.sh runs unfiltered against the finished application and exits zero.
Suite: full
Requires: executable=go; scope=test

import subprocess
import sys
from pathlib import Path

assert Path("sources/run.sh").is_file()
assert Path("sources/llm_sample.jsonl").is_file()

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
    timeout=1800,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC verification-grader ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate or substitute any file under `sources/`.
- Application code does not read `sources/run.sh` and does not embed a report row, count or total.
- The log being read is never written to.
