// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalSetExpr(ctx context.Context, env Environment, node *ast.SetExpr) (Value, error) {
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	if err := env.SetValue(node.Symbol, value); err != nil {
		return nil, env.WrapError(node.Token, err)
	}
	return value, nil
}
