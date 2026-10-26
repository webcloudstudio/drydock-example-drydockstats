# drydock DrydockStats Uat Case

Drydock uat - DrydockStats builds a Go reporter for Drydock's own LLM execution log (logs/llm.jsonl). The source is an original brief: read JSON Lines, skip malformed and duplicate records without failing, and write six CSV reports (summary, commands, targets, models, daily, failures) with Target and date filters. The supplied grader, sources/run.sh, builds ./cmd/ddkstat with go build and runs 20 checks of exit code and output against a 37-line synthetic sample with known totals. Standard library only, no database. Run with drydock uat DrydockStats. Source Repo: none, the brief is original to this kit.

## Intent

Build a command-line reporter for Drydock's LLM execution log, the Go executable `./cmd/ddkstat`,
as described in `sources/INSTRUCTIONS.md`. It reads a JSON Lines log, skips malformed and
duplicate lines without failing, and writes six CSV reports (summary, commands, targets, models,
daily, failures), each narrowed by optional Target and date filters.

Correctness is measured by `sources/run.sh`: it builds the executable with `go build` and makes
20 checks of exit code and standard output against `sources/llm_sample.jsonl`. The goal is every
check passing. The finished tool also runs against a real `logs/llm.jsonl`.

## What It Does

The Go module root builds the `ddkstat` executable with `go build ./cmd/ddkstat`, using only the standard library.

`internal/cli` turns the raw argument list into an `Options` value or a usage error. It performs no I/O and
never exits the process. Date value validation is owned by FEATURE-CliDates.md and layered onto this parser.

A date option must be a real calendar date written `YYYY-MM-DD`. Anything else is a usage error that maps to
exit code 2 in FEATURE-CliExit.md.

`cli.Run` is the single entry point behind `cmd/ddkstat/main.go`. It runs the pipeline: parse arguments, read the
log, apply filters, dispatch to the registered report writer. Report stories plug in through a registry so `Run`
never imports them.

A pure function in `internal/filter` that narrows a record slice by start date. Dates are already validated
`YYYY-MM-DD` text, so string comparison orders them correctly.

A pure function in `internal/filter` that narrows a record slice by Target. It performs no I/O and knows nothing
about reports; every report receives its output through the CLI pipeline.

Turn a JSON Lines stream into `logread.Record` values. Each non-blank UTF-8 line is one JSON object. This story establishes the reader, the `Record` type and the failed-call classification; later stories add skip rules, numeric normalization and model attribution.

Reports group by model, and the model that actually ran is more truthful than the one requested.

Aggregation sums these fields, so each record must carry a usable number even when the log omits it.

Make the reader tolerant. Bad lines are skipped and counted; they never fail the run.

`ddkstat commands [options] <log>` shows how often each Drydock command called an LLM and what it consumed.

The `daily` report shows, for each start date, how many calls ran, how many failed, how many tokens they used, and how many minutes they took.

The `failures` report lists every failed call so the operator can see which command, Target and model failed, and how.

`ddkstat models [options] <log>` shows which providers and models ran the calls and what they consumed and cost.

`ddkstat summary [options] <log>` answers how many calls ran, how many were skipped, and what they used and cost.

`ddkstat targets [options] <log>` shows what each Target consumed and cost.

One command, `sh bin/test.sh`, gates every behavioral story and the Go test suite.

Every report that prints a minutes column or a cost column uses these two pure functions in `internal/format`, so the rounding rule lives in exactly one place.

## Stack

Go

## Verification

```bash
bash bin/test.sh
```

## Next Steps

- Complete the first-run checks above.
- Run `bash bin/test.sh` before making release or deployment changes.
