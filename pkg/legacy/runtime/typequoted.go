// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
)

// QuotedValue is the type of a quoted value.
type QuotedValue struct {
	Value ast.Node
}

// Ensure *BoolValue implements Value interface.
var _ Value = (*QuotedValue)(nil)

// String implements Value.
func (q *QuotedValue) String() string {
	return fmt.Sprintf("(quote %s)", q.Value.String())
}
