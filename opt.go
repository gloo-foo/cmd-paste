package command

import gloo "github.com/gloo-foo/framework"

// pasteDelimiter is the delimiter list joined between consecutive lines (-d).
// GNU paste treats the argument as a list of single-byte delimiters cycled in
// order between joins; a single-character value behaves as one fixed separator.
type pasteDelimiter string

// PasteDelimiter sets the delimiter list used to join lines (-d). It returns the
// exported gloo.Switch flag interface rather than the unexported concrete type
// so callers always name an exported type.
func PasteDelimiter(d string) gloo.Switch[flags] { return pasteDelimiter(d) }

// pasteSerialFlag selects serial mode (-s): all lines of the stream are pasted
// into a single output row. The default (parallel) mode passes each line
// through as its own row.
type pasteSerialFlag bool

const (
	PasteSerial   pasteSerialFlag = true
	PasteNoSerial pasteSerialFlag = false
)

type flags struct {
	delimiter    pasteDelimiter
	delimiterSet bool
	serial       pasteSerialFlag
}

func (d pasteDelimiter) Configure(flags *flags) {
	flags.delimiter = d
	flags.delimiterSet = true
}

func (s pasteSerialFlag) Configure(flags *flags) { flags.serial = s }
