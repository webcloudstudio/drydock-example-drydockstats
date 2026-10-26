# FEATURE: Target Filter

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Keep only records whose job.target equals the --target value exactly, for every report. |
| Depends On  | FEATURE-ModelAttribution.md, FEATURE-CommandLine.md |
| Provides    | report.FilterTarget |
| Consumes    | logread.Record.Model, ddkstat <report> [--target] [--since] [--until] <log> |

## Purpose

`--target <name>` narrows every report to the records whose `job.target` equals `<name>`. The
comparison is exact and case-sensitive, with no trimming and no prefix or pattern matching.

## Behavior

- The filter runs after ingestion (parse, skip, dedupe) and before aggregation, for all six reports.
- A record is kept only when `job.target == name` byte for byte.
- The `skipped` count in `summary` is computed over the whole file and is not affected by the filter.
- When no record matches, each report writes its header and, where it has one, a `TOTAL` row of zeros.
- Without `--target`, no Target filtering happens.
- The filter is a pure function over the record slice (`report.FilterTarget`) and is unit-tested apart from parsing and formatting.

## Programmatic Acceptance

=== AC target-summary-counts ===
Intent: The summary report counts only the named Target's records while skipped still counts the whole file.
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

name = "Alpha"
recs = [rec(1, name), rec(2, name), rec(3, "Beta"), rec(4, "Gamma")]
bad_lines = ["this is not json"]
code, out = run(recs + bad_lines, "summary", "--target", name)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
expected = len([r for r in recs if r["job"]["target"] == name])
assert int(d["records"]) == expected
assert int(d["skipped"]) == len(bad_lines)
=== END AC target-summary-counts ===

=== AC target-exact-match ===
Intent: The Target comparison is exact and case-sensitive, so near-miss names are excluded.
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

name = "Alpha"
names = [name, name, "alpha", "Alpha2", "Alph", "ALPHA"]
recs = [rec(i, t) for i, t in enumerate(names)]
code, out = run(recs, "targets", "--target", name)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
data = [r for r in rows[1:] if r[0] != "TOTAL"]
assert [r[0] for r in data] == [name]
assert int(data[0][1]) == names.count(name)
=== END AC target-exact-match ===

=== AC target-no-match ===
Intent: A Target matching no record yields a header and a zero TOTAL row with exit zero.
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

recs = [rec(1, "Alpha"), rec(2, "Beta")]
code, out = run(recs, "commands", "--target", "NoSuchTarget")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert rows[1][0] == "TOTAL"
assert [int(x) for x in rows[1][1:6]] == [0, 0, 0, 0, 0]
assert float(rows[1][6]) == 0.0
=== END AC target-no-match ===

=== AC target-every-report ===
Intent: The Target filter narrows all six reports, not only one of them.
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

name = "Alpha"
recs = [rec(1, name), rec(2, name, status="failed"), rec(3, "Beta"), rec(4, "Beta", status="failed"), rec(5, "Gamma", day="2026-10-15")]
kept = [r for r in recs if r["job"]["target"] == name]
failed_kept = [r for r in kept if r["status"] != "succeeded"]

for report in ["commands", "targets", "models", "daily"]:
    code, out = run(recs, report, "--target", name)
    assert code == 0, report
    rows = list(csv.reader(io.StringIO(out)))
    total = rows[-1]
    assert total[0] == "TOTAL", report
    idx = rows[0].index("calls")
    assert int(total[idx]) == len(kept), report

code, out = run(recs, "summary", "--target", name)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert int(dict(zip(rows[0], rows[1]))["records"]) == len(kept)

code, out = run(recs, "failures", "--target", name)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
t = rows[0].index("target")
assert len(rows) - 1 == len(failed_kept)
assert all(r[t] == name for r in rows[1:])
=== END AC target-every-report ===

## User Acceptance

- None.

## Guardrails

- The filter must not modify, rewrite, or write to the log being read.
- The `skipped` count must never depend on `--target`.
- No Target name, count, or total may be hardcoded; every value comes from the log being read.
