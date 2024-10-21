// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalSetExpr evaluates a set expression.
func evalSetExpr(ctx context.Context, env *Environment, node *ast.SetExpr) (Value, error) {
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
