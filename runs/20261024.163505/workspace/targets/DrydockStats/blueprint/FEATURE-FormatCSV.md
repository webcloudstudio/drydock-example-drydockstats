# Feature: CSV Report Writer

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Writes report rows as CSV with a header row, LF line endings, and optional TOTAL rows. |
| Depends On  | ARCHITECTURE.md |
| Provides    | CSV writer for report rows |
| Consumes    | — |

## Purpose

Every report shares one CSV writer. It writes a header row first, then data rows, comma-separated with `\n` line endings. Zero-match reports write their header and, where the report defines one, a zero TOTAL row.

## Behavior

- The first row written is the report header, exactly as the report defines it.
- Each subsequent row is written as comma-separated fields, terminated by `\n`.
- Line endings are `\n` only; no `\r\n`.
- A zero-match report writes its header and, where the report has a TOTAL row, a TOTAL row carrying zero values.
- The writer does not add quoting to plain values. Values in this product are alphanumeric, numeric, or dates, and contain no commas, quotes, or newlines.

## Programmatic Acceptance

=== AC csv-header-and-lf ===
Intent: The CSV writer emits the header first and uses LF-only line endings.
Requires: executable=python3; scope=test

import subprocess

from drydock_go import cli_runner

log = "/tmp/ddkstat-csv-fixture.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write('{"schema_version":1,"execution_id":"e1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"build","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":1,"elapsed_ms":1000,"input_tokens":2,"output_tokens":3,"cost_usd":0.5,"model":"haiku"}}}\n')

code, out, err = cli_runner("summary", log)
print(err)
assert code == 0, (code, out, err)
assert "\r" not in out, repr(out)
lines = out.split("\n")
assert lines[0] == "records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes", repr(lines[0])
=== END AC csv-header-and-lf ===

=== AC csv-zero-match-total ===
Intent: A zero-match commands report writes its header and a zero TOTAL row, and exits 0.
Requires: executable=python3; scope=test

import subprocess

from drydock_go import cli_runner

log = "/tmp/ddkstat-csv-empty.jsonl"
with open(log, "w", encoding="utf-8") as f:
    f.write("")

code, out, err = cli_runner("commands", "--target", "NoSuchTarget", log)
print(err)
assert code == 0, (code, out, err)
lines = [line for line in out.split("\n") if line != ""]
assert lines[0] == "command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes", repr(lines[0])
assert lines[-1].startswith("TOTAL,"), repr(lines[-1])
=== END AC csv-zero-match-total ===

## User Acceptance

- None.

## Guardrails

- Standard library CSV-free writing only; no quoting logic is required by the sources.
- Output goes to standard output only; diagnostics go to standard error.
