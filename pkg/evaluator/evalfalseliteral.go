// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalFalseLiteral evaluates a FalseLiteral node and returns a BoolValue set to false.
func evalFalseLiteral(_ context.Context,
	_ runtime.Environment, _ *ast.FalseLiteral) (runtime.Value, error) {
	return &runtime.BoolValue{Value: false}, nil
}
