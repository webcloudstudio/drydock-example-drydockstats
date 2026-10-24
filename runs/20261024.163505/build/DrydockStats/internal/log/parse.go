package log

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

// Record is one kept log record with the fields the reports read.
type Record struct {
	ExecutionID  string
	StartedAt    string
	StartDate    string
	Status       string
	Command      string
	Target       string
	Provider     string
	Model        string
	ReturnCode   *int64
	TimedOut     bool
	InputTokens  int64
	CachedTokens int64
	OutputTokens int64
	CostUSD      float64
	ElapsedMS    int64
}

// Failed reports whether the record's status is anything other than succeeded.
func (r Record) Failed() bool {
	return r.Status != "succeeded"
}

// Skip names a line that Parse did not keep and why.
type Skip struct {
	Line   int
	Reason string
}

// Parsed is the outcome of parsing every line of a log.
type Parsed struct {
	Records []Record
	Skips   []Skip
}

// Parse keeps each well-formed record whose execution_id has not appeared before, and records a
// Skip for every other line. Parse never fails: a bad line is skipped, not fatal.
func Parse(lines []Line) Parsed {
	var p Parsed
	seen := make(map[string]bool)
	for _, ln := range lines {
		rec, reason := decode(ln.Bytes)
		if reason == "" && seen[rec.ExecutionID] {
			reason = "duplicate execution_id " + rec.ExecutionID
		}
		if reason != "" {
			p.Skips = append(p.Skips, Skip{Line: ln.Number, Reason: reason})
			continue
		}
		seen[rec.ExecutionID] = true
		p.Records = append(p.Records, rec)
	}
	return p
}

// decode turns one raw line into a record, or returns the reason the line is skipped.
func decode(raw []byte) (Record, string) {
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var doc map[string]any
	if err := dec.Decode(&doc); err != nil || doc == nil {
		return Record{}, "not a JSON object"
	}
	if err := dec.Decode(new(json.RawMessage)); err != io.EOF {
		return Record{}, "not a JSON object"
	}

	var rec Record
	required := []struct {
		path []string
		dst  *string
	}{
		{[]string{"execution_id"}, &rec.ExecutionID},
		{[]string{"started_at"}, &rec.StartedAt},
		{[]string{"status"}, &rec.Status},
		{[]string{"job", "command_name"}, &rec.Command},
		{[]string{"job", "target"}, &rec.Target},
	}
	for _, f := range required {
		s, ok := at(doc, f.path...).(string)
		if !ok || s == "" {
			return Record{}, "missing or empty " + strings.Join(f.path, ".")
		}
		*f.dst = s
	}

	rec.StartDate = rec.StartedAt[:min(len(rec.StartedAt), 10)]
	rec.Provider = text(doc, "job", "llm")
	rec.Model = modelOf(doc)
	rec.ReturnCode = returnCode(doc)
	rec.TimedOut, _ = at(doc, "result", "timed_out").(bool)
	rec.InputTokens = integer(doc, "result", "stats", "input_tokens")
	rec.CachedTokens = integer(doc, "result", "stats", "cached_input_tokens")
	rec.OutputTokens = integer(doc, "result", "stats", "output_tokens")
	rec.CostUSD = number(doc, "result", "stats", "cost_usd")
	rec.ElapsedMS = integer(doc, "result", "stats", "elapsed_ms")
	return rec, ""
}

// modelOf attributes a record to result.stats.model when it is present and non-empty, otherwise
// to job.model.
func modelOf(doc map[string]any) string {
	if executed := text(doc, "result", "stats", "model"); executed != "" {
		return executed
	}
	return text(doc, "job", "model")
}

// returnCode is nil when result.returncode is absent or null.
func returnCode(doc map[string]any) *int64 {
	n, ok := at(doc, "result", "returncode").(json.Number)
	if !ok {
		return nil
	}
	v, err := n.Int64()
	if err != nil {
		return nil
	}
	return &v
}

// at walks nested JSON objects by key, returning nil when any step is absent or not an object.
func at(v any, path ...string) any {
	for _, key := range path {
		obj, ok := v.(map[string]any)
		if !ok {
			return nil
		}
		v = obj[key]
	}
	return v
}

// text returns the string at path, or "" when it is absent or not a string.
func text(doc map[string]any, path ...string) string {
	s, _ := at(doc, path...).(string)
	return s
}

// integer returns the number at path as an int64, or 0 when it is absent, null, or not a number.
func integer(doc map[string]any, path ...string) int64 {
	n, ok := at(doc, path...).(json.Number)
	if !ok {
		return 0
	}
	if v, err := n.Int64(); err == nil {
		return v
	}
	f, _ := n.Float64()
	return int64(f)
}

// number returns the number at path as a float64, or 0 when it is absent, null, or not a number.
func number(doc map[string]any, path ...string) float64 {
	n, ok := at(doc, path...).(json.Number)
	if !ok {
		return 0
	}
	f, _ := n.Float64()
	return f
}
