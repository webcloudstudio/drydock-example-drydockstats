# Evidence: Block 5 · Service (block-5)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 5969
- execution id: 20261026.154218.401Z-9a1eb74b

## Stories built
- Filter records by exact job.target match. (filter-target) [story]
- Filter records by inclusive start-date range. (filter-dates) [story]

## Acceptance tooling authorization
- FEATURE-FilterTarget.md#filter-target-exact: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-case-sensitive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-no-match: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-unset: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterDates.md#filter-dates-since-inclusive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterDates.md#filter-dates-until-inclusive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterDates.md#filter-dates-both-bounds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterDates.md#filter-dates-since-after-until: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterDates.md#filter-dates-bounds-unset: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-FilterTarget.md (SP 938)
- context: ARCHITECTURE_compact.md (SP 493)
- implements: FEATURE-FilterDates.md (SP 1150)

## Build directory changes
- internal/filter/filter.go
- internal/filter/filter_test.go

## Pre-build acceptance observation
- RED: filter-target-exact (FEATURE-FilterTarget.md)
  intent: Only records whose Target equals the name are kept.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-exact.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:04.883333822-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:04.883380871-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:04.883388122-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-target-case-sensitive (FEATURE-FilterTarget.md)
  intent: Target matching is case-sensitive.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-case-sensitive.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:05.608981815-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:05.609024706-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:05.609032473-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-target-no-match (FEATURE-FilterTarget.md)
  intent: A name matching no record yields an empty result.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+439 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-no-match.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:06.33001122-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:06.33005992-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:06.330071727-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-target-unset (FEATURE-FilterTarget.md)
  intent: An empty name leaves every record in place.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-unset.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:07.043281457-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:07.043318225-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:07.043325387-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-dates-since-inclusive (FEATURE-FilterDates.md)
  intent: A record dated exactly on the since bound is kept and earlier ones are dropped.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+440 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-dates-since-inclusive.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:07.790698563-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:07.790748457-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:07.79075642-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-dates-until-inclusive (FEATURE-FilterDates.md)
  intent: A record dated exactly on the until bound is kept and later ones are dropped.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+438 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-dates-until-inclusive.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:08.405007-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:08.405048524-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:08.405056864-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-dates-both-bounds (FEATURE-FilterDates.md)
  intent: With both bounds set only records inside the inclusive range remain.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-dates-both-bounds.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:09.104823401-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:09.104866153-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:09.104877315-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-dates-since-after-until (FEATURE-FilterDates.md)
  intent: A since later than until yields an empty result without error.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-dates-since-after-until.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:09.845690309-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:09.845731394-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:09.845739034-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')
- RED: filter-dates-bounds-unset (FEATURE-FilterDates.md)
  intent: Empty bounds leave every record in place.
  return code: 1
  error: native acceptance: execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-dates-bounds-unset.py", line 4, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/filter","Action":"build-output","Output":"# ./internal/filter\\n"}\n{"ImportPath":"./internal/filter","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/filter: directory not found\\n"}\n{"ImportPath":"./internal/filter","Action":"build-fail"}\n{"Time":"2026-10-26T11:42:10.535484071-04:00","Action":"start","Package":"./internal/filter"}\n{"Time":"2026-10-26T11:42:10.535522091-04:00","Action":"output","Package":"./internal/filter","Output":"FAIL\\t./internal/filter [setup failed]\\n"}\n{"Time":"2026-10-26T11:42:10.535533464-04:00","Action":"fail","Package":"./internal/filter","Elapsed":0,"FailedBuild":"./internal/filter"}\n', '')

## Post-build programmatic acceptance
- PASS: filter-target-exact (FEATURE-FilterTarget.md)
  intent: Only records whose Target equals the name are kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-target-case-sensitive (FEATURE-FilterTarget.md)
  intent: Target matching is case-sensitive.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-target-no-match (FEATURE-FilterTarget.md)
  intent: A name matching no record yields an empty result.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-target-unset (FEATURE-FilterTarget.md)
  intent: An empty name leaves every record in place.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-dates-since-inclusive (FEATURE-FilterDates.md)
  intent: A record dated exactly on the since bound is kept and earlier ones are dropped.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-dates-until-inclusive (FEATURE-FilterDates.md)
  intent: A record dated exactly on the until bound is kept and later ones are dropped.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-dates-both-bounds (FEATURE-FilterDates.md)
  intent: With both bounds set only records inside the inclusive range remain.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-dates-since-after-until (FEATURE-FilterDates.md)
  intent: A since later than until yields an empty result without error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: filter-dates-bounds-unset (FEATURE-FilterDates.md)
  intent: Empty bounds leave every record in place.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}

## Build summary
The `internal/filter` package doesn't exist yet, so I'll create it with the two pure functions and the nine named tests. Checking the Go toolchain first.RESULT: SUCCESS

FILES CHANGED:
- internal/filter/filter.go
- internal/filter/filter_test.go

SUMMARY:
Added `internal/filter` with two pure functions, `ByTarget` and `ByDateRange`, matching the FEATURE-FilterTarget and FEATURE-FilterDates specs. Neither does I/O or modifies its input. Both return a non-nil slice, and an empty slice on no match or when `since` is after `until`.

- `ByTarget` keeps records whose `Target` equals the name exactly, case-sensitive. An empty name keeps everything.
- `ByDateRange` keeps records whose `Date` is within `[since, until]`, inclusive, compared as strings. An empty bound is unset.

Tests: the nine named tests required by the specs (`TestByTargetExact`, `TestByTargetCaseSensitive`, `TestByTargetNoMatch`, `TestByTargetUnset`, `TestSinceInclusive`, `TestUntilInclusive`, `TestRangeBothBounds`, `TestSinceAfterUntilEmpty`, `TestBoundsUnset`), plus three extra: empty input, and input-not-modified for each function. Each test reads back the returned ExecutionIDs and compares them to the expected list.

Checks run:
- `sh bin/test.sh` (`go vet ./...` and `go test ./...`): all four packages report `ok`.
- `go test -race -count=1 -v ./internal/filter/`: all 12 tests pass.

Note: `sources/run.sh` was not run. It needs the CLI wiring and report writers from later blocks, and the spec reserves it for the terminal story. No Git operations were performed.
