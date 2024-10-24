// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// Variadic represents a variadic argument list.
type Variadic struct {
	Type visitor.Type
}

// Ensure Unit implements [visitor.Type].
var _ visitor.Type = (*Variadic)(nil)

// String implements [visitor.Type].
func (t *Variadic) String() string {
	return fmt.Sprintf("(Variadic %s)", t.Type.String())
}
