// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"strconv"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalFloatLiteral evaluates a FloatLiteral node and returns a Float64Value with the parsed float.
func evalFloatLiteral(_ context.Context, _ *Environment, node *ast.FloatLiteral) (Value, error) {
	value, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return nil, wrapError(node.Token, err)
	}
	return NewFloat64Value(value), nil
}
