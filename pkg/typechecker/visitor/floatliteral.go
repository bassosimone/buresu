// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkFloatLiteral(_ context.Context, env Environment, _ *ast.FloatLiteral) (Type, error) {
	return env.NewFloat64Type(), nil
}
