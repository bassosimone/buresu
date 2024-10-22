// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInLengthTypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Sequence"},
	ReturnType: "Int",
}

// BuiltInLength is a built-in function that creates a cons cells.
func BuiltInLength(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// we're protected by type checking
	return NewIntValue(args[0].(SequenceTrait).Length()), nil
}
