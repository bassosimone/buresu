// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewIntValue implements [visitor.Environment].
func (env *Environment) NewIntValue(value int) visitor.Value {
	return &Int{value}
}

// Int represents a int value.
type Int struct {
	Value int
}

// Ensure Int implements [visitor.Value].
var _ visitor.Value = (*Int)(nil)

// String implements [visitor.Value].
func (v *Int) String() string {
	return fmt.Sprintf("%d", v.Value)
}

// Ensure Int implements [Num].
var _ Num = (*Int)(nil)

// Add implements Num.
func (v *Int) Add(other visitor.Value) (Num, error) {
	num, ok := other.(*Int)
	if !ok {
		return nil, ErrWrongArgumentType
	}
	return &Int{v.Value + num.Value}, nil
}
