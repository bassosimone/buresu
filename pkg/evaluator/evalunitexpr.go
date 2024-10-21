// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalUnitExpr evaluates a UnitExpr node and returns a UnitValue.
func evalUnitExpr(_ context.Context, _ *Environment, _ *ast.UnitExpr) (Value, error) {
	return NewUnitValue(), nil
}
