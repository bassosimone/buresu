// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

// UnitValue represents a unit value.
//
// Construct using NewUnitValue.
type UnitValue struct{}

var _ Value = (*UnitValue)(nil)

// NewUnitValue creates a new UnitValue instance.
func NewUnitValue() *UnitValue {
	return &UnitValue{}
}

// String implements Value.
func (*UnitValue) String() string {
	return "()"
}

// Type implements Value.
func (*UnitValue) Type() string {
	return "Unit"
}

var _ SequenceTrait = (*UnitValue)(nil)

// Length implements SequenceTrait.
func (*UnitValue) Length() int {
	return 0
}
