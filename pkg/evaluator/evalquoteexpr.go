// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalQuoteExpr evaluates a QuoteExpr node and returns a QuoteValue set to its content.
func evalQuoteExpr(_ context.Context,
	_ runtime.Environment, node *ast.QuoteExpr) (runtime.Value, error) {
	return &runtime.QuotedValue{Value: node.Expr}, nil
}
