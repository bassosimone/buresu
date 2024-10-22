// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"io"
)

// NewGlobalEnvironment creates a new global [*Environment] instance
// with all the built-in functions and constants.
func NewGlobalEnvironment(output io.Writer) *Environment {
	global := NewEnvironment(output)

	global.SetImplements("Bool", "Value")
	global.SetImplements("ConsCell", "Value")
	global.SetImplements("Float64", "Value")
	global.SetImplements("Int", "Value")
	global.SetImplements("Lambda", "Value")
	global.SetImplements("Quoted", "Value")
	global.SetImplements("String", "Value")
	global.SetImplements("Unit", "Value")

	global.SetImplements("String", "Sequence")
	global.SetImplements("ConsCell", "Sequence")
	global.SetImplements("Unit", "Sequence")

	_ = global.DefineValue("+", NewBuiltInFuncValue(
		"+",
		builtInAddIntTypeAnnotation.String(),
		BuiltInAddInt,
	))

	_ = global.DefineValue("+", NewBuiltInFuncValue(
		"+",
		builtInAddFloat64TypeAnnotation.String(),
		BuiltInAddFloat64,
	))

	_ = global.DefineValue("cons", NewBuiltInFuncValue(
		"cons",
		builtInConsTypeAnnotation.String(),
		BuiltInCons,
	))

	_ = global.DefineValue("display", NewBuiltInFuncValue("display", "", BuiltInDisplay))

	_ = global.DefineValue("length", NewBuiltInFuncValue(
		"length",
		builtInLengthTypeAnnotation.String(),
		BuiltInLength,
	))

	return global
}
