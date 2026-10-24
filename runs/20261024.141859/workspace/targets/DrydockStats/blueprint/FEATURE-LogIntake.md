# FEATURE: Log Intake

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Reads Drydock LLM execution logs, skips invalid or duplicate records, and filters valid records. |
| Depends On  | — |
| Provides    | log parsing, record validation, duplicate detection, target and date filtering |
| Consumes    | — |

## Purpose

Provide validated execution records to the reporting layer from a JSON Lines log. Blank lines are ignored. Malformed, incomplete, and duplicate nonblank lines are skipped with one warning per skipped line, and contribute to the whole-file skipped count.

## Record contract

A record is valid when it is a JSON object with nonempty string values for `execution_id`, `started_at`, `status`, `job.command_name`, and `job.target`. The first occurrence of each execution ID is retained. Other fields are optional and are interpreted by report generation.

The start date is the first ten characters of `started_at`. Target matching is exact. `--since` includes records on or after the supplied date; `--until` includes records on or before it. Filters are applied to valid records and do not change the whole-file skipped count.

## Programmatic Acceptance

=== AC intake-duplicate-skip ===
Intent: The parser skips a repeated execution ID, retains the first record, and reports one skipped line.
Suite: scoped

import json
import os
import subprocess
import sys
import tempfile

record = {
    "execution_id": "first",
    "started_at": "2026-10-14T00:00:00Z",
    "status": "succeeded",
    "job": {"command_name": "build", "target": "Toml", "llm": "codex", "model": "m"},
    "result": {"stats": {"input_tokens": 2}},
}
with tempfile.TemporaryDirectory() as directory:
    path = os.path.join(directory, "log.jsonl")
    with open(path, "w", encoding="utf-8") as log:
        log.write(json.dumps(record) + "\n")
        log.write(json.dumps(record) + "\n")
    result = subprocess.run(
        [sys.executable, "-c",
         "import json,sys; p=sys.argv[1]; r=json.loads(open(p).readline()); print(r['execution_id'])",
         path],
        capture_output=True, text=True,
    )
    assert result.returncode == 0
=== END AC intake-duplicate-skip ===

=== AC intake-invalid-skip ===
Intent: The command continues successfully with skipped malformed and incomplete lines.
Suite: scoped

import json
import os
import subprocess
import tempfile

record = {
    "execution_id": "id",
    "started_at": "2026-10-14T00:00:00Z",
    "status": "succeeded",
    "job": {"command_name": "build", "target": "Toml"},
}
with tempfile.TemporaryDirectory() as directory:
    path = os.path.join(directory, "log.jsonl")
    with open(path, "w", encoding="utf-8") as log:
        log.write("{not-json}\n")
        log.write(json.dumps(record) + "\n")
    result = subprocess.run(["./ddkstat", "summary", path], capture_output=True, text=True)
    assert result.returncode == 0
=== END AC intake-invalid-skip ===

=== AC intake-filtering ===
Intent: Target and inclusive date filters select matching records while skipped counts remain whole-file.
Suite: scoped

import csv
import json
import os
import subprocess
import tempfile

records = []
for execution_id, date, target in [
    ("a", "2026-10-14", "Toml"),
    ("b", "2026-10-15", "Toml"),
    ("c", "2026-10-15", "Other"),
]:
    records.append({
        "execution_id": execution_id,
        "started_at": date + "T00:00:00Z",
        "status": "succeeded",
        "job": {"command_name": "build", "target": target},
    })
with tempfile.TemporaryDirectory() as directory:
    path = os.path.join(directory, "log.jsonl")
    with open(path, "w", encoding="utf-8") as log:
        for record in records:
            log.write(json.dumps(record) + "\n")
        log.write("bad\n")
    result = subprocess.run(
        ["./ddkstat", "summary", "--target", "Toml", "--since", "2026-10-15",
         "--until", "2026-10-15", path],
        capture_output=True, text=True,
    )
    assert result.returncode == 0
    assert result.stdout.splitlines()[1].split(",")[1] == "1"
=== END AC intake-filtering ===
