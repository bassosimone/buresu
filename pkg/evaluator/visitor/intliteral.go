// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"strconv"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalIntLiteral(_ context.Context, env Environment, node *ast.IntLiteral) (Value, error) {
	value, err := strconv.Atoi(node.Value)
	if err != nil {
		return nil, env.WrapError(node.Token, err)
	}
	return env.NewIntValue(value), nil
}
