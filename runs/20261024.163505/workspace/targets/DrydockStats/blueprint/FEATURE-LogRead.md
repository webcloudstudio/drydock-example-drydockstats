# Feature: Log Read

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Reads a JSON Lines execution log from a file path as UTF-8, one line at a time, ignoring blank lines. |
| Depends On  | ARCHITECTURE.md |
| Provides    | logfile line reader for a file path |
| Consumes    | - |

## Purpose

Provide the single entry point through which every report reads Drydock's LLM execution log. The reader yields each non-blank line as raw bytes in file order. Parsing, skipping, and aggregation happen downstream.

## Behavior

- The file path is opened and read as UTF-8 bytes, one line at a time.
- Lines that are empty or contain only whitespace are ignored and are not counted as skipped.
- Each non-blank line is passed on as raw bytes, with its one-based line number preserved for diagnostics.
- A file that cannot be opened or read returns an operational error to the caller, which maps it to exit 1.
- The reader never writes to the log file.

## Programmatic Acceptance

=== AC log-read-blank-lines ===
Intent: Blank lines are ignored by the line reader and do not appear in the yielded lines.
Suite: scoped
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

from drydock_go import cli_runner

content = '{"execution_id":"a"}\n\n   \n{"execution_id":"b"}\n'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(content)
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = out.splitlines()
    assert rows[0].startswith("records,skipped,"), out
    fields = rows[1].split(",")
    assert fields[0] == "2", out
    assert fields[1] == "0", out
=== END AC log-read-blank-lines ===

=== AC log-read-utf8 ===
Intent: The line reader reads the file as UTF-8 and a multibyte character inside a valid record is kept intact.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

record = '{"schema_version":1,"execution_id":"u-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan \u00e9","target":"Caf\u00e9","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cost_usd":0.5,"elapsed_ms":1000,"input_tokens":1,"cached_input_tokens":0,"output_tokens":1}}}\n'
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "log.jsonl")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write(record)
    code, out, err = cli_runner("commands", path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    rows = out.splitlines()
    assert rows[0] == "command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes", out
    fields = rows[1].split(",")
    assert fields[0] == "plan \u00e9", out
=== END AC log-read-utf8 ===

=== AC log-read-missing-file ===
Intent: A log path that cannot be opened exits 1 with nothing on stdout.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "no-such-log.jsonl")
    code, out, err = cli_runner("summary", path)
    print(out)
    print(err)
    assert code == 1, (code, out, err)
    assert out == "", out
=== END AC log-read-missing-file ===

## User Acceptance

- None.

## Guardrails

- Blank lines are not counted as skipped (see LOG-002 and Rule 1).
- The reader never writes to the log being read.
- Standard library only.
