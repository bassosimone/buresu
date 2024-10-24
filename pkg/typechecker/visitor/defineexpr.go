// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// checkDefineExpr evaluates a define expression.
func checkDefineExpr(ctx context.Context, env Environment, node *ast.DefineExpr) (Type, error) {
	exprType, err := Check(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}

	if err := env.DefineType(node.Symbol, exprType); err != nil {
		return nil, env.WrapError(node.Token, err)
	}

	return exprType, nil
}
