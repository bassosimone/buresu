// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalLambdaExpr packages a LambdaValue capturing the current scope and returns the value.
func evalLambdaExpr(_ context.Context,
	env runtime.Environment, node *ast.LambdaExpr) (runtime.Value, error) {
	value := &runtime.LambdaValue{Closure: env, Node: node}
	return value, nil
}
