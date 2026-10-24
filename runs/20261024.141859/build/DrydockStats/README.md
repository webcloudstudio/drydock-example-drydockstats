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

Deliver the `ddkstat` command and project test runner. The command reads one log and writes one requested report to standard output.

Provide validated execution records to the reporting layer from a JSON Lines log. Blank lines are ignored. Malformed, incomplete, and duplicate nonblank lines are skipped with one warning per skipped line, and contribute to the whole-file skipped count.

Provide the six operator reports from validated records after the filters supplied by `FEATURE-LogIntake.md`. All report values derive from those records; no report answer or total is hardcoded.

## Stack

Go

## Verification

```bash
bash bin/test.sh
```

## Next Steps

- Complete the first-run checks above.
- Run `bash bin/test.sh` before making release or deployment changes.
