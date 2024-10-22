// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInAddIntTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Int"}, {Name: "Int"}},
	ReturnType: typeannotation.Type{Name: "Int"},
})

// BuiltInAddInt is a built-in function that adds integers.
func BuiltInAddInt(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	var sum int
	for _, arg := range args {
		rtx.Assert(arg.Type() == "Int", "expected Int")
		sum += arg.(*IntValue).Value
	}
	return NewIntValue(sum), nil
}
