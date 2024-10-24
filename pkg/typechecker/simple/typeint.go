// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// NewIntType implements [visitor.Environment].
func (env *Environment) NewIntType() visitor.Type {
	return &Int{}
}

// Int represents a int type.
type Int struct{}

// Ensure Int implements [visitor.Type].
var _ visitor.Type = (*Int)(nil)

// String implements [visitor.Type].
func (v *Int) String() string {
	return "Int"
}
