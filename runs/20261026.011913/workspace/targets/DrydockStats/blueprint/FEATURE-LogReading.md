# FEATURE: Log Reading

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Read the Drydock JSON Lines log line by line, ignoring blank lines and numbering every line. |
| Depends On  | ARCHITECTURE.md |
| Provides    | logread.Read, logread.ReadFile |

## Purpose

The log is JSON Lines in UTF-8, one JSON object per line. This feature delivers the line reader in
`internal/logread`: it walks the input, discards blank lines, and hands each non-blank line, with
its 1-based line number, to the record decoder that the skip-rule features extend.

## Behavior

- `logread.Read(r io.Reader, warn io.Writer) (logread.Result, error)` reads every line of `r`.
- `logread.ReadFile(path string, warn io.Writer) (logread.Result, error)` opens the file, reads it
  with `Read`, and closes it. It returns a non-nil error when the file cannot be opened or read,
  including when the path names a directory.
- A blank line, meaning empty or whitespace only, is ignored and is neither a record nor counted as
  skipped.
- Line numbers are 1-based and count blank lines, so a warning names the physical line.
- Lines longer than the default `bufio.Scanner` token limit are read in full.
- The log file is opened read-only and never written.
- Each non-blank line is decoded as one JSON object into a `logread.Record`.

## Required Unit Tests

The build writes these tests in `internal/logread`:

| Test | Behavior |
|---|---|
| `TestReadIgnoresBlankLines` | Blank lines yield no record and no skipped count. |
| `TestReadNumbersEveryLine` | The line number used for a warning counts preceding blank lines. |
| `TestReadAcceptsLongLine` | A valid record line longer than 1 MiB is read as one record. |
| `TestReadFileMissing` | `ReadFile` on a nonexistent path returns an error. |
| `TestReadFileDirectory` | `ReadFile` on a directory returns an error. |
| `TestReadNeverWritesLog` | After `ReadFile`, the log file bytes and modification time are unchanged. |

## Programmatic Acceptance

=== AC log-read-blank-lines ===
Intent: Blank lines are ignored and not counted as records or skipped lines.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestReadIgnoresBlankLines", "TestReadNumbersEveryLine"):
    result = run_test("./internal/logread", name)
    print(name, result.status)
    print(result.stdout)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC log-read-blank-lines ===

=== AC log-read-long-line ===
Intent: A record line longer than the default scanner token limit is read in full.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestReadAcceptsLongLine"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC log-read-long-line ===

=== AC log-read-file-errors ===
Intent: ReadFile reports an error for a missing path and for a directory.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestReadFileMissing", "TestReadFileDirectory"):
    result = run_test("./internal/logread", name)
    print(name, result.status)
    print(result.stdout)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC log-read-file-errors ===

=== AC log-read-read-only ===
Intent: Reading a log leaves the log file unchanged.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

name = "TestReadNeverWritesLog"
result = run_test("./internal/logread", name)
print(name, result.status)
print(result.stdout)
assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC log-read-read-only ===

## User Acceptance

- None.

## Guardrails

- The log is opened read-only and is never written, truncated, or locked for writing.
- A blank line is never counted as a record or as a skipped line.
- Standard library only.
