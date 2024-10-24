// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// NewStringType implements [visitor.Environment].
func (env *Environment) NewStringType() visitor.Type {
	return &String{}
}

// String represents a string type.
type String struct{}

// Ensure String implements [visitor.Type].
var _ visitor.Type = (*String)(nil)

// String returns the string representation of the string value.
func (v *String) String() string {
	return "String"
}
