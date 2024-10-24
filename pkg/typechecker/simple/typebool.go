// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// NewBoolType implements [visitor.Environment].
func (env *Environment) NewBoolType() visitor.Type {
	return &Bool{}
}

// Bool represents a boolean type.
type Bool struct{}

// Ensure Bool implements [visitor.Type].
var _ visitor.Type = (*Bool)(nil)

// String implements [visitor.Type].
func (v *Bool) String() string {
	return "Bool"
}
