// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// NewFloat64Type implements [visitor.Environment].
func (env *Environment) NewFloat64Type() visitor.Type {
	return &Float64{}
}

// Float64 represents a float64 type.
type Float64 struct{}

// Ensure Float64 implements [visitor.Type].
var _ visitor.Type = (*Float64)(nil)

// String implements [visitor.Type].
func (v *Float64) String() string {
	return "Float64"
}
