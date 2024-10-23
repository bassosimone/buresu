// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalSymbolName(_ context.Context, env Environment, node *ast.SymbolName) (Value, error) {
	return env.GetValue(node.Value)
}
