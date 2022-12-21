package env

import (
	"github.com/0xsuk/golox/runtime_error"
	"github.com/0xsuk/golox/token"
)

type Environment struct {
	values        map[string]interface{}
	enclosing     *Environment
	indexedValues []interface{} //TODO:?
}

type uninitialized struct{} //TODO:?

var needsInitialization = &uninitialized{}

func New(env *Environment) *Environment {
	return NewSized(env, 0)
}

func NewSized(env *Environment, size int) *Environment {
	return &Environment{values: make(map[string]interface{}), enclosing: env, indexedValues: make([]interface{}, size)}
}

func NewGlobal() *Environment {
	return New(nil)
}

func (e *Environment) Define(name string, value interface{}, index int) {
	if index == -1 {
		e.values[name] = value
	} else {
		e.indexedValues[index] = value
	}
}

func (e *Environment) DefineUninitialized(name string, index int) {
	if index == -1 {
		e.values[name] = needsInitialization
	} else {
		e.indexedValues[index] = needsInitialization
	}
}
func (e *Environment) Get(name token.Token, index int) interface{} {
	if index != -1 {
		return e.indexedValues[index]
	}

	v, ok := e.values[name.Lexeme]
	if ok {
		if v == needsInitialization {
			runtime_error.ReportAtLine(name.Line, "Uninitialized variable access: "+name.Lexeme)
			return nil
		}
		return v
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name, index)
	}

	runtime_error.ReportAtLine(name.Line, "Undefined variable '"+name.Lexeme+"'")
	return nil
}

func (e *Environment) GetAt(distance int, name token.Token, index int) interface{} {
	return e.Ancestor(distance).Get(name, index)
}

func (e *Environment) Assign(name token.Token, index int, value interface{}) {
	if index != -1 {
		e.indexedValues[index] = value
		return
	}

	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
		return
	}
	if e.enclosing != nil {
		e.enclosing.Assign(name, index, value)
		return
	}

	runtime_error.ReportAtLine(name.Line, "Undefined variable '"+name.Lexeme+"'")
}

func (e *Environment) AssignAt(distance int, index int, name token.Token, value interface{}) {
	e.Ancestor(distance).Assign(name, index, value)
}

func (e *Environment) Ancestor(distance int) *Environment {
	env := e
	for i := 0; i < distance; i++ {
		env = env.enclosing
	}
	return env
}
