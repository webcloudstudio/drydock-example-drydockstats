# Build Instructions: A Drydock LLM Log Reporter

## Objective

Build a command-line reporter for Drydock's LLM execution log. Drydock appends one JSON record to
`logs/llm.jsonl` for every LLM call it makes. The reporter reads that file and answers the
operator's questions: how many calls each command and each Target made, which providers and
models ran them, how many tokens and minutes they used, what they cost, and which calls failed.

## Stack

Implement in Go 1.22 or newer using only the standard library. The application root is a Go
module, and its executable is the main package at `./cmd/ddkstat`. Tests run under `go test`.

## Run Harness

`sources/run.sh` is the scoring entry point. It is supplied, not authored, and is staged verbatim
into the application root with the sample log. `drydock uat` runs `sh sources/run.sh` from the
completed application root and takes its exit code and output as the score.

The harness runs `go build -o <temporary path> ./cmd/ddkstat` and then runs 20 checks against
`sources/llm_sample.jsonl`. Each check compares one command's exit code and standard output with a
known answer. Standard error is never compared. The harness prints `PASS` or `FAIL` per check and
ends with:

```text
DrydockStats: NN passed, NN failed (of 20)
```

## Read-only scoring assets

These files are the exam. Do not modify, rewrite, regenerate, or substitute them:

- `sources/run.sh` — the scoring entry point and its known answers
- `sources/llm_sample.jsonl` — 37 lines: 33 valid records over 2026-10-14 to 2026-10-16, one
  blank line, and three lines that must be skipped

Application code never reads `sources/run.sh` and never embeds a report answer. Every value is
computed from the log it is given.

## Command

```text
ddkstat <report> [--target <name>] [--since <YYYY-MM-DD>] [--until <YYYY-MM-DD>] <log>
```

`<report>` comes first, options follow it, and the log path is last. Exactly one log path is
required. `<report>` is one of `summary`, `commands`, `targets`, `models`, `daily`, `failures`.

| Option | Meaning |
|---|---|
| `--target <name>` | Keep only records whose `job.target` equals `<name>` exactly. |
| `--since <date>` | Keep only records whose start date is on or after `<date>`. |
| `--until <date>` | Keep only records whose start date is on or before `<date>`. |

A date option must be a real calendar date written `YYYY-MM-DD`.

## The log

The log is JSON Lines in UTF-8: one JSON object per line. A record carries many fields. The
reporter reads only these and ignores every other field:

| Field | Type | Use |
|---|---|---|
| `execution_id` | string | Identity of the call. |
| `started_at` | string, RFC 3339 UTC, e.g. `2026-10-14T02:38:41.904Z` | Start time. Its first 10 characters are the start date. |
| `status` | string | `succeeded`, or any other value for a failed call. |
| `job.command_name` | string | Drydock command, e.g. `build`, `plan repair`. |
| `job.target` | string | Target name. |
| `job.llm` | string | Provider, e.g. `claude`, `codex`, `gemini`. |
| `job.model` | string | The requested model. |
| `result.stats.model` | string, optional | The executed model. |
| `result.stats.input_tokens` | integer, optional | Fresh input tokens. |
| `result.stats.cached_input_tokens` | integer, optional | Cache-read input tokens. |
| `result.stats.output_tokens` | integer, optional | Output tokens. |
| `result.stats.cost_usd` | number or `null`, optional | Cost in US dollars. |
| `result.stats.elapsed_ms` | integer, optional | Wall-clock duration in milliseconds. |
| `result.returncode` | integer or `null` | Exit code of the LLM process. |
| `result.timed_out` | boolean | Whether the call timed out. |

## Rules

1. **Blank lines** are ignored and not counted.
2. **Skipped lines.** A line is skipped when it is not a JSON object, when any of `execution_id`,
   `started_at`, `status`, `job.command_name`, or `job.target` is absent or not a non-empty
   string, or when its `execution_id` already appeared on an earlier line. The first occurrence of
   an `execution_id` is kept. Each skipped line writes one warning to standard error naming its
   line number. Skipping never changes the exit code.
3. **Missing numbers.** An absent or `null` token count, cost, or duration counts as zero. A
   record with no `result.stats` object at all is still a record.
4. **Model.** A record's model is `result.stats.model` when present and non-empty, otherwise
   `job.model`.
5. **Failed.** A record is failed when `status` is not `succeeded`.
6. **Filters** apply to every report. The `skipped` count in `summary` counts skipped lines in the
   whole file regardless of filters.
7. **Minutes.** A minutes column is the group's summed `elapsed_ms` written in tenths of a minute:
   `t = (ms + 3000) / 6000` in integer division, written as `t / 10`, a `.`, and `t % 10`.
   `95000` ms is `1.6`.
8. **Cost.** A cost column is the group's summed `cost_usd` written with exactly four decimal
   places, as Go's `%.4f` writes a `float64`.
9. **Ordering.** Text is compared by byte value.

## Output

Every report writes CSV to standard output: one header row, then data rows, comma-separated, `\n`
line endings. Every report except `summary` and `failures` ends with a `TOTAL` row over all
filtered records. When no record matches, a report writes its header and, where it has one, a
`TOTAL` row of zeros (`0.0` minutes, `0.0000` cost).

| Report | Header | Rows |
|---|---|---|
| `summary` | `records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes` | One row. `first` and `last` are the earliest and latest start dates, empty when no record matches. |
| `commands` | `command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes` | One per command, by calls descending, then command ascending. Then `TOTAL,…`. |
| `targets` | `target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes` | One per Target, by Target ascending. Then `TOTAL,…`. |
| `models` | `provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd` | One per provider and model pair, by provider then model ascending. Then `TOTAL,,…`. |
| `daily` | `date,calls,failed,tokens,minutes` | One per start date, ascending. `tokens` is input plus cached plus output. Then `TOTAL,…`. |
| `failures` | `started_at,target,command,model,returncode,timed_out` | One per failed record, by `started_at` then `execution_id` ascending. `returncode` is empty when `null` or absent. `timed_out` is `true` or `false`. |

A failing command writes nothing to standard output and one diagnostic line to standard error.

## Exit codes

| Code | Meaning |
|---|---|
| `0` | Success, including a log with skipped lines or no matching records. |
| `1` | Operational failure: the log cannot be opened or read. |
| `2` | Usage error: unknown report or option, missing option value, invalid date, missing or extra log path. |

## Tests

The completed application provides a POSIX-compatible `bin/test.sh` that runs `go vet ./...` and
`go test ./...` from the application root and exits zero only when both pass. The suite covers
each report, each filter, each skip rule, and each exit code using logs written by the tests.
