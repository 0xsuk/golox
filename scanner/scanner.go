package scanner

import (
	"strconv"

	"github.com/0xsuk/golox/parse_error"
	"github.com/0xsuk/golox/token"
)

var keywords = map[string]token.Type{
	"and":      token.AND,
	"class":    token.CLASS,
	"else":     token.ELSE,
	"false":    token.FALSE,
	"for":      token.FOR,
	"fun":      token.FUN,
	"if":       token.IF,
	"nil":      token.NIL,
	"or":       token.OR,
	"print":    token.PRINT,
	"return":   token.RETURN,
	"super":    token.SUPER,
	"this":     token.THIS,
	"true":     token.TRUE,
	"var":      token.VAR,
	"while":    token.WHILE,
	"break":    token.BREAK,
	"continue": token.CONTINUE,
}

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []token.Token
}

func New(source string) Scanner {
	scanner := Scanner{source: source, line: 1, tokens: make([]token.Token, 0)}
	return scanner
}

func (sc *Scanner) ScanTokens() []token.Token {
	for !sc.isAtEnd() {
		sc.start = sc.current
		sc.scanToken()
	}
	sc.tokens = append(sc.tokens, token.Token{Type: token.EOF})
	return sc.tokens
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= (len(sc.source))
}

func (sc *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (sc *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (sc *Scanner) isAlphaNumeric(c byte) bool {
	return sc.isAlpha(c) || sc.isDigit(c)
}

func (sc *Scanner) scanToken() {
	prev := sc.advance()

	switch prev {
	case '(':
		sc.addToken(token.LEFTPAREN)
	case ')':
		sc.addToken(token.RIGHTPAREN)
	case '{':
		sc.addToken(token.LEFTBRACE)
	case '}':
		sc.addToken(token.RIGHTBRACE)
	case ',':
		sc.addToken(token.COMMA)
	case '.':
		sc.addToken(token.DOT)
	case '-':
		sc.addToken(token.MINUS)
	case '+':
		sc.addToken(token.PLUS)
	case '?':
		sc.addToken(token.QMARK)
	case ':':
		sc.addToken(token.COLON)
	case ';':
		sc.addToken(token.SEMICOLON)
	case '*':
		if sc.match('*') {
			sc.addToken(token.POWER)
		} else {
			sc.addToken(token.STAR)
		}
	case '!':
		if sc.match('=') {
			sc.addToken(token.BANGEQUAL)
		} else {
			sc.addToken(token.BANG)
		}
	case '=':
		if sc.match('=') {
			sc.addToken(token.EQUALEQUAL)
		} else {
			sc.addToken(token.EQUAL)
		}
	case '<':
		if sc.match('=') {
			sc.addToken(token.LESSEQUAL)
		} else {
			sc.addToken(token.LESS)
		}
	case '>':
		if sc.match('=') {
			sc.addToken(token.GREATEREQUAL)
		} else {
			sc.addToken(token.GREATER)
		}
	case '/':
		if sc.match('/') {
			// A comment goes until the end of the line
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else {
			sc.addToken(token.SLASH)
		}
	case '\n':
		sc.line++
	case ' ', '\r', '\t':
		// do nothing
	case '"':
		sc.scanString()
	default:
		if sc.isDigit(prev) {
			sc.scanNumber()
		} else if sc.isAlpha(prev) {
			sc.scanIdentifier()
		} else {
			parse_error.ReportAtLine(sc.line, "Unexpected character"+string(prev))
		}
	}
}

func (sc *Scanner) scanString() {
	for sc.peek() == '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
		sc.advance()
	}

	if sc.isAtEnd() {
		parse_error.ReportAtLine(sc.line, "Unterminated string.")
		return
	}

	// the closing "
	sc.advance()

	value := sc.source[sc.start+1 : sc.current-1]
	sc.addTokenWithLiteral(token.STRING, value)
}

func (sc *Scanner) scanNumber() {
	for sc.isDigit(sc.peek()) {
		sc.advance()
	}

	if sc.peek() == '.' && sc.isDigit(sc.peekNext()) {
		sc.advance() //consume "."
		for sc.isDigit(sc.peek()) {
			sc.advance()
		}
	}

	number, err := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)

	if err != nil {
		parse_error.ReportAtLine(sc.line, "Invalid number '"+sc.source[sc.start:sc.current]+"'")
		return
	}

	sc.addTokenWithLiteral(token.NUMBER, number)
}

func (sc *Scanner) scanIdentifier() {
	for sc.isAlphaNumeric(sc.peek()) {
		sc.advance()
	}

	text := sc.source[sc.start:sc.current]
	tp, ok := keywords[text]
	if ok {
		sc.addToken(tp)
	} else {
		sc.addToken(token.IDENTIFIER) //variable("text") name is auto added by addTokenWithLiteral
	}
}

//advance advances current, return previous char
func (sc *Scanner) advance() byte {
	sc.current++
	return sc.source[sc.current-1]
}

func (sc *Scanner) addToken(tp token.Type) {
	sc.addTokenWithLiteral(tp, nil)
}

func (sc *Scanner) addTokenWithLiteral(tp token.Type, literal interface{}) {
	text := sc.source[sc.start:sc.current]
	sc.tokens = append(sc.tokens, token.Token{Type: tp, Lexeme: text, Literal: literal, Line: sc.line})
}

//match returns if current char is expected. If so advances current
func (sc *Scanner) match(expected byte) bool {
	if sc.isAtEnd() {
		return false
	}
	if sc.source[sc.current] != expected {
		return false
	}
	sc.advance()
	return true
}

//peek gets current char
func (sc *Scanner) peek() byte {
	if sc.isAtEnd() {
		return 0
	}
	return sc.source[sc.current]
}

//peeek gets next char
func (sc *Scanner) peekNext() byte {
	if sc.current+1 >= len(sc.source) {
		return 0
	}
	return sc.source[sc.current+1]
}
