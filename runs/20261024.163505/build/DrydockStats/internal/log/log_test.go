package log

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestReadLinesIgnoresBlankLinesAndKeepsLineNumbers(t *testing.T) {
	input := "first\n\n   \n\t\nlast-without-newline"
	lines, err := ReadLines(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("got %d lines, want 2", len(lines))
	}
	if lines[0].Number != 1 || string(lines[0].Bytes) != "first\n" {
		t.Errorf("line 0 = %d %q, want 1 \"first\\n\"", lines[0].Number, lines[0].Bytes)
	}
	if lines[1].Number != 5 || string(lines[1].Bytes) != "last-without-newline" {
		t.Errorf("line 1 = %d %q, want 5 last-without-newline", lines[1].Number, lines[1].Bytes)
	}
}

func TestReadLinesEmptyInput(t *testing.T) {
	lines, err := ReadLines(strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 0 {
		t.Errorf("got %d lines from empty input, want 0", len(lines))
	}
}

func TestReadFileMissingPathIsError(t *testing.T) {
	_, err := ReadFile(filepath.Join(t.TempDir(), "absent.jsonl"))
	if err == nil {
		t.Error("ReadFile on a missing path returned no error")
	}
}

func TestReadFileKeepsMultibyteCharacters(t *testing.T) {
	path := filepath.Join(t.TempDir(), "log.jsonl")
	line := `{"execution_id":"u","job":{"target":"Café"}}`
	if err := writeFile(path, line+"\n"); err != nil {
		t.Fatal(err)
	}
	lines, err := ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 || strings.TrimSpace(string(lines[0].Bytes)) != line {
		t.Errorf("read %v, want one line equal to %q", lines, line)
	}
}
