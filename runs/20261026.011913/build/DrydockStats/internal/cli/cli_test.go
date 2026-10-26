package cli

import (
	"bytes"
	"testing"
)

func TestRunUsageErrorsExitTwoWithEmptyStdout(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "no arguments", args: nil},
		{name: "report only", args: []string{"summary"}},
		{name: "too many positionals", args: []string{"summary", "log.jsonl", "extra"}},
		{name: "unknown option", args: []string{"--bogus", "summary", "log.jsonl"}},
		{name: "help is an unknown option", args: []string{"-h"}},
		{name: "option missing its value", args: []string{"summary", "log.jsonl", "--target"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			got := Run(tt.args, &stdout, &stderr)
			if got != 2 {
				t.Errorf("Run(%q) = %d, want 2", tt.args, got)
			}
			if stdout.Len() != 0 {
				t.Errorf("Run(%q) wrote %q to stdout, want nothing", tt.args, stdout.String())
			}
			if stderr.Len() == 0 {
				t.Errorf("Run(%q) wrote nothing to stderr, want the usage line", tt.args)
			}
		})
	}
}

func TestParseAcceptsEveryReport(t *testing.T) {
	for _, report := range []string{"summary", "commands", "targets", "models", "daily", "failures"} {
		t.Run(report, func(t *testing.T) {
			got, err := Parse([]string{report, "llm.jsonl"})
			if err != nil {
				t.Fatalf("Parse(%s, llm.jsonl) error: %v", report, err)
			}
			if got.Report != report {
				t.Errorf("Report = %q, want %q", got.Report, report)
			}
			if got.Log != "llm.jsonl" {
				t.Errorf("Log = %q, want llm.jsonl", got.Log)
			}
		})
	}
}

func TestParseOptionsFollowReport(t *testing.T) {
	want := Options{
		Report:    "summary",
		Target:    "CommonMark",
		HasTarget: true,
		Since:     "2026-10-01",
		Until:     "2026-10-31",
		Log:       "llm.jsonl",
	}
	orders := [][]string{
		{"summary", "--target", "CommonMark", "--since", "2026-10-01", "--until", "2026-10-31", "llm.jsonl"},
		{"summary", "--until", "2026-10-31", "--since", "2026-10-01", "--target", "CommonMark", "llm.jsonl"},
		{"summary", "--since", "2026-10-01", "--target", "CommonMark", "--until", "2026-10-31", "llm.jsonl"},
	}
	for _, args := range orders {
		got, err := Parse(args)
		if err != nil {
			t.Fatalf("Parse(%q) error: %v", args, err)
		}
		if got != want {
			t.Errorf("Parse(%q) = %+v, want %+v", args, got, want)
		}
	}
}

func TestParseLogPathIsLast(t *testing.T) {
	got, err := Parse([]string{"daily", "--until", "2026-10-16", "/data/logs/llm.jsonl"})
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if got.Log != "/data/logs/llm.jsonl" {
		t.Errorf("Log = %q, want /data/logs/llm.jsonl", got.Log)
	}
	if got.Report != "daily" {
		t.Errorf("Report = %q, want daily", got.Report)
	}
}

func TestParseRepeatedOptionLastWins(t *testing.T) {
	got, err := Parse([]string{"models", "--target", "Toml", "--target", "ReadingList", "llm.jsonl"})
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if got.Target != "ReadingList" {
		t.Errorf("Target = %q, want ReadingList", got.Target)
	}
	if !got.HasTarget {
		t.Errorf("HasTarget = false, want true")
	}
}

func TestParseHasTargetFalseWhenAbsent(t *testing.T) {
	got, err := Parse([]string{"summary", "llm.jsonl"})
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if got.HasTarget {
		t.Errorf("HasTarget = true, want false")
	}
}

func TestParseRejectsUnknownReport(t *testing.T) {
	if _, err := Parse([]string{"frobnicate", "llm.jsonl"}); err == nil {
		t.Error("Parse with unknown report returned no error")
	}
}

func TestParseRejectsWrongPositionalCount(t *testing.T) {
	for _, args := range [][]string{
		{"summary"},
		{"summary", "llm.jsonl", "extra"},
		{"llm.jsonl"},
	} {
		if _, err := Parse(args); err == nil {
			t.Errorf("Parse(%q) returned no error", args)
		}
	}
}

func TestRunAcceptsOptionsInAnyPosition(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "options after positionals", args: []string{"summary", "log.jsonl", "--target", "Toml"}},
		{name: "options before positionals", args: []string{"--since", "2026-10-14", "summary", "log.jsonl"}},
		{name: "options between positionals", args: []string{"summary", "--until", "2026-10-16", "log.jsonl"}},
		{name: "double dash ends options", args: []string{"summary", "--", "-log.jsonl"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			// Two positionals parse cleanly; the report itself is not yet implemented.
			if got := Run(tt.args, &stdout, &stderr); got != 1 {
				t.Errorf("Run(%q) = %d, want 1", tt.args, got)
			}
		})
	}
}
