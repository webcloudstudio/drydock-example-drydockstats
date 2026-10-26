# Evidence: Block 6 · Service (block-6)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 5594
- execution id: 20261026.154406.567Z-62543540

## Stories built
- Assemble the Run pipeline and map failures to exit codes 0, 1 and 2. (cli-exit) [story]

## Acceptance tooling authorization
- FEATURE-CliExit.md#cli-exit-usage-code: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliExit.md#cli-exit-usage-stdout-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliExit.md#cli-exit-missing-log-code: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliExit.md#cli-exit-missing-log-stdout-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliExit.md#cli-exit-skipped-lines-ok: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliExit.md#cli-exit-filter-order: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-CliExit.md (SP 1596)
- context: ARCHITECTURE_compact.md (SP 493)

## Build directory changes
- cmd/ddkstat/main.go
- internal/cli/run.go
- internal/cli/run_test.go

## Pre-build acceptance observation
- UNVERIFIED: cli-exit-usage-code (FEATURE-CliExit.md)
  intent: A usage error returns exit code 2.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:43:53.837383525-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:53.843887608-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:53.843925755-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-exit-usage-code.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:43:53.837383525-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:53.843887608-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:53.843925755-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:43:53.844467736-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:43:53.844488267-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}\n', '')
- UNVERIFIED: cli-exit-usage-stdout-empty (FEATURE-CliExit.md)
  intent: A usage error writes nothing to standard output.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:43:54.749694969-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:54.75606815-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:54.756109733-04:00","Action":"output","Package":"ddkstat/internal/… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-exit-usage-stdout-empty.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:43:54.749694969-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:54.75606815-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:54.756109733-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:43:54.756894364-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:43:54.756925244-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}\n', '')
- UNVERIFIED: cli-exit-missing-log-code (FEATURE-CliExit.md)
  intent: An unreadable log returns exit code 1.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:43:55.783163461-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:55.792182712-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:55.792220482-04:00","Action":"output","Package":"ddkstat/internal… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-exit-missing-log-code.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:43:55.783163461-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:55.792182712-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:55.792220482-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:43:55.792952827-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.010s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:43:55.792973459-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.01}\n', '')
- UNVERIFIED: cli-exit-missing-log-stdout-empty (FEATURE-CliExit.md)
  intent: An unreadable log writes nothing to standard output.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:43:56.945248237-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:56.951969998-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:56.952010619-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-exit-missing-log-stdout-empty.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:43:56.945248237-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:56.951969998-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:56.952010619-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:43:56.952799399-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:43:56.952843587-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- UNVERIFIED: cli-exit-skipped-lines-ok (FEATURE-CliExit.md)
  intent: A log containing skipped lines still exits 0.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:43:57.953338698-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:57.962328987-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:57.962378715-04:00","Action":"output","Package":"ddkstat/internal… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-exit-skipped-lines-ok.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:43:57.953338698-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:57.962328987-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:57.962378715-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:43:57.963314609-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.010s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:43:57.963338453-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.01}\n', '')
- UNVERIFIED: cli-exit-filter-order (FEATURE-CliExit.md)
  intent: The registered writer receives records narrowed by the target and date filters and the whole-file skipped count.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:43:59.05211728-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:59.058327764-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:59.058399-04:00","Action":"output","Package":"ddkstat/internal/cli… (+327 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-exit-filter-order.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:43:59.05211728-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:43:59.058327764-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:43:59.058399-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:43:59.058934022-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:43:59.058955696-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}\n', '')

## Post-build programmatic acceptance
- PASS: cli-exit-usage-code (FEATURE-CliExit.md)
  intent: A usage error returns exit code 2.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-exit-usage-stdout-empty (FEATURE-CliExit.md)
  intent: A usage error writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-exit-missing-log-code (FEATURE-CliExit.md)
  intent: An unreadable log returns exit code 1.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-exit-missing-log-stdout-empty (FEATURE-CliExit.md)
  intent: An unreadable log writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-exit-skipped-lines-ok (FEATURE-CliExit.md)
  intent: A log containing skipped lines still exits 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-exit-filter-order (FEATURE-CliExit.md)
  intent: The registered writer receives records narrowed by the target and date filters and the whole-file skipped count.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
The block needs `cli.Run` and `cli.Register`, six named tests, and `main.go` wired to call `Run`. The existing `ParseArgs`, `logread.Read`, and filter functions already provide the pieces. Writing the implementation now.Now the tests. The six named acceptance tests go in one file, plus a few finer cases (writer error, success-path output, unregistered report). They read state back through the writer and the exit code rather than matching message text.Next, `main.go` becomes the only `os.Exit` caller, delegating to `cli.Run`.The cli-exit step is implemented. All six declared acceptance tests pass, and `sh bin/test.sh` exits 0. `sh sources/run.sh` exits 1, but only because the report writers belong to later blocks.

**Changes**
- `internal/cli/run.go` (new): `Run(args, stdout, stderr) int` and `Register(name, fn)`.
  - A usage error writes one line to stderr and returns 2.
  - A log that cannot be opened or read writes one line and returns 1.
  - Otherwise it applies `ByTarget` then `ByDateRange` and calls the registered writer. The writer's output is buffered and flushed to stdout only on success. A writer error returns 1 with stdout untouched.
  - An unregistered report returns 0 with no output.
  - Skipped count is the whole-file count, so filters do not change it.
- `internal/cli/run_test.go` (new): the six declared tests, plus finer tests for skipped count vs. filters, stdout on success, writer error, and unregistered reports. They check exit codes, stdout bytes, and the records and skipped count the writer receives.
- `cmd/ddkstat/main.go`: now `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`. It is still the only `os.Exit` caller. The spec makes `main` the entry point, so I wired it here.

**Checks**
- `gofmt -l`: clean. `go vet ./...`: clean.
- `go test -race ./internal/cli -run TestRun -v`: all 9 `TestRun*` tests pass.
- `sh bin/test.sh`: exit 0.
- `sh sources/run.sh`: 7 passed, 13 failed (of 20). The 7 passing checks are E1 and U1 through U6. The 13 failing checks are A1–A6, B1–B5, and C1–C2. They fail because no report writer is registered yet, so nothing is written to stdout. Those reports are the work of later blocks. The grader's stderr also confirms the expected skip count of 3 for the sample log.

**Choices I made**
- `Register` overwrites an existing writer for the same name rather than panicking. Duplicate registration is therefore silent.
- The usage line is a fixed one-line summary and does not include the `ParseArgs` error text.
- The unregistered-report test uses `daily`, and the writer-error test uses `targets`. Tests share the global registry, so this assumes no test in this package registers `daily`.

RESULT: SUCCESS

FILES CHANGED:
- cmd/ddkstat/main.go
- internal/cli/run.go
- internal/cli/run_test.go

SUMMARY:
Implemented `cli.Run` and `cli.Register` with exit codes 0, 1, and 2. Buffered report output is written to stdout only on writer success. `main` delegates to `Run`. The six declared acceptance tests pass, along with finer tests. `bin/test.sh` exits 0. `sources/run.sh` still fails 13 of 20 checks because report writers from later blocks are not yet registered.

BLOCKERS:
- None for this step. The remaining `sources/run.sh` failures depend on the report stories in later blocks.

<blueprint-decisions>
[{"spec":"FEATURE-CliExit.md","severity":"Low","subject":"Register behavior on duplicate report name","decision":"Register overwrites any writer already registered under the same name, without panicking. Duplicate registration is silent. Is that acceptable, or should it fail loudly on replan?"}]
</blueprint-decisions>
