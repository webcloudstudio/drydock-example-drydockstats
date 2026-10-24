# FEATURE: Log Numeric Stats

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Treat absent or null token, cost, and duration values as zero, and keep records with no stats object. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDecode.md |
| Provides    | record numeric stats |
| Consumes    | log record decoder |

## Purpose

Token counts, cost, and duration are summed by every report. A value that is absent or `null` contributes zero. A record with no `result.stats` object is still a record and contributes zero to every numeric total.

## Behavior

- `input_tokens`, `cached_input_tokens`, `output_tokens`, `elapsed_ms` are read as integers; absent or `null` is zero.
- `cost_usd` is read as float64; absent or `null` is zero.
- A record without `result.stats` yields all-zero numeric stats and is still counted.

## Programmatic Acceptance

=== AC log-numeric-null-cost-is-zero ===
Intent: A cost_usd of null contributes zero to cost totals.
Requires: executable=python3; scope=test

record = {"result": {"stats": {"cost_usd": None, "input_tokens": 5}}}
cost = record["result"]["stats"].get("cost_usd")
total = float(cost) if cost is not None else 0.0
assert total == 0.0
=== END AC log-numeric-null-cost-is-zero ===

=== AC log-numeric-absent-tokens-are-zero ===
Intent: Absent token and duration fields are read as zero.
Requires: executable=python3; scope=test

record = {"result": {"stats": {"input_tokens": 7}}}
stats = record["result"]["stats"]
assert stats.get("output_tokens", 0) == 0
assert stats.get("cached_input_tokens", 0) == 0
assert stats.get("elapsed_ms", 0) == 0
assert stats["input_tokens"] == 7
=== END AC log-numeric-absent-tokens-are-zero ===

=== AC log-numeric-missing-stats-object ===
Intent: A record with no result.stats object is still a record with all-zero numeric stats.
Requires: executable=python3; scope=test

record = {"job": {"command_name": "plan"}, "result": {"returncode": 1}}
stats = record.get("result", {}).get("stats", {})
assert stats.get("input_tokens", 0) == 0
assert stats.get("cost_usd", 0) == 0
=== END AC log-numeric-missing-stats-object ===
