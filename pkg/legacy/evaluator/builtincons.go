// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/legacy/typeannotation"
)

var builtInConsTypeAnnotation = optional.Some(&typeannotation.Annotation{
	Params:     []typeannotation.Type{{Name: "Value"}, {Name: "Value"}},
	ReturnType: typeannotation.Type{Name: "Value"},
})

// BuiltInCons is a built-in function that creates cons cells.
func BuiltInCons(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(len(args) == 2, "expected 2 arguments")
	return NewConsCellValue(args[0], args[1]), nil
}
