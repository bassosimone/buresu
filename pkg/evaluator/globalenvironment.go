// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"io"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

// NewGlobalEnvironment creates a new global [*Environment] instance
// with all the built-in functions and constants.
func NewGlobalEnvironment(output io.Writer) *Environment {
	global := NewEnvironment(output)

	rtx.Assert(global.SetImplements("Bool", "Value"), "failed to set Bool implements Value")
	rtx.Assert(global.SetImplements("ConsCell", "Value"), "failed to set ConsCell implements Value")
	rtx.Assert(global.SetImplements("Float64", "Value"), "failed to set Float64 implements Value")
	rtx.Assert(global.SetImplements("Int", "Value"), "failed to set Int implements Value")
	rtx.Assert(global.SetImplements("Lambda", "Value"), "failed to set Lambda implements Value")
	rtx.Assert(global.SetImplements("Quoted", "Value"), "failed to set Quoted implements Value")
	rtx.Assert(global.SetImplements("String", "Value"), "failed to set String implements Value")
	rtx.Assert(global.SetImplements("Unit", "Value"), "failed to set Unit implements Value")

	rtx.Assert(global.SetImplements("String", "Sequence"), "failed to set String implements Sequence")
	rtx.Assert(global.SetImplements("ConsCell", "Sequence"), "failed to set ConsCell implements Sequence")
	rtx.Assert(global.SetImplements("Unit", "Sequence"), "failed to set Unit implements Sequence")

	rtx.Must(defineBuiltIn(global, "+", builtInAddIntTypeAnnotation, BuiltInAddInt))
	rtx.Must(defineBuiltIn(global, "+", builtInAddFloat64TypeAnnotation, BuiltInAddFloat64))
	rtx.Must(defineBuiltIn(global, "<", builtInLtIntTypeAnnotation, BuiltInLtInt))
	rtx.Must(defineBuiltIn(global, "<", builtInLtFloat64TypeAnnotation, BuiltInLtFloat64))
	rtx.Must(defineBuiltIn(global, ">", builtInGtIntTypeAnnotation, BuiltInGtInt))
	rtx.Must(defineBuiltIn(global, ">", builtInGtFloat64TypeAnnotation, BuiltInGtFloat64))
	rtx.Must(defineBuiltIn(global, "and", builtInAndTypeAnnotation, BuiltInAnd))
	rtx.Must(defineBuiltIn(global, "cons", builtInConsTypeAnnotation, BuiltInCons))
	rtx.Must(defineBuiltIn(global, "length", builtInLengthTypeAnnotation, BuiltInLength))
	rtx.Must(defineBuiltIn(global, "display", builtInDisplayTypeAnnotation, BuiltInDisplay))
	rtx.Must(defineBuiltIn(global, "not", builtInNotTypeAnnotation, BuiltInNot))
	rtx.Must(defineBuiltIn(global, "or", builtInOrTypeAnnotation, BuiltInOr))

	return global
}

// defineBuiltIn is a helper function to define a built-in function in the environment.
func defineBuiltIn(env *Environment, name string,
	annotation optional.Value[*typeannotation.Annotation], fx BuiltInFunc) error {
	var annotationString string
	if annotation.IsSome() {
		annotationString = annotation.Unwrap().String()
	}
	return env.DefineValue(name, NewBuiltInFuncValue(name, annotationString, fx))
}
