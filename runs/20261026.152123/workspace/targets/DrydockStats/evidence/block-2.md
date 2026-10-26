# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 11391
- execution id: 20261026.153448.328Z-8705eec6

## Stories built
- Read the JSON Lines log into records, ignoring blank lines and classifying failed calls. (log-lines) [story]
- Parse the report name, options and log path into a validated invocation. (cli-args) [story]
- Provide bin/test.sh running go vet and go test. (test-script) [story]

## Acceptance tooling authorization
- FEATURE-CliArgs.md#cli-args-report-and-log: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-options: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-repeated-option: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-option-before-report: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-unknown-report: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-unknown-option: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-missing-value: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliArgs.md#cli-args-log-path-count: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestScript.md#test-script-passes: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestScript.md#test-script-any-cwd: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-LogLines.md (SP 1192)
- context: ARCHITECTURE.md (SP 2170)
- implements: FEATURE-CliArgs.md (SP 1761)
- implements: FEATURE-TestScript.md (SP 830)
- stack: common_compact.md (SP 1653)

## Build directory changes
- internal/cli/args.go
- internal/cli/args_test.go
- internal/logread/logread.go
- internal/logread/logread_test.go

## Pre-build acceptance observation
- RED: lines-blank (FEATURE-LogLines.md)
  intent: Blank lines are ignored and not counted as records or skips.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\n"}
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\n"}
    {"ImportPath":"./internal/logread","Action":"build-fail"}
    {"Time":"2026-10-26T11:34:28.536178952-04:00","Action":"start","Package":"./internal/logread"}
    {"Time":"2026-10-26T11:34:28.536254423-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\t./internal/logread [setup failed]\n"}
    {"Time":"2026-10-26T11:34:28.536263711-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./inte… (+451 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "lines-blank.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./internal/logread","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:28.536178952-04:00","Action":"start","Package":"./internal/logread"}\n{"Time":"2026-10-26T11:34:28.536254423-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\\t./internal/logread [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:28.536263711-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}\n', '')
- RED: lines-failed-status (FEATURE-LogLines.md)
  intent: A record whose status is not succeeded is classified as failed.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\n"}
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\n"}
    {"ImportPath":"./internal/logread","Action":"build-fail"}
    {"Time":"2026-10-26T11:34:29.169022891-04:00","Action":"start","Package":"./internal/logread"}
    {"Time":"2026-10-26T11:34:29.169077028-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\t./internal/logread [setup failed]\n"}
    {"Time":"2026-10-26T11:34:29.169086437-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./inte… (+451 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "lines-failed-status.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./internal/logread","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:29.169022891-04:00","Action":"start","Package":"./internal/logread"}\n{"Time":"2026-10-26T11:34:29.169077028-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\\t./internal/logread [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:29.169086437-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}\n', '')
- RED: lines-start-date (FEATURE-LogLines.md)
  intent: The record date is the first 10 characters of started_at.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\n"}
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\n"}
    {"ImportPath":"./internal/logread","Action":"build-fail"}
    {"Time":"2026-10-26T11:34:29.833151386-04:00","Action":"start","Package":"./internal/logread"}
    {"Time":"2026-10-26T11:34:29.833217791-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\t./internal/logread [setup failed]\n"}
    {"Time":"2026-10-26T11:34:29.833236101-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./inte… (+451 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "lines-start-date.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./internal/logread","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:29.833151386-04:00","Action":"start","Package":"./internal/logread"}\n{"Time":"2026-10-26T11:34:29.833217791-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\\t./internal/logread [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:29.833236101-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}\n', '')
- RED: lines-started-at-lenient (FEATURE-LogLines.md)
  intent: started_at is accepted as any non-empty string without timestamp validation.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\n"}
    {"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\n"}
    {"ImportPath":"./internal/logread","Action":"build-fail"}
    {"Time":"2026-10-26T11:34:30.486908803-04:00","Action":"start","Package":"./internal/logread"}
    {"Time":"2026-10-26T11:34:30.487008762-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\t./internal/logread [setup failed]\n"}
    {"Time":"2026-10-26T11:34:30.487018166-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./inte… (+451 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "lines-started-at-lenient.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/logread","Action":"build-output","Output":"# ./internal/logread\\n"}\n{"ImportPath":"./internal/logread","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/logread: directory not found\\n"}\n{"ImportPath":"./internal/logread","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:30.486908803-04:00","Action":"start","Package":"./internal/logread"}\n{"Time":"2026-10-26T11:34:30.487008762-04:00","Action":"output","Package":"./internal/logread","Output":"FAIL\\t./internal/logread [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:30.487018166-04:00","Action":"fail","Package":"./internal/logread","Elapsed":0,"FailedBuild":"./internal/logread"}\n', '')
- RED: cli-args-report-and-log (FEATURE-CliArgs.md)
  intent: A report name followed by a log path parses into Options with that report and path.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-report-and-log.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:31.136766952-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:31.136830683-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:31.136840843-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-options (FEATURE-CliArgs.md)
  intent: The --target, --since and --until options populate Options.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-options.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:31.811002025-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:31.811066347-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:31.811075855-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-repeated-option (FEATURE-CliArgs.md)
  intent: A repeated option takes its last value.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-repeated-option.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:32.421110895-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:32.421158872-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:32.421170851-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-option-before-report (FEATURE-CliArgs.md)
  intent: An option placed before the report name is a usage error.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-option-before-report.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:33.023815125-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:33.023892046-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:33.023901435-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-unknown-report (FEATURE-CliArgs.md)
  intent: An unknown report name is a usage error.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-unknown-report.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:33.655402682-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:33.655461123-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:33.655469968-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-unknown-option (FEATURE-CliArgs.md)
  intent: An unknown option is a usage error.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-unknown-option.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:34.285387543-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:34.285440123-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:34.285450268-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-missing-value (FEATURE-CliArgs.md)
  intent: An option with no value is a usage error.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+409 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-missing-value.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:34.923210564-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:34.92327487-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:34.92328366-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- RED: cli-args-log-path-count (FEATURE-CliArgs.md)
  intent: A missing log path and an extra log path are both usage errors.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Actio… (+411 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-args-log-path-count.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/cli","Action":"build-output","Output":"# ./internal/cli\\n"}\n{"ImportPath":"./internal/cli","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/cli: directory not found\\n"}\n{"ImportPath":"./internal/cli","Action":"build-fail"}\n{"Time":"2026-10-26T11:34:35.594541725-04:00","Action":"start","Package":"./internal/cli"}\n{"Time":"2026-10-26T11:34:35.594596215-04:00","Action":"output","Package":"./internal/cli","Output":"FAIL\\t./internal/cli [setup failed]\\n"}\n{"Time":"2026-10-26T11:34:35.594604688-04:00","Action":"fail","Package":"./internal/cli","Elapsed":0,"FailedBuild":"./internal/cli"}\n', '')
- GREEN (prepassed): test-script-syntax (FEATURE-TestScript.md)
  intent: bin/test.sh exists and parses as POSIX sh.
  return code: 0
- GREEN (prepassed): test-script-passes (FEATURE-TestScript.md)
  intent: sh bin/test.sh exits zero when go vet and every automated test pass.
  return code: 0
  stdout:
    ok  	ddkstat/cmd/ddkstat	0.013s
- GREEN (prepassed): test-script-any-cwd (FEATURE-TestScript.md)
  intent: bin/test.sh finds the application root regardless of the caller's working directory.
  return code: 0
  stdout:
    ok  	ddkstat/cmd/ddkstat	(cached)

## Post-build programmatic acceptance
- PASS: lines-blank (FEATURE-LogLines.md)
  intent: Blank lines are ignored and not counted as records or skips.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:36:36.139665013-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:36:36.14613713-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadBlankLines"}
    {"Time":"2026-10-26T11:36:36.146164534-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadBlankLines","Output":"=== RUN   TestReadBlankLines\n"}
    {"Time":"2026-10-26T11:36:36.146186464-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadBlankLines","Output":"--- PASS: TestReadBlankLines (0.00s)\n"}
    {"Time":"2026-10-26T11:36:36.146214603-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadBlankLines","Elapsed":0}
    {"Time":"2026-10-26T11:36:36.146233679-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:36:36.146967549-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:36:36.146988956-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: lines-failed-status (FEATURE-LogLines.md)
  intent: A record whose status is not succeeded is classified as failed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:36:36.92047271-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:36:36.926728525-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadFailedStatus"}
    {"Time":"2026-10-26T11:36:36.92676546-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadFailedStatus","Output":"=== RUN   TestReadFailedStatus\n"}
    {"Time":"2026-10-26T11:36:36.92682018-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadFailedStatus","Output":"--- PASS: TestReadFailedStatus (0.00s)\n"}
    {"Time":"2026-10-26T11:36:36.926842059-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadFailedStatus","Elapsed":0}
    {"Time":"2026-10-26T11:36:36.926845297-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:36:36.927419375-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:36:36.927440481-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: lines-start-date (FEATURE-LogLines.md)
  intent: The record date is the first 10 characters of started_at.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:36:37.875587704-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:36:37.881073966-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadStartDate"}
    {"Time":"2026-10-26T11:36:37.881095626-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadStartDate","Output":"=== RUN   TestReadStartDate\n"}
    {"Time":"2026-10-26T11:36:37.881175598-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadStartDate","Output":"--- PASS: TestReadStartDate (0.00s)\n"}
    {"Time":"2026-10-26T11:36:37.881195667-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadStartDate","Elapsed":0}
    {"Time":"2026-10-26T11:36:37.881199098-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:36:37.881721196-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:36:37.881738779-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: lines-started-at-lenient (FEATURE-LogLines.md)
  intent: started_at is accepted as any non-empty string without timestamp validation.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:36:38.819967001-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:36:38.826177007-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadStartedAtNotValidated"}
    {"Time":"2026-10-26T11:36:38.826219437-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadStartedAtNotValidated","Output":"=== RUN   TestReadStartedAtNotValidated\n"}
    {"Time":"2026-10-26T11:36:38.826291278-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadStartedAtNotValidated","Output":"--- PASS: TestReadStartedAtNotValidated (0.00s)\n"}
    {"Time":"2026-10-26T11:36:38.826303702-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadStartedAtNotValidated","Elapsed":0}
    {"Time":"2026-10-26T11:36:38.826307693-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:36:38.826774595-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:36:38.82679783-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-report-and-log (FEATURE-CliArgs.md)
  intent: A report name followed by a log path parses into Options with that report and path.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-options (FEATURE-CliArgs.md)
  intent: The --target, --since and --until options populate Options.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-repeated-option (FEATURE-CliArgs.md)
  intent: A repeated option takes its last value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-option-before-report (FEATURE-CliArgs.md)
  intent: An option placed before the report name is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-unknown-report (FEATURE-CliArgs.md)
  intent: An unknown report name is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-unknown-option (FEATURE-CliArgs.md)
  intent: An unknown option is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-missing-value (FEATURE-CliArgs.md)
  intent: An option with no value is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-args-log-path-count (FEATURE-CliArgs.md)
  intent: A missing log path and an extra log path are both usage errors.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: test-script-syntax (FEATURE-TestScript.md)
  intent: bin/test.sh exists and parses as POSIX sh.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: test-script-passes (FEATURE-TestScript.md)
  intent: sh bin/test.sh exits zero when go vet and every automated test pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ok  	ddkstat/cmd/ddkstat	0.013s
    ok  	ddkstat/internal/cli	0.006s
    ok  	ddkstat/internal/logread	0.006s
- PASS: test-script-any-cwd (FEATURE-TestScript.md)
  intent: bin/test.sh finds the application root regardless of the caller's working directory.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ok  	ddkstat/cmd/ddkstat	(cached)
    ok  	ddkstat/internal/cli	(cached)
    ok  	ddkstat/internal/logread	(cached)

## Build summary
I'll start by inspecting the build directory and the staged sources.The directory already has `bin/test.sh`, `cmd/ddkstat/main.go`, `layout_test.go`, and `go.mod`. I'll read those before changing anything.The existing `bin/test.sh` already matches the test-script spec (POSIX `sh`, `set -eu`, cd to the root, `go vet` then `go test`), so I'll leave it alone. The work is `internal/logread` and `internal/cli/args`, each with tests. I'll write the reader first.Next, the logread tests. Each one calls `Read` and checks the returned values, not the warning text.Now the argument parser and its tests, following the rules in FEATURE-CliArgs.The spec names a `TestParseLogPathCount` test that I haven't written yet. I'll add it as the boundary test over zero, one, and several log paths.Both packages are written. Next I'll format, vet, and run the project gate, `sh bin/test.sh`.The gate passes. The change on disk is only gofmt moving the `encoding/json` import into sorted order. Next I'll confirm each required test ran and passed, and check the reader against the real sample log with a scratch test that I delete afterward.Both stories and the test script are done. `sh bin/test.sh` exits 0, and all 35 Go tests pass.

- **log-lines:** `logread.Read` parses the sample log into 35 records, with 1 skipped (line 16, truncated JSON) and 2 failed. The sample's 3 skips and 33 records come from duplicate handling, which a later block adds.
- **cli-args:** `cli.ParseArgs` returns `Options` or a `*UsageError`. Option order is fixed, so an option after the log path is a usage error.
- **test-script:** `bin/test.sh` was already in the repo and matched the spec (POSIX `sh`, `set -eu`, `go vet ./...` then `go test ./...`), so I left it unchanged.

I did not run `sources/run.sh`. It is the terminal check for the whole project, and this block is not the last one.

Changes are uncommitted, as the build rules require.

<reusable-compact filename="ARCHITECTURE.md">
Module `ddkstat`, go 1.22+, no `require` lines, no go.sum entries, no vendor/.
Packages: cmd/ddkstat (only os.Exit caller), internal/logread, internal/filter, internal/format, internal/report, internal/cli.
logread.Read(r io.Reader, warn io.Writer) (Result, error); Result{Records []Record, Skipped int}; error only on reader failure.
logread.Record: ExecutionID, StartedAt, Date (first 10 chars), Status, Failed (status != "succeeded"), Command, Target, Provider, Model, InputTokens, CachedTokens, OutputTokens (int64), CostUSD (float64), ElapsedMS (int64), ReturnCode (*int, nil if null/absent), TimedOut (bool).
cli.Run(args []string, stdout, stderr io.Writer) int; exit 0 ok, 1 unreadable log, 2 usage. Report writers: cli.Register(name, fn func(w io.Writer, records []logread.Record, skipped int) error).
Reports: summary, commands, targets, models, daily, failures. Options: --target, --since, --until, then one log path.
Only cmd/ddkstat/main.go calls os.Exit. Output buffered; stdout written only on success. Diagnostics to stderr, never compared.
Never write to the log; never read sources/run.sh; no hardcoded report values.
</reusable-compact>

```text
<blueprint-decisions>
[{"spec":"FEATURE-LogLines.md","severity":"Low","subject":"Malformed JSON lines","decision":"The spec does not say what to do with a non-blank line that is not a JSON object. I skip it, count it in Result.Skipped, and write a diagnostic to the warning writer, so Read never fails on bad input. Is that acceptable, or should the skip rule belong to the later skip/dedupe story?"},{"spec":"FEATURE-LogLines.md","severity":"Low","subject":"Model fields","decision":"Record.Model is taken from result.stats.model (executed) and Record.RequestedModel from job.model. The architecture lists only Model. Is that acceptable, or should model attribution choose the field?"}]
</battle>
```

RESULT: SUCCESS

FILES CHANGED:
- internal/logread/logread.go
- internal/logread/logread_test.go
- internal/cli/args.go
- internal/cli/args_test.go

SUMMARY:
- `internal/logread`: `Read` decodes JSON Lines from an `io.Reader` using `bufio` with no line-length limit. Blank lines are ignored and not counted. Each object maps to `Record`, `Failed` is `status != "succeeded"`, `Date` is the first 10 characters of `started_at`, and `started_at` is not validated. Malformed lines are skipped and counted, with a diagnostic on the warning writer. Tests cover all four named behaviors plus field mapping, nullable return code, malformed and non-object lines, reader errors, and a nil warning writer.
- `internal/cli`: `ParseArgs` returns `Options{Report, Target, Since, Until, LogPath}` or a `*UsageError`. It is a hand-rolled parser with no I/O and no `os.Exit`. Tests cover every name in the spec, including the boundary cases for zero, one, and several log paths.
- `bin/test.sh`: already present and conforming, so unchanged.
- `sh bin/test.sh` exits 0. The reader run on `sources/llm_sample.jsonl` gives 35 records, 1 skipped, and 2 failed. The grader expects 33 records and 3 skipped. The gap is the duplicate handling that a later block adds.

BLOCKERS:
- None.
