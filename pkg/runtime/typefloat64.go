// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
)

// Float64Value represents a float64 value in the runtime.
type Float64Value struct {
	Value float64
}

// Ensure *FloatValue implements Value interface.
var _ Value = (*Float64Value)(nil)

// String returns the string representation of the floating-point value.
func (v *Float64Value) String() string {
	return fmt.Sprintf("%f", v.Value)
}

// Float64SumFunc implements the `__floatSum` built-in function for integers.
var Float64SumFunc = &BuiltInFuncValue{
	Name: "__float64Sum",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("__float64Sum: expected 2 arguments, got %d", len(args))
		}

		a, ok := args[0].(*Float64Value)
		if !ok {
			return nil, fmt.Errorf("__float64Sum: expected float, got %T", args[0])
		}

		b, ok := args[1].(*Float64Value)
		if !ok {
			return nil, fmt.Errorf("__float64Sum: expected float, got %T", args[1])
		}

		return &Float64Value{Value: a.Value + b.Value}, nil
	},
}
