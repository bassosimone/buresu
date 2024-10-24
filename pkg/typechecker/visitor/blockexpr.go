// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkBlockExpr(ctx context.Context, env Environment, node *ast.BlockExpr) (Type, error) {
	var (
		err    error
		rvType Type = env.NewUnitType()
	)
	env = env.PushBlockScope() // create a new environment for the block scope
	for _, expr := range node.Exprs {
		rvType, err = Check(ctx, env, expr)
		if err != nil {
			return nil, err
		}
	}
	return rvType, nil
}
