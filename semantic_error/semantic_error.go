package semantic_error

import (
	"fmt"
	"os"
)

var HadError = false

func Print(message string) {
	fmt.Fprintf(os.Stderr, "%v\n", message)
	HadError = true
}

func Make() {}
