// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalTrueLiteral evaluates a TrueLiteral node and returns a BoolValue set to true.
func evalTrueLiteral(_ context.Context,
	_ runtime.Environment, _ *ast.TrueLiteral) (runtime.Value, error) {
	return &runtime.BoolValue{Value: true}, nil
}
