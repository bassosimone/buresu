// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"strconv"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalIntLiteral evaluates an IntLiteral node and returns an IntValue with the parsed integer.
func evalIntLiteral(_ context.Context, _ *Environment, node *ast.IntLiteral) (Value, error) {
	value, err := strconv.Atoi(node.Value)
	if err != nil {
		return nil, wrapError(node.Token, err)
	}
	return NewIntValue(value), nil
}
