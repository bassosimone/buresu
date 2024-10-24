// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewQuotedValue implements [visitor.Environment].
func (env *Environment) NewQuotedValue(node *ast.QuoteExpr) visitor.Value {
	return &Quoted{node}
}

// Quoted represents a quoted AST node.
type Quoted struct {
	Value *ast.QuoteExpr
}

// Ensure QuotedValue implements Value.
var _ visitor.Value = (*Quoted)(nil)

// String implements Value.
func (q *Quoted) String() string {
	return fmt.Sprintf("%s", q.Value.String())
}
