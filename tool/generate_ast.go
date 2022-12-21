package main

import (
	"os"
	"strings"
)

func main() {
	exprNodes := []string{
		"Assign   : Name token.Token, Value Expr, EnvIndex int, EnvDepth int",
		"Binary   : Left Expr, Operator token.Token, Right Expr",
		"Ternary  : Condition Expr, QMark token.Token, Then Expr, Colon token.Token, Else Expr",
		"Call     : Callee Expr, Paren token.Token, Arguments []Expr",
		"Get      : Object Expr, Name token.Token",
		"Grouping : Expression Expr",
		"Literal  : Value interface{}",
		"Logical  : Left Expr, Operator token.Token, Right Expr",
		"Set      : Object Expr, Name token.Token, Value Expr",
		"Super    : Keyword token.Token, Method token.Token",
		"This     : Keyword token.Token, EnvIndex int, EnvDepth int",
		"Unary    : Operator token.Token, Right Expr",
		"Variable : Name token.Token, EnvIndex int, EnvDepth int",
	}

	defineAst("ast/expr.go", "Expr", exprNodes)

	stmtNodes := []string{
		"Block      : Statements []Stmt",
		"Class      : Name token.Token, Superclass VariableExpr, Methods []FunctionStmt",
		"Expression : Expression Expr",
		"Function   : Name token.Token, Params []token.Token, Body []Stmt",
		"If         : Condition Expr, ThenBranch Stmt, ElseBranch Stmt",
		"Print      : Expression Expr",
		"Return     : Keyword token.Token, Value Expr",
		"Continue   : Token token.Token",
		"Break      : Token token.Token",
		"Var        : Name token.Token, Initializer Expr",
		"While      : Condition Expr, Body Stmt",
	}

	defineAst("ast/stmt.go", "Stmt", stmtNodes)
}

func defineAst(path string, basename string, types []string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("package ast\n")
	f.WriteString("import \"github.com/0xsuk/golox/token\"\n")

	f.WriteString("type " + basename + " interface {\n")
	f.WriteString("Accept(visitor " + basename + "Visitor)\n")
	f.WriteString("}\n")
	defineVisitor(f, basename, types)

	for _, tipe := range types {
		typeName := strings.Trim(strings.Split(tipe, ":")[0], " ")
		args := strings.Split(strings.Trim(strings.Split(tipe, ":")[1], " "), ", ") //list of [type name]
		defineType(f, basename, typeName, args)
	}
}

func defineVisitor(f *os.File, basename string, types []string) {
	f.WriteString("type " + basename + "Visitor interface {\n")

	for _, tipe := range types {
		typeName := strings.Split(tipe, " ")[0]
		f.WriteString("visit" + typeName + basename + "(" + strings.ToLower(basename) + " *" + typeName + basename + ")\n")
	}

	f.WriteString("}\n")
}

func defineType(f *os.File, basename string, typeName string, args []string) {

	f.WriteString("type " + typeName + basename + " struct {\n")
	f.WriteString(basename + "\n")

	for _, arg := range args {
		name := strings.Split(arg, " ")[0]
		tipe := strings.Split(arg, " ")[1]
		f.WriteString(name + " " + tipe + "\n")
	}

	f.WriteString("}\n")

	f.WriteString("func (" + strings.ToLower(basename) + " *" + typeName + basename + ") Accept(visitor " + basename + "Visitor) {\n")
	f.WriteString("visitor.visit" + typeName + basename + "(" + strings.ToLower(basename) + ")")
	f.WriteString("}\n")
}
