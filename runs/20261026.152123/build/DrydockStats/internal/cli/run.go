package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"ddkstat/internal/filter"
	"ddkstat/internal/logread"
)

// ReportWriter writes one report for records already narrowed by the filters. skipped is the
// whole-file count of lines that logread did not keep.
type ReportWriter func(w io.Writer, records []logread.Record, skipped int) error

// writers maps a report name to the writer that a report story registered for it.
var writers = map[string]ReportWriter{}

// Register makes fn the writer for the report called name. Each report story calls it once from
// its own package, so Run never imports a report package.
func Register(name string, fn ReportWriter) {
	writers[name] = fn
}

// Run executes one ddkstat invocation. args are the arguments that follow the program name. It
// returns the process exit code: 0 on success, 1 when the log cannot be read or a report writer
// fails, and 2 on a usage error. Run never calls os.Exit; only cmd/ddkstat does.
//
// Report output is buffered and written to stdout only when the writer succeeds. Diagnostics go to
// stderr, and skipped log lines never change the exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	opts, err := ParseArgs(args)
	if err != nil {
		fmt.Fprintf(stderr, "usage: ddkstat {summary|commands|targets|models|daily|failures} [--target NAME] [--since DATE] [--until DATE] LOG\n")
		return 2
	}

	f, err := os.Open(opts.LogPath)
	if err != nil {
		fmt.Fprintf(stderr, "error: cannot open log %s: %v\n", opts.LogPath, err)
		return 1
	}
	defer f.Close()

	res, err := logread.Read(f, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "error: cannot read log %s: %v\n", opts.LogPath, err)
		return 1
	}

	records := filter.ByDateRange(filter.ByTarget(res.Records, opts.Target), opts.Since, opts.Until)

	write, ok := writers[opts.Report]
	if !ok {
		return 0
	}

	var buf bytes.Buffer
	if err := write(&buf, records, res.Skipped); err != nil {
		fmt.Fprintf(stderr, "error: cannot write %s report: %v\n", opts.Report, err)
		return 1
	}
	if _, err := buf.WriteTo(stdout); err != nil {
		fmt.Fprintf(stderr, "error: cannot write report to stdout: %v\n", err)
		return 1
	}
	return 0
}
