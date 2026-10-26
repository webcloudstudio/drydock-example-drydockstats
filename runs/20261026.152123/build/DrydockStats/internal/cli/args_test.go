package cli_test

import (
	"errors"
	"reflect"
	"testing"

	"ddkstat/internal/cli"
)

func mustParse(t *testing.T, args ...string) cli.Options {
	t.Helper()
	opts, err := cli.ParseArgs(args)
	if err != nil {
		t.Fatalf("ParseArgs(%q): %v", args, err)
	}
	return opts
}

func expectUsageError(t *testing.T, args ...string) {
	t.Helper()
	_, err := cli.ParseArgs(args)
	var usage *cli.UsageError
	if !errors.As(err, &usage) {
		t.Errorf("ParseArgs(%q) error = %v, want *UsageError", args, err)
	}
}

func TestParseReportAndLog(t *testing.T) {
	got := mustParse(t, "summary", "log.jsonl")
	want := cli.Options{Report: "summary", LogPath: "log.jsonl"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Options = %+v, want %+v", got, want)
	}
}

func TestParseEveryKnownReport(t *testing.T) {
	for _, report := range []string{"summary", "commands", "targets", "models", "daily", "failures"} {
		if got := mustParse(t, report, "log.jsonl").Report; got != report {
			t.Errorf("Report = %q, want %q", got, report)
		}
	}
}

func TestParseOptions(t *testing.T) {
	got := mustParse(t, "daily", "--target", "Toml", "--since", "2026-10-14", "--until", "2026-10-16", "log.jsonl")
	want := cli.Options{
		Report:  "daily",
		Target:  "Toml",
		Since:   "2026-10-14",
		Until:   "2026-10-16",
		LogPath: "log.jsonl",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Options = %+v, want %+v", got, want)
	}
}

func TestParseOptionsInAnyOrder(t *testing.T) {
	got := mustParse(t, "daily", "--until", "2026-10-16", "--target", "Toml", "log.jsonl")
	if got.Until != "2026-10-16" || got.Target != "Toml" {
		t.Errorf("Until, Target = %q, %q; want 2026-10-16, Toml", got.Until, got.Target)
	}
}

func TestParseRepeatedOptionLastWins(t *testing.T) {
	got := mustParse(t, "commands", "--target", "A", "--target", "B", "log.jsonl")
	if got.Target != "B" {
		t.Errorf("Target = %q, want B", got.Target)
	}
}

func TestParseOptionBeforeReportIsUsageError(t *testing.T) {
	expectUsageError(t, "--target", "Toml", "summary", "log.jsonl")
}

func TestParseUnknownReport(t *testing.T) {
	expectUsageError(t, "frobnicate", "log.jsonl")
}

func TestParseUnknownOption(t *testing.T) {
	expectUsageError(t, "summary", "--colour", "red", "log.jsonl")
}

func TestParseMissingValue(t *testing.T) {
	expectUsageError(t, "summary", "log.jsonl", "--target")
}

func TestParseMissingValueBeforeLogPath(t *testing.T) {
	expectUsageError(t, "summary", "--since")
}

func TestParseMissingLogPath(t *testing.T) {
	expectUsageError(t, "summary")
}

func TestParseExtraLogPath(t *testing.T) {
	expectUsageError(t, "summary", "log.jsonl", "other.jsonl")
}

func TestParseLogPathCount(t *testing.T) {
	if got := mustParse(t, "summary", "one.jsonl").LogPath; got != "one.jsonl" {
		t.Errorf("LogPath = %q, want one.jsonl", got)
	}
	expectUsageError(t, "summary")
	expectUsageError(t, "summary", "one.jsonl", "two.jsonl")
	expectUsageError(t, "summary", "one.jsonl", "two.jsonl", "three.jsonl")
}

func TestParseOptionAfterLogPathIsUsageError(t *testing.T) {
	expectUsageError(t, "summary", "log.jsonl", "--target", "Toml")
}

func TestParseNoArguments(t *testing.T) {
	expectUsageError(t)
}

func TestDateValid(t *testing.T) {
	got := mustParse(t, "summary", "--since", "2026-02-28", "--until", "2028-02-29", "log.jsonl")
	if got.Since != "2026-02-28" || got.Until != "2028-02-29" {
		t.Errorf("Since, Until = %q, %q; want 2026-02-28, 2028-02-29", got.Since, got.Until)
	}
}

func TestDateInvalidMonth(t *testing.T) {
	expectUsageError(t, "summary", "--since", "2026-13-01", "log.jsonl")
	expectUsageError(t, "summary", "--until", "2026-00-10", "log.jsonl")
}

func TestDateInvalidDay(t *testing.T) {
	expectUsageError(t, "summary", "--since", "2026-02-30", "log.jsonl")
	expectUsageError(t, "summary", "--until", "2027-02-29", "log.jsonl")
	expectUsageError(t, "summary", "--until", "2026-04-31", "log.jsonl")
}

func TestDateWrongShape(t *testing.T) {
	expectUsageError(t, "summary", "--since", "2026-1-5", "log.jsonl")
	expectUsageError(t, "summary", "--since", "2026/10/14", "log.jsonl")
	expectUsageError(t, "summary", "--until", "2026-10-14T00:00:00Z", "log.jsonl")
	expectUsageError(t, "summary", "--until", "", "log.jsonl")
}

func TestDateMissingValue(t *testing.T) {
	expectUsageError(t, "summary", "--since")
	expectUsageError(t, "summary", "--until")
	expectUsageError(t, "summary", "--until", "log.jsonl")
}
