package logread

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"
	"unicode/utf8"
)

// dateLayout is the calendar date form every report keys on: the first ten characters of started_at.
const dateLayout = "2006-01-02"

// Record is one kept execution from the log.
type Record struct {
	ExecutionID  string
	StartedAt    string // RFC 3339 UTC; the first 10 characters are the start date
	Status       string
	Command      string // job.command_name
	Target       string // job.target
	Provider     string // job.llm
	Model        string // result.stats.model when non-empty, otherwise job.model
	InputTokens  int64
	CachedTokens int64
	OutputTokens int64
	CostUSD      float64
	ElapsedMS    int64
	ReturnCode   *int // nil when null or absent
	TimedOut     bool
}

// Date returns the start date, the first ten characters of StartedAt.
func (r Record) Date() string {
	if len(r.StartedAt) < len(dateLayout) {
		return ""
	}
	return r.StartedAt[:len(dateLayout)]
}

// Failed reports whether the execution did not succeed.
func (r Record) Failed() bool {
	return r.Status != "succeeded"
}

// decodeRecord decodes one trimmed, non-blank line into a Record. When the line breaks a skip
// rule it returns an error naming that rule.
func decodeRecord(line []byte) (Record, error) {
	if !utf8.Valid(line) {
		return Record{}, errors.New("not valid UTF-8")
	}
	top, ok := object(line)
	if !ok {
		return Record{}, errors.New("not a JSON object")
	}
	id, err := requiredString(top, "execution_id")
	if err != nil {
		return Record{}, err
	}
	startedAt, err := requiredString(top, "started_at")
	if err != nil {
		return Record{}, err
	}
	status, err := requiredString(top, "status")
	if err != nil {
		return Record{}, err
	}
	job, _ := object(top["job"])
	command, err := requiredString(job, "command_name")
	if err != nil {
		return Record{}, err
	}
	target, err := requiredString(job, "target")
	if err != nil {
		return Record{}, err
	}
	if !startsWithDate(startedAt) {
		return Record{}, fmt.Errorf("started_at %q does not begin with a YYYY-MM-DD date", startedAt)
	}

	result, _ := object(top["result"])
	stats, _ := object(result["stats"])

	model := optionalString(stats, "model")
	if model == "" {
		model = optionalString(job, "model")
	}
	rec := Record{
		ExecutionID:  id,
		StartedAt:    startedAt,
		Status:       status,
		Command:      command,
		Target:       target,
		Provider:     optionalString(job, "llm"),
		Model:        model,
		InputTokens:  intField(stats, "input_tokens"),
		CachedTokens: intField(stats, "cached_input_tokens"),
		OutputTokens: intField(stats, "output_tokens"),
		CostUSD:      floatField(stats, "cost_usd"),
		ElapsedMS:    intField(stats, "elapsed_ms"),
		TimedOut:     boolField(result, "timed_out"),
	}
	if code, ok := intValue(result["returncode"]); ok {
		c := int(code)
		rec.ReturnCode = &c
	}
	return rec, nil
}

// startsWithDate reports whether s begins with a real calendar date written YYYY-MM-DD.
func startsWithDate(s string) bool {
	if len(s) < len(dateLayout) {
		return false
	}
	_, err := time.Parse(dateLayout, s[:len(dateLayout)])
	return err == nil
}

// object decodes raw as a JSON object. It reports false when raw is absent, null, or not an object.
func object(raw []byte) (map[string]json.RawMessage, bool) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(raw, &m); err != nil || m == nil {
		return nil, false
	}
	return m, true
}

// requiredString returns the string at key. It fails when the value is absent, null, not a
// string, or empty.
func requiredString(obj map[string]json.RawMessage, key string) (string, error) {
	var s string
	if err := json.Unmarshal(obj[key], &s); err != nil || s == "" {
		return "", fmt.Errorf("%s is not a non-empty string", key)
	}
	return s, nil
}

// optionalString returns the string at key, or "" when the value is absent or not a string.
func optionalString(obj map[string]json.RawMessage, key string) string {
	var s string
	if err := json.Unmarshal(obj[key], &s); err != nil {
		return ""
	}
	return s
}

// isNumberToken reports whether raw is written as a JSON number. A string that looks like a
// number is not one, because json.Number would otherwise accept it.
func isNumberToken(raw []byte) bool {
	raw = bytes.TrimSpace(raw)
	return len(raw) > 0 && (raw[0] == '-' || (raw[0] >= '0' && raw[0] <= '9'))
}

// intValue decodes raw as a whole JSON number. It reports false for absent, null, wrongly typed,
// or fractional values.
func intValue(raw []byte) (int64, bool) {
	if !isNumberToken(raw) {
		return 0, false
	}
	var n json.Number
	if err := json.Unmarshal(raw, &n); err != nil {
		return 0, false
	}
	if i, err := n.Int64(); err == nil {
		return i, true
	}
	f, err := n.Float64()
	if err != nil || f != math.Trunc(f) || f < math.MinInt64 || f >= math.MaxInt64 {
		return 0, false
	}
	return int64(f), true
}

// intField returns the whole number at key, or zero when the value is absent or not a whole number.
func intField(obj map[string]json.RawMessage, key string) int64 {
	v, _ := intValue(obj[key])
	return v
}

// floatField returns the number at key, or zero when the value is absent, null, or not a number.
func floatField(obj map[string]json.RawMessage, key string) float64 {
	if !isNumberToken(obj[key]) {
		return 0
	}
	var f float64
	if err := json.Unmarshal(obj[key], &f); err != nil {
		return 0
	}
	return f
}

// boolField returns the boolean at key, or false when the value is absent or not a boolean.
func boolField(obj map[string]json.RawMessage, key string) bool {
	var b bool
	if err := json.Unmarshal(obj[key], &b); err != nil {
		return false
	}
	return b
}
