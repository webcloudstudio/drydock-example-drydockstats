# FEATURE: Summary Report

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Write the one-row summary CSV and provide the shared CSV, minutes, and cost formatters. |
| Depends On  | FEATURE-DateFilter.md |
| Provides    | ddkstat summary |
| Consumes    | report.FilterTarget, report.FilterDates, logread.Skipped |

## Purpose

`ddkstat summary <log>` answers how many records, skips, successes, failures, tokens, dollars and
minutes the log holds, after the Target and date filters. This story also defines the formatting
helpers every other report reuses.

## Output

Header, exactly:

```text
records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
```

One data row:

- `records` is the count of kept, filtered records. `succeeded` has status `succeeded`; `failed` is every other status.
- `skipped` is the count of skipped lines in the whole file, regardless of filters.
- `first` and `last` are the earliest and latest start dates (`YYYY-MM-DD`) among matching records, empty when none match.
- `input_tokens`, `cached_tokens`, `output_tokens` are sums of `input_tokens`, `cached_input_tokens`, `output_tokens`; absent values count as zero.
- `cost_usd` is the summed `cost_usd` written with `%.4f`.
- `minutes` is the summed `elapsed_ms` as tenths of a minute: `t = (ms + 3000) / 6000` in integer division, written `t/10`, `.`, `t%10`.

## Shared formatting

- CSV has one header row and `\n` line endings, with no `\r`.
- Aggregation and formatting are separate functions: aggregation returns numbers, formatting turns them into text.
- Minutes rounding applies to the group's summed milliseconds, never to per-record values.
- With no matching record: `records` 0, `first` and `last` empty, sums 0, cost `0.0000`, minutes `0.0`.

## Programmatic Acceptance

=== AC summary-columns-and-sums ===
Intent: The summary has the declared header and one row whose counts, token sums, cost and minutes derive from the fixture records.
Suite: scoped
Requires: executable=go; scope=test

import csv, io, json, os, tempfile
from drydock_go import cli_runner

def rec(i, target="T1", day="2026-10-14", status="succeeded", cmd="build", llm="claude", model="m", inp=10, cached=20, out=30, ms=60000, cost=0.25, ts=None):
    return {"execution_id": f"id-{i}", "started_at": ts or f"{day}T01:02:03.004Z", "status": status,
            "job": {"command_name": cmd, "target": target, "llm": llm, "model": model},
            "result": {"returncode": 0 if status == "succeeded" else 1, "timed_out": False,
                       "stats": {"input_tokens": inp, "cached_input_tokens": cached, "output_tokens": out, "elapsed_ms": ms, "cost_usd": cost}}}

def run(lines, *args):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "llm.jsonl")
        with open(p, "w", encoding="utf-8") as f:
            for l in lines:
                f.write((l if isinstance(l, str) else json.dumps(l)) + "\n")
        code, out, err = cli_runner(*args, p)
        assert code == 0
        print(out)
        print(err)
        return code, out

recs = [
    rec(1, day="2026-10-15", inp=1, cached=2, out=3, ms=100000, cost=0.25),
    rec(2, day="2026-10-14", inp=10, cached=20, out=30, ms=200000, cost=0.5, status="failed"),
    rec(3, day="2026-10-16", inp=100, cached=200, out=300, ms=50000, cost=1.0),
]
code, out = run(recs, "summary")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"]
assert len(rows) == 2
d = dict(zip(rows[0], rows[1]))
st = lambda r: r["result"]["stats"]
assert int(d["records"]) == len(recs)
assert int(d["succeeded"]) == len([r for r in recs if r["status"] == "succeeded"])
assert int(d["failed"]) == len([r for r in recs if r["status"] != "succeeded"])
assert d["first"] == min(r["started_at"][:10] for r in recs)
assert d["last"] == max(r["started_at"][:10] for r in recs)
assert int(d["input_tokens"]) == sum(st(r)["input_tokens"] for r in recs)
assert int(d["cached_tokens"]) == sum(st(r)["cached_input_tokens"] for r in recs)
assert int(d["output_tokens"]) == sum(st(r)["output_tokens"] for r in recs)
assert abs(float(d["cost_usd"]) - sum(st(r)["cost_usd"] for r in recs)) < 1e-9
assert len(d["cost_usd"].split(".")[1]) == 4
t = (sum(st(r)["elapsed_ms"] for r in recs) + 3000) // 6000
assert float(d["minutes"]) == t / 10
assert len(d["minutes"].split(".")[1]) == 1
=== END AC summary-columns-and-sums ===

=== AC summary-minutes-rounding ===
Intent: Minutes round the summed milliseconds to tenths with t = (ms + 3000) / 6000 in integer division.
Suite: scoped
Requires: executable=go; scope=test

import csv, io, json, os, tempfile
from drydock_go import cli_runner

def rec(i, target="T1", day="2026-10-14", status="succeeded", cmd="build", llm="claude", model="m", inp=10, cached=20, out=30, ms=60000, cost=0.25, ts=None):
    return {"execution_id": f"id-{i}", "started_at": ts or f"{day}T01:02:03.004Z", "status": status,
            "job": {"command_name": cmd, "target": target, "llm": llm, "model": model},
            "result": {"returncode": 0 if status == "succeeded" else 1, "timed_out": False,
                       "stats": {"input_tokens": inp, "cached_input_tokens": cached, "output_tokens": out, "elapsed_ms": ms, "cost_usd": cost}}}

def run(lines, *args):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "llm.jsonl")
        with open(p, "w", encoding="utf-8") as f:
            for l in lines:
                f.write((l if isinstance(l, str) else json.dumps(l)) + "\n")
        code, out, err = cli_runner(*args, p)
        assert code == 0
        print(out)
        print(err)
        return code, out

for ms in [0, 2999, 3000, 8999, 9000, 95000, 6000000]:
    code, out = run([rec(1, ms=ms)], "summary")
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    cell = dict(zip(rows[0], rows[1]))["minutes"]
    t = (ms + 3000) // 6000
    whole, frac = cell.split(".")
    assert int(whole) == t // 10, ms
    assert int(frac) == t % 10, ms
    assert len(frac) == 1, ms

# Rounding applies to the sum, not to each record.
parts = [2000, 2000, 2000]
code, out = run([rec(i, ms=m) for i, m in enumerate(parts)], "summary")
rows = list(csv.reader(io.StringIO(out)))
cell = dict(zip(rows[0], rows[1]))["minutes"]
t = (sum(parts) + 3000) // 6000
assert float(cell) == t / 10
=== END AC summary-minutes-rounding ===

=== AC summary-absent-values-count-zero ===
Intent: Records with null cost, no stats object, or null numbers still count as records and contribute zero.
Suite: scoped
Requires: executable=go; scope=test

import csv, io, json, os, tempfile
from drydock_go import cli_runner

def rec(i, target="T1", day="2026-10-14", status="succeeded", cmd="build", llm="claude", model="m", inp=10, cached=20, out=30, ms=60000, cost=0.25, ts=None):
    return {"execution_id": f"id-{i}", "started_at": ts or f"{day}T01:02:03.004Z", "status": status,
            "job": {"command_name": cmd, "target": target, "llm": llm, "model": model},
            "result": {"returncode": 0 if status == "succeeded" else 1, "timed_out": False,
                       "stats": {"input_tokens": inp, "cached_input_tokens": cached, "output_tokens": out, "elapsed_ms": ms, "cost_usd": cost}}}

def run(lines, *args):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "llm.jsonl")
        with open(p, "w", encoding="utf-8") as f:
            for l in lines:
                f.write((l if isinstance(l, str) else json.dumps(l)) + "\n")
        code, out, err = cli_runner(*args, p)
        assert code == 0
        print(out)
        print(err)
        return code, out

full = rec(1)
null_cost = rec(2, cost=None)
no_stats = rec(3)
del no_stats["result"]["stats"]
recs = [full, null_cost, no_stats]
code, out = run(recs, "summary")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == len(recs)
assert int(d["input_tokens"]) == full["result"]["stats"]["input_tokens"] + null_cost["result"]["stats"]["input_tokens"]
assert abs(float(d["cost_usd"]) - full["result"]["stats"]["cost_usd"]) < 1e-9
=== END AC summary-absent-values-count-zero ===

=== AC summary-skipped-is-whole-file ===
Intent: The skipped count covers the whole file even when a filter narrows the records, and blank lines are not counted.
Suite: scoped
Requires: executable=go; scope=test

import csv, io, json, os, tempfile
from drydock_go import cli_runner

def rec(i, target="T1", day="2026-10-14", status="succeeded", cmd="build", llm="claude", model="m", inp=10, cached=20, out=30, ms=60000, cost=0.25, ts=None):
    return {"execution_id": f"id-{i}", "started_at": ts or f"{day}T01:02:03.004Z", "status": status,
            "job": {"command_name": cmd, "target": target, "llm": llm, "model": model},
            "result": {"returncode": 0 if status == "succeeded" else 1, "timed_out": False,
                       "stats": {"input_tokens": inp, "cached_input_tokens": cached, "output_tokens": out, "elapsed_ms": ms, "cost_usd": cost}}}

def run(lines, *args):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "llm.jsonl")
        with open(p, "w", encoding="utf-8") as f:
            for l in lines:
                f.write((l if isinstance(l, str) else json.dumps(l)) + "\n")
        code, out, err = cli_runner(*args, p)
        assert code == 0
        print(out)
        print(err)
        return code, out

good = [rec(1, "Alpha"), rec(2, "Beta")]
dup = rec(1, "Alpha")
broken = ["{truncated", "42"]
lines = [good[0], "", broken[0], good[1], dup, broken[1]]
skipped_expected = len(broken) + 1
code, out = run(lines, "summary", "--target", "Alpha")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == 1
assert int(d["skipped"]) == skipped_expected
=== END AC summary-skipped-is-whole-file ===

=== AC summary-empty-and-format ===
Intent: With no matching record the summary still writes a zero row, with four-decimal cost and one-decimal minutes, using LF line endings.
Suite: scoped
Requires: executable=go; scope=test

import csv, io, json, os, tempfile
from drydock_go import cli_runner

def rec(i, target="T1", day="2026-10-14", status="succeeded", cmd="build", llm="claude", model="m", inp=10, cached=20, out=30, ms=60000, cost=0.25, ts=None):
    return {"execution_id": f"id-{i}", "started_at": ts or f"{day}T01:02:03.004Z", "status": status,
            "job": {"command_name": cmd, "target": target, "llm": llm, "model": model},
            "result": {"returncode": 0 if status == "succeeded" else 1, "timed_out": False,
                       "stats": {"input_tokens": inp, "cached_input_tokens": cached, "output_tokens": out, "elapsed_ms": ms, "cost_usd": cost}}}

def run(lines, *args):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "llm.jsonl")
        with open(p, "w", encoding="utf-8") as f:
            for l in lines:
                f.write((l if isinstance(l, str) else json.dumps(l)) + "\n")
        code, out, err = cli_runner(*args, p)
        assert code == 0
        print(out)
        print(err)
        return code, out

code, out = run([rec(1), "garbage"], "summary", "--target", "NoSuchTarget")
assert code == 0
assert out.endswith("\n")
assert "\r" not in out
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == 0
assert int(d["skipped"]) == 1
assert d["first"] == "" and d["last"] == ""
assert [int(d[k]) for k in ["succeeded", "failed", "input_tokens", "cached_tokens", "output_tokens"]] == [0, 0, 0, 0, 0]
assert float(d["cost_usd"]) == 0.0 and len(d["cost_usd"].split(".")[1]) == 4
assert float(d["minutes"]) == 0.0 and len(d["minutes"].split(".")[1]) == 1
=== END AC summary-empty-and-format ===

## User Acceptance

- None.

## Guardrails

- Every value is computed from the log being read; no row, count, or total is hardcoded.
- Application code never reads `sources/run.sh`.
- Cost is formatted with Go's `%.4f` on a `float64`.
- The log is never written to.
