# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 14761
- execution id: 20261026.013825.798Z-fb370839

## Stories built
- Stand up the Go module, the ddkstat main package, and the internal package layout. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#foundation-builds: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#foundation-vet: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 1037)
- implements: ARCHITECTURE.md (SP 3177)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/ddkstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/logread/doc.go
- internal/report/doc.go

## Pre-build acceptance observation
- RED: foundation-builds (ARCHITECTURE.md)
  intent: The module builds the ddkstat executable from ./cmd/ddkstat with the Go toolchain.
  return code: 1
  stderr:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats
    	to create a module there, run:
    	cd ../../../.. && go mod init
    
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'build', '-o', '/tmp/tmp0q2wddzv/ddkstat', './cmd/ddkstat'], returncode=1, stdout='', stderr='go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats\n\tto create a module there, run:\n\tcd ../../../.. && go mod init\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "foundation-builds.py", line 15, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: foundation-vet (ARCHITECTURE.md)
  intent: go vet passes across the whole module.
  return code: 1
  stderr:
    pattern ./...: directory prefix . does not contain main module or its selected dependencies
    
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'vet', './...'], returncode=1, stdout='', stderr='pattern ./...: directory prefix . does not contain main module or its selected dependencies\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "foundation-vet.py", line 7, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: foundation-module-stdlib-only (ARCHITECTURE.md)
  intent: go.mod declares Go 1.22 or newer and requires no third-party module.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('go.mod')
      mode = 'r'
      buffering = -1
      encoding = 'utf-8'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "foundation-module-stdlib-only.py", line 4, in <module>
        text = Path("go.mod").read_text(encoding="utf-8")
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: foundation-layout (ARCHITECTURE.md)
  intent: The main package and the three internal packages exist, and only main calls os.Exit.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "foundation-layout.py", line 3, in <module>
        assert Path("cmd/ddkstat/main.go").is_file()
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): foundation-grader-independence (ARCHITECTURE.md)
  intent: Application code never reads sources/run.sh.
  return code: 0

## Post-build programmatic acceptance
- PASS: foundation-builds (ARCHITECTURE.md)
  intent: The module builds the ddkstat executable from ./cmd/ddkstat with the Go toolchain.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: foundation-vet (ARCHITECTURE.md)
  intent: go vet passes across the whole module.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: foundation-module-stdlib-only (ARCHITECTURE.md)
  intent: go.mod declares Go 1.22 or newer and requires no third-party module.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: foundation-layout (ARCHITECTURE.md)
  intent: The main package and the three internal packages exist, and only main calls os.Exit.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: foundation-grader-independence (ARCHITECTURE.md)
  intent: Application code never reads sources/run.sh.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
Go 1.24.6 meets the 1.22 floor. The directory is empty apart from `sources/`. I'll read the grader and the sample log to see what the CLI must produce, though application code must not read `run.sh`.The foundation step needs `go.mod`, `cmd/ddkstat/main.go`, and the `cli`, `logread`, and `report` packages. The acceptance checks cover the build, `go vet`, the `go` directive, layout, `os.Exit` placement, and that application code never names `run.sh`. Argument parsing and exit codes are in scope; the readers and report formats belong to later blocks. I'll write the module and the skeleton now.Next, the tests. They check return codes and empty stdout for usage paths, as the spec requires, and avoid matching stderr wording.The "options after positionals" case exposes a real problem: Go's `flag` package stops at the first positional, so `summary log --target X` would be a usage error. Drydock's CLI should accept options in either position, so I'll parse interspersed arguments, honoring `--`.The `exitOK` constant is unused, so I'll drop it. I'll also make the accept-path tests assert the actual return value and add a `--` case.Now I'll run the formatter, the project's own test launcher, and the build. `sources/run.sh` is the terminal story and runs last, so I won't run it here.Everything passes. I'll remove the temporary binary and confirm the file list from git's point of view, without staging or committing.RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/ddkstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/logread/doc.go
- internal/report/doc.go
- bin/test.sh

SUMMARY:
The foundation is in place: a Go module with the `ddkstat` main package and the `cli`, `logread`, and `report` packages under `internal/`.

- **Module:** `go.mod` declares `module ddkstat` and `go 1.22`, with no `require` or `replace` lines. The toolchain here is Go 1.24.6, which meets the floor.
- **Entry point:** `cmd/ddkstat/main.go` is the only caller of `os.Exit`. It delegates to `cli.Run(args, stdout, stderr) int`.
- **CLI skeleton:** `internal/cli` parses arguments with a local `flag.FlagSet` in `ContinueOnError` mode. Options are accepted before, between, or after the positionals, and a bare `--` ends option parsing. A usage error (wrong positional count, unknown option, `-h`, or a flag missing its value) writes the usage line to stderr and returns 2. A valid two-positional call returns 1 with "report … is not implemented yet" on stderr, because the reports come in later blocks.
- **Packages:** `logread` and `report` contain package docs only; their parsing, aggregation, and formatting are left for later blocks.
- **Tests:** `internal/cli/cli_test.go` checks exit codes and empty stdout on the usage paths, and checks that options parse in any position. Stderr is only checked for being non-empty, never for wording.
- **Launcher:** `bin/test.sh` runs `go vet ./...` then `go test ./...`.

Checks run:
- `gofmt -l .` prints nothing.
- `sh bin/test.sh` exits 0.
- `go build ./cmd/ddkstat` succeeds.
- The foundation acceptance assertions hold by inspection: `os.Exit(` appears only in `cmd/ddkstat/main.go`, every `internal` package has `.go` files, and no non-test Go file under `cmd` or `internal` mentions `run.sh`. I checked these with grep and did not run the Python acceptance harness.
- I did not run `sources/run.sh`. It is the terminal story and runs last; it will fail until the reports are implemented.

Choices I made that you may want to redirect:
- **Module path** is `ddkstat`, since the manifest names none. Changing it means updating the import in `main.go` and in `cli.go`.
- **Not-implemented exit code** is 1, which the architecture reserves for "log cannot be opened or read." That is a placeholder and will change when the reports land.
- **Declared options** `--target`, `--since`, and `--until` are registered so they parse, but their values are not used yet.

No git add, commit, or branch operations were run; the four new top-level paths are untracked.
