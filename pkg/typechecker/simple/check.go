// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// Check evaluates a node in the AST and returns the result.
func Check(ctx context.Context, env *Environment, node ast.Node) (visitor.Type, error) {
	return visitor.Check(ctx, env, node)
}
