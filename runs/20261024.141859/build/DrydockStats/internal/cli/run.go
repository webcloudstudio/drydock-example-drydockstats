package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"drydockstats/internal/logintake"
	"drydockstats/internal/reports"
)

var reportNames = map[string]bool{"summary": true, "commands": true, "targets": true, "models": true, "daily": true, "failures": true}

func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || !reportNames[args[0]] {
		fmt.Fprintln(stderr, "usage: ddkstat <report> [--target NAME] [--since DATE] [--until DATE] LOG")
		return 2
	}
	fs := flag.NewFlagSet("ddkstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	target := fs.String("target", "", "")
	since := fs.String("since", "", "")
	until := fs.String("until", "", "")
	if err := fs.Parse(args[1:]); err != nil {
		fmt.Fprintln(stderr, "invalid arguments")
		return 2
	}
	rest := fs.Args()
	if len(rest) != 1 {
		fmt.Fprintln(stderr, "expected exactly one log path")
		return 2
	}
	for _, d := range []string{*since, *until} {
		if d != "" {
			x, err := time.Parse("2006-01-02", d)
			if err != nil || x.Format("2006-01-02") != d {
				fmt.Fprintln(stderr, "invalid date")
				return 2
			}
		}
	}
	file, err := os.Open(rest[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	defer file.Close()
	records, skipped, err := logintake.Read(file, func(n int, e error) { fmt.Fprintf(stderr, "warning: line %d: %v\n", n, e) })
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	filtered, err := logintake.Apply(records, logintake.Filters{Target: *target, Since: *since, Until: *until})
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 2
	}
	if err := reports.Write(stdout, args[0], filtered, skipped); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}
