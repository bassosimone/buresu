// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalDefineExpr evaluates a define expression.
func evalDefineExpr(ctx context.Context, env *Environment, node *ast.DefineExpr) (Value, error) {
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	if err := env.DefineValue(node.Symbol, value); err != nil {
		err := wrapError(node.Token, err)
		return nil, err
	}
	return value, nil
}
