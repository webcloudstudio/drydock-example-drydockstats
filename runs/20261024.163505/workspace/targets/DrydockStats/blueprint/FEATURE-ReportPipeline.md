# Feature: Report Pipeline

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Assemble the end-to-end ddkstat run from arguments to one written report. |
| Depends On  | ARCHITECTURE.md, FEATURE-ReadErrors.md, FEATURE-UsageErrors.md, FEATURE-FilterValidate.md, FEATURE-SummaryReport.md, FEATURE-CommandsReport.md, FEATURE-TargetsReport.md, FEATURE-ModelsReport.md, FEATURE-DailyReport.md, FEATURE-FailuresReport.md |
| Provides    | ddkstat <report> [options] <log> end-to-end |
| Consumes    | argument parser, reader, parser, filters, six reports, usage and read-failure paths |

## Purpose

Assembly feature. A single run parses arguments, reads the log, skips invalid lines, applies the active filters, aggregates, and writes exactly one selected report to standard output. Diagnostics go to standard error. Exit codes follow the contract in `sources/INSTRUCTIONS.md`: `0` for success, `1` for an unreadable log, `2` for a usage error.

## Programmatic Acceptance

=== AC pipeline-each-report-exits-zero ===
Intent: Every one of the six reports runs end to end over the sample and exits 0 with non-empty stdout.
Suite: scoped
Requires: executable=python3; scope=test

import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
for report in ["summary", "commands", "targets", "models", "daily", "failures"]:
    code, out, err = cli_runner(report, sample)
    print(report, code, err)
    assert code == 0, (report, code, out, err)
    assert out.strip() != "", report
=== END AC pipeline-each-report-exits-zero ===

=== AC pipeline-usage-error-writes-nothing ===
Intent: An unknown report exits 2 and writes nothing to standard output.
Suite: scoped
Requires: executable=python3; scope=test

import os

from drydock_go import cli_runner

sample = os.path.join("sources", "llm_sample.jsonl")
code, out, err = cli_runner("frobnicate", sample)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC pipeline-usage-error-writes-nothing ===

=== AC pipeline-unreadable-log-writes-nothing ===
Intent: A log path that does not exist exits 1 and writes nothing to standard output.
Suite: scoped
Requires: executable=python3; scope=test

import os
import tempfile

from drydock_go import cli_runner

missing = os.path.join(tempfile.mkdtemp(), "no-such-log.jsonl")
code, out, err = cli_runner("summary", missing)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
=== END AC pipeline-unreadable-log-writes-nothing ===

## User Acceptance

- None.

## Guardrails

- Only the selected report reaches standard output; diagnostics never do.
- Stdout is written only after aggregation succeeds.
