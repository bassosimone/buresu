// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkCondExpr(ctx context.Context, env Environment, node *ast.CondExpr) (Type, error) {
	var rvTypes []Type

	for _, condCase := range node.Cases {
		if err := env.CheckCondition(ctx, condCase.Predicate); err != nil {
			return nil, err
		}
		rvType, err := Check(ctx, env, condCase.Expr)
		if err != nil {
			return nil, err
		}
		rvTypes = append(rvTypes, rvType)
	}

	rvType, err := Check(ctx, env, node.ElseExpr)
	if err != nil {
		return nil, err
	}
	rvTypes = append(rvTypes, rvType)

	return env.NewUnionType(rvTypes...), nil
}
