// Package cli parses ddkstat arguments and dispatches reports.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"time"

	"drydockstats/internal/logrecord"
	"drydockstats/internal/report"
)

const (
	exitOK         = 0
	exitUnreadable = 1
	exitUsage      = 2
)

// dateLayout is the only accepted form of --since and --until.
const dateLayout = "2006-01-02"

const usageText = "usage: ddkstat <summary|commands|targets|models|daily|failures> [--target NAME] [--since DATE] [--until DATE] LOG"

// dispatch maps each report name to the function that builds its rows.
var dispatch = map[string]func(records []logrecord.Record, skipped int) [][]string{
	"summary":  func(rs []logrecord.Record, skipped int) [][]string { return report.Summary(rs, skipped) },
	"commands": func(rs []logrecord.Record, _ int) [][]string { return report.Commands(rs) },
	"targets":  func(rs []logrecord.Record, _ int) [][]string { return report.Targets(rs) },
	"models":   func(rs []logrecord.Record, _ int) [][]string { return report.Models(rs) },
	"daily":    func(rs []logrecord.Record, _ int) [][]string { return report.Daily(rs) },
	"failures": func(rs []logrecord.Record, _ int) [][]string { return report.Failures(rs) },
}

// Run executes ddkstat with args and returns the process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		return usageError(stderr, "missing report name")
	}
	name := args[0]
	build, ok := dispatch[name]
	if !ok {
		return usageError(stderr, fmt.Sprintf("unknown report %q", name))
	}

	fs := flag.NewFlagSet("ddkstat "+name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	target := fs.String("target", "", "")
	since := fs.String("since", "", "")
	until := fs.String("until", "", "")
	if err := fs.Parse(args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(stderr, usageText)
			return exitOK
		}
		return usageError(stderr, err.Error())
	}
	switch {
	case fs.NArg() == 0:
		return usageError(stderr, "missing log path")
	case fs.NArg() > 1:
		return usageError(stderr, fmt.Sprintf("unexpected argument %q after log path", fs.Arg(1)))
	}

	f := report.Filter{Target: *target}
	set := make(map[string]bool)
	fs.Visit(func(fl *flag.Flag) { set[fl.Name] = true })
	f.HasTarget = set["target"]
	if set["since"] {
		if _, err := time.Parse(dateLayout, *since); err != nil {
			return usageError(stderr, fmt.Sprintf("invalid --since %q: want YYYY-MM-DD", *since))
		}
		f.Since = *since
	}
	if set["until"] {
		if _, err := time.Parse(dateLayout, *until); err != nil {
			return usageError(stderr, fmt.Sprintf("invalid --until %q: want YYYY-MM-DD", *until))
		}
		f.Until = *until
	}

	path := fs.Arg(0)
	res, err := logrecord.ReadFile(path)
	if err != nil {
		fmt.Fprintf(stderr, "ddkstat: cannot read log %s: %v\n", path, err)
		return exitUnreadable
	}
	for _, w := range res.Warnings {
		fmt.Fprintf(stderr, "ddkstat: warning: %s\n", w)
	}

	rows := build(report.Apply(res.Records, f), res.Skipped)
	if err := report.WriteCSV(stdout, rows); err != nil {
		fmt.Fprintf(stderr, "ddkstat: cannot write report: %v\n", err)
		return exitUnreadable
	}
	return exitOK
}

// usageError writes one diagnostic line to stderr and returns the usage exit code.
func usageError(stderr io.Writer, msg string) int {
	fmt.Fprintf(stderr, "ddkstat: %s\n", msg)
	return exitUsage
}
