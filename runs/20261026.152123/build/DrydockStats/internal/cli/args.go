// Package cli parses and validates ddkstat invocations.
//
// ParseArgs performs no I/O and never exits the process.
package cli

import (
	"fmt"
	"strings"
	"time"
)

// Options is a validated ddkstat invocation.
type Options struct {
	Report  string
	Target  string
	Since   string
	Until   string
	LogPath string
}

// UsageError reports an invocation that does not match the ddkstat command line.
type UsageError struct {
	Msg string
}

func (e *UsageError) Error() string { return e.Msg }

func usageErrorf(format string, args ...any) error {
	return &UsageError{Msg: fmt.Sprintf(format, args...)}
}

var knownReports = map[string]bool{
	"summary":  true,
	"commands": true,
	"targets":  true,
	"models":   true,
	"daily":    true,
	"failures": true,
}

// checkDate requires value to be a real calendar date written YYYY-MM-DD. Parsing alone is not
// enough: time.Parse accepts shapes such as 2026-1-5, so the value must also round-trip exactly.
func checkDate(name, value string) error {
	parsed, err := time.Parse(dateLayout, value)
	if err != nil || parsed.Format(dateLayout) != value {
		return usageErrorf("option %s needs a YYYY-MM-DD calendar date, got %q", name, value)
	}
	return nil
}

const dateLayout = "2006-01-02"

// ParseArgs parses the arguments that follow the program name into Options. The order is fixed:
// the report name, then any --target, --since or --until options, then exactly one log path.
// Every failure is a *UsageError.
func ParseArgs(args []string) (Options, error) {
	if len(args) == 0 {
		return Options{}, usageErrorf("missing report name")
	}
	if !knownReports[args[0]] {
		return Options{}, usageErrorf("unknown report %q", args[0])
	}
	opts := Options{Report: args[0]}

	i := 1
	for ; i < len(args) && strings.HasPrefix(args[i], "--"); i += 2 {
		name := args[i]
		switch name {
		case "--target", "--since", "--until":
		default:
			return Options{}, usageErrorf("unknown option %q", name)
		}
		if i+1 >= len(args) {
			return Options{}, usageErrorf("option %s needs a value", name)
		}
		value := args[i+1]
		switch name {
		case "--target":
			opts.Target = value
		case "--since":
			if err := checkDate(name, value); err != nil {
				return Options{}, err
			}
			opts.Since = value
		case "--until":
			if err := checkDate(name, value); err != nil {
				return Options{}, err
			}
			opts.Until = value
		}
	}

	positional := args[i:]
	switch len(positional) {
	case 0:
		return Options{}, usageErrorf("missing log path")
	case 1:
		opts.LogPath = positional[0]
		return opts, nil
	default:
		return Options{}, usageErrorf("unexpected argument %q", positional[1])
	}
}
