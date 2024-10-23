// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalQuoteExpr(_ context.Context, env Environment, node *ast.QuoteExpr) (Value, error) {
	return env.NewQuotedValue(node), nil
}
