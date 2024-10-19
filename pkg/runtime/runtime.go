// SPDX-License-Identifier: GPL-3.0-or-later

// Package runtime provides the core data structures and interfaces for the
// interpreter's runtime environment. It includes definitions for various
// value types (e.g., boolean, integer, float, string, unit) and the
// Value interface that all runtime values must implement. The package
// also defines the Callable interface for functions and lambdas, as well
// as built-in functions that can be used within the interpreter.
package runtime

import "fmt"

// Value is a generic value managed by the runtime.
type Value interface {
	String() string
}

// BoolValue represents a boolean value in the runtime.
type BoolValue struct {
	Value bool
}

// String returns the string representation of the boolean value.
func (v *BoolValue) String() string {
	return fmt.Sprintf("%t", v.Value)
}

// FloatValue represents a floating-point value in the runtime.
type FloatValue struct {
	Value float64
}

// String returns the string representation of the floating-point value.
func (v *FloatValue) String() string {
	return fmt.Sprintf("%f", v.Value)
}

// IntValue represents an integer value in the runtime.
type IntValue struct {
	Value int
}

// String returns the string representation of the integer value.
func (v *IntValue) String() string {
	return fmt.Sprintf("%d", v.Value)
}

// StringValue represents a string value in the runtime.
type StringValue struct {
	Value string
}

// String returns the string representation of the string value.
func (v *StringValue) String() string {
	return fmt.Sprintf("%q", v.Value)
}

// UnitValue represents a unit value in the runtime.
// A unit value is typically used to represent the absence of a value.
type UnitValue struct{}

// String returns the string representation of the unit value.
func (v *UnitValue) String() string {
	return "()"
}
