// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"io"
)

// NewGlobalEnvironment creates a new global [*Environment] instance
// with all the built-in functions and constants.
func NewGlobalEnvironment(output io.Writer) *Environment {
	global := NewEnvironment(output)

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

	_ = global.DefineValue("display", NewBuiltInFuncValue("display", "", BuiltInDisplay))

	return global
}
