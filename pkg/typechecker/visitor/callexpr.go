// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkCallExpr(ctx context.Context, env Environment, node *ast.CallExpr) (Type, error) {
	// resolve and check the arguments in the current environment
	var argsTypes []Type
	for _, arg := range node.Args {
		argType, err := Check(ctx, env, arg)
		if err != nil {
			return nil, err
		}
		argsTypes = append(argsTypes, argType)
	}

	// resolve and check the callable
	//
	// it will be the callable's responsibility to push a
	// function scope for the arguments
	return env.Call(ctx, node.Callable, argsTypes...)
}
