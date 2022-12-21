package ast

import "github.com/0xsuk/golox/token"

type Stmt interface {
	Accept(visitor StmtVisitor)
}
type StmtVisitor interface {
	visitBlockStmt(stmt *BlockStmt)
	visitClassStmt(stmt *ClassStmt)
	visitExpressionStmt(stmt *ExpressionStmt)
	visitFunctionStmt(stmt *FunctionStmt)
	visitIfStmt(stmt *IfStmt)
	visitPrintStmt(stmt *PrintStmt)
	visitReturnStmt(stmt *ReturnStmt)
	visitContinueStmt(stmt *ContinueStmt)
	visitBreakStmt(stmt *BreakStmt)
	visitVarStmt(stmt *VarStmt)
	visitWhileStmt(stmt *WhileStmt)
}
type BlockStmt struct {
	Stmt
	Statements []Stmt
}

func (stmt *BlockStmt) Accept(visitor StmtVisitor) {
	visitor.visitBlockStmt(stmt)
}

type ClassStmt struct {
	Stmt
	Name       token.Token
	Superclass VariableExpr
	Methods    []FunctionStmt
}

func (stmt *ClassStmt) Accept(visitor StmtVisitor) {
	visitor.visitClassStmt(stmt)
}

type ExpressionStmt struct {
	Stmt
	Expression Expr
}

func (stmt *ExpressionStmt) Accept(visitor StmtVisitor) {
	visitor.visitExpressionStmt(stmt)
}

type FunctionStmt struct {
	Stmt
	Name   token.Token
	Params []token.Token
	Body   []Stmt
}

func (stmt *FunctionStmt) Accept(visitor StmtVisitor) {
	visitor.visitFunctionStmt(stmt)
}

type IfStmt struct {
	Stmt
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func (stmt *IfStmt) Accept(visitor StmtVisitor) {
	visitor.visitIfStmt(stmt)
}

type PrintStmt struct {
	Stmt
	Expression Expr
}

func (stmt *PrintStmt) Accept(visitor StmtVisitor) {
	visitor.visitPrintStmt(stmt)
}

type ReturnStmt struct {
	Stmt
	Keyword token.Token
	Value   Expr
}

func (stmt *ReturnStmt) Accept(visitor StmtVisitor) {
	visitor.visitReturnStmt(stmt)
}

type ContinueStmt struct {
	Stmt
	Token token.Token
}

func (stmt *ContinueStmt) Accept(visitor StmtVisitor) {
	visitor.visitContinueStmt(stmt)
}

type BreakStmt struct {
	Stmt
	Token token.Token
}

func (stmt *BreakStmt) Accept(visitor StmtVisitor) {
	visitor.visitBreakStmt(stmt)
}

type VarStmt struct {
	Stmt
	Name        token.Token
	Initializer Expr
}

func (stmt *VarStmt) Accept(visitor StmtVisitor) {
	visitor.visitVarStmt(stmt)
}

type WhileStmt struct {
	Stmt
	Condition Expr
	Body      Stmt
}

func (stmt *WhileStmt) Accept(visitor StmtVisitor) {
	visitor.visitWhileStmt(stmt)
}
