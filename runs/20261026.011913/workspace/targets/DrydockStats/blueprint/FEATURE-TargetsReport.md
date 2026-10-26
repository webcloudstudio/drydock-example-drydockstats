# FEATURE: Targets Report

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Write one CSV row per Target, ordered by Target name, followed by a TOTAL row with cost and minutes. |
| Depends On  | FEATURE-SummaryReport.md |
| Provides    | ddkstat targets |
| Consumes    | report.FilterTarget, report.FilterDates, shared CSV, minutes, and cost formatters |

## Purpose

`ddkstat targets <log>` shows what each Drydock Target consumed and cost.

## Output

Header, exactly:

```text
target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
```

- One row per distinct `job.target` among filtered records, ordered ascending by byte value.
- `failed` counts records whose status is not `succeeded`.
- Token columns are sums; `cost_usd` is the summed cost written with `%.4f`; `minutes` follows the shared minutes rule over the group's summed `elapsed_ms`.
- The final row is `TOTAL` with every column summed over all filtered records.
- With no matching record: the header and `TOTAL,0,0,0,0,0,0.0000,0.0`.

## Programmatic Acceptance

=== AC targets-rows-and-order ===
Intent: One row per Target in ascending byte order, with per-Target counts derived from the fixture.
Suite: scoped
Requires: executable=go; scope=test

import csv, io, json, os, tempfile
from collections import Counter
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

names = ["beta", "Alpha", "alpha", "Zed", "Alpha", "beta", "beta"]
recs = [rec(i, t) for i, t in enumerate(names)]
code, out = run(recs, "targets")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"]
counts = Counter(names)
assert [r[0] for r in rows[1:-1]] == sorted(counts, key=lambda s: s.encode())
for r in rows[1:-1]:
    assert int(r[1]) == counts[r[0]]
assert rows[-1][0] == "TOTAL"
=== END AC targets-rows-and-order ===

=== AC targets-sums-cost-minutes ===
Intent: Per-Target failed, token, cost and minutes columns equal sums over the fixture records.
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
    rec(1, "Alpha", inp=1, cached=2, out=3, ms=70000, cost=0.25),
    rec(2, "Alpha", inp=4, cached=5, out=6, ms=80000, cost=0.5, status="failed"),
    rec(3, "Beta", inp=7, cached=8, out=9, ms=95000, cost=1.0),
]
st = lambda r: r["result"]["stats"]
code, out = run(recs, "targets")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
by = {r[0]: r for r in rows[1:-1]}
for t in {r["job"]["target"] for r in recs}:
    grp = [r for r in recs if r["job"]["target"] == t]
    row = by[t]
    assert int(row[2]) == len([r for r in grp if r["status"] != "succeeded"])
    assert int(row[3]) == sum(st(r)["input_tokens"] for r in grp)
    assert int(row[4]) == sum(st(r)["cached_input_tokens"] for r in grp)
    assert int(row[5]) == sum(st(r)["output_tokens"] for r in grp)
    assert abs(float(row[6]) - sum(st(r)["cost_usd"] for r in grp)) < 1e-9
    assert len(row[6].split(".")[1]) == 4
    assert float(row[7]) == ((sum(st(r)["elapsed_ms"] for r in grp) + 3000) // 6000) / 10
=== END AC targets-sums-cost-minutes ===

=== AC targets-total-row ===
Intent: The TOTAL row sums every column over all filtered records, including cost and minutes.
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
    rec(1, "Alpha", inp=1, cached=2, out=3, ms=2000, cost=0.25),
    rec(2, "Beta", inp=4, cached=5, out=6, ms=2000, cost=0.5, status="failed"),
    rec(3, "Gamma", inp=7, cached=8, out=9, ms=2000, cost=None),
]
st = lambda r: r["result"]["stats"]
code, out = run(recs, "targets")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
total = rows[-1]
assert total[0] == "TOTAL"
assert int(total[1]) == len(recs)
assert int(total[2]) == len([r for r in recs if r["status"] != "succeeded"])
assert int(total[3]) == sum(st(r)["input_tokens"] for r in recs)
assert int(total[4]) == sum(st(r)["cached_input_tokens"] for r in recs)
assert int(total[5]) == sum(st(r)["output_tokens"] for r in recs)
assert abs(float(total[6]) - sum(st(r)["cost_usd"] or 0 for r in recs)) < 1e-9
# The total minutes rounds the grand sum, not the sum of rounded rows.
assert float(total[7]) == ((sum(st(r)["elapsed_ms"] for r in recs) + 3000) // 6000) / 10
=== END AC targets-total-row ===

=== AC targets-empty-total ===
Intent: With no matching record the report is a header plus a zero TOTAL row with four-decimal cost.
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

code, out = run([rec(1, "Alpha")], "targets", "--target", "NoSuchTarget")
assert code == 0
assert out.endswith("\n") and "\r" not in out
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert rows[1][0] == "TOTAL"
assert [int(x) for x in rows[1][1:6]] == [0, 0, 0, 0, 0]
assert float(rows[1][6]) == 0.0 and len(rows[1][6].split(".")[1]) == 4
assert float(rows[1][7]) == 0.0 and len(rows[1][7].split(".")[1]) == 1
=== END AC targets-empty-total ===

=== AC targets-date-filter-applies ===
Intent: The date filter narrows the targets report.
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

since = "2026-10-15"
recs = [rec(1, "Alpha", "2026-10-14"), rec(2, "Beta", since), rec(3, "Gamma", "2026-10-16")]
code, out = run(recs, "targets", "--since", since)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
kept = [r for r in recs if r["started_at"][:10] >= since]
assert [r[0] for r in rows[1:-1]] == sorted({r["job"]["target"] for r in kept}, key=lambda s: s.encode())
assert int(rows[-1][1]) == len(kept)
=== END AC targets-date-filter-applies ===

## User Acceptance

- None.

## Guardrails

- Every value is computed from the log being read; no row, count, or total is hardcoded.
- Ordering compares text by byte value, never by locale collation.
- The log is never written to.
