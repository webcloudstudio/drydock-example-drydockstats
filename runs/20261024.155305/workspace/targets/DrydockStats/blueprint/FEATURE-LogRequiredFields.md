# FEATURE: Log Required Fields

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Skip records missing a required field, with a line-numbered warning on stderr. |
| Depends On  | FEATURE-LogDecode.md |
| Provides    | required field validation |
| Consumes    | - |

## Programmatic Acceptance

=== AC log-required-fields-skip ===
Intent: A record missing job.target is counted as skipped and does not appear in the records.
Suite: scoped
Requires: executable=go; scope=test

from pathlib import Path

src = Path("internal/logrecord").read_text()
assert "job.target" in src
=== END AC log-required-fields-skip ===

## User Acceptance

- None.

## Guardrails

- Skipping never changes the exit code.
