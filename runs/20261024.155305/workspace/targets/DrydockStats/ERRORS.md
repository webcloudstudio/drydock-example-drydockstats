# BIG ERRORS — action required

- Command: `build`
- Phase: build step
- State: Failed
- Timestamp: 2026-10-24T16:15:50+00:00
- Execution ID: 20261024.161238.365Z-e58fe633
- Challenge Execution ID: -
- Classification: programmatic acceptance failed: log-required-fields-skip
- Evidence / logs: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.155305/workspace/targets/DrydockStats/evidence/block-2.md

## Diagnostic

Block "Block 2 · Service" [block-2] failed its acceptance criteria.
  Story "Skip records missing a required field, with a line-numbered warning." [log-required-fields] does not meet its own acceptance criteria:
    - AC log-required-fields-skip — A record missing job.target is counted as skipped and does not appear in the records.
        assertion: IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'
        raised at: pathlib.py:1013
        process exit code: 1
        error: IsADirectoryError: [Errno 21] Is a directory: 'internal/logrecord'
        values at failure:
          self = PosixPath('internal/logrecord')
          mode = 'r'
          buffering = -1
          encoding = 'locale'
          errors = None
          newline = None
        check stderr:
… (11 more lines truncated)

## Recovery

The agent disputes failing acceptance criteria; the claim is unverified.
Run: drydock plan verify DrydockStats
  Inspect the criteria and evidence before rebuilding.
Run: drydock plan repair DrydockStats
  Repair confirmed mechanical defects, then verify before rebuilding.
