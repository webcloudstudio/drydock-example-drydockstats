# MANIFEST: DrydockStats
updated:     2026-10-24T16:41:03+00:00
state:       approved
blocks:      5

applied_specs: |
  ARCHITECTURE.md sha256=6dfac808a9c3ea79454bc3271ac63000f1596fadf3af180b25ec2760fb9c0ac7 commit=b50614b001955581ef21eaa1da0e16ddd7bdfe98 applied_by=block-1 applied_at=2026-10-24T16:46:35+00:00 build_sha256=abe378f0658aead68a3ca632859d7d151f2356319ef88a73d531098d7ba7c584

## story 1: Establish the Go module and project tree.
id:           foundation
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-1.md
summary:      Establish the Go module and project tree.
type:         foundational
kind:         capability
phase:        1
block:        1
implements:   ARCHITECTURE.md
instructions: |
  Create go.mod at the application root for Go 1.22 with no third-party requirements, plus the internal/ and bin/ directories. Document the module boundaries in ARCHITECTURE.md. Standard library only.
stack:        go.md
stack_mode:   builder
size:         10030
provides:     go.mod module root, internal/ tree, bin/ directory
acceptance:   yes
state: closed/implemented

## story 2: Provide the ddkstat entry point and the Run dispatcher.
id: cli-entry
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Provide the ddkstat entry point and the Run dispatcher.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Entrypoint.md
covers: BLD-001
accepts: st-003
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5202
provides: cmd/ddkstat main, cli.Run(args, stdout, stderr) int
instructions: Add cmd/ddkstat/main.go, which calls internal/cli Run with os.Args[1:] and exits with its returned code. Only main calls os.Exit. The application root must build with go build ./cmd/ddkstat.
acceptance: yes
depends: foundation
state: closed/failed

## story 3: Parse the report name, options, and the single log path.
id: cli-parse
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Parse the report name, options, and the single log path.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CommandParse.md
covers: CLI-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5501
provides: argument parser returning report, options, and log path
instructions: Take the report name first, then parse options with a local flag.FlagSet using flag.ContinueOnError, then require exactly one log path. Because flag stops at the first positional argument, parse options after the report name. Unknown reports and options are passed to cli-usage.
acceptance: yes
depends: cli-entry
state: closed/failed

## story 4: Reject usage errors with exit code 2.
id: cli-usage
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Reject usage errors with exit code 2.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-UsageErrors.md
covers: CLI-002
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5367
provides: usage-error path with exit 2 and empty stdout
consumes: argument parser
instructions: An unknown report, unknown option, missing option value, repeated option, or missing or extra log path exits 2 with nothing on stdout and one diagnostic line on stderr. An unrecognized -h or --help is an unknown option.
acceptance: yes
depends: cli-parse
state: closed/failed

## story 5: Write CSV with a header, LF line endings, and TOTAL rows.
id: fmt-csv
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Write CSV with a header, LF line endings, and TOTAL rows.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FormatCSV.md
covers: FMT-003
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5661
provides: CSV writer for report rows
instructions: Write the header row first, comma-separated, with \n line endings. A zero-match report writes its header and, where it has one, a zero TOTAL row.
acceptance: yes
depends: foundation
state: closed/failed

## story 6: Format summed cost with four decimal places.
id: fmt-cost
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Format summed cost with four decimal places.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FormatCost.md
covers: FMT-002
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5674
provides: FormatCost(usd) string
instructions: Write the summed cost_usd with %.4f. A zero total renders 0.0000.
acceptance: yes
depends: foundation
state: closed/failed

## story 7: Format minutes in tenths of a minute.
id: fmt-minutes
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Format minutes in tenths of a minute.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FormatMinutes.md
covers: FMT-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5755
provides: FormatMinutes(ms) string
instructions: Compute t = (ms + 3000) / 6000 in integer division and render it as t/10, a period, and t%10. 95000 ms renders 1.6. Zero renders 0.0.
acceptance: yes
depends: foundation
state: closed/failed

## story 8: Read JSON Lines log records from a file path.
id: log-read
finding: programmatic acceptance failed: log-read-blank-lines: 0,2,0,0,,,0,0,0,0.0000,0.0
evidence: evidence/block-2.md
summary: Read JSON Lines log records from a file path.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogRead.md
covers: LOG-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6043
provides: logfile line reader for a file path
instructions: Read the file as UTF-8, one line at a time. Blank lines are ignored and are not counted as skipped. Pass each non-blank line on as raw bytes for parsing.
acceptance: yes
depends: foundation
state: closed/failed

## story 9: Report operational log read failures with exit code 1.
id: cli-read-fail
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Report operational log read failures with exit code 1.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReadErrors.md
covers: CLI-003
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4999
provides: exit 1 for an unreadable log
consumes: log reader
instructions: When the log file cannot be opened or read, exit 1 with nothing on stdout and one diagnostic line on stderr.
acceptance: yes
depends: cli-entry, log-read
state: closed/failed

## story 10: Validate date option values as real calendar dates.
id: filter-validate
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Validate date option values as real calendar dates.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FilterValidate.md
covers: FLT-003
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6342
provides: validated --since and --until values
consumes: usage-error path
instructions: Accept --since and --until only as real calendar dates written YYYY-MM-DD. Any other value is a usage error with exit 2, including an invalid date that appears after the log path.
acceptance: yes
depends: cli-usage
state: closed/failed

## story 11: Parse records and skip non-object, incomplete, and duplicate lines.
id: log-skip
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Parse records and skip non-object, incomplete, and duplicate lines.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogSkip.md
covers: LOG-002
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6767
provides: record parser with whole-file skipped count
consumes: line reader
instructions: Parse each line as a JSON object. Skip a line that is not an object, lacks a non-empty execution_id, started_at, status, job.command_name, or job.target, or repeats the execution_id of an earlier kept record. Keep the first occurrence. Write one stderr warning per skipped line naming its line number, and never write to stdout. Skipping never changes the exit code. Expose the total skipped count over the whole file.
acceptance: yes
depends: log-read
state: closed/failed

## story 12: Attribute each record to its executed or requested model.
id: log-model
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Attribute each record to its executed or requested model.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogModel.md
covers: LOG-003
accepts: st-007
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5642
provides: Record.Model
consumes: record parser
instructions: A record's model is result.stats.model when it is present and non-empty, otherwise job.model.
acceptance: yes
depends: log-skip
state: closed/failed

## story 13: Treat absent or null numeric stats as zero.
id: log-numbers
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Treat absent or null numeric stats as zero.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogNumbers.md
covers: LOG-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5753
provides: Record token, cost, and elapsed accessors defaulting to zero
consumes: record parser
instructions: A missing or null input_tokens, cached_input_tokens, output_tokens, cost_usd, or elapsed_ms counts as zero. A record with no result.stats object is still a record.
acceptance: yes
depends: log-skip
state: closed/failed

## story 14: Classify a record as failed when its status is not succeeded.
id: log-failed
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Classify a record as failed when its status is not succeeded.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogFailed.md
covers: LOG-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5888
provides: Record.Failed, nullable returncode, timed_out
consumes: record parser
instructions: A record is failed when status is any value other than succeeded. Expose its started_at, target, command, model, returncode (nullable), and timed_out for the failures report.
acceptance: yes
depends: log-skip
state: closed/failed

## story 15: Filter records by an inclusive start-date range.
id: filter-date
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Filter records by an inclusive start-date range.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FilterDate.md
covers: FLT-002
accepts: st-006
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6698
provides: date range filter predicate
consumes: record parser, parsed options
instructions: Keep records whose start date (the first 10 characters of started_at) is on or after --since and on or before --until when given. A range that matches nothing yields zero matching records and exit 0.
acceptance: yes
depends: log-skip, cli-parse
state: closed/failed

## story 16: Filter records by Target name.
id: filter-target
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Filter records by Target name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FilterTarget.md
covers: FLT-001
accepts: st-006
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6817
provides: target filter predicate
consumes: record parser, parsed options
instructions: Keep only records whose job.target equals the --target value exactly. Expose the filter as a predicate that every report applies.
acceptance: yes
depends: log-skip, cli-parse
state: closed/failed

## story 17: Produce the per-start-date report.
id: daily-report
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Produce the per-start-date report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-DailyReport.md
covers: DLY-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5882
provides: daily report
consumes: record aggregates, CSV and format helpers
instructions: Write one row per start date ascending. tokens is input plus cached plus output. Follow with a TOTAL row.
acceptance: yes
depends: filter-target, filter-date, log-numbers, log-failed, fmt-csv, fmt-minutes
state: closed/failed

## story 18: Produce the failed-record report.
id: failures-report
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Produce the failed-record report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FailuresReport.md
covers: FAL-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5885
provides: failures report
consumes: failed-record accessors, model attribution, CSV helpers
instructions: Write one row per failed record ordered by started_at then execution_id. returncode is empty when null or absent, and timed_out is true or false. This report has no TOTAL row.
acceptance: yes
depends: filter-target, filter-date, log-failed, log-model, fmt-csv
state: closed/failed

## story 19: Produce the per-Target report.
id: targets-report
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Produce the per-Target report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-TargetsReport.md
covers: TGT-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5928
provides: targets report
consumes: record aggregates, CSV and format helpers
instructions: Write one row per Target ordered ascending by byte value, with cost and minutes columns, followed by a TOTAL row.
acceptance: yes
depends: filter-target, filter-date, log-numbers, log-failed, fmt-csv, fmt-minutes, fmt-cost
state: closed/failed

## story 20: Count skipped lines over the whole file regardless of filters.
id: summary-skipped
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Count skipped lines over the whole file regardless of filters.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-SkippedCount.md
covers: SUM-002
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5984
provides: whole-file skipped count
consumes: record parser, target and date filters
instructions: The skipped count is computed over the whole log and ignores --target, --since, and --until. Test it by running the same log with and without filters.
acceptance: yes
depends: log-skip, filter-target, filter-date
state: closed/failed

## story 21: Produce the per-provider and per-model report.
id: models-report
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Produce the per-provider and per-model report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ModelsReport.md
covers: MDL-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6181
provides: models report
consumes: model attribution, record aggregates, CSV and format helpers
instructions: Write one row per provider and model pair ordered by provider then model ascending. The TOTAL row leaves the provider and model cells empty.
acceptance: yes
depends: log-model, log-numbers, log-failed, filter-target, filter-date, fmt-csv, fmt-cost
state: closed/failed

## story 22: Produce the one-row summary report.
id: summary-report
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Produce the one-row summary report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-SummaryReport.md
covers: SUM-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6379
provides: summary report
consumes: skipped count, record aggregates, CSV and format helpers
instructions: Write one row with the header records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes. first and last are the earliest and latest start dates, empty when no record matches.
acceptance: yes
depends: summary-skipped, log-failed, log-numbers, filter-target, filter-date, fmt-csv, fmt-minutes, fmt-cost
state: closed/failed

## story 23: Produce the per-command report.
id: commands-report
finding: step block-2 failed: log-read-blank-lines
evidence: evidence/block-2.md
summary: Produce the per-command report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CommandsReport.md
covers: CMD-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6477
provides: commands report
consumes: record aggregates, CSV and format helpers
instructions: Write one row per command ordered by calls descending, then command ascending by byte value, followed by a TOTAL row over all filtered records.
acceptance: yes
depends: filter-target, filter-date, log-numbers, log-failed, fmt-csv, fmt-minutes
state: closed/failed

## story 24: Assemble the end-to-end run from arguments to one written report.
id: report-pipeline
summary: Assemble the end-to-end run from arguments to one written report.
type: feature
kind: capability
phase: 1
block: 3
implements: FEATURE-ReportPipeline.md
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5619
provides: ddkstat <report> [options] <log> end-to-end
consumes: argument parser, reader, parser, filters, six reports, usage and read-failure paths
instructions: Assembly story. Parse arguments, read the log, skip invalid lines, apply filters, aggregate, and write the selected report to stdout. Intent: one run produces one report on stdout and diagnostics on stderr, with exit codes as specified.
acceptance: yes
depends: cli-read-fail, cli-usage, filter-validate, summary-report, commands-report, targets-report, models-report, daily-report, failures-report
state: pending

## story 25: Provide bin/test.sh over go vet and go test.
id: bld-test-script
summary: Provide bin/test.sh over go vet and go test.
type: service
kind: test harness
phase: 1
block: 4
implements: FEATURE-TestSuite.md
covers: BLD-002
accepts: st-002
context: ARCHITECTURE_compact.md
stack: go.md, common.md
stack_mode: builder
size: 8072
provides: bin/test.sh
consumes: all implementation packages
instructions: Write bin/test.sh as POSIX sh. It runs go vet ./... and go test ./... from the application root and exits zero only when both pass. The tests cover each report, filter, skip rule, and exit code using logs written by the tests.
acceptance: yes
depends: cli-entry, cli-parse, cli-usage, cli-read-fail, log-read, log-skip, log-model, log-numbers, log-failed, filter-target, filter-date, filter-validate, fmt-minutes, fmt-cost, fmt-csv, summary-skipped, summary-report, commands-report, targets-report, models-report, daily-report, failures-report, report-pipeline, foundation
state: pending

## story 26: Run the supplied grader as the release gate.
id:           bld-terminal
summary:      Run the supplied grader as the release gate.
type:         feature
kind:         test harness
phase:        1
block:        5
implements:   VERIFICATION.md
covers:       BLD-003
accepts:      st-001
instructions: |
  Terminal assembly story. Run sh sources/run.sh from the application root with Suite: full. It must exit 0, with all 20 checks passing against sources/llm_sample.jsonl.
stack:        go.md, common.md
stack_mode:   consumer
size:         5266
provides:     sh sources/run.sh passing
consumes:     supplied grader, every implementation story
depends:      cli-entry, cli-parse, cli-usage, cli-read-fail, log-read, log-skip, log-model, log-numbers, log-failed, filter-target, filter-date, filter-validate, fmt-minutes, fmt-cost, fmt-csv, summary-skipped, summary-report, commands-report, targets-report, models-report, daily-report, failures-report, report-pipeline, bld-test-script, foundation
acceptance:   yes
state:        pending
