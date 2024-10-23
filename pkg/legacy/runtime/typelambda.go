// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
)

// LambdaValue represents a user-defined lambda function in the source code.
type LambdaValue struct {
	// Closure is the environment in which the lambda was defined.
	Closure Environment

	// Node is the AST node representing the lambda expression.
	Node *ast.LambdaExpr
}

var (
	_ Callable = (*LambdaValue)(nil)
	_ Value    = (*LambdaValue)(nil)
)

// String implements Value.
func (fx *LambdaValue) String() string {
	return fmt.Sprintf("%s", fx.Node.String())
}

// Call invokes the user-defined function within the given environment, including the
// original closure, and returns the result of the function call.
func (fx *LambdaValue) Call(ctx context.Context, _ Environment, args ...Value) (Value, error) {
	// 1. check whether the number of arguments is correct
	if len(fx.Node.Params) != len(args) {
		err := fmt.Errorf("wrong number of arguments: expected %d, got %d", len(fx.Node.Params), len(args))
		return nil, err
	}

	// 2. create the environment for the function call, which is a child
	// of the closure environment with the parameters bound to the arguments
	closure := fx.Closure.PushFunctionScope()
	for idx, arg := range args {
		if err := closure.DefineValue(fx.Node.Params[idx], arg); err != nil {
			return nil, err
		}
	}

	// 3. evaluate the body of the lambda function in the new environment
	return closure.Eval(ctx, fx.Node.Expr)
}
