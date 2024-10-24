// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// NewUnitType implements [visitor.Environment].
func (env *Environment) NewUnitType() visitor.Type {
	return &Unit{}
}

// Unit represents a unit value.
type Unit struct{}

// Ensure Unit implements [visitor.Type].
var _ visitor.Type = (*Unit)(nil)

// String implements [visitor.Type].
func (*Unit) String() string {
	return "Unit"
}
