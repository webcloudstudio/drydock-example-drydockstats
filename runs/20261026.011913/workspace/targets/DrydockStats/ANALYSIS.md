# Blueprint Analysis: DrydockStats

## Commander Expectations

- assert the reporter answers the operator's questions about Drydock's LLM execution log (`logs/llm.jsonl`): calls, providers, models, tokens, minutes, cost, and failures
- assert the reporter writes six CSV reports (summary, commands, targets, models, daily, failures) to standard output
- assert malformed and duplicate log lines are skipped without failing the run
- assert every report is computed only from the log it is given, and Target and date filters narrow every report
- assert the supplied grader `sources/run.sh` passes against the synthetic sample
- assert the application is built with the Go standard library only and has no database

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
| LOG-001 | Read the JSON Lines log | Each non-blank UTF-8 line is parsed as one JSON object; blank lines are ignored and not counted. |
| LOG-002 | Skip malformed lines | A line that is not a JSON object, or lacks a required non-empty field, is skipped, counted as skipped, and warned about on standard error with its line number. Exit code is unchanged. |
| LOG-003 | Skip duplicate execution ids | A line whose `execution_id` already appeared on an earlier kept record is skipped and counted; the first occurrence is kept. |
| LOG-004 | Treat absent numbers as zero | Absent or `null` token counts, cost, and durations count as zero; a record without `result.stats` is still a record. |
| LOG-005 | Attribute the executed model | A record's model is `result.stats.model` when present and non-empty, otherwise `job.model`. |

### Feature: Command Line

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse the command line | `<report>` comes first, options follow, and exactly one log path comes last; the six report names and the three options are accepted. The `ddkstat` main package at `./cmd/ddkstat` is the entry point. |
| CLI-002 | Validate date options | `--since` and `--until` accept only a real calendar date written `YYYY-MM-DD`. |
| CLI-003 | Exit 2 on usage errors | An unknown report or option, a missing option value, an invalid date, or a missing or extra log path exits 2 with nothing on standard output. |
| CLI-004 | Exit 1 on an unreadable log | A log that cannot be opened or read exits 1 with nothing on standard output. |

### Feature: Reports

| ID | Story | High-level AC |
|---|---|---|
| RPT-001 | Summary report | One row: records, skipped (whole file, regardless of filters), succeeded, failed, first and last start dates, token sums, cost, and minutes. |
| RPT-002 | Commands report | One row per command ordered by calls descending then command ascending, followed by a `TOTAL` row. |
| RPT-003 | Targets report | One row per Target ordered ascending, followed by a `TOTAL` row with cost and minutes. |
| RPT-004 | Models report | One row per provider and model pair ordered by provider then model ascending, followed by a `TOTAL` row. |
| RPT-005 | Daily report | One row per start date ascending; `tokens` is input plus cached plus output; followed by a `TOTAL` row. |
| RPT-006 | Failures report | One row per failed record ordered by `started_at` then `execution_id`; `returncode` empty when null or absent; `timed_out` is `true` or `false`. |

### Feature: Filters

| ID | Story | High-level AC |
|---|---|---|
| FLT-001 | Filter by Target | `--target` keeps only records whose `job.target` equals the name exactly, and applies to every report. |
| FLT-002 | Filter by date range | `--since` and `--until` keep records whose start date is on or after or on or before the date, and apply to every report. |

### Feature: Project Verification

| ID | Story | High-level AC |
|---|---|---|
| VER-001 | Provide the test entry point | `bin/test.sh` is POSIX-compatible, runs `go vet ./...` and `go test ./...` from the application root, and exits zero only when both pass. |
| VER-002 | Pass the supplied grader | `sh sources/run.sh` builds `./cmd/ddkstat` and passes all 20 checks, ending with `DrydockStats: 20 passed, 0 failed (of 20)`. |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | CLI-001 | If the arguments are malformed, then the application shall write a usage line to standard error that names each report and each option. |
| AC-002 | CLI-004 | If the log path names a directory, then the application shall exit 1 with nothing on standard output. |
| AC-003 | LOG-002 | If a line is not valid UTF-8, then the application shall treat it as not a JSON object and skip it. |
| AC-004 | LOG-002 | If `started_at` does not begin with a real calendar date in `YYYY-MM-DD` form, then the application shall skip the line as malformed, because every report keys on the start date. |
| AC-005 | LOG-003 | If a line is skipped for any reason other than a duplicate `execution_id`, then its `execution_id` shall not register, so a later valid record with that id is kept. |
| AC-006 | LOG-004 | If a token count, cost, or duration is present with a JSON type other than number or `null`, then the application shall count that field as zero and keep the record. |
| AC-007 | RPT-005 | The daily report shall list only start dates that have at least one matching record. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/llm_sample.jsonl` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/run.sh` | conformance-harness-to-fixture | `sources/llm_sample.jsonl` | `run.sh` reads the sample from its staged path and checks 20 known answers over it | Stage both; the sample is read-only and must stay byte-identical |
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/run.sh` | Check IDs A1–U6 map to the Output, Filters, and Exit codes sections | RPT, FLT, and CLI stories cite the checks they must satisfy; VER-002 gates on the harness |
| `sources/INSTRUCTIONS.md` | instruction-to-fixture | `sources/llm_sample.jsonl` | The Rules section defines skip, blank, and missing-value behavior the sample exercises (3 skipped lines, 1 blank line) | LOG-001 through LOG-004 are verified against the sample |
| `sources/INSTRUCTIONS.md` | dependency | `sources/run.sh` Tests section requirement | `bin/test.sh` must run `go vet` and `go test` | VER-001 is a delivery story owned by the build |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | normative specification | context | prompt-only |
| `sources/run.sh` | conformance harness | context | stage |
| `sources/llm_sample.jsonl` | asset | context | stage |

## Planning Instructions

### Delivery Shape

This is a command-line pipeline. Input is a JSON Lines UTF-8 log file given as the final argument. Output is one CSV report written to standard output, selected by the first argument. The application root is a Go 1.22+ module with the main package at `./cmd/ddkstat`, using only the standard library and no database.

Execution flow: parse arguments → open the log → ingest records, applying the skip rules → apply the Target and date filters → compute one report → write CSV. The skipped count is computed over the whole file before filters apply.

Shared output rules apply to RPT-001 through RPT-006 and are written as acceptance criteria in each report story: CSV with one header row and `\n` line endings; minutes written as tenths of a minute with `t = (ms + 3000) / 6000`; cost written with `%.4f`; text ordered by byte value; a `TOTAL` row of zeros when no record matches, where the report has one.

### Story Realization Map

- **LOG-001**: `sources/INSTRUCTIONS.md` "The log" and Rule 1; `sources/llm_sample.jsonl` blank line. Capability.
- **LOG-002**: `sources/INSTRUCTIONS.md` Rule 2; `sources/llm_sample.jsonl` truncated line, `20261014.999999.000Z-deadbeef` (no `job.target`). Capability.
- **LOG-003**: `sources/INSTRUCTIONS.md` Rule 2 (first occurrence kept); `sources/llm_sample.jsonl` repeated `20261015.083525.332Z-8c438a77`. Capability.
- **LOG-004**: `sources/INSTRUCTIONS.md` Rule 3; `sources/llm_sample.jsonl` `20261016.031951.895Z-dffd7d86` (no `result.stats`). Capability.
- **LOG-005**: `sources/INSTRUCTIONS.md` Rule 4 and field table; `sources/llm_sample.jsonl` records with and without `result.stats.model`. Capability.
- **CLI-001**: `sources/INSTRUCTIONS.md` Command section; `sources/run.sh` U2, U5, U6. Capability; creates the Go module and `cmd/ddkstat`.
- **CLI-002**: `sources/INSTRUCTIONS.md` Command date rule; `sources/run.sh` U4.
- **CLI-003**: `sources/INSTRUCTIONS.md` Exit codes; `sources/run.sh` U1–U6.
- **CLI-004**: `sources/INSTRUCTIONS.md` Exit codes; `sources/run.sh` E1.
- **RPT-001**: `sources/INSTRUCTIONS.md` Output summary row and Rule 6; `sources/run.sh` A1, B1, C2.
- **RPT-002**: Output commands row; `sources/run.sh` A2, B2, C1.
- **RPT-003**: Output targets row; `sources/run.sh` A3.
- **RPT-004**: Output models row; `sources/run.sh` A4, B5.
- **RPT-005**: Output daily row; `sources/run.sh` A5, B3.
- **RPT-006**: Output failures row; `sources/run.sh` A6, B4.
- **FLT-001**: `sources/INSTRUCTIONS.md` `--target`; `sources/run.sh` B4, B5, C1.
- **FLT-002**: `sources/INSTRUCTIONS.md` `--since` and `--until`; `sources/run.sh` B1, B2, B3, C2.
- **VER-001**: `sources/INSTRUCTIONS.md` Tests section. Capability; test entry point.
- **VER-002**: `sources/INSTRUCTIONS.md` Run Harness; `sources/run.sh` in full. Terminal acceptance.

### Test and Acceptance Strategy

Focused story tests run as `go test` in the package each story owns, using logs written by the tests, and are selected with `go test -run` so they never run the whole suite. Each implementation story declares `Suite: scoped`.

VER-002 is the single terminal verification story. It declares `Suite: full`, depends on every implementation story, and proves st-001 by running `sh sources/run.sh` unfiltered. It also proves st-002 by running `sh bin/test.sh` unfiltered. A sample proves a unit works, not that the project is correct; VER-002's pass is the project's acceptance evidence.

### Sequencing and Dependencies

- CLI-001 precedes every other story: it creates the Go module and the `cmd/ddkstat` main package.
- LOG-001 → LOG-002 → LOG-003 → LOG-004 → LOG-005 in that order, since each skip rule builds on the record stream.
- CLI-002 and CLI-003 follow CLI-001. CLI-004 follows LOG-001 because it depends on opening the log.
- FLT-001 and FLT-002 follow LOG-005 and CLI-002.
- RPT-001 follows LOG-005 and FLT-002; RPT-002 through RPT-006 follow RPT-001's shared output code. RPT-004 and RPT-006 require LOG-005 for model attribution.
- VER-001 follows all CLI and RPT stories, since it runs their tests.
- VER-002 follows every other story.

No external service calls, shared mutable state, or concurrency are described; no timeout or race-condition criteria apply.

### Source Conflicts and Gaps

No cross-source conflicts exist; the brief is the only source of behavior. Unattended UAT policy is in effect: no `BLOCKERS.md` and no questionnaires are emitted. Ambiguities are resolved as follows, and these resolutions are binding for Plan and Build:

1. **Help text.** The brief names no help output. The usage line on standard error (AC-001) is the help text. `--help` is an unknown option and exits 2.
2. **Duplicate tracking.** Only records that pass every other rule register their `execution_id`. A skipped line does not block a later valid record (AC-005).
3. **Invalid UTF-8.** Treated as not a JSON object and skipped (AC-003).
4. **Bad `started_at`.** A value that does not begin with a real `YYYY-MM-DD` date is skipped as malformed (AC-004).
5. **Wrong JSON type on a numeric field.** Counted as zero and the record is kept (AC-006).
6. **Daily rows.** Only dates with at least one matching record appear (AC-007).
7. **Directory as log path.** Exits 1 as unreadable (AC-002).

## Analysis Notes
generated: 2026-10-25
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.011913/workspace/targets/DrydockStats/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 5
  stories: 19
  stack: Go 1.22 or newer, standard library only; POSIX sh for `bin/test.sh`
  display_name: not proposed
  short_description: not proposed

- Project type: `cli`. The Core Baseline help-text check is covered by AC-001.
- Unattended UAT policy in effect: no blockers and no questionnaires emitted. Ambiguities are resolved under Planning Instructions → Source Conflicts and Gaps.
- `SEA_TRIALS.md` preserves the prior contract's IDs st-001 through st-008 and adds st-009 for the grader-independence guardrail.
- The sample is internally consistent with the grader: 37 lines, 33 records, 1 blank, 3 skipped. The skipped lines are the truncated line that follows `20261014.024820.526Z-b7631290`, the `deadbeef` record with no `job.target`, and the second `20261015.083525.332Z-8c438a77`.
- `20261016.031951.895Z-dffd7d86` has no `result.stats` and is kept with zero values.
- `COMPASS.md` exists and is not pending format, so it is not emitted. `discovery-identity.json` is omitted because `DISPLAY_NAME` and `SHORT_DESCRIPTION` are set.
- `STORY_GUIDANCE.json` is omitted. The sources state no build breakdown beyond the grader, and `sources/run.sh` accepts no scope selector.
- Story count is 19, under the 80-story sanity threshold.
