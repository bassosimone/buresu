// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
	"strings"
)

// ListValue is an ordered sequence of values that
// may belong to different types. In other words, this
// is a Lisp list, not a Python list.
type ListValue struct {
	// Car is the first element of the list.
	Car Value

	// Cdr is the rest of the list, which may be nil.
	Cdr *ListValue
}

// Make sure that ListValue implements the Value interface.
var _ Value = (*ListValue)(nil)

// String returns a string representation of the list.
func (lv *ListValue) String() string {
	var buffer strings.Builder
	buffer.WriteString("(")
	for ; lv != nil; lv = lv.Cdr {
		if lv.Car != nil {
			buffer.WriteString(lv.Car.String())
		}
		if lv.Cdr != nil {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

// CarFunc implements the `car` built-in function.
// The `car` function returns the first element of a list.
// If the input is a unit value, it returns a unit value.
// If the input is not a list, it returns an error.
var CarFunc = &BuiltInFuncValue{
	Name: "car",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("car: expected 1 argument, got %d", len(args))
		}

		lv, ok := args[0].(*ListValue)
		if !ok {
			return nil, fmt.Errorf("car: expected *ListValue, got %T", args[0])
		}

		if lv.Car == nil {
			return NewUnitValue(), nil
		}

		return lv.Car, nil
	},
}

// CdrFunc implements the `cdr` built-in function.
// The `cdr` function returns the rest of the list after the first element.
// If the input is a unit value, it returns a unit value.
// If the input is not a list, it returns an error.
var CdrFunc = &BuiltInFuncValue{
	Name: "cdr",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("cdr: expected 1 argument, got %d", len(args))
		}

		lv, ok := args[0].(*ListValue)
		if !ok {
			return nil, fmt.Errorf("cdr: expected *ListValue, got %T", args[0])
		}

		if lv.Cdr == nil {
			return NewUnitValue(), nil
		}

		return lv.Cdr, nil
	},
}

// ListFunc implements the `list` built-in function.
// The `list` function constructs a list from the provided arguments.
// If no arguments are provided, it returns a unit value.
var ListFunc = &BuiltInFuncValue{
	Name: "list",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) < 1 {
			return NewUnitValue(), nil
		}

		lval := &ListValue{Car: args[0], Cdr: nil}
		for idx, lp := 1, lval; idx < len(args); lp, idx = lp.Cdr, idx+1 {
			lp.Cdr = &ListValue{Car: args[idx], Cdr: nil}
		}

		return lval, nil
	},
}
