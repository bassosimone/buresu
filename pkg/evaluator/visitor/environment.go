// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// Environment is the generic interface for the environment.
type Environment interface {
	// DefineValue defines a new symbol in the current environment.
	DefineValue(symbol string, value Value) error

	// EvalCallable attempts to evaluate a node as a callable. Generally, this
	// takes one of two forms: a lambda expression invoked inline or a symbol
	// that was previously defined as a lambda expression.
	EvalCallable(ctx context.Context, node ast.Node) (Callable, error)

	// GetValue returns the value associated with the given symbol.
	//
	// If the symbol is not found in the current environment, the parent
	// environments are searched recursively.
	GetValue(symbol string) (Value, error)

	// NewBoolValue returns a new bool value instance.
	NewBoolValue(value bool) Value

	// NewLambdaValue returns a new lambda instance.
	NewLambdaValue(node *ast.LambdaExpr) Value

	// NewFloat64Value returns a new float64 value instance.
	NewFloat64Value(value float64) Value

	// NewIntValue returns a new int value instance.
	NewIntValue(value int) Value

	// NewQuotedValue returns a new quoted value instance.
	NewQuotedValue(node *ast.QuoteExpr) Value

	// NewStringValue returns a new string value instance.
	NewStringValue(value string) Value

	// NewUnitValue returns a new unit value instance.
	NewUnitValue() Value

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

	// UnwrapBoolValue attempts to unwrap a Bool from a Value and
	// returns either the unwrapped value of an error.
	UnwrapBoolValue(value Value) (bool, error)

	// WrapError wraps an error adding contextual token information.
	WrapError(tok token.Token, err error) error
}
