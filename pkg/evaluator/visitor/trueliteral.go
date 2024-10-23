// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalTrueLiteral(_ context.Context, env Environment, _ *ast.TrueLiteral) (Value, error) {
	return env.NewBoolValue(true), nil
}
