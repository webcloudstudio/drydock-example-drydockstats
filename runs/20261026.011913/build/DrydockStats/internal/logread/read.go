package logread

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// Result is the outcome of reading one log.
type Result struct {
	Records []Record // kept records in file order
	Skipped int      // skipped lines in the whole file
}

// Read decodes every non-blank line of r as one Record. A line that breaks a skip rule, or that
// repeats an earlier kept execution_id, is counted in Skipped and reported on warn with its
// 1-based line number. Blank lines are neither records nor skipped. Only a read failure is an error.
func Read(r io.Reader, warn io.Writer) (Result, error) {
	if warn == nil {
		warn = io.Discard
	}
	var res Result
	seen := make(map[string]bool)
	br := bufio.NewReader(r)
	lineNo := 0
	for {
		// ReadBytes has no line length limit, unlike bufio.Scanner.
		line, readErr := br.ReadBytes('\n')
		if len(line) > 0 {
			lineNo++
			if text := bytes.TrimSpace(line); len(text) > 0 {
				rec, decodeErr := decodeRecord(text)
				switch {
				case decodeErr != nil:
					res.skip(warn, lineNo, decodeErr)
				case seen[rec.ExecutionID]:
					res.skip(warn, lineNo, fmt.Errorf("duplicate execution_id %q", rec.ExecutionID))
				default:
					seen[rec.ExecutionID] = true
					res.Records = append(res.Records, rec)
				}
			}
		}
		if readErr == io.EOF {
			return res, nil
		}
		if readErr != nil {
			return Result{}, readErr
		}
	}
}

// ReadFile opens path read-only, reads it with Read, and closes it. It returns an error when the
// path cannot be opened or read, including when it names a directory.
func ReadFile(path string, warn io.Writer) (Result, error) {
	f, err := os.Open(path)
	if err != nil {
		return Result{}, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return Result{}, err
	}
	if info.IsDir() {
		return Result{}, fmt.Errorf("%s is a directory", path)
	}
	return Read(f, warn)
}

// skip counts one skipped line and writes its single warning.
func (res *Result) skip(warn io.Writer, lineNo int, reason error) {
	res.Skipped++
	fmt.Fprintf(warn, "ddkstat: line %d: skipped: %v\n", lineNo, reason)
}
