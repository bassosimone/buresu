// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewFloat64Value implements [visitor.Environment].
func (env *Environment) NewFloat64Value(value float64) visitor.Value {
	return &Float64{value}
}

// Float64 represents a float64 value.
type Float64 struct {
	Value float64
}

// Ensure Float64 implements [visitor.Value].
var _ visitor.Value = (*Float64)(nil)

// String implements [visitor.Value].
func (v *Float64) String() string {
	return fmt.Sprintf("%f", v.Value)
}

// Ensure Float64 implements [Num].
var _ Num = (*Float64)(nil)

// Add implements Num.
func (v *Float64) Add(other visitor.Value) (Num, error) {
	num, ok := other.(*Float64)
	if !ok {
		return nil, ErrWrongArgumentType
	}
	return &Float64{v.Value + num.Value}, nil
}
