// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkIntLiteral(_ context.Context, env Environment, _ *ast.IntLiteral) (Type, error) {
	return env.NewIntType(), nil
}
