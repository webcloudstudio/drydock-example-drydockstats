# MANIFEST: DrydockStats
updated:     2026-10-24T14:20:46+00:00
state:       approved
blocks:      2

applied_specs: |
  FEATURE-CLI.md sha256=17994b44ae4b152599da77ede0c186b55f48d183af09f99bda28ea795c64cebb commit=b35940d9b8c2b727dfcd27d4dfe915eabbc4acba applied_by=block-2 applied_at=2026-10-24T14:27:36+00:00 build_sha256=047a487b0e794f7a947c461efad323db3c3c7c929ace834c4bdae0fdc51e4136
  FEATURE-LogIntake.md sha256=edf21aed845096c9abd10c32bdbd6b1f368b16ec5d0d9642f810bccacbc48010 commit=b35940d9b8c2b727dfcd27d4dfe915eabbc4acba applied_by=block-1 applied_at=2026-10-24T14:24:36+00:00 build_sha256=6bdc4b078baacdeb32e2cfd1b783ad1bf526e203a089932644856389e6aabd32
  FEATURE-Reports.md sha256=c8777526bd9d45cea06d603aae0f8605522a784d5c62a14681961031676872ee commit=b35940d9b8c2b727dfcd27d4dfe915eabbc4acba applied_by=block-2 applied_at=2026-10-24T14:27:36+00:00 build_sha256=6c6898649ab216e77ecd4cdf712daf2fd1a75d0aabb7a23defb82ae2e4aeaa5e
  llm_sample.jsonl sha256=19a092d921393ca299df3dbdb1b65370649bda5aaea44c2fe4f7f1080cc5f217 commit=b35940d9b8c2b727dfcd27d4dfe915eabbc4acba applied_by=block-2 applied_at=2026-10-24T14:27:36+00:00 build_sha256=19a092d921393ca299df3dbdb1b65370649bda5aaea44c2fe4f7f1080cc5f217
  run.sh sha256=0cc2ee483269886e4d61e5bfdbefec2419121f8a49a10ef33f182f855ba889d3 commit=b35940d9b8c2b727dfcd27d4dfe915eabbc4acba applied_by=block-2 applied_at=2026-10-24T14:27:36+00:00 build_sha256=0cc2ee483269886e4d61e5bfdbefec2419121f8a49a10ef33f182f855ba889d3
## story 1: Parse execution records, skip invalid and duplicate lines, and apply filters.
id: log-intake
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-1.md
summary: Parse execution records, skip invalid and duplicate lines, and apply filters.
type: service
kind: capability
phase: 1
block: 1
implements: FEATURE-LogIntake.md
covers: INTAKE-001, INTAKE-002
accepts: st-004, st-006
context: llm_sample.jsonl
stack: go.md, common.md
stack_mode: builder
size: 12367
provides: log parsing, record validation, duplicate detection, target and date filtering
instructions: Implement JSON Lines intake, validation, skipped-line counting and warnings, exact Target filtering, and inclusive date filtering. Keep skipped counts independent of report filters.
acceptance: yes
state: closed/implemented

## story 2: Aggregate and format the six CSV reports.
id: report-generation
finding: ADVISORY: implemented, not verified — every binding criterion this story passed was already green before the story built, so none measured the work. Repair the criteria in the Blueprint to exercise this story and rebuild to obtain a verdict.
evidence: evidence/block-2.md
summary: Aggregate and format the six CSV reports.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Reports.md
covers: REPORT-001, REPORT-002, REPORT-003, REPORT-004, REPORT-005
accepts: st-005, st-007
context: llm_sample.jsonl
stack: go.md, common.md
stack_mode: consumer
size: 7819
provides: summary, commands, targets, models, daily, failures CSV reports
consumes: log parsing, record validation, duplicate detection, target and date filtering
instructions: Implement the six CSV reports from filtered records, with the specified totals, ordering, formatting, model selection, and empty-result behavior. Derive every value from the supplied log.
acceptance: yes
depends: log-intake
state: closed/implemented

## story 3: Deliver and verify the command-line interface, test runner, and full conformance gate.
id: cli-delivery
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-2.md
summary: Deliver and verify the command-line interface, test runner, and full conformance gate.
type: service
kind: test harness
phase: 1
block: 2
implements: FEATURE-CLI.md
covers: CLI-001, CLI-002
accepts: st-001, st-002, st-003, st-008
context: INSTRUCTIONS.md, run.sh, llm_sample.jsonl
stack: go.md, common.md
stack_mode: consumer
size: 11392
provides: ./cmd/ddkstat, bin/test.sh
consumes: log parsing, record validation, duplicate detection, target and date filtering, summary, commands, targets, models, daily, failures CSV reports
instructions: |
  Implement argument handling, report dispatch, standard error diagnostics, conventional exit codes, and the Go test runner. Add focused automated tests for parsing, filtering, reports, and CLI behavior.
  
  As the terminal story, run the complete supplied grader and ensure it passes.
acceptance: yes
depends: log-intake, report-generation
state: closed/implemented
