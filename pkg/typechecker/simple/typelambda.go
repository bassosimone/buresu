// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// NewLambdaType implements [visitor.Environment].
func (env *Environment) NewLambdaType(node *ast.LambdaExpr) (visitor.Type, error) {
	lambda := &Callable{
		ParamsTypes: []visitor.Type{}, // set below
		ReturnType:  &Any{},
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			// create the environment for the function call, which is a child of the
			// closure environment with the parameters bound to the arguments
			closure := env.PushFunctionScope()
			for idx, arg := range args {
				// the parser guarantees that params names are not duplicated
				// so I do not see how this define could fail
				rtx.Must(closure.DefineType(node.Params[idx], arg))
			}

			// check the body of the lambda function in the new environment
			rvType, err := visitor.Check(ctx, closure, node.Expr)
			if err != nil {
				return nil, err
			}

			// merge the rvType with the values that may have been
			// returned by return statements in the lambda body
			return closure.MergeReturnTypes(rvType)
		},
		Previous: nil,
	}

	// by default configure the lambda to accept any type for the parameters
	for idx := 0; idx < len(node.Params); idx++ {
		lambda.ParamsTypes = append(lambda.ParamsTypes, &Any{})
	}

	// TODO(bassosimone): here we should parse the lambda type
	// annotation, if any, and properly fill the params

	return lambda, nil
}
