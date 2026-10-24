# Blueprint Analysis: DrydockStats

## Commander Expectations

- assert the product reports on Drydock LLM execution logs through six CSV reports with Target and date filters

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Log Intake

| ID | Story | High-level AC |
|---|---|---|
| INTAKE-001 | Read and validate JSON Lines records | Blank lines are ignored; malformed, incomplete, and duplicate records are skipped with a line-number warning, and skipped counts are retained for summary. |
| INTAKE-002 | Apply Target and date filters | Target matching is exact; inclusive start-date filters apply to every report, while summary skipped count covers the whole file. |

### Feature: Report Generation

| ID | Story | High-level AC |
|---|---|---|
| REPORT-001 | Generate summary report | Summary contains filtered record counts, skipped count, success and failure counts, date bounds, token totals, cost, and rounded minutes. |
| REPORT-002 | Generate command and Target reports | Grouped reports include calls, failures, token totals, and minutes; Target report also includes cost, and rows follow the specified ordering. |
| REPORT-003 | Generate provider and model report | Rows group by provider and executed model when present, otherwise requested model; include calls, failures, token totals, and cost in the specified ordering. |
| REPORT-004 | Generate daily report | Rows group by start date and include calls, failures, token totals, and minutes in ascending date order. |
| REPORT-005 | Generate failures report | Failed records are listed by start time and execution ID, with nullable return code rendered empty and timeout rendered as a boolean. |

### Feature: Command-Line Delivery

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Handle command usage and operational errors | Valid invocations run the requested report; malformed arguments exit 2 and unreadable logs exit 1, with no standard output on either error. |
| CLI-002 | Provide build, test, and conformance verification | The Go module builds the command using only the standard library; `bin/test.sh` runs vet and all automated tests, and the supplied grader passes every case. |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | INTAKE-001 | Missing or null token counts, cost, and duration contribute zero; records without a stats object remain valid. |
| AC-002 | REPORT-001 | Minutes are rounded to tenths using the specified integer formula, and cost is formatted to exactly four decimal places. |
| AC-003 | REPORT-002 | Command rows sort by calls descending then command ascending; Target rows sort by Target ascending. |
| AC-004 | REPORT-003 | Provider/model rows sort by provider then model ascending; total rows and empty-result output follow the required CSV format. |
| AC-005 | REPORT-004 | Empty-result reports include headers and zero totals where the report defines a total row; summary has empty first/last dates when no records match. |
| AC-006 | CLI-001 | Every report emits CSV with `\n` line endings; failing commands emit one diagnostic line to standard error. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/llm_sample.jsonl` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/run.sh` | Instructions identify the supplied script as scoring entry point and define expected report behavior. | Implement the documented interface and output; the terminal verification story runs the supplied grader. |
| `sources/run.sh` | test-kit-to-implementation | `./cmd/ddkstat` | Grader builds the command and compares report exit codes and stdout. | Build the executable before full conformance verification. |
| `sources/llm_sample.jsonl` | dependency | `sources/run.sh` | Grader reads the synthetic sample log for its known-answer checks. | Stage the sample with the grader; it is a fixture, not an implementation story. |
| `sources/INSTRUCTIONS.md` | dependency | `bin/test.sh` | Instructions require a POSIX script running `go vet ./...` and `go test ./...`. | Implement automated tests and the test runner; the terminal verification story depends on implementation stories. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | author intent | context | prompt-only |
| `sources/llm_sample.jsonl` | asset | context | stage |
| `sources/run.sh` | conformance harness | context | stage |

## Planning Instructions

### Delivery Shape

A Go command-line reporter reads a JSON Lines execution log, skips blank and invalid or duplicate records, applies optional Target and inclusive date filters, then writes one of six CSV reports to standard output. The execution flow is argument validation, log reading and record validation, filtering, aggregation, report formatting, and output. The summary retains the whole-file skipped count.

### Story Realization Map

| Story ID | Durable Blueprint scope | Evidence | Related files | Delivery need |
|---|---|---|---|---|
| INTAKE-001 | JSON Lines parsing and validation | `sources/INSTRUCTIONS.md` rules 1–3 | `sources/llm_sample.jsonl` | Capability and focused parser tests |
| INTAKE-002 | Target and date filtering | `sources/INSTRUCTIONS.md` Command and rule 6 | `sources/llm_sample.jsonl` | Capability and focused filter tests |
| REPORT-001 | Summary aggregation and CSV output | `sources/INSTRUCTIONS.md` Output: summary | `sources/run.sh` | Capability and focused report tests |
| REPORT-002 | Command and Target aggregation and CSV output | `sources/INSTRUCTIONS.md` Output: commands and targets | `sources/run.sh` | Capability and focused report tests |
| REPORT-003 | Provider/model aggregation and CSV output | `sources/INSTRUCTIONS.md` Model rule and Output: models | `sources/run.sh` | Capability and focused report tests |
| REPORT-004 | Daily aggregation and CSV output | `sources/INSTRUCTIONS.md` Output: daily | `sources/run.sh` | Capability and focused report tests |
| REPORT-005 | Failure record selection and CSV output | `sources/INSTRUCTIONS.md` Failed rule and Output: failures | `sources/run.sh` | Capability and focused report tests |
| CLI-001 | Argument parsing, exit codes, and diagnostics | `sources/INSTRUCTIONS.md` Command and Exit codes | `sources/run.sh` | Capability and focused CLI tests |
| CLI-002 | Module build, test script, and complete grader verification | `sources/INSTRUCTIONS.md` Stack and Tests; `sources/run.sh` | `sources/run.sh`, `sources/llm_sample.jsonl` | Test harness and terminal acceptance contract |

### Test and Acceptance Strategy

Keep parser, filter, report, and CLI checks focused within their stories. CLI-002 is the sole terminal verification story; it depends on every implementation story and runs the complete supplied grader. The Sea Trial verifies the same full grader by its exit status. `bin/test.sh` independently runs vet and the complete Go test suite.

### Sequencing and Dependencies

Implement parsing before filtering and aggregation. Implement the report aggregations and CLI error handling before terminal verification. The module and executable must build before either test runner or grader runs. Stage the grader and its sample fixture together.

### Source Conflicts and Gaps

None.

## Analysis Notes
generated: 2026-10-24
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/workspace/targets/DrydockStats/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 3
  stories: 9
  stack: Go 1.22 or newer; standard library only
  display_name: not proposed
  short_description: not proposed

None.
