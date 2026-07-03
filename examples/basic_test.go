package paste_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-paste"
)

func ExamplePaste_basic() {
	// echo -e "alpha\nbeta\ngamma" | paste -s -
	output, _ := testable.Test(command.Paste(command.PasteSerial), "alpha\nbeta\ngamma\n")
	fmt.Print(output)
	// Output:
	// alpha	beta	gamma
}
