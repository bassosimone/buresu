// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
)

// IntValue represents an integer value in the runtime.
type IntValue struct {
	Value int
}

// Ensure *IntValue implements Value interface.
var _ Value = (*IntValue)(nil)

// String returns the string representation of the integer value.
func (v *IntValue) String() string {
	return fmt.Sprintf("%d", v.Value)
}

// IntSumFunc implements the `__intSum` built-in function for integers.
var IntSumFunc = &BuiltInFuncValue{
	Name: "__intSum",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("__intSum: expected 2 arguments, got %d", len(args))
		}

		a, ok := args[0].(*IntValue)
		if !ok {
			return nil, fmt.Errorf("__intSum: expected integer, got %T", args[0])
		}

		b, ok := args[1].(*IntValue)
		if !ok {
			return nil, fmt.Errorf("__intSum: expected integer, got %T", args[1])
		}

		return &IntValue{Value: a.Value + b.Value}, nil
	},
}

// IntLtFunc implements the `__intLt` built-in function for integers.
var IntLtFunc = &BuiltInFuncValue{
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
