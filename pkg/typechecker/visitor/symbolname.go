// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkSymbolName(_ context.Context, env Environment, node *ast.SymbolName) (Type, error) {
	return env.GetType(node.Value)
}
