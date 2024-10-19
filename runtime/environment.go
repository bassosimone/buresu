// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"io"

	"github.com/bassosimone/buresu/ast"
)

// Environment represents the runtime environment for the interpreter,
// managing variable bindings and scope hierarchy.
type Environment interface {
	// DefineValue defines a new symbol in the current environment.
	DefineValue(symbol string, value Value) error

	// Eval evaluates the given AST node in the current environment
	// and returns the rssult of the evaluation.
	Eval(ctx context.Context, node ast.Node) (Value, error)

	// GetValue returns the value associated with the given symbol.
	//
	// If the symbol is not found in the current environment, the parent
	// environments are searched recursively.
	GetValue(symbol string) (Value, bool)

	// IsInsideFunc returns true if the environment is a function scope
	// or any of the parent environments is a function scope.
	IsInsideFunc() bool

	// Output returns the output [io.Writer] of the environment.
	Output() io.Writer

	// PushBlockScope creates a new child environment associated with
	// a block execution and returns it. The returned environment will
	// use the current environment as its parent.
	PushBlockScope() Environment

	// PushFunctionScope creates a new child environment associated with
	// a function execution and returns it. The returned environment will
	// use the current environment as its parent.
	PushFunctionScope() Environment

	// SetValue sets the value of an existing symbol in the current environment.
	SetValue(symbol string, value Value) error
}

// InitRootScope initializes a new environment with with built-in functions.
func InitRootScope(env Environment) error {
	// define boolean constants
	if err := env.DefineValue("false", &BoolValue{Value: false}); err != nil {
		return err
	}
	if err := env.DefineValue("true", &BoolValue{Value: true}); err != nil {
		return err
	}

	// define the sumInt and sumFloat functions
	if err := env.DefineValue(intSumFunc.Name, intSumFunc); err != nil {
		return err
	}
	if err := env.DefineValue(floatSumFunc.Name, floatSumFunc); err != nil {
		return err
	}

	// define the display function
	if err := env.DefineValue(displayFunc.Name, displayFunc); err != nil {
		return err
	}

	return nil
}
