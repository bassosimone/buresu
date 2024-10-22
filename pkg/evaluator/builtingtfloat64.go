// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInGtFloat64TypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Float64"}, {Name: "Float64"}},
	ReturnType: typeannotation.Type{Name: "Bool"},
})

// BuiltInGtFloat64 is a built-in function that compares two float64 values.
func BuiltInGtFloat64(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 2, "expected 2 arguments")
	_, ok := args[0].(*Float64Value)
	rtx.Assert(ok, "expected Float64")
	_, ok = args[1].(*Float64Value)
	rtx.Assert(ok, "expected Float64")

	left := args[0].(*Float64Value).Value
	right := args[1].(*Float64Value).Value

	return NewBoolValue(left > right), nil
}
