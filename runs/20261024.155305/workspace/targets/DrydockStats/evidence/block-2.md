# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-24
- resulting state: closed/failed
- story points (combined assembled cost): 29325
- execution id: 20261024.161238.365Z-e58fe633

## Stories built
- Prove the module builds the ddkstat executable with go build using only the standard library. (build-entry) [story]
- Decode the JSON Lines log, one JSON object per non-blank line. (log-decode) [story]
- Skip records missing a required field, with a line-numbered warning. (log-required-fields) [story]
- Skip repeated execution identifiers, keeping the first occurrence. (log-duplicates) [story]
- Provide the shared CSV writer and the minutes and cost formatters. (report-format) [story]
- Keep only records whose job.target equals the --target name exactly. (filter-target) [story]
- Treat absent or null token, cost, and duration values as zero. (log-numeric) [story]
- Attribute each record to its executed model, falling back to the requested model. (log-model) [story]
- Keep records whose start date falls within the inclusive --since and --until range. (filter-date) [story]
- Produce the per-Target report with a TOTAL row. (report-targets) [story]
- Produce the per-date report with a TOTAL row. (report-daily) [story]
- Produce the per-command report with a TOTAL row. (report-commands) [story]
- Produce the per-provider and model report with a TOTAL row. (report-models) [story]
- Produce the one-row summary report. (report-summary) [story]
- Produce the failed-record report. (report-failures) [story]
- Make every report write its header and a zero TOTAL row when no record matches. (report-empty-match) [story]
- Parse the report name, options, and the single log path, and wire each report. (cli-parse) [story]
- Exit 1 with no stdout when the log cannot be opened or read. (cli-unreadable-log) [story]
- Map usage errors to exit code 2 with no stdout output. (cli-usage-errors) [story]
- Exit 0 when lines are skipped or no record matches. (cli-success-partial) [story]
- Provide the POSIX bin/test.sh that runs go vet and go test over the whole module. (test-runner) [story]

## Acceptance tooling authorization
- FEATURE-BuildEntry.md#build-entry-builds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogDecode.md#log-decode-ignores-blank: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogRequiredFields.md#log-required-fields-skip: executable=go; scope=test; authorization=existing Target environment
- FEATURE-LogDuplicates.md#log-duplicates-first-kept: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ReportFormat.md#report-format-minutes-rounding: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ReportFormat.md#report-format-cost-four-places: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-exact-match: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterTarget.md#filter-target-absent-passes-all: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogNumeric.md#log-numeric-null-cost-is-zero: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogNumeric.md#log-numeric-absent-tokens-are-zero: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogNumeric.md#log-numeric-missing-stats-object: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogModel.md#log-model-prefers-executed: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogModel.md#log-model-falls-back-to-requested: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-LogModel.md#log-model-absent-provider: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterDate.md#filter-date-inclusive-range: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-FilterDate.md#filter-date-open-ended: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ReportFailures.md#report-failures-header: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ReportFailures.md#report-failures-rows: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ReportEmptyMatch.md#empty-match-commands: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-ReportEmptyMatch.md#empty-match-summary: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CliParse.md#cli-parse-dispatch-summary: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CliParse.md#cli-parse-target-option: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CliUnreadableLog.md#unreadable-missing-log: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CliUnreadableLog.md#unreadable-directory-log: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CliUsageErrors.md#usage-unknown-report: executable=python3; scope=test; authorization=existing Target environment
- FEATURE-CliUsageErrors.md#usage-impossible-date: executable=python3; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-BuildEntry.md (SP 475)
- context: ARCHITECTURE.md (SP 769)
- implements: FEATURE-LogDecode.md (SP 280)
- implements: FEATURE-LogRequiredFields.md (SP 285)
- implements: FEATURE-LogDuplicates.md (SP 294)
- implements: FEATURE-ReportFormat.md (SP 620)
- implements: FEATURE-FilterTarget.md (SP 674)
- implements: FEATURE-LogNumeric.md (SP 822)
- implements: FEATURE-LogModel.md (SP 1146)
- implements: FEATURE-FilterDate.md (SP 1015)
- implements: FEATURE-ReportTargets.md (SP 1087)
- implements: FEATURE-ReportDaily.md (SP 1189)
- implements: FEATURE-ReportCommands.md (SP 1145)
- implements: FEATURE-ReportModels.md (SP 1238)
- implements: FEATURE-ReportSummary.md (SP 1660)
- implements: FEATURE-ReportFailures.md (SP 1766)
- implements: FEATURE-ReportEmptyMatch.md (SP 1653)
- implements: FEATURE-CliParse.md (SP 1627)
- implements: FEATURE-CliUnreadableLog.md (SP 876)
- implements: FEATURE-CliUsageErrors.md (SP 944)
- implements: FEATURE-CliSuccessPartial.md (SP 1686)
- implements: FEATURE-TestRunner.md (SP 920)
- stack: common_compact.md (SP 1653)

## Build directory changes
- bin/test.sh
- ddkstat_build_probe
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/logrecord/read.go
- internal/logrecord/read_test.go
- internal/logrecord/record.go
- internal/report/filter.go
- internal/report/format.go
- internal/report/report_test.go
- internal/report/reports.go

## Pre-build acceptance observation
- GREEN (prepassed): build-entry-builds (FEATURE-BuildEntry.md)
  intent: go build produces the ddkstat executable from the application root.
  return code: 0
- GREEN (prepassed): build-entry-stdlib-only (FEATURE-BuildEntry.md)
  intent: The module declares no requirements, so every import resolves to the standard library.
  return code: 0
- GREEN (vacuous): log-decode-ignores-blank (FEATURE-LogDecode.md)
  intent: Blank lines are ignored and not counted as skipped.
  return code: 0
  note: assertion on constant literal True is always true
  stdout:
    {"execution_id":"a"}
- RED: log-required-fields-skip (FEATURE-LogRequiredFields.md)
  intent: A record missing job.target is counted as skipped and does not appear in the records.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('internal/logrecord')
      mode = 'r'
      buffering = -1
      encoding = 'locale'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-required-fields-skip.py", line 3, in <module>
        src = Path("internal/logrecord").read_text()
              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'internal/logrecord'
- RED: log-duplicates-first-kept (FEATURE-LogDuplicates.md)
  intent: The first occurrence of a duplicate execution_id is kept and later occurrences are skipped.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "log-duplicates-first-kept.py", line 3, in <module>
        assert Path("internal/logrecord").is_dir()
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): report-format-minutes-rounding (FEATURE-ReportFormat.md)
  intent: Minutes are written as tenths of a minute, rounded half up, so 95000 ms writes 1.6.
  return code: 0
- GREEN (prepassed): report-format-cost-four-places (FEATURE-ReportFormat.md)
  intent: Cost is written with exactly four decimal places.
  return code: 0
- GREEN (prepassed): filter-target-exact-match (FEATURE-FilterTarget.md)
  intent: Only records whose job.target equals the requested name exactly pass the filter.
  return code: 0
- GREEN (prepassed): filter-target-absent-passes-all (FEATURE-FilterTarget.md)
  intent: When no target is given every record passes the filter.
  return code: 0
- GREEN (prepassed): log-numeric-null-cost-is-zero (FEATURE-LogNumeric.md)
  intent: A cost_usd of null contributes zero to cost totals.
  return code: 0
- GREEN (prepassed): log-numeric-absent-tokens-are-zero (FEATURE-LogNumeric.md)
  intent: Absent token and duration fields are read as zero.
  return code: 0
- GREEN (prepassed): log-numeric-missing-stats-object (FEATURE-LogNumeric.md)
  intent: A record with no result.stats object is still a record with all-zero numeric stats.
  return code: 0
- GREEN (prepassed): log-model-prefers-executed (FEATURE-LogModel.md)
  intent: A record with a non-empty result.stats.model is attributed to that executed model.
  return code: 0
  stdout:
    placeholder
- GREEN (prepassed): log-model-falls-back-to-requested (FEATURE-LogModel.md)
  intent: A record whose result.stats.model is absent or empty is attributed to job.model.
  return code: 0
- GREEN (prepassed): log-model-absent-provider (FEATURE-LogModel.md)
  intent: A record with no job.llm is attributed an empty provider.
  return code: 0
- GREEN (prepassed): filter-date-inclusive-range (FEATURE-FilterDate.md)
  intent: Records on the since and until boundaries are kept; records outside the range are dropped.
  return code: 0
- GREEN (prepassed): filter-date-open-ended (FEATURE-FilterDate.md)
  intent: Omitting one bound leaves the range open on that side.
  return code: 0
- RED: targets-ordering-and-total (FEATURE-ReportTargets.md)
  intent: Targets rows order by target ascending byte-wise and the TOTAL row sums all filtered records.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "targets-ordering-and-total.py", line 30, in <module>
        assert rows[0] == ["target", "calls", "failed", "input_tokens", "cached_tokens",
               ~~~~^^^
    IndexError: list index out of range
- RED: daily-ordering-and-total (FEATURE-ReportDaily.md)
  intent: Daily rows group by start date ascending, tokens sum input plus cached plus output, and the TOTAL row covers all filtered records.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "daily-ordering-and-total.py", line 30, in <module>
        assert rows[0] == ["date", "calls", "failed", "tokens", "minutes"]
               ~~~~^^^
    IndexError: list index out of range
- RED: commands-ordering-and-total (FEATURE-ReportCommands.md)
  intent: Commands rows order by calls descending then byte-wise command ascending, and the TOTAL row sums all filtered records.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "commands-ordering-and-total.py", line 32, in <module>
        assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens",
               ~~~~^^^
    IndexError: list index out of range
- RED: models-ordering-and-total (FEATURE-ReportModels.md)
  intent: Models rows group by provider and attributed model, ordered byte-wise, with the TOTAL row leaving the model column empty.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "models-ordering-and-total.py", line 32, in <module>
        assert rows[0] == ["provider", "model", "calls", "failed", "input_tokens",
               ~~~~^^^
    IndexError: list index out of range
- RED: summary-empty-file (FEATURE-ReportSummary.md)
  intent: A log with no records yields one data row with zero counts and empty first and last.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-empty-file.py", line 18, in <module>
        assert rows[0] == ["records", "skipped", "succeeded", "failed", "first", "last",
               ~~~~^^^
    IndexError: list index out of range
- RED: summary-filters-and-skipped (FEATURE-ReportSummary.md)
  intent: Filtered counts exclude non-matching records while skipped counts the whole file; the first and last dates are the earliest and latest filtered start dates.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summary-filters-and-skipped.py", line 24, in <module>
        row = rows[1]
              ~~~~^^^
    IndexError: list index out of range
- RED: report-failures-header (FEATURE-ReportFailures.md)
  intent: The failures report writes the declared header on an empty match.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      expected_header = ['started_at', 'target', 'command', 'model', 'returncode', 'timed_out']
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-failures-header.py", line 23, in <module>
        assert rows[0] == expected_header
               ~~~~^^^
    IndexError: list index out of range
- RED: report-failures-rows (FEATURE-ReportFailures.md)
  intent: Each failed record yields one row with empty returncode for null and literal timed_out.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      r = {'schema_version': 1, 'execution_id': 'c', 'started_at': '2026-10-16T00:00:00.000Z', 'status': 'succeeded', 'job': {'command_name': 'build', 'target': 'Toml', 'llm': 'claude', 'model': 'sonnet'}, 'result': {'returncode': 0, 'timed_out': False}}
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "report-failures-rows.py", line 35, in <module>
        assert [r["execution_id"] if "execution_id" in r else r["started_at"] for r in rows] == [
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: empty-match-commands (FEATURE-ReportEmptyMatch.md)
  intent: The commands report writes header and zero TOTAL row when a filter matches nothing.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "empty-match-commands.py", line 27, in <module>
        assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"]
               ~~~~^^^
    IndexError: list index out of range
- RED: empty-match-summary (FEATURE-ReportEmptyMatch.md)
  intent: The summary report on an empty match writes empty first and last and zero totals.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "empty-match-summary.py", line 26, in <module>
        assert len(rows) == 1
               ^^^^^^^^^^^^^^
    AssertionError
- RED: cli-parse-dispatch-summary (FEATURE-CliParse.md)
  intent: A valid summary invocation with one log path exits 0 and writes the summary header.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-parse-dispatch-summary.py", line 24, in <module>
        assert rows[0][0] == "records"
               ~~~~^^^
    IndexError: list index out of range
- RED: cli-parse-target-option (FEATURE-CliParse.md)
  intent: The --target option is passed to the filter so only the named target's records are counted.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-parse-target-option.py", line 30, in <module>
        assert rows[0]["records"] == "1"
               ~~~~^^^
    IndexError: list index out of range
- RED: unreadable-missing-log (FEATURE-CliUnreadableLog.md)
  intent: A log path that does not exist exits 1 with empty stdout and one stderr line.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['/tmp/tmp0f68f7jj/ddkstat', 'summary', '/tmp/tmp0f68f7jj/no-such-log.jsonl'], returncode=0, stdout='', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "unreadable-missing-log.py", line 14, in <module>
        assert result.returncode == 1
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: unreadable-directory-log (FEATURE-CliUnreadableLog.md)
  intent: A log path that names a directory exits 1 with empty stdout.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['/tmp/tmpqzv_klnt/ddkstat', 'summary', '/tmp/tmpqzv_klnt'], returncode=0, stdout='', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "unreadable-directory-log.py", line 13, in <module>
        assert result.returncode == 1
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: usage-unknown-report (FEATURE-CliUsageErrors.md)
  intent: An unknown report name exits 2 with empty stdout.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['/tmp/tmpsz90ge74/ddkstat', 'frobnicate', '/dev/null'], returncode=0, stdout='', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-unknown-report.py", line 13, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: usage-impossible-date (FEATURE-CliUsageErrors.md)
  intent: A calendar-impossible --since date exits 2 with empty stdout.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['/tmp/tmp_2r_sk5_/ddkstat', 'summary', '--since', '2026-02-30', '/dev/null'], returncode=0, stdout='', stderr='')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-impossible-date.py", line 15, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-partial-skipped-exit-zero (FEATURE-CliSuccessPartial.md)
  intent: A log with malformed and duplicate lines exits 0 and its summary reports skipped lines.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-partial-skipped-exit-zero.py", line 32, in <module>
        assert len(rows) == 1
               ^^^^^^^^^^^^^^
    AssertionError
- RED: cli-partial-no-match-exit-zero (FEATURE-CliSuccessPartial.md)
  intent: A filter that matches no records exits 0 and the commands report writes its header with a zero TOTAL row.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-partial-no-match-exit-zero.py", line 27, in <module>
        assert rows[0] == ["command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"]
               ~~~~^^^
    IndexError: list index out of range
- GREEN (prepassed): test-runner-script-present (FEATURE-TestRunner.md)
  intent: bin/test.sh exists and is a POSIX sh script that sets errexit and nounset.
  return code: 0
- GREEN (prepassed): test-runner-passes-on-clean-module (FEATURE-TestRunner.md)
  intent: bin/test.sh exits 0 when go vet and go test both pass over the module.
  return code: 0
  stdout:
    ?   	drydockstats/cmd/ddkstat	[no test files]
    ok  	drydockstats/internal/cli	0.007s

## Post-build programmatic acceptance
- PASS: build-entry-builds (FEATURE-BuildEntry.md)
  intent: go build produces the ddkstat executable from the application root.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: build-entry-stdlib-only (FEATURE-BuildEntry.md)
  intent: The module declares no requirements, so every import resolves to the standard library.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- UNVERIFIED: log-decode-ignores-blank (FEATURE-LogDecode.md)
  intent: Blank lines are ignored and not counted as skipped.
  provisioning: not required
  error: excluded check: ALWAYS_PASSES vacuous
- FAIL: log-required-fields-skip (FEATURE-LogRequiredFields.md)
  intent: A record missing job.target is counted as skipped and does not appear in the records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('internal/logrecord')
      mode = 'r'
      buffering = -1
      encoding = 'locale'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "log-required-fields-skip.py", line 3, in <module>
        src = Path("internal/logrecord").read_text()
              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'
- PASS: log-duplicates-first-kept (FEATURE-LogDuplicates.md)
  intent: The first occurrence of a duplicate execution_id is kept and later occurrences are skipped.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: report-format-minutes-rounding (FEATURE-ReportFormat.md)
  intent: Minutes are written as tenths of a minute, rounded half up, so 95000 ms writes 1.6.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: report-format-cost-four-places (FEATURE-ReportFormat.md)
  intent: Cost is written with exactly four decimal places.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: filter-target-exact-match (FEATURE-FilterTarget.md)
  intent: Only records whose job.target equals the requested name exactly pass the filter.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: filter-target-absent-passes-all (FEATURE-FilterTarget.md)
  intent: When no target is given every record passes the filter.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: log-numeric-null-cost-is-zero (FEATURE-LogNumeric.md)
  intent: A cost_usd of null contributes zero to cost totals.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: log-numeric-absent-tokens-are-zero (FEATURE-LogNumeric.md)
  intent: Absent token and duration fields are read as zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: log-numeric-missing-stats-object (FEATURE-LogNumeric.md)
  intent: A record with no result.stats object is still a record with all-zero numeric stats.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: log-model-prefers-executed (FEATURE-LogModel.md)
  intent: A record with a non-empty result.stats.model is attributed to that executed model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    placeholder
- PASS: log-model-falls-back-to-requested (FEATURE-LogModel.md)
  intent: A record whose result.stats.model is absent or empty is attributed to job.model.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: log-model-absent-provider (FEATURE-LogModel.md)
  intent: A record with no job.llm is attributed an empty provider.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: filter-date-inclusive-range (FEATURE-FilterDate.md)
  intent: Records on the since and until boundaries are kept; records outside the range are dropped.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: filter-date-open-ended (FEATURE-FilterDate.md)
  intent: Omitting one bound leaves the range open on that side.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: targets-ordering-and-total (FEATURE-ReportTargets.md)
  intent: Targets rows order by target ascending byte-wise and the TOTAL row sums all filtered records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    Alpha,2,1,2,4,6,0.5000,0.0
    Zeta,1,0,1,2,3,0.2500,0.0
    TOTAL,3,1,3,6,9,0.7500,0.1
- PASS: daily-ordering-and-total (FEATURE-ReportDaily.md)
  intent: Daily rows group by start date ascending, tokens sum input plus cached plus output, and the TOTAL row covers all filtered records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    date,calls,failed,tokens,minutes
    2026-10-14,2,1,120,0.0
    2026-10-15,1,0,60,0.0
    TOTAL,3,1,180,0.1
- PASS: commands-ordering-and-total (FEATURE-ReportCommands.md)
  intent: Commands rows order by calls descending then byte-wise command ascending, and the TOTAL row sums all filtered records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
    analyze,2,0,2,4,6,0.0
    build,2,0,2,4,6,0.0
    plan,1,0,1,2,3,0.0
    TOTAL,5,0,5,10,15,0.1
- PASS: models-ordering-and-total (FEATURE-ReportModels.md)
  intent: Models rows group by provider and attributed model, ordered byte-wise, with the TOTAL row leaving the model column empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
    claude,claude-haiku-5-5,2,0,2,4,6,0.2000
    codex,gpt-x,1,0,1,2,3,0.1000
    TOTAL,,3,0,3,6,9,0.3000
- PASS: summary-empty-file (FEATURE-ReportSummary.md)
  intent: A log with no records yields one data row with zero counts and empty first and last.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    0,0,0,0,,,0,0,0,0.0000,0.0
- PASS: summary-filters-and-skipped (FEATURE-ReportSummary.md)
  intent: Filtered counts exclude non-matching records while skipped counts the whole file; the first and last dates are the earliest and latest filtered start dates.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,1,1,0,2026-10-14,2026-10-14,10,5,2,0.5000,1.0
  stderr:
    ddkstat: warning: line 3 skipped: invalid character 'o' in literal null (expecting 'u')
- PASS: report-failures-header (FEATURE-ReportFailures.md)
  intent: The failures report writes the declared header on an empty match.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
- PASS: report-failures-rows (FEATURE-ReportFailures.md)
  intent: Each failed record yields one row with empty returncode for null and literal timed_out.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    started_at,target,command,model,returncode,timed_out
    2026-10-14T02:00:00.000Z,Toml,diagnose,gemini-3-flash,1,false
    2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
- PASS: empty-match-commands (FEATURE-ReportEmptyMatch.md)
  intent: The commands report writes header and zero TOTAL row when a filter matches nothing.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
    TOTAL,0,0,0,0,0,0.0
- PASS: empty-match-summary (FEATURE-ReportEmptyMatch.md)
  intent: The summary report on an empty match writes empty first and last and zero totals.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    0,0,0,0,,,0,0,0,0.0000,0.0
- PASS: cli-parse-dispatch-summary (FEATURE-CliParse.md)
  intent: A valid summary invocation with one log path exits 0 and writes the summary header.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,0,0,0,0.0000,0.0
- PASS: cli-parse-target-option (FEATURE-CliParse.md)
  intent: The --target option is passed to the filter so only the named target's records are counted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,0,1,0,2026-10-14,2026-10-14,0,0,0,0.0000,0.0
- PASS: unreadable-missing-log (FEATURE-CliUnreadableLog.md)
  intent: A log path that does not exist exits 1 with empty stdout and one stderr line.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    ddkstat: cannot read log /tmp/tmpazbz2kp4/no-such-log.jsonl: open /tmp/tmpazbz2kp4/no-such-log.jsonl: no such file or directory
- PASS: unreadable-directory-log (FEATURE-CliUnreadableLog.md)
  intent: A log path that names a directory exits 1 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    ddkstat: cannot read log /tmp/tmpid8lcni6: read /tmp/tmpid8lcni6: is a directory
- PASS: usage-unknown-report (FEATURE-CliUsageErrors.md)
  intent: An unknown report name exits 2 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    ddkstat: unknown report "frobnicate"
- PASS: usage-impossible-date (FEATURE-CliUsageErrors.md)
  intent: A calendar-impossible --since date exits 2 with empty stdout.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stderr:
    ddkstat: invalid --since "2026-02-30": want YYYY-MM-DD
- PASS: cli-partial-skipped-exit-zero (FEATURE-CliSuccessPartial.md)
  intent: A log with malformed and duplicate lines exits 0 and its summary reports skipped lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
    1,2,1,0,2026-10-14,2026-10-14,0,0,0,0.5000,0.0
  stderr:
    ddkstat: warning: line 2 skipped: invalid character 'o' in literal null (expecting 'u')
    ddkstat: warning: line 3 skipped: duplicate execution_id "a"
- PASS: cli-partial-no-match-exit-zero (FEATURE-CliSuccessPartial.md)
  intent: A filter that matches no records exits 0 and the commands report writes its header with a zero TOTAL row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
    TOTAL,0,0,0,0,0,0.0
- PASS: test-runner-script-present (FEATURE-TestRunner.md)
  intent: bin/test.sh exists and is a POSIX sh script that sets errexit and nounset.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: test-runner-passes-on-clean-module (FEATURE-TestRunner.md)
  intent: bin/test.sh exits 0 when go vet and go test both pass over the module.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	drydockstats/cmd/ddkstat	[no test files]
    ok  	drydockstats/internal/cli	(cached)
    ok  	drydockstats/internal/logrecord	(cached)
    ok  	drydockstats/internal/report	(cached)

## Repair attempts
- attempt 0 (initial build): failed; 35/37 checks model=haiku; execution 20261024.160619.320Z-abf3502a; reason: programmatic acceptance failed: log-required-fields-skip
- attempt 1 (repair 1): failed; 35/37 checks model=haiku; execution 20261024.161107.584Z-e1bf8488; reason: programmatic acceptance failed: log-required-fields-skip
- attempt 2 (repair 2): failed; 35/37 checks model=haiku; execution 20261024.161238.365Z-e58fe633; stopped: deterministic acceptance score did not improve on 2 consecutive calls; reason: programmatic acceptance failed: log-required-fields-skip

## Failure
- summary: programmatic acceptance failed: log-required-fields-skip
- detail:
    Block "Block 2 · Service" [block-2] failed its acceptance criteria.
      Story "Skip records missing a required field, with a line-numbered warning." [log-required-fields] does not meet its own acceptance criteria:
        - AC log-required-fields-skip — A record missing job.target is counted as skipped and does not appear in the records.
            assertion: IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'
            raised at: pathlib.py:1013
            process exit code: 1
            error: IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'
            values at failure:
              self = PosixPath('internal/logrecord')
              mode = 'r'
              buffering = -1
              encoding = 'locale'
              errors = None
              newline = None
            check stderr:
              Traceback (most recent call last):
                File "log-required-fields-skip.py", line 3, in <module>
                  src = Path("internal/logrecord").read_text()
                        ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
                  with self.open(mode='r', encoding=encoding, errors=errors) as f:
                       ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
                  return io.open(self, mode, buffering, encoding, errors, newline)
                         ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
              IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'

## Build summary
The same two checks are failing, so I'll re-confirm the gate and the two assertions before reporting.I made no changes in this repair pass. The same two checks still fail, and neither can pass whatever the code does. `sh bin/test.sh` still exits 0: vet is clean and the `cli`, `logrecord` and `report` packages pass.

- **`log-required-fields-skip`:** The assertion runs `Path("internal/logrecord").read_text()`, and I re-ran it: it still raises `IsADirectoryError`. `internal/logrecord` must be a directory, because ARCHITECTURE places the package there and `log-duplicates-first-kept` asserts `is_dir()` on the same path. Missing `job.target` is covered by `TestReadSkipsRecordsMissingEachRequiredField` in `internal/logrecord/read_test.go`.
- **`log-decode-ignores-blank`:** Its body is only `assert True`, so the harness excludes it as vacuous. Blank-line handling is covered by `TestReadIgnoresBlankLinesWithoutCountingThem`.

`ddkstat_build_probe` shows up in the "files written" list this time. I didn't write it. The `cli-partial-*` assertions build into that path themselves, and it was already in the directory before this block.

RESULT: SUCCESS

FILES CHANGED:
- (none in this repair pass; the files from the first pass are unchanged)

SUMMARY:
Re-ran `sh bin/test.sh` (exit 0) and the failing directory-read assertion (`IsADirectoryError` again). No code changed because neither failing check depends on the code.

AC_BROKEN: log-required-fields-skip, log-decode-ignores-blank
FAILURE_DETAIL: `log-required-fields-skip` calls `read_text()` on the package directory `internal/logrecord`, so Python raises `IsADirectoryError` before any Go code runs. `log-decode-ignores-blank` is `assert True` and is excluded as vacuous. To fix the first, read a file inside the package, such as `internal/logrecord/record.go`. To fix the second, replace the body with a real check, such as running the built binary over a log with blank lines and reading the summary's `skipped` value.
