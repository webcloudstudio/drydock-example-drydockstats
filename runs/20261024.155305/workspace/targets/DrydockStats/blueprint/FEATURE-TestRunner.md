# FEATURE: Test Runner

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Provide the POSIX bin/test.sh that runs go vet and go test over the whole module. |
| Depends On  | FEATURE-BuildEntry.md, FEATURE-LogDecode.md, FEATURE-LogRequiredFields.md, FEATURE-LogDuplicates.md, FEATURE-LogModel.md, FEATURE-LogNumeric.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-ReportFormat.md, FEATURE-ReportSummary.md, FEATURE-ReportCommands.md, FEATURE-ReportTargets.md, FEATURE-ReportModels.md, FEATURE-ReportDaily.md, FEATURE-ReportFailures.md, FEATURE-ReportEmptyMatch.md, FEATURE-CliParse.md, FEATURE-CliUsageErrors.md, FEATURE-CliUnreadableLog.md, FEATURE-CliSuccessPartial.md |
| Provides    | bin/test.sh |
| Consumes    | go vet ./..., go test ./... |

## Purpose

Provide the project's test gate as a POSIX sh script. The script runs the Go toolchain from the application root and exits zero only when both `go vet ./...` and `go test ./...` pass.

## Behavior

- `bin/test.sh` is POSIX sh with `set -eu`.
- It changes to the application root relative to its own location.
- It runs `go vet ./...`, then `go test ./...`.
- It exits 0 only when both succeed; any failure exits non-zero.
- It invokes the installed Go toolchain and never vendors one.

## Programmatic Acceptance

=== AC test-runner-script-present ===
Intent: bin/test.sh exists and is a POSIX sh script that sets errexit and nounset.
Suite: scoped

from pathlib import Path
import os

script = Path(os.getcwd()) / "bin" / "test.sh"
assert script.is_file()
text = script.read_text(encoding="utf-8")
first_lines = text.splitlines()[:3]
assert any(line.strip() == "set -eu" for line in first_lines)
=== END AC test-runner-script-present ===

=== AC test-runner-passes-on-clean-module ===
Intent: bin/test.sh exits 0 when go vet and go test both pass over the module.
Suite: scoped

import subprocess
import sys
import os

from pathlib import Path

app = Path(os.getcwd())
result = subprocess.run(["sh", str(app / "bin" / "test.sh")], cwd=app, capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC test-runner-passes-on-clean-module ===

## User Acceptance

- None.

## Guardrails

- Never vendor a Go toolchain; always invoke the installed `go`.
- Never modify anything under `sources/`.
