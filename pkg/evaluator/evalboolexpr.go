// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
	"github.com/bassosimone/buresu/pkg/token"
)

// evalBooleanExpr evaluates a boolean expression and returns the boolean value.
func evalBooleanExpr(ctx context.Context, env runtime.Environment,
	predicate ast.Node, token token.Token) (*runtime.BoolValue, error) {
	value, err := Eval(ctx, env, predicate)
	if err != nil {
		return nil, err
	}
	boolVal, ok := value.(*runtime.BoolValue)
	if !ok {
		return nil, newError(token, "predicate must be a boolean value")
	}
	return boolVal, nil
}
