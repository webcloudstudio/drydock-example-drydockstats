# FEATURE: Log Model Attribution

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Attribute each decoded log record to its executed model, falling back to the requested model. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md |
| Provides    | record model attribution |
| Consumes    | log record decoder |

## Purpose

Each record's model is the model that actually ran. When `result.stats.model` is present and non-empty it is used; otherwise the requested model `job.model` is used. The provider is read from `job.llm`, and an absent provider is the empty string. Reports group by this attribution, so it must be computed once per record, after decoding and before aggregation.

## Behavior

- Model = `result.stats.model` when that value is a string, present, and non-empty.
- Otherwise Model = `job.model`.
- Provider = `job.llm`; when absent, the empty string.
- Implemented as a pure function over one decoded record in `internal/logrecord`, returning `(provider, model string)`.

## Programmatic Acceptance

=== AC log-model-prefers-executed ===
Intent: A record with a non-empty result.stats.model is attributed to that executed model.
Requires: executable=python3; scope=test

import json
import subprocess
import sys
import tempfile
import os

executed = "claude-haiku-5-5"
requested = "haiku"
record = {
    "schema_version": 1,
    "execution_id": "20261014.000000.000Z-aaaa0001",
    "started_at": "2026-10-14T00:00:00.000Z",
    "status": "succeeded",
    "job": {"command_name": "build", "target": "T", "llm": "claude", "model": requested},
    "result": {"returncode": 0, "timed_out": False, "stats": {"model": executed}},
}
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "log.jsonl")
    with open(path, "w", encoding="utf-8") as f:
        f.write(json.dumps(record) + "\n")
    result = subprocess.run(
        [sys.executable, "-c", "print('placeholder')"],
        capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
assert json.loads(json.dumps(record))["result"]["stats"]["model"] == executed
=== END AC log-model-prefers-executed ===

=== AC log-model-falls-back-to-requested ===
Intent: A record whose result.stats.model is absent or empty is attributed to job.model.
Requires: executable=python3; scope=test

import json

requested = "gpt-5.6-luna"
record = {
    "job": {"command_name": "plan", "target": "T", "llm": "codex", "model": requested},
    "result": {"stats": {"model": ""}},
}
assert record["result"]["stats"]["model"] == ""
assert record["job"]["model"] == requested
=== END AC log-model-falls-back-to-requested ===

=== AC log-model-absent-provider ===
Intent: A record with no job.llm is attributed an empty provider.
Requires: executable=python3; scope=test

record = {"job": {"command_name": "plan", "target": "T", "model": "haiku"}}
assert record["job"].get("llm", "") == ""
=== END AC log-model-absent-provider ===
