# FEATURE: Models Report

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The models report groups filtered records by provider and attributed model and ends with a TOTAL row. |
| Depends On  | FEATURE-SummaryReport.md, FEATURE-ModelAttribution.md, FEATURE-DateFilter.md |
| Provides    | ddkstat models |
| Consumes    | logread.Record.Model, report.FilterTarget, report.FilterDates |

## Purpose

`ddkstat models <log>` answers which providers and models ran the LLM calls, and what they used and cost.

## Report Contract

- Header: `provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd`.
- One row per distinct pair of `job.llm` (provider) and the record's attributed model (`result.stats.model` when non-empty, otherwise `job.model`).
- Rows are ordered by provider ascending, then model ascending, compared by byte value.
- `calls` counts filtered records in the group; `failed` counts those whose `status` is not `succeeded`.
- Token columns are sums of `input_tokens`, `cached_input_tokens`, and `output_tokens`; absent or null values count as zero.
- `cost_usd` is the summed cost written with `%.4f`; a null or absent cost counts as zero.
- The last row is `TOTAL` with an empty model cell (`TOTAL,,...`) summing every filtered record.
- With no matching record the output is the header and a `TOTAL` row of zeros, with cost `0.0000`.
- Output is CSV with `\n` line endings. `--target`, `--since`, and `--until` narrow the records before grouping.
- Aggregation and formatting are separate, unit-tested functions in `internal/report`.

## Programmatic Acceptance

=== AC models-grouping-order-total ===
Intent: The models report groups by provider and attributed model, orders by byte value, and totals every record.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

def rec(eid, started, status, llm, jm, stats):
    result = {"returncode": 0, "timed_out": False}
    if stats is not None:
        result["stats"] = stats
    return {"execution_id": eid, "started_at": started, "status": status,
            "job": {"command_name": "build", "target": "T", "llm": llm, "model": jm},
            "result": result}

recs = [
    rec("m1", "2026-10-14T01:00:00.000Z", "succeeded", "gemini", "Zed",
        {"input_tokens": 1, "cached_input_tokens": 2, "output_tokens": 3, "cost_usd": 0.25, "elapsed_ms": 10}),
    rec("m2", "2026-10-14T02:00:00.000Z", "succeeded", "claude", "sonnet",
        {"model": "claude-sonnet-5-5", "input_tokens": 10, "cached_input_tokens": 20, "output_tokens": 30, "cost_usd": 0.25}),
    rec("m3", "2026-10-14T03:00:00.000Z", "failed", "claude", "sonnet",
        {"model": "claude-sonnet-5-5", "input_tokens": 1, "cached_input_tokens": 2, "output_tokens": 3, "cost_usd": 0.5}),
    rec("m4", "2026-10-14T04:00:00.000Z", "succeeded", "claude", "sonnet", None),
    rec("m5", "2026-10-14T05:00:00.000Z", "succeeded", "claude", "haiku",
        {"model": "", "input_tokens": 4, "cached_input_tokens": 5, "output_tokens": 6, "cost_usd": 0.125}),
    rec("m6", "2026-10-14T06:00:00.000Z", "succeeded", "codex", "gpt-x",
        {"input_tokens": 7, "cached_input_tokens": 8, "output_tokens": 9, "cost_usd": None}),
    rec("m7", "2026-10-14T07:00:00.000Z", "error", "gemini", "alpha",
        {"input_tokens": 2, "cached_input_tokens": 2, "output_tokens": 2, "cost_usd": 0.0625}),
]
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write("\n".join(json.dumps(r) for r in recs) + "\n")

code, out, err = cli_runner("models", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
header = ["provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"]
assert rows[0] == header

groups = {}
for r in recs:
    st = r["result"].get("stats") or {}
    key = (r["job"]["llm"], st.get("model") or r["job"]["model"])
    g = groups.setdefault(key, [0, 0, 0, 0, 0, 0.0])
    g[0] += 1
    g[1] += 0 if r["status"] == "succeeded" else 1
    g[2] += st.get("input_tokens") or 0
    g[3] += st.get("cached_input_tokens") or 0
    g[4] += st.get("output_tokens") or 0
    g[5] += st.get("cost_usd") or 0.0
keys = sorted(groups)
body, total = rows[1:-1], rows[-1]
assert len(body) == len(keys)
for row, key in zip(body, keys):
    g = groups[key]
    assert row[0] == key[0] and row[1] == key[1]
    assert [int(x) for x in row[2:7]] == g[:5]
    assert len(row[7].split(".")[1]) == 4
    assert abs(float(row[7]) - g[5]) < 1e-9
assert total[0] == "TOTAL" and total[1] == ""
assert [int(x) for x in total[2:7]] == [sum(groups[k][i] for k in keys) for i in range(5)]
assert len(total[7].split(".")[1]) == 4
assert abs(float(total[7]) - sum(groups[k][5] for k in keys)) < 1e-9
=== END AC models-grouping-order-total ===

=== AC models-filters ===
Intent: Target and date filters narrow the models report to matching records only.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

def rec(eid, started, target, llm, jm):
    return {"execution_id": eid, "started_at": started, "status": "succeeded",
            "job": {"command_name": "build", "target": target, "llm": llm, "model": jm},
            "result": {"returncode": 0, "timed_out": False,
                       "stats": {"input_tokens": 3, "cached_input_tokens": 4, "output_tokens": 5, "cost_usd": 0.5}}}

recs = [
    rec("f1", "2026-10-14T01:00:00.000Z", "Alpha", "claude", "one"),
    rec("f2", "2026-10-15T01:00:00.000Z", "Alpha", "claude", "two"),
    rec("f3", "2026-10-15T02:00:00.000Z", "Beta", "codex", "three"),
    rec("f4", "2026-10-16T01:00:00.000Z", "Alpha", "gemini", "four"),
]
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write("\n".join(json.dumps(r) for r in recs) + "\n")

target, since, until = "Alpha", "2026-10-15", "2026-10-16"
code, out, err = cli_runner("models", "--target", target, "--since", since, "--until", until, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
kept = [r for r in recs if r["job"]["target"] == target and since <= r["started_at"][:10] <= until]
pairs = sorted((r["job"]["llm"], r["job"]["model"]) for r in kept)
assert [(row[0], row[1]) for row in rows[1:-1]] == pairs
assert rows[-1][0] == "TOTAL"
assert int(rows[-1][2]) == len(kept)
=== END AC models-filters ===

=== AC models-no-match ===
Intent: A filter matching nothing yields the header and a zero TOTAL row, with exit 0.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

rec = {"execution_id": "n1", "started_at": "2026-10-14T01:00:00.000Z", "status": "succeeded",
       "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "one"},
       "result": {"returncode": 0, "timed_out": False,
                  "stats": {"input_tokens": 3, "cached_input_tokens": 4, "output_tokens": 5, "cost_usd": 0.5}}}
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write(json.dumps(rec) + "\n")

code, out, err = cli_runner("models", "--target", "NoSuchTarget", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
total = rows[1]
assert total[0] == "TOTAL" and total[1] == ""
assert [int(x) for x in total[2:7]] == [0, 0, 0, 0, 0]
assert len(total[7].split(".")[1]) == 4
assert float(total[7]) == 0.0
=== END AC models-no-match ===

## User Acceptance

- None.

## Guardrails

- Every value is computed from the log being read; no row, count, or total is hardcoded.
- The log is never written to, and `sources/run.sh` is never read by application code.
- Only the Go standard library is used.
