# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-24
- resulting state: closed/failed
- story points (combined assembled cost): 35339
- execution id: 20261024.165821.563Z-d19d2cff

## Stories built
- Provide the ddkstat entry point and the Run dispatcher. (cli-entry) [story]
- Parse the report name, options, and the single log path. (cli-parse) [story]
- Reject usage errors with exit code 2. (cli-usage) [story]
- Write CSV with a header, LF line endings, and TOTAL rows. (fmt-csv) [story]
- Format summed cost with four decimal places. (fmt-cost) [story]
- Format minutes in tenths of a minute. (fmt-minutes) [story]
- Read JSON Lines log records from a file path. (log-read) [story]
- Report operational log read failures with exit code 1. (cli-read-fail) [story]
- Validate date option values as real calendar dates. (filter-validate) [story]
- Parse records and skip non-object, incomplete, and duplicate lines. (log-skip) [story]
- Attribute each record to its executed or requested model. (log-model) [story]
- Treat absent or null numeric stats as zero. (log-numbers) [story]
- Classify a record as failed when its status is not succeeded. (log-failed) [story]
- Filter records by an inclusive start-date range. (filter-date) [story]
- Filter records by Target name. (filter-target) [story]
- Produce the per-start-date report. (daily-report) [story]
- Produce the failed-record report. (failures-report) [story]
- Produce the per-Target report. (targets-report) [story]
- Count skipped lines over the whole file regardless of filters. (summary-skipped) [story]
- Produce the per-provider and per-model report. (models-report) [story]
- Produce the one-row summary report. (summary-report) [story]
- Produce the per-command report. (commands-report) [story]

## Acceptance tooling authorization
- FEATURE-Entrypoint.md#entrypoint-builds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Entrypoint.md#entrypoint-no-args-exit-2: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CommandParse.md#parse-summary-with-target: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CommandParse.md#parse-repeated-option-exit-2: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-unknown-report-exit-2: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-no-log-path-exit-2: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FormatCSV.md#csv-header-and-lf: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FormatCSV.md#csv-zero-match-total: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FormatCost.md#fmt-cost-four-places: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FormatCost.md#fmt-cost-zero: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FormatMinutes.md#fmt-minutes-example: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FormatMinutes.md#fmt-minutes-zero: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogRead.md#log-read-blank-lines: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogRead.md#log-read-utf8: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogRead.md#log-read-missing-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReadErrors.md#read-missing-log-exit-1: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterValidate.md#filter-validate-rejects-impossible-date: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterValidate.md#filter-validate-rejects-invalid-day: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterValidate.md#filter-validate-rejects-date-after-log: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogSkip.md#log-skip-counts-and-exit: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogSkip.md#log-skip-duplicate-first-kept: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogSkip.md#log-skip-warning-per-line: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogModel.md#log-model-prefers-executed: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogModel.md#log-model-falls-back-to-requested: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogNumbers.md#log-numbers-null-and-absent-are-zero: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogNumbers.md#log-numbers-no-stats-object: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogFailed.md#log-failed-status-classification: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogFailed.md#log-failed-returncode-and-timeout: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FilterDate.md#filter-date-inclusive-bounds: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterDate.md#filter-date-since-only: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterDate.md#filter-date-inverted-range-exit-zero: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-exact-match: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-absent-keeps-all: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-no-match-exit-zero: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-DailyReport.md#daily-report-whole-sample: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-DailyReport.md#daily-report-since-until: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-DailyReport.md#daily-report-zero-match: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FailuresReport.md#failures-report-whole-sample: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FailuresReport.md#failures-report-target-filter: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FailuresReport.md#failures-report-zero-match: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-TargetsReport.md#targets-ordering-and-total: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-TargetsReport.md#targets-zero-match: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-SkippedCount.md#skipped-ignores-target: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-SkippedCount.md#skipped-ignores-date: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ModelsReport.md#models-report-whole-sample: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ModelsReport.md#models-report-target-filter: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ModelsReport.md#models-report-zero-match: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-SummaryReport.md#summary-header-and-row-shape: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-SummaryReport.md#summary-values-from-fixture: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CommandsReport.md#commands-ordering-and-total: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CommandsReport.md#commands-tie-byte-order: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CommandsReport.md#commands-zero-match: executable=python3; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-Entrypoint.md (SP 639)
- context: ARCHITECTURE_compact.md (SP 476)
- implements: FEATURE-CommandParse.md (SP 821)
- implements: FEATURE-UsageErrors.md (SP 742)
- implements: FEATURE-FormatCSV.md (SP 1115)
- implements: FEATURE-FormatCost.md (SP 1160)
- implements: FEATURE-FormatMinutes.md (SP 1214)
- implements: FEATURE-LogRead.md (SP 1389)
- implements: FEATURE-ReadErrors.md (SP 467)
- implements: FEATURE-FilterValidate.md (SP 1783)
- implements: FEATURE-LogSkip.md (SP 2025)
- implements: FEATURE-LogModel.md (SP 1116)
- implements: FEATURE-LogNumbers.md (SP 1200)
- implements: FEATURE-LogFailed.md (SP 1332)
- implements: FEATURE-FilterDate.md (SP 2132)
- implements: FEATURE-FilterTarget.md (SP 2278)
- implements: FEATURE-DailyReport.md (SP 1351)
- implements: FEATURE-FailuresReport.md (SP 1327)
- implements: FEATURE-TargetsReport.md (SP 1394)
- implements: FEATURE-SkippedCount.md (SP 1436)
- implements: FEATURE-ModelsReport.md (SP 1638)
- implements: FEATURE-SummaryReport.md (SP 1805)
- implements: FEATURE-CommandsReport.md (SP 1933)

## Build directory changes
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/filter/filter.go
- internal/filter/filter_test.go
- internal/format/format.go
- internal/format/format_test.go
- internal/log/log.go
- internal/log/log_test.go
- internal/log/parse.go
- internal/log/parse_test.go
- internal/report/report.go
- internal/report/report_test.go

## Pre-build acceptance observation
- GREEN (prepassed): entrypoint-builds (FEATURE-Entrypoint.md)
  intent: The ddkstat main package builds with go build ./cmd/ddkstat.
  return code: 0
- GREEN (prepassed): entrypoint-no-args-exit-2 (FEATURE-Entrypoint.md)
  intent: Invoking ddkstat with no arguments exits 2 with empty stdout.
  return code: 0
  stdout:
    usage: ddkstat <report> <log.jsonl>
- RED: parse-summary-with-target (FEATURE-CommandParse.md)
  intent: `summary --target X <log>` parses the report, the option value, and the log path.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['/tmp/tmpcifq_w03/ddkstat', 'summary', '--target', 'X', '/tmp/tmpi4r62x65/log.jsonl'], returncode=2, stdout='', stderr='error: report "summary" is not implemented\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "parse-summary-with-target.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): parse-repeated-option-exit-2 (FEATURE-CommandParse.md)
  intent: A repeated --target option is a usage error with exit 2 and empty stdout.
  return code: 0
  stdout:
    error: report "summary" is not implemented
- GREEN (prepassed): usage-unknown-report-exit-2 (FEATURE-UsageErrors.md)
  intent: An unknown report name exits 2 with empty stdout.
  return code: 0
  stdout:
    error: report "frobnicate" is not implemented
- GREEN (prepassed): usage-no-log-path-exit-2 (FEATURE-UsageErrors.md)
  intent: A report with no log path exits 2 with empty stdout.
  return code: 0
  stdout:
    error: report "summary" is not implemented
- RED: csv-header-and-lf (FEATURE-FormatCSV.md)
  intent: The CSV writer emits the header first and uses LF-only line endings.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-header-and-lf.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: csv-zero-match-total (FEATURE-FormatCSV.md)
  intent: A zero-match commands report writes its header and a zero TOTAL row, and exits 0.
  return code: 1
  stdout:
    error: report "commands" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "commands" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-zero-match-total.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "commands" is not implemented\n')
- RED: fmt-cost-four-places (FEATURE-FormatCost.md)
  intent: The summed cost renders with exactly four decimal places in the summary report.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "fmt-cost-four-places.py", line 21, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: fmt-cost-zero (FEATURE-FormatCost.md)
  intent: A zero cost total renders as 0.0000 in the summary report.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "fmt-cost-zero.py", line 21, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: fmt-minutes-example (FEATURE-FormatMinutes.md)
  intent: The minutes formatter renders 95000 ms as 1.6 by observing a summary report that sums one record of that duration.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "fmt-minutes-example.py", line 21, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: fmt-minutes-zero (FEATURE-FormatMinutes.md)
  intent: A group with zero elapsed time renders minutes as 0.0 in the summary report.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "fmt-minutes-zero.py", line 21, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-read-blank-lines (FEATURE-LogRead.md)
  intent: Blank lines are ignored by the line reader and do not appear in the yielded lines.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-blank-lines.py", line 15, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-read-utf8 (FEATURE-LogRead.md)
  intent: The line reader reads the file as UTF-8 and a multibyte character inside a valid record is kept intact.
  return code: 1
  stdout:
    error: report "commands" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "commands" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-utf8.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "commands" is not implemented\n')
- RED: log-read-missing-file (FEATURE-LogRead.md)
  intent: A log path that cannot be opened exits 1 with nothing on stdout.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-missing-file.py", line 11, in <module>
        assert code == 1, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: read-missing-log-exit-1 (FEATURE-ReadErrors.md)
  intent: A log path that does not exist exits 1 with empty stdout.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['/tmp/tmp_3vbe7b7/ddkstat', 'summary', '/tmp/tmpg5hdmt_g/no-such-log.jsonl'], returncode=2, stdout='', stderr='error: report "summary" is not implemented\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-missing-log-exit-1.py", line 12, in <module>
        assert result.returncode == 1
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): filter-validate-rejects-impossible-date (FEATURE-FilterValidate.md)
  intent: An impossible calendar date in --since exits 2 with nothing on stdout.
  return code: 0
  stdout:
    error: report "summary" is not implemented
- GREEN (prepassed): filter-validate-rejects-invalid-day (FEATURE-FilterValidate.md)
  intent: A day that does not exist in its month (February 30) is a usage error with exit 2.
  return code: 0
  stdout:
    error: report "summary" is not implemented
- GREEN (prepassed): filter-validate-rejects-date-after-log (FEATURE-FilterValidate.md)
  intent: An invalid date that appears after the log path is still a usage error with exit 2.
  return code: 0
  stdout:
    error: report "summary" is not implemented
- RED: log-skip-counts-and-exit (FEATURE-LogSkip.md)
  intent: Non-object, incomplete, and malformed lines are skipped, counted in the summary, and do not change exit status.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-skip-counts-and-exit.py", line 15, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-skip-duplicate-first-kept (FEATURE-LogSkip.md)
  intent: A duplicate execution_id is skipped on its later occurrence and the first occurrence is kept.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-skip-duplicate-first-kept.py", line 15, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-skip-warning-per-line (FEATURE-LogSkip.md)
  intent: Each skipped line writes exactly one stderr warning naming its line number, and nothing is written to stdout for skips.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-skip-warning-per-line.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-model-prefers-executed (FEATURE-LogModel.md)
  intent: A record with result.stats.model reports that executed model in the models report rather than job.model.
  return code: 1
  stdout:
    error: report "models" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "models" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-model-prefers-executed.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "models" is not implemented\n')
- RED: log-model-falls-back-to-requested (FEATURE-LogModel.md)
  intent: A record without result.stats.model reports its job.model in the models report.
  return code: 1
  stdout:
    error: report "models" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "models" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-model-falls-back-to-requested.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "models" is not implemented\n')
- RED: log-numbers-null-and-absent-are-zero (FEATURE-LogNumbers.md)
  intent: Null and absent numeric stats count as zero in the summary totals and the record is still counted.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-numbers-null-and-absent-are-zero.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-numbers-no-stats-object (FEATURE-LogNumbers.md)
  intent: A record with no result.stats object is still counted as a record with zero numeric totals.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-numbers-no-stats-object.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-failed-status-classification (FEATURE-LogFailed.md)
  intent: A record whose status is not succeeded is counted as failed in the summary, and a succeeded record is not.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-failed-status-classification.py", line 15, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: log-failed-returncode-and-timeout (FEATURE-LogFailed.md)
  intent: The failures report shows an empty returncode for a null return code and true for timed_out.
  return code: 1
  stdout:
    error: report "failures" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "failures" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-failed-returncode-and-timeout.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "failures" is not implemented\n')
- RED: filter-date-inclusive-bounds (FEATURE-FilterDate.md)
  intent: --since and --until keep records whose start date equals either bound and exclude dates outside the range.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-date-inclusive-bounds.py", line 25, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: filter-date-since-only (FEATURE-FilterDate.md)
  intent: --since alone keeps records on or after the bound and nothing earlier.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-date-since-only.py", line 25, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: filter-date-inverted-range-exit-zero (FEATURE-FilterDate.md)
  intent: An inverted range (--since later than --until) matches nothing and exits 0.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-date-inverted-range-exit-zero.py", line 21, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: filter-target-exact-match (FEATURE-FilterTarget.md)
  intent: A record whose job.target equals the --target value is kept and a record differing only in case or trailing space is not.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-exact-match.py", line 31, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: filter-target-absent-keeps-all (FEATURE-FilterTarget.md)
  intent: Without --target every valid record is kept by the summary report.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-absent-keeps-all.py", line 23, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: filter-target-no-match-exit-zero (FEATURE-FilterTarget.md)
  intent: A --target value matching no record yields a zero-count summary with exit 0.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "filter-target-no-match-exit-zero.py", line 20, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: daily-report-whole-sample (FEATURE-DailyReport.md)
  intent: The daily report over the full sample matches the supplied known answer for every date and the TOTAL row, parsed as CSV.
  return code: 1
  stdout:
    error: report "daily" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "daily" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "daily-report-whole-sample.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "daily" is not implemented\n')
- RED: daily-report-since-until (FEATURE-DailyReport.md)
  intent: With --since and --until both set to 2026-10-15 the daily report contains only that date and its TOTAL row, parsed as CSV.
  return code: 1
  stdout:
    error: report "daily" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "daily" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "daily-report-since-until.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "daily" is not implemented\n')
- RED: daily-report-zero-match (FEATURE-DailyReport.md)
  intent: A filter matching no record writes the header and a zero TOTAL row with exit 0.
  return code: 1
  stdout:
    error: report "daily" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "daily" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "daily-report-zero-match.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "daily" is not implemented\n')
- RED: failures-report-whole-sample (FEATURE-FailuresReport.md)
  intent: The failures report over the full sample matches the two known failed records, parsed as CSV.
  return code: 1
  stdout:
    error: report "failures" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "failures" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "failures-report-whole-sample.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "failures" is not implemented\n')
- RED: failures-report-target-filter (FEATURE-FailuresReport.md)
  intent: With --target CommonMark the failures report contains only the CommonMark failed record, with an empty returncode and timed_out true, parsed as CSV.
  return code: 1
  stdout:
    error: report "failures" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "failures" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "failures-report-target-filter.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "failures" is not implemented\n')
- RED: failures-report-zero-match (FEATURE-FailuresReport.md)
  intent: A filter matching no failed record writes only the header with exit 0.
  return code: 1
  stdout:
    error: report "failures" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "failures" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "failures-report-zero-match.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "failures" is not implemented\n')
- RED: targets-ordering-and-total (FEATURE-TargetsReport.md)
  intent: Targets are ordered ascending by byte value and TOTAL sums the rows.
  return code: 1
  stdout:
    error: report "targets" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "targets" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-ordering-and-total.py", line 13, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "targets" is not implemented\n')
- RED: targets-zero-match (FEATURE-TargetsReport.md)
  intent: A filter matching nothing yields the header and a zero TOTAL row with cost and minutes, exit 0.
  return code: 1
  stdout:
    error: report "targets" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "targets" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-zero-match.py", line 12, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "targets" is not implemented\n')
- RED: skipped-ignores-target (FEATURE-SkippedCount.md)
  intent: The skipped count is identical with and without a --target filter.
  return code: 1
  stdout:
    error: report "summary" is not implemented
    
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code_all = 2
      out_all = ''
      err_all = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skipped-ignores-target.py", line 15, in <module>
        assert code_all == 0, (code_all, out_all, err_all)
               ^^^^^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: skipped-ignores-date (FEATURE-SkippedCount.md)
  intent: The skipped count is identical with and without date filters.
  return code: 1
  stdout:
    error: report "summary" is not implemented
    
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code_all = 2
      out_all = ''
      err_all = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "skipped-ignores-date.py", line 15, in <module>
        assert code_all == 0, (code_all, out_all, err_all)
               ^^^^^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: models-report-whole-sample (FEATURE-ModelsReport.md)
  intent: The models report over the full sample matches the supplied known answer for every pair and the TOTAL row, parsed as CSV.
  return code: 1
  stdout:
    error: report "models" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "models" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-report-whole-sample.py", line 13, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "models" is not implemented\n')
- RED: models-report-target-filter (FEATURE-ModelsReport.md)
  intent: With --target ReadingList the models report contains only the ReadingList claude-haiku-5-5 pair and its TOTAL row, parsed as CSV.
  return code: 1
  stdout:
    error: report "models" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "models" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-report-target-filter.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "models" is not implemented\n')
- RED: models-report-zero-match (FEATURE-ModelsReport.md)
  intent: A filter matching no record writes the header and a zero TOTAL row with exit 0.
  return code: 1
  stdout:
    error: report "models" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "models" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-report-zero-match.py", line 11, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "models" is not implemented\n')
- RED: summary-header-and-row-shape (FEATURE-SummaryReport.md)
  intent: The summary report writes the exact header and one data row with eleven fields.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-header-and-row-shape.py", line 12, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: summary-values-from-fixture (FEATURE-SummaryReport.md)
  intent: The summary report aggregates a known fixture into the expected counts, dates, sums, cost, and minutes.
  return code: 1
  stdout:
    error: report "summary" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "summary" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-values-from-fixture.py", line 20, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "summary" is not implemented\n')
- RED: commands-ordering-and-total (FEATURE-CommandsReport.md)
  intent: Commands are ordered by calls descending then name ascending, and TOTAL sums the rows.
  return code: 1
  stdout:
    error: report "commands" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "commands" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-ordering-and-total.py", line 14, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "commands" is not implemented\n')
- RED: commands-tie-byte-order (FEATURE-CommandsReport.md)
  intent: Commands with equal call counts are ordered by command name in byte order.
  return code: 1
  stdout:
    error: report "commands" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "commands" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-tie-byte-order.py", line 13, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "commands" is not implemented\n')
- RED: commands-zero-match (FEATURE-CommandsReport.md)
  intent: A filter matching nothing yields the header and a zero TOTAL row, exit 0.
  return code: 1
  stdout:
    error: report "commands" is not implemented
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'error: report "commands" is not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-zero-match.py", line 12, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'error: report "commands" is not implemented\n')

## Post-build programmatic acceptance
- PASS: entrypoint-builds (FEATURE-Entrypoint.md)
  intent: The ddkstat main package builds with go build ./cmd/ddkstat.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: entrypoint-no-args-exit-2 (FEATURE-Entrypoint.md)
  intent: Invoking ddkstat with no arguments exits 2 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: missing report name
- PASS: parse-summary-with-target (FEATURE-CommandParse.md)
  intent: `summary --target X <log>` parses the report, the option value, and the log path.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: parse-repeated-option-exit-2 (FEATURE-CommandParse.md)
  intent: A repeated --target option is a usage error with exit 2 and empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: invalid value "B" for flag -target: option given more than once
- PASS: usage-unknown-report-exit-2 (FEATURE-UsageErrors.md)
  intent: An unknown report name exits 2 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: unknown report "frobnicate"
- PASS: usage-no-log-path-exit-2 (FEATURE-UsageErrors.md)
  intent: A report with no log path exits 2 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: expected exactly one log path, got 0
- PASS: csv-header-and-lf (FEATURE-FormatCSV.md)
  intent: The CSV writer emits the header first and uses LF-only line endings.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: csv-zero-match-total (FEATURE-FormatCSV.md)
  intent: A zero-match commands report writes its header and a zero TOTAL row, and exits 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: fmt-cost-four-places (FEATURE-FormatCost.md)
  intent: The summed cost renders with exactly four decimal places in the summary report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,1,0,1,0.5000,0.0
- PASS: fmt-cost-zero (FEATURE-FormatCost.md)
  intent: A zero cost total renders as 0.0000 in the summary report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,1,0,1,0.0000,0.0
- PASS: fmt-minutes-example (FEATURE-FormatMinutes.md)
  intent: The minutes formatter renders 95000 ms as 1.6 by observing a summary report that sums one record of that duration.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,1,0,1,0.1000,1.6
- PASS: fmt-minutes-zero (FEATURE-FormatMinutes.md)
  intent: A group with zero elapsed time renders minutes as 0.0 in the summary report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,1,0,1,0.1000,0.0
- FAIL: log-read-blank-lines (FEATURE-LogRead.md)
  intent: Blank lines are ignored by the line reader and do not appear in the yielded lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    0,2,0,0,,,0,0,0,0.0000,0.0
    
    warning: skipped line 1: missing or empty started_at
    warning: skipped line 4: missing or empty started_at
  stderr:
    --- drydock: values at failure ---
      out = 'records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes\n0,2,0,0,,,0,0,0,0.0000,0.0\n'
      fields = ['0', '2', '0', '0', '', '', '0', '0', '0', '0.0000', '0.0']
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-read-blank-lines.py", line 19, in <module>
        assert fields[0] == "2", out
               ^^^^^^^^^^^^^^^^
    AssertionError: records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    0,2,0,0,,,0,0,0,0.0000,0.0
- PASS: log-read-utf8 (FEATURE-LogRead.md)
  intent: The line reader reads the file as UTF-8 and a multibyte character inside a valid record is kept intact.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
    plan é,1,0,1,0,1,0.0
    TOTAL,1,0,1,0,1,0.0
- PASS: log-read-missing-file (FEATURE-LogRead.md)
  intent: A log path that cannot be opened exits 1 with nothing on stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: cannot read log: open /tmp/tmpsglh0to6/no-such-log.jsonl: no such file or directory
- PASS: read-missing-log-exit-1 (FEATURE-ReadErrors.md)
  intent: A log path that does not exist exits 1 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: cannot read log: open /tmp/tmph3_5qw32/no-such-log.jsonl: no such file or directory
- PASS: filter-validate-rejects-impossible-date (FEATURE-FilterValidate.md)
  intent: An impossible calendar date in --since exits 2 with nothing on stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: invalid date "2026-13-01": want a real calendar date as YYYY-MM-DD
- PASS: filter-validate-rejects-invalid-day (FEATURE-FilterValidate.md)
  intent: A day that does not exist in its month (February 30) is a usage error with exit 2.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: invalid date "2026-02-30": want a real calendar date as YYYY-MM-DD
- PASS: filter-validate-rejects-date-after-log (FEATURE-FilterValidate.md)
  intent: An invalid date that appears after the log path is still a usage error with exit 2.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    error: invalid date "2026-13-01": want a real calendar date as YYYY-MM-DD
- PASS: log-skip-counts-and-exit (FEATURE-LogSkip.md)
  intent: Non-object, incomplete, and malformed lines are skipped, counted in the summary, and do not change exit status.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,3,1,0,2026-10-14,2026-10-14,1,0,1,0.1000,0.0
    
    warning: skipped line 2: not a JSON object
    warning: skipped line 3: not a JSON object
    warning: skipped line 4: missing or empty job.target
- PASS: log-skip-duplicate-first-kept (FEATURE-LogSkip.md)
  intent: A duplicate execution_id is skipped on its later occurrence and the first occurrence is kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,1,1,0,2026-10-14,2026-10-14,1,0,1,0.1000,0.0
    
    warning: skipped line 2: duplicate execution_id d-1
- PASS: log-skip-warning-per-line (FEATURE-LogSkip.md)
  intent: Each skipped line writes exactly one stderr warning naming its line number, and nothing is written to stdout for skips.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,1,1,0,2026-10-14,2026-10-14,1,0,1,0.1000,0.0
    
    warning: skipped line 1: not a JSON object
- PASS: log-model-prefers-executed (FEATURE-LogModel.md)
  intent: A record with result.stats.model reports that executed model in the models report rather than job.model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
    claude,executed-model,1,0,1,0,1,0.1000
    TOTAL,,1,0,1,0,1,0.1000
- PASS: log-model-falls-back-to-requested (FEATURE-LogModel.md)
  intent: A record without result.stats.model reports its job.model in the models report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
    claude,requested-model,1,0,1,0,1,0.1000
    TOTAL,,1,0,1,0,1,0.1000
- PASS: log-numbers-null-and-absent-are-zero (FEATURE-LogNumbers.md)
  intent: Null and absent numeric stats count as zero in the summary totals and the record is still counted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,5,0,7,0.0000,0.0
- PASS: log-numbers-no-stats-object (FEATURE-LogNumbers.md)
  intent: A record with no result.stats object is still counted as a record with zero numeric totals.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,0,0,0,0.0000,0.0
- PASS: log-failed-status-classification (FEATURE-LogFailed.md)
  intent: A record whose status is not succeeded is counted as failed in the summary, and a succeeded record is not.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    2,0,1,1,2026-10-14,2026-10-14,2,0,2,0.2000,0.0
- PASS: log-failed-returncode-and-timeout (FEATURE-LogFailed.md)
  intent: The failures report shows an empty returncode for a null return code and true for timed_out.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    2026-10-14T06:38:54.176Z,T,diagnose,sonnet,,true
- PASS: filter-date-inclusive-bounds (FEATURE-FilterDate.md)
  intent: --since and --until keep records whose start date equals either bound and exclude dates outside the range.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    3,0,3,0,2026-10-14,2026-10-16,3,0,3,0.3000,0.1
- PASS: filter-date-since-only (FEATURE-FilterDate.md)
  intent: --since alone keeps records on or after the bound and nothing earlier.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    2,0,2,0,2026-10-14,2026-10-16,2,0,2,0.2000,0.0
- PASS: filter-date-inverted-range-exit-zero (FEATURE-FilterDate.md)
  intent: An inverted range (--since later than --until) matches nothing and exits 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    0,0,0,0,,,0,0,0,0.0000,0.0
- PASS: filter-target-exact-match (FEATURE-FilterTarget.md)
  intent: A record whose job.target equals the --target value is kept and a record differing only in case or trailing space is not.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,1,0,1,0.1000,0.0
- PASS: filter-target-absent-keeps-all (FEATURE-FilterTarget.md)
  intent: Without --target every valid record is kept by the summary report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    3,0,3,0,2026-10-14,2026-10-14,3,0,3,0.3000,0.1
- PASS: filter-target-no-match-exit-zero (FEATURE-FilterTarget.md)
  intent: A --target value matching no record yields a zero-count summary with exit 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    0,0,0,0,,,0,0,0,0.0000,0.0
- PASS: daily-report-whole-sample (FEATURE-DailyReport.md)
  intent: The daily report over the full sample matches the supplied known answer for every date and the TOTAL row, parsed as CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    2026-10-14,14,1,6944443,86.3
    2026-10-15,12,1,7016674,57.8
    2026-10-16,7,0,3753306,15.3
    TOTAL,33,2,17714423,159.4
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: daily-report-since-until (FEATURE-DailyReport.md)
  intent: With --since and --until both set to 2026-10-15 the daily report contains only that date and its TOTAL row, parsed as CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    2026-10-15,12,1,7016674,57.8
    TOTAL,12,1,7016674,57.8
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: daily-report-zero-match (FEATURE-DailyReport.md)
  intent: A filter matching no record writes the header and a zero TOTAL row with exit 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    TOTAL,0,0,0,0.0
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: failures-report-whole-sample (FEATURE-FailuresReport.md)
  intent: The failures report over the full sample matches the two known failed records, parsed as CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    2026-10-14T02:38:41.904Z,Toml,diagnose,gemini-3-flash,1,false
    2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: failures-report-target-filter (FEATURE-FailuresReport.md)
  intent: With --target CommonMark the failures report contains only the CommonMark failed record, with an empty returncode and timed_out true, parsed as CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: failures-report-zero-match (FEATURE-FailuresReport.md)
  intent: A filter matching no failed record writes only the header with exit 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: targets-ordering-and-total (FEATURE-TargetsReport.md)
  intent: Targets are ordered ascending by byte value and TOTAL sums the rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: targets-zero-match (FEATURE-TargetsReport.md)
  intent: A filter matching nothing yields the header and a zero TOTAL row with cost and minutes, exit 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: skipped-ignores-target (FEATURE-SkippedCount.md)
  intent: The skipped count is identical with and without a --target filter.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    warning: skipped line 2: not a JSON object
    
    warning: skipped line 2: not a JSON object
- PASS: skipped-ignores-date (FEATURE-SkippedCount.md)
  intent: The skipped count is identical with and without date filters.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    warning: skipped line 2: not a JSON object
    
    warning: skipped line 2: not a JSON object
- PASS: models-report-whole-sample (FEATURE-ModelsReport.md)
  intent: The models report over the full sample matches the supplied known answer for every pair and the TOTAL row, parsed as CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
    claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
    claude,claude-sonnet-5-5,11,1,5547,5786766,119677,3.7919
    claude,sonnet,1,0,0,0,0,0.0000
    codex,gpt-5.6-luna,4,0,110058,2466012,40962,0.0000
    gemini,gemini-3-flash,6,1,142817,3639280,92277,0.4841
    TOTAL,,33,2,262526,17053317,398580,5.5245
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: models-report-target-filter (FEATURE-ModelsReport.md)
  intent: With --target ReadingList the models report contains only the ReadingList claude-haiku-5-5 pair and its TOTAL row, parsed as CSV.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
    claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
    TOTAL,,11,0,4104,5161259,145664,1.2485
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: models-report-zero-match (FEATURE-ModelsReport.md)
  intent: A filter matching no record writes the header and a zero TOTAL row with exit 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
    TOTAL,,0,0,0,0,0,0.0000
    
    warning: skipped line 16: not a JSON object
    warning: skipped line 28: duplicate execution_id 20261015.083525.332Z-8c438a77
    warning: skipped line 34: missing or empty job.target
- PASS: summary-header-and-row-shape (FEATURE-SummaryReport.md)
  intent: The summary report writes the exact header and one data row with eleven fields.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: summary-values-from-fixture (FEATURE-SummaryReport.md)
  intent: The summary report aggregates a known fixture into the expected counts, dates, sums, cost, and minutes.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: commands-ordering-and-total (FEATURE-CommandsReport.md)
  intent: Commands are ordered by calls descending then name ascending, and TOTAL sums the rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: commands-tie-byte-order (FEATURE-CommandsReport.md)
  intent: Commands with equal call counts are ordered by command name in byte order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: commands-zero-match (FEATURE-CommandsReport.md)
  intent: A filter matching nothing yields the header and a zero TOTAL row, exit 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Repair attempts
- attempt 0 (initial build): failed; 51/52 checks model=haiku; execution 20261024.165014.823Z-b47df7b4; reason: programmatic acceptance failed: log-read-blank-lines
- attempt 1 (repair 1): failed; 51/52 checks model=haiku; execution 20261024.165553.280Z-be6617fc; reason: programmatic acceptance failed: log-read-blank-lines
- attempt 2 (repair 2): failed; 51/52 checks model=haiku; execution 20261024.165821.563Z-d19d2cff; stopped: deterministic acceptance score did not improve on 2 consecutive calls; reason: programmatic acceptance failed: log-read-blank-lines

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: AC log-read-blank-lines contradicts the log-skip required-field rule and the run.sh sample.
- detail:
    The fixture lines carry only `execution_id` but the criterion expects `records=2, skipped=0`. `FEATURE-LogSkip.md`, `log-skip-counts-and-exit` and `run.sh` A1 all require skipping records missing `started_at`, `status`, `job.command_name` or `job.target`. The command exits 0 and the only failing assertion is `fields[0] == "2"`. To fix it, complete the fixture or loosen the required-fields rule at replan.

## Failure
- summary: programmatic acceptance failed: log-read-blank-lines
- detail:
    Block "Block 2 · Service" [block-2] failed its acceptance criteria.
      Story "Read JSON Lines log records from a file path." [log-read] does not meet its own acceptance criteria:
        - AC log-read-blank-lines — Blank lines are ignored by the line reader and do not appear in the yielded lines.
            assertion: assert fields[0] == "2", out → AssertionError: records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
            raised at: log-read-blank-lines.py:19
            process exit code: 1
            error: AssertionError: records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
            values at failure:
              out = 'records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes\n0,2,0,0,,,0,0,0,0.0000,0.0\n'
              fields = ['0', '2', '0', '0', '', '', '0', '0', '0', '0.0000', '0.0']
            observed output:
              records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
              0,2,0,0,,,0,0,0,0.0000,0.0
              warning: skipped line 1: missing or empty started_at
              warning: skipped line 4: missing or empty started_at
            check stderr:
              Traceback (most recent call last):
                File "log-read-blank-lines.py", line 19, in <module>
                  assert fields[0] == "2", out
                         ^^^^^^^^^^^^^^^^
              AssertionError: records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
              0,2,0,0,,,0,0,0,0.0000,0.0

## Build summary
`log-read-blank-lines` still fails, and I made no code changes this pass. The criterion contradicts the log-skip rules, so I'm reporting it as `AC_BROKEN` again. Everything else is green: `sh bin/test.sh` passes and `sh sources/run.sh` passes 20 of 20.

I re-ran the criterion's exact input (two records with only `execution_id`, separated by blank lines). The program exits 0 and ignores the blank lines, then prints `0,2,0,0,...`. Only the `fields[0] == "2"` assertion fails.

The criterion expects those two records to be kept, but three other sources require skipping them:
- **`FEATURE-LogSkip.md`:** a line is skipped when `started_at`, `status`, `job.command_name` or `job.target` is missing or empty.
- **`log-skip-counts-and-exit`:** it expects a record missing only `job.target` to be skipped.
- **`sources/run.sh` A1/C2:** the sample's record with no `target` must count in `skipped=3`.

Keeping `{"execution_id":"a"}` would break all three. Only a hack that special-cases this input could satisfy both, which the repair rules forbid.

AC_BROKEN: log-read-blank-lines

RESULT: FAILED

FILES CHANGED:
- (none in this pass; the files from earlier passes are unchanged)

SUMMARY:
- No changes this pass.
- All other declared checks stay green, and the implementation follows the log-skip rules.

BLOCKERS:
- `log-read-blank-lines` needs complete fixture records (add `started_at`, `status` and `job`), or the required-field rule in `FEATURE-LogSkip.md` needs to change at replan.

FAILURE_SUMMARY: AC log-read-blank-lines contradicts the log-skip required-field rule and the run.sh sample.
FAILURE_DETAIL: The fixture lines carry only `execution_id` but the criterion expects `records=2, skipped=0`. `FEATURE-LogSkip.md`, `log-skip-counts-and-exit` and `run.sh` A1 all require skipping records missing `started_at`, `status`, `job.command_name` or `job.target`. The command exits 0 and the only failing assertion is `fields[0] == "2"`. To fix it, complete the fixture or loosen the required-fields rule at replan.
