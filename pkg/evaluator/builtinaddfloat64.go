// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

var builtInAddFloat64TypeAnnotation = &typeannotation.Annotation{
	Params:     []string{"Float64", "Float64"},
	ReturnType: "Float64",
}

// BuiltInAddFloat64 is a built-in function that adds floa64.
func BuiltInAddFloat64(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// make sure we gracefully handle type checking bugs
	var sum float64
	for _, arg := range args {
		if _, ok := arg.(*Float64Value); !ok {
			return nil, fmt.Errorf("BUG: type error: expected Float64, got %s", arg.Type())
		}
		sum += arg.(*Float64Value).Value
	}
	return NewFloat64Value(sum), nil
}
