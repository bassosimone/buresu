// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import "fmt"

// BoolValue represents a boolean value in the runtime.
type BoolValue struct {
	Value bool
}

// Ensure *BoolValue implements Value interface.
var _ Value = (*BoolValue)(nil)

// String returns the string representation of the boolean value.
func (v *BoolValue) String() string {
	return fmt.Sprintf("%t", v.Value)
}
