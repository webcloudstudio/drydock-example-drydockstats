# FEATURE: Daily Report

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | The daily report summarizes filtered records per start date and ends with a TOTAL row. |
| Depends On  | FEATURE-SummaryReport.md, FEATURE-DateFilter.md |
| Provides    | ddkstat daily |
| Consumes    | report.FilterTarget, report.FilterDates |

## Purpose

`ddkstat daily <log>` shows how many calls, failures, tokens, and minutes each day used.

## Report Contract

- Header: `date,calls,failed,tokens,minutes`.
- One row per start date (the first 10 characters of `started_at`) that has at least one matching record, ordered ascending. Dates with no matching record do not appear.
- `tokens` is input plus cached plus output tokens, with absent or null counts as zero.
- `minutes` is the group's summed `elapsed_ms` written in tenths of a minute: `t = (ms + 3000) / 6000` in integer division, written as `t / 10`, a `.`, and `t % 10`.
- The last row is `TOTAL` over all filtered records, with minutes computed from the summed milliseconds.
- With no matching record the output is the header and `TOTAL,0,0,0,0.0`.
- A record without `result.stats` still counts as a call with zero tokens and minutes.
- Output is CSV with `\n` line endings; `--target`, `--since`, and `--until` narrow the records first.

## Programmatic Acceptance

=== AC daily-rows-and-total ===
Intent: Daily rows are ascending by date, with tokens and tenth-minute values derived from the fixture records.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

def rec(eid, started, status, stats):
    result = {"returncode": 0, "timed_out": False}
    if stats is not None:
        result["stats"] = stats
    return {"execution_id": eid, "started_at": started, "status": status,
            "job": {"command_name": "build", "target": "T", "llm": "claude", "model": "m"},
            "result": result}

def st(i, c, o, ms):
    return {"input_tokens": i, "cached_input_tokens": c, "output_tokens": o, "elapsed_ms": ms}

recs = [
    rec("d1", "2026-10-16T09:00:00.000Z", "succeeded", st(1, 2, 3, 95000)),
    rec("d2", "2026-10-14T01:00:00.000Z", "succeeded", st(10, 20, 30, 2999)),
    rec("d3", "2026-10-18T01:00:00.000Z", "failed", None),
    rec("d4", "2026-10-14T05:00:00.000Z", "failed", st(4, 5, 6, 3000)),
    rec("d5", "2026-10-16T10:00:00.000Z", "succeeded", st(7, 8, 9, 6000)),
]
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write("\n".join(json.dumps(r) for r in recs) + "\n")

def tenths(cell):
    whole, frac = cell.split(".")
    assert len(frac) == 1
    return int(whole) * 10 + int(frac)

def tt(ms):
    return (ms + 3000) // 6000

code, out, err = cli_runner("daily", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["date", "calls", "failed", "tokens", "minutes"]
days = {}
for r in recs:
    s = r["result"].get("stats") or {}
    g = days.setdefault(r["started_at"][:10], [0, 0, 0, 0])
    g[0] += 1
    g[1] += 0 if r["status"] == "succeeded" else 1
    g[2] += (s.get("input_tokens") or 0) + (s.get("cached_input_tokens") or 0) + (s.get("output_tokens") or 0)
    g[3] += s.get("elapsed_ms") or 0
dates = sorted(days)
assert [row[0] for row in rows[1:-1]] == dates
for row in rows[1:-1]:
    g = days[row[0]]
    assert [int(row[1]), int(row[2]), int(row[3])] == g[:3]
    assert tenths(row[4]) == tt(g[3])
total = rows[-1]
assert total[0] == "TOTAL"
assert [int(total[1]), int(total[2]), int(total[3])] == [sum(days[k][i] for k in dates) for i in range(3)]
assert tenths(total[4]) == tt(sum(days[k][3] for k in dates))
=== END AC daily-rows-and-total ===

=== AC daily-filters ===
Intent: Target and date filters narrow the daily report and a date with no matching record is not listed.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

def rec(eid, started, target):
    return {"execution_id": eid, "started_at": started, "status": "succeeded",
            "job": {"command_name": "build", "target": target, "llm": "claude", "model": "m"},
            "result": {"returncode": 0, "timed_out": False,
                       "stats": {"input_tokens": 1, "cached_input_tokens": 1, "output_tokens": 1, "elapsed_ms": 6000}}}

recs = [
    rec("g1", "2026-10-14T01:00:00.000Z", "Alpha"),
    rec("g2", "2026-10-15T01:00:00.000Z", "Beta"),
    rec("g3", "2026-10-16T01:00:00.000Z", "Alpha"),
    rec("g4", "2026-10-17T01:00:00.000Z", "Alpha"),
]
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write("\n".join(json.dumps(r) for r in recs) + "\n")

target, since, until = "Alpha", "2026-10-14", "2026-10-16"
code, out, err = cli_runner("daily", "--target", target, "--since", since, "--until", until, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
kept = [r for r in recs if r["job"]["target"] == target and since <= r["started_at"][:10] <= until]
assert [row[0] for row in rows[1:-1]] == sorted({r["started_at"][:10] for r in kept})
assert "2026-10-15" not in [row[0] for row in rows]
assert rows[-1][0] == "TOTAL" and int(rows[-1][1]) == len(kept)
=== END AC daily-filters ===

=== AC daily-no-match ===
Intent: A filter matching nothing yields the header and a zero TOTAL row, with exit 0.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from drydock_go import cli_runner

rec = {"execution_id": "z1", "started_at": "2026-10-14T01:00:00.000Z", "status": "succeeded",
       "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "m"},
       "result": {"returncode": 0, "timed_out": False}}
d = tempfile.mkdtemp()
path = os.path.join(d, "log.jsonl")
with open(path, "w") as fh:
    fh.write(json.dumps(rec) + "\n")

code, out, err = cli_runner("daily", "--since", "2027-01-01", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert rows[1][0] == "TOTAL"
assert [int(x) for x in rows[1][1:4]] == [0, 0, 0]
whole, frac = rows[1][4].split(".")
assert int(whole) == 0 and len(frac) == 1 and int(frac) == 0
=== END AC daily-no-match ===

## User Acceptance

- None.

## Guardrails

- Every value is computed from the log being read; no row, count, or total is hardcoded.
- The log is never written to, and `sources/run.sh` is never read by application code.
- Only the Go standard library is used.
