# FEATURE: Models Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the per-provider and model CSV report over filtered records, ordered by provider then model ascending, with a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md, FEATURE-LogRequiredFields.md, FEATURE-LogDuplicates.md, FEATURE-LogNumeric.md, FEATURE-LogModel.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-ReportFormat.md |
| Provides    | models report |
| Consumes    | records, FormatCost, WriteCSV |

## Purpose

Write the `models` report: one row per provider and model pair over filtered records, followed by a `TOTAL` row.

## Report Contract

Header:

```
provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
```

- Rows are grouped by the provider from `job.llm` and the model from `LOG-004` attribution.
- Rows are ordered by provider then model ascending, byte-wise.
- `failed` counts filtered records whose `status` is not `succeeded`.
- `cost_usd` is written by `FormatCost`.
- The final row is `TOTAL` with an empty model column, over all filtered records. With no matches it is `TOTAL,,0,0,0,0,0,0.0000`.

## Programmatic Acceptance

=== AC models-ordering-and-total ===
Intent: Models rows group by provider and attributed model, ordered byte-wise, with the TOTAL row leaving the model column empty.
Suite: scoped

import csv
import io
import os
import sys
import tempfile

from drydock_go import cli_runner

def rec(eid, llm, job_model, stats_model=None):
    stats = '"input_tokens":1,"cached_input_tokens":2,"output_tokens":3,"cost_usd":0.1,"elapsed_ms":1000'
    if stats_model is not None:
        stats = '"model":"%s",' % stats_model + stats
    return ('{"schema_version":1,"execution_id":"%s","started_at":"2026-10-14T01:00:00.000Z",'
            '"status":"succeeded","job":{"command_name":"build","target":"T","llm":"%s","model":"%s"},'
            '"result":{"returncode":0,"timed_out":false,"stats":{%s}}}') % (eid, llm, job_model, stats)

lines = [
    rec("1", "codex", "gpt-x"),
    rec("2", "claude", "haiku", stats_model="claude-haiku-5-5"),
    rec("3", "claude", "haiku", stats_model="claude-haiku-5-5"),
]

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "log.jsonl")
    with open(path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")
    code, out, err = cli_runner("models", path)
    print(out)
    print(err, file=sys.stderr)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[0] == ["provider", "model", "calls", "failed", "input_tokens",
                       "cached_tokens", "output_tokens", "cost_usd"]
    pairs = [(r[0], r[1]) for r in rows[1:-1]]
    assert pairs == [("claude", "claude-haiku-5-5"), ("codex", "gpt-x")]
    total = rows[-1]
    assert total[0] == "TOTAL"
    assert total[1] == ""
    assert total[2] == "3"
=== END AC models-ordering-and-total ===

## User Acceptance

- None.

## Guardrails

- The models report must not read `sources/run.sh` or embed any expected answer.
- Model attribution follows FEATURE-LogModel exactly: `result.stats.model` when non-empty, otherwise `job.model`.
