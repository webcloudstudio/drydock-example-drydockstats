# FEATURE: Start Date Filter

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Keep records whose start date falls within the inclusive --since and --until range. |
| Depends On  | ARCHITECTURE.md, FEATURE-FilterTarget.md |
| Provides    | date range filter |
| Consumes    | record stream after target filter |

## Purpose

`--since` and `--until` keep records whose start date lies in an inclusive range. The start date is the first 10 characters of `started_at`. The CLI validates the option values before this filter runs.

## Behavior

- Start date = first `min(10, len(started_at))` characters of `started_at`.
- With `--since D`, keep records whose start date is greater than or equal to `D`.
- With `--until D`, keep records whose start date is less than or equal to `D`.
- Comparison is byte-wise on `YYYY-MM-DD` text, which orders the same as calendar dates.
- Options that are not given do not constrain the range.

## Programmatic Acceptance

=== AC filter-date-inclusive-range ===
Intent: Records on the since and until boundaries are kept; records outside the range are dropped.
Requires: executable=python3; scope=test

def start_date(started_at):
    return started_at[:10]

def keep(record, since, until):
    d = start_date(record["started_at"])
    if since is not None and d < since:
        return False
    if until is not None and d > until:
        return False
    return True

records = [
    {"started_at": "2026-10-13T23:59:59.000Z"},
    {"started_at": "2026-10-14T00:00:00.000Z"},
    {"started_at": "2026-10-15T12:00:00.000Z"},
    {"started_at": "2026-10-16T00:00:00.000Z"},
    {"started_at": "2026-10-17T00:00:00.000Z"},
]
kept = [r["started_at"] for r in records if keep(r, "2026-10-14", "2026-10-16")]
assert kept == [
    "2026-10-14T00:00:00.000Z",
    "2026-10-15T12:00:00.000Z",
    "2026-10-16T00:00:00.000Z",
]
=== END AC filter-date-inclusive-range ===

=== AC filter-date-open-ended ===
Intent: Omitting one bound leaves the range open on that side.
Requires: executable=python3; scope=test

def start_date(started_at):
    return started_at[:10]

def keep(record, since, until):
    d = start_date(record["started_at"])
    if since is not None and d < since:
        return False
    if until is not None and d > until:
        return False
    return True

records = [
    {"started_at": "2026-10-13T00:00:00.000Z"},
    {"started_at": "2026-10-20T00:00:00.000Z"},
]
assert [r for r in records if keep(r, "2026-10-14", None)] == [records[1]]
assert [r for r in records if keep(r, None, "2026-10-14")] == [records[0]]
=== END AC filter-date-open-ended ===
