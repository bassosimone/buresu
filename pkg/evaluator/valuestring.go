// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"encoding/json"
	"fmt"
)

// StringValue represents a string value.
//
// Construct using NewStringValue.
type StringValue struct {
	Value string
}

var _ Value = (*StringValue)(nil)

// NewStringValue creates a new [*StringValue] instance.
func NewStringValue(value string) *StringValue {
	return &StringValue{value}
}

// String returns the string representation of the string value.
func (v *StringValue) String() string {
	repr, _ := json.Marshal(v.Value)
	return fmt.Sprintf("%s", repr)
}

// Type implements Value.
func (*StringValue) Type() string {
	return "String"
}

var _ SequenceTrait = (*StringValue)(nil)

// Length returns the length of the string value.
func (v *StringValue) Length() int {
	return len(v.Value)
}
