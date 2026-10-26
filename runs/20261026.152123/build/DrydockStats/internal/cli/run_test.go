package cli_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"ddkstat/internal/cli"
	"ddkstat/internal/logread"
)

// logLine returns one valid JSON Lines record with the given identifier, start date and target.
func logLine(id, date, target string) string {
	return `{"execution_id":"` + id + `","started_at":"` + date + `T02:38:41.904Z","status":"succeeded",` +
		`"job":{"command_name":"build","target":"` + target + `","llm":"claude","model":"claude-haiku-5-5"},` +
		`"result":{"returncode":0,"timed_out":false,"stats":{"input_tokens":1,"cached_input_tokens":2,` +
		`"output_tokens":3,"cost_usd":0.5,"elapsed_ms":4}}}`
}

// writeLog writes lines to a fresh file under t.TempDir and returns its path.
func writeLog(t *testing.T, lines ...string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "log.jsonl")
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		t.Fatalf("write log: %v", err)
	}
	return path
}

// capture registers a writer under report that records what it was given and writes text.
type capture struct {
	records []logread.Record
	skipped int
	calls   int
}

func registerCapture(report string, text string) *capture {
	c := &capture{}
	cli.Register(report, func(w io.Writer, records []logread.Record, skipped int) error {
		c.calls++
		c.records = records
		c.skipped = skipped
		_, err := io.WriteString(w, text)
		return err
	})
	return c
}

func recordIDs(records []logread.Record) []string {
	ids := make([]string, 0, len(records))
	for _, rec := range records {
		ids = append(ids, rec.ExecutionID)
	}
	return ids
}

func TestRunUsageErrorExit2(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if got := cli.Run([]string{"frobnicate", "log.jsonl"}, &stdout, &stderr); got != 2 {
		t.Errorf("Run exit = %d, want 2", got)
	}
}

func TestRunUsageErrorWritesNothingToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	cli.Run([]string{"summary"}, &stdout, &stderr)
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
	if n := strings.Count(stderr.String(), "\n"); n != 1 {
		t.Errorf("stderr has %d lines, want exactly 1", n)
	}
}

func TestRunMissingLogExit1(t *testing.T) {
	var stdout, stderr bytes.Buffer
	missing := filepath.Join(t.TempDir(), "absent.jsonl")
	if got := cli.Run([]string{"summary", missing}, &stdout, &stderr); got != 1 {
		t.Errorf("Run exit = %d, want 1", got)
	}
}

func TestRunMissingLogWritesNothingToStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	missing := filepath.Join(t.TempDir(), "absent.jsonl")
	cli.Run([]string{"summary", missing}, &stdout, &stderr)
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
	if n := strings.Count(stderr.String(), "\n"); n != 1 {
		t.Errorf("stderr has %d lines, want exactly 1", n)
	}
}

func TestRunSkippedLinesStillExit0(t *testing.T) {
	log := writeLog(t, logLine("a", "2026-10-14", "Toml"), "{not json", logLine("b", "2026-10-15", "Toml"))
	c := registerCapture("summary", "ok\n")

	var stdout, stderr bytes.Buffer
	if got := cli.Run([]string{"summary", log}, &stdout, &stderr); got != 0 {
		t.Fatalf("Run exit = %d, want 0", got)
	}
	if c.skipped != 1 {
		t.Errorf("writer skipped = %d, want 1", c.skipped)
	}
	if want := []string{"a", "b"}; !reflect.DeepEqual(recordIDs(c.records), want) {
		t.Errorf("writer records = %v, want %v", recordIDs(c.records), want)
	}
}

func TestRunFilterOrder(t *testing.T) {
	log := writeLog(t,
		logLine("early-toml", "2026-10-13", "Toml"),
		logLine("in-toml", "2026-10-14", "Toml"),
		logLine("in-other", "2026-10-15", "CommonMark"),
		logLine("late-toml", "2026-10-17", "Toml"),
		"{malformed",
		logLine("in-toml-2", "2026-10-16", "Toml"),
	)
	c := registerCapture("summary", "")

	var stdout, stderr bytes.Buffer
	got := cli.Run([]string{"summary", "--target", "Toml", "--since", "2026-10-14", "--until", "2026-10-16", log}, &stdout, &stderr)
	if got != 0 {
		t.Fatalf("Run exit = %d, want 0", got)
	}
	if want := []string{"in-toml", "in-toml-2"}; !reflect.DeepEqual(recordIDs(c.records), want) {
		t.Errorf("writer records = %v, want %v", recordIDs(c.records), want)
	}
	if c.skipped != 1 {
		t.Errorf("writer skipped = %d, want whole-file count 1", c.skipped)
	}
}

func TestRunSkippedCountIgnoresFilters(t *testing.T) {
	log := writeLog(t, "{bad one", logLine("a", "2026-10-14", "Toml"), "{bad two")
	c := registerCapture("summary", "")

	var stdout, stderr bytes.Buffer
	cli.Run([]string{"summary", "--target", "NoSuchTarget", log}, &stdout, &stderr)
	if len(c.records) != 0 {
		t.Errorf("writer records = %v, want none", recordIDs(c.records))
	}
	if c.skipped != 2 {
		t.Errorf("writer skipped = %d, want 2", c.skipped)
	}
}

func TestRunWriterOutputReachesStdoutOnSuccess(t *testing.T) {
	log := writeLog(t, logLine("a", "2026-10-14", "Toml"))
	registerCapture("summary", "report body\n")

	var stdout, stderr bytes.Buffer
	if got := cli.Run([]string{"summary", log}, &stdout, &stderr); got != 0 {
		t.Fatalf("Run exit = %d, want 0", got)
	}
	if stdout.String() != "report body\n" {
		t.Errorf("stdout = %q, want the writer's output", stdout.String())
	}
}

func TestRunWriterErrorExit1WritesNothingToStdout(t *testing.T) {
	log := writeLog(t, logLine("a", "2026-10-14", "Toml"))
	cli.Register("targets", func(w io.Writer, records []logread.Record, skipped int) error {
		io.WriteString(w, "partial output that must not escape\n")
		return errors.New("boom")
	})

	var stdout, stderr bytes.Buffer
	if got := cli.Run([]string{"targets", log}, &stdout, &stderr); got != 1 {
		t.Errorf("Run exit = %d, want 1", got)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout.String())
	}
}
