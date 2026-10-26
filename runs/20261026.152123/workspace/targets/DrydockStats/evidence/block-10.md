# Evidence: Block 10 · Feature (block-10)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 4101
- execution id: 20261026.155809.167Z-3d56068f

## Stories built
- Run the supplied grader against the finished tool as the terminal gate. (verification) [story]

## Acceptance tooling authorization
- VERIFICATION.md#verification-grader: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: VERIFICATION.md (SP 796)

## Pre-build acceptance observation
- GREEN (prepassed): verification-grader (VERIFICATION.md)
  intent: The supplied grader sources/run.sh runs unfiltered against the finished application and exits zero.
  return code: 0
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.kqpRcnQkSV/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 20 passed, 0 failed (of 20)

## Post-build programmatic acceptance
- PASS: verification-grader (VERIFICATION.md)
  intent: The supplied grader sources/run.sh runs unfiltered against the finished application and exits zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.a57NBNoJ0p/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261026.152123/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 20 passed, 0 failed (of 20)

## Build summary
The application already exists, so this step is verification. I'll run the grader unmodified.The grader passed. I'll confirm the run left no changes behind in the working tree.RESULT: SUCCESS

The supplied grader passes unmodified. `sh sources/run.sh` built `./cmd/ddkstat` and ran all 20 checks against `sources/llm_sample.jsonl`: 20 passed, 0 failed, exit 0. I did not change any file, and the working tree is still clean.

FILES CHANGED:
- None

SUMMARY:
This was a verification-only block. The existing implementation already satisfies it, so I made no application changes. Checks performed:
- `sh sources/run.sh`, run unfiltered from the application root: exit 0, 20 of 20 checks passed. The checks cover all six reports (summary, commands, targets, models, daily, failures), the Target and date filters, the empty-filter case, the missing-log operational failure (exit 1), and the usage errors (exit 2).
- `git status` after the run: no changes.

I did not run `sh bin/test.sh` or `go test`. The verification story's declared gate is the grader alone, and this block's scope was to run it.

BLOCKERS:
- None
