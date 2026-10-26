package logread

import "testing"

func TestModelPrefersExecuted(t *testing.T) {
	rec := baseRecord("a")
	stats(rec)["model"] = "claude-sonnet-5-5"
	rec["job"].(map[string]any)["model"] = "sonnet"
	if got := readOne(t, rec).Model; got != "claude-sonnet-5-5" {
		t.Errorf("Model = %q, want claude-sonnet-5-5", got)
	}
}

func TestModelFallsBackWhenStatsAbsent(t *testing.T) {
	rec := baseRecord("a")
	delete(rec["result"].(map[string]any), "stats")
	if got := readOne(t, rec).Model; got != "sonnet" {
		t.Errorf("Model = %q, want sonnet", got)
	}
}

func TestModelFallsBackWhenEmpty(t *testing.T) {
	tests := map[string]func(map[string]any){
		"empty string": func(r map[string]any) { stats(r)["model"] = "" },
		"null":         func(r map[string]any) { stats(r)["model"] = nil },
		"absent":       func(r map[string]any) { delete(stats(r), "model") },
		"not a string": func(r map[string]any) { stats(r)["model"] = 5 },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			rec := baseRecord("a")
			mutate(rec)
			if got := readOne(t, rec).Model; got != "sonnet" {
				t.Errorf("Model = %q, want sonnet", got)
			}
		})
	}
}

func TestProviderIsJobLLM(t *testing.T) {
	rec := baseRecord("a")
	rec["job"].(map[string]any)["llm"] = "codex"
	if got := readOne(t, rec).Provider; got != "codex" {
		t.Errorf("Provider = %q, want codex", got)
	}
}
