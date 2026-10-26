# MANIFEST: DrydockStats
updated: 2026-10-26T01:35:56+00:00
state: approved
blocks:      7

applied_specs: |
  ARCHITECTURE.md sha256=a0aeaa9a1d2e57962bfcb2e82424f63ca75c6855cf36381965614c04c5e57aef commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=6e16fccf3c9884d7e113a56d3e64338da379d13abf702f2642af13e65ca33625
  FEATURE-AbsentNumbers.md sha256=8ee4d5892f98bc900d606a64553292f79d3735c84b590c20c399687300703b08 commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=6a8aa748ced0231abc8afe366917cbc6375eabc4ce10fd8d287b3a2fb55a47f1
  FEATURE-CommandLine.md sha256=bf96bc5be74f0e722ba79b7a59643e60e522b8a6ecedeef848515f39660c2f0b commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-2 applied_at=2026-10-26T01:41:07+00:00 build_sha256=3f5ecc329e51fec37b2e8d121d2ea3a67d8cd66f526d653183645fd89e3bae61
  FEATURE-DateOptions.md sha256=c98478ee142a949bcc778da9a304cb7c56e29b96f62bc204829d95c87e4b417d commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=00cda4f54f0b236a693a4db5c74d4dc2304bab2efd477f2a0d3d6d5954e4e931
  FEATURE-DuplicateIds.md sha256=3bbbce41bacf61c3a3aeb6d6490f0e97b22aa98c16399a5edaff69d4fb81b59e commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=4106dba36b8d0ccb8e2c86bf3d1159f1159b70b6aa043722e40f7d85abcef64b
  FEATURE-LogReading.md sha256=cc7c0ad7698305c3fa30d0679541bb0c5936f27ebd7390adafacce8ddc9f6b62 commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=01a957b4eecadf97472db3a241123c32b0945a3639db0fe9d92b1ce4ac0eb1d2
  FEATURE-MalformedLines.md sha256=2474666520a22238abb59e2b3ea99265c988cfb799fe95f9c5b88e09cd3b263b commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=5d0054e860bb888705f98e69df209bf62e09678949a19c6d34c8669d31fe81c5
  FEATURE-ModelAttribution.md sha256=e95eb544f396630dd259db597f526072dad60e081165c88555820a541304af6e commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=dde0f4c05efa067783d7dcf371996c2d952beefb0f6e917a0f0bbc09f36a3b88
  FEATURE-UnreadableLog.md sha256=11601896229409d22ed603bb577463bd1f0645ee32cbb8f8266fef8e736843fe commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=15ce1b6731bb3c4f8102a4d170d75598591f8f30bb2d915075d09f1d54b789d8
  FEATURE-UsageErrors.md sha256=f17687313a718ece26436d659b8f87cc4d9a84877c8426968341c587f1297430 commit=12c7525ced1dea949f498303c05e235a6068b5d3 applied_by=block-3 applied_at=2026-10-26T01:46:50+00:00 build_sha256=43d2b0585f8165f4a986fa22be9c501327babf617c66dc0ba5740d895086f928

## story 1: Stand up the Go module, the ddkstat main package, and the internal package layout.
id:           foundation
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-1.md
summary:      Stand up the Go module, the ddkstat main package, and the internal package layout.
type:         foundational
kind:         capability
phase:        1
block:        1
implements:   ARCHITECTURE.md
instructions: |
  Create the root go.mod (Go 1.22 or newer, standard library only, no third-party modules).
  Create cmd/ddkstat/main.go as the only caller of os.Exit, delegating to
  cli.Run(args []string, stdout, stderr io.Writer) int in internal/cli.
  Lay out internal/logread (parse, skip, dedupe), internal/report (aggregate and format as
  separate functions), and internal/cli. Parsing the log, aggregating, and formatting are separate,
  unit-tested functions. No database. Never read sources/run.sh from application code.
stack:        common.md, go.md
stack_mode:   builder
size:         14546
provides:     go.mod, cmd/ddkstat, internal/cli, internal/logread, internal/report
acceptance:   yes
state: closed/implemented

## story 2: Parse the command line into a report name, options, and one log path.
id: cli-parse
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-2.md
summary: Parse the command line into a report name, options, and one log path.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CommandLine.md
covers: CLI-001
accepts: st-003
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9944
provides: ddkstat <report> [--target] [--since] [--until] <log>
instructions: |
  Use a local flag-free or flag.FlagSet with ContinueOnError parser so the report comes first,
  options follow, and exactly one log path comes last. Accept summary, commands, targets, models,
  daily, failures and the options --target, --since, --until. Produce a typed options struct.
  go build ./cmd/ddkstat must succeed with only the standard library.
acceptance: yes
depends: foundation
state: closed/implemented

## story 3: Read the JSON Lines log line by line, ignoring blank lines.
id: log-read
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Read the JSON Lines log line by line, ignoring blank lines.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-LogReading.md
covers: LOG-001
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9490
provides: logread.Read
instructions: |
  Read the file as UTF-8 JSON Lines. Each non-blank line is decoded as one JSON object; blank lines
  are ignored and not counted. Track the 1-based line number of each line for warnings. Support
  lines longer than the default bufio token limit. Never write to the log.
acceptance: yes
depends: foundation
state: closed/implemented

## story 4: Validate --since and --until as real YYYY-MM-DD dates.
id: cli-dates
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Validate --since and --until as real YYYY-MM-DD dates.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-DateOptions.md
covers: CLI-002
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9023
provides: cli date validation
instructions: |
  Accept only a real calendar date in exactly YYYY-MM-DD form (reject 2026-13-01, 2026-02-30,
  and non-padded forms). Use time.Parse with the 2006-01-02 layout plus a length check.
acceptance: yes
depends: cli-parse
state: closed/implemented

## story 5: Exit 1 with nothing on stdout when the log cannot be opened or read.
id: cli-unreadable-log
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Exit 1 with nothing on stdout when the log cannot be opened or read.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-UnreadableLog.md
covers: CLI-004
accepts: st-008
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9057
provides: exit code 1
instructions: |
  A missing file, a directory, or any read error exits 1, writes one diagnostic line to standard
  error, and writes nothing to standard output. Buffer report output so a late read error cannot
  leave partial stdout.
acceptance: yes
depends: log-read, cli-parse
state: closed/implemented

## story 6: Exit 2 with nothing on stdout for every usage error.
id: cli-usage-errors
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Exit 2 with nothing on stdout for every usage error.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-UsageErrors.md
covers: CLI-003
accepts: st-008
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9531
provides: exit code 2
instructions: |
  An unknown report or option (including --help), a missing option value, an invalid date, or a
  missing or extra log path exits 2 with empty stdout. Write a usage line naming each report and
  each option to standard error.
acceptance: yes
depends: cli-dates
state: closed/implemented

## story 7: Skip malformed lines with a stderr warning and a skipped count.
id: log-skip-malformed
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Skip malformed lines with a stderr warning and a skipped count.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-MalformedLines.md
covers: LOG-002
accepts: st-004
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9798
provides: logread.Skipped
instructions: |
  Skip a line that is not a JSON object, is not valid UTF-8, lacks a non-empty string in any of
  execution_id, started_at, status, job.command_name, job.target, or whose started_at does not
  begin with a real YYYY-MM-DD calendar date. Count it as skipped and write one warning naming the
  line number to standard error. Skipping never changes the exit code.
acceptance: yes
depends: log-read
state: closed/implemented

## story 8: Skip lines repeating an earlier kept execution_id.
id: log-skip-duplicates
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Skip lines repeating an earlier kept execution_id.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-DuplicateIds.md
covers: LOG-003
accepts: st-004
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9371
provides: logread.Dedupe
instructions: |
  Keep the first occurrence of an execution_id; skip and count later repeats with a warning. Only
  records that pass every other rule register their id, so a line skipped as malformed does not
  block a later valid record with the same id.
acceptance: yes
depends: log-skip-malformed
state: closed/implemented

## story 9: Treat absent, null, or wrongly typed numeric fields as zero.
id: log-absent-numbers
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Treat absent, null, or wrongly typed numeric fields as zero.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-AbsentNumbers.md
covers: LOG-004
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9699
provides: logread.Record numeric fields
instructions: |
  Token counts, cost_usd, and elapsed_ms that are absent, null, or of a JSON type other than number
  count as zero. A record without result.stats is still a record. returncode null or absent is
  recorded as absent; timed_out defaults to false.
acceptance: yes
depends: log-skip-duplicates
state: closed/implemented

## story 10: Attribute each record to the executed model, falling back to the requested one.
id: log-model
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-3.md
summary: Attribute each record to the executed model, falling back to the requested one.
type: service
kind: capability
phase: 2
block: 3
implements: FEATURE-ModelAttribution.md
covers: LOG-005
accepts: st-007
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9095
provides: logread.Record.Model
instructions: |
  A record's model is result.stats.model when present and non-empty, otherwise job.model. The
  provider is job.llm.
acceptance: yes
depends: log-absent-numbers
state: closed/implemented

## story 11: Keep only records whose job.target equals --target exactly.
id: filter-target
finding: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report: AssertionError
evidence: evidence/block-4.md
summary: Keep only records whose job.target equals --target exactly.
type: service
kind: capability
phase: 3
block: 4
implements: FEATURE-TargetFilter.md
covers: FLT-001
accepts: st-006
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 11117
provides: report.FilterTarget
instructions: |
  Exact, case-sensitive match on job.target, applied before every report. The summary skipped
  count is not affected.
acceptance: yes
depends: log-model, cli-parse
state: closed/failed

## story 12: Keep records whose start date is within --since and --until inclusive.
id: filter-dates
finding: programmatic acceptance failed: date-since-until-inclusive, date-one-sided-bounds, date-time-of-day-ignored, date-every-report, date-with-target-and-empty-range: AssertionError
evidence: evidence/block-4.md
summary: Keep records whose start date is within --since and --until inclusive.
type: service
kind: capability
phase: 3
block: 4
implements: FEATURE-DateFilter.md
covers: FLT-002
accepts: st-006
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 12197
provides: report.FilterDates
instructions: |
  The start date is the first 10 characters of started_at. Compare as strings; --since is
  inclusive lower bound, --until inclusive upper bound. Applies to every report.
acceptance: yes
depends: filter-target, cli-dates
state: closed/failed

## story 13: Summary report plus the shared CSV, minutes, and cost formatters.
id: report-summary
summary: Summary report plus the shared CSV, minutes, and cost formatters.
type: service
kind: capability
phase: 4
block: 5
implements: FEATURE-SummaryReport.md
covers: RPT-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 12576
provides: ddkstat summary
instructions: |
  One row: records, skipped (whole file, regardless of filters), succeeded, failed, first, last
  (earliest and latest start dates, empty when none match), token sums, cost_usd, minutes.
  Implement shared formatters: minutes as t=(ms+3000)/6000 written t/10 "." t%10; cost with %.4f;
  CSV with one header and \n line endings.
acceptance: yes
depends: filter-dates
state: pending

## story 14: Failures report listing each failed record.
id: report-failures
summary: Failures report listing each failed record.
type: service
kind: capability
phase: 4
block: 5
implements: FEATURE-FailuresReport.md
covers: RPT-006
accepts: st-005
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 10184
provides: ddkstat failures
instructions: |
  One row per record whose status is not succeeded, ordered by started_at then execution_id
  ascending. returncode is empty when null or absent; timed_out is true or false. No TOTAL row.
acceptance: yes
depends: report-summary
state: pending

## story 15: Daily report with TOTAL row.
id: report-daily
summary: Daily report with TOTAL row.
type: service
kind: capability
phase: 4
block: 5
implements: FEATURE-DailyReport.md
covers: RPT-005
accepts: st-005
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 10410
provides: ddkstat daily
instructions: |
  One row per start date ascending, only dates with at least one matching record. tokens is input
  plus cached plus output. Then a TOTAL row.
acceptance: yes
depends: report-summary
state: pending

## story 16: Models report by provider and model with TOTAL row.
id: report-models
summary: Models report by provider and model with TOTAL row.
type: service
kind: capability
phase: 4
block: 5
implements: FEATURE-ModelsReport.md
covers: RPT-004
accepts: st-005
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 10908
provides: ddkstat models
instructions: |
  One row per provider and attributed model pair ordered by provider then model ascending, then a
  TOTAL row whose model cell is empty (TOTAL,,...).
acceptance: yes
depends: report-summary
state: pending

## story 17: Commands report with TOTAL row.
id: report-commands
summary: Commands report with TOTAL row.
type: service
kind: capability
phase: 4
block: 5
implements: FEATURE-CommandsReport.md
covers: RPT-002
accepts: st-005
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 11768
provides: ddkstat commands
instructions: |
  One row per command ordered by calls descending then command ascending (byte order), followed by
  a TOTAL row. With no matching record write the header and a zero TOTAL row (minutes 0.0).
acceptance: yes
depends: report-summary
state: pending

## story 18: Targets report with TOTAL row.
id: report-targets
summary: Targets report with TOTAL row.
type: service
kind: capability
phase: 4
block: 5
implements: FEATURE-TargetsReport.md
covers: RPT-003
accepts: st-005
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 12017
provides: ddkstat targets
instructions: |
  One row per Target ascending by byte value, then TOTAL with cost_usd and minutes (0.0000 and
  0.0 when empty).
acceptance: yes
depends: report-summary
state: pending

## story 19: Provide the POSIX bin/test.sh running go vet and go test.
id: test-entry-point
summary: Provide the POSIX bin/test.sh running go vet and go test.
type: service
kind: test harness
phase: 5
block: 6
implements: FEATURE-TestEntryPoint.md
covers: VER-001
accepts: st-002
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 8838
provides: bin/test.sh
instructions: |
  bin/test.sh is POSIX sh, changes to the application root, runs go vet ./... then go test ./...,
  and exits zero only when both pass. The Go suite covers each report, each filter, each skip
  rule, and each exit code using logs written by the tests.
acceptance: yes
depends: report-commands, report-targets, report-models, report-daily, report-failures, cli-usage-errors, cli-unreadable-log
state: pending

## story 20: Terminal assembly: the supplied grader passes all 20 checks.
id: grader-verification
summary: Terminal assembly: the supplied grader passes all 20 checks.
type: feature
kind: test harness
phase: 5
block: 7
implements: FEATURE-GraderVerification.md
covers: VER-002
accepts: st-001
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 9094
instructions: |
  Assembly and intent: with every member story closed, run sh sources/run.sh unfiltered and sh
  bin/test.sh unfiltered. Fix integration seams in the application only; never modify anything
  under sources/. The grader must end with the line reporting 20 passed, 0 failed.
acceptance: yes
depends: test-entry-point, report-summary, filter-dates, log-model, cli-unreadable-log
state: pending
