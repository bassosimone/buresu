// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalDefineExpr evaluates a define expression.
func evalDefineExpr(ctx context.Context, env Environment, node *ast.DefineExpr) (Value, error) {
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	if err := env.DefineValue(node.Symbol, value); err != nil {
		return nil, env.WrapError(node.Token, err)
	}
	return value, nil
}
