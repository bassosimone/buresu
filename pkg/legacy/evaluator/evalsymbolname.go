// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalSymbolName evaluates a symbol name and returns its value by
// searching in the current environment and in the parent environments.
func evalSymbolName(_ context.Context, env *Environment, node *ast.SymbolName) (Value, error) {
	value, found := env.GetValue(node.Value)
	if !found {
		return nil, newError(node.Token, "symbol %s not defined", node.Value)
	}
	return value, nil
}
