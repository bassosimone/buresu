// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

type errReturn struct {
	value Value
}

func (errReturn) Error() string {
	return "return statement"
}

func evalReturnStmt(ctx context.Context, env Environment, node *ast.ReturnStmt) (Value, error) {
	// the parse guarantees that a return! statement only happens inside a lambda
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	return nil, &errReturn{value: value}
}
