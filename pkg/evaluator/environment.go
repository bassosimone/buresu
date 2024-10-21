// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"fmt"
	"io"
)

const (
	// environmentFlagScopeFunc indicates that the scope is a function scope.
	environmentFlagScopeFunc = 1 << iota
)

// Environment is the environment used by the evaluator.
//
// Use [NewEnvironment] to construct.
type Environment struct {
	// flags contains flags describing this environment.
	flags int

	// output is the writer used to write output.
	output io.Writer

	// parent is a pointer to the parent environment.
	//
	// The root environment has a nil parent.
	parent *Environment

	// symbols contains the symbols defined in the current environment.
	symbols map[string]Value
}

// NewEnvironment creates a new [*Environment] instance.
func NewEnvironment(output io.Writer) *Environment {
	return &Environment{
		flags:   0,
		output:  output,
		parent:  nil,
		symbols: make(map[string]Value),
	}
}

// isInsideFunc returns true if the environment is a function scope
// or any of the parent environments is a function scope.
func (env *Environment) IsInsideFunc() bool {
	if env.flags&environmentFlagScopeFunc != 0 {
		return true
	}
	if env.parent != nil {
		return env.parent.IsInsideFunc()
	}
	return false
}

// pushFunctionScope creates a new child environment associated with
// a function execution and returns it. The returned environment will
// use the current environment as its parent.
func (env *Environment) pushFunctionScope() *Environment {
	return env.pushScope(environmentFlagScopeFunc)
}

// pushBlockScope creates a new child environment associated with
// a block execution and returns it. The returned environment will
// use the current environment as its parent.
func (env *Environment) pushBlockScope() *Environment {
	return env.pushScope(0)
}

// pushScope creates a new child environment with the given flags and returns it.
func (env *Environment) pushScope(flags int) *Environment {
	return &Environment{
		flags:   flags,
		output:  env.output,
		parent:  env,
		symbols: make(map[string]Value),
	}
}

// GetValue returns the value associated with the given symbol.
//
// If the symbol is not found in the current environment, the parent
// environments are searched recursively.
func (env *Environment) GetValue(symbol string) (Value, bool) {
	if value, ok := env.symbols[symbol]; ok {
		return value, true
	}
	if env.parent != nil {
		return env.parent.GetValue(symbol)
	}
	return NewUnitValue(), false
}

// DefineValue defines a new symbol in the current environment.
func (env *Environment) DefineValue(symbol string, value Value) error {
	if _, found := env.symbols[symbol]; found {
		return fmt.Errorf("symbol %s already defined", symbol)
	}
	env.symbols[symbol] = value
	return nil
}

// SetValue sets the value of an existing symbol in the current environment.
func (env *Environment) SetValue(symbol string, value Value) error {
	// attempt to set the value in the current environment first
	if _, found := env.symbols[symbol]; found {
		env.symbols[symbol] = value
		return nil
	}

	// otherwise attempt to set the value in the parent environment
	if env.parent != nil {
		return env.parent.SetValue(symbol, value)
	}

	// as a base case, say that the symbol does not exist
	return fmt.Errorf("symbol %s not defined", symbol)
}
