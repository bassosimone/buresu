// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/legacy/typeannotation"
)

var builtInLtFloat64TypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Float64"}, {Name: "Float64"}},
	ReturnType: typeannotation.Type{Name: "Bool"},
})

// BuiltInLtFloat64 is a built-in function that compares two float64 values.
func BuiltInLtFloat64(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 2, "expected 2 arguments")
	_, ok := args[0].(*Float64Value)
	rtx.Assert(ok, "expected Float64")
	_, ok = args[1].(*Float64Value)
	rtx.Assert(ok, "expected Float64")

	left := args[0].(*Float64Value).Value
	right := args[1].(*Float64Value).Value

	return NewBoolValue(left < right), nil
}

var builtInLtIntTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Int"}, {Name: "Int"}},
	ReturnType: typeannotation.Type{Name: "Bool"},
})

// BuiltInLtInt is a built-in function that compares two integers.
func BuiltInLtInt(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 2, "expected 2 arguments")
	_, ok := args[0].(*IntValue)
	rtx.Assert(ok, "expected Int")
	_, ok = args[1].(*IntValue)
	rtx.Assert(ok, "expected Int")

	return NewBoolValue(args[0].(*IntValue).Value < args[1].(*IntValue).Value), nil
}
