// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalUnitExpr evaluates a UnitExpr node and returns a UnitValue.
func evalUnitExpr(_ context.Context,
	_ runtime.Environment, _ *ast.UnitExpr) (runtime.Value, error) {
	return runtime.NewUnitValue(), nil
}
