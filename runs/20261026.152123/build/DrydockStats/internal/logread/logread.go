// Package logread reads Drydock's JSON Lines LLM execution log into Records.
//
// It never writes to the log it reads. Blank lines are ignored; malformed, incomplete and
// duplicate lines are skipped with a diagnostic on the warning writer.
package logread

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Record is one execution from the log, normalized for reporting.
type Record struct {
	ExecutionID    string
	StartedAt      string
	Date           string
	Status         string
	Failed         bool
	Command        string
	Target         string
	Provider       string
	RequestedModel string
	Model          string
	InputTokens    int64
	CachedTokens   int64
	OutputTokens   int64
	CostUSD        float64
	ElapsedMS      int64
	ReturnCode     *int
	TimedOut       bool
}

// Skip describes one physical line that Read did not keep.
type Skip struct {
	Line   int    // 1-based physical line number, counting blank lines
	Reason string // why the line was skipped
}

// Result is what Read produced from a log: the decoded records and the skipped lines.
type Result struct {
	Records []Record
	Skipped int
	Skips   []Skip
}

// wireLine mirrors one JSON Lines object. The required string fields are typed so that a wrong
// type skips the line. Every other field is kept raw and decoded leniently by parseLine, so a
// value of the wrong JSON type never costs the whole line.
type wireLine struct {
	ExecutionID string `json:"execution_id"`
	StartedAt   string `json:"started_at"`
	Status      string `json:"status"`
	Job         struct {
		CommandName string          `json:"command_name"`
		Target      string          `json:"target"`
		LLM         json.RawMessage `json:"llm"`
		Model       json.RawMessage `json:"model"`
	} `json:"job"`
	Result json.RawMessage `json:"result"`
}

// wireResult mirrors the result object. Each member is raw; see wireLine.
type wireResult struct {
	ReturnCode json.RawMessage `json:"returncode"`
	TimedOut   json.RawMessage `json:"timed_out"`
	Stats      json.RawMessage `json:"stats"`
}

// wireStats mirrors the result.stats object. Each member is raw; see wireLine.
type wireStats struct {
	CachedInputTokens json.RawMessage `json:"cached_input_tokens"`
	ElapsedMS         json.RawMessage `json:"elapsed_ms"`
	InputTokens       json.RawMessage `json:"input_tokens"`
	OutputTokens      json.RawMessage `json:"output_tokens"`
	CostUSD           json.RawMessage `json:"cost_usd"`
	Model             json.RawMessage `json:"model"`
}

// Read decodes every non-blank line of r as one log record. Malformed, incomplete and duplicate
// lines are counted in Result.Skipped, listed in Result.Skips and described on warn. A duplicate
// is a later line whose execution_id an earlier kept line already registered. The error is non-nil
// only when reading from r fails.
func Read(r io.Reader, warn io.Writer) (Result, error) {
	if warn == nil {
		warn = io.Discard
	}
	var res Result
	seen := make(map[string]bool)
	br := bufio.NewReader(r)
	for lineNo := 1; ; lineNo++ {
		line, err := br.ReadBytes('\n')
		if trimmed := bytes.TrimSpace(line); len(trimmed) > 0 {
			rec, perr := parseLine(trimmed)
			if perr == nil && seen[rec.ExecutionID] {
				perr = errors.New("duplicate execution_id")
			}
			if perr != nil {
				res.Skipped++
				res.Skips = append(res.Skips, Skip{Line: lineNo, Reason: perr.Error()})
				fmt.Fprintf(warn, "line %d: skipped: %v\n", lineNo, perr)
			} else {
				seen[rec.ExecutionID] = true
				res.Records = append(res.Records, rec)
			}
		}
		if errors.Is(err, io.EOF) {
			return res, nil
		}
		if err != nil {
			return res, err
		}
	}
}

// parseLine decodes one non-blank, whitespace-trimmed line into a Record. It fails when the line
// is not a JSON object or when a required field is absent, empty or not a string. Every other
// field is optional: an absent, null or wrong-typed value takes its zero value, and a missing
// result or stats object leaves the record with zero numbers.
//
// Model is result.stats.model when that is a non-empty string, otherwise job.model.
func parseLine(data []byte) (Record, error) {
	if data[0] != '{' {
		return Record{}, errors.New("not a JSON object")
	}
	var w wireLine
	if err := json.Unmarshal(data, &w); err != nil {
		return Record{}, err
	}
	required := []struct{ name, value string }{
		{"execution_id", w.ExecutionID},
		{"started_at", w.StartedAt},
		{"status", w.Status},
		{"job.command_name", w.Job.CommandName},
		{"job.target", w.Job.Target},
	}
	for _, field := range required {
		if field.value == "" {
			return Record{}, fmt.Errorf("missing or empty %s", field.name)
		}
	}
	var res wireResult
	decodeObject(w.Result, &res)
	var stats wireStats
	decodeObject(res.Stats, &stats)

	requested := stringValue(w.Job.Model)
	model := stringValue(stats.Model)
	if model == "" {
		model = requested
	}
	return Record{
		ExecutionID:    w.ExecutionID,
		StartedAt:      w.StartedAt,
		Date:           firstRunes(w.StartedAt, 10),
		Status:         w.Status,
		Failed:         w.Status != "succeeded",
		Command:        w.Job.CommandName,
		Target:         w.Job.Target,
		Provider:       stringValue(w.Job.LLM),
		RequestedModel: requested,
		Model:          model,
		InputTokens:    intValue(stats.InputTokens),
		CachedTokens:   intValue(stats.CachedInputTokens),
		OutputTokens:   intValue(stats.OutputTokens),
		CostUSD:        floatValue(stats.CostUSD),
		ElapsedMS:      intValue(stats.ElapsedMS),
		ReturnCode:     returnCode(res.ReturnCode),
		TimedOut:       boolValue(res.TimedOut),
	}, nil
}

// decodeObject decodes raw into v when raw holds a JSON object. Absent, null and non-object values
// leave v at its zero value.
func decodeObject(raw json.RawMessage, v any) {
	if len(raw) > 0 {
		_ = json.Unmarshal(raw, v)
	}
}

// intValue returns raw as an int64, or zero when raw is absent, null or not an integer.
func intValue(raw json.RawMessage) int64 {
	var n int64
	_ = json.Unmarshal(raw, &n)
	return n
}

// floatValue returns raw as a float64, or zero when raw is absent, null or not a number.
func floatValue(raw json.RawMessage) float64 {
	var f float64
	_ = json.Unmarshal(raw, &f)
	return f
}

// stringValue returns raw as a string, or the empty string when raw is absent, null or not a string.
func stringValue(raw json.RawMessage) string {
	var s string
	_ = json.Unmarshal(raw, &s)
	return s
}

// boolValue returns raw as a bool, or false when raw is absent, null or not a bool.
func boolValue(raw json.RawMessage) bool {
	var b bool
	_ = json.Unmarshal(raw, &b)
	return b
}

// returnCode returns raw as a return code. It is nil when raw is absent, null or not an integer.
func returnCode(raw json.RawMessage) *int {
	var n int
	if json.Unmarshal(raw, &n) != nil || string(raw) == "null" {
		return nil
	}
	return &n
}

// firstRunes returns the first n characters of s, or all of s when it is shorter.
func firstRunes(s string, n int) string {
	r := []rune(s)
	if len(r) > n {
		r = r[:n]
	}
	return string(r)
}
