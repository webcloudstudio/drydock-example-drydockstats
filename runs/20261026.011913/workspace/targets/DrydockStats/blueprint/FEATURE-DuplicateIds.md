# FEATURE: Duplicate Ids

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Skip any line whose execution_id already appeared on an earlier kept record, keeping the first occurrence. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogReading.md, FEATURE-MalformedLines.md |
| Provides    | logread.Dedupe |

## Purpose

Drydock may write the same execution twice. This feature makes the reader keep the first
occurrence of each `execution_id` and skip every later repeat, counting it in `Result.Skipped` and
warning on the warning writer with its line number, without failing the run.

## Behavior

- A line is skipped as a duplicate when its `execution_id` equals that of an earlier kept record.
- The first occurrence is kept, and later repeats are skipped even when their other fields differ.
- Only a record that passes every other rule registers its `execution_id`. A line skipped as
  malformed does not register its id, so a later valid record with the same id is kept.
- Each duplicate writes one warning naming its 1-based line number.
- Duplicate detection is exact string equality on `execution_id`.
- The skipped count includes duplicates and malformed lines together, over the whole file.

## Required Unit Tests

The build writes these tests in `internal/logread`:

| Test | Behavior |
|---|---|
| `TestDuplicateFirstOccurrenceKept` | Two lines with one id yield one record, the first, and one skipped. |
| `TestDuplicateKeepsFirstEvenWhenFieldsDiffer` | The kept record carries the first line's field values. |
| `TestSkippedLineDoesNotRegisterID` | A malformed line followed by a valid line with the same id keeps the valid line. |
| `TestDuplicateWarnsWithLineNumber` | Each duplicate produces one warning naming its physical line number. |
| `TestDuplicateCountedWithMalformed` | Duplicates and malformed lines are both counted in `Skipped`. |

## Programmatic Acceptance

=== AC duplicate-first-kept ===
Intent: The first occurrence of an execution_id is kept and later repeats are skipped and counted.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestDuplicateFirstOccurrenceKept", "TestDuplicateKeepsFirstEvenWhenFieldsDiffer"):
    result = run_test("./internal/logread", name)
    print(name, result.status)
    print(result.stdout)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC duplicate-first-kept ===

=== AC duplicate-skipped-line-does-not-register ===
Intent: A line skipped for another reason does not register its id, so a later valid record with that id is kept.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestSkippedLineDoesNotRegisterID"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC duplicate-skipped-line-does-not-register ===

=== AC duplicate-warning-and-count ===
Intent: Each duplicate warns once with its line number and is counted together with malformed lines.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestDuplicateWarnsWithLineNumber", "TestDuplicateCountedWithMalformed"):
    result = run_test("./internal/logread", name)
    print(name, result.status)
    print(result.stdout)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC duplicate-warning-and-count ===

## User Acceptance

- None.

## Guardrails

- The first occurrence of an `execution_id` is always the one kept.
- Skipping a duplicate never changes the exit code and never writes to standard output.
- Standard library only.
