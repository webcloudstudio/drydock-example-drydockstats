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

const recordA = `{"schema_version":1,"execution_id":"a","started_at":"2026-10-14T01:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Alpha","llm":"claude","model":"haiku"},"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":10,"cached_input_tokens":5,"output_tokens":2,"cost_usd":0.5,"elapsed_ms":60000}}}`
const recordB = `{"schema_version":1,"execution_id":"b","started_at":"2026-10-16T01:00:00.000Z","status":"failed","job":{"command_name":"diagnose","target":"Beta","llm":"claude","model":"haiku"},"result":{"returncode":null,"timed_out":true,"stats":{"input_tokens":3,"cached_input_tokens":0,"output_tokens":1,"cost_usd":0.25,"elapsed_ms":0}}}`

// writeLog writes lines to a fresh log file and returns its path.
func writeLog(t *testing.T, lines ...string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "log.jsonl")
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// invoke runs Run and returns the exit code with what it wrote to each stream.
func invoke(args ...string) (int, string, string) {
	var stdout, stderr bytes.Buffer
	code := Run(args, &stdout, &stderr)
	return code, stdout.String(), stderr.String()
}

// parseCSV reads stdout back as CSV rows.
func parseCSV(t *testing.T, out string) [][]string {
	t.Helper()
	rows, err := csv.NewReader(strings.NewReader(out)).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%s", err, out)
	}
	return rows
}

func TestUsageErrorsExitTwoWithEmptyStdout(t *testing.T) {
	log := writeLog(t, recordA)
	cases := map[string][]string{
		"no arguments":          {},
		"unknown report":        {"frobnicate", log},
		"unknown option":        {"summary", "--colour", "red", log},
		"option missing value":  {"summary", "--target"},
		"missing log path":      {"summary"},
		"extra log path":        {"summary", log, log},
		"path before option":    {"summary", log, "--target", "Alpha"},
		"impossible since date": {"summary", "--since", "2026-02-30", log},
		"month 13 until date":   {"summary", "--until", "2026-13-01", log},
		"malformed since date":  {"summary", "--since", "10/14/2026", log},
		"empty since date":      {"summary", "--since", "", log},
	}
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			code, out, _ := invoke(args...)
			if code != 2 {
				t.Errorf("exit = %d, want 2", code)
			}
			if out != "" {
				t.Errorf("stdout = %q, want empty", out)
			}
		})
	}
}

func TestUsageErrorIsCheckedBeforeLogIsOpened(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "absent.jsonl")
	code, _, _ := invoke("summary", "--since", "2026-02-30", missing)
	if code != 2 {
		t.Errorf("exit = %d, want 2 (usage error takes precedence over the missing log)", code)
	}
}

func TestUnreadableLogExitsOneWithEmptyStdout(t *testing.T) {
	cases := map[string]string{
		"missing file": filepath.Join(t.TempDir(), "absent.jsonl"),
		"directory":    t.TempDir(),
	}
	for name, path := range cases {
		t.Run(name, func(t *testing.T) {
			code, out, errOut := invoke("summary", path)
			if code != 1 {
				t.Errorf("exit = %d, want 1", code)
			}
			if out != "" {
				t.Errorf("stdout = %q, want empty", out)
			}
			if strings.Count(errOut, "\n") != 1 {
				t.Errorf("stderr = %q, want exactly one line", errOut)
			}
		})
	}
}

func TestHelpFlagExitsZeroWithoutReport(t *testing.T) {
	code, out, _ := invoke("summary", "-h")
	if code != 0 {
		t.Errorf("exit = %d, want 0", code)
	}
	if out != "" {
		t.Errorf("stdout = %q, want empty", out)
	}
}

func TestEveryReportWritesItsHeaderOnAValidLog(t *testing.T) {
	log := writeLog(t, recordA)
	headers := map[string]string{
		"summary":  "records",
		"commands": "command",
		"targets":  "target",
		"models":   "provider",
		"daily":    "date",
		"failures": "started_at",
	}
	for name, first := range headers {
		t.Run(name, func(t *testing.T) {
			code, out, _ := invoke(name, log)
			if code != 0 {
				t.Fatalf("exit = %d, want 0", code)
			}
			rows := parseCSV(t, out)
			if len(rows) == 0 || rows[0][0] != first {
				t.Errorf("header = %v, want first column %q", rows, first)
			}
		})
	}
}

func TestSummaryReportsCountsAndSkippedLines(t *testing.T) {
	log := writeLog(t, recordA, "not json", recordA, recordB)
	code, out, errOut := invoke("summary", log)
	if code != 0 {
		t.Fatalf("exit = %d, want 0 with skipped lines", code)
	}
	if errOut == "" {
		t.Errorf("stderr is empty; skipped lines should produce warnings")
	}
	rows := parseCSV(t, out)
	if len(rows) != 2 {
		t.Fatalf("rows = %d, want header plus one", len(rows))
	}
	got := rows[1]
	want := []string{"2", "2", "1", "1", "2026-10-14", "2026-10-16", "13", "5", "3", "0.7500", "1.0"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("summary = %v, want %v", got, want)
	}
}

func TestTargetOptionNarrowsEveryReport(t *testing.T) {
	log := writeLog(t, recordA, recordB)
	code, out, _ := invoke("summary", "--target", "Alpha", log)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, out)
	if rows[1][0] != "1" || rows[1][1] != "0" {
		t.Errorf("records/skipped = %s/%s, want 1/0 (skipped is whole-file)", rows[1][0], rows[1][1])
	}
}

func TestDateOptionsNarrowTheRange(t *testing.T) {
	log := writeLog(t, recordA, recordB)
	code, out, _ := invoke("daily", "--since", "2026-10-15", "--until", "2026-10-16", log)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, out)
	if len(rows) != 3 || rows[1][0] != "2026-10-16" {
		t.Errorf("rows = %v, want only the 2026-10-16 group and TOTAL", rows)
	}
}

func TestNoMatchExitsZeroWithHeaderAndZeroTotal(t *testing.T) {
	log := writeLog(t, recordA)
	code, out, _ := invoke("commands", "--target", "NoSuchTarget", log)
	if code != 0 {
		t.Fatalf("exit = %d, want 0 on empty match", code)
	}
	rows := parseCSV(t, out)
	if len(rows) != 2 {
		t.Fatalf("rows = %d, want header plus TOTAL", len(rows))
	}
	if !reflect.DeepEqual(rows[1], []string{"TOTAL", "0", "0", "0", "0", "0", "0.0"}) {
		t.Errorf("TOTAL = %v, want zero row", rows[1])
	}
}

func TestFailuresReportListsOnlyFailedRecords(t *testing.T) {
	log := writeLog(t, recordA, recordB)
	code, out, _ := invoke("failures", log)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, out)
	if len(rows) != 2 {
		t.Fatalf("rows = %d, want header plus one failure", len(rows))
	}
	if rows[1][4] != "" || rows[1][5] != "true" {
		t.Errorf("failure row = %v, want empty returncode and timed_out true", rows[1])
	}
}

func TestFailuresEmptyMatchWritesHeaderOnly(t *testing.T) {
	log := writeLog(t, recordA)
	code, out, _ := invoke("failures", log)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, out)
	if len(rows) != 1 {
		t.Errorf("rows = %d, want header only", len(rows))
	}
}

func TestEmptyLogWritesSummaryZeroRow(t *testing.T) {
	log := writeLog(t)
	code, out, _ := invoke("summary", log)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	rows := parseCSV(t, out)
	if rows[1][0] != "0" || rows[1][4] != "" || rows[1][5] != "" {
		t.Errorf("summary = %v, want zero records with empty first and last", rows[1])
	}
}
