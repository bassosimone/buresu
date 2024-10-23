// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// MockEnvironment is a mock implementation of the Environment interface
// used for testing purposes.
type MockEnvironment struct {
	insideFunc bool
	values     map[string]Value
}

// NewMockEnvironment creates a new instance of MockEnvironment.
func NewMockEnvironment() *MockEnvironment {
	return &MockEnvironment{values: make(map[string]Value)}
}

// DefineValue defines a new symbol in the mock environment.
func (env *MockEnvironment) DefineValue(symbol string, value Value) error {
	env.values[symbol] = value
	return nil
}

// EvalCallable attempts to evaluate a node as a callable in the mock environment.
func (env *MockEnvironment) EvalCallable(ctx context.Context, node ast.Node) (Callable, error) {
	symbol, ok := node.(*ast.SymbolName)
	if !ok {
		return nil, fmt.Errorf("node is not a symbol")
	}
	value, exists := env.values[symbol.Value]
	if !exists {
		return nil, fmt.Errorf("symbol not found")
	}
	callable, ok := value.(Callable)
	if !ok {
		return nil, fmt.Errorf("value is not callable")
	}
	return callable, nil
}

// GetValue returns the value associated with the given symbol in the mock environment.
func (env *MockEnvironment) GetValue(symbol string) (Value, error) {
	value, exists := env.values[symbol]
	if !exists {
		return nil, fmt.Errorf("symbol not found")
	}
	return value, nil
}

// IsInsideFunc returns true if the mock environment is a function scope.
func (env *MockEnvironment) IsInsideFunc() bool {
	return env.insideFunc
}

// NewBoolValue returns a new boolean value instance in the mock environment.
func (env *MockEnvironment) NewBoolValue(value bool) Value {
	return MockValue{value: value}
}

// NewLambdaValue returns a new lambda instance in the mock environment.
func (env *MockEnvironment) NewLambdaValue(node *ast.LambdaExpr) Value {
	return MockValue{value: node}
}

// NewFloat64Value returns a new float64 value instance in the mock environment.
func (env *MockEnvironment) NewFloat64Value(value float64) Value {
	return MockValue{value: value}
}

// NewIntValue returns a new int value instance in the mock environment.
func (env *MockEnvironment) NewIntValue(value int) Value {
	return MockValue{value: value}
}

// NewQuotedValue returns a new quoted value instance in the mock environment.
func (env *MockEnvironment) NewQuotedValue(node *ast.QuoteExpr) Value {
	return MockValue{value: node}
}

// NewStringValue returns a new string value instance in the mock environment.
func (env *MockEnvironment) NewStringValue(value string) Value {
	return MockValue{value: value}
}

// NewUnitValue returns a new unit value instance in the mock environment.
func (env *MockEnvironment) NewUnitValue() Value {
	return MockValue{value: nil}
}

// PushBlockScope creates a new child environment for a block scope in the mock environment.
func (env *MockEnvironment) PushBlockScope() Environment {
	return env
}

// PushFunctionScope creates a new child environment for a function scope in the mock environment.
func (env *MockEnvironment) PushFunctionScope() Environment {
	return env
}

// SetValue sets the value of an existing symbol in the mock environment.
func (env *MockEnvironment) SetValue(symbol string, value Value) error {
	env.values[symbol] = value
	return nil
}

// UnwrapBoolValue attempts to unwrap a boolean from a value in the mock environment.
func (env *MockEnvironment) UnwrapBoolValue(value Value) (bool, error) {
	boolVal, ok := value.(MockValue)
	if !ok {
		return false, fmt.Errorf("value is not a boolean")
	}
	return boolVal.value.(bool), nil
}

// WrapError wraps an error with contextual token information in the mock environment.
func (env *MockEnvironment) WrapError(tok token.Token, err error) error {
	return fmt.Errorf("%s: %w", tok.Value, err)
}

// MockValue is a mock implementation of the Value interface used for testing purposes.
type MockValue struct {
	value any
}

// String converts the MockValue to a string representation.
func (v MockValue) String() string {
	return fmt.Sprintf("%v", v.value)
}

// MockCallable is a mock implementation of the Callable interface used for testing purposes.
type MockCallable struct {
	fn func(ctx context.Context, args ...Value) (Value, error)
}

// NewMockCallable creates a new instance of MockCallable.
func NewMockCallable(fn func(ctx context.Context, args ...Value) (Value, error)) Callable {
	return MockCallable{fn: fn}
}

// Call invokes the mock callable with the given arguments.
func (c MockCallable) Call(ctx context.Context, args ...Value) (Value, error) {
	return c.fn(ctx, args...)
}

// String converts the MockCallable to a string representation.
func (c MockCallable) String() string {
	return "mockCallable"
}
