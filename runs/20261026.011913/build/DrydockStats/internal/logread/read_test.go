package logread

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestReadIgnoresBlankLines(t *testing.T) {
	input := "\n   \n\t\n" + lineOf(t, baseRecord("a")) + "\n\n  \r\n" + lineOf(t, baseRecord("b")) + "\n\n"
	res, warn := readText(t, input)
	if got := ids(res); !equalStrings(got, []string{"a", "b"}) {
		t.Errorf("kept ids = %v, want [a b]", got)
	}
	if res.Skipped != 0 {
		t.Errorf("Skipped = %d, want 0", res.Skipped)
	}
	if warn != "" {
		t.Errorf("warnings = %q, want none", warn)
	}
}

func TestReadNumbersEveryLine(t *testing.T) {
	// Two blank lines precede the malformed line, so it is physical line 3.
	input := "\n\n{not json\n" + lineOf(t, baseRecord("a")) + "\n"
	res, warn := readText(t, input)
	if got := warnedLines(t, warn); !equalInts(got, []int{3}) {
		t.Errorf("warned lines = %v, want [3]", got)
	}
	if res.Skipped != 1 || len(res.Records) != 1 {
		t.Errorf("Skipped=%d kept=%d, want 1 and 1", res.Skipped, len(res.Records))
	}
}

func TestReadAcceptsLongLine(t *testing.T) {
	rec := baseRecord("long")
	rec["prompt"] = map[string]any{"bytes": 1, "pad": strings.Repeat("x", 2<<20)}
	res, warn := readText(t, lineOf(t, rec)+"\n")
	if len(res.Records) != 1 || res.Records[0].ExecutionID != "long" {
		t.Fatalf("kept records = %+v, want the one long record", res.Records)
	}
	if res.Skipped != 0 || warn != "" {
		t.Errorf("Skipped=%d warn=%q, want 0 and empty", res.Skipped, warn)
	}
}

func TestReadFileMissing(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-log.jsonl")
	if _, err := ReadFile(missing, io.Discard); err == nil {
		t.Error("ReadFile on a missing path returned no error")
	}
}

func TestReadFileDirectory(t *testing.T) {
	if _, err := ReadFile(t.TempDir(), io.Discard); err == nil {
		t.Error("ReadFile on a directory returned no error")
	}
}

func TestReadNeverWritesLog(t *testing.T) {
	path := filepath.Join(t.TempDir(), "llm.jsonl")
	content := []byte(lineOf(t, baseRecord("a")) + "\n{bad\n")
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatal(err)
	}
	past := time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC)
	if err := os.Chtimes(path, past, past); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ReadFile(path, io.Discard)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if len(res.Records) != 1 || res.Skipped != 1 {
		t.Fatalf("kept=%d skipped=%d, want 1 and 1", len(res.Records), res.Skipped)
	}

	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, content) {
		t.Error("log bytes changed after ReadFile")
	}
	if !after.ModTime().Equal(before.ModTime()) {
		t.Errorf("log modification time changed from %v to %v", before.ModTime(), after.ModTime())
	}
}

func TestRecordDateIsStartedAtPrefix(t *testing.T) {
	rec := readOne(t, baseRecord("a"))
	if got := rec.Date(); got != "2026-10-14" {
		t.Errorf("Date() = %q, want 2026-10-14", got)
	}
}

func TestRecordFailedIsStatusNotSucceeded(t *testing.T) {
	ok := readOne(t, baseRecord("a"))
	if ok.Failed() {
		t.Error("succeeded record reports Failed() = true")
	}
	bad := baseRecord("b")
	bad["status"] = "failed"
	if !readOne(t, bad).Failed() {
		t.Error("failed record reports Failed() = false")
	}
}
