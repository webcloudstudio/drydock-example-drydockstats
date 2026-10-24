package logintake

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

type Record struct {
	ExecutionID string `json:"execution_id"`
	StartedAt   string `json:"started_at"`
	Status      string `json:"status"`
	Job         struct {
		CommandName string `json:"command_name"`
		Target      string `json:"target"`
		LLM         string `json:"llm"`
		Model       string `json:"model"`
	} `json:"job"`
	Result struct {
		Returncode *int `json:"returncode"`
		TimedOut   bool `json:"timed_out"`
		Stats      struct {
			InputTokens       int64   `json:"input_tokens"`
			CachedInputTokens int64   `json:"cached_input_tokens"`
			OutputTokens      int64   `json:"output_tokens"`
			ElapsedMS         int64   `json:"elapsed_ms"`
			CostUSD           float64 `json:"cost_usd"`
			Model             string  `json:"model"`
		} `json:"stats"`
	} `json:"result"`
}

type Filters struct{ Target, Since, Until string }

func Read(r io.Reader, warn func(line int, err error)) ([]Record, int, error) {
	s := bufio.NewScanner(r)
	s.Buffer(make([]byte, 64*1024), 16*1024*1024)
	seen := map[string]bool{}
	var records []Record
	skipped, line := 0, 0
	for s.Scan() {
		line++
		b := s.Bytes()
		if strings.TrimSpace(string(b)) == "" {
			continue
		}
		var raw map[string]json.RawMessage
		var rec Record
		err := json.Unmarshal(b, &raw)
		if err == nil && raw == nil {
			err = fmt.Errorf("record must be an object")
		}
		if err == nil {
			err = json.Unmarshal(b, &rec)
		}
		if err == nil && (rec.ExecutionID == "" || rec.StartedAt == "" || rec.Status == "" || rec.Job.CommandName == "" || rec.Job.Target == "") {
			err = fmt.Errorf("record is incomplete")
		}
		if err == nil && seen[rec.ExecutionID] {
			err = fmt.Errorf("duplicate execution_id")
		}
		if err != nil {
			skipped++
			if warn != nil {
				warn(line, err)
			}
			continue
		}
		seen[rec.ExecutionID] = true
		records = append(records, rec)
	}
	if err := s.Err(); err != nil {
		return nil, skipped, err
	}
	return records, skipped, nil
}

func Apply(records []Record, f Filters) ([]Record, error) {
	for _, d := range []string{f.Since, f.Until} {
		if d != "" {
			if _, err := time.Parse("2006-01-02", d); err != nil {
				return nil, fmt.Errorf("invalid date %q", d)
			}
		}
	}
	out := make([]Record, 0, len(records))
	for _, r := range records {
		date := r.StartedAt
		if len(date) < 10 {
			continue
		}
		date = date[:10]
		if f.Target != "" && r.Job.Target != f.Target {
			continue
		}
		if f.Since != "" && date < f.Since {
			continue
		}
		if f.Until != "" && date > f.Until {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}
