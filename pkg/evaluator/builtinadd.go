// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInAddFloat64TypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Float64"}, {Name: "Float64"}},
	ReturnType: typeannotation.Type{Name: "Float64"},
})

// BuiltInAddFloat64 is a built-in function that adds float64.
func BuiltInAddFloat64(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	var sum float64
	for _, arg := range args {
		_, ok := arg.(*Float64Value)
		rtx.Assert(ok, "expected Float64")
		sum += arg.(*Float64Value).Value
	}
	return NewFloat64Value(sum), nil
}

var builtInAddIntTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Int"}, {Name: "Int"}},
	ReturnType: typeannotation.Type{Name: "Int"},
})

// BuiltInAddInt is a built-in function that adds integers.
func BuiltInAddInt(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	var sum int
	for _, arg := range args {
		_, ok := arg.(*IntValue)
		rtx.Assert(ok, "expected Int")
		sum += arg.(*IntValue).Value
	}
	return NewIntValue(sum), nil
}
