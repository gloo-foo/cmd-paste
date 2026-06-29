package command_test

import (
	"fmt"
	"testing"

	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/assertion"

	command "github.com/gloo-foo/cmd-paste"
)

// Default (parallel) mode on a single stream passes each line through as its
// own row — GNU `paste -` with one stream is a no-op on the line structure.
func TestPaste_DefaultParallelPassesEachLineThrough(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(), "alpha\nbeta\ngamma\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"alpha", "beta", "gamma"})
}

func TestPaste_DefaultParallelExplicitNoSerial(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(command.PasteNoSerial), "a\nb\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"a", "b"})
}

// Serial mode (-s) collapses the whole stream into one tab-joined row.
func TestPaste_SerialDefaultTab(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(command.PasteSerial), "alpha\nbeta\ngamma\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"alpha\tbeta\tgamma"})
}

func TestPaste_SerialCustomDelimiter(t *testing.T) {
	lines, err := testable.TestLines(
		command.Paste(command.PasteSerial, command.PasteDelimiter(",")),
		"one\ntwo\nthree\n",
	)
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"one,two,three"})
}

// A multi-character -d list is cycled byte by byte between joins: ",;" yields
// "a,b;c,d;e" for five lines.
func TestPaste_SerialCyclesDelimiterList(t *testing.T) {
	lines, err := testable.TestLines(
		command.Paste(command.PasteSerial, command.PasteDelimiter(",;")),
		"a\nb\nc\nd\ne\n",
	)
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"a,b;c,d;e"})
}

// An empty -d list joins with no separator at all.
func TestPaste_SerialEmptyDelimiterJoinsWithNothing(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(command.PasteSerial, command.PasteDelimiter("")), "a\nb\nc\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"abc"})
}

// Empty input produces no output row in either mode — paste emits nothing for
// an empty stream.
func TestPaste_SerialEmptyInputProducesNoRow(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(command.PasteSerial), "")
	assertion.NoError(t, err)
	assertion.Empty(t, lines)
}

func TestPaste_DefaultEmptyInputProducesNoRow(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(), "")
	assertion.NoError(t, err)
	assertion.Empty(t, lines)
}

// A single line in serial mode is emitted unchanged with no trailing delimiter.
func TestPaste_SerialSingleLine(t *testing.T) {
	lines, err := testable.TestLines(command.Paste(command.PasteSerial), "only\n")
	assertion.NoError(t, err)
	assertion.Lines(t, lines, []string{"only"})
}

func TestPaste_SerialTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		opts     []any
		input    string
		expected []string
	}{
		{"tab two lines", []any{command.PasteSerial}, "a\nb\n", []string{"a\tb"}},
		{"comma three lines", []any{command.PasteSerial, command.PasteDelimiter(",")}, "x\ny\nz\n", []string{"x,y,z"}},
		{
			"space delimiter",
			[]any{command.PasteSerial, command.PasteDelimiter(" ")},
			"hello\nworld\n",
			[]string{"hello world"},
		},
		{"single line no join", []any{command.PasteSerial}, "solo\n", []string{"solo"}},
		{"empty delimiter", []any{command.PasteSerial, command.PasteDelimiter("")}, "a\nb\nc\n", []string{"abc"}},
		{"parallel passthrough", nil, "a\nb\nc\n", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := testable.TestLines(command.Paste(tt.opts...), tt.input)
			assertion.NoError(t, err)
			assertion.Lines(t, lines, tt.expected)
		})
	}
}

func ExamplePaste() {
	// Default (parallel) mode: each line of a single stream is its own row.
	output, _ := testable.Test(command.Paste(), "alpha\nbeta\ngamma\n")
	fmt.Print(output)
	// Output:
	// alpha
	// beta
	// gamma
}

func ExamplePaste_serial() {
	// Serial mode (-s) joins every line into one tab-separated row.
	output, _ := testable.Test(command.Paste(command.PasteSerial), "alpha\nbeta\ngamma\n")
	fmt.Print(output)
	// Output:
	// alpha	beta	gamma
}

func ExamplePaste_serialDelimiter() {
	// Serial mode with a delimiter list cycled between joins.
	output, _ := testable.Test(command.Paste(command.PasteSerial, command.PasteDelimiter(",;")), "a\nb\nc\nd\n")
	fmt.Print(output)
	// Output:
	// a,b;c,d
}
