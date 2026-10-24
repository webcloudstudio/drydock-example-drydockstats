# Scorecard: DrydockStats

## Specification

> **NOT YET RUN**
> drydock score spec DrydockStats

## Acceptance

> **PASSED**
> drydock score ac DrydockStats · 2026-10-24T14:28:19.314+00:00

PASS 5   FAIL 0   UNVERIFIED 0   PREPASSED 3   (8 AC total)

## Build

> **RECORDED**
> drydock score build DrydockStats

Blocks: 2   Repaired: 1   Failed: 0
Tokens: 1,021,000 (cached 849,024, output 12,905)   Cache hit rate: 84.2%

## Release

> **RELEASED**
> drydock score release DrydockStats · 2026-10-24T14:29:07.940+00:00

- Verdict: PASSED

DrydockStats: PASSED — 8 of 8

  st-001  MET       ["sh", "sources/run.sh"] exited 0; sources/run.sh: gate output: DrydockStats: 20 passed, 0 failed (of 20)
  st-002  MET       bin/test.sh: runs go vet ./... and go test ./...; probe: bin/test.sh exited zero
  st-003  MET       go.mod: module drydockstats; go 1.22; probe: go build ./cmd/ddkstat exited zero
  st-004  MET       internal/logintake/intake.go: Read counts invalid and duplicate records as skipped; probe: Read returned two valid records and skipped count 2; summary included the skip count
  st-005  MET       sources/run.sh: PASS A1 through PASS A6
  st-006  MET       sources/run.sh: PASS B1 through PASS B5; probe: Apply returned only the record matching target T on 2026-10-15
  st-007  MET       internal/reports/reports.go: model selection prefers Result.Stats.Model and falls back to Job.Model; probe: models output contained both executed and fallback models
  st-008  MET       internal/cli/run.go: unreadable logs return 1; malformed arguments return 2; probe: missing log returned 1 with zero stdout bytes; invalid date returned 2 with zero stdout bytes

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed ddkstat shall make sh sources/run.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | ["sh", "sources/run.sh"] exited 0; sources/run.sh: gate output: DrydockStats: 20 passed, 0 failed (of 20) |
| st-002 | technical | The application shall provide a POSIX-compatible bin/test.sh that exits zero when go vet and every automated test pass. | MET | bin/test.sh: runs go vet ./... and go test ./...; probe: bin/test.sh exited zero |
| st-003 | technical | The application root shall be a Go module from which go build ./cmd/ddkstat produces the ddkstat executable using only the standard library. | MET | go.mod: module drydockstats; go 1.22; probe: go build ./cmd/ddkstat exited zero |
| st-004 | behavioral | If a log line is not a JSON object, lacks a required field, or repeats an earlier execution_id, then the application shall skip it, count it in the summary, and still exit zero. | MET | internal/logintake/intake.go: Read counts invalid and duplicate records as skipped; probe: Read returned two valid records and skipped count 2; summary included the skip count |
| st-005 | behavioral | The application shall compute the summary, commands, targets, models, daily, and failures reports from the log, matching the known answers in sources/run.sh byte for byte. | MET | sources/run.sh: PASS A1 through PASS A6 |
| st-006 | behavioral | When --target, --since, or --until is given, the application shall compute every report over only the matching records. | MET | sources/run.sh: PASS B1 through PASS B5; probe: Apply returned only the record matching target T on 2026-10-15 |
| st-007 | behavioral | The application shall attribute a record to result.stats.model when it is present and to job.model otherwise. | MET | internal/reports/reports.go: model selection prefers Result.Stats.Model and falls back to Job.Model; probe: models output contained both executed and fallback models |
| st-008 | behavioral | If the log cannot be read, then the application shall exit 1; if the arguments are malformed, then the application shall exit 2; in both cases it shall write nothing to standard output. | MET | internal/cli/run.go: unreadable logs return 1; malformed arguments return 2; probe: missing log returned 1 with zero stdout bytes; invalid date returned 2 with zero stdout bytes |

**Material decisions:**

- Malformed input response contract — 
