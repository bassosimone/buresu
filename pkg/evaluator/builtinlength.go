// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInLengthTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Sequence"}},
	ReturnType: typeannotation.Type{Name: "Int"},
})

// BuiltInLength is a built-in function that returns the length of a sequence.
func BuiltInLength(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 1, "expected 1 argument")
	rtx.Assert(args[0].Type() == "Sequence", "expected Sequence")
	return NewIntValue(args[0].(SequenceTrait).Length()), nil
}
