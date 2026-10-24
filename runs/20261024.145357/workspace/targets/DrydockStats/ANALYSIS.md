# Blueprint Analysis: DrydockStats

## Commander Expectations

- assert the product reads Drydock's LLM execution log (`logs/llm.jsonl`) and reports on it
- assert the product runs with the standard library only and no database
- assert malformed and duplicate log lines never fail the run
- assert the six CSV reports match the supplied known answers
- assert the supplied grader `sources/run.sh` passes unmodified

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
| LOG-001 | Read JSON Lines log records from a file path | Each non-blank line is decoded as one JSON object; the file is read as UTF-8 |
| LOG-002 | Skip blank lines without counting them | Blank lines are ignored and not counted in `skipped` |
| LOG-003 | Skip lines that are not JSON objects | A non-object line is skipped, counted in `skipped`, and warned on stderr with its line number; exit code is unchanged |
| LOG-004 | Skip lines missing required fields | A line lacking any of `execution_id`, `started_at`, `status`, `job.command_name`, `job.target` (or with a non-string or empty value) is skipped and counted |
| LOG-005 | Skip duplicate execution IDs | A line whose `execution_id` already appeared is skipped; the first occurrence is kept |
| LOG-006 | Treat missing or null numeric stats as zero | Absent or null token, cost, and duration values contribute zero; a record without `result.stats` is still a record |
| LOG-007 | Resolve the record model | Model is `result.stats.model` when present and non-empty, otherwise `job.model` |
| LOG-008 | Classify failed records | A record with `status` other than `succeeded` is failed |

### Feature: Filters

| ID | Story | High-level AC |
|---|---|---|
| FLT-001 | Filter records by target | `--target <name>` keeps only records whose `job.target` equals the name exactly |
| FLT-002 | Filter records by start date lower bound | `--since <date>` keeps records whose start date (first 10 characters of `started_at`) is on or after the date |
| FLT-003 | Filter records by start date upper bound | `--until <date>` keeps records whose start date is on or before the date |
| FLT-004 | Apply filters to every report | Each report is computed over only the filtered records; `summary.skipped` counts skipped lines over the whole file |

### Feature: Reports

| ID | Story | High-level AC |
|---|---|---|
| RPT-001 | Write the summary report | One CSV row with `records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes` |
| RPT-002 | Write the commands report | One row per command ordered by calls descending then command ascending, followed by a `TOTAL` row |
| RPT-003 | Write the targets report | One row per target ordered ascending, followed by a `TOTAL` row |
| RPT-004 | Write the models report | One row per provider and model pair ordered ascending, followed by a `TOTAL,,…` row |
| RPT-005 | Write the daily report | One row per start date ascending with `tokens` = input + cached + output, followed by a `TOTAL` row |
| RPT-006 | Write the failures report | One row per failed record ordered by `started_at` then `execution_id`; `returncode` empty when null or absent |
| RPT-007 | Format minutes and cost columns | Minutes are `t = (ms + 3000) / 6000` written as `t/10.t%10`; cost is written with four decimal places (`%.4f`) |
| RPT-008 | Emit zero-match reports | A report with no matching records writes its header and, where applicable, a zero `TOTAL` row |

### Feature: Command Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse the report name and log path | `<report>` comes first, options follow, exactly one log path is last; `<report>` must be one of the six reports |
| CLI-002 | Validate option values | Unknown options, missing option values, and invalid `YYYY-MM-DD` calendar dates are usage errors |
| CLI-003 | Map errors to exit codes | Unreadable log exits 1 with nothing on stdout; usage errors exit 2 with nothing on stdout; success exits 0 |
| CLI-004 | Build the ddkstat executable | `go build ./cmd/ddkstat` produces the executable from a Go 1.22+ module using only the standard library |

### Feature: Verification

| ID | Story | High-level AC |
|---|---|---|
| VER-001 | Provide the POSIX test entry point | `bin/test.sh` runs `go vet ./...` and `go test ./...` from the application root and exits zero only when both pass |
| VER-002 | Cover reports, filters, skips, and exit codes with tests | Automated tests write their own logs and cover each report, filter, skip rule, and exit code |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | LOG-003 | A skipped non-object line writes exactly one warning to stderr naming its line number |
| AC-002 | LOG-004 | The skipped-line rule applies to `execution_id`, `started_at`, `status`, `job.command_name`, and `job.target` |
| AC-003 | RPT-007 | `95000` ms renders as `1.6` minutes |
| AC-004 | RPT-002 | The `commands` report TOTAL row sums all filtered records |
| AC-005 | CLI-003 | A failing command writes one diagnostic line to stderr and nothing to stdout |
| AC-006 | CLI-001 | Exactly one log path is accepted; zero or two paths is a usage error |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/llm_sample.jsonl` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/run.sh` | test-kit-to-implementation | ddkstat executable | `sources/run.sh` builds `./cmd/ddkstat` and runs 20 checks | Executable must exist and match known answers before acceptance |
| `sources/llm_sample.jsonl` | dependency | `sources/run.sh` | `LOG="$DATA/llm_sample.jsonl"` | Sample must be staged with the grader |
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/run.sh` | Rules 1–9 and output tables define the checked behavior | Implementation follows the instruction rules exactly |
| Log ingestion (LOG-*) | parser-to-normalizer | Filters (FLT-*) and Reports (RPT-*) | Rule 2 skip logic feeds `summary.skipped` and every report | Ingestion precedes filtering and reporting |
| Reports (RPT-*) | dependency | Log ingestion (LOG-*) | Every report reads normalized records | Reports depend on ingestion |
| CLI (CLI-*) | dependency | Reports (RPT-*) | `<report>` selects the report implementation | CLI dispatches to reports |
| Verification (VER-*) | test-kit-to-implementation | All feature groups | `bin/test.sh` runs `go test ./...` | Test suite must be written after the implementation it covers |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | author intent | compass | prompt-only |
| `sources/run.sh` | conformance harness | context | stage |
| `sources/llm_sample.jsonl` | test helper | context | stage |

## Planning Instructions

### Delivery Shape

The system is a single-process, read-only CLI reporter. Inputs: one JSON Lines log path, one report name, and up to three filter options. Output: one CSV report on stdout; diagnostics on stderr; exit code 0, 1, or 2. Execution flow: parse arguments → open and read log → decode and validate lines → apply filters → compute the selected report → write CSV.

### Story Realization Map

- LOG-001 through LOG-008: `sources/INSTRUCTIONS.md` Log and Rules 1–5; evidence in `sources/llm_sample.jsonl`. Capability, not integration.
- FLT-001 through FLT-004: `sources/INSTRUCTIONS.md` Command table and Rule 6; checked by `sources/run.sh` B1–B5, C1–C2.
- RPT-001 through RPT-008: `sources/INSTRUCTIONS.md` Output table and Rules 7–9; checked by `sources/run.sh` A1–A6, C1–C2.
- CLI-001 through CLI-004: `sources/INSTRUCTIONS.md` Command, Stack, and Exit codes; checked by `sources/run.sh` U1–U6, E1.
- VER-001 and VER-002: `sources/INSTRUCTIONS.md` Tests; capability for the test harness.

### Test and Acceptance Strategy

Focused story tests run per story with their own written logs. The terminal verification story VER-001 gates on the full suite and declares `Suite: full`. Feature stories run only the slice they own and declare `Suite: scoped`. The supplied grader `sources/run.sh` is the acceptance verdict and is covered by `st-001`.

### Sequencing and Dependencies

Build order: LOG → FLT → RPT → CLI → VER. The CLI depends on reports, and reports depend on ingestion and filters. The terminal VER-001 depends on every implementation story.

### Source Conflicts and Gaps

None. The sources define one consistent behavior. Where the sources are silent, the narrowest behavior consistent with them is recorded: stderr diagnostics for skipped lines use the message format named in Rule 2; the test suite covers only behaviors the sources name.

## Analysis Notes
generated: 2026-10-24
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 5
  stories: 29
  stack: Go 1.22+ (standard library only)
  display_name: not proposed
  short_description: not proposed

None.
