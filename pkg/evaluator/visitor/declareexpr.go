// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalDeclareExpr evaluates a declare expression.
func evalDeclareExpr(_ context.Context, env Environment, _ *ast.DeclareExpr) (Value, error) {
	// The evaluator is not concerned at all with declare expressions because
	// we handle these expressions inside the typechecker. As far as the evaluator is
	// concerned, a declare expression is just a way to get the unit value.
	return env.NewUnitValue(), nil
}
