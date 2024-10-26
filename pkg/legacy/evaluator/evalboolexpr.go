// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// evalBooleanExpr evaluates a boolean expression and returns a boolean value.
func evalBooleanExpr(ctx context.Context, env *Environment,
	predicate ast.Node, token token.Token) (*BoolValue, error) {
	value, err := Eval(ctx, env, predicate)
	if err != nil {
		return nil, err
	}
	boolVal, ok := value.(*BoolValue)
	if !ok {
		return nil, newError(token, "predicate must be a boolean value")
	}
	return boolVal, nil
}
