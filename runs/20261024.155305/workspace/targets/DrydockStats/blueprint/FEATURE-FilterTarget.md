# FEATURE: Target Filter

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Keep only records whose job.target equals the --target value exactly. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogDuplicates.md |
| Provides    | target filter |
| Consumes    | record stream after skip rules |

## Purpose

The `--target` option narrows every report to records for one Target. The comparison is exact and case-sensitive. When the option is absent every record passes. The filter receives records only after the skip rules have run, so the whole-file `skipped` count is unaffected.

## Behavior

- With `--target X`, a record passes when `job.target == X` as a byte-exact string comparison.
- Without `--target`, every record passes.
- The filter never changes the skipped count.

## Programmatic Acceptance

=== AC filter-target-exact-match ===
Intent: Only records whose job.target equals the requested name exactly pass the filter.
Requires: executable=python3; scope=test

def keep(record, target):
    return target is None or record["job"]["target"] == target

records = [
    {"job": {"target": "Toml"}},
    {"job": {"target": "toml"}},
    {"job": {"target": "Toml "}},
]
kept = [r for r in records if keep(r, "Toml")]
assert kept == [{"job": {"target": "Toml"}}]
=== END AC filter-target-exact-match ===

=== AC filter-target-absent-passes-all ===
Intent: When no target is given every record passes the filter.
Requires: executable=python3; scope=test

def keep(record, target):
    return target is None or record["job"]["target"] == target

records = [{"job": {"target": "A"}}, {"job": {"target": "B"}}]
assert [r for r in records if keep(r, None)] == records
=== END AC filter-target-absent-passes-all ===
