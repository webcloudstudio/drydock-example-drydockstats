# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 12358
- execution id: 20261024.160355.997Z-746b93d1

## Stories built
- Create the Go module root, the ddkstat main package, and the internal CLI package. (go-foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#architecture-module-root: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 1037)
- implements: ARCHITECTURE.md (SP 769)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/ddkstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go

## Pre-build acceptance observation
- RED: architecture-module-root (ARCHITECTURE.md)
  intent: The module root declares the module path drydockstats and the Go 1.22 directive with no requires.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('go.mod')
      mode = 'r'
      buffering = -1
      encoding = 'locale'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-module-root.py", line 5, in <module>
        go_mod = Path("go.mod").read_text()
                 ^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'

## Post-build programmatic acceptance
- PASS: architecture-module-root (ARCHITECTURE.md)
  intent: The module root declares the module path drydockstats and the Go 1.22 directive with no requires.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
The build directory has only `.git` and `sources/`, so this is a fresh module. Go 1.24.6 meets the 1.22 floor. I'll create the module root, the main package, the CLI package, and a test script with tests.Files are written. Next I'll check formatting, vet, and the test script, and confirm the module-root acceptance assertion.RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/ddkstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- bin/test.sh

SUMMARY:
- `go.mod` declares module `drydockstats` with `go 1.22` and no `require` lines. The module-root acceptance check (`architecture-module-root`) passes.
- `cmd/ddkstat/main.go` is the only `main` package. It calls `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`, so it is the only caller of `os.Exit`.
- `internal/cli/cli.go` exposes `Run(args []string, stdout, stderr io.Writer) int`, returning 0 as the placeholder.
- `internal/cli/cli_test.go` is a table-driven test that checks `Run` returns 0 for no, one, and many arguments. It reads the return value directly.
- `bin/test.sh` runs `go test ./...` from the application root, per the COMPASS verification protocol.
- Verification: `gofmt -l .` prints nothing, `go vet ./...` is clean, `go build ./cmd/ddkstat` succeeds, and `sh bin/test.sh` exits 0 (cli package passes).
- I did not run `sources/run.sh`. COMPASS reserves it for the terminal story, and this is the foundation block. Nothing was committed.

Choices I made beyond the spec:
- I added `bin/test.sh`. COMPASS requires it for every behavioral story, and the spec did not name it.
- The `Run` parameters are unused in the placeholder. Go accepts that, and the signature is the contract the later blocks build on.
