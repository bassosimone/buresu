// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalCallExpr evaluates a call expression.
func evalCallExpr(ctx context.Context,
	env runtime.Environment, node *ast.CallExpr) (runtime.Value, error) {
	// 1. evaluate the arguments in the current environment
	var args []runtime.Value
	for _, arg := range node.Args {
		value, err := Eval(ctx, env, arg)
		if err != nil {
			return nil, err
		}
		args = append(args, value)
	}

	// 2. fetch the actual callable value
	maybeCallable, err := Eval(ctx, env, node.Callable)
	if err != nil {
		return nil, err
	}
	callable, ok := maybeCallable.(runtime.Callable)
	if !ok {
		return nil, newError(node.Token, fmt.Sprintf("cannot call a %T", maybeCallable))
	}

	// 3. closures create their own stack frame, while builtin functions
	// shouldn't need a stack, but let's create one anyway.
	if _, ok := callable.(*runtime.BuiltInFuncValue); ok {
		env = env.PushFunctionScope()
	}

	// 4. actually invoke the callable
	return callable.Call(ctx, env, args...)
}

// evalLambdaExpr packages a LambdaValue capturing the current scope and returns the value.
func evalLambdaExpr(_ context.Context,
	env runtime.Environment, node *ast.LambdaExpr) (runtime.Value, error) {
	value := &runtime.LambdaValue{Closure: env, Node: node}
	return value, nil
}
