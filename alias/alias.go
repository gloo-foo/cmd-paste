// Package alias provides unprefixed type aliases for paste command flags.
// This allows users to import and use shorter names:
//
//	import "github.com/gloo-foo/cmd-paste/alias"
//	paste.Paste(alias.Delimiter(","))
package alias

import command "github.com/gloo-foo/cmd-paste"

// Paste is the command constructor.
var Paste = command.Paste

// Delimiter sets the join delimiter (-d flag).
var Delimiter = command.PasteDelimiter

// -s flag: serial (paste one file at a time)
const Serial = command.PasteSerial

// default: parallel
const NoSerial = command.PasteNoSerial
