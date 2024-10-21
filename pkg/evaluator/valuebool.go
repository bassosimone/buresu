// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import "fmt"

// BoolValue represents a boolean value.
//
// Construct using NewBoolValue.
type BoolValue struct {
	Value bool
}

var _ Value = (*BoolValue)(nil)

// NewBoolValue creates a new [*BoolValue] instance.
func NewBoolValue(value bool) *BoolValue {
	return &BoolValue{value}
}

// String returns the string representation of the boolean value.
func (v *BoolValue) String() string {
	return fmt.Sprintf("%t", v.Value)
}
