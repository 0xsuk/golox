package parser

import (
	"github.com/0xsuk/golox/ast"
	"github.com/0xsuk/golox/parse_error"
	"github.com/0xsuk/golox/token"
)

type Parser struct {
	tokens  []token.Token
	current int
	inloop  bool //TODO: ?
}

func New(tokens []token.Token) Parser {
	return Parser{tokens, 0, false}
}

func (p *Parser) Parse() []ast.Stmt {
	statements := make([]ast.Stmt, 0)

	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
}

func (p *Parser) declaration() ast.Stmt {
	var stmt ast.Stmt

	defer func() {
		err := recover()
		if err != nil {
			parse_error.Report(err.(string))
			p.synchronize()
			stmt = nil
		}
	}()

	if p.match(token.CLASS) {
		stmt = nil
	} else if p.match(token.VAR) {
		stmt = nil
	} else if p.match(token.FUN) {
	} else {
		stmt = p.statement()
	}

	return stmt
}

func (p *Parser) statement() ast.Stmt {
	return p.expressionStatement()
}

func (p *Parser) expressionStatement() ast.Stmt {
	expr := p.expression()
	p.consume(token.SEMICOLON, "Expected ';' after value.")

	return &ast.ExpressionStmt{Expression: expr}
}

func (p *Parser) expression() ast.Expr {
	return p.comma()
}

func (p *Parser) comma() ast.Expr {
	expr := p.assignment()

	for p.match(",") {
		operator := p.previous()
		right := p.assignment()
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) assignment() ast.Expr {
	expr := p.or()

	if p.match(token.EQUAL) {
		equals := p.previous()
		value := p.assignment()

		if variable, ok := expr.(*ast.VariableExpr); ok {
			return &ast.AssignExpr{Name: variable.Name, Value: value, EnvIndex: -1, EnvDepth: -1}
		} else if get, ok := expr.(*ast.GetExpr); ok {
			return &ast.SetExpr{Object: get.Object, Name: get.Name, Value: value}
		}

		panic(parse_error.FormatByToken(equals, "Invalid assignment target."))
	}
	return expr
}

func (p *Parser) or() ast.Expr {
	expr := p.and()

	for p.match(token.OR) {
		operator := p.previous()
		right := p.and()
		expr = &ast.LogicalExpr{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) and() ast.Expr {
	expr := p.ternary()
	for p.match(token.AND) {
		operator := p.previous()
		right := p.ternary()
		expr = &ast.LogicalExpr{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) ternary() ast.Expr {
	cond := p.equality()
	if p.match("?") {
		qmark := p.previous()
		thenClause := p.expression()
		p.consume(token.COLON, "Expected ':' in ternary operator.")
		colon := p.previous()
		elseClause := p.expression()
		return &ast.TernaryExpr{Condition: cond, QMark: qmark, Then: thenClause, Colon: colon, Else: elseClause}
	}
	return cond
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()

	for p.match(token.BANGEQUAL, token.EQUALEQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) comparison() ast.Expr {
	expr := p.addition()

	for p.match(token.GREATER, token.GREATEREQUAL, token.LESS, token.LESSEQUAL) {
		operator := p.previous()
		right := p.addition()
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) addition() ast.Expr {
	expr := p.multiplication()

	for p.match(token.PLUS, token.MINUS) {
		operator := p.previous()
		right := p.multiplication()
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) multiplication() ast.Expr {
	expr := p.unary()

	for p.match(token.STAR, token.SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &ast.UnaryExpr{Operator: operator, Right: right}
	}

	return p.power()
}

func (p *Parser) power() ast.Expr {
	expr := p.call()

	for p.match(token.POWER) {
		operator := p.previous()
		right := p.unary()
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) call() ast.Expr {
	expr := p.primary()

	for {
		if p.match(token.LEFTPAREN) {
			expr = p.finishCall(expr)
		} else if p.match(token.DOT) {
			name := p.consume(token.IDENTIFIER, "Expected property name after '.'")
			expr = &ast.GetExpr{Object: expr, Name: name}
		} else {
			break
		}
	}
	return expr
}

func (p *Parser) finishCall(callee ast.Expr) ast.Expr {
	args := make([]ast.Expr, 0)
	if !p.check(token.RIGHTPAREN) {
		for {
			arg := p.assignment() // we don't want the comma operator here
			if len(args) >= 8 {
				parse_error.FormatByToken(p.peek(), "Cannot have more than 8 arguments.")
			}
			args = append(args, arg)
			if !p.match(token.COMMA) {
				break
			}
		}
	}

	paren := p.consume(token.RIGHTPAREN, "Expected ')' after arguments.")
	return &ast.CallExpr{Callee: callee, Paren: paren, Arguments: args}
}

func (p *Parser) primary() ast.Expr {
	if p.match(token.FALSE) {
		return &ast.LiteralExpr{Value: false}
	} else if p.match(token.TRUE) {
		return &ast.LiteralExpr{Value: true}
	} else if p.match(token.NIL) {
		return &ast.LiteralExpr{Value: nil}
	} else if p.match(token.NUMBER, token.STRING) {
		return &ast.LiteralExpr{Value: p.previous().Literal}
	} else if p.match(token.SUPER) {
		keyword := p.previous()
		p.consume(token.DOT, "Expected '.' after 'super'.")
		method := p.consume(token.IDENTIFIER, "Expected superclass method name.")
		return &ast.SuperExpr{Keyword: keyword, Method: method}
	} else if p.match(token.THIS) {
		return &ast.ThisExpr{Keyword: p.previous(), EnvIndex: -1, EnvDepth: -1}
	} else if p.match(token.LEFTPAREN) {
		expr := p.expression()
		p.consume(token.RIGHTPAREN, "Expected ')' after expression.")
		return &ast.GroupingExpr{Expression: expr}
	} else if p.match(token.IDENTIFIER) {
		return &ast.VariableExpr{Name: p.previous(), EnvIndex: -1, EnvDepth: -1}
	}
	panic(parse_error.FormatByToken(p.peek(), "Expected expression."))
}

func (p *Parser) check(tp token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tp
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) consume(tp token.Type, message string) token.Token {
	if p.check(tp) {
		return p.advance()
	}
	panic(parse_error.FormatByToken(p.peek(), message))
}

func (p *Parser) match(types ...token.Type) bool {
	for _, tp := range types {
		if p.check(tp) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case token.CLASS, token.FUN, token.VAR, token.FOR, token.IF, token.WHILE, token.PRINT, token.RETURN:
			return
		}
		p.advance()
	}
}
