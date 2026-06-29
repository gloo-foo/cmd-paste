package alias_test

import (
	"slices"
	"testing"

	"github.com/gloo-foo/testable"

	paste "github.com/gloo-foo/cmd-paste/alias"
)

// The alias package re-exports the constructor and flag constants under
// unprefixed names. A mis-wired re-export (say, Serial bound to the disabled
// constant, or Delimiter bound to the wrong function) compiles cleanly, so only
// behavior can prove the wiring. Each test exercises one re-export and asserts
// the GNU paste output it must produce.

const serialInput = "alpha\nbeta\ngamma\n"

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_DefaultPassesEachLineThrough(t *testing.T) {
	// No flag: default (parallel) mode keeps each line as its own row.
	lines, err := testable.TestLines(paste.Paste(), serialInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"alpha", "beta", "gamma"})
}

func TestAlias_SerialJoinsWithTab(t *testing.T) {
	// Serial re-exports the -s enabled constant: all lines collapse to one
	// tab-joined row.
	lines, err := testable.TestLines(paste.Paste(paste.Serial), serialInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"alpha\tbeta\tgamma"})
}

func TestAlias_NoSerialMatchesDefault(t *testing.T) {
	// NoSerial is the disabled form of -s: it must behave like passing no flag.
	lines, err := testable.TestLines(paste.Paste(paste.NoSerial), serialInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"alpha", "beta", "gamma"})
}

func TestAlias_DelimiterSetsTheSeparator(t *testing.T) {
	// Delimiter re-exports the -d constructor: serial joins use the given list.
	lines, err := testable.TestLines(paste.Paste(paste.Serial, paste.Delimiter(",")), serialInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"alpha,beta,gamma"})
}
