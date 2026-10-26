# FEATURE: Date Filter

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Keep only records whose start date falls within the inclusive --since and --until bounds, for every report. |
| Depends On  | FEATURE-TargetFilter.md, FEATURE-DateOptions.md |
| Provides    | report.FilterDates |
| Consumes    | report.FilterTarget, cli date validation |

## Purpose

`--since <date>` and `--until <date>` narrow every report to records whose start date lies on or
after `--since` and on or before `--until`. Both bounds are inclusive and optional.

## Behavior

- A record's start date is the first 10 characters of `started_at`.
- Dates are `YYYY-MM-DD`, so comparison as strings equals chronological comparison.
- The time of day never affects the outcome: a record at `23:59:59.999Z` on the `--until` date is kept.
- `--since` and `--until` may be combined with each other and with `--target`; all conditions must hold.
- A range matching nothing yields a header and, where the report has one, a `TOTAL` row of zeros.
- The `skipped` count in `summary` is whole-file and not affected by the date filter.
- The filter is a pure function over the record slice (`report.FilterDates`), unit-tested apart from parsing and formatting.

## Programmatic Acceptance

=== AC date-since-until-inclusive ===
Intent: Both bounds are inclusive and records outside them are excluded.
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

days = ["2026-10-13", "2026-10-14", "2026-10-14", "2026-10-15", "2026-10-16", "2026-10-17"]
recs = [rec(i, day=dy) for i, dy in enumerate(days)]
since, until = "2026-10-14", "2026-10-16"
code, out = run(recs, "daily", "--since", since, "--until", until)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
data = [r for r in rows[1:] if r[0] != "TOTAL"]
expected_days = sorted({dy for dy in days if since <= dy <= until})
assert [r[0] for r in data] == expected_days
for r in data:
    assert int(r[1]) == days.count(r[0])
total = rows[-1]
assert total[0] == "TOTAL"
assert int(total[1]) == len([dy for dy in days if since <= dy <= until])
=== END AC date-since-until-inclusive ===

=== AC date-one-sided-bounds ===
Intent: --since alone and --until alone each bound only one side.
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

days = ["2026-10-13", "2026-10-14", "2026-10-15", "2026-10-15"]
recs = [rec(i, day=dy) for i, dy in enumerate(days)]
bound = "2026-10-14"

code, out = run(recs, "summary", "--since", bound)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == len([dy for dy in days if dy >= bound])
assert d["first"] == min(dy for dy in days if dy >= bound)

code, out = run(recs, "summary", "--until", bound)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == len([dy for dy in days if dy <= bound])
assert d["last"] == max(dy for dy in days if dy <= bound)
=== END AC date-one-sided-bounds ===

=== AC date-time-of-day-ignored ===
Intent: Only the date part of started_at decides membership, not the time of day.
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

day = "2026-10-15"
recs = [rec(1, ts=day + "T00:00:00.000Z"), rec(2, ts=day + "T23:59:59.999Z"), rec(3, day="2026-10-14", ts="2026-10-14T23:59:59.999Z"), rec(4, ts="2026-10-16T00:00:00.000Z")]
code, out = run(recs, "summary", "--since", day, "--until", day)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == len([r for r in recs if r["started_at"][:10] == day])
assert d["first"] == day and d["last"] == day
=== END AC date-time-of-day-ignored ===

=== AC date-every-report ===
Intent: The date filter narrows all six reports and leaves the skipped count untouched.
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

day = "2026-10-15"
recs = [rec(1, day="2026-10-14", status="failed"), rec(2, day=day), rec(3, day=day, status="failed"), rec(4, day="2026-10-16", status="failed")]
bad = ["{broken", "[1, 2]"]
kept = [r for r in recs if r["started_at"][:10] == day]
failed_kept = [r for r in kept if r["status"] != "succeeded"]
args = ("--since", day, "--until", day)

for report in ["commands", "targets", "models", "daily"]:
    code, out = run(recs + bad, report, *args)
    assert code == 0, report
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[-1][0] == "TOTAL", report
    assert int(rows[-1][rows[0].index("calls")]) == len(kept), report

code, out = run(recs + bad, "summary", *args)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
d = dict(zip(rows[0], rows[1]))
assert int(d["records"]) == len(kept)
assert int(d["skipped"]) == len(bad)

code, out = run(recs + bad, "failures", *args)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) - 1 == len(failed_kept)
i = rows[0].index("started_at")
assert all(r[i][:10] == day for r in rows[1:])
=== END AC date-every-report ===

=== AC date-with-target-and-empty-range ===
Intent: Date and Target filters combine, and a range matching nothing yields a zero TOTAL row.
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

name, day = "Alpha", "2026-10-15"
recs = [rec(1, name, day), rec(2, "Beta", day), rec(3, name, "2026-10-16"), rec(4, name, day)]
code, out = run(recs, "commands", "--target", name, "--since", day, "--until", day)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
expected = len([r for r in recs if r["job"]["target"] == name and r["started_at"][:10] == day])
assert int(rows[-1][1]) == expected

code, out = run(recs, "commands", "--since", "2030-01-01")
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2 and rows[1][0] == "TOTAL"
assert [int(x) for x in rows[1][1:6]] == [0, 0, 0, 0, 0]
=== END AC date-with-target-and-empty-range ===

## User Acceptance

- None.

## Guardrails

- The start date is always the first 10 characters of `started_at`; the filter must not parse the time zone or shift dates.
- The `skipped` count must never depend on `--since` or `--until`.
- The filter must not modify, rewrite, or write to the log being read.
