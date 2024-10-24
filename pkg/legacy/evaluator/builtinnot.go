// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/legacy/typeannotation"
)

var builtInNotTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Bool"}},
	ReturnType: typeannotation.Type{Name: "Bool"},
})

func BuiltInNot(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 1, "expected 1 argument")
	boolVal, ok := args[0].(*BoolValue)
	rtx.Assert(ok, "expected Bool")
	return NewBoolValue(!boolVal.Value), nil
}
