# Blueprint Analysis: DrydockStats

## Commander Expectations

- assert ddkstat reports Drydock's LLM execution log as six CSV reports: summary, commands, targets, models, daily, and failures
- assert malformed and duplicate log lines are skipped without failing the run
- assert every report honours the `--target`, `--since`, and `--until` filters
- assert the application uses only the Go standard library and no database
- assert the supplied grader `sources/run.sh` is the acceptance verdict

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Log Ingestion

| ID | Story | High-level AC |
|---|---|---|
| LOG-001 | Read the JSON Lines log | Each non-blank UTF-8 line is one JSON record. Blank lines are ignored and uncounted. A record is failed when `status` is not `succeeded`. |
| LOG-002 | Skip malformed and duplicate lines | A line that is not a JSON object, lacks a non-empty required field, or repeats an earlier `execution_id` is skipped with a stderr warning naming its line number. The exit code is unaffected. |
| LOG-003 | Treat missing numbers as zero | Absent or `null` token counts, cost, and duration count as zero. A record with no `result.stats` object is still a record. |
| LOG-004 | Attribute the executed model | A record's model is `result.stats.model` when present and non-empty, otherwise `job.model`. |

### Feature: Command Line

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse the report, options, and log path | Usage is `ddkstat <report> [options] <log>`. Exactly one log path is required. An unknown report or option is a usage error. |
| CLI-002 | Validate date options | `--since` and `--until` accept only real calendar dates in `YYYY-MM-DD`. An invalid date or a missing value is a usage error. |
| CLI-003 | Map failures to exit codes | An unreadable log exits 1. A usage error exits 2. A failing command writes nothing to stdout and one line to stderr. |

### Feature: Filters

| ID | Story | High-level AC |
|---|---|---|
| FLT-001 | Filter by target | `--target` keeps records whose `job.target` equals the name exactly. The filter applies to all six reports. |
| FLT-002 | Filter by start date range | `--since` and `--until` keep records whose start date is on or after, or on or before, the date. Both bounds are inclusive. The filter applies to all six reports. |

### Feature: Summary Report

| ID | Story | High-level AC |
|---|---|---|
| SUM-001 | Write the summary report | One row: records, whole-file skipped count, succeeded, failed, first and last start dates, token sums, cost as `%.4f`, and minutes. `first` and `last` are empty when no record matches. |

### Feature: Commands Report

| ID | Story | High-level AC |
|---|---|---|
| CMD-001 | Write the commands report | One row per command, ordered by calls descending then command ascending, followed by a `TOTAL` row. Minutes use tenths-of-a-minute formatting. |

### Feature: Targets Report

| ID | Story | High-level AC |
|---|---|---|
| TGT-001 | Write the targets report | One row per target, ordered by target ascending, followed by a `TOTAL` row with cost. |

### Feature: Models Report

| ID | Story | High-level AC |
|---|---|---|
| MOD-001 | Write the models report | One row per provider and model pair, ordered by provider then model ascending, followed by a `TOTAL` row with an empty model column. |

### Feature: Daily Report

| ID | Story | High-level AC |
|---|---|---|
| DAY-001 | Write the daily report | One row per start date ascending. `tokens` is input plus cached plus output. A `TOTAL` row follows. |

### Feature: Failures Report

| ID | Story | High-level AC |
|---|---|---|
| FAIL-001 | Write the failures report | One row per failed record ordered by `started_at` then `execution_id`. `returncode` is empty when null or absent. `timed_out` is `true` or `false`. |

### Feature: Build and Verification

| ID | Story | High-level AC |
|---|---|---|
| BLD-001 | Build ddkstat from the standard library | The Go module root builds `./cmd/ddkstat` with `go build`. The module declares no third-party requirements. |
| BLD-002 | Run the Go test suite through bin/test.sh | A POSIX `bin/test.sh` runs `go vet ./...` and `go test ./...` from the application root and exits zero only when both pass. |
| BLD-003 | Pass the supplied acceptance harness | `sh sources/run.sh` exits zero with all 20 checks passing against the staged sample. |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | LOG-002 | Each skipped-line warning on stderr uses the 1-based physical line number, counting blank lines. |
| AC-002 | LOG-002 | A malformed line does not register its `execution_id`, so a later line with that id is kept. |
| AC-003 | LOG-002 | Duplicate detection runs before filters. A later repeat is skipped even when its first occurrence is filtered out. |
| AC-004 | LOG-001 | `started_at` is accepted as any non-empty string. Its first 10 characters form the date, with no further timestamp validation. |
| AC-005 | LOG-003 | A non-numeric value in a numeric `result.stats` field counts as zero. |
| AC-006 | CLI-001 | A repeated option takes its last value. |
| AC-007 | CLI-001 | An option placed before the report name is a usage error. |
| AC-008 | CLI-003 | A usage error writes a one-line usage summary to stderr, so every command has its help text available. |
| AC-009 | FLT-002 | When `--since` is later than `--until`, the report is empty and the exit code is 0. |
| AC-010 | CMD-001 | Text values are written to CSV without quoting or escaping. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/llm_sample.jsonl` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| sources/INSTRUCTIONS.md | instruction-to-test | sources/run.sh | Rules and output tables map to checks A1–U6 | BLD-003 cites both; implementation follows the instruction |
| sources/run.sh | test-kit-to-implementation | cmd/ddkstat | `go build -o … ./cmd/ddkstat`, then 20 checks | BLD-001 precedes BLD-003 |
| sources/llm_sample.jsonl | dependency | sources/run.sh | `LOG="$DATA/llm_sample.jsonl"` | The fixture is staged before BLD-003 runs |
| sources/INSTRUCTIONS.md | dependency | sources/llm_sample.jsonl | Log field table and skip rules | LOG-001 through LOG-004 read this schema |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | normative specification | context | prompt-only |
| sources/run.sh | conformance harness | context | stage |
| sources/llm_sample.jsonl | asset | context | stage |

## Planning Instructions

### Delivery Shape

The system is a command-line pipeline with one input and one output. The input is the log file `logs/llm.jsonl`, plus the command arguments. The outputs are one CSV report on stdout, one warning per skipped line on stderr, and an exit code of 0, 1, or 2.

The execution flow is: parse arguments, then read the log line by line, then parse and validate each line, then drop duplicates, then apply filters, then aggregate, then format the report.

### Story Realization Map

- LOG-001: `sources/INSTRUCTIONS.md` Log and Rules 1, 5. Capability. Needs focused tests with test-written logs.
- LOG-002: Rule 2; `sources/llm_sample.jsonl` lines 16, 28, 34. Capability. Needs a test for each skip reason.
- LOG-003: Rule 3. Capability.
- LOG-004: Rule 4; sample line 35 has no stats. Capability.
- CLI-001: Command section. Capability.
- CLI-002: Command section, date rule. Capability.
- CLI-003: Exit codes section; `run.sh` E1, U1–U6. Capability.
- FLT-001: Filters table. Capability.
- FLT-002: Filters table. Capability.
- SUM-001: Output table, summary row; `run.sh` A1, B1, C2. Capability.
- CMD-001: Output table; Rules 7–9; `run.sh` A2, B2, C1. Capability.
- TGT-001: `run.sh` A3. Capability.
- MOD-001: `run.sh` A4, B5. Capability.
- DAY-001: `run.sh` A5, B3. Capability.
- FAIL-001: `run.sh` A6, B4. Capability.
- BLD-001: Stack section. Build capability.
- BLD-002: Tests section. Test harness.
- BLD-003: Run Harness section; `run.sh`. Acceptance contract.

### Test and Acceptance Strategy

Focused story tests cover each story's rule using logs written by the tests, and each story runs only its own tests. Story-level acceptance uses the `go test -run` selector for its owned tests and declares `Suite: scoped`.

BLD-003 is the single terminal verification story. It depends on every implementation story and declares `Suite: full`. It runs `sh sources/run.sh` unfiltered and gates on the exit status.

### Sequencing and Dependencies

- BLD-001 precedes every story that compiles.
- LOG-001 precedes LOG-002, which precedes LOG-003 and LOG-004.
- CLI-001 precedes CLI-002 and CLI-003.
- FLT-001 and FLT-002 precede all six report stories.
- The LOG stories precede all six report stories.
- The fixture `sources/llm_sample.jsonl` is staged before BLD-003.
- BLD-002 precedes BLD-003.

### Source Conflicts and Gaps

There are no cross-source conflicts. Every source definition of the log schema, skip rules, and output columns agrees with `run.sh`.

The sample file has 36 records and one blank line in 37 physical lines. Line 16 is truncated, line 28 repeats `execution_id` from line 27, and line 34 has no `job.target`. That gives 33 valid records and 3 skipped lines, which matches A1. The per-day counts in A5 also match the rows. Byte-for-byte totals are left to the BLD-003 run.

## Analysis Notes
generated: 2026-10-26
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/workspace/targets/DrydockStats/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 10
  stories: 18
  stack: Go 1.22 or newer
  display_name: not proposed
  short_description: not proposed

Observations and choices made under the unattended policy:

- Project type is `cli`. Core baseline is covered by AC-008.
- Security: there is no authentication, no secret, and no external service. The log carries prompt sizes, not prompt bodies. Nothing routes to a questionnaire.
- Persistence: none. The product reads a file and writes stdout.
- Concurrency: none. The process is single-threaded and has no shared state.
- Story count is 18, below the 80 threshold, so no story-count question is emitted.
- No numeric threshold or measurement is stated, so no measurement criteria or `## Questions` entries are needed.
- Ambiguities resolved to the narrowest reading: repeated options take the last value (AC-006); options before the report are a usage error (AC-007); `--since` after `--until` is an empty result, not an error (AC-009); `started_at` is not validated beyond being a non-empty string (AC-004); non-numeric stats count as zero (AC-005); CSV values are written raw (AC-010).
