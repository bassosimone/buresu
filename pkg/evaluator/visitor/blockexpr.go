// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalBlockExpr(ctx context.Context, env Environment, node *ast.BlockExpr) (Value, error) {
	var (
		err    error
		result Value = env.NewUnitValue()
	)
	env = env.PushBlockScope() // create a new environment for the block scope
	for _, expr := range node.Exprs {
		result, err = Eval(ctx, env, expr)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
