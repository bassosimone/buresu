// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

// IntValue represents a lambda value.
//
// Construct using NewLambdaValue.
type LambdaValue struct {
	// AnnotationPrefix is the type annotation prefix of the lambda function.
	AnnotationPrefix string

	// Closure is the environment in which the lambda function was defined.
	Closure *Environment

	// Node is the AST node representing the lambda function.
	Node *ast.LambdaExpr
}

var _ Value = (*LambdaValue)(nil)

// NewLambdaValue creates a new [*LambdaValue] instance.
func NewLambdaValue(env *Environment, node *ast.LambdaExpr) *LambdaValue {
	var annotation string
	if ap, err := typeannotation.Parse(node.Docs); err == nil && ap != nil {
		annotation = ap.ArgumentsAnnotationPrefix()
	}
	return &LambdaValue{annotation, env, node}
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

// TypeAnnotationPrefix implements CallableTrait.
func (fx *LambdaValue) TypeAnnotationPrefix() string {
	return fx.AnnotationPrefix
}

// Type implements Value.
func (fx *LambdaValue) Type() string {
	return "<lambda>"
}
