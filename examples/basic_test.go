package paste_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-paste"
	"github.com/gloo-foo/testable"
)

func ExamplePaste_basic() {
	// echo -e "alpha\nbeta\ngamma" | paste -s -
	output, _ := testable.Test(command.Paste(command.PasteSerial), "alpha\nbeta\ngamma\n")
	fmt.Print(output)
	// Output:
	// alpha	beta	gamma
}
