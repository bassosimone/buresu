// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// NewQuotedType implements [visitor.Environment].
func (env *Environment) NewQuotedType(node *ast.QuoteExpr) visitor.Type {
	return &Any{}
}
