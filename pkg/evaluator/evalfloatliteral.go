// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"strconv"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalFloatLiteral evaluates a FloatLiteral node and returns a FloatValue with the parsed float.
func evalFloatLiteral(_ context.Context,
	_ runtime.Environment, node *ast.FloatLiteral) (runtime.Value, error) {
	value, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return nil, wrapError(node.Token, err)
	}
	return &runtime.Float64Value{Value: value}, nil
}
