<!-- Compacted from ARCHITECTURE.md sha256=6dfac808a9c3ea79454bc3271ac63000f1596fadf3af180b25ec2760fb9c0ac7 on 2026-10-24 by drydock build agent -->

Module: `ddkstat`, `go 1.22`, no `require` lines (standard library only).
Layout: `cmd/ddkstat/main.go` (process startup only; the sole `os.Exit`); `internal/cli` (args, usage errors, dispatch); `internal/log` (JSON Lines read, parse, skip rules, model attribution, numeric defaults); `internal/filter` (target and inclusive date-range predicates); `internal/report` (aggregation and six CSV reports); `internal/format` (minutes in tenths, cost `%.4f`, CSV writer); `bin/` (launchers, holds `test.sh`).
Interface: `cli.Run(args []string, stdout, stderr io.Writer) int` — 0 success, 1 operational read failure, 2 usage error.
Boundaries: only `cmd/ddkstat` calls `os.Exit`; no package outside `cmd/ddkstat` writes to the real stdout/stderr; stdout carries only CSV; diagnostics go to stderr.
Acceptance: `go build -o <out> ./cmd/ddkstat` from the root must succeed; `go.mod` must contain `module ddkstat` and `go 1.22` with no `require`.
Guardrails: no third-party modules; never read `sources/run.sh` from application code; never hardcode report values; never write to the log being read.
