# Blueprint Analysis: DrydockStats

## Commander Expectations

- assert the product reads Drydock's LLM execution log (`logs/llm.jsonl`) and writes six CSV reports to standard output
- assert the product runs from a single command with a report name, optional filters, and a log path
- assert malformed, blank, and duplicate log lines never make the product fail
- assert every report is computed from the log it is given, never from embedded answers
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
| LOG-001 | Read the JSON Lines log and decode each record | Each non-blank line is decoded as one JSON object; undecodable lines are reported as skipped |
| LOG-002 | Skip records missing required fields | A line lacking any of `execution_id`, `started_at`, `status`, `job.command_name`, `job.target` is skipped with a stderr warning naming its line number |
| LOG-003 | Skip duplicate execution identifiers | A repeated `execution_id` is skipped; the first occurrence is kept |
| LOG-004 | Attribute the executed model | Model is `result.stats.model` when present and non-empty, otherwise `job.model` |
| LOG-005 | Treat missing numeric stats as zero | Absent or null token, cost, and duration values count as zero; a record without `result.stats` is still a record |

### Feature: Filtering

| ID | Story | High-level AC |
|---|---|---|
| FLT-001 | Filter records by target | `--target` keeps only records whose `job.target` equals the name exactly |
| FLT-002 | Filter records by start date | `--since` and `--until` keep records whose start date falls within the inclusive range; dates must be real `YYYY-MM-DD` values |

### Feature: Reports

| ID | Story | High-level AC |
|---|---|---|
| RPT-001 | Produce the summary report | One row with records, skipped, succeeded, failed, first, last, tokens, cost, minutes; skipped counts the whole file regardless of filters |
| RPT-002 | Produce the commands report | One row per command ordered by calls descending then command ascending, followed by a TOTAL row |
| RPT-003 | Produce the targets report | One row per Target ordered ascending, followed by a TOTAL row |
| RPT-004 | Produce the models report | One row per provider and model pair ordered ascending, followed by a TOTAL row |
| RPT-005 | Produce the daily report | One row per start date ascending with tokens as input plus cached plus output, followed by a TOTAL row |
| RPT-006 | Produce the failures report | One row per failed record ordered by `started_at` then `execution_id`; returncode empty when null or absent |
| RPT-007 | Format minutes and cost columns | Minutes written as tenths of a minute via `t = (ms + 3000) / 6000`; cost written with `%.4f` |
| RPT-008 | Write empty-match output | A report with no matching records writes its header and, where applicable, a zero TOTAL row |

### Feature: Command Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse report, options, and log path | Report comes first, options follow, exactly one log path is last; `<report>` is one of six names |
| CLI-002 | Map usage errors to exit code 2 | Unknown report or option, missing option value, invalid date, or missing or extra log path exits 2 with no stdout |
| CLI-003 | Map unreadable logs to exit code 1 | A log that cannot be opened or read exits 1 with no stdout and one stderr line |
| CLI-004 | Report success on partial or empty data | Skipped lines and no matching records still exit 0 |

### Feature: Build and Verification

| ID | Story | High-level AC |
|---|---|---|
| BLD-001 | Provide the Go module and ddkstat entry point | Module root builds `./cmd/ddkstat` with `go build` using only the standard library |
| BLD-002 | Provide the POSIX test runner | `bin/test.sh` runs `go vet ./...` and `go test ./...` and exits zero only when both pass |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | LOG-001 | A blank line is ignored and not counted as skipped |
| AC-002 | LOG-001 | A line that is not a JSON object is skipped and counted in `skipped` |
| AC-003 | LOG-002 | Each skipped line writes exactly one stderr warning naming its line number |
| AC-004 | LOG-003 | The first occurrence of a duplicate `execution_id` is kept and later occurrences are skipped |
| AC-005 | LOG-004 | A record whose `result.stats.model` is empty or absent is attributed to `job.model` |
| AC-006 | LOG-005 | A `cost_usd` of `null` contributes zero to cost totals |
| AC-007 | FLT-001 | `--target` matching is an exact, case-sensitive comparison |
| AC-008 | FLT-002 | `--since` and `--until` apply to the start date, the first 10 characters of `started_at` |
| AC-009 | RPT-001 | The `skipped` count in `summary` is unaffected by any filter |
| AC-010 | RPT-001 | `first` and `last` are empty when no record matches |
| AC-011 | RPT-002 | Commands ordering uses byte-wise comparison for ties on call count |
| AC-012 | RPT-006 | `timed_out` is written as the literal `true` or `false` |
| AC-013 | RPT-007 | A minutes value of `95000` ms is written as `1.6` |
| AC-014 | RPT-008 | An empty match writes `0.0` minutes and `0.0000` cost in its TOTAL row |
| AC-015 | CLI-002 | `--since` and `--until` reject any value that is not a real calendar date in `YYYY-MM-DD` form |
| AC-016 | CLI-003 | A failing command writes nothing to standard output |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/llm_sample.jsonl` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/run.sh` | instruction-to-test | `sources/INSTRUCTIONS.md` | run.sh checks A1–U6 encode the rules stated in INSTRUCTIONS.md | Grader is staged verbatim; the application is built against the stated rules |
| `sources/llm_sample.jsonl` | test-kit-to-implementation | `sources/run.sh` | run.sh expected outputs are computed from the 37-line sample | Sample is staged as the grader's input; application must not embed answers |
| Log ingestion (LOG-*) | parser-to-normalizer | Filtering (FLT-*) and Reports (RPT-*) | Rules 1–4 in INSTRUCTIONS.md produce normalized records consumed by every report | Ingestion precedes filtering and reporting |
| Reports (RPT-*) | dependency | Command line interface (CLI-*) | INSTRUCTIONS.md Command and Output sections | Reports are selected and filtered through the CLI |
| Build and verification (BLD-*) | build-before-test | All implementation stories | INSTRUCTIONS.md Stack and Tests sections | Module and entry point precede the test runner |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | normative specification | compass | prompt-only |
| `sources/run.sh` | conformance harness | context | stage |
| `sources/llm_sample.jsonl` | conformance harness | context | stage |

## Planning Instructions

### Delivery Shape

The system is a single-process, read-only command-line reporter. Input is one JSON Lines log file. Output is one CSV report per invocation written to standard output, with diagnostics on standard error. Execution flow: parse arguments, open the log, decode and validate each line, apply filters, aggregate, and write the selected report. The application root is a Go module; the executable is `./cmd/ddkstat`.

### Story Realization Map

- LOG-001 through LOG-005: `sources/INSTRUCTIONS.md` Log and Rules 1–4; reads `sources/llm_sample.jsonl`.
- FLT-001, FLT-002: `sources/INSTRUCTIONS.md` Command and Rule 6; grader checks B1–B5, B3.
- RPT-001 through RPT-008: `sources/INSTRUCTIONS.md` Output and Rules 3, 7, 8; grader checks A1–A6, B2–B5, C1–C2.
- CLI-001 through CLI-004: `sources/INSTRUCTIONS.md` Command and Exit codes; grader checks E1, U1–U6.
- BLD-001, BLD-002: `sources/INSTRUCTIONS.md` Stack and Tests; grader builds via `go build` and runs checks.

### Test and Acceptance Strategy

Focused story tests are written by the application under `go test`, covering each report, filter, skip rule, and exit code over logs the tests write. Final verification is the supplied grader `sources/run.sh`, executed once as the terminal story gate. The terminal verification story declares `Suite: full`. All other stories run only their own slice and declare `Suite: scoped`. The sample is a unit input; it does not prove the product in general.

### Sequencing and Dependencies

Build module and entry point (BLD-001) before the test runner (BLD-002). Log ingestion (LOG-*) precedes filtering (FLT-*) and reports (RPT-*). CLI (CLI-*) wiring depends on reports. Terminal verification gates on every implementation story and runs `sh sources/run.sh`.

### Source Conflicts and Gaps

None. The sources are internally consistent and state the precedence explicitly where a rule applies. Unresolved items are resolved narrowly in Surfaced Acceptance Criteria above: blank-line handling, case sensitivity of target matching, and stderr content for skip warnings follow the stated rules without extension.

## Analysis Notes
generated: 2026-10-24
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.155305/workspace/targets/DrydockStats/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 5
  stories: 22
  stack: Go 1.22 or newer, standard library only
  display_name: DrydockStats
  short_description: A standard-library Go command-line reporter that reads Drydock's LLM execution log and writes six CSV reports with target and date filters.

Unattended UAT policy applied. Ambiguities resolved narrowly from the sources: unknown report names and options are usage errors, exact case-sensitive matching is used for targets, and all dates use UTC start dates from the first 10 characters of `started_at`.
