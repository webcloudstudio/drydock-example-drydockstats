# FEATURE: Log Decode

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Decode the JSON Lines log into records, one JSON object per non-blank line. |
| Depends On  | ARCHITECTURE.md |
| Provides    | log record decoder |
| Consumes    | - |

## Programmatic Acceptance

=== AC log-decode-ignores-blank ===
Intent: Blank lines are ignored and not counted as skipped.
Suite: scoped
Requires: executable=go; scope=test

import json
import subprocess
import sys

log = '{"execution_id":"a"}\n\n'
print(log)
assert True
=== END AC log-decode-ignores-blank ===

## User Acceptance

- None.

## Guardrails

- Failure to open or read the file is returned as an error, never as a skipped line.
