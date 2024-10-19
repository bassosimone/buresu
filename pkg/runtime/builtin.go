// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
	"strings"
)

// displayFunc implements the `display` built-in function.
var displayFunc = &BuiltInFuncValue{
	Name: "display",
	Fx: func(_ context.Context, env Environment, args ...Value) (Value, error) {
		var builder strings.Builder
		for idx, arg := range args {
			stringer, ok := arg.(fmt.Stringer)
			if !ok {
				return nil, fmt.Errorf("display: cannot convert argument to string: %v", arg)
			}
			fmt.Fprintf(&builder, "%s", stringer.String())
			if idx < len(args)-1 {
				builder.WriteString(" ")
			}
		}
		fmt.Fprintln(env.Output(), builder.String())
		return &UnitValue{}, nil
	},
}

// intSumFunc implements the `__intSum` built-in function for integers.
var intSumFunc = &BuiltInFuncValue{
	Name: "__intSum",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("sumInt: expected 2 arguments, got %d", len(args))
		}

		a, ok := args[0].(*IntValue)
		if !ok {
			return nil, fmt.Errorf("sumInt: expected integer, got %T", args[0])
		}

		b, ok := args[1].(*IntValue)
		if !ok {
			return nil, fmt.Errorf("sumInt: expected integer, got %T", args[1])
		}

		return &IntValue{Value: a.Value + b.Value}, nil
	},
}

// floatSumFunc implements the `__floatSum` built-in function for integers.
var floatSumFunc = &BuiltInFuncValue{
	Name: "__floatSum",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("sumFloat: expected 2 arguments, got %d", len(args))
		}

		a, ok := args[0].(*FloatValue)
		if !ok {
			return nil, fmt.Errorf("sumFloat: expected float, got %T", args[0])
		}

		b, ok := args[1].(*FloatValue)
		if !ok {
			return nil, fmt.Errorf("sumFloat: expected float, got %T", args[1])
		}

		return &FloatValue{Value: a.Value + b.Value}, nil
	},
}

// intLtFunc implements the `__intLt` built-in function for integers.
var intLtFunc = &BuiltInFuncValue{
	Name: "__intLt",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("sumInt: expected 2 arguments, got %d", len(args))
		}

		a, ok := args[0].(*IntValue)
		if !ok {
			return nil, fmt.Errorf("sumInt: expected integer, got %T", args[0])
		}

		b, ok := args[1].(*IntValue)
		if !ok {
			return nil, fmt.Errorf("sumInt: expected integer, got %T", args[1])
		}

		return &BoolValue{Value: a.Value < b.Value}, nil
	},
}
