package logrecord

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// line builds one valid record line, with overrides applied by the caller.
const validLine = `{"schema_version":1,"execution_id":"%s","started_at":"2026-10-14T01:00:00.000Z",` +
	`"status":"succeeded","job":{"command_name":"build","target":"Toml","llm":"claude","model":"haiku"},` +
	`"result":{"returncode":0,"timed_out":false,"stats":{"model":"claude-haiku-5-5","input_tokens":10,` +
	`"cached_input_tokens":20,"output_tokens":30,"cost_usd":0.25,"elapsed_ms":1500}}}`

func valid(id string) string {
	return strings.Replace(validLine, "%s", id, 1)
}

func mustRead(t *testing.T, input string) Result {
	t.Helper()
	res, err := Read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Read returned error: %v", err)
	}
	return res
}

func TestReadDecodesOneRecordPerLine(t *testing.T) {
	res := mustRead(t, valid("a")+"\n"+valid("b")+"\n")
	if len(res.Records) != 2 {
		t.Fatalf("records = %d, want 2", len(res.Records))
	}
	if res.Records[0].ExecutionID != "a" || res.Records[1].ExecutionID != "b" {
		t.Errorf("ids = %q, %q; want a, b in file order", res.Records[0].ExecutionID, res.Records[1].ExecutionID)
	}
	if res.Skipped != 0 {
		t.Errorf("skipped = %d, want 0", res.Skipped)
	}
}

func TestReadIgnoresBlankLinesWithoutCountingThem(t *testing.T) {
	res := mustRead(t, "\n"+valid("a")+"\n   \n\t\n"+valid("b")+"\n\n")
	if len(res.Records) != 2 {
		t.Errorf("records = %d, want 2", len(res.Records))
	}
	if res.Skipped != 0 {
		t.Errorf("skipped = %d, want 0", res.Skipped)
	}
}

func TestReadEmptyInputYieldsNoRecords(t *testing.T) {
	res := mustRead(t, "")
	if len(res.Records) != 0 || res.Skipped != 0 {
		t.Errorf("got %d records, %d skipped; want 0, 0", len(res.Records), res.Skipped)
	}
}

func TestReadSkipsLinesThatAreNotJSONObjects(t *testing.T) {
	cases := map[string]string{
		"not json":         "not json at all",
		"json array":       `[1,2,3]`,
		"json scalar":      `42`,
		"trailing data":    valid("x") + " extra",
		"truncated object": `{"execution_id":"x"`,
	}
	for name, bad := range cases {
		t.Run(name, func(t *testing.T) {
			res := mustRead(t, bad+"\n"+valid("ok")+"\n")
			if res.Skipped != 1 {
				t.Errorf("skipped = %d, want 1", res.Skipped)
			}
			if len(res.Records) != 1 || res.Records[0].ExecutionID != "ok" {
				t.Errorf("records = %+v, want only ok", res.Records)
			}
		})
	}
}

func TestReadWarningNamesTheLineNumber(t *testing.T) {
	res := mustRead(t, valid("a")+"\n\nnot json\n")
	if len(res.Warnings) != 1 {
		t.Fatalf("warnings = %d, want 1", len(res.Warnings))
	}
	if !strings.HasPrefix(res.Warnings[0], "line 3 ") {
		t.Errorf("warning = %q, want it to name line 3", res.Warnings[0])
	}
}

func TestReadSkipsRecordsMissingEachRequiredField(t *testing.T) {
	cases := map[string]string{
		"execution_id":     `{"started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"b","target":"T"}}`,
		"started_at":       `{"execution_id":"x","status":"succeeded","job":{"command_name":"b","target":"T"}}`,
		"status":           `{"execution_id":"x","started_at":"2026-10-14T00:00:00.000Z","job":{"command_name":"b","target":"T"}}`,
		"job.command_name": `{"execution_id":"x","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"target":"T"}}`,
		"job.target":       `{"execution_id":"x","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"b"}}`,
		"empty job.target": `{"execution_id":"x","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"b","target":""}}`,
		"null target":      `{"execution_id":"x","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"b","target":null}}`,
		"non-string id":    `{"execution_id":7,"started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"b","target":"T"}}`,
		"no job object":    `{"execution_id":"x","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded"}`,
	}
	for name, bad := range cases {
		t.Run(name, func(t *testing.T) {
			res := mustRead(t, bad+"\n")
			if len(res.Records) != 0 {
				t.Errorf("records = %d, want 0", len(res.Records))
			}
			if res.Skipped != 1 {
				t.Errorf("skipped = %d, want 1", res.Skipped)
			}
		})
	}
}

func TestReadDuplicateExecutionIDKeepsFirstOccurrence(t *testing.T) {
	first := strings.Replace(valid("dup"), `"status":"succeeded"`, `"status":"failed"`, 1)
	second := valid("dup")
	res := mustRead(t, first+"\n"+second+"\n"+valid("other")+"\n")
	if len(res.Records) != 2 {
		t.Fatalf("records = %d, want 2", len(res.Records))
	}
	if res.Records[0].Status != "failed" {
		t.Errorf("kept status = %q, want the first occurrence's failed", res.Records[0].Status)
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
}

func TestReadDuplicateCheckIgnoresLinesThatFailedValidation(t *testing.T) {
	invalidFirst := `{"execution_id":"dup","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"b"}}`
	res := mustRead(t, invalidFirst+"\n"+valid("dup")+"\n")
	if len(res.Records) != 1 || res.Records[0].ExecutionID != "dup" {
		t.Errorf("records = %+v, want the later valid dup kept", res.Records)
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
}

func TestReadNullAndAbsentNumericsAreZero(t *testing.T) {
	nullCost := `{"execution_id":"n","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded",` +
		`"job":{"command_name":"b","target":"T"},"result":{"stats":{"cost_usd":null,"input_tokens":null}}}`
	noStats := `{"execution_id":"m","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded",` +
		`"job":{"command_name":"b","target":"T"},"result":{"returncode":0}}`
	noResult := `{"execution_id":"r","started_at":"2026-10-14T00:00:00.000Z","status":"failed",` +
		`"job":{"command_name":"b","target":"T"}}`
	res := mustRead(t, nullCost+"\n"+noStats+"\n"+noResult+"\n")
	if len(res.Records) != 3 {
		t.Fatalf("records = %d, want 3; records with no stats are still records", len(res.Records))
	}
	for _, r := range res.Records {
		if r.CostUSD != 0 || r.InputTokens != 0 || r.CachedTokens != 0 || r.OutputTokens != 0 || r.ElapsedMS != 0 {
			t.Errorf("record %s numerics = %+v, want all zero", r.ExecutionID, r)
		}
	}
}

func TestReadNonNumericStatsAreSkipped(t *testing.T) {
	bad := `{"execution_id":"s","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded",` +
		`"job":{"command_name":"b","target":"T"},"result":{"stats":{"input_tokens":"ten"}}}`
	res := mustRead(t, bad+"\n")
	if len(res.Records) != 0 || res.Skipped != 1 {
		t.Errorf("records = %d, skipped = %d; want 0, 1", len(res.Records), res.Skipped)
	}
}

func TestReadModelPrefersExecutedModel(t *testing.T) {
	res := mustRead(t, valid("a")+"\n")
	r := res.Records[0]
	if r.Model != "claude-haiku-5-5" {
		t.Errorf("model = %q, want executed claude-haiku-5-5", r.Model)
	}
	if r.Provider != "claude" {
		t.Errorf("provider = %q, want claude", r.Provider)
	}
}

func TestReadModelFallsBackToRequestedWhenExecutedMissingOrEmpty(t *testing.T) {
	empty := strings.Replace(valid("e"), `"model":"claude-haiku-5-5"`, `"model":""`, 1)
	noStats := `{"execution_id":"f","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded",` +
		`"job":{"command_name":"b","target":"T","llm":"codex","model":"gpt-5.6-luna"}}`
	res := mustRead(t, empty+"\n"+noStats+"\n")
	if res.Records[0].Model != "haiku" {
		t.Errorf("empty stats.model: model = %q, want requested haiku", res.Records[0].Model)
	}
	if res.Records[1].Model != "gpt-5.6-luna" {
		t.Errorf("absent stats.model: model = %q, want requested gpt-5.6-luna", res.Records[1].Model)
	}
}

func TestReadAbsentProviderIsEmpty(t *testing.T) {
	line := `{"execution_id":"p","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded",` +
		`"job":{"command_name":"b","target":"T","model":"haiku"}}`
	res := mustRead(t, line+"\n")
	if res.Records[0].Provider != "" {
		t.Errorf("provider = %q, want empty", res.Records[0].Provider)
	}
}

func TestReadReturnCodeNullIsNil(t *testing.T) {
	line := `{"execution_id":"rc","started_at":"2026-10-14T00:00:00.000Z","status":"failed",` +
		`"job":{"command_name":"b","target":"T"},"result":{"returncode":null,"timed_out":true}}`
	res := mustRead(t, line+"\n")
	r := res.Records[0]
	if r.ReturnCode != nil {
		t.Errorf("returncode = %d, want nil for null", *r.ReturnCode)
	}
	if !r.TimedOut {
		t.Errorf("timed_out = false, want true")
	}
}

func TestReadStartDateIsFirstTenCharacters(t *testing.T) {
	res := mustRead(t, valid("a")+"\n")
	if got := res.Records[0].StartDate(); got != "2026-10-14" {
		t.Errorf("StartDate = %q, want 2026-10-14", got)
	}
	short := Record{StartedAt: "2026"}
	if got := short.StartDate(); got != "2026" {
		t.Errorf("StartDate of short value = %q, want 2026", got)
	}
}

func TestReadFileOpensAndDecodes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "log.jsonl")
	if err := os.WriteFile(path, []byte(valid("a")+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if len(res.Records) != 1 {
		t.Errorf("records = %d, want 1", len(res.Records))
	}
}

func TestReadFileMissingIsError(t *testing.T) {
	_, err := ReadFile(filepath.Join(t.TempDir(), "absent.jsonl"))
	if err == nil {
		t.Fatal("ReadFile on missing path returned nil error")
	}
}

func TestReadFileDirectoryIsError(t *testing.T) {
	_, err := ReadFile(t.TempDir())
	if err == nil {
		t.Fatal("ReadFile on a directory returned nil error")
	}
}

func TestReadOverlongLineIsError(t *testing.T) {
	long := `{"execution_id":"` + strings.Repeat("x", maxLineBytes+1) + `"}`
	_, err := Read(strings.NewReader(long))
	if err == nil {
		t.Fatal("Read of an overlong line returned nil error")
	}
}
