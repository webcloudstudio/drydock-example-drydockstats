# Blueprint Analysis: DrydockStats

## Commander Expectations

- assert the product builds a command-line reporter named `ddkstat` from Drydock's LLM execution log
- assert the reporter reads `logs/llm.jsonl` as JSON Lines and computes every report from the log it is given
- assert malformed and duplicate records are skipped without changing the exit code
- assert the supplied grader `sources/run.sh` passes against `sources/llm_sample.jsonl`
- assert the product uses only the Go standard library and no database

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
| LOG-001 | Read JSON Lines log records | Each non-blank line is parsed as one JSON object; the log path is the single required positional argument. |
| LOG-002 | Skip malformed and duplicate lines | Lines that are not objects, lack a required non-empty field, or repeat an earlier `execution_id` are skipped, counted, and warned on stderr without changing the exit code. |
| LOG-003 | Attribute the executed model | A record's model is `result.stats.model` when present and non-empty, otherwise `job.model`. |
| LOG-004 | Treat missing numbers as zero | Absent or null token counts, cost, and duration count as zero; a record without `result.stats` is still a record. |

### Feature: Command Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse the report name and log path | `<report>` comes first, options follow, and exactly one log path is last; `summary`, `commands`, `targets`, `models`, `daily`, `failures` are accepted. |
| CLI-002 | Apply `--target` filter | Only records whose `job.target` equals the value exactly are kept. |
| CLI-003 | Apply `--since` and `--until` filters | Records are kept by start date (first 10 characters of `started_at`) inclusively; dates must be real `YYYY-MM-DD` calendar dates. |
| CLI-004 | Map usage errors to exit code 2 | Unknown report or option, missing option value, invalid date, or missing or extra log path exits 2 with no stdout. |
| CLI-005 | Map unreadable log to exit code 1 | A log that cannot be opened or read exits 1 with no stdout. |

### Feature: Reports

| ID | Story | High-level AC |
|---|---|---|
| RPT-001 | Summary report | One row with `records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes`; `skipped` counts the whole file regardless of filters. |
| RPT-002 | Commands report | One row per command ordered by calls descending then command ascending, plus a `TOTAL` row. |
| RPT-003 | Targets report | One row per Target ordered ascending, plus a `TOTAL` row. |
| RPT-004 | Models report | One row per provider and model pair ordered ascending, plus a `TOTAL` row. |
| RPT-005 | Daily report | One row per start date ascending with `tokens` as input plus cached plus output, plus a `TOTAL` row. |
| RPT-006 | Failures report | One row per failed record ordered by `started_at` then `execution_id`; no `TOTAL` row. |
| RPT-007 | Empty-match output | A report with no matching records writes its header and, where it has one, a zero `TOTAL` row (`0.0` minutes, `0.0000` cost). |
| RPT-008 | Minutes and cost formatting | Minutes are `t = (ms + 3000) / 6000` written as `t/10` `.` `t%10`; cost is `%.4f`. |

### Feature: Build and Verification

| ID | Story | High-level AC |
|---|---|---|
| BLD-001 | Go module and executable | Application root is a Go 1.22+ module; `go build ./cmd/ddkstat` produces the executable with only the standard library. |
| BLD-002 | Test suite script | `bin/test.sh` is POSIX-compatible and exits zero only when `go vet ./...` and `go test ./...` both pass. |
| BLD-003 | Grader passes | `sh sources/run.sh` exits zero against the supplied sample log. |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | LOG-001 | A blank line is ignored and not counted as a record or a skip. |
| AC-002 | LOG-002 | Each skipped line writes one stderr warning naming its line number, and stdout is unaffected by the warning. |
| AC-003 | LOG-002 | The first occurrence of an `execution_id` is kept; later duplicates are skipped. |
| AC-004 | LOG-004 | A record whose `result.stats` object is absent is counted as a record with zero tokens, cost, and duration. |
| AC-005 | CLI-004 | A failing command writes nothing to stdout and exactly one diagnostic line to stderr. |
| AC-006 | RPT-001 | `first` and `last` are the earliest and latest start dates, empty when no record matches. |
| AC-007 | RPT-006 | `returncode` is empty when `null` or absent; `timed_out` is written as `true` or `false`. |
| AC-008 | RPT-004 | The models report's TOTAL row begins `TOTAL,,` (empty provider column). |
| AC-009 | BLD-001 | `sources/run.sh` is staged verbatim into the application root and never read by application code. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/llm_sample.jsonl` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/run.sh` | test-kit-to-implementation | `cmd/ddkstat` executable | `sources/run.sh` builds `./cmd/ddkstat` and runs 20 checks | Executable must exist before grader; BLD-003 depends on all report stories. |
| `sources/llm_sample.jsonl` | fixture-to-implementation | Log ingestion and every report | `sources/llm_sample.jsonl` (37 lines, 33 valid) | Fixture drives all known-answer checks. |
| `sources/INSTRUCTIONS.md` | dependency | All stories | `sources/INSTRUCTIONS.md` Rules 1–9, Output, Exit codes | Behavioral spec for every story. |
| LOG-002 | parser-to-normalizer | RPT-001 through RPT-006 | Skip rule governs `skipped` count and filter-independent behavior | Ingestion precedes all reports. |
| CLI-004, CLI-005 | dependency | `ddkstat` argument handling | `sources/INSTRUCTIONS.md` Exit codes table | Exit code mapping shared by all commands. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | author intent | compass | none |
| `sources/run.sh` | conformance harness | context | stage |
| `sources/llm_sample.jsonl` | conformance fixture | context | stage |

## Planning Instructions

### Delivery Shape

The system is a single-binary Go CLI. Inputs: one JSON Lines log file (`logs/llm.jsonl` or any path given). Outputs: CSV on stdout, diagnostics on stderr. Required flow: parse arguments → read and skip records → apply filters → aggregate → emit the selected report. Exit 0 on success including skipped lines, 1 on unreadable log, 2 on usage error.

### Story Realization Map

- LOG-001–LOG-004: `internal` log package; cited `sources/INSTRUCTIONS.md` Log and Rules 1–5; fixture `sources/llm_sample.jsonl`; capability stories.
- CLI-001–CLI-005: `cmd/ddkstat` main package; cited `sources/INSTRUCTIONS.md` Command and Exit codes; capability stories.
- RPT-001–RPT-008: report package; cited `sources/INSTRUCTIONS.md` Output and Rules 6–9; known answers in `sources/run.sh`.
- BLD-001: Go module root; cited `sources/INSTRUCTIONS.md` Stack; build gate.
- BLD-002: test harness script `bin/test.sh`; cited `sources/INSTRUCTIONS.md` Tests; verification gate.
- BLD-003: terminal verification; depends on every implementation story; runs `sh sources/run.sh`.

### Test and Acceptance Strategy

Focused story tests cover each report, filter, skip rule, and exit code using logs written by tests. The terminal story BLD-003 gates on the complete grader run (`Suite: full`). No story runs the grader except BLD-003.

### Sequencing and Dependencies

Build (BLD-001) precedes all tests. Log ingestion (LOG-001–LOG-004) precedes CLI filters and reports. CLI-001 precedes CLI-002–CLI-005. All RPT stories precede BLD-002 and BLD-003. BLD-003 is terminal and depends on every implementation story.

### Source Conflicts and Gaps

None. The sources define one consistent behavior for skip rules, model attribution, filters, and output formats. The `models` provider column for the `sonnet` entry in `sources/run.sh` A4 is an observed fixture value and is not a conflict.

## Analysis Notes
generated: 2026-10-24
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.143257/workspace/targets/DrydockStats/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 4
  stories: 16
  stack: Go 1.22+, standard library only
  display_name: not proposed
  short_description: not proposed

None.
