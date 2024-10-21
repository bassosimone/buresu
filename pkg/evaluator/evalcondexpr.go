// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalCondExpr evaluates a conditional expression by evaluating each case's
// predicate until one evaluates to true, then evaluating and returning the
// corresponding expression. If no predicates are true, it evaluates and
// returns the else expression. Note that the predicates must strictly be
// boolean values, otherwise an error is returned.
func evalCondExpr(ctx context.Context,
	env runtime.Environment, node *ast.CondExpr) (runtime.Value, error) {
	for _, condCase := range node.Cases {
		boolVal, err := evalBooleanExpr(ctx, env, condCase.Predicate, node.Token)
		if err != nil {
			return nil, err
		}
		if boolVal.Value {
			return Eval(ctx, env, condCase.Expr)
		}
	}
	return Eval(ctx, env, node.ElseExpr)
}
