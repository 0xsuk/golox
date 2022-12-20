package runtimeerror

import (
	"fmt"
	"os"
)

var HadError = false

func ErrorAtLine(line int, message string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error :%s\n", line, message)
	HadError = true
}
