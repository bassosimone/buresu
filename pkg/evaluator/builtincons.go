// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInConsTypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Value", "Value"},
	ReturnType: "Value",
}

// BuiltInCons is a built-in function that creates a cons cells.
func BuiltInCons(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// make sure we gracefully handle type checking bugs
	if len(args) != 2 {
		return nil, fmt.Errorf("BUG: expected 2 arguments, got %d", len(args))
	}
	return NewConsCellValue(args[0], args[1]), nil
}
