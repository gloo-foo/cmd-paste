package paste_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-paste"
	"github.com/gloo-foo/testable"
)

func ExamplePaste_customDelimiter() {
	// echo -e "one\ntwo\nthree" | paste -s -d, -
	output, _ := testable.Test(command.Paste(command.PasteSerial, command.PasteDelimiter(",")), "one\ntwo\nthree\n")
	fmt.Print(output)
	// Output:
	// one,two,three
}
