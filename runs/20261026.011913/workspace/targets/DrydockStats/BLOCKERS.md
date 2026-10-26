# Blockers: Uat Stage
- target: DrydockStats
- source: drydock uat DrydockStats
- declared: 2026-10-25T21:50:50-04:00

<!-- drydock-state: {"owner": "uat DrydockStats", "execution": "", "attempts": 1} -->

## blocker-001: initial-build-1
- kind: environmental
- class: uat stage
- auto-continue: no

### Error

drydock build DrydockStats exited 1.

stderr:
error: build failed for DrydockStats: 2 step(s)
  filter-target — closed/failed: programmatic acceptance failed: target-summary-counts, target-exact-match, target-no-match, target-every-report, date-since-until-inclusive, date-one-sided-bounds, date-time-of-day-ignored, date-every-report, date-with-target-and-empty-range

Evidence:
  stderr: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.011913/evidence/commands/12-initial-build-1.stderr.log
  stdout: /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.011913/evidence/commands/12-initial-build-1.stdout.log

### Why blocked

The build did not complete. UAT stopped before acceptance tests and scoring because the application is unfinished. Resolve the build error above before rerunning UAT.
