# FEATURE: Malformed Lines

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Skip lines that are not valid records, count them, and warn on standard error without failing the run. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogReading.md |
| Provides    | logread.Result.Skipped |

## Purpose

A malformed line must never fail the run. This feature adds the record validation rules to the
reader: a line that fails any rule is skipped, counted in `Result.Skipped`, and reported with one
warning on the warning writer naming its line number. Skipping never changes the exit code.

## Behavior

A non-blank line is skipped when any of these holds:

- It is not a JSON object, including a truncated line, a JSON array, a bare scalar, and `null`.
- It is not valid UTF-8; it is treated as not a JSON object.
- Any of `execution_id`, `started_at`, `status`, `job.command_name`, or `job.target` is absent or
  not a non-empty string.
- `started_at` does not begin with a real calendar date written `YYYY-MM-DD`, for example
  `2026-13-01` or `2026-02-30`, because every report keys on the start date.

Each skipped line writes exactly one warning to the warning writer that contains its 1-based line
number. A skipped line never registers its `execution_id`, so a later valid record with that id is
kept (see `FEATURE-DuplicateIds.md`). The skipped count covers the whole file and is independent of
any filter applied later.

## Required Unit Tests

The build writes these tests in `internal/logread`:

| Test | Behavior |
|---|---|
| `TestSkipNotJSONObject` | Truncated JSON, an array, a scalar, and `null` are each skipped. |
| `TestSkipInvalidUTF8` | A line with invalid UTF-8 bytes is skipped. |
| `TestSkipMissingRequiredField` | For each of the five required fields, absent, empty, and non-string values skip the line. |
| `TestSkipBadStartedAt` | A `started_at` without a real leading `YYYY-MM-DD` date skips the line. |
| `TestSkipWarnsWithLineNumber` | One warning per skipped line names its physical line number, counting blank lines. |
| `TestSkipDoesNotAffectValidRecords` | Valid lines around skipped lines are all kept, in file order. |

## Programmatic Acceptance

=== AC malformed-not-object ===
Intent: Lines that are not JSON objects are skipped and counted rather than failing the read.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestSkipNotJSONObject"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC malformed-not-object ===

=== AC malformed-invalid-utf8 ===
Intent: A line that is not valid UTF-8 is skipped as not a JSON object.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestSkipInvalidUTF8"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC malformed-invalid-utf8 ===

=== AC malformed-required-fields ===
Intent: A line lacking a non-empty string in any required field is skipped.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestSkipMissingRequiredField"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC malformed-required-fields ===

=== AC malformed-started-at ===
Intent: A started_at that does not begin with a real calendar date skips the line.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestSkipBadStartedAt"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC malformed-started-at ===

=== AC malformed-warning-and-count ===
Intent: Each skipped line is warned about once with its line number and valid neighbours are kept.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestSkipWarnsWithLineNumber", "TestSkipDoesNotAffectValidRecords"):
    result = run_test("./internal/logread", name)
    print(name, result.status)
    print(result.stdout)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC malformed-warning-and-count ===

## User Acceptance

- None.

## Guardrails

- Skipping never changes the exit code and never writes to standard output.
- A skipped line never registers its `execution_id`.
- Warnings go to standard error only and are never compared.
- Standard library only.
