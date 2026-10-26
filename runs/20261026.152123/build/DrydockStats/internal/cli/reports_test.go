package cli_test

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"strings"
	"testing"

	"ddkstat/internal/cli"
	"ddkstat/internal/report"
)

// runReport registers the real report writers, which earlier tests may have replaced, then runs
// report over a log written by the test. It returns the exit code and the parsed CSV rows.
func runReport(t *testing.T, args []string, lines ...string) (int, [][]string) {
	t.Helper()
	report.Register()
	var stdout, stderr bytes.Buffer
	code := cli.Run(append(args, writeLog(t, lines...)), &stdout, &stderr)
	if code != 0 {
		return code, nil
	}
	rows, err := csv.NewReader(strings.NewReader(stdout.String())).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%s", err, stdout.String())
	}
	return code, rows
}

func TestRunSummaryReport(t *testing.T) {
	code, rows := runReport(t, []string{"summary"},
		logLine("a", "2026-10-14", "Toml"),
		logLine("b", "2026-10-16", "CommonMark"),
		"{not json",
	)
	if code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	want := [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"2", "1", "2", "0", "2026-10-14", "2026-10-16", "2", "4", "6", "1.0000", "0.0"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %v, want %v", rows, want)
	}
}

func TestRunTargetsReport(t *testing.T) {
	code, rows := runReport(t, []string{"targets"},
		logLine("a", "2026-10-14", "Toml"),
		logLine("b", "2026-10-15", "CommonMark"),
	)
	if code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	want := [][]string{
		{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"CommonMark", "1", "0", "1", "2", "3", "0.5000", "0.0"},
		{"Toml", "1", "0", "1", "2", "3", "0.5000", "0.0"},
		{"TOTAL", "2", "0", "2", "4", "6", "1.0000", "0.0"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %v, want %v", rows, want)
	}
}

func TestRunCommandsReport(t *testing.T) {
	code, rows := runReport(t, []string{"commands"},
		logLine("a", "2026-10-14", "Toml"),
		logLine("b", "2026-10-15", "CommonMark"),
	)
	if code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	want := [][]string{
		{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"},
		{"build", "2", "0", "2", "4", "6", "0.0"},
		{"TOTAL", "2", "0", "2", "4", "6", "0.0"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %v, want %v", rows, want)
	}
}

func TestRunModelsReport(t *testing.T) {
	code, rows := runReport(t, []string{"models"},
		logLine("a", "2026-10-14", "Toml"),
		logLine("b", "2026-10-15", "CommonMark"),
	)
	if code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	want := [][]string{
		{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"},
		{"claude", "claude-haiku-5-5", "2", "0", "2", "4", "6", "1.0000"},
		{"TOTAL", "", "2", "0", "2", "4", "6", "1.0000"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %v, want %v", rows, want)
	}
}

func TestRunDailyReport(t *testing.T) {
	code, rows := runReport(t, []string{"daily"},
		logLine("a", "2026-10-14", "Toml"),
		logLine("b", "2026-10-15", "CommonMark"),
	)
	if code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	want := [][]string{
		{"date", "calls", "failed", "tokens", "minutes"},
		{"2026-10-14", "1", "0", "6", "0.0"},
		{"2026-10-15", "1", "0", "6", "0.0"},
		{"TOTAL", "2", "0", "12", "0.0"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %v, want %v", rows, want)
	}
}

func TestRunFailuresReport(t *testing.T) {
	failed := strings.Replace(logLine("b", "2026-10-15", "CommonMark"), `"succeeded"`, `"failed"`, 1)
	code, rows := runReport(t, []string{"failures"}, logLine("a", "2026-10-14", "Toml"), failed)
	if code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	want := [][]string{
		{"started_at", "target", "command", "model", "returncode", "timed_out"},
		{"2026-10-15T02:38:41.904Z", "CommonMark", "build", "claude-haiku-5-5", "0", "false"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %v, want %v", rows, want)
	}
}
