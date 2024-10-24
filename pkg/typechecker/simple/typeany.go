// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// Any represents any type.
type Any struct{}

// Ensure Any implements [visitor.Type].
var _ visitor.Type = (*Any)(nil)

// String implements [visitor.Type].
func (v *Any) String() string {
	return "Any"
}
