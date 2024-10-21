// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalTrueLiteral evaluates a TrueLiteral node and returns a BoolValue set to true.
func evalTrueLiteral(_ context.Context, _ *Environment, _ *ast.TrueLiteral) (Value, error) {
	return NewBoolValue(true), nil
}
