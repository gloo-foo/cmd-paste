// Package alias provides unprefixed names for paste command flags.
//
//	import paste "github.com/gloo-foo/cmd-paste/alias"
//	paste.Paste(paste.Serial, paste.Delimiter(","))
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-paste"
)

// Paste merges input lines according to GNU paste semantics; see the command
// package for the flag set.
func Paste(opts ...any) gloo.Command[[]byte, []byte] { return command.Paste(opts...) }

// Delimiter is the -d delimiter-list flag joined between consecutive lines.
type Delimiter = command.PasteDelimiter

// -s flag: serial (paste one file at a time)
const Serial = command.PasteSerial

// default: parallel
const NoSerial = command.PasteNoSerial
