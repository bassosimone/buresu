// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"strconv"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalFloatLiteral(_ context.Context, env Environment, node *ast.FloatLiteral) (Value, error) {
	value, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return nil, env.WrapError(node.Token, err)
	}
	return env.NewFloat64Value(value), nil
}
