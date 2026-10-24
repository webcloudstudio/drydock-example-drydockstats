# MANIFEST: DrydockStats
updated:     2026-10-24T16:00:15+00:00
state:       approved
blocks:      3

applied_specs: |
  ARCHITECTURE.md sha256=fb4369df7298d66d79ee1fffaf8d0587a2db410ba5753b4d0a5558755fba11e9 commit=84da886ee8f51a1a70fa56769698abd073dfa15b applied_by=block-1 applied_at=2026-10-24T16:04:24+00:00 build_sha256=7cc7b4a743294f824abeb2c85e57cfe11c5bffc608bba25a35669b58f9fec327

## story 1: Create the Go module root, the ddkstat main package, and the internal CLI package.
id:           go-foundation
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-1.md
summary:      Create the Go module root, the ddkstat main package, and the internal CLI package.
type:         foundational
kind:         capability
phase:        1
block:        1
implements:   ARCHITECTURE.md
instructions: |
  Create go.mod at the application root with module path drydockstats, a go 1.22 directive, and no require lines.

  Create cmd/ddkstat/main.go. Its main calls os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr)) and is the only place that calls os.Exit.

  Create package internal/cli exposing Run(args []string, stdout, stderr io.Writer) int, returning 0 as a placeholder.

  Use the standard library only. Never read sources/run.sh from application code, never embed a report answer, and never write to the log being read.
stack:        common.md, go.md
stack_mode:   builder
size:         12476
provides:     go module drydockstats, cmd/ddkstat main package, internal/cli Run
acceptance:   yes
state: closed/implemented

## story 2: Prove the module builds the ddkstat executable with go build using only the standard library.
id: build-entry
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Prove the module builds the ddkstat executable with go build using only the standard library.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-BuildEntry.md
covers: BLD-001
accepts: st-003
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4647
provides: ddkstat executable
instructions: |
  Ensure go build -o <path> ./cmd/ddkstat succeeds from the application root with go.mod declaring no requires.
  
  Ensure every import in the module is from the standard library. Remove any non-standard import introduced by earlier stories.
  
  The build output path is supplied by the caller and never points into a Blueprint, evidence, or sources directory.
acceptance: yes
depends: go-foundation
state: closed/failed

## story 3: Decode the JSON Lines log, one JSON object per non-blank line.
id: log-decode
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Decode the JSON Lines log, one JSON object per non-blank line.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogDecode.md
covers: LOG-001
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4536
provides: log record decoder
instructions: |
  In package internal/logrecord, read the log as UTF-8 one line at a time with bufio.Scanner and a buffer large enough for long lines.
  
  Ignore blank lines without counting them. Decode each other line with encoding/json. A line that is not a JSON object is counted in skipped and writes one stderr warning naming its line number.
  
  Return the decoded records in file order with the skipped count. Failure to open or read the file is returned as an error to the caller.
acceptance: yes
depends: go-foundation
state: closed/failed

## story 4: Skip records missing a required field, with a line-numbered warning.
id: log-required-fields
finding: programmatic acceptance failed: log-required-fields-skip: IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'
evidence: evidence/block-2.md
summary: Skip records missing a required field, with a line-numbered warning.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogRequiredFields.md
covers: LOG-002
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4449
provides: required field validation
instructions: |
  Implement a pure validation function over one decoded object. The fields execution_id, started_at, status, job.command_name, and job.target must each be present and a non-empty string.
  
  A line that fails validation is counted in skipped and writes one stderr warning naming its line number. Skipping never changes the exit code.
acceptance: yes
depends: log-decode
state: closed/failed

## story 5: Skip repeated execution identifiers, keeping the first occurrence.
id: log-duplicates
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Skip repeated execution identifiers, keeping the first occurrence.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogDuplicates.md
covers: LOG-003
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4446
provides: duplicate execution identifier filter
instructions: |
  Keep a set of execution_id values seen so far, in file order. Run duplicate detection only on lines that passed required-field validation.
  
  A later line whose execution_id has already appeared is counted in skipped and writes one stderr warning naming its line number. The first occurrence is kept.
acceptance: yes
depends: log-required-fields
state: closed/failed

## story 6: Provide the shared CSV writer and the minutes and cost formatters.
id: report-format
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Provide the shared CSV writer and the minutes and cost formatters.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportFormat.md
covers: RPT-007
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4778
provides: FormatMinutes, FormatCost, WriteCSV
instructions: |
  FormatMinutes(ms int64) string computes t = (ms + 3000) / 6000 in integer division and writes t/10, a period, and t%10.
  
  FormatCost(usd float64) string writes the result of fmt.Sprintf with %.4f.
  
  WriteCSV(w io.Writer, rows [][]string) error uses encoding/csv with \n line endings and writes every row it is given.
acceptance: yes
depends: go-foundation
state: closed/failed

## story 7: Keep only records whose job.target equals the --target name exactly.
id: filter-target
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Keep only records whose job.target equals the --target name exactly.
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
size: 4787
provides: target filter
instructions: |
  Apply the --target predicate as an exact, case-sensitive string comparison with job.target. When the option is absent every record passes.
  
  The filter receives records after the skip rules have run, so the whole-file skipped count is unaffected.
acceptance: yes
depends: log-duplicates
state: closed/failed

## story 8: Treat absent or null token, cost, and duration values as zero.
id: log-numeric
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Treat absent or null token, cost, and duration values as zero.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogNumeric.md
covers: LOG-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4955
provides: record numeric stats
instructions: |
  Read result.stats.input_tokens, cached_input_tokens, output_tokens, cost_usd (float64), and elapsed_ms from each record. An absent or null value counts as zero.
  
  A record with no result.stats object is still a record, with every numeric value zero.
acceptance: yes
depends: log-decode
state: closed/failed

## story 9: Attribute each record to its executed model, falling back to the requested model.
id: log-model
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Attribute each record to its executed model, falling back to the requested model.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-LogModel.md
covers: LOG-004
accepts: st-007
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5256
provides: record model attribution
instructions: |
  Set a record's model to result.stats.model when that string is present and non-empty, otherwise to job.model.
  
  Read the provider from job.llm. An absent provider value is the empty string.
acceptance: yes
depends: log-decode
state: closed/failed

## story 10: Keep records whose start date falls within the inclusive --since and --until range.
id: filter-date
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Keep records whose start date falls within the inclusive --since and --until range.
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
size: 5298
provides: date range filter
instructions: |
  A record's start date is the first 10 characters of started_at. Keep the record when that text is on or after --since and on or before --until, for each option that is given.
  
  Compare as YYYY-MM-DD text, which orders the same as calendar dates. The CLI validates the option values before this filter runs.
acceptance: yes
depends: filter-target
state: closed/failed

## story 11: Produce the per-Target report with a TOTAL row.
id: report-targets
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Produce the per-Target report with a TOTAL row.
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
size: 5212
provides: targets report
instructions: |
  Group filtered records by job.target. Write the header target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes.
  
  Order rows by target ascending, byte-wise. End with a TOTAL row over all filtered records.
acceptance: yes
depends: filter-date, report-format, log-numeric
state: closed/failed

## story 12: Produce the per-date report with a TOTAL row.
id: report-daily
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Produce the per-date report with a TOTAL row.
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
size: 5300
provides: daily report
instructions: |
  Group filtered records by start date. Write the header date,calls,failed,tokens,minutes, where tokens is input plus cached plus output.
  
  Order rows by date ascending. End with a TOTAL row.
acceptance: yes
depends: filter-date, report-format, log-numeric
state: closed/failed

## story 13: Produce the per-command report with a TOTAL row.
id: report-commands
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Produce the per-command report with a TOTAL row.
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
size: 5303
provides: commands report
instructions: |
  Group filtered records by job.command_name. Write the header command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes.
  
  Order rows by calls descending, then command ascending by byte-wise comparison. Count failed as records whose status is not succeeded. End with a TOTAL row over all filtered records.
acceptance: yes
depends: filter-date, report-format, log-numeric
state: closed/failed

## story 14: Produce the per-provider and model report with a TOTAL row.
id: report-models
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Produce the per-provider and model report with a TOTAL row.
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
size: 5373
provides: models report
instructions: |
  Group filtered records by provider and model from log-model. Write the header provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd.
  
  Order rows by provider then model ascending, byte-wise. The TOTAL row has an empty model column.
acceptance: yes
depends: filter-date, report-format, log-model, log-numeric
state: closed/failed

## story 15: Produce the one-row summary report.
id: report-summary
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Produce the one-row summary report.
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
size: 5826
provides: summary report
instructions: |
  Write the header records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes followed by one data row.
  
  skipped counts the whole file regardless of filters. Records and token totals are over filtered records. first and last are the earliest and latest start dates, empty when no record matches.
acceptance: yes
depends: filter-date, report-format, log-numeric
state: closed/failed

## story 16: Produce the failed-record report.
id: report-failures
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Produce the failed-record report.
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
size: 5904
provides: failures report
instructions: |
  Emit one row per failed filtered record with the header started_at,target,command,model,returncode,timed_out. There is no TOTAL row.
  
  Order by started_at then execution_id ascending. returncode is empty when null or absent. timed_out is the literal true or false.
acceptance: yes
depends: filter-date, log-model
state: closed/failed

## story 17: Make every report write its header and a zero TOTAL row when no record matches.
id: report-empty-match
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Make every report write its header and a zero TOTAL row when no record matches.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReportEmptyMatch.md
covers: RPT-008
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5815
provides: empty-match report output
instructions: |
  Verify and, where needed, correct each report so that a filter matching no records still writes its header.
  
  Reports with a TOTAL row write a zero TOTAL row with 0.0 minutes and 0.0000 cost. The summary writes empty first and last. The failures report writes its header only.
  
  The skipped count still reflects the whole file.
acceptance: yes
depends: report-summary, report-commands, report-targets, report-models, report-daily, report-failures
state: closed/failed

## story 18: Parse the report name, options, and the single log path, and wire each report.
id: cli-parse
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Parse the report name, options, and the single log path, and wire each report.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliParse.md
covers: CLI-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5809
provides: ddkstat command line
instructions: |
  In internal/cli, parse Run(args, stdout, stderr) with a local flag.FlagSet configured with flag.ContinueOnError.
  
  The first argument is the report name, one of summary, commands, targets, models, daily, or failures. Options follow it, and exactly one log path is the final argument.
  
  Pass --target, --since, and --until to the filters, then call the matching report function.
acceptance: yes
depends: build-entry, filter-date, report-empty-match
state: closed/failed

## story 19: Exit 1 with no stdout when the log cannot be opened or read.
id: cli-unreadable-log
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Exit 1 with no stdout when the log cannot be opened or read.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliUnreadableLog.md
covers: CLI-003
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 4986
provides: unreadable log exit 1
instructions: |
  When the log cannot be opened or a read fails, return exit code 1, write exactly one stderr line, and write nothing to stdout.
  
  Build the full report in memory and write stdout only after the whole log has been read successfully.
acceptance: yes
depends: cli-parse, log-decode
state: closed/failed

## story 20: Map usage errors to exit code 2 with no stdout output.
id: cli-usage-errors
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Map usage errors to exit code 2 with no stdout output.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliUsageErrors.md
covers: CLI-002
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5115
provides: usage error exit 2
instructions: |
  Return exit code 2, write one stderr diagnostic, and write nothing to stdout for an unknown report, an unknown option, a missing option value, a missing or extra log path, or a date that is not a real calendar date.
  
  Validate --since and --until with time.Parse using the layout 2006-01-02. time.Parse rejects impossible dates such as 2026-02-30.
acceptance: yes
depends: cli-parse
state: closed/failed

## story 21: Exit 0 when lines are skipped or no record matches.
id: cli-success-partial
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Exit 0 when lines are skipped or no record matches.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CliSuccessPartial.md
covers: CLI-004
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 5800
provides: success exit 0 on partial data
instructions: |
  Confirm that a log with skipped lines or no matching records exits 0 and writes its report to stdout, with skip warnings on stderr only.
  
  Correct any path that returns a non-zero code for these cases.
acceptance: yes
depends: cli-unreadable-log
state: closed/failed

## story 22: Provide the POSIX bin/test.sh that runs go vet and go test over the whole module.
id: test-runner
finding: step block-2 failed: log-required-fields-skip
evidence: evidence/block-2.md
summary: Provide the POSIX bin/test.sh that runs go vet and go test over the whole module.
type: service
kind: test harness
phase: 1
block: 2
implements: FEATURE-TestRunner.md
covers: BLD-002
accepts: st-002
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 6737
provides: bin/test.sh
instructions: |
  Create bin/test.sh as a POSIX sh script with set -eu. It changes to the application root relative to its own location, runs go vet ./..., then go test ./..., and exits zero only when both pass.
  
  The script invokes the installed toolchain and never vendors one. Tests written by earlier stories write their logs under t.TempDir().
acceptance: yes
depends: build-entry, log-decode, log-required-fields, log-duplicates, log-model, log-numeric, filter-target, filter-date, report-format, report-summary, report-commands, report-targets, report-models, report-daily, report-failures, report-empty-match, cli-parse, cli-usage-errors, cli-unreadable-log, cli-success-partial
state: closed/failed

## story 23: Run the supplied grader and require exit zero as the release verdict.
id: grader-gate
summary: Run the supplied grader and require exit zero as the release verdict.
type: feature
kind: test harness
phase: 1
block: 3
implements: FEATURE-Acceptance.md
accepts: st-001
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 6312
instructions: |
  Assembly and intent. Once every member story is verified, run sh sources/run.sh from the application root. Its exit status is the sole verdict.
  
  Do not edit anything under sources/.
acceptance: yes
depends: test-runner, cli-usage-errors, cli-success-partial
state: pending
