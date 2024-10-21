// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// TODO(bassosimone): we need to handle the return statement
// inside the block or the function expression.

// errReturn is a special value that is returned when a return statement is
// encountered. It is used to signal the interpreter that the current function
// has returned early. This type also carries the value to return.
type errReturn struct {
	value Value
}

// Error implements the error interface for errReturn.
func (errReturn) Error() string {
	return "return statement"
}

// evalReturnStmt evaluates a return statement.
func evalReturnStmt(ctx context.Context, env *Environment, node *ast.ReturnStmt) (Value, error) {
	if !env.IsInsideFunc() {
		return nil, newError(node.Token, "return statement outside of function")
	}
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	return nil, &errReturn{value: value}
}
