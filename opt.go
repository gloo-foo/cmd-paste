package command

// PasteDelimiter is the delimiter list joined between consecutive lines (-d).
// GNU paste treats the argument as a list of single-byte delimiters cycled in
// order between joins; a single-character value behaves as one fixed separator.
type PasteDelimiter string

// pasteSerialFlag selects serial mode (-s): all lines of the stream are pasted
// into a single output row. The default (parallel) mode passes each line
// through as its own row.
type pasteSerialFlag bool

const (
	PasteSerial   pasteSerialFlag = true
	PasteNoSerial pasteSerialFlag = false
)

// flags is the option set folded from a Paste call's option values. The
// hasDelimiter marker keeps an explicit empty -d list distinguishable from
// "not set" (which defaults to tab).
type flags struct {
	delimiter    PasteDelimiter
	hasDelimiter bool
	isSerial     pasteSerialFlag
}

// with folds one option value into the flag set. Values of any other type are
// ignored: paste operates on the single input stream and takes no positional
// arguments.
func (f flags) with(o any) flags {
	switch v := o.(type) {
	case PasteDelimiter:
		f.delimiter = v
		f.hasDelimiter = true
	case pasteSerialFlag:
		f.isSerial = v
	}
	return f
}

// fold collapses the Paste option values into the flag set.
func fold(opts []any) flags {
	var f flags
	for _, o := range opts {
		f = f.with(o)
	}
	return f
}
