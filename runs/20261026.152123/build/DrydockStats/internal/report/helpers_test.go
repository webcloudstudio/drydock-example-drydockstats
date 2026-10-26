package report

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"ddkstat/internal/logread"
)

// rec returns a succeeded record with fixed numbers, so each test varies only the fields it names.
func rec(id, started, command, target string) logread.Record {
	return logread.Record{
		ExecutionID:    id,
		StartedAt:      started,
		Date:           started[:10],
		Status:         "succeeded",
		Command:        command,
		Target:         target,
		Provider:       "claude",
		RequestedModel: "claude-haiku-5-5",
		Model:          "claude-haiku-5-5",
		InputTokens:    1,
		CachedTokens:   2,
		OutputTokens:   3,
		CostUSD:        0.5,
		ElapsedMS:      60000,
		ReturnCode:     intPtr(0),
	}
}

// asFailed returns r with the status and failure fields a failed record carries.
func asFailed(r logread.Record) logread.Record {
	r.Status = "failed"
	r.Failed = true
	r.ReturnCode = intPtr(1)
	return r
}

func intPtr(n int) *int { return &n }

// writeFunc is the signature shared by every report writer.
type writeFunc func(io.Writer, []logread.Record, int) error

// readCSV runs write over records and reads the bytes it wrote back as CSV rows.
func readCSV(t *testing.T, write writeFunc, records []logread.Record, skipped int) [][]string {
	t.Helper()
	var buf bytes.Buffer
	if err := write(&buf, records, skipped); err != nil {
		t.Fatalf("write: %v", err)
	}
	rows, err := csv.NewReader(&buf).ReadAll()
	if err != nil {
		t.Fatalf("output is not CSV: %v", err)
	}
	return rows
}

// wantRows fails the test unless got equals want.
func wantRows(t *testing.T, got, want [][]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %v, want %v", got, want)
	}
}

// logLine returns one valid JSON Lines record with the given identifier, start date and target.
func logLine(id, date, target string) string {
	return `{"execution_id":"` + id + `","started_at":"` + date + `T02:38:41.904Z","status":"succeeded",` +
		`"job":{"command_name":"build","target":"` + target + `","llm":"claude","model":"claude-haiku-5-5"},` +
		`"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":1,"cached_input_tokens":2,` +
		`"output_tokens":3,"cost_usd":0.5,"elapsed_ms":4}}}`
}

// writeLog writes lines to a fresh log file under t.TempDir and returns its path.
func writeLog(t *testing.T, lines ...string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "log.jsonl")
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		t.Fatalf("write log: %v", err)
	}
	return path
}
