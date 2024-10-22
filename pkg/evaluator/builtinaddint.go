// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInAddIntTypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Int", "Int"},
	ReturnType: "Int",
}

// BuiltInAddInt is a built-in function that adds integers.
func BuiltInAddInt(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// make sure we gracefully handle type checking bugs
	var sum int
	for _, arg := range args {
		if _, ok := arg.(*IntValue); !ok {
			return nil, fmt.Errorf("BUG: type error: expected Int, got %s", arg.Type())
		}
		sum += arg.(*IntValue).Value
	}
	return NewIntValue(sum), nil
}
