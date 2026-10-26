# Evidence: Block 4 · Service (block-4)

- block type: block
- date: 2026-10-25
- resulting state: closed/failed
- story points (combined assembled cost): 12961
- execution id: 20261026.015000.461Z-1930d24a

## Stories built
- Keep only records whose job.target equals --target exactly. (filter-target) [story]
- Keep records whose start date is within --since and --until inclusive. (filter-dates) [story]

## Acceptance tooling authorization
- FEATURE-TargetFilter.md#target-summary-counts: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TargetFilter.md#target-exact-match: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TargetFilter.md#target-no-match: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TargetFilter.md#target-every-report: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateFilter.md#date-since-until-inclusive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateFilter.md#date-one-sided-bounds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateFilter.md#date-time-of-day-ignored: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateFilter.md#date-every-report: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateFilter.md#date-with-target-and-empty-range: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-TargetFilter.md (SP 3292)
- context: ARCHITECTURE_compact.md (SP 387)
- stack: common_compact.md (SP 1653)
- implements: FEATURE-DateFilter.md (SP 4379)

## Build directory changes
- internal/report/filter.go
- internal/report/filter_test.go

## Pre-build acceptance observation
- RED: target-summary-counts (FEATURE-TargetFilter.md)
  intent: The summary report counts only the named Target's records while skipped still counts the whole file.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-summary-counts.py", line 25, in <module>
        code, out = run(recs + bad_lines, "summary", "--target", name)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-summary-counts.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: target-exact-match (FEATURE-TargetFilter.md)
  intent: The Target comparison is exact and case-sensitive, so near-miss names are excluded.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-exact-match.py", line 25, in <module>
        code, out = run(recs, "targets", "--target", name)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-exact-match.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: target-no-match (FEATURE-TargetFilter.md)
  intent: A Target matching no record yields a header and a zero TOTAL row with exit zero.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-no-match.py", line 23, in <module>
        code, out = run(recs, "commands", "--target", "NoSuchTarget")
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-no-match.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: target-every-report (FEATURE-TargetFilter.md)
  intent: The Target filter narrows all six reports, not only one of them.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-every-report.py", line 28, in <module>
        code, out = run(recs, report, "--target", name)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-every-report.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: date-since-until-inclusive (FEATURE-DateFilter.md)
  intent: Both bounds are inclusive and records outside them are excluded.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-since-until-inclusive.py", line 25, in <module>
        code, out = run(recs, "daily", "--since", since, "--until", until)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-since-until-inclusive.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: date-one-sided-bounds (FEATURE-DateFilter.md)
  intent: --since alone and --until alone each bound only one side.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-one-sided-bounds.py", line 26, in <module>
        code, out = run(recs, "summary", "--since", bound)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-one-sided-bounds.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: date-time-of-day-ignored (FEATURE-DateFilter.md)
  intent: Only the date part of started_at decides membership, not the time of day.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-time-of-day-ignored.py", line 24, in <module>
        code, out = run(recs, "summary", "--since", day, "--until", day)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-time-of-day-ignored.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: date-every-report (FEATURE-DateFilter.md)
  intent: The date filter narrows all six reports and leaves the skipped count untouched.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-every-report.py", line 30, in <module>
        code, out = run(recs + bad, report, *args)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-every-report.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: date-with-target-and-empty-range (FEATURE-DateFilter.md)
  intent: Date and Target filters combine, and a range matching nothing yields a zero TOTAL row.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-with-target-and-empty-range.py", line 24, in <module>
        code, out = run(recs, "commands", "--target", name, "--since", day, "--until", day)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-with-target-and-empty-range.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- FAIL: target-summary-counts (FEATURE-TargetFilter.md)
  intent: The summary report counts only the named Target's records while skipped still counts the whole file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-summary-counts.py", line 25, in <module>
        code, out = run(recs + bad_lines, "summary", "--target", name)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-summary-counts.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: target-exact-match (FEATURE-TargetFilter.md)
  intent: The Target comparison is exact and case-sensitive, so near-miss names are excluded.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-exact-match.py", line 25, in <module>
        code, out = run(recs, "targets", "--target", name)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-exact-match.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: target-no-match (FEATURE-TargetFilter.md)
  intent: A Target matching no record yields a header and a zero TOTAL row with exit zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-no-match.py", line 23, in <module>
        code, out = run(recs, "commands", "--target", "NoSuchTarget")
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-no-match.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: target-every-report (FEATURE-TargetFilter.md)
  intent: The Target filter narrows all six reports, not only one of them.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "target-every-report.py", line 28, in <module>
        code, out = run(recs, report, "--target", name)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "target-every-report.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: date-since-until-inclusive (FEATURE-DateFilter.md)
  intent: Both bounds are inclusive and records outside them are excluded.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-since-until-inclusive.py", line 25, in <module>
        code, out = run(recs, "daily", "--since", since, "--until", until)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-since-until-inclusive.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: date-one-sided-bounds (FEATURE-DateFilter.md)
  intent: --since alone and --until alone each bound only one side.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-one-sided-bounds.py", line 26, in <module>
        code, out = run(recs, "summary", "--since", bound)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-one-sided-bounds.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: date-time-of-day-ignored (FEATURE-DateFilter.md)
  intent: Only the date part of started_at decides membership, not the time of day.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-time-of-day-ignored.py", line 24, in <module>
        code, out = run(recs, "summary", "--since", day, "--until", day)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-time-of-day-ignored.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: date-every-report (FEATURE-DateFilter.md)
  intent: The date filter narrows all six reports and leaves the skipped count untouched.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-every-report.py", line 30, in <module>
        code, out = run(recs + bad, report, *args)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-every-report.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError
- FAIL: date-with-target-and-empty-range (FEATURE-DateFilter.md)
  intent: Date and Target filters combine, and a range matching nothing yields a zero TOTAL row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-with-target-and-empty-range.py", line 24, in <module>
        code, out = run(recs, "commands", "--target", name, "--since", day, "--until", day)
                    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "date-with-target-and-empty-range.py", line 17, in run
        assert code == 0
               ^^^^^^^^^
    AssertionError

## Repair attempts
- attempt 0 (initial build): failed; 0/9 checks model=claude-haiku-5-5; execution 20261026.014747.893Z-5f32a598; reason: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report, date-since-until-inclusive, date-one-sided-bou…
- attempt 1 (repair 1): failed; 0/9 checks model=claude-haiku-5-5; execution 20261026.014909.776Z-88dd7a72; reason: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report, date-since-until-inclusive, date-one-sided-bou…
- attempt 2 (repair 2): failed; 0/9 checks model=claude-haiku-5-5; execution 20261026.015000.461Z-1930d24a; stopped: deterministic acceptance score did not improve on 2 consecutive calls; reason: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report, date-since-until-inclusive, date-one-sided-bou…

## Agent self-report (advisory)
The build agent declared a failure. This report is advisory; Drydock's execution and deterministic checks determine the block's outcome.
- summary: Block 4 checks depend on Block 5 report commands, which are not built.
- detail:
    All nine failing checks run report subcommands through the CLI. `internal/cli/cli.go` returns exit 1 with "not implemented yet" because no report aggregation exists. Those reports are specified in the Block 5 features, so they belong to a later block. Before rerunning Block 4, build Block 5, then route the CLI to the reports and apply the two filters before aggregation. Alternatively, explicitly add the reports to Block 4 and update the Manifest.

## Failure
- summary: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report, date-since-until-inclusive, date-one-sided-bounds, date-time-of-day-ignored, date-every-report, date-with-target-and-empty-range
- detail:
    Block "Block 4 · Service" [block-4] failed its acceptance criteria.
      Story "Keep only records whose job.target equals --target exactly." [filter-target] does not meet its own acceptance criteria:
        - AC target-summary-counts — The summary report counts only the named Target's records while skipped still counts the whole file.
            assertion: assert code == 0 → AssertionError
            raised at: target-summary-counts.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "target-summary-counts.py", line 25, in <module>
                  code, out = run(recs + bad_lines, "summary", "--target", name)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "target-summary-counts.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC target-exact-match — The Target comparison is exact and case-sensitive, so near-miss names are excluded.
            assertion: assert code == 0 → AssertionError
            raised at: target-exact-match.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "target-exact-match.py", line 25, in <module>
                  code, out = run(recs, "targets", "--target", name)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "target-exact-match.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC target-no-match — A Target matching no record yields a header and a zero TOTAL row with exit zero.
            assertion: assert code == 0 → AssertionError
            raised at: target-no-match.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "target-no-match.py", line 23, in <module>
                  code, out = run(recs, "commands", "--target", "NoSuchTarget")
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "target-no-match.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC target-every-report — The Target filter narrows all six reports, not only one of them.
            assertion: assert code == 0 → AssertionError
            raised at: target-every-report.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "target-every-report.py", line 28, in <module>
                  code, out = run(recs, report, "--target", name)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "target-every-report.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
      Story "Keep records whose start date is within --since and --until inclusive." [filter-dates] does not meet its own acceptance criteria:
        - AC date-since-until-inclusive — Both bounds are inclusive and records outside them are excluded.
            assertion: assert code == 0 → AssertionError
            raised at: date-since-until-inclusive.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "date-since-until-inclusive.py", line 25, in <module>
                  code, out = run(recs, "daily", "--since", since, "--until", until)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "date-since-until-inclusive.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC date-one-sided-bounds — --since alone and --until alone each bound only one side.
            assertion: assert code == 0 → AssertionError
            raised at: date-one-sided-bounds.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "date-one-sided-bounds.py", line 26, in <module>
                  code, out = run(recs, "summary", "--since", bound)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "date-one-sided-bounds.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC date-time-of-day-ignored — Only the date part of started_at decides membership, not the time of day.
            assertion: assert code == 0 → AssertionError
            raised at: date-time-of-day-ignored.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "date-time-of-day-ignored.py", line 24, in <module>
                  code, out = run(recs, "summary", "--since", day, "--until", day)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "date-time-of-day-ignored.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC date-every-report — The date filter narrows all six reports and leaves the skipped count untouched.
            assertion: assert code == 0 → AssertionError
            raised at: date-every-report.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "date-every-report.py", line 30, in <module>
                  code, out = run(recs + bad, report, *args)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "date-every-report.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError
        - AC date-with-target-and-empty-range — Date and Target filters combine, and a range matching nothing yields a zero TOTAL row.
            assertion: assert code == 0 → AssertionError
            raised at: date-with-target-and-empty-range.py:17
            process exit code: 1
            values at failure:
              code = 1
            check stderr:
              Traceback (most recent call last):
                File "date-with-target-and-empty-range.py", line 24, in <module>
                  code, out = run(recs, "commands", "--target", name, "--since", day, "--until", day)
                              ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                File "date-with-target-and-empty-range.py", line 17, in run
                  assert code == 0
                         ^^^^^^^^^
              AssertionError

## Build summary
This repair prompt matches the previous pass. I'll check whether anything has changed in the build directory since then, such as Block 5 landing.RESULT: FAILED

Nothing has changed since the last pass, so I made no edits. The nine failing checks still cannot pass in Block 4.

**Current state**
- The Drydock checkpoint `1410985 Build block-4 … [failed]` holds the filter work: `internal/report/filter.go` and `filter_test.go`.
- `go test ./...` passes for all packages, including the 12 filter unit tests.
- `internal/cli/cli.go:122` still prints `report "<name>" is not implemented yet` and exits 1. No report code exists yet.

**Why the checks fail**
- Every failing check calls `summary`, `commands`, `targets`, `models`, `daily`, or `failures` through the CLI.
- Those six reports are owned by Block 5 in the Manifest (`FEATURE-SummaryReport`, `FailuresReport`, `DailyReport`, `ModelsReport`, `CommandsReport`, `TargetsReport`).
- The failing checks are correct, so this is not `AC_BROKEN`. They need Block 5 code.

**What would unblock it**
1. Build Block 5 first. Then route the CLI to the reports and apply `report.FilterTarget` and `report.FilterDates` before aggregation. This is the scope-clean route.
2. Or explicitly widen Block 4 to include the six reports. That mixes two Manifest blocks, so the Manifest should be updated to match.

I did not take option 2 without your decision, because it goes beyond "implement only this step".

FILES CHANGED:
- None in this pass.

SUMMARY:
Repair pass re-checked the build directory. The filter functions and their unit tests are complete and green. The nine declared Block 4 checks exit 1 because the report commands they call are Block 5 features and are not yet implemented.

BLOCKERS:
- The six report commands (summary, commands, targets, models, daily, failures) are Block 5 features and do not exist yet. Block 4's declared checks cannot pass until they do.

FAILURE_SUMMARY: Block 4 checks depend on Block 5 report commands, which are not built.
FAILURE_DETAIL: All nine failing checks run report subcommands through the CLI. `internal/cli/cli.go` returns exit 1 with "not implemented yet" because no report aggregation exists. Those reports are specified in the Block 5 features, so they belong to a later block. Before rerunning Block 4, build Block 5, then route the CLI to the reports and apply the two filters before aggregation. Alternatively, explicitly add the reports to Block 4 and update the Manifest.
