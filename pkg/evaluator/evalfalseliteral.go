// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalFalseLiteral evaluates a FalseLiteral node and returns a BoolValue set to false.
func evalFalseLiteral(_ context.Context, _ *Environment, _ *ast.FalseLiteral) (Value, error) {
	return NewBoolValue(false), nil
}
