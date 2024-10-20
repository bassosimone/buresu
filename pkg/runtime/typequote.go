// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import "github.com/bassosimone/buresu/pkg/ast"

// QuoteValue is the type of a quoted value.
type QuoteValue struct {
	Value ast.Node
}

// Ensure *BoolValue implements Value interface.
var _ Value = (*QuoteValue)(nil)

// String implements Value.
func (q *QuoteValue) String() string {
	return q.Value.String()
}
