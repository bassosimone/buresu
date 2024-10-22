// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

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
