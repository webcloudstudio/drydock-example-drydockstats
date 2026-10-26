# FEATURE: Report Daily

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Writes the daily CSV report with one row per start date and a TOTAL row. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliExit.md, FEATURE-ValueFormat.md |
| Provides    | ddkstat daily |
| Consumes    | cli.Register, format.Minutes, logread.Record |

## Purpose

The `daily` report shows, for each start date, how many calls ran, how many failed, how many tokens they used, and how many minutes they took.

## Behavior

- Header: `date,calls,failed,tokens,minutes`.
- The report covers the records that remain after the filters are applied.
- One row per start date (the first 10 characters of `started_at`), in ascending byte order.
- `calls` is the record count. `failed` is the count of records whose status is not `succeeded`.
- `tokens` is input plus cached plus output tokens, summed over the group.
- `minutes` is the group's summed `elapsed_ms` formatted by `format.Minutes`.
- A final `TOTAL` row covers all filtered records.
- When no record matches, the report writes the header and a `TOTAL` row of zeros with `0.0` minutes.
- Aggregation (`AggregateDaily`) and writing (`WriteDaily`) are separate functions in `internal/report`.
- The writer is registered with `cli.Register("daily", ...)`.
- Output is CSV, `\n` line endings, values unquoted.

## Programmatic Acceptance

=== AC report-daily-unit ===
Intent: The native unit tests for daily aggregation, ordering, totals, empty input and writing all pass.
Requires: executable=go; scope=test

from drydock_go import run_test

for name in (
    "TestDailyOrdering",
    "TestDailyTokensSum",
    "TestDailyTotal",
    "TestDailyEmpty",
    "TestDailyWrite",
):
    result = run_test("./internal/report", name)
    print(name, result.status)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC report-daily-unit ===

=== AC report-daily-rows ===
Intent: The daily command emits one row per start date with calls, failed, tokens and minutes derived from the log, followed by a TOTAL row.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from decimal import Decimal

from drydock_go import cli_runner


def rec(n, started, cmd, target, status, tin, tc, tout, ms):
    return {
        "schema_version": 1,
        "execution_id": "id-%03d" % n,
        "started_at": started,
        "status": status,
        "job": {"command_name": cmd, "target": target, "llm": "claude", "model": "haiku"},
        "result": {
            "returncode": 0 if status == "succeeded" else 1,
            "timed_out": False,
            "stats": {
                "input_tokens": tin,
                "cached_input_tokens": tc,
                "output_tokens": tout,
                "elapsed_ms": ms,
                "cost_usd": 0.5,
                "model": "claude-haiku-5-5",
            },
        },
    }


recs = [
    rec(1, "2026-10-15T08:00:00.000Z", "build", "Alpha", "succeeded", 10, 200, 30, 95000),
    rec(2, "2026-10-14T09:00:00.000Z", "plan", "Beta", "failed", 1, 2, 3, 20000),
    rec(3, "2026-10-14T10:00:00.000Z", "build", "Alpha", "succeeded", 100, 0, 7, 61000),
    rec(4, "2026-10-16T11:00:00.000Z", "analyze", "Beta", "succeeded", 5, 5, 5, 0),
    rec(5, "2026-10-15T12:00:00.000Z", "build", "Beta", "succeeded", 0, 40, 0, 125000),
]

work = tempfile.mkdtemp()
path = os.path.join(work, "log.jsonl")
with open(path, "w", encoding="utf-8") as fh:
    for r in recs:
        fh.write(json.dumps(r) + "\n")


def expected(group):
    total_ms = sum(r["result"]["stats"]["elapsed_ms"] for r in group)
    tenths = (total_ms + 3000) // 6000
    return (
        len(group),
        sum(1 for r in group if r["status"] != "succeeded"),
        sum(
            r["result"]["stats"]["input_tokens"]
            + r["result"]["stats"]["cached_input_tokens"]
            + r["result"]["stats"]["output_tokens"]
            for r in group
        ),
        Decimal(tenths) / Decimal(10),
    )


code, out, err = cli_runner("daily", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
header, body = rows[0], rows[1:]
assert header == ["date", "calls", "failed", "tokens", "minutes"]

dates = sorted({r["started_at"][:10] for r in recs})
assert [row[0] for row in body[:-1]] == dates
for row in body[:-1]:
    group = [r for r in recs if r["started_at"][:10] == row[0]]
    calls, failed, tokens, minutes = expected(group)
    assert int(row[1]) == calls
    assert int(row[2]) == failed
    assert int(row[3]) == tokens
    assert Decimal(row[4]) == minutes

total = body[-1]
calls, failed, tokens, minutes = expected(recs)
assert total[0] == "TOTAL"
assert int(total[1]) == calls
assert int(total[2]) == failed
assert int(total[3]) == tokens
assert Decimal(total[4]) == minutes
=== END AC report-daily-rows ===

=== AC report-daily-filters ===
Intent: The daily command honours --target, --since and --until, computing every row and the TOTAL over only the matching records.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from decimal import Decimal

from drydock_go import cli_runner


def rec(n, started, target, status, tin, ms):
    return {
        "schema_version": 1,
        "execution_id": "id-%03d" % n,
        "started_at": started,
        "status": status,
        "job": {"command_name": "build", "target": target, "llm": "claude", "model": "haiku"},
        "result": {
            "returncode": 0,
            "timed_out": False,
            "stats": {
                "input_tokens": tin,
                "cached_input_tokens": 0,
                "output_tokens": 0,
                "elapsed_ms": ms,
                "cost_usd": 0.1,
                "model": "claude-haiku-5-5",
            },
        },
    }


recs = [
    rec(1, "2026-10-14T01:00:00.000Z", "Alpha", "succeeded", 11, 60000),
    rec(2, "2026-10-15T01:00:00.000Z", "Alpha", "failed", 22, 60000),
    rec(3, "2026-10-15T02:00:00.000Z", "Beta", "succeeded", 33, 60000),
    rec(4, "2026-10-16T01:00:00.000Z", "Alpha", "succeeded", 44, 60000),
]

work = tempfile.mkdtemp()
path = os.path.join(work, "log.jsonl")
with open(path, "w", encoding="utf-8") as fh:
    for r in recs:
        fh.write(json.dumps(r) + "\n")

target = "Alpha"
since = "2026-10-15"
until = "2026-10-16"
kept = [
    r for r in recs
    if r["job"]["target"] == target and since <= r["started_at"][:10] <= until
]

code, out, err = cli_runner("daily", "--target", target, "--since", since, "--until", until, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
body = rows[1:]
assert [row[0] for row in body[:-1]] == sorted({r["started_at"][:10] for r in kept})
assert body[-1][0] == "TOTAL"
assert int(body[-1][1]) == len(kept)
assert int(body[-1][2]) == sum(1 for r in kept if r["status"] != "succeeded")
assert int(body[-1][3]) == sum(r["result"]["stats"]["input_tokens"] for r in kept)
tenths = (sum(r["result"]["stats"]["elapsed_ms"] for r in kept) + 3000) // 6000
assert Decimal(body[-1][4]) == Decimal(tenths) / Decimal(10)
=== END AC report-daily-filters ===

=== AC report-daily-empty ===
Intent: When no record matches the filter, the daily command exits 0 and writes the header and a TOTAL row of zeros.
Requires: executable=go; scope=test

import csv
import io
import json
import os
import tempfile
from decimal import Decimal

from drydock_go import cli_runner

record = {
    "schema_version": 1,
    "execution_id": "id-001",
    "started_at": "2026-10-14T01:00:00.000Z",
    "status": "succeeded",
    "job": {"command_name": "build", "target": "Alpha", "llm": "claude", "model": "haiku"},
    "result": {"returncode": 0, "timed_out": False, "stats": {"input_tokens": 9, "elapsed_ms": 60000}},
}
work = tempfile.mkdtemp()
path = os.path.join(work, "log.jsonl")
with open(path, "w", encoding="utf-8") as fh:
    fh.write(json.dumps(record) + "\n")

code, out, err = cli_runner("daily", "--target", "NoSuchTarget", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["date", "calls", "failed", "tokens", "minutes"]
assert len(rows) == 2
total = rows[1]
assert total[0] == "TOTAL"
assert int(total[1]) == 0
assert int(total[2]) == 0
assert int(total[3]) == 0
assert Decimal(total[4]) == Decimal(0)
=== END AC report-daily-empty ===

## User Acceptance

- None.

## Guardrails

- The report reads only values computed from the log; no row, count or total is hardcoded.
- The report never writes to the log.
- Text values are written without quoting or escaping.
