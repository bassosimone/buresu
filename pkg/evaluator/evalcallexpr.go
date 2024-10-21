// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalCallExpr evaluates a call expression.
func evalCallExpr(ctx context.Context, env *Environment, node *ast.CallExpr) (Value, error) {
	// 1. evaluate the arguments in the current environment
	var args []Value
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
	callable, ok := maybeCallable.(CallableTrait)
	if !ok {
		return nil, newError(node.Token, fmt.Sprintf("cannot call a %T", maybeCallable))
	}

	// 3. actually invoke the callable
	return callable.Call(ctx, env, args...)
}
