# ARCHITECTURE: DrydockStats

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Package layout, boundaries and pipeline of the ddkstat command-line reporter for Drydock's LLM execution log. |
| Depends On  | None |
| Provides    | go.mod, internal/, cmd/ddkstat/, bin/ |
| Consumes    | None |

## Purpose

ddkstat reads a JSON Lines log (`logs/llm.jsonl`) and writes one of six CSV reports to standard output. It is a single-process, single-threaded CLI. It has no database, no network access and no configuration files. It never writes to the log it reads.

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`:

| Technology | Where it applies |
|---|---|
| Go 1.22 or newer, standard library only | All application code under `cmd/` and `internal/`. |
| go test and go vet | Test runner behind `bin/test.sh`; tests live beside the code in `*_test.go`. |
| JSON Lines | Input format, decoded with `encoding/json`, one object per physical line. |
| POSIX sh | `bin/test.sh` and the supplied grader `sources/run.sh`. |

No third-party module is permitted. `go.mod` declares the module `ddkstat`, a `go` directive of 1.22 or newer, and no `require` lines.

## Layout

```text
go.mod                    module ddkstat
cmd/ddkstat/main.go       package main; the only caller of os.Exit
internal/logread/         read, validate, dedupe and normalize log records
internal/filter/          pure record filters (target, date range)
internal/format/          minutes and cost text formatting
internal/report/          one aggregate-and-write function per report
internal/cli/             Run, argument parsing, report registry
bin/test.sh               go vet ./... then go test ./...
sources/                  read-only scoring assets; never touched by application code
```

## Module Ownership

| Module | Owns | May access |
|---|---|---|
| `internal/logread` | Opening-agnostic reading of an `io.Reader`, line classification, skip and duplicate rules, numeric normalization, model attribution | `encoding/json`, `bufio`, `io`, `strings` |
| `internal/filter` | Target and date-range predicates over `[]logread.Record` | `logread` types only |
| `internal/format` | `Minutes` and `Cost` text | `fmt` |
| `internal/report` | Aggregation and CSV writing for each report | `logread`, `format`, `io` |
| `internal/cli` | Argument parsing, exit-code mapping, opening the log file read-only, report registry | `os`, `flag`, `time`, all internal packages |
| `cmd/ddkstat` | Process start and `os.Exit` | `internal/cli` |

Persistence, configuration files, file stores and external services: none. The log file is opened read-only by `internal/cli` and passed to `logread` as an `io.Reader`.

## Pipeline

Parse arguments, read the log, drop skipped and duplicate lines, filter, aggregate, format. Parsing, aggregating and formatting are separate functions, each unit-tested.

## Shared Types and Interfaces

`logread.Read(r io.Reader, warn io.Writer) (Result, error)` returns `Result{Records []Record, Skipped int}`. The error is non-nil only when reading from `r` fails.

`logread.Record` carries: `ExecutionID`, `StartedAt`, `Date` (first 10 characters of `StartedAt`), `Status`, `Failed`, `Command`, `Target`, `Provider`, `Model`, `InputTokens`, `CachedTokens`, `OutputTokens` (all `int64`), `CostUSD` (`float64`), `ElapsedMS` (`int64`), `ReturnCode` (`*int`, nil when null or absent) and `TimedOut` (`bool`).

`cli.Run(args []string, stdout, stderr io.Writer) int` is the single entry point; exit codes are 0 success, 1 unreadable log, 2 usage error. Report writers register with `cli.Register(name, fn)` where `fn(w io.Writer, records []logread.Record, skipped int) error`.

## Technical Decisions

- Text is compared by byte value; CSV values are written without quoting.
- A failing command writes nothing to standard output and one diagnostic line to standard error; stdout is buffered and flushed only on success.
- Diagnostics (skip warnings, usage line, errors) go to standard error and are never compared.
- No report answer, count or total is embedded in application code; every value comes from the log being read.

## Programmatic Acceptance

=== AC arch-go-mod ===
Intent: The root go.mod declares module ddkstat, Go 1.22 or newer, and no third-party requirement.

from pathlib import Path

text = Path("go.mod").read_text()
lines = [ln.strip() for ln in text.splitlines()]
assert any(ln.startswith("module ") for ln in lines), text
go_lines = [ln for ln in lines if ln.startswith("go ")]
assert len(go_lines) == 1, text
major, minor = go_lines[0].split()[1].split(".")[:2]
assert (int(major), int(minor)) >= (1, 22), go_lines
assert not any(ln.startswith("require") for ln in lines), text
=== END AC arch-go-mod ===

=== AC arch-layout ===
Intent: The directories every later story builds into exist.

from pathlib import Path

for d in ("cmd/ddkstat", "internal", "bin"):
    assert Path(d).is_dir(), d
=== END AC arch-layout ===

=== AC arch-no-vendor ===
Intent: The module is standard-library only and has no go.sum dependency entries.

from pathlib import Path

sum_path = Path("go.sum")
if sum_path.exists():
    assert sum_path.read_text().strip() == "", sum_path.read_text()
assert not Path("vendor").exists()
=== END AC arch-no-vendor ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module and no database.
- Only `cmd/ddkstat/main.go` calls `os.Exit`.
- Application code never reads `sources/run.sh` and never embeds a report row, count or total.
- The log being read is never written to.
- Nothing under `sources/` is modified, regenerated or substituted.
