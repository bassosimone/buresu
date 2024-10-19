// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/ast"
	"github.com/bassosimone/buresu/runtime"
)

// evalDefineExpr evaluates a define expression.
func evalDefineExpr(ctx context.Context,
	env runtime.Environment, node *ast.DefineExpr) (runtime.Value, error) {
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	if err := env.DefineValue(node.Symbol, value); err != nil {
		err := wrapError(node.Token, err)
		return nil, err
	}
	return value, nil
}

// evalSetExpr evaluates a set expression.
func evalSetExpr(ctx context.Context,
	env runtime.Environment, node *ast.SetExpr) (runtime.Value, error) {
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	if err := env.SetValue(node.Symbol, value); err != nil {
		err := wrapError(node.Token, err)
		return nil, err
	}
	return value, nil
}

// evalSymbolName evaluates a symbol name and returns its value by
// searching in the current environment and in the parent environments.
func evalSymbolName(_ context.Context,
	env runtime.Environment, node *ast.SymbolName) (runtime.Value, error) {
	value, found := env.GetValue(node.Value)
	if !found {
		return nil, newError(node.Token, "symbol %s not defined", node.Value)
	}
	return value, nil
}
