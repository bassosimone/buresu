// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
)

// BuiltInAddFloat64 is a built-in function that adds floa64.
func BuiltInAddFloat64(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	var sum float64
	for _, arg := range args {
		sum += arg.(*Float64Value).Value // we're protected by type checking
	}
	return NewFloat64Value(sum), nil
}
