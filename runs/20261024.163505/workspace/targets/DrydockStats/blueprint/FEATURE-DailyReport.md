# Feature: Daily Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the per-start-date CSV report over the filtered log records. |
| Depends On  | ARCHITECTURE.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-LogNumbers.md, FEATURE-LogFailed.md, FEATURE-FormatCSV.md, FEATURE-FormatMinutes.md |
| Provides    | daily report |
| Consumes    | record aggregates, CSV and format helpers |

## Purpose

Write one CSV row per start date over the records that pass the active `--target`, `--since`, and `--until` filters. A record's start date is the first 10 characters of `started_at`. Tokens per row are input plus cached plus output.

## Output

Header row, in this exact order:

```
date,calls,failed,tokens,minutes
```

- One data row per distinct start date, ordered ascending by byte value of the `YYYY-MM-DD` string.
- `tokens` is input plus cached plus output tokens summed over the date's records.
- `minutes` is the summed `elapsed_ms` formatted by the minutes rule (tenths of a minute).
- A final `TOTAL` row over all filtered records.
- A report with no matching records writes the header and a `TOTAL,0,0,0,0.0` row.

## Programmatic Acceptance

=== AC daily-report-whole-sample ===
Intent: The daily report over the full sample matches the supplied known answer for every date and the TOTAL row, parsed as CSV.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("daily", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["date", "calls", "failed", "tokens", "minutes"], rows[0]
assert rows[-1][0] == "TOTAL" and rows[-1][1] == "33" and rows[-1][2] == "2", rows[-1]
dates = [r[0] for r in rows[1:-1]]
assert dates == sorted(dates), dates
=== END AC daily-report-whole-sample ===

=== AC daily-report-since-until ===
Intent: With --since and --until both set to 2026-10-15 the daily report contains only that date and its TOTAL row, parsed as CSV.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("daily", "--since", "2026-10-15", "--until", "2026-10-15", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 3, rows
assert rows[1][0] == "2026-10-15", rows[1]
assert rows[2][0] == "TOTAL" and rows[2][1] == rows[1][1], rows
=== END AC daily-report-since-until ===

=== AC daily-report-zero-match ===
Intent: A filter matching no record writes the header and a zero TOTAL row with exit 0.
Suite: scoped
Requires: executable=python3; scope=test

import csv
import io
import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("daily", "--target", "NoSuchTarget", sample)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1][0] == "TOTAL" and rows[1][1] == "0", rows[1]
=== END AC daily-report-zero-match ===

## User Acceptance

- None.

## Guardrails

- Dates are compared by byte value of the `YYYY-MM-DD` string only; no time-zone conversion is applied.
- Values are computed from the log read, never from a hardcoded row or count.
