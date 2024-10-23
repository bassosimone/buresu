// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/evaluator/visitor"

// NewStringValue implements [visitor.Environment].
func (env *Environment) NewStringValue(value string) visitor.Value {
	return &String{value}
}

// String represents a string value.
type String struct {
	Value string
}

// Ensure String implements [visitor.Value].
var _ visitor.Value = (*String)(nil)

// String returns the string representation of the string value.
func (v *String) String() string {
	return v.Value
}

// Ensure String implements [Seq].
var _ Seq = (*String)(nil)

// Length implements [Seq].
func (v *String) Length() int {
	return len(v.Value)
}
