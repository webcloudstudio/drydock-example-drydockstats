package logrecord

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// maxLineBytes bounds one log line; a longer line is a read failure, not a skipped record.
const maxLineBytes = 64 << 20

// Result is everything read from one log.
type Result struct {
	Records  []Record
	Skipped  int
	Warnings []string
}

// ReadFile opens the log at path and reads it with Read.
func ReadFile(path string) (Result, error) {
	f, err := os.Open(path)
	if err != nil {
		return Result{}, err
	}
	defer f.Close()
	return Read(f)
}

// Read decodes JSON Lines from r in file order. Blank lines are ignored. A line that is not a
// valid record is counted in Skipped and described in Warnings, naming its line number. A
// repeated execution_id is skipped after the first occurrence. Read returns an error only when
// r itself cannot be read.
func Read(r io.Reader) (Result, error) {
	var res Result
	seen := make(map[string]bool)

	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), maxLineBytes)
	line := 0
	for sc.Scan() {
		line++
		text := sc.Text()
		if strings.TrimSpace(text) == "" {
			continue
		}

		obj, err := decodeObject(text)
		if err != nil {
			res.skip(line, err)
			continue
		}
		rec, err := parseRecord(obj)
		if err != nil {
			res.skip(line, err)
			continue
		}
		if seen[rec.ExecutionID] {
			res.skip(line, fmt.Errorf("duplicate execution_id %q", rec.ExecutionID))
			continue
		}
		seen[rec.ExecutionID] = true
		res.Records = append(res.Records, rec)
	}
	if err := sc.Err(); err != nil {
		return Result{}, err
	}
	return res, nil
}

// skip counts one skipped line and records its warning.
func (r *Result) skip(line int, err error) {
	r.Skipped++
	r.Warnings = append(r.Warnings, fmt.Sprintf("line %d skipped: %v", line, err))
}
