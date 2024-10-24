// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkQuoteExpr(_ context.Context, env Environment, node *ast.QuoteExpr) (Type, error) {
	return env.NewQuotedType(node), nil
}
