// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInLengthTypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Sequence"},
	ReturnType: "Int",
}

// BuiltInLength is a built-in function that creates a cons cells.
func BuiltInLength(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// make sure we gracefully handle type checking bugs
	if len(args) != 1 {
		return nil, fmt.Errorf("BUG: expected 1 argument, got %d", len(args))
	}
	if _, ok := args[0].(SequenceTrait); !ok {
		return nil, fmt.Errorf("BUG: type error: expected Sequence, got %s", args[0].Type())
	}
	return NewIntValue(args[0].(SequenceTrait).Length()), nil
}
