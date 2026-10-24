# MANIFEST: DrydockStats
updated: 2026-10-24T14:44:43+00:00
state: approved
blocks:      5

## story 1: Stand up the Go module, the bin/test.sh launcher, and a placeholder package.
id:           go-foundation
summary:      Stand up the Go module, the bin/test.sh launcher, and a placeholder package.
type:         foundational
kind:         capability
phase:        1
block:        1
implements:   ARCHITECTURE.md
instructions: |
  Create the root go.mod with module path drydock.local/ddkstat, a go 1.22 directive, and no require directives.

  Create internal/doc/doc.go as a placeholder package so go vet ./... and go test ./... always have a package to match.

  Create bin/test.sh as a POSIX sh launcher that runs go vet ./... and then go test ./... from the application root, and exits non-zero if either fails. Do not vendor a toolchain.

  Record in ARCHITECTURE.md the module ownership: internal/log for ingestion, internal/report for aggregation and formatting, internal/cli for argument handling and dispatch, cmd/ddkstat for process startup only.
stack:        go.md, common.md
stack_mode:   builder
size:         13188
provides:     go.mod module drydock.local/ddkstat, bin/test.sh launcher
acceptance:   yes
state:        pending

## story 2: Produce the ddkstat executable from ./cmd/ddkstat using only the standard library.
id: build-executable
summary: Produce the ddkstat executable from ./cmd/ddkstat using only the standard library.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Executable.md
covers: BLD-001
accepts: st-003
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6166
provides: cmd/ddkstat main package, internal/cli Run entry point
instructions: |
  Create cmd/ddkstat/main.go as package main. It calls cli.Run(os.Args[1:], os.Stdout, os.Stderr) and passes its int result to os.Exit. It is the only place os.Exit is called.
  
  Create internal/cli/run.go exporting Run(args []string, stdout, stderr io.Writer) int. For now it writes nothing to stdout and returns 2 for every call; cli-parse replaces the body.
  
  Use only the standard library. Confirm go.mod still has no require directives.
acceptance: yes
depends: go-foundation
state: pending

## story 3: Provide the minutes, cost, and CSV writing helpers used by every report.
id: report-format
summary: Provide the minutes, cost, and CSV writing helpers used by every report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportFormat.md
covers: RPT-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5690
provides: internal/report FormatMinutes, FormatCost, CSV writer
instructions: |
  Minutes use t = (ms + 3000) / 6000 in integer division, written as t/10, a dot, and t%10. For example, 95000 ms renders as 1.6.
  
  Cost uses %.4f for a float64. Write CSV with \n line endings and one header row.
  
  Keep formatting separate from aggregation so both are unit-tested on their own.
acceptance: yes
depends: go-foundation
state: pending

## story 4: Parse the report name, options, and the single log path in the required order.
id: cli-parse
summary: Parse the report name, options, and the single log path in the required order.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliParse.md
covers: CLI-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5760
provides: internal/cli Run dispatch, report name table, option grammar
instructions: |
  Accept report as the first argument, followed by options, with exactly one log path last. The six report names are summary, commands, targets, models, daily, and failures.
  
  Use a local flag.FlagSet with flag.ContinueOnError and bytes.Buffer in tests. Declare --target, --since, and --until so they parse. Their semantics are implemented by cli-target-filter and cli-date-filters.
  
  Own the report dispatch table. Each report story registers its name here. Sequencing the report stories after report-summary avoids conflicting edits to this table.
acceptance: yes
depends: build-executable
state: pending

## story 5: Read the log path as JSON Lines, yielding each non-blank line with its line number.
id: log-reader
summary: Read the log path as JSON Lines, yielding each non-blank line with its line number.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogReader.md
covers: LOG-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5761
provides: internal/log ReadLines, internal/log Line
instructions: |
  In internal/log, read the file as UTF-8 and yield each non-blank line with its 1-based line number. Blank lines are ignored and are not counted anywhere.
  
  Do not parse JSON or classify lines here. Classification belongs to log-skip. Return an error when the file cannot be opened or read so the caller can map it to exit code 1.
  
  Accept an io.Reader in the core function so tests can use bytes.Buffer; the file open is a thin wrapper.
acceptance: yes
depends: go-foundation
state: pending

## story 6: Map an unreadable log to exit code 1 with no stdout and one stderr line.
id: cli-read-errors
summary: Map an unreadable log to exit code 1 with no stdout and one stderr line.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliReadErrors.md
covers: CLI-005
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5399
provides: Exit code 1 for operational failures
instructions: |
  When the log cannot be opened or a read fails, exit 1, write nothing to stdout, and write one diagnostic line to stderr.
  
  Treat a missing file as this operational failure, not as a usage error. The path has already passed argument validation.
acceptance: yes
depends: cli-parse, log-reader
state: pending

## story 7: Attribute each record to the executed model, falling back to the requested model.
id: log-model
summary: Attribute each record to the executed model, falling back to the requested model.
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
size: 5401
provides: Record.Model derivation in internal/log
instructions: |
  A record's Model is result.stats.model when that field is present and non-empty, otherwise job.model.
  
  Implement the derivation inside the record normalizer that log-skip produces, so every report reads Model from the same field.
acceptance: yes
depends: log-reader
state: pending

## story 8: Map every usage error to exit code 2 with no stdout and one stderr line.
id: cli-usage-errors
summary: Map every usage error to exit code 2 with no stdout and one stderr line.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliUsageErrors.md
covers: CLI-004
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5775
provides: Exit code 2 for usage errors
instructions: |
  Usage errors are an unknown report, an unknown option, a missing option value, an invalid date that is not a real YYYY-MM-DD calendar date, and a missing or extra log path.
  
  Each exits 2, writes nothing to stdout, and writes exactly one diagnostic line to stderr.
  
  Validate dates with time.Parse using the layout 2006-01-02 and reject any string that does not round-trip exactly.
acceptance: yes
depends: cli-parse
state: pending

## story 9: Zero-fill absent numeric fields and derive the failed flag on each normalized record.
id: log-numbers
summary: Zero-fill absent numeric fields and derive the failed flag on each normalized record.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogNumbers.md
covers: LOG-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5788
provides: Record numeric fields, Record.Failed, Record.ReturnCode, Record.TimedOut
instructions: |
  An absent or null input_tokens, cached_input_tokens, output_tokens, cost_usd, or elapsed_ms counts as zero. A record with no result.stats object is still a record, with zero tokens, cost, and duration.
  
  Set Failed when status is not succeeded. Keep returncode (nil when null or absent) and timed_out for the failures report.
  
  Keep cost as float64 and durations as integer milliseconds so the formatter in report-format can apply the Rule 7 and Rule 8 rounding.
acceptance: yes
depends: log-reader
state: pending

## story 10: Classify each non-blank line and skip malformed or duplicate records with a warning.
id: log-skip
summary: Classify each non-blank line and skip malformed or duplicate records with a warning.
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
size: 5861
provides: internal/log Normalize, internal/log Record, skipped count
instructions: |
  A line is skipped when it is not a JSON object, when any of execution_id, started_at, status, job.command_name, or job.target is absent or not a non-empty string, or when its execution_id already appeared on an earlier line.
  
  The first occurrence of an execution_id is kept. Each skipped line writes one warning to stderr naming its line number, and stdout is never touched.
  
  Return the kept records and the total skipped count for the whole file. Skipping never changes the exit code. Filters are applied after this stage, so the skipped count is unaffected by filters.
acceptance: yes
depends: log-reader
state: pending

## story 11: Keep only records whose job.target equals the --target value exactly.
id: cli-target-filter
summary: Keep only records whose job.target equals the --target value exactly.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliTargetFilter.md
covers: CLI-002
accepts: st-006
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6104
provides: Target filter applied before every report
instructions: |
  Apply an exact, case-sensitive equality test on job.target after skipping and before aggregation.
  
  Filtering applies to every report. The skipped count in summary is never filtered.
acceptance: yes
depends: cli-parse, log-skip
state: pending

## story 12: Keep records whose start date falls inclusively within --since and --until.
id: cli-date-filters
summary: Keep records whose start date falls inclusively within --since and --until.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliDateFilters.md
covers: CLI-003
accepts: st-006
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6221
provides: Date filter over the first 10 characters of started_at
instructions: |
  The start date is the first 10 characters of started_at. Both bounds are inclusive and compare as strings, which is correct for YYYY-MM-DD.
  
  Assume the date values are already validated. Invalid-date rejection (exit 2) is owned by cli-usage-errors.
acceptance: yes
depends: cli-parse, log-skip
state: pending

## story 13: Produce the one-row summary report and register the summary name.
id: report-summary
summary: Produce the one-row summary report and register the summary name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportSummary.md
covers: RPT-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6149
provides: summary report, dispatch entry for summary
instructions: |
  Emit one row with the header records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes.
  
  skipped is the whole-file count and ignores filters. first and last are the earliest and latest start dates, empty when no record matches.
  
  Register the summary name in the cli-parse dispatch table.
acceptance: yes
depends: report-format, log-skip, log-model, log-numbers, cli-target-filter, cli-date-filters, cli-read-errors, cli-parse
state: pending

## story 14: Produce the per-start-date report with a TOTAL row, and register the daily name.
id: report-daily
summary: Produce the per-start-date report with a TOTAL row, and register the daily name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportDaily.md
covers: RPT-005
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5878
provides: daily report, dispatch entry for daily
instructions: |
  Header: date,calls,failed,tokens,minutes. Group by the first 10 characters of started_at, ascending. tokens is input plus cached plus output.
  
  Register the daily name in the dispatch table.
acceptance: yes
depends: report-summary, report-format, log-model, log-numbers, cli-target-filter, cli-date-filters, cli-read-errors
state: pending

## story 15: Produce the per-Target report with a TOTAL row, and register the targets name.
id: report-targets
summary: Produce the per-Target report with a TOTAL row, and register the targets name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportTargets.md
covers: RPT-003
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5909
provides: targets report, dispatch entry for targets
instructions: |
  Header: target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes. One row per Target ascending by byte value, then a TOTAL row.
  
  Register the targets name in the dispatch table.
acceptance: yes
depends: report-summary, report-format, log-model, log-numbers, cli-target-filter, cli-date-filters, cli-read-errors
state: pending

## story 16: Produce the per-command report with a TOTAL row, and register the commands name.
id: report-commands
summary: Produce the per-command report with a TOTAL row, and register the commands name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportCommands.md
covers: RPT-002
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5933
provides: commands report, dispatch entry for commands
instructions: |
  Header: command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes. Group by job.command_name, ordered by calls descending then command ascending by byte value. Append a TOTAL row.
  
  Register the commands name in the dispatch table.
acceptance: yes
depends: report-summary, report-format, log-model, log-numbers, cli-target-filter, cli-date-filters, cli-read-errors
state: pending

## story 17: Produce the per-provider and model report with an empty-provider TOTAL row, and register the models name.
id: report-models
summary: Produce the per-provider and model report with an empty-provider TOTAL row, and register the models name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportModels.md
covers: RPT-004
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5939
provides: models report, dispatch entry for models
instructions: |
  Header: provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd. Group by provider and model, ordered by provider then model ascending.
  
  The TOTAL row begins TOTAL,, with an empty provider column.
  
  Register the models name in the dispatch table.
acceptance: yes
depends: report-summary, report-format, log-model, log-numbers, cli-target-filter, cli-date-filters, cli-read-errors
state: pending

## story 18: Produce the failed-record report with no TOTAL row, and register the failures name.
id: report-failures
summary: Produce the failed-record report with no TOTAL row, and register the failures name.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportFailures.md
covers: RPT-006
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 7716
provides: failures report, dispatch entry for failures
instructions: |
  Header: started_at,target,command,model,returncode,timed_out. One row per failed record, ordered by started_at then execution_id ascending. There is no TOTAL row.
  
  returncode is empty when null or absent. timed_out is written as true or false.
  
  Register the failures name in the dispatch table.
acceptance: yes
depends: report-summary, report-format, log-model, log-numbers, cli-target-filter, cli-date-filters, cli-read-errors
state: pending

## story 19: Assemble and verify that every report writes its header and zero TOTAL row when nothing matches.
id: report-empty-match
summary: Assemble and verify that every report writes its header and zero TOTAL row when nothing matches.
type: feature
kind: capability
phase: 1
block: 3
implements: FEATURE-ReportEmptyMatch.md
covers: RPT-007
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 9545
provides: empty-match behavior across all six reports
instructions: |
  Assemble the empty-match behavior across the reports. Each report run with a filter that matches nothing writes its header.
  
  Reports that have a TOTAL row write zeros: 0.0 minutes and 0.0000 cost. The summary row has empty first and last. failures writes only its header.
  
  Assert these through parsed CSV field values, not substrings.
acceptance: yes
depends: report-summary, report-commands, report-targets, report-models, report-daily, report-failures
state: pending

## story 20: Finalize bin/test.sh and add end-to-end coverage across reports, filters, skip rules, and exit codes.
id: test-suite
summary: Finalize bin/test.sh and add end-to-end coverage across reports, filters, skip rules, and exit codes.
type: service
kind: test harness
phase: 1
block: 4
implements: FEATURE-TestSuite.md
covers: BLD-002
accepts: st-002
context: ARCHITECTURE_compact.md
stack: go.md, common.md
stack_mode: consumer
size: 7868
provides: bin/test.sh verification gate, end-to-end test in cmd/ddkstat
instructions: |
  Finalize bin/test.sh so it is POSIX-compatible, runs go vet ./... and go test ./... from the application root, and exits zero only when both pass.
  
  Add an end-to-end test under cmd/ddkstat that writes logs in test code and exercises each report, each filter, each skip rule, and each exit code through cli.Run.
  
  Do not modify any file under sources/.
acceptance: yes
depends: go-foundation, build-executable, log-reader, log-skip, log-model, log-numbers, cli-parse, cli-target-filter, cli-date-filters, cli-usage-errors, cli-read-errors, report-format, report-summary, report-commands, report-targets, report-models, report-daily, report-failures, report-empty-match
state: pending

## story 21: Run the supplied grader sh sources/run.sh against the completed application as the terminal gate.
id: grader-run
summary: Run the supplied grader sh sources/run.sh against the completed application as the terminal gate.
type: feature
kind: test harness
phase: 1
block: 5
implements: FEATURE-Grader.md
covers: BLD-003
accepts: st-001
context: ARCHITECTURE_compact.md
stack: common.md
stack_mode: consumer
size: 7229
provides: terminal verification gate
instructions: |
  This is the terminal story and runs last. Every other story is a transitive dependency.
  
  Its acceptance runs sh sources/run.sh from the application root with Suite: full, and asserts the exit status is zero. Do not re-type the 20-check total or any known answer; the harness exit status is the verdict.
  
  Do not modify sources/run.sh or sources/llm_sample.jsonl. If checks fail, report the captured output and fix the application, not the grader.
acceptance: yes
depends: go-foundation, build-executable, log-reader, log-skip, log-model, log-numbers, cli-parse, cli-target-filter, cli-date-filters, cli-usage-errors, cli-read-errors, report-format, report-summary, report-commands, report-targets, report-models, report-daily, report-failures, report-empty-match, test-suite
state: pending
