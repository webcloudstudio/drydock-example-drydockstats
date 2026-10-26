package logread_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"ddkstat/internal/logread"
)

// entry builds one JSON Lines object with the fields Read consumes.
func entry(id, status, startedAt string) string {
	return fmt.Sprintf(`{"schema_version":1,"execution_id":%q,"started_at":%q,"status":%q,`+
		`"job":{"command_name":"build","target":"Toml","llm":"claude","model":"haiku"},`+
		`"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":10,"cached_input_tokens":20,`+
		`"output_tokens":30,"cost_usd":0.5,"elapsed_ms":40,"model":"claude-haiku-5-5"}}}`,
		id, startedAt, status)
}

func read(t *testing.T, input string) logread.Result {
	t.Helper()
	res, err := logread.Read(strings.NewReader(input), &bytes.Buffer{})
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	return res
}

func TestReadBlankLines(t *testing.T) {
	input := "\n" +
		"   \n" +
		entry("a", "succeeded", "2026-10-14T06:37:54.176Z") + "\n" +
		"\t\n" +
		"\n" +
		entry("b", "succeeded", "2026-10-15T06:37:54.176Z") + "\n" +
		"\n"
	res := read(t, input)
	if len(res.Records) != 2 {
		t.Fatalf("records = %d, want 2", len(res.Records))
	}
	if res.Records[0].ExecutionID != "a" || res.Records[1].ExecutionID != "b" {
		t.Errorf("ids = %q, %q; want a, b", res.Records[0].ExecutionID, res.Records[1].ExecutionID)
	}
	if res.Skipped != 0 {
		t.Errorf("skipped = %d, want 0 (blank lines are not skips)", res.Skipped)
	}
}

func TestReadBlankOnlyInput(t *testing.T) {
	res := read(t, "\n\n   \n")
	if len(res.Records) != 0 || res.Skipped != 0 {
		t.Errorf("got %d records, %d skipped; want 0, 0", len(res.Records), res.Skipped)
	}
}

func TestReadEmptyInput(t *testing.T) {
	res := read(t, "")
	if len(res.Records) != 0 || res.Skipped != 0 {
		t.Errorf("got %d records, %d skipped; want 0, 0", len(res.Records), res.Skipped)
	}
}

func TestReadLastLineWithoutNewline(t *testing.T) {
	res := read(t, entry("a", "succeeded", "2026-10-14T06:37:54.176Z"))
	if len(res.Records) != 1 || res.Records[0].ExecutionID != "a" {
		t.Errorf("records = %+v, want one record a", res.Records)
	}
}

func TestReadFailedStatus(t *testing.T) {
	input := entry("ok", "succeeded", "2026-10-14T06:00:00Z") + "\n" +
		entry("bad", "failed", "2026-10-14T06:00:00Z") + "\n" +
		entry("slow", "timeout", "2026-10-14T06:00:00Z") + "\n"
	res := read(t, input)
	want := map[string]bool{"ok": false, "bad": true, "slow": true}
	if len(res.Records) != len(want) {
		t.Fatalf("records = %d, want %d", len(res.Records), len(want))
	}
	for _, rec := range res.Records {
		if rec.Failed != want[rec.ExecutionID] {
			t.Errorf("%s: Failed = %t, want %t", rec.ExecutionID, rec.Failed, want[rec.ExecutionID])
		}
	}
}

func TestReadStartDate(t *testing.T) {
	res := read(t, entry("a", "succeeded", "2026-10-14T06:37:54.176Z"))
	if got := res.Records[0].Date; got != "2026-10-14" {
		t.Errorf("Date = %q, want 2026-10-14", got)
	}
}

func TestReadStartDateShorterThanTenCharacters(t *testing.T) {
	res := read(t, entry("a", "succeeded", "2026-1"))
	if got := res.Records[0].Date; got != "2026-1" {
		t.Errorf("Date = %q, want 2026-1", got)
	}
}

func TestReadStartedAtNotValidated(t *testing.T) {
	res := read(t, entry("a", "succeeded", "not a timestamp"))
	if len(res.Records) != 1 {
		t.Fatalf("records = %d, want 1", len(res.Records))
	}
	if res.Records[0].StartedAt != "not a timestamp" {
		t.Errorf("StartedAt = %q, want the raw value", res.Records[0].StartedAt)
	}
	if res.Records[0].Date != "not a time" {
		t.Errorf("Date = %q, want first 10 characters", res.Records[0].Date)
	}
	if res.Skipped != 0 {
		t.Errorf("skipped = %d, want 0", res.Skipped)
	}
}

func TestReadMapsFields(t *testing.T) {
	rec := read(t, entry("a", "succeeded", "2026-10-14T06:37:54.176Z")).Records[0]
	if rec.Command != "build" || rec.Target != "Toml" || rec.Provider != "claude" {
		t.Errorf("command/target/provider = %q/%q/%q", rec.Command, rec.Target, rec.Provider)
	}
	if rec.RequestedModel != "haiku" || rec.Model != "claude-haiku-5-5" {
		t.Errorf("requested/executed model = %q/%q", rec.RequestedModel, rec.Model)
	}
	if rec.InputTokens != 10 || rec.CachedTokens != 20 || rec.OutputTokens != 30 {
		t.Errorf("tokens = %d/%d/%d, want 10/20/30", rec.InputTokens, rec.CachedTokens, rec.OutputTokens)
	}
	if rec.CostUSD != 0.5 || rec.ElapsedMS != 40 {
		t.Errorf("cost/elapsed = %v/%d, want 0.5/40", rec.CostUSD, rec.ElapsedMS)
	}
	if rec.ReturnCode == nil || *rec.ReturnCode != 0 || rec.TimedOut {
		t.Errorf("returncode/timed_out = %v/%t, want 0/false", rec.ReturnCode, rec.TimedOut)
	}
}

func TestReadNullReturnCodeIsNil(t *testing.T) {
	line := strings.Replace(entry("a", "failed", "2026-10-14T06:00:00Z"),
		`"returncode":0`, `"returncode":null`, 1)
	rec := read(t, line).Records[0]
	if rec.ReturnCode != nil {
		t.Errorf("ReturnCode = %d, want nil", *rec.ReturnCode)
	}
}

func TestReadMalformedLineIsSkipped(t *testing.T) {
	input := entry("a", "succeeded", "2026-10-14T06:00:00Z") + "\n" +
		"{not json\n" +
		entry("b", "succeeded", "2026-10-14T06:00:00Z") + "\n"
	res := read(t, input)
	if len(res.Records) != 2 {
		t.Errorf("records = %d, want 2", len(res.Records))
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
}

func TestReadNonObjectLineIsSkipped(t *testing.T) {
	res := read(t, "[1,2,3]\n"+"42\n")
	if len(res.Records) != 0 || res.Skipped != 2 {
		t.Errorf("got %d records, %d skipped; want 0, 2", len(res.Records), res.Skipped)
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("disk on fire") }

func TestReadReaderErrorIsReturned(t *testing.T) {
	if _, err := logread.Read(failingReader{}, &bytes.Buffer{}); err == nil {
		t.Error("Read returned nil error for a failing reader")
	}
}

func TestReadNilWarnWriter(t *testing.T) {
	res, err := logread.Read(strings.NewReader("{bad\n"), nil)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
}

func TestSkipNotObject(t *testing.T) {
	input := "[1,2,3]\n" + "\"text\"\n" + "42\n" + "true\n" + "null\n" +
		entry("a", "succeeded", "2026-10-14T06:00:00Z") + "\n"
	res := read(t, input)
	if len(res.Records) != 1 || res.Records[0].ExecutionID != "a" {
		t.Errorf("records = %+v, want only a", res.Records)
	}
	if res.Skipped != 5 {
		t.Errorf("skipped = %d, want 5", res.Skipped)
	}
}

func TestSkipMissingField(t *testing.T) {
	valid := entry("a", "succeeded", "2026-10-14T06:00:00Z")
	cases := map[string]string{
		"execution_id":     strings.Replace(valid, `"execution_id":"a",`, "", 1),
		"started_at":       strings.Replace(valid, `"started_at":"2026-10-14T06:00:00Z",`, "", 1),
		"status":           strings.Replace(valid, `"status":"succeeded",`, "", 1),
		"job.command_name": strings.Replace(valid, `"command_name":"build",`, "", 1),
		"job.target":       strings.Replace(valid, `"target":"Toml",`, "", 1),
	}
	for field, line := range cases {
		if line == valid {
			t.Fatalf("%s: test input was not altered", field)
		}
		res := read(t, line+"\n")
		if len(res.Records) != 0 || res.Skipped != 1 {
			t.Errorf("without %s: %d records, %d skipped; want 0, 1", field, len(res.Records), res.Skipped)
		}
	}
}

func TestSkipEmptyStringField(t *testing.T) {
	cases := map[string]string{
		"execution_id":      entry("", "succeeded", "2026-10-14T06:00:00Z"),
		"started_at":        entry("a", "succeeded", ""),
		"status":            entry("a", "", "2026-10-14T06:00:00Z"),
		"job.command_name":  strings.Replace(entry("a", "succeeded", "2026-10-14T06:00:00Z"), `"command_name":"build"`, `"command_name":""`, 1),
		"job.target":        strings.Replace(entry("a", "succeeded", "2026-10-14T06:00:00Z"), `"target":"Toml"`, `"target":""`, 1),
		"status as number":  strings.Replace(entry("a", "succeeded", "2026-10-14T06:00:00Z"), `"status":"succeeded"`, `"status":5`, 1),
		"execution_id null": strings.Replace(entry("a", "succeeded", "2026-10-14T06:00:00Z"), `"execution_id":"a"`, `"execution_id":null`, 1),
	}
	for field, line := range cases {
		res := read(t, line+"\n")
		if len(res.Records) != 0 || res.Skipped != 1 {
			t.Errorf("%s: %d records, %d skipped; want 0, 1", field, len(res.Records), res.Skipped)
		}
	}
}

func TestSkipDuplicateKeepsFirst(t *testing.T) {
	input := entry("a", "succeeded", "2026-10-14T06:00:00Z") + "\n" +
		entry("a", "failed", "2026-10-15T06:00:00Z") + "\n"
	res := read(t, input)
	if len(res.Records) != 1 {
		t.Fatalf("records = %d, want 1", len(res.Records))
	}
	if rec := res.Records[0]; rec.Status != "succeeded" || rec.Date != "2026-10-14" {
		t.Errorf("kept record status/date = %q/%q, want the first occurrence", rec.Status, rec.Date)
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
}

func TestMalformedDoesNotRegisterID(t *testing.T) {
	input := entry("a", "", "2026-10-14T06:00:00Z") + "\n" +
		entry("a", "succeeded", "2026-10-14T06:00:00Z") + "\n"
	res := read(t, input)
	if len(res.Records) != 1 || res.Records[0].Status != "succeeded" {
		t.Fatalf("records = %+v, want the later valid line a kept", res.Records)
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
}

func TestSkipWarningLineNumber(t *testing.T) {
	input := "\n" +
		"{bad\n" +
		entry("a", "succeeded", "2026-10-14T06:00:00Z") + "\n" +
		"\n" +
		"[1]\n"
	var warn bytes.Buffer
	res, err := logread.Read(strings.NewReader(input), &warn)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if len(res.Skips) != 2 || res.Skips[0].Line != 2 || res.Skips[1].Line != 5 {
		t.Fatalf("skip lines = %+v, want lines 2 and 5", res.Skips)
	}
	if res.Skipped != len(res.Skips) {
		t.Errorf("skipped = %d, want %d", res.Skipped, len(res.Skips))
	}
	if warnings := bytes.Count(warn.Bytes(), []byte("\n")); warnings != res.Skipped {
		t.Errorf("warning lines = %d, want one per skipped line (%d)", warnings, res.Skipped)
	}
}

// withResult builds one valid line whose result member is the given raw JSON. The job model is
// "haiku" and the line is otherwise complete.
func withResult(result string) string {
	return `{"execution_id":"a","started_at":"2026-10-14T06:00:00Z","status":"succeeded",` +
		`"job":{"command_name":"build","target":"Toml","llm":"claude","model":"haiku"},` +
		`"result":` + result + `}`
}

// readOne reads input and requires exactly one kept record and no skips.
func readOne(t *testing.T, input string) logread.Record {
	t.Helper()
	res := read(t, input)
	if len(res.Records) != 1 || res.Skipped != 0 {
		t.Fatalf("records = %d, skipped = %d; want 1, 0", len(res.Records), res.Skipped)
	}
	return res.Records[0]
}

func requireZeroNumbers(t *testing.T, rec logread.Record) {
	t.Helper()
	if rec.InputTokens != 0 || rec.CachedTokens != 0 || rec.OutputTokens != 0 ||
		rec.CostUSD != 0 || rec.ElapsedMS != 0 {
		t.Errorf("numbers = %d/%d/%d/%v/%d, want all zero",
			rec.InputTokens, rec.CachedTokens, rec.OutputTokens, rec.CostUSD, rec.ElapsedMS)
	}
}

func TestNumbersAbsentAreZero(t *testing.T) {
	rec := readOne(t, withResult(`{"returncode":0,"timed_out":false,"stats":{"model":"claude-haiku-5-5"}}`))
	requireZeroNumbers(t, rec)
	if rec.Model != "claude-haiku-5-5" {
		t.Errorf("Model = %q, want the stats model", rec.Model)
	}
}

func TestNumbersNullAreZero(t *testing.T) {
	rec := readOne(t, withResult(`{"returncode":0,"timed_out":false,"stats":{`+
		`"input_tokens":null,"cached_input_tokens":null,"output_tokens":null,`+
		`"cost_usd":null,"elapsed_ms":null}}`))
	requireZeroNumbers(t, rec)
}

func TestNumbersNonNumericAreZero(t *testing.T) {
	rec := readOne(t, withResult(`{"returncode":0,"timed_out":false,"stats":{`+
		`"input_tokens":"many","cached_input_tokens":true,"output_tokens":[1],`+
		`"cost_usd":"0.5","elapsed_ms":{"ms":3}}}`))
	requireZeroNumbers(t, rec)
}

func TestNumbersNonNumericKeepsOtherFields(t *testing.T) {
	rec := readOne(t, withResult(`{"stats":{"input_tokens":"many","output_tokens":30,"cost_usd":0.5}}`))
	if rec.InputTokens != 0 || rec.OutputTokens != 30 || rec.CostUSD != 0.5 {
		t.Errorf("input/output/cost = %d/%d/%v, want 0/30/0.5", rec.InputTokens, rec.OutputTokens, rec.CostUSD)
	}
}

func TestNoStatsObjectStillRecord(t *testing.T) {
	cases := map[string]string{
		"result without stats": `{"returncode":0,"timed_out":false}`,
		"stats not an object":  `{"returncode":0,"stats":"none"}`,
		"result not an object": `"none"`,
	}
	for name, result := range cases {
		rec := readOne(t, withResult(result))
		requireZeroNumbers(t, rec)
		if rec.ExecutionID != "a" {
			t.Errorf("%s: ExecutionID = %q, want a", name, rec.ExecutionID)
		}
	}
}

func TestNoResultObjectStillRecord(t *testing.T) {
	line := `{"execution_id":"a","started_at":"2026-10-14T06:00:00Z","status":"succeeded",` +
		`"job":{"command_name":"build","target":"Toml","llm":"claude","model":"haiku"}}`
	rec := readOne(t, line)
	requireZeroNumbers(t, rec)
	if rec.ReturnCode != nil || rec.TimedOut {
		t.Errorf("returncode/timed_out = %v/%t, want nil/false", rec.ReturnCode, rec.TimedOut)
	}
}

func TestReturnCodeNullIsNil(t *testing.T) {
	rec := readOne(t, withResult(`{"returncode":null,"timed_out":false,"stats":{}}`))
	if rec.ReturnCode != nil {
		t.Errorf("ReturnCode = %d, want nil", *rec.ReturnCode)
	}
}

func TestReturnCodeAbsentIsNil(t *testing.T) {
	rec := readOne(t, withResult(`{"timed_out":false,"stats":{}}`))
	if rec.ReturnCode != nil {
		t.Errorf("ReturnCode = %d, want nil", *rec.ReturnCode)
	}
}

func TestModelPrefersExecuted(t *testing.T) {
	rec := readOne(t, withResult(`{"stats":{"model":"claude-haiku-5-5"}}`))
	if rec.Model != "claude-haiku-5-5" || rec.RequestedModel != "haiku" {
		t.Errorf("Model/RequestedModel = %q/%q, want claude-haiku-5-5/haiku", rec.Model, rec.RequestedModel)
	}
}

func TestModelFallsBackWhenAbsent(t *testing.T) {
	rec := readOne(t, withResult(`{"stats":{"input_tokens":1}}`))
	if rec.Model != "haiku" {
		t.Errorf("Model = %q, want haiku", rec.Model)
	}
}

func TestModelFallsBackWhenEmpty(t *testing.T) {
	rec := readOne(t, withResult(`{"stats":{"model":""}}`))
	if rec.Model != "haiku" {
		t.Errorf("Model = %q, want haiku", rec.Model)
	}
}

func TestModelFallsBackWhenNotString(t *testing.T) {
	rec := readOne(t, withResult(`{"stats":{"model":7}}`))
	if rec.Model != "haiku" {
		t.Errorf("Model = %q, want haiku", rec.Model)
	}
}

func TestModelFallsBackWithoutStats(t *testing.T) {
	rec := readOne(t, withResult(`{"returncode":0}`))
	if rec.Model != "haiku" {
		t.Errorf("Model = %q, want haiku", rec.Model)
	}
}

func TestModelEmptyWhenJobModelAbsent(t *testing.T) {
	line := `{"execution_id":"a","started_at":"2026-10-14T06:00:00Z","status":"succeeded",` +
		`"job":{"command_name":"build","target":"Toml","llm":"claude"},` +
		`"result":{"stats":{}}}`
	rec := readOne(t, line)
	if rec.Model != "" || rec.RequestedModel != "" {
		t.Errorf("Model/RequestedModel = %q/%q, want empty", rec.Model, rec.RequestedModel)
	}
}

func TestReadDedupesWholeFile(t *testing.T) {
	var b strings.Builder
	for i := 0; i < 1000; i++ {
		b.WriteString(entry(fmt.Sprintf("id-%d", i), "succeeded", "2026-10-14T06:00:00Z") + "\n")
	}
	b.WriteString(entry("id-0", "failed", "2026-10-16T06:00:00Z") + "\n")
	res := read(t, b.String())
	if len(res.Records) != 1000 {
		t.Errorf("records = %d, want 1000", len(res.Records))
	}
	if res.Skipped != 1 {
		t.Errorf("skipped = %d, want 1", res.Skipped)
	}
	for _, rec := range res.Records {
		if rec.ExecutionID == "id-0" && rec.Failed {
			t.Errorf("id-0 kept the later duplicate's status")
		}
	}
}
