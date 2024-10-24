// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/legacy/typeannotation"
)

var builtInFormatTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Value"}},
	ReturnType: typeannotation.Type{Name: "String"},
})

// BuiltInFormat is a built-in function that formats its unique argument.
func BuiltInFormat(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(env != nil, "environment must not be nil")
	rtx.Assert(env.output != nil, "environment output must not be nil")
	rtx.Assert(len(args) == 1, "expected exactly one argument")
	return NewStringValue(args[0].String()), nil
}
