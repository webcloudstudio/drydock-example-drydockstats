package logread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// baseRecord returns a valid log record as a generic JSON object. Tests change one field at a time.
func baseRecord(id string) map[string]any {
	return map[string]any{
		"schema_version": 1,
		"execution_id":   id,
		"started_at":     "2026-10-14T06:37:54.176Z",
		"status":         "succeeded",
		"job": map[string]any{
			"command_name": "build",
			"target":       "Toml",
			"llm":          "claude",
			"model":        "sonnet",
		},
		"prompt": map[string]any{"bytes": 191082},
		"result": map[string]any{
			"returncode": 0,
			"timed_out":  false,
			"stats": map[string]any{
				"cached_input_tokens": 263315,
				"elapsed_ms":          54872,
				"input_tokens":        843,
				"output_tokens":       7751,
				"turns":               16,
				"cost_usd":            0.0659295,
				"model":               "claude-sonnet-5-5",
			},
		},
	}
}

// lineOf encodes rec as one JSON Lines line without the trailing newline.
func lineOf(t *testing.T, rec map[string]any) string {
	t.Helper()
	b, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("marshal record: %v", err)
	}
	return string(b)
}

// readText reads input and returns the result and everything written to the warning writer.
func readText(t *testing.T, input string) (Result, string) {
	t.Helper()
	var warn bytes.Buffer
	res, err := Read(strings.NewReader(input), &warn)
	if err != nil {
		t.Fatalf("Read returned error: %v", err)
	}
	return res, warn.String()
}

// readOne reads a single line holding rec and returns the one record kept from it.
func readOne(t *testing.T, rec map[string]any) Record {
	t.Helper()
	res, _ := readText(t, lineOf(t, rec)+"\n")
	if len(res.Records) != 1 {
		t.Fatalf("kept %d records, want 1 (skipped %d)", len(res.Records), res.Skipped)
	}
	return res.Records[0]
}

// warnedLines returns the 1-based line numbers named by the warnings, one per warning line.
func warnedLines(t *testing.T, warn string) []int {
	t.Helper()
	if warn == "" {
		return nil
	}
	var got []int
	for _, l := range strings.Split(strings.TrimSuffix(warn, "\n"), "\n") {
		var n int
		if _, err := fmt.Sscanf(l, "ddkstat: line %d:", &n); err != nil {
			t.Fatalf("warning %q does not name a line: %v", l, err)
		}
		got = append(got, n)
	}
	return got
}

// ids returns the execution_id of each kept record, in order.
func ids(res Result) []string {
	var got []string
	for _, r := range res.Records {
		got = append(got, r.ExecutionID)
	}
	return got
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
