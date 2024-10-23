// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"errors"

	"github.com/bassosimone/buresu/pkg/ast"
)

// Callable drefines the traits shared by all callables.
type Callable interface {
	// Call invokes this callable and returns the resulting value.
	Call(ctx context.Context, args ...Value) (Value, error)

	// A Callable is also a Value.
	Value
}

func evalCallExpr(ctx context.Context, env Environment, node *ast.CallExpr) (Value, error) {
	// 1. evaluate the arguments in the current environment
	var args []Value
	for _, arg := range node.Args {
		value, err := Eval(ctx, env, arg)
		if err != nil {
			return nil, err
		}
		args = append(args, value)
	}

	// 2. fetch and invoke the callable
	callable, err := env.EvalCallable(ctx, node.Callable)
	if err != nil {
		return nil, err
	}
	result, err := callable.Call(ctx, args...)

	// 3. handle early return
	var retErr *errReturn
	if errors.As(err, &retErr) {
		return retErr.value, nil
	}
	return result, err
}
