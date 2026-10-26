# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 13838
- execution id: 20261026.153331.491Z-c5c8a0e5

## Stories built
- Stand up the Go module, directory layout and test runner every other story builds into. (foundation) [story]

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 1037)
- implements: ARCHITECTURE.md (SP 2170)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/ddkstat/layout_test.go
- cmd/ddkstat/main.go
- go.mod

## Pre-build acceptance observation
- RED: arch-go-mod (ARCHITECTURE.md)
  intent: The root go.mod declares module ddkstat, Go 1.22 or newer, and no third-party requirement.
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
      File "arch-go-mod.py", line 3, in <module>
        text = Path("go.mod").read_text()
               ^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: arch-layout (ARCHITECTURE.md)
  intent: The directories every later story builds into exist.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      d = 'cmd/ddkstat'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-layout.py", line 4, in <module>
        assert Path(d).is_dir(), d
               ^^^^^^^^^^^^^^^^
    AssertionError: cmd/ddkstat
- GREEN (prepassed): arch-no-vendor (ARCHITECTURE.md)
  intent: The module is standard-library only and has no go.sum dependency entries.
  return code: 0

## Post-build programmatic acceptance
- PASS: arch-go-mod (ARCHITECTURE.md)
  intent: The root go.mod declares module ddkstat, Go 1.22 or newer, and no third-party requirement.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-layout (ARCHITECTURE.md)
  intent: The directories every later story builds into exist.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-no-vendor (ARCHITECTURE.md)
  intent: The module is standard-library only and has no go.sum dependency entries.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
The block is the foundation: `go.mod`, the `cmd/ddkstat`, `internal/` and `bin/` directories, and a test runner. `bin/test.sh` is what COMPASS's verification protocol calls, and ARCHITECTURE fixes its contents as `go vet ./...` then `go test ./...`. The `arch-*` acceptance checks cover the go.mod contents, the directories, and the absence of go.sum and vendor. I'll create those files now.The runner passes, but with no test files it proves nothing about the layout. I'll add a Go test that reads the layout back from disk, so `bin/test.sh` checks the same things the acceptance assertions do.The go-version check I just wrote is wrong: the prefix test would accept `1.20` and `1.21`. I'll replace it with a numeric comparison.The file's current contents match what I intended (the `fmt` import came from my own edit). Final check that the module is clean and the runner still passes:RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/ddkstat/main.go
- cmd/ddkstat/layout_test.go
- bin/test.sh

SUMMARY:
The foundation is in place: a root Go module, the `cmd/ddkstat`, `internal/` and `bin/` directories, and a test runner.

- `go.mod` declares `module ddkstat` and `go 1.22`, with no `require` lines. There is no `go.sum` and no `vendor/`.
- `cmd/ddkstat/main.go` is an empty `package main` stub. It is the only package that may call `os.Exit`, and the reporting pipeline will replace it in later blocks.
- `internal/` is an empty directory for the packages later blocks add. Git does not track empty directories, so it will not appear in a commit until a later block adds a file to it.
- `bin/test.sh` is a POSIX sh runner that runs `go vet ./...` then `go test ./...`, as ARCHITECTURE specifies. It takes no tee log, since that would create runtime files in the build directory.
- `cmd/ddkstat/layout_test.go` reads the module and layout back from disk and checks the same three things as the `arch-*` acceptance assertions: the module name and Go version, the directories, and the absence of vendor and go.sum entries.

Checks run: `gofmt -l .` printed nothing, `go vet ./...` passed, and `sh bin/test.sh` exited 0 with all three layout tests passing. I did not run `sources/run.sh`, because it is the terminal story and needs the full `ddkstat` implementation. I did not commit, per the build instructions.

Choices to review:
- I placed the layout test in `cmd/ddkstat` because the root has no Go package and `internal/` should hold only application code. A later block can move it if that is wrong.
- I used POSIX sh for `bin/test.sh` to follow ARCHITECTURE, which overrides the bash template in the stack.
