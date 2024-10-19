// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/ast"
	"github.com/bassosimone/buresu/runtime"
	"github.com/bassosimone/buresu/token"
)

// evalBlockExpr evaluates a block expression by evaluating each expression in
// the block sequentially and returning the result of the last expression.
//
// Note that block creates a new environment for the block scope.
func evalBlockExpr(ctx context.Context,
	env runtime.Environment, node *ast.BlockExpr) (runtime.Value, error) {
	var (
		err    error
		result runtime.Value = &runtime.UnitValue{}
	)
	env = env.PushBlockScope() // create a new environment for the block scope
	for _, expr := range node.Exprs {
		result, err = Eval(ctx, env, expr)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// evalCondExpr evaluates a conditional expression by evaluating each case's
// predicate until one evaluates to true, then evaluating and returning the
// corresponding expression. If no predicates are true, it evaluates and
// returns the else expression. Note that the predicates must strictly be
// boolean values, otherwise an error is returned.
func evalCondExpr(ctx context.Context,
	env runtime.Environment, node *ast.CondExpr) (runtime.Value, error) {
	for _, condCase := range node.Cases {
		boolVal, err := evalBooleanPredicate(ctx, env, condCase.Predicate, node.Token)
		if err != nil {
			return nil, err
		}
		if boolVal.Value {
			return Eval(ctx, env, condCase.Expr)
		}
	}
	return Eval(ctx, env, node.ElseExpr)
}

// evalBooleanPredicate evaluates a predicate and ensures it is a boolean value.
func evalBooleanPredicate(ctx context.Context, env runtime.Environment,
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

// errReturn is a special value that is returned when a return statement is
// encountered. It is used to signal the interpreter that the current function
// has returned early. This type also carries the value to return.
type errReturn struct {
	value runtime.Value
}

// Error implements the error interface for errReturn.
func (errReturn) Error() string {
	return "return statement"
}

// evalReturnStmt evaluates a return statement.
func evalReturnStmt(ctx context.Context, env runtime.Environment,
	node *ast.ReturnStmt) (runtime.Value, error) {
	if !env.IsInsideFunc() {
		return nil, newError(node.Token, "return statement outside of function")
	}
	value, err := Eval(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	return nil, &errReturn{value: value}
}

// evalWhileExpr evaluates a while expression by repeatedly evaluating the predicate
// and the body expression as long as the predicate evaluates to true. Note that while
// always returns the singleton unit value to the caller.
func evalWhileExpr(ctx context.Context, env runtime.Environment,
	node *ast.WhileExpr) (runtime.Value, error) {
	for {
		boolVal, err := evalBooleanPredicate(ctx, env, node.Predicate, node.Token)
		if err != nil {
			return nil, err
		}
		if !boolVal.Value {
			break
		}
		if _, err = Eval(ctx, env, node.Expr); err != nil {
			return nil, err
		}
	}
	return &runtime.UnitValue{}, nil
}
