# COMPASS: DrydockStats

## Compass

Build a command-line reporter for Drydock's LLM execution log, the Go executable `./cmd/ddkstat`,
as described in `sources/INSTRUCTIONS.md`. It reads a JSON Lines log, skips malformed and
duplicate lines without failing, and writes six CSV reports (summary, commands, targets, models,
daily, failures), each narrowed by optional Target and date filters.

Correctness is measured by `sources/run.sh`: it builds the executable with `go build` and makes
20 checks of exit code and standard output against `sources/llm_sample.jsonl`. The goal is every
check passing. The finished tool also runs against a real `logs/llm.jsonl`.

## Constraints

- Implement in Go 1.22 or newer with the standard library only. No third-party modules.
- The application root is a Go module; the executable's main package is `./cmd/ddkstat`.
- Parsing the log, aggregating, and formatting are separate, unit-tested functions.
- Diagnostics go to standard error and are never compared.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Do not read `sources/run.sh` from application code, and do not hardcode a report row, a count,
  or a total. Every value comes from the log being read.
- Never write to the log being read.

## Verification Protocol

| Story kind | Command | Asserts |
|---|---|---|
| Every behavioral story | `sh bin/test.sh` | exit `0` |
| Terminal story (once, last) | `sh sources/run.sh` | exit `0` |

<!-- drydock:build-write-guardrail:start -->
## Build Write Guardrail

- Authorized build directory: `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/build/DrydockStats`
- Authorized Target directory: `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats`
- Build agents have permission to create, modify, and remove files required by the active build block inside these authorized directories.
- No path outside these authorized directories may be modified.
- Protected Drydock artifacts:
  - `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats/blueprint/`
  - `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats/MANIFEST.md`
  - `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats/COMPASS.md`
  - `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats/QuarterDeck/`
  - `/mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.145357/workspace/targets/DrydockStats/evidence/`
<!-- drydock:build-write-guardrail:end -->
