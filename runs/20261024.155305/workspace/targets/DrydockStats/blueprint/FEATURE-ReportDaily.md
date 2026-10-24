# FEATURE: Daily Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the per-date CSV report over filtered records, ordered by start date ascending, with tokens as input plus cached plus output and a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md, FEATURE-LogRequiredFields.md, FEATURE-LogDuplicates.md, FEATURE-LogNumeric.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-ReportFormat.md |
| Provides    | daily report |
| Consumes    | records, FormatMinutes, WriteCSV |

## Purpose

Write the `daily` report: one row per start date over filtered records, followed by a `TOTAL` row.

## Report Contract

Header:

```
date,calls,failed,tokens,minutes
```

- Rows are grouped by start date, the first 10 characters of `started_at`.
- Rows are ordered by `date` ascending.
- `failed` counts filtered records whose `status` is not `succeeded`.
- `tokens` is `input_tokens` plus `cached_input_tokens` plus `output_tokens`, summed over the group.
- `minutes` is written by `FormatMinutes` over summed `elapsed_ms`.
- The final row is `TOTAL` over all filtered records. With no matches it is `TOTAL,0,0,0,0.0`.

## Programmatic Acceptance

=== AC daily-ordering-and-total ===
Intent: Daily rows group by start date ascending, tokens sum input plus cached plus output, and the TOTAL row covers all filtered records.
Suite: scoped

import csv
import io
import os
import sys
import tempfile

from drydock_go import cli_runner

def rec(eid, started, status="succeeded"):
    return ('{"schema_version":1,"execution_id":"%s","started_at":"%s",'
            '"status":"%s","job":{"command_name":"build","target":"T","llm":"claude","model":"haiku"},'
            '"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":10,"cached_input_tokens":20,'
            '"output_tokens":30,"cost_usd":0.1,"elapsed_ms":1000}}}') % (eid, started, status)

lines = [
    rec("1", "2026-10-15T09:00:00.000Z"),
    rec("2", "2026-10-14T09:00:00.000Z"),
    rec("3", "2026-10-14T10:00:00.000Z", status="failed"),
]

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "log.jsonl")
    with open(path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")
    code, out, err = cli_runner("daily", path)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[0] == ["date", "calls", "failed", "tokens", "minutes"]
    dates = [r[0] for r in rows[1:-1]]
    assert dates == ["2026-10-14", "2026-10-15"]
    first_group = rows[1]
    assert first_group[1] == "2"
    assert first_group[2] == "1"
    assert first_group[3] == "120"
    total = rows[-1]
    assert total[0] == "TOTAL"
    assert total[1] == "3"
    assert total[3] == "180"
=== END AC daily-ordering-and-total ===

## User Acceptance

- None.

## Guardrails

- The daily report must not read `sources/run.sh` or embed any expected answer.
- Dates are grouped as text of the first 10 characters of `started_at`, never re-parsed into local time.
