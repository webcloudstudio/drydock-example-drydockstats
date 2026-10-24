// Package log reads Drydock's JSON Lines execution log and parses its lines into records.
package log

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Line is one non-blank line of the log with its one-based line number.
type Line struct {
	Number int
	Bytes  []byte
}

// ReadFile reads the log at path as UTF-8 lines, ignoring blank lines.
func ReadFile(path string) ([]Line, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadLines(f)
}

// ReadLines reads r one line at a time and returns every non-blank line in file order.
func ReadLines(r io.Reader) ([]Line, error) {
	br := bufio.NewReader(r)
	var lines []Line
	for n := 1; ; n++ {
		raw, err := br.ReadBytes('\n')
		if len(bytes.TrimSpace(raw)) > 0 {
			lines = append(lines, Line{Number: n, Bytes: raw})
		}
		if err == io.EOF {
			return lines, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
