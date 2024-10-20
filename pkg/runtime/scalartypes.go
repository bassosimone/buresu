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

// FloatValue represents a floating-point value in the runtime.
type FloatValue struct {
	Value float64
}

// Ensure *FloatValue implements Value interface.
var _ Value = (*FloatValue)(nil)

// String returns the string representation of the floating-point value.
func (v *FloatValue) String() string {
	return fmt.Sprintf("%f", v.Value)
}

// IntValue represents an integer value in the runtime.
type IntValue struct {
	Value int
}

// Ensure *IntValue implements Value interface.
var _ Value = (*IntValue)(nil)

// String returns the string representation of the integer value.
func (v *IntValue) String() string {
	return fmt.Sprintf("%d", v.Value)
}

// StringValue represents a string value in the runtime.
type StringValue struct {
	Value string
}

// Ensure *StringValue implements Value interface.
var _ Value = (*StringValue)(nil)

// String returns the string representation of the string value.
func (v *StringValue) String() string {
	return fmt.Sprintf("%q", v.Value)
}

// UnitValue represents a unit value in the runtime.
// A unit value is typically used to represent the absence of a value.
type UnitValue struct{}

// Ensure *UnitValue implements Value interface.
var _ Value = (*UnitValue)(nil)

// String returns the string representation of the unit value.
func (v *UnitValue) String() string {
	return "()"
}
