package command

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// tabDelimiter is the default separator GNU paste places between joined lines.
const tabDelimiter = "\t"

// Paste merges input lines according to GNU paste semantics.
//
// Flags:
//   - PasteSerial (-s): join every line of the stream into a single row.
//   - PasteDelimiter (-d): set the delimiter list cycled between joins.
//
// In serial mode the whole stream collapses to one row whose lines are joined
// by the delimiter list (cycled byte by byte). In the default (parallel) mode a
// single input stream passes through unchanged: each line is its own row.
func Paste(opts ...any) gloo.Command[[]byte, []byte] {
	f := gloo.NewParameters[gloo.File, flags](opts...).Flags
	join := joiner(delimiterFor(f))
	return patterns.Accumulate(transform(bool(f.serial), join))
}

// delimiterFor resolves the active delimiter list: the configured value when
// -d was supplied, otherwise the default tab.
func delimiterFor(f flags) string {
	if f.delimiterSet {
		return string(f.delimiter)
	}
	return tabDelimiter
}

// transform selects the row-building strategy for the active mode. Serial mode
// reduces the whole stream to a single joined row (none for empty input);
// parallel mode returns the lines untouched.
func transform(serial bool, join func([][]byte) []byte) func([][]byte) ([][]byte, error) {
	if !serial {
		return passthrough
	}
	return func(lines [][]byte) ([][]byte, error) {
		if len(lines) == 0 {
			return lines, nil
		}
		return [][]byte{join(lines)}, nil
	}
}

// passthrough emits each input line as its own row.
func passthrough(lines [][]byte) ([][]byte, error) { return lines, nil }

// joiner builds a row from lines, inserting the next delimiter byte from the
// cycled list before every line after the first. An empty list joins with no
// separator.
func joiner(delimiter string) func([][]byte) []byte {
	return func(lines [][]byte) []byte {
		var row []byte
		for i, line := range lines {
			row = append(row, separator(delimiter, i)...)
			row = append(row, line...)
		}
		return row
	}
}

// separator returns the delimiter byte preceding the line at index i: nothing
// before the first line or when the list is empty, otherwise the byte at the
// cycled position.
func separator(delimiter string, i int) []byte {
	if i == 0 || delimiter == "" {
		return nil
	}
	return []byte{delimiter[(i-1)%len(delimiter)]}
}
