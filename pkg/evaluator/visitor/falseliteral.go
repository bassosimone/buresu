// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalFalseLiteral(_ context.Context, env Environment, _ *ast.FalseLiteral) (Value, error) {
	return env.NewBoolValue(false), nil
}
