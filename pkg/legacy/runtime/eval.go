// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
)

// EvalFunc implements the `eval` built-in function.
var EvalFunc = &BuiltInFuncValue{
	Name: "eval",
	Fx: func(ctx context.Context, env Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("eval: expected 1 argument, got %d", len(args))
		}
		expr, ok := args[0].(*QuotedValue)
		if !ok {
			return nil, fmt.Errorf("eval: expected a *QuotedValue, got %T", args[0])
		}
		return env.Eval(ctx, expr.Value)
	},
}
