// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
)

// IntValue represents a quoted AST node.
//
// Construct using NewIntValue.
type QuotedValue struct {
	Value ast.Node
}

var _ Value = (*IntValue)(nil)

// NewQuotedValue creates a new [*QuotedValue] instance.
func NewQuotedValue(value ast.Node) *QuotedValue {
	return &QuotedValue{value}
}

// String implements Value.
func (q *QuotedValue) String() string {
	return fmt.Sprintf("(quote %s)", q.Value.String())
}

// Type implements Value.
func (*QuotedValue) Type() string {
	return "<quoted value>"
}
