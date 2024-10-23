// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewLambdaValue implements [visitor.Environment].
func (env *Environment) NewLambdaValue(node *ast.LambdaExpr) visitor.Value {
	return &Lambda{env, node}
}

// Lambda represents a lambda function.
type Lambda struct {
	// Closure is the environment in which the lambda function was defined.
	Closure *Environment

	// Node is the AST node representing the lambda function.
	Node *ast.LambdaExpr
}

// Ensure Lambda implements [visitor.Callable].
var _ visitor.Callable = (*Lambda)(nil)

// Call implements [visitor.Callable].
func (lv *Lambda) Call(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
	// 1. check whether the number of arguments is correct
	if len(lv.Node.Params) != len(args) {
		err := fmt.Errorf("wrong number of arguments: expected %d, got %d", len(lv.Node.Params), len(args))
		return nil, err
	}

	// 2. create the environment for the function call, which is a child
	// of the closure environment with the parameters bound to the arguments
	closure := lv.Closure.PushFunctionScope()
	for idx, arg := range args {
		if err := closure.DefineValue(lv.Node.Params[idx], arg); err != nil {
			return nil, err
		}
	}

	// 3. evaluate the body of the lambda function in the new environment
	return visitor.Eval(ctx, closure, lv.Node.Expr)
}

// Ensure Lambda implements [visitor.Value].
var _ visitor.Value = (*Lambda)(nil)

// String implements [visitor.Value].
func (fx *Lambda) String() string {
	return fmt.Sprintf("%s", fx.Node.String())
}
