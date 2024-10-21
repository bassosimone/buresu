// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalStringLiteral evaluates a StringLiteral node and returns a StringValue with the node's value.
func evalStringLiteral(_ context.Context,
	_ runtime.Environment, node *ast.StringLiteral) (runtime.Value, error) {
	return &runtime.StringValue{Value: node.Value}, nil
}
