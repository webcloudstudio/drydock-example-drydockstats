package reports

import (
	"bytes"
	"drydockstats/internal/logintake"
	"strings"
	"testing"
)

func TestAllReportHeadersAndEmptyTotals(t *testing.T) {
	wants := map[string]string{"summary": "records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes\n0,4,0,0,,,0,0,0,0.0000,0.0\n", "commands": "command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes\nTOTAL,0,0,0,0,0,0.0\n", "targets": "target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes\nTOTAL,0,0,0,0,0,0.0000,0.0\n", "models": "provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd\nTOTAL,,0,0,0,0,0,0.0000\n", "daily": "date,calls,failed,tokens,minutes\nTOTAL,0,0,0,0.0\n", "failures": "started_at,target,command,model,returncode,timed_out\n"}
	for name, want := range wants {
		t.Run(name, func(t *testing.T) {
			var b bytes.Buffer
			if err := Write(&b, name, nil, 4); err != nil {
				t.Fatal(err)
			}
			if b.String() != want {
				t.Fatalf("got %q want %q", b.String(), want)
			}
		})
	}
}

func TestModelsPreferExecutedModelAndSort(t *testing.T) {
	var a, b logintake.Record
	a.Job.LLM = "z"
	a.Job.Model = "requested"
	a.Result.Stats.Model = "executed"
	a.Status = "succeeded"
	b.Job.LLM = "a"
	b.Job.Model = "fallback"
	b.Status = "failed"
	var out bytes.Buffer
	if err := Write(&out, "models", []logintake.Record{a, b}, 0); err != nil {
		t.Fatal(err)
	}
	want := "provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd\na,fallback,1,1,0,0,0,0.0000\nz,executed,1,0,0,0,0,0.0000\nTOTAL,,2,1,0,0,0,0.0000\n"
	if out.String() != want {
		t.Fatalf("got %q", out.String())
	}
}

func TestCSVQuotesNamesAndRoundsMinutes(t *testing.T) {
	var r logintake.Record
	r.Job.CommandName = "build,fast"
	r.Status = "succeeded"
	r.Result.Stats.ElapsedMS = 3000
	var out bytes.Buffer
	if err := Write(&out, "commands", []logintake.Record{r}, 0); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"build,fast",1,0,0,0,0,0.1`) {
		t.Fatalf("%q", out.String())
	}
}
