# Evidence: Block 9 · Service (block-9)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 15632
- execution id: 20261026.155116.930Z-7029aa95

## Stories built
- Write the one-row summary report. (report-summary) [story]
- Write the targets report with cost and a TOTAL row. (report-targets) [story]
- Write the commands report with a TOTAL row. (report-commands) [story]
- Write the models report grouped by provider and model with a TOTAL row. (report-models) [story]
- Write the failures report listing each failed record. (report-failures) [story]
- Write the daily report with token totals and a TOTAL row. (report-daily) [story]

## Acceptance tooling authorization
- FEATURE-ReportFailures.md#report-failures-unit: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportFailures.md#report-failures-rows: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportFailures.md#report-failures-filters: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportFailures.md#report-failures-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportDaily.md#report-daily-unit: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportDaily.md#report-daily-rows: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportDaily.md#report-daily-filters: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportDaily.md#report-daily-empty: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-ReportSummary.md (SP 1338)
- context: ARCHITECTURE_compact.md (SP 493)
- implements: FEATURE-ReportTargets.md (SP 1133)
- implements: FEATURE-ReportCommands.md (SP 1165)
- implements: FEATURE-ReportModels.md (SP 1376)
- implements: FEATURE-ReportFailures.md (SP 2644)
- implements: FEATURE-ReportDaily.md (SP 3319)

## Build directory changes
- cmd/ddkstat/main.go
- internal/cli/reports_test.go
- internal/cli/run_test.go
- internal/report/commands.go
- internal/report/commands_test.go
- internal/report/daily.go
- internal/report/daily_test.go
- internal/report/failures.go
- internal/report/failures_test.go
- internal/report/helpers_test.go
- internal/report/models.go
- internal/report/models_test.go
- internal/report/report.go
- internal/report/summary.go
- internal/report/summary_test.go
- internal/report/targets.go
- internal/report/targets_test.go

## Pre-build acceptance observation
- RED: summary-aggregate (FEATURE-ReportSummary.md)
  intent: SummaryAggregate sums counts, tokens, cost and duration over records built by the test.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:05.227056184-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:05.227141439-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:05.227159886-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-aggregate.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:05.227056184-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:05.227141439-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:05.227159886-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: summary-empty-first-last (FEATURE-ReportSummary.md)
  intent: With no matching record, first and last are empty and counters are zero.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:06.00560554-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:06.005647919-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:06.005655841-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+440 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-empty-first-last.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:06.00560554-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:06.005647919-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:06.005655841-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: summary-skipped-ignores-filters (FEATURE-ReportSummary.md)
  intent: The skipped count stays the whole-file count regardless of target and date filters.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:06.874551176-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:06.874622734-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:06.874638863-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-skipped-ignores-filters.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:06.874551176-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:06.874622734-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:06.874638863-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: summary-write (FEATURE-ReportSummary.md)
  intent: WriteSummary writes the declared header and one row with the declared column order.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:07.811317138-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:07.81136218-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:07.811370134-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+440 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-write.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:07.811317138-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:07.81136218-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:07.811370134-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- UNVERIFIED: summary-run-via-cli (FEATURE-ReportSummary.md)
  intent: cli.Run dispatches the summary report over a test-written log and exits zero with a parseable CSV.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:50:09.055684918-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:50:09.061803915-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:50:09.061849696-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:50:09.062338838-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:50:09.062355191-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:50:09.055684918-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:09.061803915-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:09.061849696-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-run-via-cli.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:50:09.055684918-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:09.061803915-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:09.061849696-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:50:09.062338838-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:50:09.062355191-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}\n', '')
- RED: targets-ordering (FEATURE-ReportTargets.md)
  intent: Rows are ordered by target ascending in byte order.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:09.968331278-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:09.968380676-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:09.968388277-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-ordering.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:09.968331278-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:09.968380676-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:09.968388277-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: targets-total-cost (FEATURE-ReportTargets.md)
  intent: The TOTAL row carries the summed cost formatted to four places.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:10.869353487-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:10.869405727-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:10.869412887-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-total-cost.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:10.869353487-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:10.869405727-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:10.869412887-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: targets-empty (FEATURE-ReportTargets.md)
  intent: With no matching record the report is the header plus a zero TOTAL row.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:11.748152757-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:11.748192825-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:11.748203885-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:11.748152757-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:11.748192825-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:11.748203885-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: targets-write (FEATURE-ReportTargets.md)
  intent: The writer emits the declared header and columns.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:12.750437487-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:12.750482792-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:12.750493643-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-write.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:12.750437487-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:12.750482792-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:12.750493643-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- UNVERIFIED: targets-run-via-cli (FEATURE-ReportTargets.md)
  intent: cli.Run dispatches the targets report over a test-written log, exits zero and writes parseable CSV.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:50:14.02220488-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:50:14.03100595-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:50:14.0310505-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:50:14.03165285-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.009s [no tests to run]\n"}
    {"Time":"2026-10-26T11:50:14.031675177-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:50:14.02220488-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:14.03100595-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:14.0310505-04:00","Action":"output","Package":"ddkstat/internal/cli… (+326 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-run-via-cli.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:50:14.02220488-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:14.03100595-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:14.0310505-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:50:14.03165285-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.009s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:50:14.031675177-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.009}\n', '')
- RED: commands-ordering (FEATURE-ReportCommands.md)
  intent: Rows are ordered by calls descending then command ascending, including ties.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:14.959310009-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:14.959358733-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:14.959366144-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-ordering.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:14.959310009-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:14.959358733-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:14.959366144-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: commands-total (FEATURE-ReportCommands.md)
  intent: The TOTAL row sums calls, failures, tokens and minutes over all filtered records.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:15.910235599-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:15.910302172-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:15.910310554-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-total.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:15.910235599-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:15.910302172-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:15.910310554-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: commands-empty (FEATURE-ReportCommands.md)
  intent: With no matching record the report is the header plus a zero TOTAL row.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:16.842343084-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:16.842391636-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:16.842400499-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:16.842343084-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:16.842391636-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:16.842400499-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: commands-write (FEATURE-ReportCommands.md)
  intent: The writer emits the declared header and columns.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:17.701862688-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:17.701907021-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:17.701915726-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-write.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:17.701862688-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:17.701907021-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:17.701915726-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- UNVERIFIED: commands-run-via-cli (FEATURE-ReportCommands.md)
  intent: cli.Run dispatches the commands report over a test-written log, exits zero and writes parseable CSV.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:50:19.174333071-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:50:19.180453369-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:50:19.180493098-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:50:19.181080076-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-26T11:50:19.181099971-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:50:19.174333071-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:19.180453369-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:19.180493098-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-run-via-cli.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:50:19.174333071-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:19.180453369-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:19.180493098-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:50:19.181080076-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:50:19.181099971-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}\n', '')
- RED: models-grouping-executed-model (FEATURE-ReportModels.md)
  intent: Records group under the executed model, falling back to the requested model.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:20.177013258-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:20.177070186-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:20.177086708-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-grouping-executed-model.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:20.177013258-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:20.177070186-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:20.177086708-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: models-ordering (FEATURE-ReportModels.md)
  intent: Rows are ordered by provider then model ascending in byte order.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:21.005062556-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:21.00511932-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:21.005126764-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+440 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-ordering.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:21.005062556-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:21.00511932-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:21.005126764-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: models-total-empty-model-column (FEATURE-ReportModels.md)
  intent: The TOTAL row leaves the model column empty and sums every numeric column.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:21.814268173-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:21.814314594-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:21.81432343-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+440 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-total-empty-model-column.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:21.814268173-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:21.814314594-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:21.81432343-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: models-empty (FEATURE-ReportModels.md)
  intent: With no matching record the report is the header plus a zero TOTAL row.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:22.748811494-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:22.748852662-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:22.748859977-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:22.748811494-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:22.748852662-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:22.748859977-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: models-write (FEATURE-ReportModels.md)
  intent: The writer emits the declared header and columns.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    {"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\n"}
    {"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\n"}
    {"ImportPath":"./internal/report","Action":"build-fail"}
    {"Time":"2026-10-26T11:50:23.775230515-04:00","Action":"start","Package":"./internal/report"}
    {"Time":"2026-10-26T11:50:23.775284673-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\t./internal/report [setup failed]\n"}
    {"Time":"2026-10-26T11:50:23.775293403-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-write.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:23.775230515-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:23.775284673-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:23.775293403-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- UNVERIFIED: models-run-via-cli (FEATURE-ReportModels.md)
  intent: cli.Run dispatches the models report over a test-written log, exits zero and writes parseable CSV.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-26T11:50:24.856046332-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:50:24.862711771-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-26T11:50:24.862769865-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:50:24.863525276-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-26T11:50:24.863545539-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-26T11:50:24.856046332-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:24.862711771-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:24.862769865-04:00","Action":"output","Package":"ddkstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-run-via-cli.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-26T11:50:24.856046332-04:00","Action":"start","Package":"ddkstat/internal/cli"}\n{"Time":"2026-10-26T11:50:24.862711771-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-26T11:50:24.862769865-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-26T11:50:24.863525276-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \\tddkstat/internal/cli\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-26T11:50:24.863545539-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.008}\n', '')
- RED: report-failures-unit (FEATURE-ReportFailures.md)
  intent: The native unit tests for failure selection, ordering, null returncode, timed_out text and empty output all pass.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    TestFailuresOrdering execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+441 chars)
      name = 'TestFailuresOrdering'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-failures-unit.py", line 12, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestFailuresOrdering', 'execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:25.683855811-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:25.683899629-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:25.683908337-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: report-failures-rows (FEATURE-ReportFailures.md)
  intent: The failures command lists only failed records ordered by started_at then execution_id, with model fallback, empty null returncode and true/false timed_out.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-failures-rows.py", line 43, in <module>
        assert rows[0] == ["started_at", "target", "command", "model", "returncode", "timed_out"]
               ~~~~^^^
    IndexError: list index out of range
- RED: report-failures-filters (FEATURE-ReportFailures.md)
  intent: The failures command honours --target and the inclusive --since/--until date range.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      r = {'schema_version': 1, 'execution_id': 'id-4', 'started_at': '2026-10-16T01:00:00.000Z', 'status': 'failed', 'job': {'command_name': 'build', 'target': 'Alpha', 'llm': 'claude', 'model': 'm'}, 'result': {'returncode': 1, 'timed_out': False, 'stats': {'elapsed_ms': 1}}}
      kept = [{'schema_version': 1, 'execution_id': 'id-2', 'started_at': '2026-10-15T01:00:00.000Z', 'status': 'failed', 'job': {'command_name': 'build', 'target': 'Alpha', 'llm': 'claude', 'model': 'm'}, 'result': {'returncode': 1, 'timed_out': False, 'stats': {'elapsed_ms': 1}}}]
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-failures-filters.py", line 46, in <module>
        assert [row[0] for row in rows[1:]] == [r["started_at"] for r in kept]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: report-failures-empty (FEATURE-ReportFailures.md)
  intent: When no failed record matches, the failures command exits 0 and writes only the header row.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-failures-empty.py", line 27, in <module>
        assert len(rows) == 1
               ^^^^^^^^^^^^^^
    AssertionError
- RED: report-daily-unit (FEATURE-ReportDaily.md)
  intent: The native unit tests for daily aggregation, ordering, totals, empty input and writing all pass.
  return code: 1
  error: native acceptance: execution-error
  stdout:
    TestDailyOrdering execution-error
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "execution-error"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='execution-error', return_code=1, stdout='{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal… (+439 chars)
      name = 'TestDailyOrdering'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-daily-unit.py", line 12, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestDailyOrdering', 'execution-error', '{"ImportPath":"./internal/report","Action":"build-output","Output":"# ./internal/report\\n"}\n{"ImportPath":"./internal/report","Action":"build-output","Output":"stat /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/internal/report: directory not found\\n"}\n{"ImportPath":"./internal/report","Action":"build-fail"}\n{"Time":"2026-10-26T11:50:35.7297846-04:00","Action":"start","Package":"./internal/report"}\n{"Time":"2026-10-26T11:50:35.729874026-04:00","Action":"output","Package":"./internal/report","Output":"FAIL\\t./internal/report [setup failed]\\n"}\n{"Time":"2026-10-26T11:50:35.729882532-04:00","Action":"fail","Package":"./internal/report","Elapsed":0,"FailedBuild":"./internal/report"}\n', '')
- RED: report-daily-rows (FEATURE-ReportDaily.md)
  intent: The daily command emits one row per start date with calls, failed, tokens and minutes derived from the log, followed by a TOTAL row.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-daily-rows.py", line 69, in <module>
        header, body = rows[0], rows[1:]
                       ~~~~^^^
    IndexError: list index out of range
- RED: report-daily-filters (FEATURE-ReportDaily.md)
  intent: The daily command honours --target, --since and --until, computing every row and the TOTAL over only the matching records.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      body = []
      r = {'schema_version': 1, 'execution_id': 'id-004', 'started_at': '2026-10-16T01:00:00.000Z', 'status': 'succeeded', 'job': {'command_name': 'build', 'target': 'Alpha', 'llm': 'claude', 'model': 'haiku'}, 'result': {'returncode': 0, 'timed_out': False, 'stats': {'input_tokens': 44, 'cached_input_tokens': 0, 'output_tokens': 0, 'elapsed_ms': 60000, 'cost_usd': 0.1, 'model': 'claude-haiku-5-5'}}}
      kept = [{'schema_version': 1, 'execution_id': 'id-002', 'started_at': '2026-10-15T01:00:00.000Z', 'status': 'failed', 'job': {'command_name': 'build', 'target': 'Alpha', 'llm': 'claude', 'model': 'haiku'}, 'result': {'returncode': 0, 'timed_out': False, 'stats': {'input_tokens': 22, 'cached_input_tokens': 0, 'output_tokens': 0, 'elapsed_ms': 60000, 'cost_usd': 0.1, 'model': 'claude-haiku-5-5'}}}, {'schem… (+387 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-daily-filters.py", line 60, in <module>
        assert [row[0] for row in body[:-1]] == sorted({r["started_at"][:10] for r in kept})
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: report-daily-empty (FEATURE-ReportDaily.md)
  intent: When no record matches the filter, the daily command exits 0 and writes the header and a TOTAL row of zeros.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-daily-empty.py", line 28, in <module>
        assert rows[0] == ["date", "calls", "failed", "tokens", "minutes"]
               ~~~~^^^
    IndexError: list index out of range

## Post-build programmatic acceptance
- PASS: summary-aggregate (FEATURE-ReportSummary.md)
  intent: SummaryAggregate sums counts, tokens, cost and duration over records built by the test.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:38.04536849-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:38.052436254-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestSummaryAggregate"}
    {"Time":"2026-10-26T11:55:38.052473898-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummaryAggregate","Output":"=== RUN   TestSummaryAggregate\n"}
    {"Time":"2026-10-26T11:55:38.052502945-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummaryAggregate","Output":"--- PASS: TestSummaryAggregate (0.00s)\n"}
    {"Time":"2026-10-26T11:55:38.052505723-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestSummaryAggregate","Elapsed":0}
    {"Time":"2026-10-26T11:55:38.052508449-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:38.053528054-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.008s\n"}
    {"Time":"2026-10-26T11:55:38.053551018-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summary-empty-first-last (FEATURE-ReportSummary.md)
  intent: With no matching record, first and last are empty and counters are zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:39.344604072-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:39.353193438-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestSummaryEmptyFirstLast"}
    {"Time":"2026-10-26T11:55:39.353232555-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummaryEmptyFirstLast","Output":"=== RUN   TestSummaryEmptyFirstLast\n"}
    {"Time":"2026-10-26T11:55:39.353240679-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummaryEmptyFirstLast","Output":"--- PASS: TestSummaryEmptyFirstLast (0.00s)\n"}
    {"Time":"2026-10-26T11:55:39.353247992-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestSummaryEmptyFirstLast","Elapsed":0}
    {"Time":"2026-10-26T11:55:39.353250928-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:39.353651094-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.009s\n"}
    {"Time":"2026-10-26T11:55:39.353670018-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summary-skipped-ignores-filters (FEATURE-ReportSummary.md)
  intent: The skipped count stays the whole-file count regardless of target and date filters.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:40.718786764-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:40.725723934-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestSummarySkippedIgnoresFilters"}
    {"Time":"2026-10-26T11:55:40.725745341-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummarySkippedIgnoresFilters","Output":"=== RUN   TestSummarySkippedIgnoresFilters\n"}
    {"Time":"2026-10-26T11:55:40.72616187-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummarySkippedIgnoresFilters","Output":"--- PASS: TestSummarySkippedIgnoresFilters (0.00s)\n"}
    {"Time":"2026-10-26T11:55:40.726180207-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestSummarySkippedIgnoresFilters","Elapsed":0}
    {"Time":"2026-10-26T11:55:40.726184269-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:40.726894231-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.008s\n"}
    {"Time":"2026-10-26T11:55:40.726914346-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summary-write (FEATURE-ReportSummary.md)
  intent: WriteSummary writes the declared header and one row with the declared column order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:42.168061331-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:42.173968194-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestSummaryWrite"}
    {"Time":"2026-10-26T11:55:42.173990204-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummaryWrite","Output":"=== RUN   TestSummaryWrite\n"}
    {"Time":"2026-10-26T11:55:42.174013157-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestSummaryWrite","Output":"--- PASS: TestSummaryWrite (0.00s)\n"}
    {"Time":"2026-10-26T11:55:42.174015074-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestSummaryWrite","Elapsed":0}
    {"Time":"2026-10-26T11:55:42.174018012-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:42.174563891-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.006s\n"}
    {"Time":"2026-10-26T11:55:42.174571395-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summary-run-via-cli (FEATURE-ReportSummary.md)
  intent: cli.Run dispatches the summary report over a test-written log and exits zero with a parseable CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:43.382213845-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:55:43.388304327-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestRunSummaryReport"}
    {"Time":"2026-10-26T11:55:43.388386243-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunSummaryReport","Output":"=== RUN   TestRunSummaryReport\n"}
    {"Time":"2026-10-26T11:55:43.388789209-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunSummaryReport","Output":"--- PASS: TestRunSummaryReport (0.00s)\n"}
    {"Time":"2026-10-26T11:55:43.388812447-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestRunSummaryReport","Elapsed":0}
    {"Time":"2026-10-26T11:55:43.38881676-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:43.389588282-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-26T11:55:43.38960916-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: targets-ordering (FEATURE-ReportTargets.md)
  intent: Rows are ordered by target ascending in byte order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:44.786283485-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:44.792782524-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestTargetsOrdering"}
    {"Time":"2026-10-26T11:55:44.792803989-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsOrdering","Output":"=== RUN   TestTargetsOrdering\n"}
    {"Time":"2026-10-26T11:55:44.79282155-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsOrdering","Output":"--- PASS: TestTargetsOrdering (0.00s)\n"}
    {"Time":"2026-10-26T11:55:44.792823606-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestTargetsOrdering","Elapsed":0}
    {"Time":"2026-10-26T11:55:44.792826416-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:44.793399931-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.007s\n"}
    {"Time":"2026-10-26T11:55:44.793428819-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: targets-total-cost (FEATURE-ReportTargets.md)
  intent: The TOTAL row carries the summed cost formatted to four places.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:46.117118953-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:46.123732599-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestTargetsTotalCost"}
    {"Time":"2026-10-26T11:55:46.123761365-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsTotalCost","Output":"=== RUN   TestTargetsTotalCost\n"}
    {"Time":"2026-10-26T11:55:46.123784032-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsTotalCost","Output":"--- PASS: TestTargetsTotalCost (0.00s)\n"}
    {"Time":"2026-10-26T11:55:46.12378738-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestTargetsTotalCost","Elapsed":0}
    {"Time":"2026-10-26T11:55:46.123790475-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:46.124331065-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.007s\n"}
    {"Time":"2026-10-26T11:55:46.124355055-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: targets-empty (FEATURE-ReportTargets.md)
  intent: With no matching record the report is the header plus a zero TOTAL row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:47.46788623-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:47.475140565-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestTargetsEmpty"}
    {"Time":"2026-10-26T11:55:47.475172984-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsEmpty","Output":"=== RUN   TestTargetsEmpty\n"}
    {"Time":"2026-10-26T11:55:47.475196302-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsEmpty","Output":"--- PASS: TestTargetsEmpty (0.00s)\n"}
    {"Time":"2026-10-26T11:55:47.475198193-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestTargetsEmpty","Elapsed":0}
    {"Time":"2026-10-26T11:55:47.475201872-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:47.476364864-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.008s\n"}
    {"Time":"2026-10-26T11:55:47.476386187-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: targets-write (FEATURE-ReportTargets.md)
  intent: The writer emits the declared header and columns.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:48.880820749-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:48.886578005-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestTargetsWrite"}
    {"Time":"2026-10-26T11:55:48.886600907-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsWrite","Output":"=== RUN   TestTargetsWrite\n"}
    {"Time":"2026-10-26T11:55:48.886620278-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestTargetsWrite","Output":"--- PASS: TestTargetsWrite (0.00s)\n"}
    {"Time":"2026-10-26T11:55:48.886622385-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestTargetsWrite","Elapsed":0}
    {"Time":"2026-10-26T11:55:48.886625771-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:48.887287406-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.006s\n"}
    {"Time":"2026-10-26T11:55:48.887308284-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: targets-run-via-cli (FEATURE-ReportTargets.md)
  intent: cli.Run dispatches the targets report over a test-written log, exits zero and writes parseable CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:49.936045497-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:55:49.941696058-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestRunTargetsReport"}
    {"Time":"2026-10-26T11:55:49.941716351-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunTargetsReport","Output":"=== RUN   TestRunTargetsReport\n"}
    {"Time":"2026-10-26T11:55:49.942141498-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunTargetsReport","Output":"--- PASS: TestRunTargetsReport (0.00s)\n"}
    {"Time":"2026-10-26T11:55:49.942155114-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestRunTargetsReport","Elapsed":0}
    {"Time":"2026-10-26T11:55:49.942158051-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:49.942802994-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-26T11:55:49.942830712-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: commands-ordering (FEATURE-ReportCommands.md)
  intent: Rows are ordered by calls descending then command ascending, including ties.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:51.334883777-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:51.341482002-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestCommandsOrdering"}
    {"Time":"2026-10-26T11:55:51.341517209-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsOrdering","Output":"=== RUN   TestCommandsOrdering\n"}
    {"Time":"2026-10-26T11:55:51.341524859-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsOrdering","Output":"--- PASS: TestCommandsOrdering (0.00s)\n"}
    {"Time":"2026-10-26T11:55:51.341526673-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestCommandsOrdering","Elapsed":0}
    {"Time":"2026-10-26T11:55:51.341529209-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:51.342071821-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.007s\n"}
    {"Time":"2026-10-26T11:55:51.342094493-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: commands-total (FEATURE-ReportCommands.md)
  intent: The TOTAL row sums calls, failures, tokens and minutes over all filtered records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:52.684570631-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:52.688490228-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestCommandsTotal"}
    {"Time":"2026-10-26T11:55:52.688523621-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsTotal","Output":"=== RUN   TestCommandsTotal\n"}
    {"Time":"2026-10-26T11:55:52.688556638-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsTotal","Output":"--- PASS: TestCommandsTotal (0.00s)\n"}
    {"Time":"2026-10-26T11:55:52.688559488-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestCommandsTotal","Elapsed":0}
    {"Time":"2026-10-26T11:55:52.688562676-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:52.689241524-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.004s\n"}
    {"Time":"2026-10-26T11:55:52.689264795-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: commands-empty (FEATURE-ReportCommands.md)
  intent: With no matching record the report is the header plus a zero TOTAL row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:53.498484026-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:53.509648449-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestCommandsEmpty"}
    {"Time":"2026-10-26T11:55:53.509670608-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsEmpty","Output":"=== RUN   TestCommandsEmpty\n"}
    {"Time":"2026-10-26T11:55:53.509687237-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsEmpty","Output":"--- PASS: TestCommandsEmpty (0.00s)\n"}
    {"Time":"2026-10-26T11:55:53.509689222-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestCommandsEmpty","Elapsed":0}
    {"Time":"2026-10-26T11:55:53.509692232-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:53.510291426-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.012s\n"}
    {"Time":"2026-10-26T11:55:53.510309295-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.012}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: commands-write (FEATURE-ReportCommands.md)
  intent: The writer emits the declared header and columns.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:54.771821454-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:54.777892158-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestCommandsWrite"}
    {"Time":"2026-10-26T11:55:54.777918317-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsWrite","Output":"=== RUN   TestCommandsWrite\n"}
    {"Time":"2026-10-26T11:55:54.777939365-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestCommandsWrite","Output":"--- PASS: TestCommandsWrite (0.00s)\n"}
    {"Time":"2026-10-26T11:55:54.777941683-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestCommandsWrite","Elapsed":0}
    {"Time":"2026-10-26T11:55:54.777944481-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:54.778587651-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.007s\n"}
    {"Time":"2026-10-26T11:55:54.778608579-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: commands-run-via-cli (FEATURE-ReportCommands.md)
  intent: cli.Run dispatches the commands report over a test-written log, exits zero and writes parseable CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:56.107165842-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:55:56.111559358-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestRunCommandsReport"}
    {"Time":"2026-10-26T11:55:56.111592773-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunCommandsReport","Output":"=== RUN   TestRunCommandsReport\n"}
    {"Time":"2026-10-26T11:55:56.112037507-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunCommandsReport","Output":"--- PASS: TestRunCommandsReport (0.00s)\n"}
    {"Time":"2026-10-26T11:55:56.112062873-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestRunCommandsReport","Elapsed":0}
    {"Time":"2026-10-26T11:55:56.112069086-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:56.114414723-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-26T11:55:56.114436733-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: models-grouping-executed-model (FEATURE-ReportModels.md)
  intent: Records group under the executed model, falling back to the requested model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:57.500089759-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:57.505981835-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestModelsGroupingUsesExecutedModel"}
    {"Time":"2026-10-26T11:55:57.506002108-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsGroupingUsesExecutedModel","Output":"=== RUN   TestModelsGroupingUsesExecutedModel\n"}
    {"Time":"2026-10-26T11:55:57.506023-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsGroupingUsesExecutedModel","Output":"--- PASS: TestModelsGroupingUsesExecutedModel (0.00s)\n"}
    {"Time":"2026-10-26T11:55:57.506024913-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestModelsGroupingUsesExecutedModel","Elapsed":0}
    {"Time":"2026-10-26T11:55:57.506033419-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:57.506602837-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.006s\n"}
    {"Time":"2026-10-26T11:55:57.506620868-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: models-ordering (FEATURE-ReportModels.md)
  intent: Rows are ordered by provider then model ascending in byte order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:55:58.889701986-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:55:58.895950195-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestModelsOrdering"}
    {"Time":"2026-10-26T11:55:58.895973967-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsOrdering","Output":"=== RUN   TestModelsOrdering\n"}
    {"Time":"2026-10-26T11:55:58.895990219-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsOrdering","Output":"--- PASS: TestModelsOrdering (0.00s)\n"}
    {"Time":"2026-10-26T11:55:58.895992466-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestModelsOrdering","Elapsed":0}
    {"Time":"2026-10-26T11:55:58.895994921-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:55:58.896707589-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.007s\n"}
    {"Time":"2026-10-26T11:55:58.896729573-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: models-total-empty-model-column (FEATURE-ReportModels.md)
  intent: The TOTAL row leaves the model column empty and sums every numeric column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:56:00.120961631-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:56:00.126471071-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestModelsTotalEmptyModelColumn"}
    {"Time":"2026-10-26T11:56:00.126513169-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsTotalEmptyModelColumn","Output":"=== RUN   TestModelsTotalEmptyModelColumn\n"}
    {"Time":"2026-10-26T11:56:00.126527778-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsTotalEmptyModelColumn","Output":"--- PASS: TestModelsTotalEmptyModelColumn (0.00s)\n"}
    {"Time":"2026-10-26T11:56:00.126529623-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestModelsTotalEmptyModelColumn","Elapsed":0}
    {"Time":"2026-10-26T11:56:00.126532365-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:56:00.127114706-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.006s\n"}
    {"Time":"2026-10-26T11:56:00.12713553-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: models-empty (FEATURE-ReportModels.md)
  intent: With no matching record the report is the header plus a zero TOTAL row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:56:01.393334791-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:56:01.399089439-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestModelsEmpty"}
    {"Time":"2026-10-26T11:56:01.399108442-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsEmpty","Output":"=== RUN   TestModelsEmpty\n"}
    {"Time":"2026-10-26T11:56:01.399126242-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsEmpty","Output":"--- PASS: TestModelsEmpty (0.00s)\n"}
    {"Time":"2026-10-26T11:56:01.399127969-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestModelsEmpty","Elapsed":0}
    {"Time":"2026-10-26T11:56:01.399130782-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:56:01.399664078-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.006s\n"}
    {"Time":"2026-10-26T11:56:01.399682298-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: models-write (FEATURE-ReportModels.md)
  intent: The writer emits the declared header and columns.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:56:02.710745495-04:00","Action":"start","Package":"ddkstat/internal/report"}
    {"Time":"2026-10-26T11:56:02.716477709-04:00","Action":"run","Package":"ddkstat/internal/report","Test":"TestModelsWrite"}
    {"Time":"2026-10-26T11:56:02.716521302-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsWrite","Output":"=== RUN   TestModelsWrite\n"}
    {"Time":"2026-10-26T11:56:02.71652821-04:00","Action":"output","Package":"ddkstat/internal/report","Test":"TestModelsWrite","Output":"--- PASS: TestModelsWrite (0.00s)\n"}
    {"Time":"2026-10-26T11:56:02.716531833-04:00","Action":"pass","Package":"ddkstat/internal/report","Test":"TestModelsWrite","Elapsed":0}
    {"Time":"2026-10-26T11:56:02.71653482-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"PASS\n"}
    {"Time":"2026-10-26T11:56:02.717261083-04:00","Action":"output","Package":"ddkstat/internal/report","Output":"ok  \tddkstat/internal/report\t0.006s\n"}
    {"Time":"2026-10-26T11:56:02.717291811-04:00","Action":"pass","Package":"ddkstat/internal/report","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: models-run-via-cli (FEATURE-ReportModels.md)
  intent: cli.Run dispatches the models report over a test-written log, exits zero and writes parseable CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-26T11:56:03.852523385-04:00","Action":"start","Package":"ddkstat/internal/cli"}
    {"Time":"2026-10-26T11:56:03.857931414-04:00","Action":"run","Package":"ddkstat/internal/cli","Test":"TestRunModelsReport"}
    {"Time":"2026-10-26T11:56:03.857962602-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunModelsReport","Output":"=== RUN   TestRunModelsReport\n"}
    {"Time":"2026-10-26T11:56:03.858440576-04:00","Action":"output","Package":"ddkstat/internal/cli","Test":"TestRunModelsReport","Output":"--- PASS: TestRunModelsReport (0.00s)\n"}
    {"Time":"2026-10-26T11:56:03.85846832-04:00","Action":"pass","Package":"ddkstat/internal/cli","Test":"TestRunModelsReport","Elapsed":0}
    {"Time":"2026-10-26T11:56:03.858472049-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-26T11:56:03.859188167-04:00","Action":"output","Package":"ddkstat/internal/cli","Output":"ok  \tddkstat/internal/cli\t0.006s\n"}
    {"Time":"2026-10-26T11:56:03.859211446-04:00","Action":"pass","Package":"ddkstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: report-failures-unit (FEATURE-ReportFailures.md)
  intent: The native unit tests for failure selection, ordering, null returncode, timed_out text and empty output all pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestFailuresOrdering pass
    TestFailuresNullReturncodeEmpty pass
    TestFailuresTimedOutText pass
    TestFailuresOnlyFailed pass
    TestFailuresEmpty pass
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: report-failures-rows (FEATURE-ReportFailures.md)
  intent: The failures command lists only failed records ordered by started_at then execution_id, with model fallback, empty null returncode and true/false timed_out.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    2026-10-13T23:00:00.000Z,Beta,build,executed-model,124,true
    2026-10-14T02:00:00.000Z,Alpha,plan,executed-model,2,false
    2026-10-14T02:00:00.000Z,Alpha,build,requested-model,1,false
    2026-10-15T06:00:00.000Z,Beta,diagnose,executed-model,,true
- PASS: report-failures-filters (FEATURE-ReportFailures.md)
  intent: The failures command honours --target and the inclusive --since/--until date range.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    2026-10-15T01:00:00.000Z,Alpha,build,m,1,false
- PASS: report-failures-empty (FEATURE-ReportFailures.md)
  intent: When no failed record matches, the failures command exits 0 and writes only the header row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
- PASS: report-daily-unit (FEATURE-ReportDaily.md)
  intent: The native unit tests for daily aggregation, ordering, totals, empty input and writing all pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestDailyOrdering pass
    TestDailyTokensSum pass
    TestDailyTotal pass
    TestDailyEmpty pass
    TestDailyWrite pass
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: report-daily-rows (FEATURE-ReportDaily.md)
  intent: The daily command emits one row per start date with calls, failed, tokens and minutes derived from the log, followed by a TOTAL row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    2026-10-14,2,1,113,1.4
    2026-10-15,2,0,280,3.7
    2026-10-16,1,0,15,0.0
    TOTAL,5,1,408,5.0
- PASS: report-daily-filters (FEATURE-ReportDaily.md)
  intent: The daily command honours --target, --since and --until, computing every row and the TOTAL over only the matching records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    2026-10-15,1,1,22,1.0
    2026-10-16,1,0,44,1.0
    TOTAL,2,1,66,2.0
- PASS: report-daily-empty (FEATURE-ReportDaily.md)
  intent: When no record matches the filter, the daily command exits 0 and writes the header and a TOTAL row of zeros.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    TOTAL,0,0,0,0.0

## Build summary
I'll start by inspecting the build directory and the staged sources.The report package doesn't exist yet. I'll read the existing CLI, filter, format, and logread code to see what's already there.The existing `cli` package already has the registry and the filter path. What's missing is the `internal/report` package and the registration hook. I'll confirm the existing tests pass before I change anything.The baseline is green. Next I'll write the `internal/report` package: a shared base file, then one file per report.I used a nonexistent `recordFields` type in `Add`. I'll fix it to `logread.Record` before going further.Next, `main.go` needs the blank import so the binary registers the writers. The existing "unregistered report" test in `internal/cli` can no longer hold once every report is registered, so I'll remove that one test.Now the CLI-level reachability tests for the `internal/cli` acceptance checks. Each one re-registers the real writers first, because the existing tests overwrite the registry globally.Next, the `internal/report` unit tests. I'll start with the shared helpers: a record builder, a CSV read-back helper, and a log writer for the filter test.Now the summary tests. The skipped-count test goes through `cli.Run` on a real log file, so the filter path is exercised, not just the aggregate.I left junk in `TestSummarySkippedIgnoresFilters`: a stray `report := cli.Run` and a check that reads back a nil-record writer instead of the CLI's stdout. I'll rewrite it to parse the CLI's own output.Now the remaining report tests, one file per report: targets, commands, models, failures, and daily.All the test files are written. I'll run gofmt, vet, and the full test suite.Tests pass. gofmt flags `run_test.go`, most likely a blank line left by the removed test. I'll check the diff before applying it.The only gofmt change was a trailing blank line I created when I removed the test, so the diff is mine. Vet is clean. Next I'll run the staged grader, `sources/run.sh`, which the COMPASS names as the terminal check. It only reads the sample log.All 20 grader checks pass. Last, I'll confirm every test name the specs require exists and passes, and that nothing under `sources/` changed.All six report writers are built and registered, and the grader passes 20 of 20.

Checks run:
- `sh bin/test.sh` (`go vet` plus `go test ./...`) passes for every package.
- `sh sources/run.sh` exits 0 with 20 of 20 passing.
- All 31 test names the specs require (the 27 report unit tests plus the 4 `internal/cli` reachability tests) exist and pass.
- `sources/` is unchanged.

Choices I made:
- **Registration:** the `report` package registers its writers from `init()`, and `cmd/ddkstat/main.go` imports it with a blank import. `cli.Run` still never imports `report`.
- **Re-registration in tests:** I exported `report.Register()` and call it from the new `internal/cli` tests. The existing CLI tests overwrite registry entries and never restore them, so the new tests would otherwise see a stale writer.
- **Removed test:** I deleted `TestRunUnregisteredReportExit0WritesNothing` from `internal/cli/run_test.go`. It used `daily` as the unregistered report, and once every report is registered that case no longer exists. The `nil` writer branch in `run.go` is still there, but nothing exercises it now.
- **Shared accumulator:** the grouped reports share one `Totals` type, so the TOTAL rows and the per-group rows are computed the same way.
- **Output:** the report writers write unquoted text, as the specs require. The tests read the output back with `encoding/csv` rather than matching substrings.

I did not commit anything.

Files changed:
- cmd/ddkstat/main.go
- internal/cli/run_test.go
- internal/cli/reports_test.go
- internal/report/report.go
- internal/report/summary.go
- internal/report/targets.go
- internal/report/commands.go
- internal/report/models.go
- internal/report/failures.go
- internal/report/daily.go
- internal/report/helpers_test.go
- internal/report/summary_test.go
- internal/report/targets_test.go
- internal/report/commands_test.go
- internal/report/models_test.go
- internal/report/failures_test.go
- internal/report/daily_test.go

<blueprint-decisions>
[{"spec":"FEATURE-ReportSummary.md","severity":"Low","subject":"Writer registration and test isolation","decision":"Report writers register from an init() in internal/report, with an exported report.Register() that re-registers all six. cmd/ddkstat imports the package for its side effect. The new internal/cli tests call report.Register() because existing tests overwrite registry entries globally. I removed TestRunUnregisteredReportExit0WritesNothing, since every known report now has a writer. Is that acceptable, or should registration move on replan?"}]
</blueprint-decisions>

RESULT: SUCCESS

FILES CHANGED:
- cmd/ddkstat/main.go
- internal/cli/run_test.go
- internal/cli/reports_test.go
- internal/report/report.go
- internal/report/summary.go
- internal/report/targets.go
- internal/report/commands.go
- internal/report/models.go
- internal/report/failures.go
- internal/report/daily.go
- internal/report/helpers_test.go
- internal/report/summary_test.go
- internal/report/targets_test.go
- internal/report/commands_test.go
- internal/report/models_test.go
- internal/report/failures_test.go
- internal/report/daily_test.go

SUMMARY:
Implemented the six ddkstat reports (summary, targets, commands, models, failures, daily) as an `internal/report` package. Each report has a pure aggregate or selection function and a CSV writer registered with `cli.Register`. `go vet ./...` is clean, `go test ./...` passes in every package, and `sh sources/run.sh` passes 20 of 20 against the staged sample log. All 31 required test names pass. Nothing was committed.

BLOCKERS:
- None.
