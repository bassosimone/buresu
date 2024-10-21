// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalSetExpr evaluates a set expression.
func evalSetExpr(ctx context.Context,
	env runtime.Environment, node *ast.SetExpr) (runtime.Value, error) {
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	if err := env.SetValue(node.Symbol, value); err != nil {
		err := wrapError(node.Token, err)
		return nil, err
	}
	return value, nil
}
