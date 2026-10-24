package logintake

import (
	"strings"
	"testing"
)

func TestReadSkipsInvalidAndDuplicateAndRetainsFirst(t *testing.T) {
	in := `{"execution_id":"x","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"build","target":"T"}}
not-json
{"execution_id":"x","started_at":"2026-10-15T00:00:00Z","status":"failed","job":{"command_name":"build","target":"T"}}
`
	warnings := 0
	rs, n, err := Read(strings.NewReader(in), func(int, error) { warnings++ })
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 || warnings != 2 || len(rs) != 1 || rs[0].StartedAt != "2026-10-14T00:00:00Z" {
		t.Fatalf("records=%+v skipped=%d warnings=%d", rs, n, warnings)
	}
}
func TestReadIgnoresBlankAndRejectsIncomplete(t *testing.T) {
	rs, n, err := Read(strings.NewReader("\n{}\n"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(rs) != 0 || n != 1 {
		t.Fatalf("%d %d", len(rs), n)
	}
}
func TestApplyExactTargetAndInclusiveDates(t *testing.T) {
	mk := func(date, target string) Record {
		var r Record
		r.StartedAt = date + "T00:00:00Z"
		r.Job.Target = target
		return r
	}
	got, err := Apply([]Record{mk("2026-10-14", "T"), mk("2026-10-15", "T"), mk("2026-10-15", "Other")}, Filters{Target: "T", Since: "2026-10-15", Until: "2026-10-15"})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Job.Target != "T" {
		t.Fatalf("%+v", got)
	}
}
func TestApplyRejectsBadDate(t *testing.T) {
	if _, err := Apply(nil, Filters{Since: "2026-13-01"}); err == nil {
		t.Fatal("expected error")
	}
}
