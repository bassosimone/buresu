// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"io"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

// NewGlobalEnvironment creates a new global [*Environment] instance
// with all the built-in functions and constants.
func NewGlobalEnvironment(output io.Writer) *Environment {
	global := NewEnvironment(output)

	_ = global.DefineValue("+", NewBuiltInFuncValue(
		"+",
		(&typeannotation.Annotation{Params: []string{"Int", "Int"}}).ArgumentsAnnotationPrefix(),
		BuiltInAddInt,
	))

	_ = global.DefineValue("+", NewBuiltInFuncValue(
		"+",
		(&typeannotation.Annotation{Params: []string{"Float64", "Float64"}}).ArgumentsAnnotationPrefix(),
		BuiltInAddFloat64,
	))

	_ = global.DefineValue("display", NewBuiltInFuncValue("display", "", BuiltInDisplay))

	return global
}
