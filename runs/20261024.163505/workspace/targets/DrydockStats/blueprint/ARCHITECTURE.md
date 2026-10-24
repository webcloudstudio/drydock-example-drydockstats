# Architecture: DrydockStats

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Module layout and boundaries for the ddkstat command-line reporter. |
| Depends On  | COMPASS.md |
| Provides    | ddkstat executable, internal/cli Run dispatcher |
| Consumes    | — |

## Technology Stack

- Go 1.22 or newer, standard library only (`go.mod` declares `go 1.22`, no `require` lines).
- Shell: POSIX `sh` for `bin/test.sh`. The supplied grader `sources/run.sh` is staged and never modified.

## Module Layout

```
go.mod                 module ddkstat, go 1.22
cmd/ddkstat/main.go    process startup only; calls cli.Run(os.Args[1:], os.Stdout, os.Stderr) and os.Exit on the returned code
internal/cli/          argument parsing, usage errors, read-failure reporting, dispatch to reports
internal/log/          JSON Lines reading, record parsing, skip rules, model attribution, numeric defaults
internal/filter/       target and inclusive date-range predicates
internal/report/       aggregation and the six CSV reports
internal/format/       minutes (tenths), cost (%.4f), CSV writer
bin/                   launchers; holds test.sh
```

## Module Ownership

| Boundary | Owner | May access |
|---|---|---|
| Process exit | `cmd/ddkstat` | `os.Exit` only |
| Argument parsing and usage errors | `internal/cli` | `flag.FlagSet` (ContinueOnError) |
| Log file I/O | `internal/log` | `os.Open`, `bufio` reader |
| Output formatting | `internal/format` | `encoding/csv`, `fmt` |

## Interfaces

- `cli.Run(args []string, stdout, stderr io.Writer) int` returns the exit code: 0 success, 1 operational read failure, 2 usage error.
- Only `cmd/ddkstat` calls `os.Exit`. No package outside `cmd/ddkstat` writes to the process's real standard streams.

## Technical Decisions

- The executable is built with `go build ./cmd/ddkstat`. No third-party modules, no database.
- Diagnostics go to stderr; stdout carries only the CSV report.

## Programmatic Acceptance

=== AC arch-go-module ===
Intent: The application root is a Go module named ddkstat that declares go 1.22 and has no require directives.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import pathlib

root = pathlib.Path(".")
text = (root / "go.mod").read_text()
assert "module ddkstat" in text.splitlines()[0] or any(
    line.strip() == "module ddkstat" for line in text.splitlines()
)
assert any(line.strip() == "go 1.22" or line.strip().startswith("go 1.2") for line in text.splitlines())
assert "require" not in text
=== END AC arch-go-module ===

=== AC arch-main-builds ===
Intent: The ddkstat executable builds from the application root with go build using only the standard library.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
result = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC arch-main-builds ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party modules.
- Do not read `sources/run.sh` from application code, and do not hardcode any report value.
- Never write to the log being read.
