// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"
)

// BuiltInFunc is a function that is built-in to the evaluator.
type BuiltInFunc func(ctx context.Context, env *Environment, args ...Value) (Value, error)

// BuiltInFuncValue is a built-in function value.
type BuiltInFuncValue struct {
	// Name is the name of the function.
	Name string

	// Fx is the function itself.
	Fx BuiltInFunc

	// Prefix is the type annotation prefix.
	Prefix string
}

// NewBuiltInFuncValue creates a new built-in function value.
func NewBuiltInFuncValue(name, prefix string, fx BuiltInFunc) *BuiltInFuncValue {
	return &BuiltInFuncValue{Name: name, Fx: fx, Prefix: prefix}
}

var (
	_ CallableTrait = (*BuiltInFuncValue)(nil)
	_ Value         = (*BuiltInFuncValue)(nil)
)

// Call implements CallableTrait.
func (bf *BuiltInFuncValue) Call(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// 1. push a new frame for the function call
	env = env.pushFunctionScope()

	// 2. invoke the built-in function proper
	return bf.Fx(ctx, env, args...)
}

// String implements Value.
func (bf *BuiltInFuncValue) String() string {
	return fmt.Sprintf("<built-in function %s>", bf.Name)
}

// Type implements Value.
func (bf *BuiltInFuncValue) Type() string {
	return "<built-in function>"
}

// TypeAnnotationPrefix implements CallableTrait.
func (bf *BuiltInFuncValue) TypeAnnotationPrefix() string {
	return bf.Prefix
}
