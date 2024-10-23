// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInAndTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Bool"}, {Name: "Bool"}},
	ReturnType: typeannotation.Type{Name: "Bool"},
})

// BuiltInAnd is a built-in function that performs logical AND.
func BuiltInAnd(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 2, "expected 2 arguments")
	_, ok := args[0].(*BoolValue)
	rtx.Assert(ok, "expected Bool")
	_, ok = args[1].(*BoolValue)
	rtx.Assert(ok, "expected Bool")

	left := args[0].(*BoolValue).Value
	right := args[1].(*BoolValue).Value

	return NewBoolValue(left && right), nil
}
