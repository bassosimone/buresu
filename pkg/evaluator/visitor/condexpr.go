// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalCondExpr(ctx context.Context, env Environment, node *ast.CondExpr) (Value, error) {
	for _, condCase := range node.Cases {
		expr, err := Eval(ctx, env, condCase.Predicate)
		if err != nil {
			return nil, err
		}
		condition, err := env.UnwrapBoolValue(expr)
		if err != nil {
			return nil, env.WrapError(node.Token, err)
		}
		if condition {
			return Eval(ctx, env, condCase.Expr)
		}
	}
	return Eval(ctx, env, node.ElseExpr)
}
