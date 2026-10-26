# FEATURE: Build Executable

| Field       | Value |
|-------------|-------|
| Version     | 20261026 V1 |
| Description | Wires the ddkstat main package and keeps the module standard-library only. |
| Depends On  | ARCHITECTURE.md, FEATURE-CliExit.md |
| Provides    | cmd/ddkstat |
| Consumes    | cli.Run |

## Purpose

The Go module root builds the `ddkstat` executable with `go build ./cmd/ddkstat`, using only the standard library.

## Behavior

- `cmd/ddkstat/main.go` is `package main` and contains only `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`.
- Only this file calls `os.Exit`.
- `go.mod` declares a module path and a Go version of 1.22 or newer, and has no `require` lines.
- `go build ./cmd/ddkstat` from the module root exits 0.
- With no arguments the executable exits 2 with empty standard output.

## Programmatic Acceptance

=== AC build-exe-builds ===
Intent: go build of ./cmd/ddkstat from the module root succeeds.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

out_dir = tempfile.mkdtemp()
binary = os.path.join(out_dir, "ddkstat-build-check")
result = subprocess.run(
    ["go", "build", "-o", binary, "./cmd/ddkstat"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert os.path.isfile(binary)
=== END AC build-exe-builds ===

=== AC build-exe-stdlib-only ===
Intent: The module declares no third-party requirements.

from pathlib import Path

go_mod = Path("go.mod")
assert go_mod.is_file()
lines = [line.strip() for line in go_mod.read_text(encoding="utf-8").splitlines()]
print(lines)
assert any(line.startswith("module ") for line in lines)
assert not any(line.startswith("require") for line in lines)
assert not Path("go.sum").exists() or Path("go.sum").stat().st_size == 0
assert Path("cmd/ddkstat/main.go").is_file()
=== END AC build-exe-stdlib-only ===

=== AC build-exe-usage-exit ===
Intent: The built executable exits 2 and writes nothing to standard output when run with no arguments or an unknown report.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner()
print(out)
print(err)
assert code == 2
assert out == ""

code, out, err = cli_runner("frobnicate", "no-such-log.jsonl")
print(out)
print(err)
assert code == 2
assert out == ""
=== END AC build-exe-usage-exit ===

## User Acceptance

- None.

## Guardrails

- No third-party module and no database.
- `main.go` holds no application logic.
- Application code never reads `sources/run.sh`.
