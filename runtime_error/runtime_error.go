package runtime_error

import (
	"fmt"
	"os"
)

var HadError = false

func ReportAtLine(line int, message string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error :%s\n", line, message)
	HadError = true
}
