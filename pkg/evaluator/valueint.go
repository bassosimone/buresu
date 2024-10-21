// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import "fmt"

// IntValue represents a int value.
//
// Construct using NewIntValue.
type IntValue struct {
	Value int
}

var _ Value = (*IntValue)(nil)

// NewIntValue creates a new [*IntValue] instance.
func NewIntValue(value int) *IntValue {
	return &IntValue{value}
}

// String returns the string representation of the integer value.
func (v *IntValue) String() string {
	return fmt.Sprintf("%d", v.Value)
}
