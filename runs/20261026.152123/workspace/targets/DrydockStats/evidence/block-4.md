# Evidence: Block 4 · Service (block-4)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 5994
- execution id: 20261026.153951.393Z-ca9280c6

## Stories built
- Treat absent, null or non-numeric token, cost and duration values as zero. (log-numbers) [story]
- Attribute each record to the executed model, falling back to the requested model. (log-model) [story]

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-LogNumbers.md (SP 1163)
- context: ARCHITECTURE_compact.md (SP 493)
- implements: FEATURE-LogModel.md (SP 957)

## Build directory changes
- internal/logread/logread.go
- internal/logread/logread_test.go

## Pre-build acceptance observation
- UNVERIFIED: numbers-absent (FEATURE-LogNumbers.md)
  intent: Absent numeric fields count as zero.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:33.925228194-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:33.930661967-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:33.930703016-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:33.931403768-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:33.931426987-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:33.925228194-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:33.930661967-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:33.930703016-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numbers-absent.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:33.925228194-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:33.930661967-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:33.930703016-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:33.931403768-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:33.931426987-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}\n', '')
- UNVERIFIED: numbers-null (FEATURE-LogNumbers.md)
  intent: Null numeric fields, including a null cost_usd, count as zero.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:34.919355776-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:34.923964118-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:34.924012765-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:34.924681604-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.005s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:34.92471597-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:34.919355776-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:34.923964118-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:34.924012765-04:00","Action":"output","Package":"ddkstat/… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numbers-null.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:34.919355776-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:34.923964118-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:34.924012765-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:34.924681604-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:34.92471597-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.005}\n', '')
- UNVERIFIED: numbers-non-numeric (FEATURE-LogNumbers.md)
  intent: A non-numeric value in a numeric stats field counts as zero.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:35.97941796-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:35.985889497-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:35.985931396-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:35.986713462-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:35.986732691-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:35.97941796-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:35.985889497-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:35.985931396-04:00","Action":"output","Package":"ddkstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numbers-non-numeric.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:35.97941796-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:35.985889497-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:35.985931396-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:35.986713462-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:35.986732691-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')
- UNVERIFIED: numbers-no-stats (FEATURE-LogNumbers.md)
  intent: A record with no result.stats object is still a kept record with zero numbers.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:37.130317915-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:37.136500947-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:37.136554389-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:37.137133866-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:37.137154165-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:37.130317915-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:37.136500947-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:37.136554389-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numbers-no-stats.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:37.130317915-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:37.136500947-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:37.136554389-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:37.137133866-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:37.137154165-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')
- UNVERIFIED: numbers-returncode-null (FEATURE-LogNumbers.md)
  intent: A null or absent returncode is represented as no return code.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:38.305189142-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:38.312182877-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:38.312226359-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:38.312899768-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:38.312923135-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:38.305189142-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:38.312182877-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:38.312226359-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numbers-returncode-null.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:38.305189142-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:38.312182877-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:38.312226359-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:38.312899768-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:38.312923135-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}\n', '')
- UNVERIFIED: model-prefers-executed (FEATURE-LogModel.md)
  intent: result.stats.model is used when present and non-empty.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:39.465104773-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:39.471508054-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:39.471550767-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:39.472440376-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:39.472458256-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:39.465104773-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:39.471508054-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:39.471550767-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-prefers-executed.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:39.465104773-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:39.471508054-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:39.471550767-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:39.472440376-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:39.472458256-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')
- UNVERIFIED: model-fallback-absent (FEATURE-LogModel.md)
  intent: job.model is used when result.stats.model is absent.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:40.585169467-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:40.590631712-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:40.590669974-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:40.591365067-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:40.591378025-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:40.585169467-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:40.590631712-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:40.590669974-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-fallback-absent.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:40.585169467-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:40.590631712-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:40.590669974-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:40.591365067-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:40.591378025-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}\n', '')
- UNVERIFIED: model-fallback-empty (FEATURE-LogModel.md)
  intent: job.model is used when result.stats.model is an empty string.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:41.644756583-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:41.650951472-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:41.650997654-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:41.651564204-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:41.651584543-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:41.644756583-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:41.650951472-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:41.650997654-04:00","Action":"output","Package":"ddkstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-fallback-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:41.644756583-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:41.650951472-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:41.650997654-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:41.651564204-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:41.651584543-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')
- UNVERIFIED: model-fallback-no-stats (FEATURE-LogModel.md)
  intent: job.model is used when the record has no result.stats object.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:39:42.802095858-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:39:42.808377566-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:39:42.808412627-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:39:42.808916917-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:39:42.80893579-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:39:42.802095858-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:42.808377566-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:42.808412627-04:00","Action":"output","Package":"ddkstat/… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "model-fallback-no-stats.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:39:42.802095858-04:00","Action":"start","Package":"ddkstat/internal/logread"}\n{"Time":"2026-10-26T11:39:42.808377566-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:39:42.808412627-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:39:42.808916917-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \\tddkstat/internal/logread\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:39:42.80893579-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}\n', '')

## Post-build programmatic acceptance
- PASS: numbers-absent (FEATURE-LogNumbers.md)
  intent: Absent numeric fields count as zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:20.359957086-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:20.366540436-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestNumbersAbsentAreZero"}
    {"Time":"2026-10-26T11:41:20.366590569-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNumbersAbsentAreZero","Output":"=== RUN   TestNumbersAbsentAreZero\n"}
    {"Time":"2026-10-26T11:41:20.366628783-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNumbersAbsentAreZero","Output":"--- PASS: TestNumbersAbsentAreZero (0.00s)\n"}
    {"Time":"2026-10-26T11:41:20.366633906-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestNumbersAbsentAreZero","Elapsed":0}
    {"Time":"2026-10-26T11:41:20.366638077-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:20.367236566-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:41:20.367269756-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numbers-null (FEATURE-LogNumbers.md)
  intent: Null numeric fields, including a null cost_usd, count as zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:21.273032364-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:21.2802881-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestNumbersNullAreZero"}
    {"Time":"2026-10-26T11:41:21.28031714-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNumbersNullAreZero","Output":"=== RUN   TestNumbersNullAreZero\n"}
    {"Time":"2026-10-26T11:41:21.280373773-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNumbersNullAreZero","Output":"--- PASS: TestNumbersNullAreZero (0.00s)\n"}
    {"Time":"2026-10-26T11:41:21.280394887-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestNumbersNullAreZero","Elapsed":0}
    {"Time":"2026-10-26T11:41:21.280398697-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:21.28096127-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:41:21.280982142-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numbers-non-numeric (FEATURE-LogNumbers.md)
  intent: A non-numeric value in a numeric stats field counts as zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:22.176050097-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:22.181979003-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestNumbersNonNumericAreZero"}
    {"Time":"2026-10-26T11:41:22.182008439-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNumbersNonNumericAreZero","Output":"=== RUN   TestNumbersNonNumericAreZero\n"}
    {"Time":"2026-10-26T11:41:22.18205744-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNumbersNonNumericAreZero","Output":"--- PASS: TestNumbersNonNumericAreZero (0.00s)\n"}
    {"Time":"2026-10-26T11:41:22.182072076-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestNumbersNonNumericAreZero","Elapsed":0}
    {"Time":"2026-10-26T11:41:22.182076367-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:22.182754882-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:41:22.182773515-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numbers-no-stats (FEATURE-LogNumbers.md)
  intent: A record with no result.stats object is still a kept record with zero numbers.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:23.072400792-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:23.078442257-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestNoStatsObjectStillRecord"}
    {"Time":"2026-10-26T11:41:23.078486101-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNoStatsObjectStillRecord","Output":"=== RUN   TestNoStatsObjectStillRecord\n"}
    {"Time":"2026-10-26T11:41:23.078515263-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestNoStatsObjectStillRecord","Output":"--- PASS: TestNoStatsObjectStillRecord (0.00s)\n"}
    {"Time":"2026-10-26T11:41:23.078536279-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestNoStatsObjectStillRecord","Elapsed":0}
    {"Time":"2026-10-26T11:41:23.078540862-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:23.079158794-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:41:23.079178997-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numbers-returncode-null (FEATURE-LogNumbers.md)
  intent: A null or absent returncode is represented as no return code.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:23.937559925-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:23.943845602-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestReturnCodeNullIsNil"}
    {"Time":"2026-10-26T11:41:23.943864729-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReturnCodeNullIsNil","Output":"=== RUN   TestReturnCodeNullIsNil\n"}
    {"Time":"2026-10-26T11:41:23.943956792-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestReturnCodeNullIsNil","Output":"--- PASS: TestReturnCodeNullIsNil (0.00s)\n"}
    {"Time":"2026-10-26T11:41:23.943982264-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestReturnCodeNullIsNil","Elapsed":0}
    {"Time":"2026-10-26T11:41:23.943987358-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:23.94462563-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.007s\n"}
    {"Time":"2026-10-26T11:41:23.944636766-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-prefers-executed (FEATURE-LogModel.md)
  intent: result.stats.model is used when present and non-empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:24.667234711-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:24.674238342-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted"}
    {"Time":"2026-10-26T11:41:24.674272777-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted","Output":"=== RUN   TestModelPrefersExecuted\n"}
    {"Time":"2026-10-26T11:41:24.674294877-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted","Output":"--- PASS: TestModelPrefersExecuted (0.00s)\n"}
    {"Time":"2026-10-26T11:41:24.674319042-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelPrefersExecuted","Elapsed":0}
    {"Time":"2026-10-26T11:41:24.674330872-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:24.675161275-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.008s\n"}
    {"Time":"2026-10-26T11:41:24.675183395-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-fallback-absent (FEATURE-LogModel.md)
  intent: job.model is used when result.stats.model is absent.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:25.563438211-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:25.569277844-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenAbsent"}
    {"Time":"2026-10-26T11:41:25.569319101-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenAbsent","Output":"=== RUN   TestModelFallsBackWhenAbsent\n"}
    {"Time":"2026-10-26T11:41:25.569326313-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenAbsent","Output":"--- PASS: TestModelFallsBackWhenAbsent (0.00s)\n"}
    {"Time":"2026-10-26T11:41:25.569328701-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenAbsent","Elapsed":0}
    {"Time":"2026-10-26T11:41:25.569333342-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:25.569994539-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:41:25.570014715-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-fallback-empty (FEATURE-LogModel.md)
  intent: job.model is used when result.stats.model is an empty string.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:26.483794568-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:26.48998489-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty"}
    {"Time":"2026-10-26T11:41:26.490019688-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty","Output":"=== RUN   TestModelFallsBackWhenEmpty\n"}
    {"Time":"2026-10-26T11:41:26.490060296-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty","Output":"--- PASS: TestModelFallsBackWhenEmpty (0.00s)\n"}
    {"Time":"2026-10-26T11:41:26.490065421-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWhenEmpty","Elapsed":0}
    {"Time":"2026-10-26T11:41:26.490069024-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:26.490687766-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:41:26.490704496-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: model-fallback-no-stats (FEATURE-LogModel.md)
  intent: job.model is used when the record has no result.stats object.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:41:27.300741242-04:00","Action":"start","Package":"ddkstat/internal/logread"}
    {"Time":"2026-10-26T11:41:27.306466504-04:00","Action":"run","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWithoutStats"}
    {"Time":"2026-10-26T11:41:27.306494213-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWithoutStats","Output":"=== RUN   TestModelFallsBackWithoutStats\n"}
    {"Time":"2026-10-26T11:41:27.306536223-04:00","Action":"output","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWithoutStats","Output":"--- PASS: TestModelFallsBackWithoutStats (0.00s)\n"}
    {"Time":"2026-10-26T11:41:27.306541265-04:00","Action":"pass","Package":"ddkstat/internal/logread","Test":"TestModelFallsBackWithoutStats","Elapsed":0}
    {"Time":"2026-10-26T11:41:27.306556294-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"PASS\n"}
    {"Time":"2026-10-26T11:41:27.307100516-04:00","Action":"output","Package":"ddkstat/internal/logread","Output":"ok  \tddkstat/internal/logread\t0.006s\n"}
    {"Time":"2026-10-26T11:41:27.307123399-04:00","Action":"pass","Package":"ddkstat/internal/logread","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
Now the existing implementation, which the previous blocks produced.The existing `logread` package has two gaps against the specs. A string or boolean in a numeric field fails the whole line, so the line is skipped instead of counted as zero. And `Record.Model` is only `stats.model`, with no fallback to `job.model`. I'll fix both in `logread.go`, then add the named tests.Now the tests. I'll append them to the existing `logread_test.go`, which needs a Read first.Now I'll run the project's own gate (`bin/test.sh`, which is `go vet` plus `go test`) and check formatting.All gates pass. The sample log has 3 expected skips (check A1), and this change could alter that count if any sample line has a wrong-typed numeric. I'll check that with a throwaway test file that I delete straight after.The `logread` numeric and model rules are implemented, and `sh bin/test.sh` passes (exit 0). `sources/run.sh` is not expected to pass yet, since `cmd/ddkstat` is still a stub until a later block.

Changes in `internal/logread/logread.go`:
- **Numbers:** absent, `null`, or wrong-typed `input_tokens`, `cached_input_tokens`, `output_tokens`, `cost_usd` and `elapsed_ms` now count as zero. Before, a string in one of these fields dropped the whole line.
- **Missing objects:** a missing or non-object `result` or `stats` still yields a kept record with zero numbers.
- **Return code:** `returncode` that is `null` or absent is nil.
- **Model:** `Record.Model` is `stats.model` when it is a non-empty string, otherwise `job.model`. `RequestedModel` still holds `job.model`.
- Required string fields (`execution_id`, `started_at`, `status`, `job.command_name`, `job.target`) still skip the line when they are missing, empty, or the wrong type.

Tests added to `internal/logread/logread_test.go` under the names the specs require: `TestNumbersAbsentAreZero`, `TestNumbersNullAreZero`, `TestNumbersNonNumericAreZero`, `TestNoStatsObjectStillRecord`, `TestReturnCodeNullIsNil`, and the four `TestModel*` tests. I also added a few extra boundary tests. The new tests pass, and so does the existing suite.

I also checked the staged sample log with a temporary test file, which I deleted afterwards. It still gives 33 records and 3 skips, which matches check A1 in `run.sh`.

Choices I made that the specs don't settle:
- A non-integer `returncode` is treated as nil, and a non-boolean `timed_out` is treated as false.
- A non-string `job.llm` or `job.model` becomes an empty string instead of skipping the line.
- Token and elapsed fields must be JSON integers. A value such as `12.0` counts as zero.

I did not commit anything.

RESULT: SUCCESS

FILES CHANGED:
- internal/logread/logread.go
- internal/logread/logread_test.go

SUMMARY:
Implemented log-numbers and log-model in `internal/logread`. Numeric fields are decoded leniently and default to zero. `Record.Model` falls back to `job.model`. Added the named tests for both stories plus boundary cases. `sh bin/test.sh` exits 0. The sample log still yields 33 records and 3 skips. `sources/run.sh` is not run because `cmd/ddkstat` is still a stub, which belongs to a later block.

<blueprint-decisions>
[{"spec":"FEATURE-LogNumbers.md","severity":"Low","subject":"Wrong-typed returncode, timed_out and token values","decision":"A non-integer returncode is treated as nil, a non-boolean timed_out as false, and a non-integer token or elapsed value as zero. A float such as 12.0 in a token field counts as zero. The spec only named absent, null and non-numeric values, so I took the strictest reading. Is that acceptable?"}]
</blueprint-decisions>
