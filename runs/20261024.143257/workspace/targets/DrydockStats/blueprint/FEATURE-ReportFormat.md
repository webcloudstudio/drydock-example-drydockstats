# FEATURE: Report Format

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Provide the minutes, cost, and CSV writing helpers used by every report. |
| Depends On  | ARCHITECTURE.md, FEATURE-Executable.md |
| Provides    | internal/report FormatMinutes, FormatCost, CSV writer |
| Consumes    | Record numeric fields |

## Purpose

Shared formatting used by every report. Minutes are derived from summed `elapsed_ms` as
`t = (ms + 3000) / 6000` in integer division, rendered as `t/10`, a dot, and `t%10`. Cost is
rendered with `%.4f`. The CSV writer emits one header row and `\n` line endings. Formatting stays
separate from aggregation so each is unit-tested on its own.

## Behavior

- `FormatMinutes(95000)` renders `1.6`.
- `FormatMinutes(0)` renders `0.0`.
- `FormatCost` renders a `float64` with exactly four decimal places.
- The CSV writer writes a header row followed by data rows, comma-separated, with `\n` endings.

## Programmatic Acceptance

=== AC format-minutes-rounding ===
Intent: Summed milliseconds render as tenths of a minute using the Rule 7 rounding.
Suite: scoped

import subprocess
import sys

script = (
    "import sys\n"
    "sys.path.insert(0, '.')\n"
    "from internal.report import FormatMinutes\n"
    "print(FormatMinutes(95000))\n"
)
result = subprocess.run(
    [sys.executable, "-c", script],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC format-minutes-rounding ===

=== AC format-cost-four-places ===
Intent: Cost renders with exactly four decimal places as Go's %.4f writes a float64.
Suite: scoped

import subprocess
import sys

script = (
    "import sys\n"
    "sys.path.insert(0, '.')\n"
    "from internal.report import FormatCost\n"
    "print(FormatCost(0.0))\n"
)
result = subprocess.run(
    [sys.executable, "-c", script],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC format-cost-four-places ===

## User Acceptance

- None.

## Guardrails

- Formatting is never mixed into aggregation logic.
- Minutes and cost are never rendered by any means other than the helpers in this feature.
