# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 20274
- execution id: 20261026.014211.743Z-0d70a833

## Stories built
- Read the JSON Lines log line by line, ignoring blank lines. (log-read) [story]
- Validate --since and --until as real YYYY-MM-DD dates. (cli-dates) [story]
- Exit 1 with nothing on stdout when the log cannot be opened or read. (cli-unreadable-log) [story]
- Exit 2 with nothing on stdout for every usage error. (cli-usage-errors) [story]
- Skip malformed lines with a stderr warning and a skipped count. (log-skip-malformed) [story]
- Skip lines repeating an earlier kept execution_id. (log-skip-duplicates) [story]
- Treat absent, null, or wrongly typed numeric fields as zero. (log-absent-numbers) [story]
- Attribute each record to the executed model, falling back to the requested one. (log-model) [story]

## Acceptance tooling authorization
- FEATURE-LogReading.md#log-read-blank-lines: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogReading.md#log-read-long-line: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogReading.md#log-read-file-errors: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogReading.md#log-read-read-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateOptions.md#date-options-accepts-real: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateOptions.md#date-options-rejects-malformed: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateOptions.md#date-options-rejects-impossible: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableLog.md#unreadable-log-missing-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableLog.md#unreadable-log-directory: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableLog.md#unreadable-log-no-partial-output: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-errors-exit-2: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-errors-before-open: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-errors-usage-line: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MalformedLines.md#malformed-not-object: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MalformedLines.md#malformed-invalid-utf8: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MalformedLines.md#malformed-required-fields: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MalformedLines.md#malformed-started-at: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MalformedLines.md#malformed-warning-and-count: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DuplicateIds.md#duplicate-first-kept: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DuplicateIds.md#duplicate-skipped-line-does-not-register: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DuplicateIds.md#duplicate-warning-and-count: executable=go; scope=test; authorization=existing Target environment
- FEATURE-AbsentNumbers.md#absent-numbers-no-stats: executable=go; scope=test; authorization=existing Target environment
- FEATURE-AbsentNumbers.md#absent-numbers-null-cost: executable=go; scope=test; authorization=existing Target environment
- FEATURE-AbsentNumbers.md#absent-numbers-absent-fields: executable=go; scope=test; authorization=existing Target environment
- FEATURE-AbsentNumbers.md#absent-numbers-wrong-type: executable=go; scope=test; authorization=existing Target environment
- FEATURE-AbsentNumbers.md#absent-numbers-returncode: executable=go; scope=test; authorization=existing Target environment
- FEATURE-AbsentNumbers.md#absent-numbers-timed-out: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ModelAttribution.md#model-attribution-executed: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ModelAttribution.md#model-attribution-no-stats: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ModelAttribution.md#model-attribution-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ModelAttribution.md#model-attribution-provider: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-LogReading.md (SP 1545)
- context: ARCHITECTURE.md (SP 3177)
- stack: common_compact.md (SP 1653)
- implements: FEATURE-DateOptions.md (SP 1104)
- implements: FEATURE-UnreadableLog.md (SP 1126)
- implements: FEATURE-UsageErrors.md (SP 1604)
- implements: FEATURE-MalformedLines.md (SP 1824)
- implements: FEATURE-DuplicateIds.md (SP 1435)
- implements: FEATURE-AbsentNumbers.md (SP 1767)
- implements: FEATURE-ModelAttribution.md (SP 1202)

## Build directory changes
- internal/cli/cli.go
- internal/cli/dates_test.go
- internal/cli/run_test.go
- internal/logread/absent_test.go
- internal/logread/duplicate_test.go
- internal/logread/helpers_test.go
- internal/logread/model_test.go
- internal/logread/read.go
- internal/logread/read_test.go
- internal/logread/record.go
- internal/logread/skip_test.go

## Pre-build acceptance observation
- UNVERIFIED: log-read-blank-lines (FEATURE-LogReading.md)
  intent: Blank lines are ignored and not counted as records or skipped lines.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestReadIgnoresBlankLines not-run
    {"Time":"2026-10-25T21:41:14.714210514-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:14.714294633-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:14.714301612-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:14.714210514-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:14.714294633-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:14.714301612-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestReadIgnoresBlankLines'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-blank-lines.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestReadIgnoresBlankLines', 'not-run', '{"Time":"2026-10-25T21:41:14.714210514-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:14.714294633-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:14.714301612-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: log-read-long-line (FEATURE-LogReading.md)
  intent: A record line longer than the default scanner token limit is read in full.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestReadAcceptsLongLine not-run
    {"Time":"2026-10-25T21:41:15.520860588-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:15.520942989-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:15.520950718-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:15.520860588-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:15.520942989-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:15.520950718-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestReadAcceptsLongLine'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-long-line.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestReadAcceptsLongLine', 'not-run', '{"Time":"2026-10-25T21:41:15.520860588-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:15.520942989-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:15.520950718-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: log-read-file-errors (FEATURE-LogReading.md)
  intent: ReadFile reports an error for a missing path and for a directory.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestReadFileMissing not-run
    {"Time":"2026-10-25T21:41:16.38506772-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:16.385135542-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:16.385140109-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:16.38506772-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:16.385135542-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:16.385140109-04:00","Action":"skip","Packa… (+68 chars)
      name = 'TestReadFileMissing'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-file-errors.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestReadFileMissing', 'not-run', '{"Time":"2026-10-25T21:41:16.38506772-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:16.385135542-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:16.385140109-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: log-read-read-only (FEATURE-LogReading.md)
  intent: Reading a log leaves the log file unchanged.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestReadNeverWritesLog not-run
    {"Time":"2026-10-25T21:41:17.239124942-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:17.239234488-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:17.239246847-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:17.239124942-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:17.239234488-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:17.239246847-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestReadNeverWritesLog'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-read-only.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestReadNeverWritesLog', 'not-run', '{"Time":"2026-10-25T21:41:17.239124942-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:17.239234488-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:17.239246847-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: date-options-accepts-real (FEATURE-DateOptions.md)
  intent: Real YYYY-MM-DD dates, including a leap day, are accepted by the date validator.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:18.139626746-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:41:18.146545625-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T21:41:18.146622647-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:41:18.14718949-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T21:41:18.147217214-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:18.139626746-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:18.146545625-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:18.146622647-04:00","Action":"output","Package":"ddkstat/internal… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-options-accepts-real.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:18.139626746-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:18.146545625-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:18.146622647-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-25T21:41:18.14718949-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T21:41:18.147217214-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- UNVERIFIED: date-options-rejects-malformed (FEATURE-DateOptions.md)
  intent: Non-padded, compact, padded, empty, and non-date values are rejected by the date validator.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:18.999168939-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:41:19.006160711-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T21:41:19.006243102-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:41:19.006852972-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T21:41:19.006883873-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:18.999168939-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:19.006160711-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:19.006243102-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-options-rejects-malformed.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:18.999168939-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:19.006160711-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:19.006243102-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-25T21:41:19.006852972-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T21:41:19.006883873-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- UNVERIFIED: date-options-rejects-impossible (FEATURE-DateOptions.md)
  intent: Well-formed but impossible dates such as month 13 or 30 February are rejected.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:19.937785382-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:41:19.9447958-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T21:41:19.944844455-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:41:19.945452796-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T21:41:19.945472354-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:19.937785382-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:19.9447958-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:19.944844455-04:00","Action":"output","Package":"ddkstat/internal/c… (+329 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-options-rejects-impossible.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:19.937785382-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:19.9447958-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:19.944844455-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-25T21:41:19.945452796-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T21:41:19.945472354-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- GREEN (prepassed): unreadable-log-missing-file (FEATURE-UnreadableLog.md)
  intent: A missing log exits 1 with empty standard output for every report.
  return code: 0
  stdout:
    summary 1
    ddkstat: report "summary" is not implemented yet
    
    commands 1
    ddkstat: report "commands" is not implemented yet
    
    targets 1
    ddkstat: report "targets" is not implemented yet
    
    models 1
    ddkstat: report "models" is not implemented yet
    
    daily 1
    ddkstat: report "daily" is not implemented yet
    
    failures 1
    ddkstat: report "failures" is not implemented yet
- GREEN (prepassed): unreadable-log-directory (FEATURE-UnreadableLog.md)
  intent: A directory named as the log exits 1 with empty standard output.
  return code: 0
  stdout:
    1
    ddkstat: report "summary" is not implemented yet
- UNVERIFIED: unreadable-log-no-partial-output (FEATURE-UnreadableLog.md)
  intent: A read error partway through the log leaves standard output empty.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:25.174786993-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:41:25.181555042-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T21:41:25.181599504-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:41:25.182556387-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T21:41:25.182582997-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:25.174786993-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:25.181555042-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:25.181599504-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "unreadable-log-no-partial-output.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:25.174786993-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:25.181555042-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:25.181599504-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-25T21:41:25.182556387-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T21:41:25.182582997-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- RED: usage-errors-exit-2 (FEATURE-UsageErrors.md)
  intent: Each kind of usage error exits 2 and writes nothing to standard output, with an existing log path supplied.
  return code: 1
  stdout:
    no-args 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    unknown-report 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-log 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    unknown-option 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    help-option 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    extra-log 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-value-target 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-value-since 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-value-until 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    bad-since-month 1
    ddkstat: report "summary" is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
      name = 'bad-since-month'
      out = ''
      err = 'ddkstat: report "summary" is not implemented yet\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-errors-exit-2.py", line 31, in <module>
        assert code == 2, (name, code, out, err)
               ^^^^^^^^^
    AssertionError: ('bad-since-month', 1, '', 'ddkstat: report "summary" is not implemented yet\n')
- RED: usage-errors-before-open (FEATURE-UsageErrors.md)
  intent: A usage error takes precedence over an unreadable log, so it exits 2 rather than 1.
  return code: 1
  stdout:
    ['frobnicate', '/tmp/ddkstat-no-such-log-usage.jsonl'] 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    ['summary', '--since', '2026-13-01', '/tmp/ddkstat-no-such-log-usage.jsonl'] 1
    ddkstat: report "summary" is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
      args = ['summary', '--since', '2026-13-01', '/tmp/ddkstat-no-such-log-usage.jsonl']
      out = ''
      err = 'ddkstat: report "summary" is not implemented yet\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-errors-before-open.py", line 14, in <module>
        assert code == 2, (args, code, out, err)
               ^^^^^^^^^
    AssertionError: (['summary', '--since', '2026-13-01', '/tmp/ddkstat-no-such-log-usage.jsonl'], 1, '', 'ddkstat: report "summary" is not implemented yet\n')
- UNVERIFIED: usage-errors-usage-line (FEATURE-UsageErrors.md)
  intent: The usage line names every report and every option.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:31.143238687-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:41:31.150243377-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T21:41:31.15029969-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:41:31.151094264-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T21:41:31.151133797-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:31.143238687-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:31.150243377-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:31.15029969-04:00","Action":"output","Package":"ddkstat/internal/… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-errors-usage-line.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:31.143238687-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-25T21:41:31.150243377-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T21:41:31.15029969-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-25T21:41:31.151094264-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T21:41:31.151133797-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- UNVERIFIED: malformed-not-object (FEATURE-MalformedLines.md)
  intent: Lines that are not JSON objects are skipped and counted rather than failing the read.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestSkipNotJSONObject not-run
    {"Time":"2026-10-25T21:41:31.997226678-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:31.997296696-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:31.997307483-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:31.997226678-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:31.997296696-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:31.997307483-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestSkipNotJSONObject'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "malformed-not-object.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestSkipNotJSONObject', 'not-run', '{"Time":"2026-10-25T21:41:31.997226678-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:31.997296696-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:31.997307483-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: malformed-invalid-utf8 (FEATURE-MalformedLines.md)
  intent: A line that is not valid UTF-8 is skipped as not a JSON object.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestSkipInvalidUTF8 not-run
    {"Time":"2026-10-25T21:41:32.879235832-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:32.879297692-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:32.879301634-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:32.879235832-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:32.879297692-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:32.879301634-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestSkipInvalidUTF8'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "malformed-invalid-utf8.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestSkipInvalidUTF8', 'not-run', '{"Time":"2026-10-25T21:41:32.879235832-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:32.879297692-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:32.879301634-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: malformed-required-fields (FEATURE-MalformedLines.md)
  intent: A line lacking a non-empty string in any required field is skipped.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestSkipMissingRequiredField not-run
    {"Time":"2026-10-25T21:41:33.72040897-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:33.720476433-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:33.720480755-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:33.72040897-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:33.720476433-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:33.720480755-04:00","Action":"skip","Packa… (+68 chars)
      name = 'TestSkipMissingRequiredField'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "malformed-required-fields.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestSkipMissingRequiredField', 'not-run', '{"Time":"2026-10-25T21:41:33.72040897-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:33.720476433-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:33.720480755-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: malformed-started-at (FEATURE-MalformedLines.md)
  intent: A started_at that does not begin with a real calendar date skips the line.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestSkipBadStartedAt not-run
    {"Time":"2026-10-25T21:41:34.561073411-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:34.56114316-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:34.561149156-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:34.561073411-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:34.56114316-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:34.561149156-04:00","Action":"skip","Packa… (+68 chars)
      name = 'TestSkipBadStartedAt'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "malformed-started-at.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestSkipBadStartedAt', 'not-run', '{"Time":"2026-10-25T21:41:34.561073411-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:34.56114316-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:34.561149156-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: malformed-warning-and-count (FEATURE-MalformedLines.md)
  intent: Each skipped line is warned about once with its line number and valid neighbours are kept.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestSkipWarnsWithLineNumber not-run
    {"Time":"2026-10-25T21:41:35.416539923-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:35.416608985-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:35.416632211-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:35.416539923-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:35.416608985-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:35.416632211-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestSkipWarnsWithLineNumber'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "malformed-warning-and-count.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestSkipWarnsWithLineNumber', 'not-run', '{"Time":"2026-10-25T21:41:35.416539923-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:35.416608985-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:35.416632211-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: duplicate-first-kept (FEATURE-DuplicateIds.md)
  intent: The first occurrence of an execution_id is kept and later repeats are skipped and counted.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestDuplicateFirstOccurrenceKept not-run
    {"Time":"2026-10-25T21:41:36.290092392-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:36.29015737-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:36.290162886-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:36.290092392-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:36.29015737-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:36.290162886-04:00","Action":"skip","Packa… (+68 chars)
      name = 'TestDuplicateFirstOccurrenceKept'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "duplicate-first-kept.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestDuplicateFirstOccurrenceKept', 'not-run', '{"Time":"2026-10-25T21:41:36.290092392-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:36.29015737-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:36.290162886-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: duplicate-skipped-line-does-not-register (FEATURE-DuplicateIds.md)
  intent: A line skipped for another reason does not register its id, so a later valid record with that id is kept.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestSkippedLineDoesNotRegisterID not-run
    {"Time":"2026-10-25T21:41:37.17130233-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:37.171404918-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:37.171414767-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:37.17130233-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:37.171404918-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:37.171414767-04:00","Action":"skip","Packa… (+68 chars)
      name = 'TestSkippedLineDoesNotRegisterID'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "duplicate-skipped-line-does-not-register.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestSkippedLineDoesNotRegisterID', 'not-run', '{"Time":"2026-10-25T21:41:37.17130233-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:37.171404918-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:37.171414767-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: duplicate-warning-and-count (FEATURE-DuplicateIds.md)
  intent: Each duplicate warns once with its line number and is counted together with malformed lines.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestDuplicateWarnsWithLineNumber not-run
    {"Time":"2026-10-25T21:41:38.059831111-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:38.059904976-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:38.059909497-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:38.059831111-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:38.059904976-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:38.059909497-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestDuplicateWarnsWithLineNumber'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "duplicate-warning-and-count.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestDuplicateWarnsWithLineNumber', 'not-run', '{"Time":"2026-10-25T21:41:38.059831111-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:38.059904976-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:38.059909497-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: absent-numbers-no-stats (FEATURE-AbsentNumbers.md)
  intent: A record with no result.stats object is kept as a record with zero numeric fields.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:38.975171548-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:38.975240865-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:38.975245582-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:38.975171548-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:38.975240865-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:38.975245582-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "absent-numbers-no-stats.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:38.975171548-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:38.975240865-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:38.975245582-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: absent-numbers-null-cost (FEATURE-AbsentNumbers.md)
  intent: A null cost_usd counts as zero while the other numeric fields of the record are still summed.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:39.789337077-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:39.789487972-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:39.78951078-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:39.789337077-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:39.789487972-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:39.78951078-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "absent-numbers-null-cost.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:39.789337077-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:39.789487972-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:39.78951078-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: absent-numbers-absent-fields (FEATURE-AbsentNumbers.md)
  intent: Each absent token count, cost, and duration counts as zero independently of the others.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:40.578904123-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:40.57896842-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:40.578972675-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:40.578904123-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:40.57896842-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:40.578972675-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "absent-numbers-absent-fields.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:40.578904123-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:40.57896842-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:40.578972675-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: absent-numbers-wrong-type (FEATURE-AbsentNumbers.md)
  intent: A numeric field with a JSON type other than number or null counts as zero and keeps the record.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:41.378506718-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:41.378585225-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:41.378590861-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:41.378506718-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:41.378585225-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:41.378590861-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "absent-numbers-wrong-type.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:41.378506718-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:41.378585225-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:41.378590861-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: absent-numbers-returncode (FEATURE-AbsentNumbers.md)
  intent: A null or absent returncode is recorded as absent rather than zero.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:42.186435185-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:42.186513363-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:42.186518216-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:42.186435185-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:42.186513363-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:42.186518216-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "absent-numbers-returncode.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:42.186435185-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:42.186513363-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:42.186518216-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: absent-numbers-timed-out (FEATURE-AbsentNumbers.md)
  intent: An absent timed_out defaults to false.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:42.950685031-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:42.950776485-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:42.950782488-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:42.950685031-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:42.950776485-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:42.950782488-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "absent-numbers-timed-out.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:42.950685031-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:42.950776485-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:42.950782488-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: model-attribution-executed (FEATURE-ModelAttribution.md)
  intent: A record with a non-empty result.stats.model is attributed to that model rather than job.model.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:43.77085127-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:43.770927382-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:43.770932642-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:43.77085127-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:43.770927382-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:43.770932642-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-attribution-executed.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:43.77085127-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:43.770927382-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:43.770932642-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: model-attribution-no-stats (FEATURE-ModelAttribution.md)
  intent: A record without result.stats is attributed to job.model.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:44.584281893-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:44.584347868-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:44.584353632-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:44.584281893-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:44.584347868-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:44.584353632-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-attribution-no-stats.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:44.584281893-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:44.584347868-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:44.584353632-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: model-attribution-empty (FEATURE-ModelAttribution.md)
  intent: A null, absent, or empty result.stats.model falls back to job.model.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:45.432251158-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:45.432334368-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:45.432340732-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:45.432251158-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:45.432334368-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:45.432340732-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-attribution-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:45.432251158-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:45.432334368-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:45.432340732-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')
- UNVERIFIED: model-attribution-provider (FEATURE-ModelAttribution.md)
  intent: The provider of a record is read from job.llm.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T21:41:46.290393392-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:41:46.29045557-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \tddkstat/internal/logread\t[no test files]\n"}
    {"Time":"2026-10-25T21:41:46.290460124-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T21:41:46.290393392-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:46.29045557-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:46.290460124-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-attribution-provider.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T21:41:46.290393392-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-25T21:41:46.29045557-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"?   \\tddkstat/internal/logread\\t[no test files]\\n"}\n{"Time":"2026-10-25T21:41:46.290460124-04:00","Action":"skip","Package":"ddkstat/internal/logread","Elapsed":0}\n', '')

## Post-build programmatic acceptance
- PASS: log-read-blank-lines (FEATURE-LogReading.md)
  intent: Blank lines are ignored and not counted as records or skipped lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestReadIgnoresBlankLines pass
    {"Time":"2026-10-25T21:45:52.574241303-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:45:52.581291873-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadIgnoresBlankLines"}
    {"Time":"2026-10-25T21:45:52.58134062-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadIgnoresBlankLines","Output":"=== RUN   TestReadIgnoresBlankLines\n"}
    {"Time":"2026-10-25T21:45:52.581413782-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadIgnoresBlankLines","Output":"--- PASS: TestReadIgnoresBlankLines (0.00s)\n"}
    {"Time":"2026-10-25T21:45:52.581436504-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadIgnoresBlankLines","Elapsed":0}
    {"Time":"2026-10-25T21:45:52.581443744-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:52.582077609-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:45:52.582104417-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
    
    TestReadNumbersEveryLine pass
    {"Time":"2026-10-25T21:45:53.331627809-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:45:53.338242436-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadNumbersEveryLine"}
    {"Time":"2026-10-25T21:45:53.338279254-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadNumbersEveryLine","Output":"=== RUN   TestReadNumbersEveryLine\n"}
    {"Time":"2026-10-25T21:45:53.338362393-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadNumbersEveryLine","Output":"--- PASS: TestReadNumbersEveryLine (0.00s)\n"}
    {"Time":"2026-10-25T21:45:53.338387636-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadNumbersEveryLine","Elapsed":0}
    {"Time":"2026-10-25T21:45:53.338393654-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:53.339091828-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:45:53.339116401-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: log-read-long-line (FEATURE-LogReading.md)
  intent: A record line longer than the default scanner token limit is read in full.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestReadAcceptsLongLine pass
    {"Time":"2026-10-25T21:45:54.459925578-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:45:54.465842309-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadAcceptsLongLine"}
    {"Time":"2026-10-25T21:45:54.465888594-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadAcceptsLongLine","Output":"=== RUN   TestReadAcceptsLongLine\n"}
    {"Time":"2026-10-25T21:45:54.48118652-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadAcceptsLongLine","Output":"--- PASS: TestReadAcceptsLongLine (0.02s)\n"}
    {"Time":"2026-10-25T21:45:54.481222273-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadAcceptsLongLine","Elapsed":0.02}
    {"Time":"2026-10-25T21:45:54.481233033-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:54.483094401-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.023s\n"}
    {"Time":"2026-10-25T21:45:54.483118217-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.023}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: log-read-file-errors (FEATURE-LogReading.md)
  intent: ReadFile reports an error for a missing path and for a directory.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestReadFileMissing pass
    {"Time":"2026-10-25T21:45:55.559100386-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:45:55.566489943-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadFileMissing"}
    {"Time":"2026-10-25T21:45:55.566545418-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadFileMissing","Output":"=== RUN   TestReadFileMissing\n"}
    {"Time":"2026-10-25T21:45:55.567047226-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadFileMissing","Output":"--- PASS: TestReadFileMissing (0.00s)\n"}
    {"Time":"2026-10-25T21:45:55.567078534-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadFileMissing","Elapsed":0}
    {"Time":"2026-10-25T21:45:55.567083175-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:55.567858802-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:45:55.567881505-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.009}
    
    TestReadFileDirectory pass
    {"Time":"2026-10-25T21:45:56.319632514-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:45:56.326391608-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadFileDirectory"}
    {"Time":"2026-10-25T21:45:56.326423458-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadFileDirectory","Output":"=== RUN   TestReadFileDirectory\n"}
    {"Time":"2026-10-25T21:45:56.326807713-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadFileDirectory","Output":"--- PASS: TestReadFileDirectory (0.00s)\n"}
    {"Time":"2026-10-25T21:45:56.326832637-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadFileDirectory","Elapsed":0}
    {"Time":"2026-10-25T21:45:56.326838722-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:56.327473215-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:45:56.327492558-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: log-read-read-only (FEATURE-LogReading.md)
  intent: Reading a log leaves the log file unchanged.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestReadNeverWritesLog pass
    {"Time":"2026-10-25T21:45:57.401558439-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:45:57.406969447-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadNeverWritesLog"}
    {"Time":"2026-10-25T21:45:57.407022612-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadNeverWritesLog","Output":"=== RUN   TestReadNeverWritesLog\n"}
    {"Time":"2026-10-25T21:45:57.407503094-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadNeverWritesLog","Output":"--- PASS: TestReadNeverWritesLog (0.00s)\n"}
    {"Time":"2026-10-25T21:45:57.407525748-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadNeverWritesLog","Elapsed":0}
    {"Time":"2026-10-25T21:45:57.407529335-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:57.408244944-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-25T21:45:57.408274355-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-options-accepts-real (FEATURE-DateOptions.md)
  intent: Real YYYY-MM-DD dates, including a leap day, are accepted by the date validator.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:45:58.560860158-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:45:58.568549411-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates"}
    {"Time":"2026-10-25T21:45:58.568604413-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates","Output":"=== RUN   TestParseDateAcceptsRealDates\n"}
    {"Time":"2026-10-25T21:45:58.568612606-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-10-14"}
    {"Time":"2026-10-25T21:45:58.568652769-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-10-14","Output":"=== RUN   TestParseDateAcceptsRealDates/2026-10-14\n"}
    {"Time":"2026-10-25T21:45:58.568702401-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-10-14","Output":"--- PASS: TestParseDateAcceptsRealDates/2026-10-14 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:58.568727783-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-10-14","Elapsed":0}
    {"Time":"2026-10-25T21:45:58.568736691-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-01-01"}
    {"Time":"2026-10-25T21:45:58.568738468-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-01-01","Output":"=== RUN   TestParseDateAcceptsRealDates/2026-01-01\n"}
    {"Time":"2026-10-25T21:45:58.568740661-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-01-01","Output":"--- PASS: TestParseDateAcceptsRealDates/2026-01-01 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:58.568742425-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-01-01","Elapsed":0}
    {"Time":"2026-10-25T21:45:58.568745455-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-12-31"}
    {"Time":"2026-10-25T21:45:58.568746728-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-12-31","Output":"=== RUN   TestParseDateAcceptsRealDates/2026-12-31\n"}
    {"Time":"2026-10-25T21:45:58.568748531-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-12-31","Output":"--- PASS: TestParseDateAcceptsRealDates/2026-12-31 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:58.568749801-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2026-12-31","Elapsed":0}
    {"Time":"2026-10-25T21:45:58.568750973-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2024-02-29"}
    {"Time":"2026-10-25T21:45:58.568752049-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2024-02-29","Output":"=== RUN   TestParseDateAcceptsRealDates/2024-02-29\n"}
    {"Time":"2026-10-25T21:45:58.568753421-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2024-02-29","Output":"--- PASS: TestParseDateAcceptsRealDates/2024-02-29 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:58.568754709-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2024-02-29","Elapsed":0}
    {"Time":"2026-10-25T21:45:58.568755946-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2000-02-29"}
    {"Time":"2026-10-25T21:45:58.568756997-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2000-02-29","Output":"=== RUN   TestParseDateAcceptsRealDates/2000-02-29\n"}
    {"Time":"2026-10-25T21:45:58.568758372-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2000-02-29","Output":"--- PASS: TestParseDateAcceptsRealDates/2000-02-29 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:58.568759594-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates/2000-02-29","Elapsed":0}
    {"Time":"2026-10-25T21:45:58.568760957-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates","Output":"--- PASS: TestParseDateAcceptsRealDates (0.00s)\n"}
    {"Time":"2026-10-25T21:45:58.568904561-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateAcceptsRealDates","Elapsed":0}
    {"Time":"2026-10-25T21:45:58.568913813-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:58.570162034-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.009s\n"}
    {"Time":"2026-10-25T21:45:58.570187046-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-options-rejects-malformed (FEATURE-DateOptions.md)
  intent: Non-padded, compact, padded, empty, and non-date values are rejected by the date validator.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:45:59.583161329-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:45:59.588849489-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed"}
    {"Time":"2026-10-25T21:45:59.588899415-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed","Output":"=== RUN   TestParseDateRejectsMalformed\n"}
    {"Time":"2026-10-25T21:45:59.588920039-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-1-5"}
    {"Time":"2026-10-25T21:45:59.588922102-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-1-5","Output":"=== RUN   TestParseDateRejectsMalformed/2026-1-5\n"}
    {"Time":"2026-10-25T21:45:59.588926763-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-1-5","Output":"--- PASS: TestParseDateRejectsMalformed/2026-1-5 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588928884-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-1-5","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.58893194-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-5"}
    {"Time":"2026-10-25T21:45:59.588933209-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-5","Output":"=== RUN   TestParseDateRejectsMalformed/2026-10-5\n"}
    {"Time":"2026-10-25T21:45:59.588935138-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-5","Output":"--- PASS: TestParseDateRejectsMalformed/2026-10-5 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588936556-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-5","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588937951-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/20261014"}
    {"Time":"2026-10-25T21:45:59.588939058-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/20261014","Output":"=== RUN   TestParseDateRejectsMalformed/20261014\n"}
    {"Time":"2026-10-25T21:45:59.588940645-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/20261014","Output":"--- PASS: TestParseDateRejectsMalformed/20261014 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588942028-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/20261014","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588943365-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026/10/14"}
    {"Time":"2026-10-25T21:45:59.588945068-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026/10/14","Output":"=== RUN   TestParseDateRejectsMalformed/2026/10/14\n"}
    {"Time":"2026-10-25T21:45:59.588946677-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026/10/14","Output":"--- PASS: TestParseDateRejectsMalformed/2026/10/14 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588947964-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026/10/14","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588949237-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14_"}
    {"Time":"2026-10-25T21:45:59.588950356-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14_","Output":"=== RUN   TestParseDateRejectsMalformed/2026-10-14_\n"}
    {"Time":"2026-10-25T21:45:59.588951942-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14_","Output":"--- PASS: TestParseDateRejectsMalformed/2026-10-14_ (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588954493-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14_","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588955905-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/_2026-10-14"}
    {"Time":"2026-10-25T21:45:59.58895705-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/_2026-10-14","Output":"=== RUN   TestParseDateRejectsMalformed/_2026-10-14\n"}
    {"Time":"2026-10-25T21:45:59.588958538-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/_2026-10-14","Output":"--- PASS: TestParseDateRejectsMalformed/_2026-10-14 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588959846-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/_2026-10-14","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588961108-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14T00"}
    {"Time":"2026-10-25T21:45:59.588962198-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14T00","Output":"=== RUN   TestParseDateRejectsMalformed/2026-10-14T00\n"}
    {"Time":"2026-10-25T21:45:59.588963783-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14T00","Output":"--- PASS: TestParseDateRejectsMalformed/2026-10-14T00 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588965625-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-14T00","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588967665-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-144"}
    {"Time":"2026-10-25T21:45:59.588968745-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-144","Output":"=== RUN   TestParseDateRejectsMalformed/2026-10-144\n"}
    {"Time":"2026-10-25T21:45:59.588970185-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-144","Output":"--- PASS: TestParseDateRejectsMalformed/2026-10-144 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588971472-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/2026-10-144","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588972754-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/#00"}
    {"Time":"2026-10-25T21:45:59.588973817-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/#00","Output":"=== RUN   TestParseDateRejectsMalformed/#00\n"}
    {"Time":"2026-10-25T21:45:59.588975317-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/#00","Output":"--- PASS: TestParseDateRejectsMalformed/#00 (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588976773-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/#00","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588978168-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/yesterday"}
    {"Time":"2026-10-25T21:45:59.588979341-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/yesterday","Output":"=== RUN   TestParseDateRejectsMalformed/yesterday\n"}
    {"Time":"2026-10-25T21:45:59.588980878-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/yesterday","Output":"--- PASS: TestParseDateRejectsMalformed/yesterday (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588982296-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed/yesterday","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588983752-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed","Output":"--- PASS: TestParseDateRejectsMalformed (0.00s)\n"}
    {"Time":"2026-10-25T21:45:59.588985258-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsMalformed","Elapsed":0}
    {"Time":"2026-10-25T21:45:59.588987971-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:45:59.590424194-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T21:45:59.590455964-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-options-rejects-impossible (FEATURE-DateOptions.md)
  intent: Well-formed but impossible dates such as month 13 or 30 February are rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:00.586150325-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:46:00.593262415-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates"}
    {"Time":"2026-10-25T21:46:00.593296778-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates","Output":"=== RUN   TestParseDateRejectsImpossibleDates\n"}
    {"Time":"2026-10-25T21:46:00.593316472-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-13-01"}
    {"Time":"2026-10-25T21:46:00.593318062-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-13-01","Output":"=== RUN   TestParseDateRejectsImpossibleDates/2026-13-01\n"}
    {"Time":"2026-10-25T21:46:00.593321827-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-13-01","Output":"--- PASS: TestParseDateRejectsImpossibleDates/2026-13-01 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593323837-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-13-01","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.59332739-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-00-10"}
    {"Time":"2026-10-25T21:46:00.593328652-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-00-10","Output":"=== RUN   TestParseDateRejectsImpossibleDates/2026-00-10\n"}
    {"Time":"2026-10-25T21:46:00.593332013-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-00-10","Output":"--- PASS: TestParseDateRejectsImpossibleDates/2026-00-10 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593333523-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-00-10","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.593334991-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-10-00"}
    {"Time":"2026-10-25T21:46:00.593336089-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-10-00","Output":"=== RUN   TestParseDateRejectsImpossibleDates/2026-10-00\n"}
    {"Time":"2026-10-25T21:46:00.593343582-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-10-00","Output":"--- PASS: TestParseDateRejectsImpossibleDates/2026-10-00 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593345025-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-10-00","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.593346901-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-02-30"}
    {"Time":"2026-10-25T21:46:00.59334798-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-02-30","Output":"=== RUN   TestParseDateRejectsImpossibleDates/2026-02-30\n"}
    {"Time":"2026-10-25T21:46:00.593364259-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-02-30","Output":"--- PASS: TestParseDateRejectsImpossibleDates/2026-02-30 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593376516-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-02-30","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.593378597-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2025-02-29"}
    {"Time":"2026-10-25T21:46:00.593380298-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2025-02-29","Output":"=== RUN   TestParseDateRejectsImpossibleDates/2025-02-29\n"}
    {"Time":"2026-10-25T21:46:00.593382947-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2025-02-29","Output":"--- PASS: TestParseDateRejectsImpossibleDates/2025-02-29 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593384441-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2025-02-29","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.593385786-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-04-31"}
    {"Time":"2026-10-25T21:46:00.593386905-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-04-31","Output":"=== RUN   TestParseDateRejectsImpossibleDates/2026-04-31\n"}
    {"Time":"2026-10-25T21:46:00.593388475-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-04-31","Output":"--- PASS: TestParseDateRejectsImpossibleDates/2026-04-31 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593389916-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates/2026-04-31","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.593391374-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates","Output":"--- PASS: TestParseDateRejectsImpossibleDates (0.00s)\n"}
    {"Time":"2026-10-25T21:46:00.593392799-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestParseDateRejectsImpossibleDates","Elapsed":0}
    {"Time":"2026-10-25T21:46:00.593394606-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:00.594517224-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s\n"}
    {"Time":"2026-10-25T21:46:00.594547875-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: unreadable-log-missing-file (FEATURE-UnreadableLog.md)
  intent: A missing log exits 1 with empty standard output for every report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    summary 1
    ddkstat: cannot read log: open /tmp/ddkstat-no-such-log-unreadable.jsonl: no such file or directory
    
    commands 1
    ddkstat: cannot read log: open /tmp/ddkstat-no-such-log-unreadable.jsonl: no such file or directory
    
    targets 1
    ddkstat: cannot read log: open /tmp/ddkstat-no-such-log-unreadable.jsonl: no such file or directory
    
    models 1
    ddkstat: cannot read log: open /tmp/ddkstat-no-such-log-unreadable.jsonl: no such file or directory
    
    daily 1
    ddkstat: cannot read log: open /tmp/ddkstat-no-such-log-unreadable.jsonl: no such file or directory
    
    failures 1
    ddkstat: cannot read log: open /tmp/ddkstat-no-such-log-unreadable.jsonl: no such file or directory
- PASS: unreadable-log-directory (FEATURE-UnreadableLog.md)
  intent: A directory named as the log exits 1 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    1
    ddkstat: cannot read log: /tmp/tmpogtz10af is a directory
- PASS: unreadable-log-no-partial-output (FEATURE-UnreadableLog.md)
  intent: A read error partway through the log leaves standard output empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:06.231144657-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:46:06.238487321-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestReadErrorLeavesStdoutEmpty"}
    {"Time":"2026-10-25T21:46:06.238544027-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestReadErrorLeavesStdoutEmpty","Output":"=== RUN   TestReadErrorLeavesStdoutEmpty\n"}
    {"Time":"2026-10-25T21:46:06.238553831-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestReadErrorLeavesStdoutEmpty","Output":"--- PASS: TestReadErrorLeavesStdoutEmpty (0.00s)\n"}
    {"Time":"2026-10-25T21:46:06.238555968-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestReadErrorLeavesStdoutEmpty","Elapsed":0}
    {"Time":"2026-10-25T21:46:06.238559274-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:06.239771995-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.008s\n"}
    {"Time":"2026-10-25T21:46:06.239817249-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: usage-errors-exit-2 (FEATURE-UsageErrors.md)
  intent: Each kind of usage error exits 2 and writes nothing to standard output, with an existing log path supplied.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    no-args 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    unknown-report 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-log 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    unknown-option 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    help-option 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    extra-log 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-value-target 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-value-since 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    missing-value-until 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    bad-since-month 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    bad-until-day 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    bad-since-form 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
- PASS: usage-errors-before-open (FEATURE-UsageErrors.md)
  intent: A usage error takes precedence over an unreadable log, so it exits 2 rather than 1.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ['frobnicate', '/tmp/ddkstat-no-such-log-usage.jsonl'] 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    ['summary', '--since', '2026-13-01', '/tmp/ddkstat-no-such-log-usage.jsonl'] 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
    
    ['summary', '--colour', 'red', '/tmp/ddkstat-no-such-log-usage.jsonl'] 2
    usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]
- PASS: usage-errors-usage-line (FEATURE-UsageErrors.md)
  intent: The usage line names every report and every option.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:11.889302934-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-25T21:46:11.896111848-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestUsageLineNamesReportsAndOptions"}
    {"Time":"2026-10-25T21:46:11.896154429-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestUsageLineNamesReportsAndOptions","Output":"=== RUN   TestUsageLineNamesReportsAndOptions\n"}
    {"Time":"2026-10-25T21:46:11.896170859-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestUsageLineNamesReportsAndOptions","Output":"--- PASS: TestUsageLineNamesReportsAndOptions (0.00s)\n"}
    {"Time":"2026-10-25T21:46:11.896173405-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestUsageLineNamesReportsAndOptions","Elapsed":0}
    {"Time":"2026-10-25T21:46:11.896176288-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:11.897010289-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T21:46:11.897037669-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: malformed-not-object (FEATURE-MalformedLines.md)
  intent: Lines that are not JSON objects are skipped and counted rather than failing the read.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestSkipNotJSONObject pass
    {"Time":"2026-10-25T21:46:13.036410606-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:13.043215347-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject"}
    {"Time":"2026-10-25T21:46:13.043267591-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject","Output":"=== RUN   TestSkipNotJSONObject\n"}
    {"Time":"2026-10-25T21:46:13.043288039-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/truncated"}
    {"Time":"2026-10-25T21:46:13.043289561-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/truncated","Output":"=== RUN   TestSkipNotJSONObject/truncated\n"}
    {"Time":"2026-10-25T21:46:13.043297785-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/truncated","Output":"--- PASS: TestSkipNotJSONObject/truncated (0.00s)\n"}
    {"Time":"2026-10-25T21:46:13.043299622-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/truncated","Elapsed":0}
    {"Time":"2026-10-25T21:46:13.043303886-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/array"}
    {"Time":"2026-10-25T21:46:13.043305113-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/array","Output":"=== RUN   TestSkipNotJSONObject/array\n"}
    {"Time":"2026-10-25T21:46:13.043306988-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/array","Output":"--- PASS: TestSkipNotJSONObject/array (0.00s)\n"}
    {"Time":"2026-10-25T21:46:13.043308521-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/array","Elapsed":0}
    {"Time":"2026-10-25T21:46:13.043309929-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/number"}
    {"Time":"2026-10-25T21:46:13.043311147-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/number","Output":"=== RUN   TestSkipNotJSONObject/number\n"}
    {"Time":"2026-10-25T21:46:13.043312751-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/number","Output":"--- PASS: TestSkipNotJSONObject/number (0.00s)\n"}
    {"Time":"2026-10-25T21:46:13.043314132-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/number","Elapsed":0}
    {"Time":"2026-10-25T21:46:13.043315413-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/string"}
    {"Time":"2026-10-25T21:46:13.043316568-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/string","Output":"=== RUN   TestSkipNotJSONObject/string\n"}
    {"Time":"2026-10-25T21:46:13.04331807-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/string","Output":"--- PASS: TestSkipNotJSONObject/string (0.00s)\n"}
    {"Time":"2026-10-25T21:46:13.043319783-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/string","Elapsed":0}
    {"Time":"2026-10-25T21:46:13.043321117-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/null"}
    {"Time":"2026-10-25T21:46:13.043322247-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/null","Output":"=== RUN   TestSkipNotJSONObject/null\n"}
    {"Time":"2026-10-25T21:46:13.043323726-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/null","Output":"--- PASS: TestSkipNotJSONObject/null (0.00s)\n"}
    {"Time":"2026-10-25T21:46:13.043325039-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject/null","Elapsed":0}
    {"Time":"2026-10-25T21:46:13.043326539-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject","Output":"--- PASS: TestSkipNotJSONObject (0.00s)\n"}
    {"Time":"2026-10-25T21:46:13.043330086-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotJSONObject","Elapsed":0}
    {"Time":"2026-10-25T21:46:13.043331724-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:13.043955988-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:13.043984382-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: malformed-invalid-utf8 (FEATURE-MalformedLines.md)
  intent: A line that is not valid UTF-8 is skipped as not a JSON object.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestSkipInvalidUTF8 pass
    {"Time":"2026-10-25T21:46:14.129520553-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:14.135700463-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipInvalidUTF8"}
    {"Time":"2026-10-25T21:46:14.13574112-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipInvalidUTF8","Output":"=== RUN   TestSkipInvalidUTF8\n"}
    {"Time":"2026-10-25T21:46:14.135767074-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipInvalidUTF8","Output":"--- PASS: TestSkipInvalidUTF8 (0.00s)\n"}
    {"Time":"2026-10-25T21:46:14.135794991-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipInvalidUTF8","Elapsed":0}
    {"Time":"2026-10-25T21:46:14.135809648-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:14.136694372-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:14.136727285-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: malformed-required-fields (FEATURE-MalformedLines.md)
  intent: A line lacking a non-empty string in any required field is skipped.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestSkipMissingRequiredField pass
    {"Time":"2026-10-25T21:46:15.240171596-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:15.24639399-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField"}
    {"Time":"2026-10-25T21:46:15.246458076-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField","Output":"=== RUN   TestSkipMissingRequiredField\n"}
    {"Time":"2026-10-25T21:46:15.246469566-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_number"}
    {"Time":"2026-10-25T21:46:15.246472561-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_number","Output":"=== RUN   TestSkipMissingRequiredField/execution_id_number\n"}
    {"Time":"2026-10-25T21:46:15.246482556-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_number","Output":"--- PASS: TestSkipMissingRequiredField/execution_id_number (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246536665-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_number","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246557257-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_null"}
    {"Time":"2026-10-25T21:46:15.246559941-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_null","Output":"=== RUN   TestSkipMissingRequiredField/started_at_null\n"}
    {"Time":"2026-10-25T21:46:15.246567071-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_null","Output":"--- PASS: TestSkipMissingRequiredField/started_at_null (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246569315-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_null","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.24657268-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_empty"}
    {"Time":"2026-10-25T21:46:15.246574079-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_empty","Output":"=== RUN   TestSkipMissingRequiredField/command_name_empty\n"}
    {"Time":"2026-10-25T21:46:15.246620724-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_empty","Output":"--- PASS: TestSkipMissingRequiredField/command_name_empty (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246634319-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_empty","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.24663768-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_absent"}
    {"Time":"2026-10-25T21:46:15.24663926-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_absent","Output":"=== RUN   TestSkipMissingRequiredField/target_absent\n"}
    {"Time":"2026-10-25T21:46:15.246679001-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_absent","Output":"--- PASS: TestSkipMissingRequiredField/target_absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246710012-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246715172-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_absent"}
    {"Time":"2026-10-25T21:46:15.246717055-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_absent","Output":"=== RUN   TestSkipMissingRequiredField/job_absent\n"}
    {"Time":"2026-10-25T21:46:15.246719533-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_absent","Output":"--- PASS: TestSkipMissingRequiredField/job_absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246721922-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246725017-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_not_an_object"}
    {"Time":"2026-10-25T21:46:15.246726301-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_not_an_object","Output":"=== RUN   TestSkipMissingRequiredField/job_not_an_object\n"}
    {"Time":"2026-10-25T21:46:15.246758164-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_not_an_object","Output":"--- PASS: TestSkipMissingRequiredField/job_not_an_object (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246771774-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/job_not_an_object","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246773866-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_absent"}
    {"Time":"2026-10-25T21:46:15.246775362-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_absent","Output":"=== RUN   TestSkipMissingRequiredField/execution_id_absent\n"}
    {"Time":"2026-10-25T21:46:15.246832934-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_absent","Output":"--- PASS: TestSkipMissingRequiredField/execution_id_absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246847932-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246850321-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_empty"}
    {"Time":"2026-10-25T21:46:15.246851813-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_empty","Output":"=== RUN   TestSkipMissingRequiredField/execution_id_empty\n"}
    {"Time":"2026-10-25T21:46:15.24688074-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_empty","Output":"--- PASS: TestSkipMissingRequiredField/execution_id_empty (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246895371-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/execution_id_empty","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246897842-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_absent"}
    {"Time":"2026-10-25T21:46:15.246899454-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_absent","Output":"=== RUN   TestSkipMissingRequiredField/started_at_absent\n"}
    {"Time":"2026-10-25T21:46:15.246935766-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_absent","Output":"--- PASS: TestSkipMissingRequiredField/started_at_absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246949313-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/started_at_absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246951953-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_absent"}
    {"Time":"2026-10-25T21:46:15.2469535-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_absent","Output":"=== RUN   TestSkipMissingRequiredField/status_absent\n"}
    {"Time":"2026-10-25T21:46:15.246956488-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_absent","Output":"--- PASS: TestSkipMissingRequiredField/status_absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.246958132-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.246960194-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_boolean"}
    {"Time":"2026-10-25T21:46:15.246961384-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_boolean","Output":"=== RUN   TestSkipMissingRequiredField/status_boolean\n"}
    {"Time":"2026-10-25T21:46:15.247000612-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_boolean","Output":"--- PASS: TestSkipMissingRequiredField/status_boolean (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.247028807-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/status_boolean","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.247031833-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_absent"}
    {"Time":"2026-10-25T21:46:15.247033502-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_absent","Output":"=== RUN   TestSkipMissingRequiredField/command_name_absent\n"}
    {"Time":"2026-10-25T21:46:15.247036755-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_absent","Output":"--- PASS: TestSkipMissingRequiredField/command_name_absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.247045682-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/command_name_absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.247047356-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_null"}
    {"Time":"2026-10-25T21:46:15.24704999-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_null","Output":"=== RUN   TestSkipMissingRequiredField/target_null\n"}
    {"Time":"2026-10-25T21:46:15.247081874-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_null","Output":"--- PASS: TestSkipMissingRequiredField/target_null (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.247096765-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField/target_null","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.247099624-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField","Output":"--- PASS: TestSkipMissingRequiredField (0.00s)\n"}
    {"Time":"2026-10-25T21:46:15.247104623-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingRequiredField","Elapsed":0}
    {"Time":"2026-10-25T21:46:15.24710913-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:15.248441483-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:15.248482781-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: malformed-started-at (FEATURE-MalformedLines.md)
  intent: A started_at that does not begin with a real calendar date skips the line.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestSkipBadStartedAt pass
    {"Time":"2026-10-25T21:46:16.28706486-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:16.293225743-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt"}
    {"Time":"2026-10-25T21:46:16.293270264-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt","Output":"=== RUN   TestSkipBadStartedAt\n"}
    {"Time":"2026-10-25T21:46:16.293289104-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-13-01T00:00:00.000Z"}
    {"Time":"2026-10-25T21:46:16.293291124-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-13-01T00:00:00.000Z","Output":"=== RUN   TestSkipBadStartedAt/2026-13-01T00:00:00.000Z\n"}
    {"Time":"2026-10-25T21:46:16.293336266-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-13-01T00:00:00.000Z","Output":"--- PASS: TestSkipBadStartedAt/2026-13-01T00:00:00.000Z (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293358153-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-13-01T00:00:00.000Z","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.293363472-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-02-30T00:00:00.000Z"}
    {"Time":"2026-10-25T21:46:16.293365032-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-02-30T00:00:00.000Z","Output":"=== RUN   TestSkipBadStartedAt/2026-02-30T00:00:00.000Z\n"}
    {"Time":"2026-10-25T21:46:16.293367162-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-02-30T00:00:00.000Z","Output":"--- PASS: TestSkipBadStartedAt/2026-02-30T00:00:00.000Z (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293372884-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-02-30T00:00:00.000Z","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.29337496-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2025-02-29T00:00:00.000Z"}
    {"Time":"2026-10-25T21:46:16.293376108-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2025-02-29T00:00:00.000Z","Output":"=== RUN   TestSkipBadStartedAt/2025-02-29T00:00:00.000Z\n"}
    {"Time":"2026-10-25T21:46:16.293420762-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2025-02-29T00:00:00.000Z","Output":"--- PASS: TestSkipBadStartedAt/2025-02-29T00:00:00.000Z (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293434428-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2025-02-29T00:00:00.000Z","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.29343737-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-1-5T00:00:00.000Z"}
    {"Time":"2026-10-25T21:46:16.293438812-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-1-5T00:00:00.000Z","Output":"=== RUN   TestSkipBadStartedAt/2026-1-5T00:00:00.000Z\n"}
    {"Time":"2026-10-25T21:46:16.293470742-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-1-5T00:00:00.000Z","Output":"--- PASS: TestSkipBadStartedAt/2026-1-5T00:00:00.000Z (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293481819-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/2026-1-5T00:00:00.000Z","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.293483671-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/20261014T000000Z"}
    {"Time":"2026-10-25T21:46:16.293485038-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/20261014T000000Z","Output":"=== RUN   TestSkipBadStartedAt/20261014T000000Z\n"}
    {"Time":"2026-10-25T21:46:16.293538367-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/20261014T000000Z","Output":"--- PASS: TestSkipBadStartedAt/20261014T000000Z (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293551342-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/20261014T000000Z","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.293554533-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/yesterday"}
    {"Time":"2026-10-25T21:46:16.293556801-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/yesterday","Output":"=== RUN   TestSkipBadStartedAt/yesterday\n"}
    {"Time":"2026-10-25T21:46:16.293586022-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/yesterday","Output":"--- PASS: TestSkipBadStartedAt/yesterday (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293600888-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt/yesterday","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.2936037-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt","Output":"--- PASS: TestSkipBadStartedAt (0.00s)\n"}
    {"Time":"2026-10-25T21:46:16.293641485-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipBadStartedAt","Elapsed":0}
    {"Time":"2026-10-25T21:46:16.293662714-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:16.294335689-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:16.294356745-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: malformed-warning-and-count (FEATURE-MalformedLines.md)
  intent: Each skipped line is warned about once with its line number and valid neighbours are kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestSkipWarnsWithLineNumber pass
    {"Time":"2026-10-25T21:46:17.457076572-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:17.463827461-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipWarnsWithLineNumber"}
    {"Time":"2026-10-25T21:46:17.463874371-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipWarnsWithLineNumber","Output":"=== RUN   TestSkipWarnsWithLineNumber\n"}
    {"Time":"2026-10-25T21:46:17.463983687-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipWarnsWithLineNumber","Output":"--- PASS: TestSkipWarnsWithLineNumber (0.00s)\n"}
    {"Time":"2026-10-25T21:46:17.463999474-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipWarnsWithLineNumber","Elapsed":0}
    {"Time":"2026-10-25T21:46:17.464003178-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:17.464891801-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:17.464918096-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
    
    TestSkipDoesNotAffectValidRecords pass
    {"Time":"2026-10-25T21:46:18.206442232-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:18.21276192-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipDoesNotAffectValidRecords"}
    {"Time":"2026-10-25T21:46:18.212832024-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipDoesNotAffectValidRecords","Output":"=== RUN   TestSkipDoesNotAffectValidRecords\n"}
    {"Time":"2026-10-25T21:46:18.21292552-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipDoesNotAffectValidRecords","Output":"--- PASS: TestSkipDoesNotAffectValidRecords (0.00s)\n"}
    {"Time":"2026-10-25T21:46:18.212968764-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipDoesNotAffectValidRecords","Elapsed":0}
    {"Time":"2026-10-25T21:46:18.212982016-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:18.213528934-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:18.213555263-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: duplicate-first-kept (FEATURE-DuplicateIds.md)
  intent: The first occurrence of an execution_id is kept and later repeats are skipped and counted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestDuplicateFirstOccurrenceKept pass
    {"Time":"2026-10-25T21:46:19.267993549-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:19.273584797-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestDuplicateFirstOccurrenceKept"}
    {"Time":"2026-10-25T21:46:19.273616766-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateFirstOccurrenceKept","Output":"=== RUN   TestDuplicateFirstOccurrenceKept\n"}
    {"Time":"2026-10-25T21:46:19.27373872-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateFirstOccurrenceKept","Output":"--- PASS: TestDuplicateFirstOccurrenceKept (0.00s)\n"}
    {"Time":"2026-10-25T21:46:19.273749191-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestDuplicateFirstOccurrenceKept","Elapsed":0}
    {"Time":"2026-10-25T21:46:19.273756-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:19.274443528-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-25T21:46:19.274470242-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
    
    TestDuplicateKeepsFirstEvenWhenFieldsDiffer pass
    {"Time":"2026-10-25T21:46:20.034127211-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:20.041458979-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestDuplicateKeepsFirstEvenWhenFieldsDiffer"}
    {"Time":"2026-10-25T21:46:20.041506838-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateKeepsFirstEvenWhenFieldsDiffer","Output":"=== RUN   TestDuplicateKeepsFirstEvenWhenFieldsDiffer\n"}
    {"Time":"2026-10-25T21:46:20.04156382-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateKeepsFirstEvenWhenFieldsDiffer","Output":"--- PASS: TestDuplicateKeepsFirstEvenWhenFieldsDiffer (0.00s)\n"}
    {"Time":"2026-10-25T21:46:20.041570364-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestDuplicateKeepsFirstEvenWhenFieldsDiffer","Elapsed":0}
    {"Time":"2026-10-25T21:46:20.041577111-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:20.042279521-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:20.042312375-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: duplicate-skipped-line-does-not-register (FEATURE-DuplicateIds.md)
  intent: A line skipped for another reason does not register its id, so a later valid record with that id is kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestSkippedLineDoesNotRegisterID pass
    {"Time":"2026-10-25T21:46:21.107569803-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:21.114028834-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkippedLineDoesNotRegisterID"}
    {"Time":"2026-10-25T21:46:21.114069838-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkippedLineDoesNotRegisterID","Output":"=== RUN   TestSkippedLineDoesNotRegisterID\n"}
    {"Time":"2026-10-25T21:46:21.114159681-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkippedLineDoesNotRegisterID","Output":"--- PASS: TestSkippedLineDoesNotRegisterID (0.00s)\n"}
    {"Time":"2026-10-25T21:46:21.114183319-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkippedLineDoesNotRegisterID","Elapsed":0}
    {"Time":"2026-10-25T21:46:21.114190091-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:21.115043072-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:21.115063913-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: duplicate-warning-and-count (FEATURE-DuplicateIds.md)
  intent: Each duplicate warns once with its line number and is counted together with malformed lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestDuplicateWarnsWithLineNumber pass
    {"Time":"2026-10-25T21:46:22.238794759-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:22.245913456-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestDuplicateWarnsWithLineNumber"}
    {"Time":"2026-10-25T21:46:22.245960841-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateWarnsWithLineNumber","Output":"=== RUN   TestDuplicateWarnsWithLineNumber\n"}
    {"Time":"2026-10-25T21:46:22.246106359-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateWarnsWithLineNumber","Output":"--- PASS: TestDuplicateWarnsWithLineNumber (0.00s)\n"}
    {"Time":"2026-10-25T21:46:22.246129705-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestDuplicateWarnsWithLineNumber","Elapsed":0}
    {"Time":"2026-10-25T21:46:22.246135502-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:22.247023535-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:22.247060553-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
    
    TestDuplicateCountedWithMalformed pass
    {"Time":"2026-10-25T21:46:22.942765874-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:22.949826071-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestDuplicateCountedWithMalformed"}
    {"Time":"2026-10-25T21:46:22.949889641-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateCountedWithMalformed","Output":"=== RUN   TestDuplicateCountedWithMalformed\n"}
    {"Time":"2026-10-25T21:46:22.949992701-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestDuplicateCountedWithMalformed","Output":"--- PASS: TestDuplicateCountedWithMalformed (0.00s)\n"}
    {"Time":"2026-10-25T21:46:22.950010025-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestDuplicateCountedWithMalformed","Elapsed":0}
    {"Time":"2026-10-25T21:46:22.950016398-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:22.951031958-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:22.95105815-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: absent-numbers-no-stats (FEATURE-AbsentNumbers.md)
  intent: A record with no result.stats object is kept as a record with zero numeric fields.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:24.048049197-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:24.055167337-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentStatsIsStillRecord"}
    {"Time":"2026-10-25T21:46:24.055212678-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentStatsIsStillRecord","Output":"=== RUN   TestAbsentStatsIsStillRecord\n"}
    {"Time":"2026-10-25T21:46:24.055421633-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentStatsIsStillRecord","Output":"--- PASS: TestAbsentStatsIsStillRecord (0.00s)\n"}
    {"Time":"2026-10-25T21:46:24.055463841-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentStatsIsStillRecord","Elapsed":0}
    {"Time":"2026-10-25T21:46:24.055499413-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:24.05613031-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:24.056159203-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: absent-numbers-null-cost (FEATURE-AbsentNumbers.md)
  intent: A null cost_usd counts as zero while the other numeric fields of the record are still summed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:25.187826021-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:25.193067016-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestNullCostCountsZero"}
    {"Time":"2026-10-25T21:46:25.19312156-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNullCostCountsZero","Output":"=== RUN   TestNullCostCountsZero\n"}
    {"Time":"2026-10-25T21:46:25.193173845-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNullCostCountsZero","Output":"--- PASS: TestNullCostCountsZero (0.00s)\n"}
    {"Time":"2026-10-25T21:46:25.193188034-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestNullCostCountsZero","Elapsed":0}
    {"Time":"2026-10-25T21:46:25.19319183-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:25.193692056-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-25T21:46:25.193715601-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: absent-numbers-absent-fields (FEATURE-AbsentNumbers.md)
  intent: Each absent token count, cost, and duration counts as zero independently of the others.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:26.231945888-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:26.238677851-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero"}
    {"Time":"2026-10-25T21:46:26.238702883-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero","Output":"=== RUN   TestAbsentNumbersCountZero\n"}
    {"Time":"2026-10-25T21:46:26.238719996-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/input_tokens"}
    {"Time":"2026-10-25T21:46:26.238724389-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/input_tokens","Output":"=== RUN   TestAbsentNumbersCountZero/input_tokens\n"}
    {"Time":"2026-10-25T21:46:26.238801718-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/input_tokens","Output":"--- PASS: TestAbsentNumbersCountZero/input_tokens (0.00s)\n"}
    {"Time":"2026-10-25T21:46:26.238815235-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/input_tokens","Elapsed":0}
    {"Time":"2026-10-25T21:46:26.238819848-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cached_input_tokens"}
    {"Time":"2026-10-25T21:46:26.238821431-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cached_input_tokens","Output":"=== RUN   TestAbsentNumbersCountZero/cached_input_tokens\n"}
    {"Time":"2026-10-25T21:46:26.238864406-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cached_input_tokens","Output":"--- PASS: TestAbsentNumbersCountZero/cached_input_tokens (0.00s)\n"}
    {"Time":"2026-10-25T21:46:26.238877725-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cached_input_tokens","Elapsed":0}
    {"Time":"2026-10-25T21:46:26.23888016-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/output_tokens"}
    {"Time":"2026-10-25T21:46:26.238881655-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/output_tokens","Output":"=== RUN   TestAbsentNumbersCountZero/output_tokens\n"}
    {"Time":"2026-10-25T21:46:26.238929254-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/output_tokens","Output":"--- PASS: TestAbsentNumbersCountZero/output_tokens (0.00s)\n"}
    {"Time":"2026-10-25T21:46:26.238941418-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/output_tokens","Elapsed":0}
    {"Time":"2026-10-25T21:46:26.238943541-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cost_usd"}
    {"Time":"2026-10-25T21:46:26.238944957-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cost_usd","Output":"=== RUN   TestAbsentNumbersCountZero/cost_usd\n"}
    {"Time":"2026-10-25T21:46:26.238988104-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cost_usd","Output":"--- PASS: TestAbsentNumbersCountZero/cost_usd (0.00s)\n"}
    {"Time":"2026-10-25T21:46:26.239000417-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/cost_usd","Elapsed":0}
    {"Time":"2026-10-25T21:46:26.239002393-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/elapsed_ms"}
    {"Time":"2026-10-25T21:46:26.239003886-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/elapsed_ms","Output":"=== RUN   TestAbsentNumbersCountZero/elapsed_ms\n"}
    {"Time":"2026-10-25T21:46:26.239058313-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/elapsed_ms","Output":"--- PASS: TestAbsentNumbersCountZero/elapsed_ms (0.00s)\n"}
    {"Time":"2026-10-25T21:46:26.239070755-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero/elapsed_ms","Elapsed":0}
    {"Time":"2026-10-25T21:46:26.239072982-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero","Output":"--- PASS: TestAbsentNumbersCountZero (0.00s)\n"}
    {"Time":"2026-10-25T21:46:26.23907494-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestAbsentNumbersCountZero","Elapsed":0}
    {"Time":"2026-10-25T21:46:26.239077122-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:26.239773029-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:26.239801161-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: absent-numbers-wrong-type (FEATURE-AbsentNumbers.md)
  intent: A numeric field with a JSON type other than number or null counts as zero and keeps the record.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:27.321648099-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:27.328552744-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero"}
    {"Time":"2026-10-25T21:46:27.328582049-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero","Output":"=== RUN   TestWrongTypedNumbersCountZero\n"}
    {"Time":"2026-10-25T21:46:27.328601703-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/string"}
    {"Time":"2026-10-25T21:46:27.328603452-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/string","Output":"=== RUN   TestWrongTypedNumbersCountZero/string\n"}
    {"Time":"2026-10-25T21:46:27.32870472-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/string","Output":"--- PASS: TestWrongTypedNumbersCountZero/string (0.00s)\n"}
    {"Time":"2026-10-25T21:46:27.328727922-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/string","Elapsed":0}
    {"Time":"2026-10-25T21:46:27.328732115-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/boolean"}
    {"Time":"2026-10-25T21:46:27.328733451-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/boolean","Output":"=== RUN   TestWrongTypedNumbersCountZero/boolean\n"}
    {"Time":"2026-10-25T21:46:27.32873717-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/boolean","Output":"--- PASS: TestWrongTypedNumbersCountZero/boolean (0.00s)\n"}
    {"Time":"2026-10-25T21:46:27.32876304-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/boolean","Elapsed":0}
    {"Time":"2026-10-25T21:46:27.328765225-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/array"}
    {"Time":"2026-10-25T21:46:27.328767209-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/array","Output":"=== RUN   TestWrongTypedNumbersCountZero/array\n"}
    {"Time":"2026-10-25T21:46:27.32883125-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/array","Output":"--- PASS: TestWrongTypedNumbersCountZero/array (0.00s)\n"}
    {"Time":"2026-10-25T21:46:27.328852389-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/array","Elapsed":0}
    {"Time":"2026-10-25T21:46:27.328857152-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/object"}
    {"Time":"2026-10-25T21:46:27.328858621-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/object","Output":"=== RUN   TestWrongTypedNumbersCountZero/object\n"}
    {"Time":"2026-10-25T21:46:27.328907117-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/object","Output":"--- PASS: TestWrongTypedNumbersCountZero/object (0.00s)\n"}
    {"Time":"2026-10-25T21:46:27.328921126-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero/object","Elapsed":0}
    {"Time":"2026-10-25T21:46:27.328923811-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero","Output":"--- PASS: TestWrongTypedNumbersCountZero (0.00s)\n"}
    {"Time":"2026-10-25T21:46:27.328925818-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestWrongTypedNumbersCountZero","Elapsed":0}
    {"Time":"2026-10-25T21:46:27.328927193-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:27.329843515-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:27.32986455-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: absent-numbers-returncode (FEATURE-AbsentNumbers.md)
  intent: A null or absent returncode is recorded as absent rather than zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:28.403639755-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:28.409211751-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReturncodeNullOrAbsentIsAbsent"}
    {"Time":"2026-10-25T21:46:28.409254173-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReturncodeNullOrAbsentIsAbsent","Output":"=== RUN   TestReturncodeNullOrAbsentIsAbsent\n"}
    {"Time":"2026-10-25T21:46:28.40934304-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReturncodeNullOrAbsentIsAbsent","Output":"--- PASS: TestReturncodeNullOrAbsentIsAbsent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:28.409357143-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReturncodeNullOrAbsentIsAbsent","Elapsed":0}
    {"Time":"2026-10-25T21:46:28.409360609-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:28.41033815-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-25T21:46:28.410357416-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: absent-numbers-timed-out (FEATURE-AbsentNumbers.md)
  intent: An absent timed_out defaults to false.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:29.501671604-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:29.508660263-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestTimedOutDefaultsFalse"}
    {"Time":"2026-10-25T21:46:29.508710091-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestTimedOutDefaultsFalse","Output":"=== RUN   TestTimedOutDefaultsFalse\n"}
    {"Time":"2026-10-25T21:46:29.508785417-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestTimedOutDefaultsFalse","Output":"--- PASS: TestTimedOutDefaultsFalse (0.00s)\n"}
    {"Time":"2026-10-25T21:46:29.508790291-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestTimedOutDefaultsFalse","Elapsed":0}
    {"Time":"2026-10-25T21:46:29.508794787-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:29.509335734-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:29.509368894-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-attribution-executed (FEATURE-ModelAttribution.md)
  intent: A record with a non-empty result.stats.model is attributed to that model rather than job.model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:30.587552795-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:30.595569644-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted"}
    {"Time":"2026-10-25T21:46:30.595622152-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted","Output":"=== RUN   TestModelPrefersExecuted\n"}
    {"Time":"2026-10-25T21:46:30.595663867-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted","Output":"--- PASS: TestModelPrefersExecuted (0.00s)\n"}
    {"Time":"2026-10-25T21:46:30.595669941-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted","Elapsed":0}
    {"Time":"2026-10-25T21:46:30.595675361-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:30.596413353-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:30.596436222-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-attribution-no-stats (FEATURE-ModelAttribution.md)
  intent: A record without result.stats is attributed to job.model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:31.656680682-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:31.663784847-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenStatsAbsent"}
    {"Time":"2026-10-25T21:46:31.663840091-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenStatsAbsent","Output":"=== RUN   TestModelFallsBackWhenStatsAbsent\n"}
    {"Time":"2026-10-25T21:46:31.663863096-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenStatsAbsent","Output":"--- PASS: TestModelFallsBackWhenStatsAbsent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:31.663867043-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenStatsAbsent","Elapsed":0}
    {"Time":"2026-10-25T21:46:31.663870165-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:31.664975617-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:31.665001277-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-attribution-empty (FEATURE-ModelAttribution.md)
  intent: A null, absent, or empty result.stats.model falls back to job.model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:32.727272763-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:32.734123627-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty"}
    {"Time":"2026-10-25T21:46:32.734201074-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty","Output":"=== RUN   TestModelFallsBackWhenEmpty\n"}
    {"Time":"2026-10-25T21:46:32.734213438-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/absent"}
    {"Time":"2026-10-25T21:46:32.734216528-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/absent","Output":"=== RUN   TestModelFallsBackWhenEmpty/absent\n"}
    {"Time":"2026-10-25T21:46:32.734225543-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/absent","Output":"--- PASS: TestModelFallsBackWhenEmpty/absent (0.00s)\n"}
    {"Time":"2026-10-25T21:46:32.734229855-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/absent","Elapsed":0}
    {"Time":"2026-10-25T21:46:32.734235203-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/not_a_string"}
    {"Time":"2026-10-25T21:46:32.7342374-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/not_a_string","Output":"=== RUN   TestModelFallsBackWhenEmpty/not_a_string\n"}
    {"Time":"2026-10-25T21:46:32.734240266-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/not_a_string","Output":"--- PASS: TestModelFallsBackWhenEmpty/not_a_string (0.00s)\n"}
    {"Time":"2026-10-25T21:46:32.734242844-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/not_a_string","Elapsed":0}
    {"Time":"2026-10-25T21:46:32.734245687-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/empty_string"}
    {"Time":"2026-10-25T21:46:32.73424787-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/empty_string","Output":"=== RUN   TestModelFallsBackWhenEmpty/empty_string\n"}
    {"Time":"2026-10-25T21:46:32.73426652-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/empty_string","Output":"--- PASS: TestModelFallsBackWhenEmpty/empty_string (0.00s)\n"}
    {"Time":"2026-10-25T21:46:32.734297845-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/empty_string","Elapsed":0}
    {"Time":"2026-10-25T21:46:32.734304637-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/null"}
    {"Time":"2026-10-25T21:46:32.734307509-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/null","Output":"=== RUN   TestModelFallsBackWhenEmpty/null\n"}
    {"Time":"2026-10-25T21:46:32.734314953-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/null","Output":"--- PASS: TestModelFallsBackWhenEmpty/null (0.00s)\n"}
    {"Time":"2026-10-25T21:46:32.734317754-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty/null","Elapsed":0}
    {"Time":"2026-10-25T21:46:32.734320609-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty","Output":"--- PASS: TestModelFallsBackWhenEmpty (0.00s)\n"}
    {"Time":"2026-10-25T21:46:32.734323948-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty","Elapsed":0}
    {"Time":"2026-10-25T21:46:32.734327403-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:32.735032539-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-25T21:46:32.735060468-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-attribution-provider (FEATURE-ModelAttribution.md)
  intent: The provider of a record is read from job.llm.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T21:46:33.885565403-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-25T21:46:33.892528434-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestProviderIsJobLLM"}
    {"Time":"2026-10-25T21:46:33.892560715-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestProviderIsJobLLM","Output":"=== RUN   TestProviderIsJobLLM\n"}
    {"Time":"2026-10-25T21:46:33.89273512-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestProviderIsJobLLM","Output":"--- PASS: TestProviderIsJobLLM (0.00s)\n"}
    {"Time":"2026-10-25T21:46:33.892758271-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestProviderIsJobLLM","Elapsed":0}
    {"Time":"2026-10-25T21:46:33.892762149-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-25T21:46:33.893585038-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-25T21:46:33.893610286-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
Block 3 covers the log reader (`internal/logread`) and the CLI's date validation, unreadable-log and usage-error exits. The report package belongs to a later block, so a readable log still ends in the existing "not implemented" path. I'll start with the CLI.I added an `os` import that the CLI doesn't use, so I'll remove it and then write the logread package.`record.go` uses `bytes` and `math` without importing them, so I'll fix the import block.Next, the logread tests. Each named test from the feature specs goes in, plus a shared helper that builds records as maps so each test can change one field.Now the CLI tests: the three named date tests, the usage-line test, and the read-failure test.Now I'll format, vet, and run the full suite with the project's own test script, then the grader.The one failure is a bad test case on my side: `2026-10-14` on its own does begin with a real date, which is what the spec requires, so it should be kept. I'll remove that case from the skip table.The 20-check grader fails only where a report is requested: the report package belongs to a later block, so those checks cannot pass yet. The parts this block owns pass (E1, U1–U6), and the skip counts on the sample match the expected `skipped` value of 3. I'll confirm each named acceptance test exists and ran.RESULT: SUCCESS

Block 3 (Service) is built. Every declared acceptance test in the log-reading, malformed-line, duplicate, absent-number, model, date, unreadable-log and usage-error specs runs and passes, and `sh bin/test.sh` (vet plus `go test ./...`) is green. Nothing was committed.

The grader `sources/run.sh` is not green yet: 7 of 20 pass, and the 13 failures are all report output. The report package is a later block, so those checks can't pass here. The skip counts on the sample match what the grader expects (3 skipped).

**Choices I made**
- **Readable log, report not implemented:** after a successful read, `ddkstat` still prints "not implemented" and exits 1. That is the existing placeholder, kept until the report block lands.
- **Date validation:** `--since` and `--until` are validated in `Parse` only when given, so `--since ""` is a usage error and only the last occurrence of a repeated option is checked.
- **Testable CLI:** `Run` delegates to `execute`, which takes a loader function, so the partial-read test can feed a reader that fails midway.
- **Warnings:** written as `ddkstat: line N: skipped: <reason>`.

**Files changed**
- `internal/cli/cli.go`: date validation in `Parse`, unreadable-log handling via `execute`.
- `internal/logread/read.go`: `Read`, `ReadFile`, `Result`, skip and duplicate rules.
- `internal/logread/record.go`: `Record`, JSON field decoding, model fallback, numeric fields that count as zero.
- Tests: `internal/logread/{helpers,read,skip,duplicate,absent,model}_test.go`, `internal/cli/{dates,run}_test.go`.

`gofmt` is clean and `go vet` reports nothing.

<reusable-compact filename="ARCHITECTURE.md">
logread: Read(io.Reader, io.Writer) (Result, error); ReadFile(path, io.Writer) (Result, error); Result{Records []Record; Skipped int}; Record{ExecutionID, StartedAt, Status, Command, Target, Provider, Model string; InputTokens, CachedTokens, OutputTokens, ElapsedMS int64; CostUSD float64; ReturnCode *int; TimedOut bool}; Record.Date(), Record.Failed(). Skip when not a JSON object, invalid UTF-8, a required string (execution_id, started_at, status, job.command_name, job.target) absent/empty/non-string, or started_at not starting with a real YYYY-MM-DD. Duplicate execution_id keeps first; only kept records register ids. Numeric fields absent/null/wrong-typed/fractional are 0. Model = result.stats.model if non-empty, else job.model; Provider = job.llm.
cli: Run(args, stdout, stderr) int; exit 0/1/2; --since/--until must be real YYYY-MM-DD when given.
</reusable-compact>
