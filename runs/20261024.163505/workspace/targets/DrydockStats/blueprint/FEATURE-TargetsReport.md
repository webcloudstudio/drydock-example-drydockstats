# Feature: Targets Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces one row per Target ordered ascending, with cost and minutes columns, plus a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-LogNumbers.md, FEATURE-LogFailed.md, FEATURE-FormatCSV.md, FEATURE-FormatMinutes.md, FEATURE-FormatCost.md |
| Provides    | targets report |
| Consumes    | record aggregates, CSV and format helpers |

## Purpose

The targets report shows, per Target, how many calls ran, how many failed, and what tokens, cost, and minutes they used.

## Behavior

- Header: `target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`.
- One row per distinct `job.target` among filtered records, ordered ascending by byte value.
- `cost_usd` is written with `%.4f`; `minutes` is written in tenths of a minute.
- A final `TOTAL` row aggregates all filtered records.
- A zero-match run writes the header and one row `TOTAL,0,0,0,0,0,0.0000,0.0`.

## Programmatic Acceptance

=== AC targets-ordering-and-total ===
Intent: Targets are ordered ascending by byte value and TOTAL sums the rows.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-targets-order.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Zeta","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":1,"elapsed_ms":60000,"input_tokens":1,"output_tokens":1,"cost_usd":0.1,"model":"haiku"}}}\n')
    f.write('{"schema_version":1,"execution_id":"e2","started_at":"2026-10-14T06:01:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Alpha","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":1,"elapsed_ms":60000,"input_tokens":1,"output_tokens":1,"cost_usd":0.1,"model":"haiku"}}}\n')

code, out, err = cli_runner("targets", log)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[1][0] == "Alpha", rows[1]
assert rows[2][0] == "Zeta", rows[2]
assert rows[3][0] == "TOTAL", rows[3]
assert rows[3][1] == "2", rows[3]
assert rows[3][6] == "0.2000", rows[3]
=== END AC targets-ordering-and-total ===

=== AC targets-zero-match ===
Intent: A filter matching nothing yields the header and a zero TOTAL row with cost and minutes, exit 0.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-targets-zero.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')

code, out, err = cli_runner("targets", "--target", "NoSuchTarget", log)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"], rows[0]
assert rows[1][0] == "TOTAL", rows[1]
assert rows[1][6] == "0.0000", rows[1]
assert rows[1][7] == "0.0", rows[1]
=== END AC targets-zero-match ===

## User Acceptance

- None.

## Guardrails

- Targets are grouped by the exact `job.target` value; no normalization or case folding.
- Filters apply to every row and to TOTAL.
