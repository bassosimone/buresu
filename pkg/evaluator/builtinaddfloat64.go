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
		rtx.Assert(arg.Type() == "Float64", "expected Float64")
		sum += arg.(*Float64Value).Value
	}
	return NewFloat64Value(sum), nil
}
