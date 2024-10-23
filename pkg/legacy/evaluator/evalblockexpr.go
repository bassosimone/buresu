// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalBlockExpr evaluates a block expression by evaluating each expression in
// the block sequentially and returning the result of the last expression.
//
// Note that block creates a new environment for the block scope.
func evalBlockExpr(ctx context.Context, env *Environment, node *ast.BlockExpr) (Value, error) {
	var (
		err    error
		result Value = NewUnitValue()
	)
	env = env.pushBlockScope() // create a new environment for the block scope
	for _, expr := range node.Exprs {
		result, err = Eval(ctx, env, expr)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
