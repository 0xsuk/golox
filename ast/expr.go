package ast

import "github.com/0xsuk/golox/token"

type Expr interface {
	Accept(visitor ExprVisitor)
}
type ExprVisitor interface {
	visitAssignExpr(expr *AssignExpr)
	visitBinaryExpr(expr *BinaryExpr)
	visitTernaryExpr(expr *TernaryExpr)
	visitCallExpr(expr *CallExpr)
	visitGetExpr(expr *GetExpr)
	visitGroupingExpr(expr *GroupingExpr)
	visitLiteralExpr(expr *LiteralExpr)
	visitLogicalExpr(expr *LogicalExpr)
	visitSetExpr(expr *SetExpr)
	visitSuperExpr(expr *SuperExpr)
	visitThisExpr(expr *ThisExpr)
	visitUnaryExpr(expr *UnaryExpr)
	visitVariableExpr(expr *VariableExpr)
}
type AssignExpr struct {
	Expr
	Name     token.Token
	Value    Expr
	EnvIndex int
	EnvDepth int
}

func (expr *AssignExpr) Accept(visitor ExprVisitor) {
	visitor.visitAssignExpr(expr)
}

type BinaryExpr struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr *BinaryExpr) Accept(visitor ExprVisitor) {
	visitor.visitBinaryExpr(expr)
}

type TernaryExpr struct {
	Expr
	Condition Expr
	QMark     token.Token
	Then      Expr
	Colon     token.Token
	Else      Expr
}

func (expr *TernaryExpr) Accept(visitor ExprVisitor) {
	visitor.visitTernaryExpr(expr)
}

type CallExpr struct {
	Expr
	Callee    Expr
	Paren     token.Token
	Arguments []Expr
}

func (expr *CallExpr) Accept(visitor ExprVisitor) {
	visitor.visitCallExpr(expr)
}

type GetExpr struct {
	Expr
	Object Expr
	Name   token.Token
}

func (expr *GetExpr) Accept(visitor ExprVisitor) {
	visitor.visitGetExpr(expr)
}

type GroupingExpr struct {
	Expr
	Expression Expr
}

func (expr *GroupingExpr) Accept(visitor ExprVisitor) {
	visitor.visitGroupingExpr(expr)
}

type LiteralExpr struct {
	Expr
	Value interface{}
}

func (expr *LiteralExpr) Accept(visitor ExprVisitor) {
	visitor.visitLiteralExpr(expr)
}

type LogicalExpr struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr *LogicalExpr) Accept(visitor ExprVisitor) {
	visitor.visitLogicalExpr(expr)
}

type SetExpr struct {
	Expr
	Object Expr
	Name   token.Token
	Value  Expr
}

func (expr *SetExpr) Accept(visitor ExprVisitor) {
	visitor.visitSetExpr(expr)
}

type SuperExpr struct {
	Expr
	Keyword token.Token
	Method  token.Token
}

func (expr *SuperExpr) Accept(visitor ExprVisitor) {
	visitor.visitSuperExpr(expr)
}

type ThisExpr struct {
	Expr
	Keyword  token.Token
	EnvIndex int
	EnvDepth int
}

func (expr *ThisExpr) Accept(visitor ExprVisitor) {
	visitor.visitThisExpr(expr)
}

type UnaryExpr struct {
	Expr
	Operator token.Token
	Right    Expr
}

func (expr *UnaryExpr) Accept(visitor ExprVisitor) {
	visitor.visitUnaryExpr(expr)
}

type VariableExpr struct {
	Expr
	Name     token.Token
	EnvIndex int
	EnvDepth int
}

func (expr *VariableExpr) Accept(visitor ExprVisitor) {
	visitor.visitVariableExpr(expr)
}
