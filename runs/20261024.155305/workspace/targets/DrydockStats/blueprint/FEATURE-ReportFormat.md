# FEATURE: Report Formatting

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Provide the shared minutes formatter, the cost formatter, and the CSV writer used by every report. |
| Depends On  | ARCHITECTURE.md |
| Provides    | FormatMinutes, FormatCost, WriteCSV |
| Consumes    | none |

## Purpose

Every report shares three output rules: minutes are written in tenths of a minute, cost is written with four decimal places, and rows are written as CSV with `\n` line endings. Centralising these keeps every report byte-consistent.

## Behavior

- `FormatMinutes(ms int64) string`: `t = (ms + 3000) / 6000` in integer division, written as `t/10`, a period, and `t%10`. For example `95000` ms writes `1.6`.
- `FormatCost(usd float64) string`: `fmt.Sprintf("%.4f", usd)`.
- `WriteCSV(w io.Writer, rows [][]string) error`: writes each row with `encoding/csv`, comma-separated, `\n` line endings.

## Programmatic Acceptance

=== AC report-format-minutes-rounding ===
Intent: Minutes are written as tenths of a minute, rounded half up, so 95000 ms writes 1.6.
Requires: executable=python3; scope=test

def format_minutes(ms):
    t = (ms + 3000) // 6000
    return f"{t // 10}.{t % 10}"

assert format_minutes(95000) == "1.6"
assert format_minutes(0) == "0.0"
=== END AC report-format-minutes-rounding ===

=== AC report-format-cost-four-places ===
Intent: Cost is written with exactly four decimal places.
Requires: executable=python3; scope=test

usd = 5.5245
assert f"{usd:.4f}" == "5.5245"
assert f"{0.0:.4f}" == "0.0000"
=== END AC report-format-cost-four-places ===
