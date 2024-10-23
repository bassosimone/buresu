// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"errors"

	"github.com/bassosimone/buresu/pkg/ast"
)

type errReturn struct {
	value Value
}

func (errReturn) Error() string {
	return "return statement"
}

func evalReturnStmt(ctx context.Context, env Environment, node *ast.ReturnStmt) (Value, error) {
	if !env.IsInsideFunc() {
		return nil, env.WrapError(node.Token, errors.New("return statement outside of function"))
	}
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	return nil, &errReturn{value: value}
}
