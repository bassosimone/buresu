// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

// evalQuoteExpr evaluates a QuoteExpr node and returns a QuoteValue set to its content.
func evalQuoteExpr(_ context.Context, _ *Environment, node *ast.QuoteExpr) (Value, error) {
	return NewQuotedValue(node.Expr), nil
}
