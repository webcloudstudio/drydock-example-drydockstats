# Evidence: Block 8 · Service (block-8)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 4734
- execution id: 20261026.154815.733Z-f0d51a42

## Stories built
- Provide shared minutes and cost text formatters. (value-format) [story]

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-ValueFormat.md (SP 972)
- context: ARCHITECTURE_compact.md (SP 493)

## Build directory changes
- internal/format/format.go
- internal/format/format_test.go

## Pre-build acceptance observation
- RED: format-minutes-rounding (FEATURE-ValueFormat.md)
  intent: Minutes applies the (ms + 3000) / 6000 tenths rule across the rounding boundaries.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\n"}
    {"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\n"}
    {"ImportPath":"./internal/format","Action":"build-fail"}
    {"Time":"2026-10-26T11:48:08.033479036-04:00","Action":"start","Package":"./internal/format"}
    {"Time":"2026-10-26T11:48:08.033537412-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\t./internal/format [setup failed]\n"}
    {"Time":"2026-10-26T11:48:08.033546764-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "format-minutes-rounding.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal/format","Action":"build-fail"}\n{"Time":"2026-10-26T11:48:08.033479036-04:00","Action":"start","Package":"./internal/format"}\n{"Time":"2026-10-26T11:48:08.033537412-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\\t./internal/format [setup failed]\\n"}\n{"Time":"2026-10-26T11:48:08.033546764-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}\n', '')
- RED: format-minutes-zero (FEATURE-ValueFormat.md)
  intent: Minutes of zero milliseconds is written as tenths with one decimal digit.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\n"}
    {"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\n"}
    {"ImportPath":"./internal/format","Action":"build-fail"}
    {"Time":"2026-10-26T11:48:08.813983081-04:00","Action":"start","Package":"./internal/format"}
    {"Time":"2026-10-26T11:48:08.814029786-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\t./internal/format [setup failed]\n"}
    {"Time":"2026-10-26T11:48:08.814038798-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "format-minutes-zero.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal/format","Action":"build-fail"}\n{"Time":"2026-10-26T11:48:08.813983081-04:00","Action":"start","Package":"./internal/format"}\n{"Time":"2026-10-26T11:48:08.814029786-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\\t./internal/format [setup failed]\\n"}\n{"Time":"2026-10-26T11:48:08.814038798-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}\n', '')
- RED: format-cost-four-places (FEATURE-ValueFormat.md)
  intent: Cost always writes exactly four decimal places.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\n"}
    {"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\n"}
    {"ImportPath":"./internal/format","Action":"build-fail"}
    {"Time":"2026-10-26T11:48:09.54382739-04:00","Action":"start","Package":"./internal/format"}
    {"Time":"2026-10-26T11:48:09.543906574-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\t./internal/format [setup failed]\n"}
    {"Time":"2026-10-26T11:48:09.54391625-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal… (+439 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "format-cost-four-places.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal/format","Action":"build-fail"}\n{"Time":"2026-10-26T11:48:09.54382739-04:00","Action":"start","Package":"./internal/format"}\n{"Time":"2026-10-26T11:48:09.543906574-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\\t./internal/format [setup failed]\\n"}\n{"Time":"2026-10-26T11:48:09.54391625-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}\n', '')
- RED: format-cost-zero (FEATURE-ValueFormat.md)
  intent: Cost of zero is written with four decimal places.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\n"}
    {"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\n"}
    {"ImportPath":"./internal/format","Action":"build-fail"}
    {"Time":"2026-10-26T11:48:10.242962781-04:00","Action":"start","Package":"./internal/format"}
    {"Time":"2026-10-26T11:48:10.243014991-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\t./internal/format [setup failed]\n"}
    {"Time":"2026-10-26T11:48:10.243024636-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "format-cost-zero.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/format","Action":"build-output","Output":"# ./internal/format\\n"}\n{"ImportPath":"./internal/format","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/format: directory not found\\n"}\n{"ImportPath":"./internal/format","Action":"build-fail"}\n{"Time":"2026-10-26T11:48:10.242962781-04:00","Action":"start","Package":"./internal/format"}\n{"Time":"2026-10-26T11:48:10.243014991-04:00","Action":"output","Package":"./internal/format","Output":"FAIL\\t./internal/format [setup failed]\\n"}\n{"Time":"2026-10-26T11:48:10.243024636-04:00","Action":"fail","Package":"./internal/format","Elapsed":0,"FailedBuild":"./internal/format"}\n', '')

## Post-build programmatic acceptance
- PASS: format-minutes-rounding (FEATURE-ValueFormat.md)
  intent: Minutes applies the (ms + 3000) / 6000 tenths rule across the rounding boundaries.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:48:49.562972882-04:00","Action":"start","Package":"ddkstat/internal/format"}
    {"Time":"2026-10-26T11:48:49.570218485-04:00","Action":"run","Package":"ddkstat/internal/format","Test":"TestMinutesRounding"}
    {"Time":"2026-10-26T11:48:49.570240485-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestMinutesRounding","Output":"=== RUN   TestMinutesRounding\n"}
    {"Time":"2026-10-26T11:48:49.570257922-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestMinutesRounding","Output":"--- PASS: TestMinutesRounding (0.00s)\n"}
    {"Time":"2026-10-26T11:48:49.57026112-04:00","Action":"pass","Package":"ddkstat/internal/format","Test":"TestMinutesRounding","Elapsed":0}
    {"Time":"2026-10-26T11:48:49.570263946-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"PASS\n"}
    {"Time":"2026-10-26T11:48:49.570896651-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"ok  \tddkstat/internal/format\t0.008s\n"}
    {"Time":"2026-10-26T11:48:49.570918756-04:00","Action":"pass","Package":"ddkstat/internal/format","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: format-minutes-zero (FEATURE-ValueFormat.md)
  intent: Minutes of zero milliseconds is written as tenths with one decimal digit.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:48:50.49192818-04:00","Action":"start","Package":"ddkstat/internal/format"}
    {"Time":"2026-10-26T11:48:50.497822287-04:00","Action":"run","Package":"ddkstat/internal/format","Test":"TestMinutesZero"}
    {"Time":"2026-10-26T11:48:50.497863269-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestMinutesZero","Output":"=== RUN   TestMinutesZero\n"}
    {"Time":"2026-10-26T11:48:50.497887652-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestMinutesZero","Output":"--- PASS: TestMinutesZero (0.00s)\n"}
    {"Time":"2026-10-26T11:48:50.497889367-04:00","Action":"pass","Package":"ddkstat/internal/format","Test":"TestMinutesZero","Elapsed":0}
    {"Time":"2026-10-26T11:48:50.497892777-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"PASS\n"}
    {"Time":"2026-10-26T11:48:50.49837503-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"ok  \tddkstat/internal/format\t0.006s\n"}
    {"Time":"2026-10-26T11:48:50.498400071-04:00","Action":"pass","Package":"ddkstat/internal/format","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: format-cost-four-places (FEATURE-ValueFormat.md)
  intent: Cost always writes exactly four decimal places.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:48:51.590335245-04:00","Action":"start","Package":"ddkstat/internal/format"}
    {"Time":"2026-10-26T11:48:51.594757294-04:00","Action":"run","Package":"ddkstat/internal/format","Test":"TestCostFourPlaces"}
    {"Time":"2026-10-26T11:48:51.594788558-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestCostFourPlaces","Output":"=== RUN   TestCostFourPlaces\n"}
    {"Time":"2026-10-26T11:48:51.594809016-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestCostFourPlaces","Output":"--- PASS: TestCostFourPlaces (0.00s)\n"}
    {"Time":"2026-10-26T11:48:51.594811133-04:00","Action":"pass","Package":"ddkstat/internal/format","Test":"TestCostFourPlaces","Elapsed":0}
    {"Time":"2026-10-26T11:48:51.594814661-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"PASS\n"}
    {"Time":"2026-10-26T11:48:51.595335097-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"ok  \tddkstat/internal/format\t0.005s\n"}
    {"Time":"2026-10-26T11:48:51.595360033-04:00","Action":"pass","Package":"ddkstat/internal/format","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: format-cost-zero (FEATURE-ValueFormat.md)
  intent: Cost of zero is written with four decimal places.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:48:52.576779473-04:00","Action":"start","Package":"ddkstat/internal/format"}
    {"Time":"2026-10-26T11:48:52.582915839-04:00","Action":"run","Package":"ddkstat/internal/format","Test":"TestCostZero"}
    {"Time":"2026-10-26T11:48:52.582937802-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestCostZero","Output":"=== RUN   TestCostZero\n"}
    {"Time":"2026-10-26T11:48:52.58295643-04:00","Action":"output","Package":"ddkstat/internal/format","Test":"TestCostZero","Output":"--- PASS: TestCostZero (0.00s)\n"}
    {"Time":"2026-10-26T11:48:52.582959655-04:00","Action":"pass","Package":"ddkstat/internal/format","Test":"TestCostZero","Elapsed":0}
    {"Time":"2026-10-26T11:48:52.582964001-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"PASS\n"}
    {"Time":"2026-10-26T11:48:52.583540474-04:00","Action":"output","Package":"ddkstat/internal/format","Output":"ok  \tddkstat/internal/format\t0.007s\n"}
    {"Time":"2026-10-26T11:48:52.583561881-04:00","Action":"pass","Package":"ddkstat/internal/format","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- internal/format/format.go
- internal/format/format_test.go

SUMMARY:
Added the `internal/format` package with the two pure formatters from FEATURE-ValueFormat.md:
- `Minutes(ms int64) string` computes `t = (ms + 3000) / 6000` in integer division and returns `t/10`, `.`, `t%10`. So 95000 gives `1.6` and 0 gives `0.0`.
- `Cost(usd float64) string` returns `fmt.Sprintf("%.4f", usd)`.

Tests in `format_test.go`, named as the four declared acceptance checks:
- `TestMinutesRounding` checks boundary values (0, 2999, 3000, 5999, 6000, 8999, 9000, 95000, and others). Expected values are computed from the formula in the test, plus spot checks for 95000, 2999 and 3000.
- `TestMinutesZero` checks `Minutes(0) == "0.0"`.
- `TestCostFourPlaces` checks that `Cost` matches `%.4f` and that the fractional part is exactly four digits, read back from the returned string.
- `TestCostZero` checks `Cost(0) == "0.0000"`.

Checks run:
- `gofmt -l internal/format` printed nothing.
- `go vet ./internal/format` passed.
- `go test -count=1 -v ./internal/format`: all four tests PASS.
- `sh bin/test.sh` (vet and test on all packages): all packages `ok`, no failures. The script's own exit status was not captured because I piped it through `tail`; the output shows no failing package.

No changes were made outside `internal/format/`, and nothing under `sources/` was touched. No Git operations were run.

BLOCKERS:
- None.
