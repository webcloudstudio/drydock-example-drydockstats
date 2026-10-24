# FEATURE: Targets Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the per-Target CSV report over filtered records, ordered by Target ascending, with a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md, FEATURE-LogRequiredFields.md, FEATURE-LogDuplicates.md, FEATURE-LogNumeric.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-ReportFormat.md |
| Provides    | targets report |
| Consumes    | records, FormatCost, FormatMinutes, WriteCSV |

## Purpose

Write the `targets` report: one row per `job.target` over filtered records, followed by a `TOTAL` row.

## Report Contract

Header:

```
target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
```

- Rows are grouped by `job.target`.
- Rows are ordered by `target` ascending, byte-wise.
- `failed` counts filtered records whose `status` is not `succeeded`.
- `cost_usd` is written by `FormatCost`; `minutes` by `FormatMinutes`.
- The final row is `TOTAL` over all filtered records. With no matches it is `TOTAL,0,0,0,0,0,0.0000,0.0`.

## Programmatic Acceptance

=== AC targets-ordering-and-total ===
Intent: Targets rows order by target ascending byte-wise and the TOTAL row sums all filtered records.
Suite: scoped

import csv
import io
import os
import sys
import tempfile

from drydock_go import cli_runner

def rec(eid, target, status="succeeded"):
    return ('{"schema_version":1,"execution_id":"%s","started_at":"2026-10-14T01:00:00.000Z",'
            '"status":"%s","job":{"command_name":"build","target":"%s","llm":"claude","model":"haiku"},'
            '"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":1,"cached_input_tokens":2,'
            '"output_tokens":3,"cost_usd":0.25,"elapsed_ms":1000}}}') % (eid, status, target)

lines = [
    rec("1", "Zeta"),
    rec("2", "Alpha"),
    rec("3", "Alpha", status="failed"),
]

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "log.jsonl")
    with open(path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")
    code, out, err = cli_runner("targets", path)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[0] == ["target", "calls", "failed", "input_tokens", "cached_tokens",
                       "output_tokens", "cost_usd", "minutes"]
    targets = [r[0] for r in rows[1:-1]]
    assert targets == ["Alpha", "Zeta"]
    total = rows[-1]
    assert total[0] == "TOTAL"
    assert total[1] == "3"
    assert total[2] == "1"
=== END AC targets-ordering-and-total ===

## User Acceptance

- None.

## Guardrails

- The targets report must not read `sources/run.sh` or embed any expected answer.
- Target matching for grouping is exact and case-sensitive.
