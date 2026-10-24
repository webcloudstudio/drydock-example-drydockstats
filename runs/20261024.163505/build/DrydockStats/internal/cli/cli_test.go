package cli

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const (
	recAlpha = `{"schema_version":1,"execution_id":"a-1","started_at":"2026-10-14T06:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Alpha","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"cached_input_tokens":10,"elapsed_ms":60000,"input_tokens":5,"output_tokens":7,"cost_usd":0.25,"model":"haiku"}}}`
	recBeta  = `{"schema_version":1,"execution_id":"a-2","started_at":"2026-10-16T06:00:00.000Z","status":"failed","job":{"command_name":"plan","target":"Beta","llm":"claude","model":"sonnet"},"result":{"returncode":null,"timed_out":true,"stats":{"cached_input_tokens":20,"elapsed_ms":35000,"input_tokens":6,"output_tokens":8,"cost_usd":0.5}}}`
	recBad   = `not json`
)

// writeLog writes lines to a fresh temp log and returns its path.
func writeLog(t *testing.T, lines ...string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "log.jsonl")
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

// runReport runs ddkstat and returns the exit code and the CSV rows read back from stdout.
func runReport(t *testing.T, args ...string) (int, [][]string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	code := Run(args, &stdout, &stderr)
	if code != exitOK {
		return code, nil
	}
	rows, err := csv.NewReader(&stdout).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%s", err, stdout.String())
	}
	return code, rows
}

func TestRunUsageErrorsExitTwoWithEmptyStdout(t *testing.T) {
	log := writeLog(t)
	tests := []struct {
		name string
		args []string
	}{
		{"no arguments", nil},
		{"unknown report", []string{"frobnicate", log}},
		{"unknown flag", []string{"summary", "-nope", log}},
		{"help flag is unknown", []string{"summary", "--help", log}},
		{"missing log path", []string{"summary"}},
		{"extra log path", []string{"summary", log, log}},
		{"missing option value", []string{"summary", log, "--target"}},
		{"repeated option", []string{"summary", "--target", "A", "--target", "B", log}},
		{"impossible --since date", []string{"summary", "--since", "2026-13-01", log}},
		{"impossible --until date", []string{"summary", log, "--until", "2026-02-30"}},
		{"non-ISO date", []string{"summary", "--since", "2026/10/14", log}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if got := Run(tt.args, &stdout, &stderr); got != exitUsage {
				t.Errorf("Run(%q) = %d, want %d", tt.args, got, exitUsage)
			}
			if stdout.Len() != 0 {
				t.Errorf("Run(%q) wrote %q to stdout, want nothing", tt.args, stdout.String())
			}
		})
	}
}

func TestRunUnreadableLogExitsOne(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-log.jsonl")
	var stdout, stderr bytes.Buffer
	if got := Run([]string{"summary", missing}, &stdout, &stderr); got != exitReadFail {
		t.Errorf("Run = %d, want %d", got, exitReadFail)
	}
	if stdout.Len() != 0 {
		t.Errorf("wrote %q to stdout, want nothing", stdout.String())
	}
}

func TestSummaryReportRoundTrip(t *testing.T) {
	log := writeLog(t, recAlpha, recBad, recBeta)
	code, rows := runReport(t, "summary", log)
	if code != exitOK {
		t.Fatalf("exit = %d, want 0", code)
	}
	want := [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"2", "1", "1", "1", "2026-10-14", "2026-10-16", "11", "30", "15", "0.7500", "1.6"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("summary rows = %q, want %q", rows, want)
	}
}

func TestOptionsAfterLogPathAreApplied(t *testing.T) {
	log := writeLog(t, recAlpha, recBeta)
	code, rows := runReport(t, "summary", log, "--target", "Alpha")
	if code != exitOK {
		t.Fatalf("exit = %d, want 0", code)
	}
	if got := rows[1][0]; got != "1" {
		t.Errorf("records = %q, want 1 for --target Alpha", got)
	}
}

func TestTargetFilterAppliesToEveryRecordAndSkippedStaysWholeFile(t *testing.T) {
	log := writeLog(t, recAlpha, recBad, recBeta)
	_, rows := runReport(t, "summary", "--target", "Beta", log)
	if rows[1][0] != "1" || rows[1][1] != "1" {
		t.Errorf("records,skipped = %q,%q; want 1,1", rows[1][0], rows[1][1])
	}
}

func TestEmptyTargetMatchesNothing(t *testing.T) {
	log := writeLog(t, recAlpha)
	code, rows := runReport(t, "commands", "--target", "", log)
	if code != exitOK {
		t.Fatalf("exit = %d, want 0", code)
	}
	want := [][]string{
		{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"},
		{"TOTAL", "0", "0", "0", "0", "0", "0.0"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %q, want %q", rows, want)
	}
}

func TestFailuresReportWritesEmptyReturncodeForNull(t *testing.T) {
	log := writeLog(t, recAlpha, recBeta)
	_, rows := runReport(t, "failures", log)
	want := [][]string{
		{"started_at", "target", "command", "model", "returncode", "timed_out"},
		{"2026-10-16T06:00:00.000Z", "Beta", "plan", "sonnet", "", "true"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("rows = %q, want %q", rows, want)
	}
}

func TestEveryReportRunsOverEmptyLog(t *testing.T) {
	log := writeLog(t)
	for _, name := range []string{"summary", "commands", "targets", "models", "daily", "failures"} {
		t.Run(name, func(t *testing.T) {
			code, rows := runReport(t, name, log)
			if code != exitOK {
				t.Fatalf("exit = %d, want 0", code)
			}
			if len(rows) == 0 {
				t.Fatal("no header row written")
			}
		})
	}
}
