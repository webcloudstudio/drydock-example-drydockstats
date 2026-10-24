// Package cli parses ddkstat arguments and dispatches to the reports.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"

	"ddkstat/internal/filter"
	"ddkstat/internal/format"
	ddklog "ddkstat/internal/log"
	"ddkstat/internal/report"
)

// Exit codes returned by Run.
const (
	exitOK       = 0
	exitReadFail = 1
	exitUsage    = 2
)

// reports maps each report name to its builder. skipped is the whole-file skipped count.
var reports = map[string]func(recs []ddklog.Record, skipped int) [][]string{
	"summary":  func(recs []ddklog.Record, skipped int) [][]string { return report.Summary(recs, skipped) },
	"commands": func(recs []ddklog.Record, _ int) [][]string { return report.Commands(recs) },
	"targets":  func(recs []ddklog.Record, _ int) [][]string { return report.Targets(recs) },
	"models":   func(recs []ddklog.Record, _ int) [][]string { return report.Models(recs) },
	"daily":    func(recs []ddklog.Record, _ int) [][]string { return report.Daily(recs) },
	"failures": func(recs []ddklog.Record, _ int) [][]string { return report.Failures(recs) },
}

// once is a string flag that may be given at most once.
type once struct {
	value string
	set   bool
}

func (o *once) String() string { return o.value }

func (o *once) Set(s string) error {
	if o.set {
		return errors.New("option given more than once")
	}
	o.value, o.set = s, true
	return nil
}

// Run parses args, writes the report to stdout, writes diagnostics to stderr, and returns the exit code.
// The first argument is the report name; options follow it and the single log path may appear anywhere
// after the report name.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		return usageError(stderr, "missing report name")
	}
	name := args[0]
	build, ok := reports[name]
	if !ok {
		return usageError(stderr, fmt.Sprintf("unknown report %q", name))
	}
	opts, logPath, err := parseArgs(name, args[1:])
	if err != nil {
		return usageError(stderr, err.Error())
	}

	lines, err := ddklog.ReadFile(logPath)
	if err != nil {
		fmt.Fprintf(stderr, "error: cannot read log: %v\n", err)
		return exitReadFail
	}
	parsed := ddklog.Parse(lines)
	for _, s := range parsed.Skips {
		fmt.Fprintf(stderr, "warning: skipped line %d: %s\n", s.Line, s.Reason)
	}

	rows := build(filter.Apply(opts, parsed.Records), len(parsed.Skips))
	if err := format.WriteCSV(stdout, rows); err != nil {
		fmt.Fprintf(stderr, "error: cannot write report: %v\n", err)
		return exitReadFail
	}
	return exitOK
}

// parseArgs parses the options of one report and its single log path. flag stops at the first
// positional argument, so parsing resumes after each one to accept options placed after the log path.
func parseArgs(name string, args []string) (filter.Options, string, error) {
	var target, since, until once
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Var(&target, "target", "keep only records for this Target")
	fs.Var(&since, "since", "keep records starting on or after this date")
	fs.Var(&until, "until", "keep records starting on or before this date")

	var positional []string
	rest := args
	for {
		if err := fs.Parse(rest); err != nil {
			return filter.Options{}, "", err
		}
		rest = fs.Args()
		if len(rest) == 0 {
			break
		}
		positional = append(positional, rest[0])
		rest = rest[1:]
	}
	if len(positional) != 1 {
		return filter.Options{}, "", fmt.Errorf("expected exactly one log path, got %d", len(positional))
	}
	for _, d := range []once{since, until} {
		if d.set && !filter.ValidDate(d.value) {
			return filter.Options{}, "", fmt.Errorf("invalid date %q: want a real calendar date as YYYY-MM-DD", d.value)
		}
	}

	opts := filter.Options{
		Target:    target.value,
		HasTarget: target.set,
		Since:     since.value,
		Until:     until.value,
	}
	return opts, positional[0], nil
}

// usageError writes one diagnostic line to stderr and returns the usage exit code.
func usageError(stderr io.Writer, msg string) int {
	fmt.Fprintf(stderr, "error: %s\n", msg)
	return exitUsage
}
