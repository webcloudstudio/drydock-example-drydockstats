# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 6888
- execution id: 20261026.153728.666Z-8cc11b51

## Stories built
- Skip malformed, incomplete and duplicate lines with a stderr warning and no effect on the exit code. (log-skip) [story]
- Validate --since and --until as real YYYY-MM-DD calendar dates. (cli-dates) [story]

## Acceptance tooling authorization
- FEATURE-CliDates.md#cli-dates-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliDates.md#cli-dates-invalid-month: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliDates.md#cli-dates-invalid-day: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliDates.md#cli-dates-wrong-shape: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliDates.md#cli-dates-missing-value: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-LogSkip.md (SP 1748)
- context: ARCHITECTURE_compact.md (SP 493)
- implements: FEATURE-CliDates.md (SP 1110)

## Build directory changes
- internal/cli/args.go
- internal/cli/args_test.go
- internal/logread/logread.go
- internal/logread/logread_test.go

## Pre-build acceptance observation
- UNVERIFIED: skip-not-object (FEATURE-LogSkip.md)
  intent: A line that is not a JSON object is skipped and counted.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:07.671129377-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:07.677010377-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:07.677043496-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:07.677558862-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:07.677581327-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:07.671129377-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:07.677010377-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:07.677043496-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-not-object.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:07.671129377-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:07.677010377-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:07.677043496-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:07.677558862-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:07.677581327-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}\n', '')
- UNVERIFIED: skip-missing-field (FEATURE-LogSkip.md)
  intent: A line lacking a required field is skipped and counted.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:08.659463967-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:08.665774596-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:08.665830037-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:08.666391124-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:08.666411434-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:08.659463967-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:08.665774596-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:08.665830037-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-missing-field.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:08.659463967-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:08.665774596-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:08.665830037-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:08.666391124-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:08.666411434-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')
- UNVERIFIED: skip-empty-string-field (FEATURE-LogSkip.md)
  intent: A required field holding an empty or non-string value causes the line to be skipped.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:09.677093626-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:09.683281297-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:09.683321205-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:09.683898472-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:09.683923984-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:09.677093626-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:09.683281297-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:09.683321205-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-empty-string-field.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:09.677093626-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:09.683281297-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:09.683321205-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:09.683898472-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:09.683923984-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')
- UNVERIFIED: skip-duplicate-first-kept (FEATURE-LogSkip.md)
  intent: A repeated execution_id is skipped and the first occurrence is kept.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:10.677423892-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:10.68240495-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:10.682447262-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:10.683023249-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.005s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:10.68304458-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:10.677423892-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:10.68240495-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:10.682447262-04:00","Action":"output","Package":"ddkstat/i… (+353 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-duplicate-first-kept.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:10.677423892-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:10.68240495-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:10.682447262-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:10.683023249-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:10.68304458-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}\n', '')
- UNVERIFIED: skip-malformed-no-register (FEATURE-LogSkip.md)
  intent: A malformed line does not register its execution_id, so a later valid line with that id is kept.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:11.713029018-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:11.718365618-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:11.718406843-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:11.718896716-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:11.71891489-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:11.713029018-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:11.718365618-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:11.718406843-04:00","Action":"output","Package":"ddkstat/… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-malformed-no-register.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:11.713029018-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:11.718365618-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:11.718406843-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:11.718896716-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:11.71891489-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}\n', '')
- UNVERIFIED: skip-warning-line-number (FEATURE-LogSkip.md)
  intent: Each warning names the 1-based physical line number, counting blank lines.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:12.81375101-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:12.819322656-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:12.819371055-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:12.820010885-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:12.820031614-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:12.81375101-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:12.819322656-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:12.819371055-04:00","Action":"output","Package":"ddkstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-warning-line-number.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:12.81375101-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:12.819322656-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:12.819371055-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:12.820010885-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:12.820031614-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}\n', '')
- UNVERIFIED: skip-dedupe-whole-file (FEATURE-LogSkip.md)
  intent: Duplicate detection covers the whole file and does not depend on any later filter.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:37:13.881793118-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:37:13.88612831-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:37:13.886174306-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:37:13.886784753-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.005s [no tests to run]\n"}
    {"Time":"2026-10-26T11:37:13.88680521-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:13.881793118-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:13.88612831-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:13.886174306-04:00","Action":"output","Package":"ddkstat/i… (+353 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skip-dedupe-whole-file.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:13.881793118-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:37:13.88612831-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:13.886174306-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:13.886784753-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:13.88680521-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.005}\n', '')
- UNVERIFIED: cli-dates-valid (FEATURE-CliDates.md)
  intent: A real calendar date is accepted and stored unchanged.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:14.978863947-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:14.983776377-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:14.983814564-04:00","Action":"output","Package":"ddkstat/internal… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-dates-valid.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:14.978863947-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:14.983776377-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:14.983814564-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:14.984322609-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:14.98434166-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.005}\n', '')
- UNVERIFIED: cli-dates-invalid-month (FEATURE-CliDates.md)
  intent: A date with an impossible month is a usage error.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:15.985586655-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:15.990658142-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:15.990712332-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-dates-invalid-month.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:15.985586655-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:15.990658142-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:15.990712332-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:15.991200849-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:15.991225872-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.006}\n', '')
- UNVERIFIED: cli-dates-invalid-day (FEATURE-CliDates.md)
  intent: A date with an impossible day of month is a usage error.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:16.962199794-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:16.967786681-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:16.967828441-04:00","Action":"output","Package":"ddkstat/internal… (+329 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-dates-invalid-day.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:16.962199794-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:16.967786681-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:16.967828441-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:16.96835041-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:16.96836462-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.006}\n', '')
- UNVERIFIED: cli-dates-wrong-shape (FEATURE-CliDates.md)
  intent: A date not written as zero-padded YYYY-MM-DD is a usage error.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:17.956094048-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:17.961847612-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:17.961895466-04:00","Action":"output","Package":"ddkstat/internal… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-dates-wrong-shape.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:17.956094048-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:17.961847612-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:17.961895466-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:17.962396119-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:17.96241959-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.006}\n', '')
- UNVERIFIED: cli-dates-missing-value (FEATURE-CliDates.md)
  intent: A date option with no value is a usage error.
  return code: 1
  error: unverified acceptance: not-run
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:37:19.034146631-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:19.03886161-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:19.038895744-04:00","Action":"output","Package":"ddkstat/internal/… (+330 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-dates-missing-value.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:37:19.034146631-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:37:19.03886161-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:37:19.038895744-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:37:19.039480023-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:37:19.039513454-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.005}\n', '')

## Post-build programmatic acceptance
- PASS: skip-not-object (FEATURE-LogSkip.md)
  intent: A line that is not a JSON object is skipped and counted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:49.817144012-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:49.824794353-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipNotObject"}
    {"Time":"2026-10-26T11:38:49.824815865-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotObject","Output":"=== RUN   TestSkipNotObject\n"}
    {"Time":"2026-10-26T11:38:49.824891363-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipNotObject","Output":"--- PASS: TestSkipNotObject (0.00s)\n"}
    {"Time":"2026-10-26T11:38:49.824910279-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipNotObject","Elapsed":0}
    {"Time":"2026-10-26T11:38:49.82491385-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:49.825598169-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-26T11:38:49.825614888-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: skip-missing-field (FEATURE-LogSkip.md)
  intent: A line lacking a required field is skipped and counted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:51.039057907-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:51.045451666-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipMissingField"}
    {"Time":"2026-10-26T11:38:51.045474104-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingField","Output":"=== RUN   TestSkipMissingField\n"}
    {"Time":"2026-10-26T11:38:51.045564665-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipMissingField","Output":"--- PASS: TestSkipMissingField (0.00s)\n"}
    {"Time":"2026-10-26T11:38:51.04557848-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipMissingField","Elapsed":0}
    {"Time":"2026-10-26T11:38:51.045581386-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:51.046340101-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:38:51.046354335-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: skip-empty-string-field (FEATURE-LogSkip.md)
  intent: A required field holding an empty or non-string value causes the line to be skipped.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:51.962445485-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:51.968569535-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipEmptyStringField"}
    {"Time":"2026-10-26T11:38:51.968595953-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipEmptyStringField","Output":"=== RUN   TestSkipEmptyStringField\n"}
    {"Time":"2026-10-26T11:38:51.968735148-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipEmptyStringField","Output":"--- PASS: TestSkipEmptyStringField (0.00s)\n"}
    {"Time":"2026-10-26T11:38:51.968759903-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipEmptyStringField","Elapsed":0}
    {"Time":"2026-10-26T11:38:51.968764843-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:51.969188728-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:38:51.969203911-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: skip-duplicate-first-kept (FEATURE-LogSkip.md)
  intent: A repeated execution_id is skipped and the first occurrence is kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:53.114796854-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:53.120788964-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipDuplicateKeepsFirst"}
    {"Time":"2026-10-26T11:38:53.120819552-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipDuplicateKeepsFirst","Output":"=== RUN   TestSkipDuplicateKeepsFirst\n"}
    {"Time":"2026-10-26T11:38:53.12088362-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipDuplicateKeepsFirst","Output":"--- PASS: TestSkipDuplicateKeepsFirst (0.00s)\n"}
    {"Time":"2026-10-26T11:38:53.120907136-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipDuplicateKeepsFirst","Elapsed":0}
    {"Time":"2026-10-26T11:38:53.120912205-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:53.121632495-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:38:53.121658339-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: skip-malformed-no-register (FEATURE-LogSkip.md)
  intent: A malformed line does not register its execution_id, so a later valid line with that id is kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:54.213424751-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:54.220437547-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestMalformedDoesNotRegisterID"}
    {"Time":"2026-10-26T11:38:54.220467536-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestMalformedDoesNotRegisterID","Output":"=== RUN   TestMalformedDoesNotRegisterID\n"}
    {"Time":"2026-10-26T11:38:54.220552104-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestMalformedDoesNotRegisterID","Output":"--- PASS: TestMalformedDoesNotRegisterID (0.00s)\n"}
    {"Time":"2026-10-26T11:38:54.220564746-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestMalformedDoesNotRegisterID","Elapsed":0}
    {"Time":"2026-10-26T11:38:54.220569434-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:54.2213244-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-26T11:38:54.221345534-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: skip-warning-line-number (FEATURE-LogSkip.md)
  intent: Each warning names the 1-based physical line number, counting blank lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:55.314544732-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:55.321192459-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestSkipWarningLineNumber"}
    {"Time":"2026-10-26T11:38:55.321215183-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipWarningLineNumber","Output":"=== RUN   TestSkipWarningLineNumber\n"}
    {"Time":"2026-10-26T11:38:55.321317079-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestSkipWarningLineNumber","Output":"--- PASS: TestSkipWarningLineNumber (0.00s)\n"}
    {"Time":"2026-10-26T11:38:55.321350518-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestSkipWarningLineNumber","Elapsed":0}
    {"Time":"2026-10-26T11:38:55.321355803-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:55.321900071-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:38:55.321918299-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: skip-dedupe-whole-file (FEATURE-LogSkip.md)
  intent: Duplicate detection covers the whole file and does not depend on any later filter.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:38:56.332457734-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:38:56.338606165-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReadDedupesWholeFile"}
    {"Time":"2026-10-26T11:38:56.338642139-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadDedupesWholeFile","Output":"=== RUN   TestReadDedupesWholeFile\n"}
    {"Time":"2026-10-26T11:38:56.344147631-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReadDedupesWholeFile","Output":"--- PASS: TestReadDedupesWholeFile (0.01s)\n"}
    {"Time":"2026-10-26T11:38:56.344166295-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReadDedupesWholeFile","Elapsed":0.01}
    {"Time":"2026-10-26T11:38:56.344171353-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:38:56.345173989-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.012s\n"}
    {"Time":"2026-10-26T11:38:56.345192764-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.013}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-dates-valid (FEATURE-CliDates.md)
  intent: A real calendar date is accepted and stored unchanged.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-dates-invalid-month (FEATURE-CliDates.md)
  intent: A date with an impossible month is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-dates-invalid-day (FEATURE-CliDates.md)
  intent: A date with an impossible day of month is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-dates-wrong-shape (FEATURE-CliDates.md)
  intent: A date not written as zero-padded YYYY-MM-DD is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: cli-dates-missing-value (FEATURE-CliDates.md)
  intent: A date option with no value is a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
The build block covers two stories. `logread` already skips non-objects and invalid JSON, but it has no required-field check, no duplicate detection, and no structured record of skip line numbers. `cli` has no date validation. I'll start with the reader.Now the reader tests. Each one reads back the returned `Result` (records, counts, skip line numbers) rather than matching warning text.The reader tests pass. The notice says the test file changed on disk; that's my own append plus `gofmt -w`. I'll confirm the content before moving on.The reader side is in place. Next, date validation in `ParseArgs`.Now the five date tests named in the spec, appended to the existing CLI test file.Both stories are implemented, and the new tests pass: the full `bin/test.sh` (vet plus all packages) exits 0.

I didn't run `sources/run.sh`. It's the terminal story, and the spec says to run it once, last. It would also fail here because `cmd/ddkstat/main.go` is still an empty stub (later blocks wire that up).

**What changed**

- **log-skip** (`internal/logread/logread.go`): a line is now skipped when a required field is absent, empty, or not a string. Duplicate `execution_id`s are skipped across the whole file, keeping the first occurrence. Only kept lines register their ID, so a malformed line never blocks a later valid one. Each skip still gets one warning line on the warning writer.
- **cli-dates** (`internal/cli/args.go`): `--since` and `--until` values must round-trip through `time.Parse("2006-01-02")`. Impossible dates, wrong shapes, and empty values are `*UsageError`s. A missing value was already a usage error.
- **Tests**: I added the seven required logread tests and the five required cli date tests, named as the spec names them. Each one reads back the returned `Result` or `Options` rather than matching warning text.

**Choices you may want to redirect**

- I added `Result.Skips []Skip` (line number and reason) alongside `Result.Skipped`. The line-number test checks that structured field, not stderr text, so the oracle is a return value. ARCHITECTURE.md doesn't mention this field, so it's an additive change to the reader's public shape.
- A line whose non-required field has the wrong JSON type (for example `returncode` as a string) is still skipped as a whole line. That was already the behavior, and the spec doesn't ask to change it.

<blueprint-decisions>
[{"spec":"FEATURE-LogSkip.md","severity":"Low","subject":"Structured skip record","decision":"Warnings were required on the warning writer only. I also added Result.Skips, a slice of {Line, Reason}, so the warning line-number test can read the line back as a value instead of matching text. Skipped is still the count. Is this acceptable, or should it move out of Result on replan?"}]
</blueprint-decisions>

RESULT: SUCCESS

FILES CHANGED:
- internal/logread/logread.go
- internal/logread/logread_test.go
- internal/cli/args.go
- internal/cli/args_test.go

SUMMARY:
Implemented the log-skip and cli-dates stories. The reader skips non-objects, incomplete or empty required fields, and duplicate execution_ids across the whole file, with one warning per skip and no effect on the exit code. A malformed line does not register its ID. ParseArgs rejects --since/--until values that are not real YYYY-MM-DD calendar dates. Checks run: `go vet ./...` and `go test ./...` via `sh bin/test.sh` both pass (exit 0). All twelve required test names (seven log-skip, five cli-dates) run and pass. `sources/run.sh` was not run; it is the terminal story and main.go is still a stub.

BLOCKERS:
- None.
