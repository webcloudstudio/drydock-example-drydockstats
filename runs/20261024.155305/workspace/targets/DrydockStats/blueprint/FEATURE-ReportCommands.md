# FEATURE: Commands Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the per-command CSV report over filtered records, ordered by calls descending then command ascending, with a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md, FEATURE-LogRequiredFields.md, FEATURE-LogDuplicates.md, FEATURE-LogNumeric.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-ReportFormat.md |
| Provides    | commands report |
| Consumes    | records, FormatMinutes, WriteCSV |

## Purpose

Write the `commands` report: one row per `job.command_name` over filtered records, followed by a `TOTAL` row.

## Report Contract

Header:

```
command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
```

- Rows are grouped by `job.command_name`.
- Rows are ordered by `calls` descending, then `command` ascending by byte-wise comparison (AC-011).
- `failed` counts filtered records whose `status` is not `succeeded`.
- `minutes` is written by `FormatMinutes` over summed `elapsed_ms`.
- The final row is `TOTAL` over all filtered records. With no matches it is `TOTAL,0,0,0,0,0,0.0`.

## Programmatic Acceptance

=== AC commands-ordering-and-total ===
Intent: Commands rows order by calls descending then byte-wise command ascending, and the TOTAL row sums all filtered records.
Suite: scoped

import csv
import io
import os
import sys
import tempfile

from drydock_go import cli_runner

def rec(eid, cmd, status="succeeded"):
    return ('{"schema_version":1,"execution_id":"%s","started_at":"2026-10-14T01:00:00.000Z",'
            '"status":"%s","job":{"command_name":"%s","target":"T","llm":"claude","model":"haiku"},'
            '"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":1,"cached_input_tokens":2,'
            '"output_tokens":3,"cost_usd":0.1,"elapsed_ms":1000}}}') % (eid, status, cmd)

lines = [
    rec("1", "plan"),
    rec("2", "build"),
    rec("3", "build"),
    rec("4", "analyze"),
    rec("5", "analyze"),
]

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "log.jsonl")
    with open(path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")
    code, out, err = cli_runner("commands", path)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens",
                       "output_tokens", "minutes"]
    commands = [r[0] for r in rows[1:-1]]
    assert commands == ["analyze", "build", "plan"]
    total = rows[-1]
    assert total[0] == "TOTAL"
    assert total[1] == "5"
    assert total[3] == "5"
=== END AC commands-ordering-and-total ===

## User Acceptance

- None.

## Guardrails

- The commands report must not read `sources/run.sh` or embed any expected answer.
- Ordering ties on call count are resolved by byte-wise command comparison, never locale collation.
