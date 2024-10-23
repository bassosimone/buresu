// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalLambdaExpr packages a LambdaValue capturing the current scope and returns the value.
func evalLambdaExpr(_ context.Context, env *Environment, node *ast.LambdaExpr) (Value, error) {
	return NewLambdaValue(env, node), nil
}
