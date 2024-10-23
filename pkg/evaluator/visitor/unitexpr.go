// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalUnitExpr(_ context.Context, env Environment, _ *ast.UnitExpr) (Value, error) {
	return env.NewUnitValue(), nil
}
