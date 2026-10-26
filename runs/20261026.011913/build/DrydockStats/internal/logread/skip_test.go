package logread

import (
	"bytes"
	"testing"
)

func TestSkipNotJSONObject(t *testing.T) {
	tests := []struct {
		name string
		line string
	}{
		{name: "truncated", line: `{"execution_id":`},
		{name: "array", line: `[1,2,3]`},
		{name: "number", line: `42`},
		{name: "string", line: `"text"`},
		{name: "null", line: `null`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := readText(t, tt.line+"\n")
			if len(res.Records) != 0 || res.Skipped != 1 {
				t.Errorf("kept=%d skipped=%d, want 0 and 1", len(res.Records), res.Skipped)
			}
		})
	}
}

func TestSkipInvalidUTF8(t *testing.T) {
	// The id carries a lone 0xff byte. json.Marshal would replace it, so splice the raw bytes in.
	line := bytes.Replace([]byte(lineOf(t, baseRecord("bad-utf8"))), []byte("bad-utf8"), []byte{'x', 0xff, 'y'}, 1)
	res, _ := readText(t, string(line)+"\n")
	if len(res.Records) != 0 || res.Skipped != 1 {
		t.Errorf("kept=%d skipped=%d, want 0 and 1", len(res.Records), res.Skipped)
	}
}

func TestSkipMissingRequiredField(t *testing.T) {
	// Each mutation breaks one required field: absent, empty, a non-string, or null.
	mutations := map[string]func(map[string]any){
		"execution_id absent": func(r map[string]any) { delete(r, "execution_id") },
		"execution_id empty":  func(r map[string]any) { r["execution_id"] = "" },
		"execution_id number": func(r map[string]any) { r["execution_id"] = 7 },
		"started_at absent":   func(r map[string]any) { delete(r, "started_at") },
		"started_at null":     func(r map[string]any) { r["started_at"] = nil },
		"status absent":       func(r map[string]any) { delete(r, "status") },
		"status boolean":      func(r map[string]any) { r["status"] = true },
		"command_name absent": func(r map[string]any) { delete(r["job"].(map[string]any), "command_name") },
		"command_name empty":  func(r map[string]any) { r["job"].(map[string]any)["command_name"] = "" },
		"target absent":       func(r map[string]any) { delete(r["job"].(map[string]any), "target") },
		"target null":         func(r map[string]any) { r["job"].(map[string]any)["target"] = nil },
		"job absent":          func(r map[string]any) { delete(r, "job") },
		"job not an object":   func(r map[string]any) { r["job"] = "build" },
	}
	for name, mutate := range mutations {
		t.Run(name, func(t *testing.T) {
			rec := baseRecord("a")
			mutate(rec)
			res, _ := readText(t, lineOf(t, rec)+"\n")
			if len(res.Records) != 0 || res.Skipped != 1 {
				t.Errorf("kept=%d skipped=%d, want 0 and 1", len(res.Records), res.Skipped)
			}
		})
	}
}

func TestSkipBadStartedAt(t *testing.T) {
	for _, startedAt := range []string{
		"2026-13-01T00:00:00.000Z",
		"2026-02-30T00:00:00.000Z",
		"2025-02-29T00:00:00.000Z",
		"2026-1-5T00:00:00.000Z",
		"20261014T000000Z",
		"yesterday",
	} {
		t.Run(startedAt, func(t *testing.T) {
			rec := baseRecord("a")
			rec["started_at"] = startedAt
			res, _ := readText(t, lineOf(t, rec)+"\n")
			if len(res.Records) != 0 || res.Skipped != 1 {
				t.Errorf("kept=%d skipped=%d, want 0 and 1", len(res.Records), res.Skipped)
			}
		})
	}
}

func TestSkipWarnsWithLineNumber(t *testing.T) {
	input := lineOf(t, baseRecord("a")) + "\n" +
		"{bad\n" +
		lineOf(t, baseRecord("b")) + "\n" +
		"\n" +
		"[]\n"
	res, warn := readText(t, input)
	if got := warnedLines(t, warn); !equalInts(got, []int{2, 5}) {
		t.Errorf("warned lines = %v, want [2 5]", got)
	}
	if res.Skipped != 2 {
		t.Errorf("Skipped = %d, want 2", res.Skipped)
	}
}

func TestSkipDoesNotAffectValidRecords(t *testing.T) {
	input := lineOf(t, baseRecord("a")) + "\n" +
		"{bad\n" +
		lineOf(t, baseRecord("b")) + "\n" +
		"null\n" +
		lineOf(t, baseRecord("c")) + "\n"
	res, _ := readText(t, input)
	if got := ids(res); !equalStrings(got, []string{"a", "b", "c"}) {
		t.Errorf("kept ids = %v, want [a b c]", got)
	}
	if res.Skipped != 2 {
		t.Errorf("Skipped = %d, want 2", res.Skipped)
	}
}
