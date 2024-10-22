// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInConsTypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Value", "Value"},
	ReturnType: "Value",
}

// BuiltInCons is a built-in function that creates a cons cells.
func BuiltInCons(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// we're protected by type checking
	return NewConsCellValue(args[0], args[1]), nil
}
