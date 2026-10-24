# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 9949
- execution id: 20261024.164604.424Z-baf6e842

## Stories built
- Establish the Go module and project tree. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#arch-go-module: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#arch-main-builds: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 1037)
- implements: ARCHITECTURE.md (SP 1258)

## Build directory changes
- bin/test.sh
- cmd/ddkstat/main.go
- go.mod
- internal/.gitkeep
- internal/cli/cli.go
- internal/cli/cli_test.go

## Pre-build acceptance observation
- RED: arch-go-module (ARCHITECTURE.md)
  intent: The application root is a Go module named ddkstat that declares go 1.22 and has no require directives.
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
      File "arch-go-module.py", line 5, in <module>
        text = (root / "go.mod").read_text()
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: arch-main-builds (ARCHITECTURE.md)
  intent: The ddkstat executable builds from the application root with go build using only the standard library.
  return code: 1
  stdout:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats
    	to create a module there, run:
    	cd ../../../.. && go mod init
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'build', '-o', '/tmp/tmp9u6cer9s/ddkstat', './cmd/ddkstat'], returncode=1, stdout='', stderr='go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats\n\tto create a module there, run:\n\tcd ../../../.. && go mod init\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-main-builds.py", line 9, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: arch-go-module (ARCHITECTURE.md)
  intent: The application root is a Go module named ddkstat that declares go 1.22 and has no require directives.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-main-builds (ARCHITECTURE.md)
  intent: The ddkstat executable builds from the application root with go build using only the standard library.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Repair attempts
- attempt 0 (initial build): failed; 1/2 checks model=haiku; execution 20261024.164537.480Z-76e318fb; reason: programmatic acceptance failed: arch-main-builds
- attempt 1 (repair 1): built; 2/2 checks model=haiku; execution 20261024.164604.424Z-baf6e842

## Build summary
The failing check is `arch-main-builds`, which needs `./cmd/ddkstat` to exist and compile. I'll add the thin `main` that the architecture specifies, plus a minimal `internal/cli.Run` with its usage-error path. The reporting logic belongs to later blocks.RESULT: SUCCESS

FILES CHANGED:
- cmd/ddkstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go

SUMMARY:
- The repair fixes `arch-main-builds`. `cmd/ddkstat/main.go` is a thin `main` that calls `cli.Run(os.Args[1:], os.Stdout, os.Stderr)` and passes the result to `os.Exit`. It is the only place that exits.
- `internal/cli/cli.go` has `Run(args []string, stdout, stderr io.Writer) int` using a `flag.FlagSet` with `ContinueOnError`. It returns 2 for a bad flag, for no arguments (usage on stderr), and for any report name. No report is implemented in this foundation step, so the report-name case is a stub that returns 2. Reporting logic is left for the later blocks.
- `internal/cli/cli_test.go` is a table-driven test. It checks the returned exit code and that stdout is empty for the three usage-error cases.
- Checks run: `gofmt -l .` printed nothing; `go vet ./...` passed; `go build ./cmd/ddkstat` succeeded; `sh bin/test.sh` exited 0 (`internal/cli` ok, `cmd/ddkstat` has no test files).
- I did not run `sources/run.sh`. It is the terminal story's gate and is not part of this block.
- `ARCHITECTURE.md` is still not edited. It sits under the protected `blueprint/` path, and its module layout already matches what was built.

BLOCKERS:
- None.

<reusable-compact filename="ARCHITECTURE.md">
Module: `ddkstat`, `go 1.22`, no `require` lines (standard library only).
Layout: `cmd/ddkstat/main.go` (process startup only; the sole `os.Exit`); `internal/cli` (args, usage errors, dispatch); `internal/log` (JSON Lines read, parse, skip rules, model attribution, numeric defaults); `internal/filter` (target and inclusive date-range predicates); `internal/report` (aggregation and six CSV reports); `internal/format` (minutes in tenths, cost `%.4f`, CSV writer); `bin/` (launchers, holds `test.sh`).
Interface: `cli.Run(args []string, stdout, stderr io.Writer) int` — 0 success, 1 operational read failure, 2 usage error.
Boundaries: only `cmd/ddkstat` calls `os.Exit`; no package outside `cmd/ddkstat` writes to the real stdout/stderr; stdout carries only CSV; diagnostics go to stderr.
Acceptance: `go build -o <out> ./cmd/ddkstat` from the root must succeed; `go.mod` must contain `module ddkstat` and `go 1.22` with no `require`.
Guardrails: no third-party modules; never read `sources/run.sh` from application code; never hardcode report values; never write to the log being read.
</reusable-compact>

<blueprint-decisions>
[{"spec":"ARCHITECTURE.md","severity":"Low","subject":"Stub dispatch in internal/cli for the foundation step","decision":"Run returns usage code 2 for any report name, because no reports exist yet. Later blocks replace this with real dispatch. Should the foundation instead return a distinct code for not-yet-implemented reports?"}]
</blueprint-decisions>
