// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkSetExpr(ctx context.Context, env Environment, node *ast.SetExpr) (Type, error) {
	exprType, err := Check(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}

	if err := env.SetType(node.Symbol, exprType); err != nil {
		return nil, env.WrapError(node.Token, err)
	}

	return exprType, nil
}
