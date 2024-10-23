// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import "fmt"

// Float64Value represents a float64 value.
//
// Construct using NewFloat64Value.
type Float64Value struct {
	Value float64
}

var _ Value = (*Float64Value)(nil)

// NewFloat64Value creates a new [*Float64Value] instance.
func NewFloat64Value(value float64) *Float64Value {
	return &Float64Value{value}
}

// String returns the string representation of the floating-point value.
func (v *Float64Value) String() string {
	return fmt.Sprintf("%f", v.Value)
}

// Type implements Value.
func (*Float64Value) Type() string {
	return "Float64"
}
