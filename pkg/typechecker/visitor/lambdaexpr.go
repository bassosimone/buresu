// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalLambdaExpr(_ context.Context, env Environment, node *ast.LambdaExpr) (Type, error) {
	return env.NewLambdaType(node)
}
