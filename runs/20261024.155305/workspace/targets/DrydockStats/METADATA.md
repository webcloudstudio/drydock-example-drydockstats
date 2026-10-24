# AUTHORITATIVE PROJECT METADATA — FIELDS SHOULD BE CURRENT

name: DrydockStats
display_name: drydock DrydockStats Uat Case
short_description: Drydock uat - DrydockStats builds a Go reporter for Drydock's own LLM execution log (logs/llm.jsonl). The source is an original brief: read JSON Lines, skip malformed and duplicate records without failing, and write six CSV reports (summary, commands, targets, models, daily, failures) with Target and date filters. The supplied grader, sources/run.sh, builds ./cmd/ddkstat with go build and runs 20 checks of exit code and output against a 37-line synthetic sample with known totals. Standard library only, no database. Run with drydock uat DrydockStats. Source Repo: none, the brief is original to this kit.
stack: Go
version: 0.09
build_state: building
build_sub_state: running
last_analyzed: 2026-10-24
last_planned: 2026-10-24
last_built: 
build_dir: 
brand: 
git_repo: 
release_tag: 
updated: 2026-10-24T16:03:48+00:00
