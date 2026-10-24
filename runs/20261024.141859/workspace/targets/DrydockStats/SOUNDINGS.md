# Soundings

Per-assertion acceptance board, one row per Blueprint Programmatic Acceptance check.
`drydock plan` projects every assertion as `— UNVERIFIED`; `drydock score ac` sets the
Status column from deterministic proof results. A rerun of `drydock plan` resets Status to
`— UNVERIFIED`; rescore to refresh.

| Status | Blueprint | AC Id | Text | Binding | Baseline | Previously Passed | Last Block | Evidence | Verified At |
|---|---|---|---|---|---|---|---|---|---|
| ~ PREPASSED | FEATURE-LogIntake.md | intake-duplicate-skip | The parser skips a repeated execution ID, retains the first record, and reports one skipped line. | yes | PREBUILD_PASSES empty-tree | yes | block-2 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-24T14:28:06+00:00 |
| ✓ PASS | FEATURE-LogIntake.md | intake-invalid-skip | The command continues successfully with skipped malformed and incomplete lines. | yes | VALID product-missing | yes | block-2 |  | 2026-10-24T14:28:06+00:00 |
| ✓ PASS | FEATURE-LogIntake.md | intake-filtering | Target and inclusive date filters select matching records while skipped counts remain whole-file. | yes | VALID product-missing | yes | block-2 |  | 2026-10-24T14:28:06+00:00 |
| ~ PREPASSED | FEATURE-Reports.md | reports-conformance | The six reports, filters, skipped-line handling, formatting, ordering, and empty-result output pass the supplied known-answer suite. | yes | VALID existing-behavior | yes | block-2 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-24T14:28:06+00:00 |
| ~ PREPASSED | FEATURE-Reports.md | model-selection | The report aggregation assigns a record to its executed model when present and its requested model otherwise. | yes | VALID existing-behavior | yes | block-2 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-24T14:28:06+00:00 |
| ✓ PASS | FEATURE-CLI.md | cli-error-contract | Missing arguments and an unreadable log produce their specified exit statuses with no standard output. | yes | VALID baseline-red | yes | block-2 |  | 2026-10-24T14:28:06+00:00 |
| ✓ PASS | FEATURE-CLI.md | test-runner | The project test runner exits successfully after vet and all Go tests pass. | yes | VALID baseline-red | yes | block-2 |  | 2026-10-24T14:28:06+00:00 |
| ✓ PASS | FEATURE-CLI.md | full-conformance | The terminal story passes all 20 supplied grader checks. | yes | VALID baseline-red | yes | block-2 |  | 2026-10-24T14:28:06+00:00 |
