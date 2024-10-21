// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
)

// IntValue represents a lambda value.
//
// Construct using NewLambdaValue.
type LambdaValue struct {
	// Closure is the environment in which the lambda function was defined.
	Closure *Environment

	// Node is the AST node representing the lambda function.
	Node *ast.LambdaExpr
}

var _ Value = (*LambdaValue)(nil)

// NewLambdaValue creates a new [*LambdaValue] instance.
func NewLambdaValue(env *Environment, node *ast.LambdaExpr) *LambdaValue {
	return &LambdaValue{env, node}
}

var _ CallableTrait = (*LambdaValue)(nil)

// Call implements CallableTrait.
func (lv *LambdaValue) Call(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// 1. check whether the number of arguments is correct
	if len(lv.Node.Params) != len(args) {
		err := fmt.Errorf("wrong number of arguments: expected %d, got %d", len(lv.Node.Params), len(args))
		return nil, err
	}

	// 2. create the environment for the function call, which is a child
	// of the closure environment with the parameters bound to the arguments
	closure := lv.Closure.pushFunctionScope()
	for idx, arg := range args {
		if err := closure.DefineValue(lv.Node.Params[idx], arg); err != nil {
			return nil, err
		}
	}

	// 3. evaluate the body of the lambda function in the new environment
	return Eval(ctx, closure, lv.Node.Expr)
}

// String implements Value.
func (fx *LambdaValue) String() string {
	return fmt.Sprintf("%s", fx.Node.String())
}
