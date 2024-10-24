// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkFalseLiteral(_ context.Context, env Environment, _ *ast.FalseLiteral) (Type, error) {
	return env.NewBoolType(), nil
}
