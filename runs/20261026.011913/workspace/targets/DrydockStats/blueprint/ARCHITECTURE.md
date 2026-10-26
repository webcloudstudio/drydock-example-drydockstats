# ARCHITECTURE: DrydockStats

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Module layout, package boundaries, shared types, and technical decisions for the ddkstat LLM log reporter. |
| Depends On  | None |
| Provides    | go.mod, cmd/ddkstat, internal/cli, internal/logread, internal/report |

## Purpose

`ddkstat` is a command-line reporter for Drydock's LLM execution log (`logs/llm.jsonl`). It reads a
JSON Lines file, skips malformed and duplicate lines, narrows the records by optional Target and
date filters, and writes one of six CSV reports to standard output. It has no database and no
network access, and it never writes to the log it reads.

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`, which remains the decision of record.

| Technology | Where it applies |
|---|---|
| Go 1.22 or newer, standard library only | All application code; no third-party modules |
| `go test` and `go vet` | The test runner behind `bin/test.sh` |
| JSON Lines | The input format, decoded with `encoding/json` |
| POSIX sh | `bin/test.sh` and the supplied grader `sources/run.sh` |

## Module Layout

```text
go.mod                    module declaration, Go 1.22 or newer, no require lines
cmd/ddkstat/main.go       package main; the only caller of os.Exit
internal/cli/             Run, argument parsing, exit codes, usage line
internal/logread/         reading, skip rules, de-duplication, record model
internal/report/          filters, aggregation, and CSV formatting as separate functions
bin/test.sh               POSIX launcher for go vet and go test
```

## Package Boundaries

| Package | Owns | May import |
|---|---|---|
| `cmd/ddkstat` | Process start; passes `os.Args[1:]`, `os.Stdout`, `os.Stderr` to `cli.Run`; exits with its result | `internal/cli` |
| `internal/cli` | Argument parsing, date validation, exit codes, orchestration of read, filter, aggregate, format | `internal/logread`, `internal/report` |
| `internal/logread` | Opening and reading the log, line numbering, skip and duplicate rules, the `Record` model | standard library |
| `internal/report` | Filtering, aggregation per report, minutes and cost formatting, CSV writing | `internal/logread` |

Parsing the log, aggregating, and formatting are separate functions, each unit-tested. Business
logic lives in `internal/logread` and `internal/report`; `internal/cli` only wires them together.

## Shared Types and Interfaces

```go
// internal/logread
type Record struct {
    ExecutionID  string
    StartedAt    string // RFC 3339 UTC; the first 10 characters are the start date
    Status       string
    Command      string // job.command_name
    Target       string // job.target
    Provider     string // job.llm
    Model        string // result.stats.model when non-empty, otherwise job.model
    InputTokens  int64
    CachedTokens int64
    OutputTokens int64
    CostUSD      float64
    ElapsedMS    int64
    ReturnCode   *int // nil when null or absent
    TimedOut     bool
}

func (r Record) Date() string   // StartedAt[:10]
func (r Record) Failed() bool   // Status != "succeeded"

type Result struct {
    Records []Record // kept records in file order
    Skipped int      // skipped lines in the whole file
}

func Read(r io.Reader, warn io.Writer) (Result, error)
func ReadFile(path string, warn io.Writer) (Result, error)
```

```go
// internal/cli
type Options struct {
    Report string
    Target string
    HasTarget bool
    Since  string // YYYY-MM-DD or empty
    Until  string // YYYY-MM-DD or empty
    Log    string
}

func Parse(args []string) (Options, error)
func Run(args []string, stdout, stderr io.Writer) int
```

`Run` uses a local `flag.FlagSet` with `flag.ContinueOnError` or an equivalent hand parser, never
the global flag set, and is tested with `bytes.Buffer`. Report output is buffered and written to
standard output only after the log has been read successfully, so a failing command leaves
standard output empty.

## Execution Flow

parse arguments, open the log, ingest records applying the skip rules, apply the Target and date
filters, compute one report, write CSV. The `skipped` count is computed over the whole file before
any filter applies.

## Shared Output Rules

CSV with one header row and `\n` line endings; minutes written as tenths with
`t = (ms + 3000) / 6000` in integer division and formatted `t/10`, `.`, `t%10`; cost written with
`%.4f`; text ordered by byte value; a `TOTAL` row of zeros when no record matches, where the report
has one.

## Exit Codes

| Code | Meaning |
|---|---|
| 0 | Success, including skipped lines or no matching records |
| 1 | The log cannot be opened or read |
| 2 | Usage error |

Diagnostics go to standard error and are never compared.

## Persistence Ownership

| Boundary | Owner module |
|---|---|
| Persistence, configuration, file stores, external services | None; the application has no database and no external service. The log is read-only input. |

## Technical Decisions

- Records that fail any rule never register their `execution_id`, so a later valid record with that
  id is kept.
- `--help` is an unknown option and exits 2 with the usage line on standard error.
- A repeated option keeps its last occurrence.
- A directory given as the log path is an unreadable log and exits 1.
- A numeric field with a JSON type other than number or `null` counts as zero.

## Programmatic Acceptance

=== AC foundation-builds ===
Intent: The module builds the ddkstat executable from ./cmd/ddkstat with the Go toolchain.
Suite: scoped
Requires: executable=go; scope=test

import os
import subprocess
import sys
import tempfile

with tempfile.TemporaryDirectory() as workdir:
    binary = os.path.join(workdir, "ddkstat")
    result = subprocess.run(
        ["go", "build", "-o", binary, "./cmd/ddkstat"],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    assert os.path.isfile(binary)
=== END AC foundation-builds ===

=== AC foundation-vet ===
Intent: go vet passes across the whole module.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["go", "vet", "./..."], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC foundation-vet ===

=== AC foundation-module-stdlib-only ===
Intent: go.mod declares Go 1.22 or newer and requires no third-party module.
Suite: scoped

import re
from pathlib import Path

text = Path("go.mod").read_text(encoding="utf-8")
match = re.search(r"^go (\d+)\.(\d+)", text, re.M)
assert match is not None
assert (int(match.group(1)), int(match.group(2))) >= (1, 22)
for line in text.splitlines():
    stripped = line.strip()
    assert not stripped.startswith("require"), stripped
    assert not stripped.startswith("replace"), stripped
=== END AC foundation-module-stdlib-only ===

=== AC foundation-layout ===
Intent: The main package and the three internal packages exist, and only main calls os.Exit.
Suite: scoped

from pathlib import Path

assert Path("cmd/ddkstat/main.go").is_file()
for package in ("cli", "logread", "report"):
    directory = Path("internal") / package
    assert directory.is_dir(), package
    assert list(directory.glob("*.go")), package

for source in Path("internal").rglob("*.go"):
    assert "os.Exit(" not in source.read_text(encoding="utf-8"), str(source)
=== END AC foundation-layout ===

=== AC foundation-grader-independence ===
Intent: Application code never reads sources/run.sh.
Suite: scoped

from pathlib import Path

for root in ("cmd", "internal"):
    for source in Path(root).rglob("*.go"):
        if source.name.endswith("_test.go"):
            continue
        assert "run.sh" not in source.read_text(encoding="utf-8"), str(source)
=== END AC foundation-grader-independence ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module and no database.
- Never modify, rewrite, or substitute any file under `sources/`.
- Application code never reads `sources/run.sh` and never embeds a report row, count, or total.
- Never write to the log being read.
- Only `cmd/ddkstat` calls `os.Exit`.
