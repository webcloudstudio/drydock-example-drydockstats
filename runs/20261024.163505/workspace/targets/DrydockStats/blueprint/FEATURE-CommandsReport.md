# Feature: Commands Report

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produces one row per command, ordered by calls then command name, plus a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-FilterTarget.md, FEATURE-FilterDate.md, FEATURE-LogNumbers.md, FEATURE-LogFailed.md, FEATURE-FormatCSV.md, FEATURE-FormatMinutes.md |
| Provides    | commands report |
| Consumes    | record aggregates, CSV and format helpers |

## Purpose

The commands report shows which Drydock commands ran, how often, how many failed, and what tokens and minutes they used.

## Behavior

- Header: `command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes`.
- One row per distinct `job.command_name` among filtered records.
- Rows are ordered by `calls` descending, then `command` ascending by byte value (AC-006).
- Each row's `minutes` is written in tenths of a minute (FEATURE-FormatMinutes.md).
- A final `TOTAL` row aggregates all filtered records.
- A zero-match run writes the header and one row `TOTAL,0,0,0,0,0,0.0`.

## Programmatic Acceptance

=== AC commands-ordering-and-total ===
Intent: Commands are ordered by calls descending then name ascending, and TOTAL sums the rows.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-commands-order.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:00:00.000Z","status":"succeeded","job":{"command_name":"plan","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":1,"elapsed_ms":60000,"input_tokens":1,"output_tokens":1,"cost_usd":0.1,"model":"haiku"}}}\n')
    f.write('{"schema_version":1,"execution_id":"e2","started_at":"2026-10-14T06:01:00.000Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":1,"elapsed_ms":60000,"input_tokens":1,"output_tokens":1,"cost_usd":0.1,"model":"haiku"}}}\n')
    f.write('{"schema_version":1,"execution_id":"e3","started_at":"2026-10-14T06:02:00.000Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":1,"elapsed_ms":60000,"input_tokens":1,"output_tokens":1,"cost_usd":0.1,"model":"haiku"}}}\n')

code, out, err = cli_runner("commands", log)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[1][0] == "build", rows[1]
assert rows[2][0] == "plan", rows[2]
assert rows[3][0] == "TOTAL", rows[3]
assert rows[3][1] == "3", rows[3]
assert rows[3][3] == "3", rows[3]
=== END AC commands-ordering-and-total ===

=== AC commands-tie-byte-order ===
Intent: Commands with equal call counts are ordered by command name in byte order.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-commands-tie.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:00:00.000Z","status":"succeeded","job":{"command_name":"score release","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')
    f.write('{"schema_version":1,"execution_id":"e2","started_at":"2026-10-14T06:01:00.000Z","status":"succeeded","job":{"command_name":"analyze","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')

code, out, err = cli_runner("commands", log)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[1][0] == "analyze", rows[1]
assert rows[2][0] == "score release", rows[2]
=== END AC commands-tie-byte-order ===

=== AC commands-zero-match ===
Intent: A filter matching nothing yields the header and a zero TOTAL row, exit 0.
Requires: executable=python3; scope=test

import csv
import io

from drydock_go import cli_runner

log = "/tmp/ddkstat-commands-zero.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"A","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false}}\n')

code, out, err = cli_runner("commands", "--target", "NoSuchTarget", log)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"], rows[0]
assert rows[1][0] == "TOTAL", rows[1]
assert rows[1][1] == "0", rows[1]
assert rows[1][6] == "0.0", rows[1]
=== END AC commands-zero-match ===

## User Acceptance

- None.

## Guardrails

- Commands are grouped by the exact `job.command_name` value; no normalization or case folding.
- Filters apply to every row and to TOTAL.
