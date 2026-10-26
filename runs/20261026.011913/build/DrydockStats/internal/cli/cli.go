// Package cli parses ddkstat arguments and orchestrates a single run.
package cli

import (
	"flag"
	"fmt"
	"io"
	"time"

	"ddkstat/internal/logread"
)

const usage = "usage: ddkstat <summary|commands|targets|models|daily|failures> <log> [--target NAME] [--since YYYY-MM-DD] [--until YYYY-MM-DD]"

// Exit codes, as declared in the architecture.
const (
	exitRead  = 1
	exitUsage = 2
)

// reportNames lists the six reports in the order the architecture declares them.
var reportNames = []string{"summary", "commands", "targets", "models", "daily", "failures"}

// Options is the parsed ddkstat command line.
type Options struct {
	Report    string
	Target    string
	HasTarget bool
	Since     string // YYYY-MM-DD or empty
	Until     string // YYYY-MM-DD or empty
	Log       string
}

// Parse reads the command line (program name excluded) into Options. The report comes first,
// the three options may follow in any order, and exactly one log path comes last. A given
// --since or --until must be a real YYYY-MM-DD date, or Parse returns an error.
func Parse(args []string) (Options, error) {
	fs := flag.NewFlagSet("ddkstat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	target := fs.String("target", "", "")
	since := fs.String("since", "", "")
	until := fs.String("until", "", "")
	positional, err := parseInterspersed(fs, args)
	if err != nil {
		return Options{}, err
	}
	if len(positional) != 2 {
		return Options{}, fmt.Errorf("want one report and one log path, got %d positional arguments", len(positional))
	}
	if !isReport(positional[0]) {
		return Options{}, fmt.Errorf("unknown report %q", positional[0])
	}
	set := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) {
		set[f.Name] = true
	})
	opts := Options{
		Report:    positional[0],
		Target:    *target,
		HasTarget: set["target"],
		Log:       positional[1],
	}
	if set["since"] {
		if opts.Since, err = validateDate(*since); err != nil {
			return Options{}, fmt.Errorf("--since: %w", err)
		}
	}
	if set["until"] {
		if opts.Until, err = validateDate(*until); err != nil {
			return Options{}, fmt.Errorf("--until: %w", err)
		}
	}
	return opts, nil
}

// dateLayout is the only accepted form of --since and --until: a real calendar date, YYYY-MM-DD.
const dateLayout = "2006-01-02"

// validateDate returns value unchanged when it is exactly ten characters and a real calendar
// date in the YYYY-MM-DD layout, so leap days are honored. It never normalizes a value.
func validateDate(value string) (string, error) {
	if len(value) != len(dateLayout) {
		return "", fmt.Errorf("%q is not a YYYY-MM-DD date", value)
	}
	if _, err := time.Parse(dateLayout, value); err != nil {
		return "", fmt.Errorf("%q is not a real calendar date", value)
	}
	return value, nil
}

// isReport reports whether name is one of the six report names.
func isReport(name string) bool {
	for _, r := range reportNames {
		if r == name {
			return true
		}
	}
	return false
}

// Run executes ddkstat with args (the program name excluded) and returns the process exit code.
// Usage errors write the usage line to stderr and return 2; nothing reaches stdout.
func Run(args []string, stdout, stderr io.Writer) int {
	opts, err := Parse(args)
	if err != nil {
		fmt.Fprintln(stderr, usage)
		return exitUsage
	}
	return execute(opts, stdout, stderr, func(warn io.Writer) (logread.Result, error) {
		return logread.ReadFile(opts.Log, warn)
	})
}

// execute reads the log through load and returns the exit code. A log that cannot be read exits
// 1 with one diagnostic on stderr. Report output will go to stdout only after the whole log has
// been read, so a failing read never leaves partial output.
func execute(opts Options, stdout, stderr io.Writer, load func(warn io.Writer) (logread.Result, error)) int {
	if _, err := load(stderr); err != nil {
		fmt.Fprintf(stderr, "ddkstat: cannot read log: %v\n", err)
		return exitRead
	}
	fmt.Fprintf(stderr, "ddkstat: report %q is not implemented yet\n", opts.Report)
	return exitRead
}

// parseInterspersed parses flags that may appear before, between, or after positional
// arguments, as the flag package alone stops at the first positional. A bare "--" ends
// option parsing, and everything after it is positional.
func parseInterspersed(fs *flag.FlagSet, args []string) ([]string, error) {
	var positional []string
	for {
		if err := fs.Parse(args); err != nil {
			return nil, err
		}
		rest := fs.Args()
		if consumed := len(args) - len(rest); consumed > 0 && args[consumed-1] == "--" {
			return append(positional, rest...), nil
		}
		if len(rest) == 0 {
			return positional, nil
		}
		positional = append(positional, rest[0])
		args = rest[1:]
	}
}
