// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// Eval evaluates a node in the AST and returns the result.
func Eval(ctx context.Context, env *Environment, node ast.Node) (visitor.Value, error) {
	return visitor.Eval(ctx, env, node)
}
