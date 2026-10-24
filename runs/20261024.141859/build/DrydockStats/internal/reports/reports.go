package reports

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strconv"

	"drydockstats/internal/logintake"
)

type totals struct {
	calls, failed, input, cached, output, elapsed int64
	cost                                          float64
}

func (t *totals) add(r logintake.Record) {
	t.calls++
	if r.Status != "succeeded" {
		t.failed++
	}
	t.input += r.Result.Stats.InputTokens
	t.cached += r.Result.Stats.CachedInputTokens
	t.output += r.Result.Stats.OutputTokens
	t.elapsed += r.Result.Stats.ElapsedMS
	t.cost += r.Result.Stats.CostUSD
}
func (t totals) minutes() string {
	n := (t.elapsed + 3000) / 6000
	return fmt.Sprintf("%d.%d", n/10, n%10)
}
func num(n int64) string    { return strconv.FormatInt(n, 10) }
func cost(n float64) string { return fmt.Sprintf("%.4f", n) }
func write(w io.Writer, rows [][]string) error {
	c := csv.NewWriter(w)
	c.UseCRLF = false
	for _, r := range rows {
		if err := c.Write(r); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}

func Write(w io.Writer, name string, rs []logintake.Record, skipped int) error {
	base := totals{}
	for _, r := range rs {
		base.add(r)
	}
	if name == "summary" {
		first, last := "", ""
		for _, r := range rs {
			d := r.StartedAt[:10]
			if first == "" || d < first {
				first = d
			}
			if d > last {
				last = d
			}
		}
		succeeded := base.calls - base.failed
		return write(w, [][]string{{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"}, {num(base.calls), strconv.Itoa(skipped), num(succeeded), num(base.failed), first, last, num(base.input), num(base.cached), num(base.output), cost(base.cost), base.minutes()}})
	}
	if name == "failures" {
		failed := make([]logintake.Record, 0)
		for _, r := range rs {
			if r.Status != "succeeded" {
				failed = append(failed, r)
			}
		}
		sort.Slice(failed, func(i, j int) bool {
			if failed[i].StartedAt != failed[j].StartedAt {
				return failed[i].StartedAt < failed[j].StartedAt
			}
			return failed[i].ExecutionID < failed[j].ExecutionID
		})
		rows := [][]string{{"started_at", "target", "command", "model", "returncode", "timed_out"}}
		for _, r := range failed {
			model := r.Result.Stats.Model
			if model == "" {
				model = r.Job.Model
			}
			rc := ""
			if r.Result.Returncode != nil {
				rc = strconv.Itoa(*r.Result.Returncode)
			}
			rows = append(rows, []string{r.StartedAt, r.Job.Target, r.Job.CommandName, model, rc, strconv.FormatBool(r.Result.TimedOut)})
		}
		return write(w, rows)
	}
	groups := map[string]*totals{}
	labels := map[string][]string{}
	for _, r := range rs {
		var key string
		var labs []string
		switch name {
		case "commands":
			key = r.Job.CommandName
			labs = []string{key}
		case "targets":
			key = r.Job.Target
			labs = []string{key}
		case "models":
			m := r.Result.Stats.Model
			if m == "" {
				m = r.Job.Model
			}
			labs = []string{r.Job.LLM, m}
			key = r.Job.LLM + "\x00" + m
		case "daily":
			key = r.StartedAt[:10]
			labs = []string{key}
		default:
			return fmt.Errorf("unknown report %q", name)
		}
		if groups[key] == nil {
			groups[key] = &totals{}
			labels[key] = labs
		}
		groups[key].add(r)
	}
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	if name == "commands" {
		sort.Slice(keys, func(i, j int) bool {
			a, b := groups[keys[i]], groups[keys[j]]
			if a.calls != b.calls {
				return a.calls > b.calls
			}
			return keys[i] < keys[j]
		})
	} else {
		sort.Slice(keys, func(i, j int) bool {
			if name == "models" {
				return keys[i] < keys[j]
			}
			return keys[i] < keys[j]
		})
	}
	header := map[string][]string{"commands": {"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"}, "targets": {"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"}, "models": {"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"}, "daily": {"date", "calls", "failed", "tokens", "minutes"}}[name]
	rows := [][]string{header}
	for _, k := range keys {
		t := groups[k]
		row := append([]string{}, labels[k]...)
		switch name {
		case "commands":
			row = append(row, num(t.calls), num(t.failed), num(t.input), num(t.cached), num(t.output), t.minutes())
		case "targets":
			row = append(row, num(t.calls), num(t.failed), num(t.input), num(t.cached), num(t.output), cost(t.cost), t.minutes())
		case "models":
			row = append(row, num(t.calls), num(t.failed), num(t.input), num(t.cached), num(t.output), cost(t.cost))
		case "daily":
			row = append(row, num(t.calls), num(t.failed), num(t.input+t.cached+t.output), t.minutes())
		}
		rows = append(rows, row)
	}
	row := []string{"TOTAL"}
	switch name {
	case "commands":
		row = append(row, num(base.calls), num(base.failed), num(base.input), num(base.cached), num(base.output), base.minutes())
	case "targets":
		row = append(row, num(base.calls), num(base.failed), num(base.input), num(base.cached), num(base.output), cost(base.cost), base.minutes())
	case "models":
		row = append(row, "", num(base.calls), num(base.failed), num(base.input), num(base.cached), num(base.output), cost(base.cost))
	case "daily":
		row = append(row, num(base.calls), num(base.failed), num(base.input+base.cached+base.output), base.minutes())
	}
	rows = append(rows, row)
	return write(w, rows)
}
