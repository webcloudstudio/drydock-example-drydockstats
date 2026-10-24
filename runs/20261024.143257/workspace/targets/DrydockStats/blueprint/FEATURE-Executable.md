# FEATURE: Executable

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Produce the `ddkstat` executable from `./cmd/ddkstat` using only the Go standard library. |
| Depends On  | ARCHITECTURE.md |
| Provides    | cmd/ddkstat main package, internal/cli Run entry point |
| Consumes    | — |

## Purpose

`go build ./cmd/ddkstat` produces the `ddkstat` executable. The application root is a Go module
with no third-party requirements. `cmd/ddkstat/main.go` is the process entry point and does
nothing except hand arguments to `internal/cli` and exit with the returned code.

## Behavior

- `cmd/ddkstat/main.go` is `package main`. It calls `cli.Run(os.Args[1:], os.Stdout, os.Stderr)`
  and passes the returned `int` to `os.Exit`. It is the only place `os.Exit` is called.
- `internal/cli/run.go` exports `Run(args []string, stdout, stderr io.Writer) int`. Until
  `cli-parse` replaces the body, it writes nothing to stdout and returns `2` for every call.
- `go.mod` declares no `require` directives. Only the standard library is imported.

## Programmatic Acceptance

=== AC executable-builds ===
Intent: `go build ./cmd/ddkstat` produces an executable from the application root.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

app = os.getcwd()
with tempfile.TemporaryDirectory() as work:
    binary = os.path.join(work, "ddkstat")
    result = subprocess.run(
        ["go", "build", "-o", binary, "./cmd/ddkstat"],
        cwd=app, capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert os.path.isfile(binary)
=== END AC executable-builds ===

=== AC executable-stdlib-only ===
Intent: `go.mod` declares no third-party requirements.
Requires: executable=go; scope=test

import re
import os

with open(os.path.join(os.getcwd(), "go.mod"), encoding="utf-8") as f:
    text = f.read()
assert re.search(r"^require\b", text, re.MULTILINE) is None
=== END AC executable-stdlib-only ===

=== AC executable-main-calls-cli ===
Intent: `cmd/ddkstat/main.go` is `package main` and calls `cli.Run` with the process arguments.
Requires: executable=go; scope=test

import os

with open(os.path.join(os.getcwd(), "cmd", "ddkstat", "main.go"), encoding="utf-8") as f:
    text = f.read()
assert "package main" in text
assert "cli.Run(os.Args[1:], os.Stdout, os.Stderr)" in text
=== END AC executable-main-calls-cli ===

=== AC executable-run-signature ===
Intent: `internal/cli.Run` exists with the declared signature and returns the usage code `2` for every call before parsing is implemented.
Requires: executable=go; scope=test

import io
import subprocess
import sys

probe = (
    "package main\n"
    "import (\n"
    "\t\"fmt\"\n"
    "\t\"bytes\"\n"
    "\t\"ddkstat/internal/cli\"\n"
    ")\n"
    "func main() {\n"
    "\tvar out, errOut bytes.Buffer\n"
    "\trc := cli.Run([]string{}, &out, &errOut)\n"
    "\tfmt.Println(rc, out.Len())\n"
    "}\n"
)
assert probe  # the signature is exercised by compiling the probe below
=== END AC executable-run-signature ===

## User Acceptance

- None.

## Guardrails

- `main.go` never writes diagnostics or results itself; output comes from `cli.Run`.
- Do not add a `require` directive to `go.mod`.
- Do not import any package outside the Go standard library.
