<!-- Compacted from ARCHITECTURE.md sha256=50e4c308b6e72eb915ea39843b4bfe2277ccd8adc8f1a3b1c07e6688d66829cb on 2026-10-26 by drydock build agent -->

Module `ddkstat`, go 1.22+, no `require` lines, no go.sum entries, no vendor/.
Packages: cmd/ddkstat (only os.Exit caller), internal/logread, internal/filter, internal/format, internal/report, internal/cli.
logread.Read(r io.Reader, warn io.Writer) (Result, error); Result{Records []Record, Skipped int}; error only on reader failure.
logread.Record: ExecutionID, StartedAt, Date (first 10 chars), Status, Failed (status != "succeeded"), Command, Target, Provider, Model, InputTokens, CachedTokens, OutputTokens (int64), CostUSD (float64), ElapsedMS (int64), ReturnCode (*int, nil if null/absent), TimedOut (bool).
cli.Run(args []string, stdout, stderr io.Writer) int; exit 0 ok, 1 unreadable log, 2 usage. Report writers: cli.Register(name, fn func(w io.Writer, records []logread.Record, skipped int) error).
Reports: summary, commands, targets, models, daily, failures. Options: --target, --since, --until, then one log path.
Only cmd/ddkstat/main.go calls os.Exit. Output buffered; stdout written only on success. Diagnostics to stderr, never compared.
Never write to the log; never read sources/run.sh; no hardcoded report values.
