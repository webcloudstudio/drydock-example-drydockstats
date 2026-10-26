# Scorecard: DrydockStats

## Specification

> **NOT YET RUN**
> drydock score spec DrydockStats

## Acceptance

> **PASSED**
> drydock score ac DrydockStats · 2026-10-26T16:03:57.282+00:00

PASS 83   FAIL 0   UNVERIFIED 0   PREPASSED 8   (91 AC total)

## Build

> **RECORDED**
> drydock score build DrydockStats

Blocks: 10   Repaired: 0   Failed: 0
Tokens: 3,025,929 (cached 2,897,865, output 127,912)   Cache hit rate: 100.0%

## Release

> **RELEASED**
> drydock score release DrydockStats · 2026-10-26T16:05:00.626+00:00

- Verdict: PASSED

DrydockStats: PASSED — 8 of 8

  st-001  MET       ["sh", "sources/run.sh"] exited 0; governed_gate full: sh sources/run.sh return_code 0, 20 passed, 0 failed; sources/run.sh compares every report with cmp and exits 0 only when all checks pass
  st-002  MET       bin/test.sh: #!/bin/sh, set -eu, runs go vet ./... then go test ./...; probe run: ok ddkstat/cmd/ddkstat, internal/cli, internal/filter, internal/format, internal/logread, internal/report; rc=0; test results were reported as (cached); the cache is keyed on file contents, so they reflect the current tree
  st-003  MET       go.mod: module ddkstat, go 1.22; go build -o ... ./cmd/ddkstat rc=0; go list -deps ./cmd/ddkstat: no dotted import paths, so stdlib only; cmd/ddkstat/main.go imports only ddkstat/internal/* and os
  st-004  MET       grader judgement
  st-005  MET       gate stdout: PASS A1 summary, PASS A2 commands, PASS A3 targets, PASS A4 models, PASS A5 daily, PASS A6 failures; sources/run.sh check() uses cmp -s on stdout and the expected exit code
  st-006  MET       gate: PASS B1 summary --target --since, PASS B2 commands --until, PASS B3 daily --since --until, PASS B4 failures --target, PASS B5 models --target; probe: `commands --target NoSuch` -> TOTAL,0,...; `daily --until 2026-10-01` -> only 2026-10-01 row; internal/cli/run.go applies filter.ByTarget and filter.ByDateRange before every report writer
  st-007  MET       probe: `models mixed.jsonl` -> `claude,claude-x,1,...` (record a, stats.model present) and `claude,haiku-job,1,...` (record c, no stats.model); internal/logread/logread.go parseLine: model = stats.model if non-empty, otherwise job.model; gate A4 also passes with the known models table
  st-008  MET       probe: `summary nope.jsonl` -> rc=1 stdout_bytes=0, stderr 'cannot open log'; probe: `summary --since 2026-1-5 mixed.jsonl` -> rc=2 stdout_bytes=0, stderr usage line; gate E1 (rc 1) and U1–U6 (rc 2) all PASS with empty expected stdout

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed ddkstat shall make sh sources/run.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | ["sh", "sources/run.sh"] exited 0; governed_gate full: sh sources/run.sh return_code 0, 20 passed, 0 failed; sources/run.sh compares every report with cmp and exits 0 only when all checks pass |
| st-002 | technical | The application shall provide a POSIX-compatible bin/test.sh that exits zero when go vet and every automated test pass. | MET | bin/test.sh: #!/bin/sh, set -eu, runs go vet ./... then go test ./...; probe run: ok ddkstat/cmd/ddkstat, internal/cli, internal/filter, internal/format, internal/logread, internal/report; rc=0; test results were reported as (cached); the cache is keyed on file contents, so they reflect the current tree |
| st-003 | technical | The application root shall be a Go module from which go build ./cmd/ddkstat produces the ddkstat executable using only the standard library. | MET | go.mod: module ddkstat, go 1.22; go build -o ... ./cmd/ddkstat rc=0; go list -deps ./cmd/ddkstat: no dotted import paths, so stdlib only; cmd/ddkstat/main.go imports only ddkstat/internal/* and os |
| st-004 | behavioral | If a log line is not a JSON object, lacks a required field, or repeats an earlier execution_id, then the application shall skip it, count it in the summary, and still exit zero. | MET | grader judgement |
| st-005 | behavioral | The application shall compute the summary, commands, targets, models, daily, and failures reports from the log, matching the known answers in sources/run.sh byte for byte. | MET | gate stdout: PASS A1 summary, PASS A2 commands, PASS A3 targets, PASS A4 models, PASS A5 daily, PASS A6 failures; sources/run.sh check() uses cmp -s on stdout and the expected exit code |
| st-006 | behavioral | When --target, --since, or --until is given, the application shall compute every report over only the matching records. | MET | gate: PASS B1 summary --target --since, PASS B2 commands --until, PASS B3 daily --since --until, PASS B4 failures --target, PASS B5 models --target; probe: `commands --target NoSuch` -> TOTAL,0,...; `daily --until 2026-10-01` -> only 2026-10-01 row; internal/cli/run.go applies filter.ByTarget and filter.ByDateRange before every report writer |
| st-007 | behavioral | The application shall attribute a record to result.stats.model when it is present and to job.model otherwise. | MET | probe: `models mixed.jsonl` -> `claude,claude-x,1,...` (record a, stats.model present) and `claude,haiku-job,1,...` (record c, no stats.model); internal/logread/logread.go parseLine: model = stats.model if non-empty, otherwise job.model; gate A4 also passes with the known models table |
| st-008 | behavioral | If the log cannot be read, then the application shall exit 1; if the arguments are malformed, then the application shall exit 2; in both cases it shall write nothing to standard output. | MET | probe: `summary nope.jsonl` -> rc=1 stdout_bytes=0, stderr 'cannot open log'; probe: `summary --since 2026-1-5 mixed.jsonl` -> rc=2 stdout_bytes=0, stderr usage line; gate E1 (rc 1) and U1–U6 (rc 2) all PASS with empty expected stdout |

**Ranked improvements:**
1. Add malformed and duplicate fixtures to the grader (or to bin/test.sh) so st-004 is enforced by a gate, not only by an ad hoc probe; the shipped sample has only 3 skips and never exercises a missing required field.
2. Add a --target or --since/--until check for the targets report in run.sh; B1–B5 cover summary, commands, daily, failures and models but not targets, though the filter is applied before every writer.
3. Add a fixture where result.stats.model is absent so the job.model fallback in st-007 is pinned by the gate, not only by the known-answers table.
