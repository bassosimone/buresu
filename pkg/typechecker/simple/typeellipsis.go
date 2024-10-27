// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// NewEllipsisType implements [visitor.Environment].
func (env *Environment) NewEllipsisType() visitor.Type {
	return &Ellipsis{}
}

// Ellipsis represents a unit value.
type Ellipsis struct{}

// Ensure Ellipsis implements [visitor.Type].
var _ visitor.Type = (*Ellipsis)(nil)

// String implements [visitor.Type].
func (*Ellipsis) String() string {
	return "Ellipsis"
}
