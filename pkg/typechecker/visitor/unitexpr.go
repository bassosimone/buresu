// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkUnitExpr(_ context.Context, env Environment, _ *ast.UnitExpr) (Type, error) {
	return env.NewUnitType(), nil
}
