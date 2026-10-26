# Sea Trials: DrydockStats

## Policy

| Consequence | On FAIL | On INCONCLUSIVE |
|---|---|---|
| blocks  | fail   | attest |
| scores  | score  | score  |
| attests | report | report |

## st-001: The supplied grader passes
Type: technical
Required: yes
Criterion: The completed ddkstat shall make sh sources/run.sh exit zero; that script's exit status is the sole acceptance verdict.
Testability: deterministic
Consequence: blocks
Verification: proof
Command: ["sh", "sources/run.sh"]
Pattern: ubiquitous

## st-002: Test suite passes
Type: technical
Required: yes
Criterion: The application shall provide a POSIX-compatible bin/test.sh that exits zero when go vet and every automated test pass.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-003: The executable builds with go build
Type: technical
Required: yes
Criterion: The application root shall be a Go module from which go build ./cmd/ddkstat produces the ddkstat executable using only the standard library.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-004: Malformed and duplicate lines are skipped, not fatal
Type: behavioral
Required: yes
Criterion: If a log line is not a JSON object, lacks a required field, or repeats an earlier execution_id, then the application shall skip it, count it in the summary, and still exit zero.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: unwanted

## st-005: The six reports match the known answers
Type: behavioral
Required: yes
Criterion: The application shall compute the summary, commands, targets, models, daily, and failures reports from the log, matching the known answers in sources/run.sh byte for byte.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-006: Filters narrow every report
Type: behavioral
Required: yes
Criterion: When --target, --since, or --until is given, the application shall compute every report over only the matching records.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: event

## st-007: The executed model is preferred
Type: behavioral
Required: yes
Criterion: The application shall attribute a record to result.stats.model when it is present and to job.model otherwise.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-008: Errors map to exit codes
Type: behavioral
Required: yes
Criterion: If the log cannot be read, then the application shall exit 1; if the arguments are malformed, then the application shall exit 2; in both cases it shall write nothing to standard output.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: unwanted
