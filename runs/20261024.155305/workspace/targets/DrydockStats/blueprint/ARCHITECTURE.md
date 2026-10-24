# ARCHITECTURE: DrydockStats

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Module layout, package boundaries, and the read-only CLI data flow for the DrydockStats reporter. |
| Depends On  | - |
| Provides    | - |
| Consumes    | - |

## Module Layout

- `go.mod` at the application root: module `drydockstats`, `go 1.22`, no `require` lines.
- `cmd/ddkstat/main.go`: the only `main` package. It calls `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))` and is the only caller of `os.Exit`.
- `internal/cli`: argument parsing, option validation, report dispatch, and exit-code mapping. Exposes `Run(args []string, stdout, stderr io.Writer) int`.
- `internal/logrecord`: JSON Lines decoding, required-field validation, duplicate detection, model attribution, and numeric defaults.

## Technology Stack

- Go 1.22 or newer, standard library only, per TECHNOLOGY_STACK.md.
- No database, no persistence layer, and no `DATABASE.md`.

## Module Ownership

| Boundary | Owner | Allowed access |
|---|---|---|
| Log file reading | `internal/logrecord` | `os.Open`, `bufio.Scanner` |
| Argument and exit handling | `internal/cli` | `flag.FlagSet`, `time.Parse` |
| Process exit | `cmd/ddkstat` | `os.Exit` |

## Technical Decisions

- The log is read-only input. The tool never writes to it.
- Each report is built in memory and written to stdout only after the whole log is read successfully.

## Programmatic Acceptance

=== AC architecture-module-root ===
Intent: The module root declares the module path drydockstats and the Go 1.22 directive with no requires.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import sys
from pathlib import Path

go_mod = Path("go.mod").read_text()
assert "module drydockstats" in go_mod
assert "go 1.22" in go_mod
assert "require" not in go_mod
=== END AC architecture-module-root ===

## User Acceptance

- None.

## Guardrails

- The log file is never written.
- No third-party modules are declared.
