// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/legacy/typeannotation"
)

var builtInOrTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Bool"}, {Name: "Bool"}},
	ReturnType: typeannotation.Type{Name: "Bool"},
})

// BuiltInOr is a built-in function that performs logical OR on boolean values.
func BuiltInOr(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 2, "expected 2 arguments")
	_, ok := args[0].(*BoolValue)
	rtx.Assert(ok, "expected Bool")
	_, ok = args[1].(*BoolValue)
	rtx.Assert(ok, "expected Bool")

	bool1 := args[0].(*BoolValue).Value
	bool2 := args[1].(*BoolValue).Value

	return NewBoolValue(bool1 || bool2), nil
}
