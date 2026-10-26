# FEATURE: Log Skip

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Skip malformed, incomplete and duplicate log lines with a stderr warning and no effect on the exit code. |
| Depends On  | ARCHITECTURE.md, FEATURE-LogLines.md |
| Provides    | logread.Read skip rules, logread.Result.Skipped |
| Consumes    | logread.Read |

## Purpose

Make the reader tolerant. Bad lines are skipped and counted; they never fail the run.

## Behavior

`logread.Read` returns the kept records and `Result.Skipped`, the whole-file count of skipped lines. A line is skipped when:

- it is not a JSON object (invalid JSON, or valid JSON that is an array, string, number, boolean or null);
- any of `execution_id`, `started_at`, `status`, `job.command_name`, `job.target` is absent or not a non-empty string;
- its `execution_id` was already registered by an earlier kept line.

Further rules:

- Each skipped line writes one warning to the supplied warning writer naming its 1-based physical line number, counting blank lines.
- A skipped line does not register its `execution_id`; a malformed line therefore never causes a later line with that id to be skipped.
- The first occurrence of an `execution_id` is kept.
- Duplicate detection runs over the whole file inside `Read`, before any filter is applied elsewhere. A later repeat is skipped even when its first occurrence is later filtered out.
- Skipping never changes the exit code; `Read` returns a non-nil error only for an I/O failure.

## Required Go tests

TestSkipNotObject, TestSkipMissingField, TestSkipDuplicateKeepsFirst, TestMalformedDoesNotRegisterID, TestSkipWarningLineNumber, TestReadDedupesWholeFile, TestSkipEmptyStringField.

## Programmatic Acceptance

=== AC skip-not-object ===
Intent: A line that is not a JSON object is skipped and counted.

from drydock_go import run_test

result = run_test("./internal/logread", "TestSkipNotObject")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-not-object ===

=== AC skip-missing-field ===
Intent: A line lacking a required field is skipped and counted.

from drydock_go import run_test

result = run_test("./internal/logread", "TestSkipMissingField")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-missing-field ===

=== AC skip-empty-string-field ===
Intent: A required field holding an empty or non-string value causes the line to be skipped.

from drydock_go import run_test

result = run_test("./internal/logread", "TestSkipEmptyStringField")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-empty-string-field ===

=== AC skip-duplicate-first-kept ===
Intent: A repeated execution_id is skipped and the first occurrence is kept.

from drydock_go import run_test

result = run_test("./internal/logread", "TestSkipDuplicateKeepsFirst")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-duplicate-first-kept ===

=== AC skip-malformed-no-register ===
Intent: A malformed line does not register its execution_id, so a later valid line with that id is kept.

from drydock_go import run_test

result = run_test("./internal/logread", "TestMalformedDoesNotRegisterID")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-malformed-no-register ===

=== AC skip-warning-line-number ===
Intent: Each warning names the 1-based physical line number, counting blank lines.

from drydock_go import run_test

result = run_test("./internal/logread", "TestSkipWarningLineNumber")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-warning-line-number ===

=== AC skip-dedupe-whole-file ===
Intent: Duplicate detection covers the whole file and does not depend on any later filter.

from drydock_go import run_test

result = run_test("./internal/logread", "TestReadDedupesWholeFile")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC skip-dedupe-whole-file ===

## User Acceptance

- None.

## Guardrails

- Skipping never changes the exit code.
- A malformed line never registers its `execution_id`.
- Warnings go only to the warning writer (standard error in the CLI) and are never compared.
