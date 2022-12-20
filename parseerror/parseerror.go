package parseerror

import (
	"fmt"
	"os"

	"github.com/0xsuk/golox/token"
)

var HadError = false

func ErrorAtLine(line int, message string) {
	report(line, "", message)
}

func ErrorAtToken(tok token.Token, message string) {
	if tok.Type == token.EOF {
		report(tok.Line, " at end", message)
	} else {
		report(tok.Line, " at "+tok.Lexeme, message)
	}
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error '%s': %s", line, where, message)
	HadError = true
}
