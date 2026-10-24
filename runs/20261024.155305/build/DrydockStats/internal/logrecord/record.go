// Package logrecord decodes Drydock's JSON Lines execution log into records.
package logrecord

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Record is one validated execution from the log, with numeric defaults applied.
type Record struct {
	ExecutionID  string
	StartedAt    string
	Status       string
	Target       string
	Command      string
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

// Succeeded reports whether the record's status is "succeeded".
func (r Record) Succeeded() bool {
	return r.Status == "succeeded"
}

// StartDate returns the first 10 characters of started_at, the YYYY-MM-DD text used for grouping and filtering.
func (r Record) StartDate() string {
	if len(r.StartedAt) < 10 {
		return r.StartedAt
	}
	return r.StartedAt[:10]
}

// requiredFields are the dotted paths that must each hold a non-empty string.
var requiredFields = []string{"execution_id", "started_at", "status", "job.command_name", "job.target"}

// checkRequired reports the first required field that is absent, null, empty, or not a string.
func checkRequired(obj map[string]any) error {
	for _, path := range requiredFields {
		if s, ok := stringAt(obj, path); !ok || s == "" {
			return fmt.Errorf("missing required field %s", path)
		}
	}
	return nil
}

// parseRecord validates one decoded JSON object and converts it to a Record.
func parseRecord(obj map[string]any) (Record, error) {
	if err := checkRequired(obj); err != nil {
		return Record{}, err
	}

	var r Record
	r.ExecutionID, _ = stringAt(obj, "execution_id")
	r.StartedAt, _ = stringAt(obj, "started_at")
	r.Status, _ = stringAt(obj, "status")
	r.Command, _ = stringAt(obj, "job.command_name")
	r.Target, _ = stringAt(obj, "job.target")
	r.Provider, _ = stringAt(obj, "job.llm")

	// The executed model wins; the requested model is the fallback.
	r.Model, _ = stringAt(obj, "job.model")
	if executed, _ := stringAt(obj, "result.stats.model"); executed != "" {
		r.Model = executed
	}

	if _, present := valueAt(obj, "result.returncode"); present {
		rc, err := intAt(obj, "result.returncode")
		if err != nil {
			return Record{}, err
		}
		r.ReturnCode = &rc
	}

	if v, present := valueAt(obj, "result.timed_out"); present {
		timedOut, isBool := v.(bool)
		if !isBool {
			return Record{}, errors.New("result.timed_out is not a boolean")
		}
		r.TimedOut = timedOut
	}

	counts := []struct {
		path string
		dst  *int64
	}{
		{"result.stats.input_tokens", &r.InputTokens},
		{"result.stats.cached_input_tokens", &r.CachedTokens},
		{"result.stats.output_tokens", &r.OutputTokens},
		{"result.stats.elapsed_ms", &r.ElapsedMS},
	}
	for _, c := range counts {
		v, err := intAt(obj, c.path)
		if err != nil {
			return Record{}, err
		}
		*c.dst = v
	}

	cost, err := floatAt(obj, "result.stats.cost_usd")
	if err != nil {
		return Record{}, err
	}
	r.CostUSD = cost
	return r, nil
}

// decodeObject parses one line as a single JSON object, keeping numbers exact.
func decodeObject(line string) (map[string]any, error) {
	dec := json.NewDecoder(bytes.NewReader([]byte(line)))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	if _, err := dec.Token(); err != io.EOF {
		return nil, errors.New("trailing data after JSON value")
	}
	obj, ok := v.(map[string]any)
	if !ok {
		return nil, errors.New("not a JSON object")
	}
	return obj, nil
}

// valueAt returns the value at a dotted path. Absent and null values report false.
func valueAt(obj map[string]any, path string) (any, bool) {
	var cur any = obj
	for _, key := range strings.Split(path, ".") {
		m, isMap := cur.(map[string]any)
		if !isMap {
			return nil, false
		}
		next, found := m[key]
		if !found {
			return nil, false
		}
		cur = next
	}
	if cur == nil {
		return nil, false
	}
	return cur, true
}

// stringAt returns the string at path and whether one was present.
func stringAt(obj map[string]any, path string) (string, bool) {
	v, ok := valueAt(obj, path)
	if !ok {
		return "", false
	}
	s, isString := v.(string)
	return s, isString
}

// intAt returns the integer at path, or zero when the value is absent or null.
func intAt(obj map[string]any, path string) (int64, error) {
	v, ok := valueAt(obj, path)
	if !ok {
		return 0, nil
	}
	n, isNumber := v.(json.Number)
	if !isNumber {
		return 0, fmt.Errorf("%s is not a number", path)
	}
	i, err := n.Int64()
	if err != nil {
		return 0, fmt.Errorf("%s is not an integer", path)
	}
	return i, nil
}

// floatAt returns the number at path, or zero when the value is absent or null.
func floatAt(obj map[string]any, path string) (float64, error) {
	v, ok := valueAt(obj, path)
	if !ok {
		return 0, nil
	}
	n, isNumber := v.(json.Number)
	if !isNumber {
		return 0, fmt.Errorf("%s is not a number", path)
	}
	f, err := n.Float64()
	if err != nil {
		return 0, fmt.Errorf("%s is not a number", path)
	}
	return f, nil
}
