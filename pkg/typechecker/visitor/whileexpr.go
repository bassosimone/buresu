// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkWhileExpr(ctx context.Context, env Environment, node *ast.WhileExpr) (Type, error) {
	if err := env.CheckCondition(ctx, node.Predicate); err != nil {
		return nil, err
	}

	if _, err := Check(ctx, env, node.Expr); err != nil {
		return nil, err
	}

	return env.NewUnitType(), nil
}
