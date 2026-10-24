package log

import (
	"os"
	"strings"
	"testing"
)

const validLine = `{"schema_version":1,"execution_id":"e-1","started_at":"2026-10-14T06:37:54.176Z","status":"succeeded","job":{"command_name":"plan","target":"T","llm":"claude","model":"requested"},"result":{"returncode":0,"timed_out":false,"stats":{"model":"executed","cached_input_tokens":2,"elapsed_ms":1500,"input_tokens":3,"output_tokens":4,"cost_usd":0.5}}}`

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0o600)
}

// parseText reads and parses log text the same way the CLI does.
func parseText(t *testing.T, text string) Parsed {
	t.Helper()
	lines, err := ReadLines(strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	return Parse(lines)
}

func TestParseKeepsValidRecordFields(t *testing.T) {
	p := parseText(t, validLine+"\n")
	if len(p.Records) != 1 || len(p.Skips) != 0 {
		t.Fatalf("records=%d skips=%d, want 1 and 0", len(p.Records), len(p.Skips))
	}
	r := p.Records[0]
	if r.ExecutionID != "e-1" || r.StartDate != "2026-10-14" || r.Command != "plan" || r.Target != "T" || r.Provider != "claude" {
		t.Errorf("identity fields wrong: %+v", r)
	}
	if r.InputTokens != 3 || r.CachedTokens != 2 || r.OutputTokens != 4 || r.ElapsedMS != 1500 || r.CostUSD != 0.5 {
		t.Errorf("numeric fields wrong: %+v", r)
	}
}

func TestParseSkipsBadLinesWithLineNumbers(t *testing.T) {
	tests := []struct {
		name string
		line string
	}{
		{"not JSON", `garbage`},
		{"JSON array", `[1,2,3]`},
		{"JSON null", `null`},
		{"trailing data after object", validLine + ` x`},
		{"missing execution_id", `{"started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"b","target":"T"}}`},
		{"empty target", `{"execution_id":"x","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"b","target":""}}`},
		{"non-string status", `{"execution_id":"x","started_at":"2026-10-14T00:00:00Z","status":3,"job":{"command_name":"b","target":"T"}}`},
		{"missing job", `{"execution_id":"x","started_at":"2026-10-14T00:00:00Z","status":"succeeded"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parseText(t, validLine+"\n"+tt.line+"\n")
			if len(p.Records) != 1 {
				t.Errorf("kept %d records, want 1", len(p.Records))
			}
			if len(p.Skips) != 1 || p.Skips[0].Line != 2 {
				t.Errorf("skips = %+v, want one skip on line 2", p.Skips)
			}
		})
	}
}

func TestParseDuplicateKeepsFirstOccurrence(t *testing.T) {
	second := strings.Replace(validLine, `"status":"succeeded"`, `"status":"failed"`, 1)
	p := parseText(t, validLine+"\n"+second+"\n")
	if len(p.Records) != 1 || p.Records[0].Status != "succeeded" {
		t.Fatalf("records = %+v, want only the first succeeded occurrence", p.Records)
	}
	if len(p.Skips) != 1 || !strings.Contains(p.Skips[0].Reason, "duplicate") {
		t.Errorf("skips = %+v, want one duplicate skip", p.Skips)
	}
}

func TestParseDuplicateOfSkippedLineIsNotRememberedAsSeen(t *testing.T) {
	incomplete := `{"execution_id":"e-9","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"b"}}`
	p := parseText(t, incomplete+"\n"+strings.Replace(validLine, "e-1", "e-9", 1)+"\n")
	if len(p.Records) != 1 {
		t.Errorf("kept %d records, want 1: a skipped line must not block its execution_id", len(p.Records))
	}
}

func TestModelAttribution(t *testing.T) {
	tests := []struct {
		name  string
		stats string
		want  string
	}{
		{"executed model preferred", `"model":"executed",`, "executed"},
		{"empty executed model falls back", `"model":"",`, "requested"},
		{"absent executed model falls back", ``, "requested"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line := strings.Replace(validLine, `"model":"executed",`, tt.stats, 1)
			p := parseText(t, line+"\n")
			if len(p.Records) != 1 || p.Records[0].Model != tt.want {
				t.Errorf("model = %+v, want %q", p.Records, tt.want)
			}
		})
	}
}

func TestNullAndAbsentNumericStatsAreZero(t *testing.T) {
	line := `{"execution_id":"n","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"b","target":"T"},"result":{"stats":{"cost_usd":null,"elapsed_ms":null,"input_tokens":5}}}`
	p := parseText(t, line+"\n")
	if len(p.Records) != 1 {
		t.Fatalf("record dropped for null stats: %+v", p.Skips)
	}
	r := p.Records[0]
	if r.CostUSD != 0 || r.ElapsedMS != 0 || r.InputTokens != 5 || r.CachedTokens != 0 || r.OutputTokens != 0 {
		t.Errorf("numeric defaults wrong: %+v", r)
	}
}

func TestRecordWithoutStatsIsKept(t *testing.T) {
	line := `{"execution_id":"n","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"b","target":"T"},"result":{"returncode":0}}`
	p := parseText(t, line+"\n")
	if len(p.Records) != 1 || p.Records[0].InputTokens != 0 {
		t.Errorf("records = %+v, want one zeroed record", p.Records)
	}
}

func TestFailedClassification(t *testing.T) {
	tests := map[string]bool{"succeeded": false, "failed": true, "timeout": true, "SUCCEEDED": true}
	for status, want := range tests {
		line := strings.Replace(validLine, `"status":"succeeded"`, `"status":"`+status+`"`, 1)
		p := parseText(t, line+"\n")
		if len(p.Records) != 1 || p.Records[0].Failed() != want {
			t.Errorf("status %q: failed = %v, want %v", status, p.Records[0].Failed(), want)
		}
	}
}

func TestReturnCodeNullableAndTimedOut(t *testing.T) {
	nullRC := `{"execution_id":"f","started_at":"2026-10-14T00:00:00Z","status":"failed","job":{"command_name":"b","target":"T"},"result":{"returncode":null,"timed_out":true}}`
	p := parseText(t, nullRC+"\n")
	if len(p.Records) != 1 || p.Records[0].ReturnCode != nil || !p.Records[0].TimedOut {
		t.Errorf("record = %+v, want nil returncode and timed out", p.Records)
	}

	withRC := strings.Replace(nullRC, `"returncode":null`, `"returncode":7`, 1)
	p = parseText(t, withRC+"\n")
	if len(p.Records) != 1 || p.Records[0].ReturnCode == nil || *p.Records[0].ReturnCode != 7 {
		t.Errorf("record = %+v, want returncode 7", p.Records)
	}
}

func TestStartDateIsFirstTenCharacters(t *testing.T) {
	short := strings.Replace(validLine, `"2026-10-14T06:37:54.176Z"`, `"2026"`, 1)
	p := parseText(t, short+"\n")
	if len(p.Records) != 1 || p.Records[0].StartDate != "2026" {
		t.Errorf("start date = %+v, want 2026", p.Records)
	}
}
