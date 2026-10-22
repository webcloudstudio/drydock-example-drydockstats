# DrydockStats

A Drydock UAT kit that builds a Go reporter for Drydock's own LLM execution log. The brief is
original to this kit. Complexity: 2 of 5, about 5 to 7 stories.

| Path | Role |
|---|---|
| `sources/INSTRUCTIONS.md` | The specification imported into the Target. |
| `sources/llm_sample.jsonl` | Synthetic `schema_version` 1 log: 33 valid records, 3 skipped lines, 1 blank line. |
| `bin/run.sh` | The grader. Staged into the build as `sources/run.sh`; `acceptance.full` runs it. |
| `tools/` | The sample generator and the reference that produced the answers. Not imported. |
| `inputs/` | Frozen `SEA_TRIALS.md`, `TECHNOLOGY_STACK.md`, `COMPASS.md`, `METADATA.md`. |

## Grading

`bin/run.sh [APP]` runs `go build ./cmd/ddkstat` in `APP` (default: the current directory) and
makes 20 checks of exit code and standard output:

| Group | Checks |
|---|---|
| A | all six reports over the whole sample |
| B | `--target`, `--since`, `--until`, alone and combined |
| C | a filter that matches nothing |
| E | unreadable log |
| U | usage errors |

## The sample

Three Targets (CommonMark, ReadingList, Toml) over three days, on providers `claude`, `codex`, and
`gemini`. Known traps:

| Trap | Line content |
|---|---|
| Executed vs requested model | `codex` records carry no `result.stats.model`; one `claude` record has no `result.stats` at all and reports as `sonnet`. |
| Null cost | `codex` records have `cost_usd: null`. |
| Skipped lines | a truncated JSON line, a repeated `execution_id`, a record missing `job.target`. |
| Failures | one exit `1`; one timeout with `returncode: null`. |
| Cost rounding | the ReadingList cost sums to `1.2485499` (`1.2485`); rounding each record first gives `1.2487`. |

After the build, `./ddkstat targets <drydock>/logs/llm.jsonl` reports on the real log.
