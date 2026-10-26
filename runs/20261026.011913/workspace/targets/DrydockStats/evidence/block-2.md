# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 10080
- execution id: 20261026.013956.190Z-7a197839

## Stories built
- Parse the command line into a report name, options, and one log path. (cli-parse) [story]

## Acceptance tooling authorization
- FEATURE-CommandLine.md#cli-parse-every-report: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CommandLine.md#cli-parse-options: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CommandLine.md#cli-parse-unit-tests: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-CommandLine.md (SP 1974)
- context: ARCHITECTURE.md (SP 3177)
- stack: common_compact.md (SP 1653)

## Build directory changes
- internal/cli/cli.go
- internal/cli/cli_test.go

## Pre-build acceptance observation
- GREEN (prepassed): cli-parse-every-report (FEATURE-CommandLine.md)
  intent: Every one of the six report names with a log path is accepted as a valid command line, so none exits with the usage-error code.
  return code: 0
  stdout:
    summary 1 ddkstat: report "summary" is not implemented yet
    
    commands 1 ddkstat: report "commands" is not implemented yet
    
    targets 1 ddkstat: report "targets" is not implemented yet
    
    models 1 ddkstat: report "models" is not implemented yet
    
    daily 1 ddkstat: report "daily" is not implemented yet
    
    failures 1 ddkstat: report "failures" is not implemented yet
- GREEN (prepassed): cli-parse-options (FEATURE-CommandLine.md)
  intent: The three options between the report and the log path, in either order, are accepted and do not exit with the usage-error code.
  return code: 0
  stdout:
    ['--target', 'CommonMark', '--since', '2026-10-01', '--until', '2026-10-31'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--until', '2026-10-31', '--since', '2026-10-01', '--target', 'CommonMark'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--since', '2026-10-01'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--until', '2026-10-31'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--target', 'CommonMark'] 1 ddkstat: report "summary" is not implemented yet
- UNVERIFIED: cli-parse-unit-tests (FEATURE-CommandLine.md)
  intent: The parser unit tests required by this feature exist and pass.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestParseAcceptsEveryReport not-run
    {"Time":"2026-10-25T21:39:49.645635588-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:39:49.652692261-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T21:39:49.652733507-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:39:49.653934009-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T21:39:49.653968606-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:39:49.645635588-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:39:49.652692261-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:39:49.652733507-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
      name = 'TestParseAcceptsEveryReport'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-parse-unit-tests.py", line 13, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestParseAcceptsEveryReport', 'not-run', '{"Time":"2026-10-25T21:39:49.645635588-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:39:49.652692261-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:39:49.652733507-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-25T21:39:49.653934009-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T21:39:49.653968606-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')

## Post-build programmatic acceptance
- PASS: cli-parse-every-report (FEATURE-CommandLine.md)
  intent: Every one of the six report names with a log path is accepted as a valid command line, so none exits with the usage-error code.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    summary 1 ddkstat: report "summary" is not implemented yet
    
    commands 1 ddkstat: report "commands" is not implemented yet
    
    targets 1 ddkstat: report "targets" is not implemented yet
    
    models 1 ddkstat: report "models" is not implemented yet
    
    daily 1 ddkstat: report "daily" is not implemented yet
    
    failures 1 ddkstat: report "failures" is not implemented yet
- PASS: cli-parse-options (FEATURE-CommandLine.md)
  intent: The three options between the report and the log path, in either order, are accepted and do not exit with the usage-error code.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ['--target', 'CommonMark', '--since', '2026-10-01', '--until', '2026-10-31'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--until', '2026-10-31', '--since', '2026-10-01', '--target', 'CommonMark'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--since', '2026-10-01'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--until', '2026-10-31'] 1 ddkstat: report "summary" is not implemented yet
    
    ['--target', 'CommonMark'] 1 ddkstat: report "summary" is not implemented yet
- PASS: cli-parse-unit-tests (FEATURE-CommandLine.md)
  intent: The parser unit tests required by this feature exist and pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestParseAcceptsEveryReport pass
    {"Time":"2026-10-25T21:40:57.820980616-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:40:57.828289303-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport"}
    {"Time":"2026-10-25T21:40:57.828382498-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport","Output":"=== RUN   TestParseAcceptsEveryReport\n"}
    {"Time":"2026-10-25T21:40:57.828395489-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/summary"}
    {"Time":"2026-10-25T21:40:57.828398998-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/summary","Output":"=== RUN   TestParseAcceptsEveryReport/summary\n"}
    {"Time":"2026-10-25T21:40:57.828407429-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/summary","Output":"--- PASS: TestParseAcceptsEveryReport/summary (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828411478-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/summary","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.828419807-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/commands"}
    {"Time":"2026-10-25T21:40:57.828422334-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/commands","Output":"=== RUN   TestParseAcceptsEveryReport/commands\n"}
    {"Time":"2026-10-25T21:40:57.828425523-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/commands","Output":"--- PASS: TestParseAcceptsEveryReport/commands (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828428253-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/commands","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.828430681-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/targets"}
    {"Time":"2026-10-25T21:40:57.828432714-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/targets","Output":"=== RUN   TestParseAcceptsEveryReport/targets\n"}
    {"Time":"2026-10-25T21:40:57.828435454-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/targets","Output":"--- PASS: TestParseAcceptsEveryReport/targets (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828437877-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/targets","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.828440208-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/models"}
    {"Time":"2026-10-25T21:40:57.828443381-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/models","Output":"=== RUN   TestParseAcceptsEveryReport/models\n"}
    {"Time":"2026-10-25T21:40:57.828446096-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/models","Output":"--- PASS: TestParseAcceptsEveryReport/models (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828448562-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/models","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.82845083-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/daily"}
    {"Time":"2026-10-25T21:40:57.828492194-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/daily","Output":"=== RUN   TestParseAcceptsEveryReport/daily\n"}
    {"Time":"2026-10-25T21:40:57.828501751-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/daily","Output":"--- PASS: TestParseAcceptsEveryReport/daily (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828505808-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/daily","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.828508944-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/failures"}
    {"Time":"2026-10-25T21:40:57.828511255-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/failures","Output":"=== RUN   TestParseAcceptsEveryReport/failures\n"}
    {"Time":"2026-10-25T21:40:57.828514235-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/failures","Output":"--- PASS: TestParseAcceptsEveryReport/failures (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828516928-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport/failures","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.828519724-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport","Output":"--- PASS: TestParseAcceptsEveryReport (0.00s)\n"}
    {"Time":"2026-10-25T21:40:57.828522491-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseAcceptsEveryReport","Elapsed":0}
    {"Time":"2026-10-25T21:40:57.828554154-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:40:57.829195165-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s\n"}
    {"Time":"2026-10-25T21:40:57.829252192-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
    
    TestParseOptionsFollowReport pass
    {"Time":"2026-10-25T21:40:58.426829015-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:40:58.43380693-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseOptionsFollowReport"}
    {"Time":"2026-10-25T21:40:58.433859581-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseOptionsFollowReport","Output":"=== RUN   TestParseOptionsFollowReport\n"}
    {"Time":"2026-10-25T21:40:58.433869597-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseOptionsFollowReport","Output":"--- PASS: TestParseOptionsFollowReport (0.00s)\n"}
    {"Time":"2026-10-25T21:40:58.433872471-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseOptionsFollowReport","Elapsed":0}
    {"Time":"2026-10-25T21:40:58.433875922-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:40:58.434785236-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s\n"}
    {"Time":"2026-10-25T21:40:58.434810175-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
    
    TestParseLogPathIsLast pass
    {"Time":"2026-10-25T21:40:59.014581496-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:40:59.021743759-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseLogPathIsLast"}
    {"Time":"2026-10-25T21:40:59.021784737-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseLogPathIsLast","Output":"=== RUN   TestParseLogPathIsLast\n"}
    {"Time":"2026-10-25T21:40:59.021824254-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseLogPathIsLast","Output":"--- PASS: TestParseLogPathIsLast (0.00s)\n"}
    {"Time":"2026-10-25T21:40:59.021828644-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseLogPathIsLast","Elapsed":0}
    {"Time":"2026-10-25T21:40:59.021831964-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:40:59.022299937-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T21:40:59.022334785-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
    
    TestParseRepeatedOptionLastWins pass
    {"Time":"2026-10-25T21:40:59.61056276-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:40:59.617908105-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseRepeatedOptionLastWins"}
    {"Time":"2026-10-25T21:40:59.617958211-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseRepeatedOptionLastWins","Output":"=== RUN   TestParseRepeatedOptionLastWins\n"}
    {"Time":"2026-10-25T21:40:59.617967276-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseRepeatedOptionLastWins","Output":"--- PASS: TestParseRepeatedOptionLastWins (0.00s)\n"}
    {"Time":"2026-10-25T21:40:59.617969387-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseRepeatedOptionLastWins","Elapsed":0}
    {"Time":"2026-10-25T21:40:59.617972804-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:40:59.618599075-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s\n"}
    {"Time":"2026-10-25T21:40:59.618626494-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
Some code already exists from the earlier block. I'll read it before changing anything.The scaffold from the earlier block already has `Run`, the interspersed-argument parser, and some tests. This step needs the typed `Options` and `Parse` that the spec names, so I'll build those on top of the existing parser and leave the rest alone.Next, the four unit tests the spec names, plus negative-path tests for unknown reports and positional counts. Each reads the parsed `Options` back and compares fields.Now I'll run the project's own test launcher, gofmt, and build the executable. Then I'll check the spec's CLI behavior directly against the binary.Build and test suite pass. I'll confirm the four spec-named tests ran rather than being skipped or missing.RESULT: SUCCESS

FILES CHANGED:
- internal/cli/cli.go
- internal/cli/cli_test.go

SUMMARY:
- **`internal/cli/cli.go`**: Added the typed `Options` struct from ARCHITECTURE.md and `Parse(args []string) (Options, error)`. `Parse` reuses the existing interspersed `flag.FlagSet` parser (`ContinueOnError`, output discarded), so options can appear in any order between the report and the log. It requires exactly two positionals, validates the report name against the six reports, and sets `HasTarget` via `fs.Visit` so an explicit `--target ""` still counts as given. `Run` now calls `Parse`, so a usage error returns 2 with the usage line on stderr and stdout empty. A valid command line still returns 1 with the "not implemented yet" message, since the report logic belongs to later blocks.
- **`internal/cli/cli_test.go`**: Added the four spec-named tests (`TestParseAcceptsEveryReport`, `TestParseOptionsFollowReport`, `TestParseLogPathIsLast`, `TestParseRepeatedOptionLastWins`), which read back the `Options` struct and compare fields. I also added `TestParseHasTargetFalseWhenAbsent`, `TestParseRejectsUnknownReport`, and `TestParseRejectsWrongPositionalCount`.

Checks run:
- `sh bin/test.sh` (go vet and go test ./...): exit 0. The four named tests ran and passed.
- `gofmt -l`: clean.
- `go build -o ... ./cmd/ddkstat`: succeeds with the standard library only.
- Binary behavior: all six reports and the option orderings from `cli-parse-options` return exit 1, not 2. `frobnicate` and `--target` with no value return 2.
- `sh sources/run.sh`: exits 1 as expected. Its checks cover aggregation, filters, and usage errors that belong to later blocks.

Choices I made:
- I kept the existing flag-based parser instead of rewriting it as a hand parser.
- Parse rejects unknown report names with an error. The spec says the parser must reject argument lists it cannot accept, and the report name is part of that.
- Date values are not validated in `Parse`. The spec puts date validation in a separate feature.

BLOCKERS:
- None for this block. I did not run the `drydock_go` helpers (`cli_runner`, `run_test`) because they are not on disk here. I reproduced their checks directly with `go test -run` and by running the built binary.
