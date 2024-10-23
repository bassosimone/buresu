// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/evaluator/visitor"

// NewUnitValue implements [visitor.Environment].
func (env *Environment) NewUnitValue() visitor.Value {
	return &Unit{}
}

// Unit represents a unit value.
type Unit struct{}

// Ensure Unit implements [visitor.Value].
var _ visitor.Value = (*Unit)(nil)

// String implements [visitor.Value].
func (*Unit) String() string {
	return "()"
}

// Ensure Unit implements [Seq].
var _ Seq = (*Unit)(nil)

// Length implements [Seq].
func (*Unit) Length() int {
	return 0
}
