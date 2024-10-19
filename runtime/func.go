// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/ast"
)

// Callable is the common interface for callables types in the runtime.
type Callable interface {
	Call(ctx context.Context, env Environment, args ...Value) (Value, error)
}

// LambdaValue represents a user-defined lambda function in the source code.
type LambdaValue struct {
	// Closure is the environment in which the lambda was defined.
	Closure Environment

	// Node is the AST node representing the lambda expression.
	Node *ast.LambdaExpr
}

var (
	_ Callable = (*LambdaValue)(nil)
	_ Value    = (*LambdaValue)(nil)
)

// String implements Value.
func (fx *LambdaValue) String() string {
	return fmt.Sprintf("%s: (lambda %s %s)", fx.Node.Token, fx.Node.Params, fx.Node.Expr)
}

// Call invokes the user-defined function within the given environment, including the
// original closure, and returns the result of the function call.
func (fx *LambdaValue) Call(ctx context.Context, _ Environment, args ...Value) (Value, error) {
	// 1. check whether the number of arguments is correct
	if len(fx.Node.Params) != len(args) {
		err := fmt.Errorf("wrong number of arguments: expected %d, got %d", len(fx.Node.Params), len(args))
		return nil, err
	}

	// 2. create the environment for the function call, which is a child
	// of the closure environment with the parameters bound to the arguments
	closure := fx.Closure.PushFunctionScope()
	for idx, arg := range args {
		if err := closure.DefineValue(fx.Node.Params[idx], arg); err != nil {
			return nil, err
		}
	}

	// 3. evaluate the body of the lambda function in the new environment
	return closure.Eval(ctx, fx.Node.Expr)
}

// BuiltInFuncValue is a function that is built-in in the runtime.
//
// The zero value is not ready for use.
type BuiltInFuncValue struct {
	// Name is the name of the built-in function.
	Name string

	// Fx is the actual function that implements the built-in function.
	Fx func(ctx context.Context, env Environment, args ...Value) (Value, error)
}

var (
	_ Callable = (*BuiltInFuncValue)(nil)
	_ Value    = (*BuiltInFuncValue)(nil)
)

// String returns the `<builtin: %p>` string representation of the built-in function.
func (fx *BuiltInFuncValue) String() string {
	return fmt.Sprintf("<builtin: %s>", fx.Name)
}

// Call calls the built-in function with the given arguments.
func (fx *BuiltInFuncValue) Call(ctx context.Context, env Environment, args ...Value) (Value, error) {
	return fx.Fx(ctx, env, args...)
}
