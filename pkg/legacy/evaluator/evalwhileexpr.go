// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalWhileExpr evaluates a while expression by repeatedly evaluating the predicate
// and the body expression as long as the predicate evaluates to true. Note that while
// always returns the singleton unit value to the caller.
func evalWhileExpr(ctx context.Context, env *Environment, node *ast.WhileExpr) (Value, error) {
	for {
		boolVal, err := evalBooleanExpr(ctx, env, node.Predicate, node.Token)
		if err != nil {
			return nil, err
		}
		if !boolVal.Value {
			break
		}
		if _, err = Eval(ctx, env, node.Expr); err != nil {
			return nil, err
		}
	}
	return NewUnitValue(), nil
}
