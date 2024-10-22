// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInAddIntTypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Int", "Int"},
	ReturnType: "Int",
}

// BuiltInAddInt is a built-in function that adds integers.
func BuiltInAddInt(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	var sum int
	for _, arg := range args {
		sum += arg.(*IntValue).Value // we're protected by type checking
	}
	return NewIntValue(sum), nil
}
