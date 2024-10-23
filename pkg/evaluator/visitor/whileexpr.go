// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalWhileExpr(ctx context.Context, env Environment, node *ast.WhileExpr) (Value, error) {
	for {
		expr, err := Eval(ctx, env, node.Predicate)
		if err != nil {
			return nil, err
		}
		boolVal, err := env.UnwrapBoolValue(expr)
		if err != nil {
			return nil, err
		}
		if !boolVal {
			break
		}
		if _, err = Eval(ctx, env, node.Expr); err != nil {
			return nil, err
		}
	}
	return env.NewUnitValue(), nil
}
