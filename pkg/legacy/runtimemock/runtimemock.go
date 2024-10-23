// SPDX-License-Identifier: GPL-3.0-or-later

package runtimemock

import (
	"context"
	"io"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/legacy/runtime"
)

// MockEnvironment is a mock implementation of the runtime.Environment interface.
// It allows for customizable behavior by setting function pointers for each method.
type MockEnvironment struct {
	// MockDefineValue is a function pointer that mocks the DefineValue method.
	MockDefineValue func(symbol string, value runtime.Value) error

	// MockEval is a function pointer that mocks the Eval method.
	MockEval func(ctx context.Context, node ast.Node) (runtime.Value, error)

	// MockGetValue is a function pointer that mocks the GetValue method.
	MockGetValue func(symbol string) (runtime.Value, bool)

	// MockIsInsideFunc is a function pointer that mocks the IsInsideFunc method.
	MockIsInsideFunc func() bool

	// MockOutput is a function pointer that mocks the Output method.
	MockOutput func() io.Writer

	// MockPushBlockScope is a function pointer that mocks the PushBlockScope method.
	MockPushBlockScope func() runtime.Environment

	// MockPushFunctionScope is a function pointer that mocks the PushFunctionScope method.
	MockPushFunctionScope func() runtime.Environment

	// MockSetValue is a function pointer that mocks the SetValue method.
	MockSetValue func(symbol string, value runtime.Value) error
}

// DefineValue calls the MockDefineValue function pointer.
func (env *MockEnvironment) DefineValue(symbol string, value runtime.Value) error {
	return env.MockDefineValue(symbol, value)
}

// Eval calls the MockEval function pointer.
func (env *MockEnvironment) Eval(ctx context.Context, node ast.Node) (runtime.Value, error) {
	return env.MockEval(ctx, node)
}

// GetValue calls the MockGetValue function pointer.
func (env *MockEnvironment) GetValue(symbol string) (runtime.Value, bool) {
	return env.MockGetValue(symbol)
}

// IsInsideFunc calls the MockIsInsideFunc function pointer.
func (env *MockEnvironment) IsInsideFunc() bool {
	return env.MockIsInsideFunc()
}

// Output calls the MockOutput function pointer.
func (env *MockEnvironment) Output() io.Writer {
	return env.MockOutput()
}

// PushBlockScope calls the MockPushBlockScope function pointer.
func (env *MockEnvironment) PushBlockScope() runtime.Environment {
	return env.MockPushBlockScope()
}

// PushFunctionScope calls the MockPushFunctionScope function pointer.
func (env *MockEnvironment) PushFunctionScope() runtime.Environment {
	return env.MockPushFunctionScope()
}

// SetValue calls the MockSetValue function pointer.
func (env *MockEnvironment) SetValue(symbol string, value runtime.Value) error {
	return env.MockSetValue(symbol, value)
}
