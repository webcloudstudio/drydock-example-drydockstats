package cli

import (
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"ddkstat/internal/logread"
)

// failingReader yields one complete record line, then fails the read.
type failingReader struct {
	data []byte
}

func (f *failingReader) Read(p []byte) (int, error) {
	if len(f.data) == 0 {
		return 0, errors.New("disk on fire")
	}
	n := copy(p, f.data)
	f.data = f.data[n:]
	return n, nil
}

func TestReadErrorLeavesStdoutEmpty(t *testing.T) {
	const record = `{"execution_id":"a","started_at":"2026-10-14T00:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Toml"}}` + "\n"
	reader := &failingReader{data: []byte(record)}
	opts := Options{Report: "summary", Log: "partial.jsonl"}

	var stdout, stderr bytes.Buffer
	got := execute(opts, &stdout, &stderr, func(warn io.Writer) (logread.Result, error) {
		return logread.Read(reader, warn)
	})
	if got != 1 {
		t.Errorf("execute = %d, want 1", got)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want nothing", stdout.String())
	}
}

func TestRunMissingLogExitsOneWithEmptyStdout(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-log.jsonl")
	for _, report := range reportNames {
		var stdout, stderr bytes.Buffer
		if got := Run([]string{report, missing}, &stdout, &stderr); got != 1 {
			t.Errorf("Run(%s, missing) = %d, want 1", report, got)
		}
		if stdout.Len() != 0 {
			t.Errorf("Run(%s, missing) wrote %q to stdout", report, stdout.String())
		}
	}
}

func TestRunDirectoryLogExitsOneWithEmptyStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if got := Run([]string{"summary", t.TempDir()}, &stdout, &stderr); got != 1 {
		t.Errorf("Run(summary, directory) = %d, want 1", got)
	}
	if stdout.Len() != 0 {
		t.Errorf("Run(summary, directory) wrote %q to stdout", stdout.String())
	}
}

func TestUsageLineNamesReportsAndOptions(t *testing.T) {
	// Split the usage line on its syntax characters, so each name must appear as its own token.
	tokens := make(map[string]bool)
	for _, tok := range strings.FieldsFunc(usage, func(r rune) bool {
		return strings.ContainsRune(" <>|[]", r)
	}) {
		tokens[tok] = true
	}
	for _, report := range []string{"summary", "commands", "targets", "models", "daily", "failures"} {
		if !tokens[report] {
			t.Errorf("usage line does not name report %q", report)
		}
	}
	for _, option := range []string{"--target", "--since", "--until"} {
		if !tokens[option] {
			t.Errorf("usage line does not name option %q", option)
		}
	}
}
