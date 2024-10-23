// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalStringLiteral evaluates a StringLiteral node and returns a StringValue with the node's value.
func evalStringLiteral(_ context.Context, _ *Environment, node *ast.StringLiteral) (Value, error) {
	return NewStringValue(node.Value), nil
}
