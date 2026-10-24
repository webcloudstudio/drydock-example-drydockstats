package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunExitClassesAndOutputContract(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "input.jsonl")
	if err := os.WriteFile(log, []byte(`{"execution_id":"x","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"build","target":"T"}}`+"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name      string
		args      []string
		want      int
		hasOutput bool
	}{{"usage", nil, 2, false}, {"unreadable", []string{"summary", filepath.Join(dir, "missing")}, 1, false}, {"success", []string{"summary", log}, 0, true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out, errout bytes.Buffer
			got := Run(tt.args, &out, &errout)
			if got != tt.want {
				t.Fatalf("exit=%d want=%d", got, tt.want)
			}
			if (out.Len() > 0) != tt.hasOutput {
				t.Fatalf("stdout=%q", out.String())
			}
		})
	}
}

func TestRunRejectsInvalidDateWithoutStdout(t *testing.T) {
	var out, errout bytes.Buffer
	if got := Run([]string{"summary", "--since", "2026-02-30", "log"}, &out, &errout); got != 2 {
		t.Fatalf("exit=%d", got)
	}
	if out.Len() != 0 {
		t.Fatalf("stdout=%q", out.String())
	}
}
