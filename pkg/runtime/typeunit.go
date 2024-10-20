// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

// UnitValue represents a unit value in the runtime.
// A unit value is typically used to represent the absence of a value.
type UnitValue struct{}

// Ensure *UnitValue implements Value interface.
var _ Value = (*UnitValue)(nil)

// String returns the string representation of the unit value.
func (v *UnitValue) String() string {
	return "()"
}
