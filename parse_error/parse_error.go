package parse_error

import (
	"fmt"
	"os"

	"github.com/0xsuk/golox/token"
)

var HadError = false

func Format(line int, where string, message string) string {
	return fmt.Sprintf("[line %v] Error%s: %s", line, where, message)
}

func FormatByToken(tok token.Token, message string) string {
	if tok.Type == token.EOF {
		return Format(tok.Line, " at end", message)
	}
	return Format(tok.Line, " at '"+tok.Lexeme+"'", message)
}

func ReportAtLine(line int, message string) {
	Report(Format(line, "", message))
}

func ReportAtToken(tok token.Token, message string) {
	if tok.Type == token.EOF {
		Report(Format(tok.Line, " at end", message))
	} else {
		Report(FormatByToken(tok, message))
	}
}

func Report(message string) {
	fmt.Fprintln(os.Stderr, message)
	HadError = true
}
