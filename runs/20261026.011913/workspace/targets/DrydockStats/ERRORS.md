# BIG ERRORS — action required

- Command: `build`
- Phase: build step
- State: Failed
- Timestamp: 2026-10-26T01:50:44+00:00
- Execution ID: 20261026.015000.461Z-1930d24a
- Challenge Execution ID: -
- Classification: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report, date-since-until-inclusive, date-one-sided-bounds, date-time-of-day-ignored, date-every-report, date-with-target-and-empty-ran…
- Evidence / logs: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.011913/workspace/targets/DrydockStats/evidence/block-4.md

## Diagnostic

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
… (125 more lines truncated)

## Recovery

Run: drydock build DrydockStats
  to continue the build. This story resumes where it left off and is retried against the checks it failed.
