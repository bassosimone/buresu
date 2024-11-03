// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// checkDeclareExpr evaluates a declare expression.
func checkDeclareExpr(ctx context.Context, env Environment, node *ast.DeclareExpr) (Type, error) {
	// From the point of view of the typechecker, a declare expression is
	// equivalent to a define expression: it sets the type of a given lambda
	// just like define does. The distinction between declare and define is
	// useful only for the interpreter, which will ignore declare expressions.
	exprType, err := Check(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}

	if err := env.DefineType(node.Symbol, exprType); err != nil {
		return nil, env.WrapError(node.Token, err)
	}

	return exprType, nil
}
