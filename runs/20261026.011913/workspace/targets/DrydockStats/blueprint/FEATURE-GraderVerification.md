# FEATURE: Grader Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Terminal assembly that proves the supplied grader and the project test entry point both pass against the finished application. |
| Depends On  | FEATURE-TestEntryPoint.md, FEATURE-SummaryReport.md, FEATURE-DateFilter.md, FEATURE-ModelAttribution.md, FEATURE-UnreadableLog.md |
| Provides    | sh sources/run.sh passing |
| Consumes    | bin/test.sh, cmd/ddkstat |

## Purpose

With every member story closed, the finished `ddkstat` is graded by the supplied `sources/run.sh`, which builds `./cmd/ddkstat` and makes 20 checks of exit code and standard output over `sources/llm_sample.jsonl`. This story is the integration seam: any defect the grader reveals is fixed in the application, never in `sources/`.

## Assembly and Intent

- Run `sh sources/run.sh` unfiltered from the application root; it must exit zero.
- Run `sh bin/test.sh` unfiltered; it must exit zero.
- Application code never reads `sources/run.sh` and never embeds a report answer; every value comes from the log being read.
- The module uses the Go standard library only, with no `require` entries in `go.mod`.
- `sources/run.sh` and `sources/llm_sample.jsonl` are read-only scoring assets and remain unmodified.

## Programmatic Acceptance

=== AC grader-passes ===
Intent: The supplied grader builds ddkstat and passes all of its checks, so it exits zero.
Suite: full
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["sh", "sources/run.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC grader-passes ===

=== AC project-suite-passes ===
Intent: The project test entry point exits zero over the complete application.
Suite: full
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC project-suite-passes ===

=== AC application-independent-of-grader ===
Intent: Application code does not reference the grader and the module has no third-party requirements.

from pathlib import Path

sources = []
for root in ("cmd", "internal"):
    for p in Path(root).rglob("*.go"):
        if not p.name.endswith("_test.go"):
            sources.append(p)
assert sources
for p in sources:
    assert "run.sh" not in p.read_text(encoding="utf-8"), str(p)
mod = Path("go.mod").read_text(encoding="utf-8").splitlines()
assert not [line for line in mod if line.strip().startswith("require")]
=== END AC application-independent-of-grader ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Do not hardcode a report row, a count, or a total in application code.
- Never write to the log being read.
