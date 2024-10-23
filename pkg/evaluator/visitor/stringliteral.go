// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalStringLiteral(_ context.Context, env Environment, node *ast.StringLiteral) (Value, error) {
	return env.NewStringValue(node.Value), nil
}
