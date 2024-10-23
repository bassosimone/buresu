// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
	"slices"
	"strings"
)

// ConsCell represents a pair of values, similar to a node in a linked list.
// It is used to build lists by chaining these pairs together.
type ConsCell struct {
	// car is the first value in the pair.
	// When building lists, this is the current element in the list.
	car Value

	// cdr is the second value in the pair.
	// When building lists, this is the pointer to the rest of the list, also a value.
	cdr Value
}

// NewConsCell creates a new ConsCell with the given car and cdr values.
func NewConsCell(car, cdr Value) *ConsCell {
	return &ConsCell{car: car, cdr: cdr}
}

// ConsFunc is a built-in function that constructs a pair (ConsCell) from two provided arguments.
var ConsFunc = &BuiltInFuncValue{
	Name: "cons",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("cons: expects exactly 2 arguments, got %d", len(args))
		}
		if args[0] == nil || args[1] == nil {
			return nil, fmt.Errorf("cons: nil values are not allowed")
		}
		return NewConsCell(args[0], args[1]), nil
	},
}

// Car returns the first value (car) in the ConsCell.
func (cc *ConsCell) Car() Value {
	return cc.car
}

// CarFunc is a built-in function that returns the first element (car) of a ConsCell.
var CarFunc = &BuiltInFuncValue{
	Name: "car",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("car: expected 1 argument, got %d", len(args))
		}

		cell, ok := args[0].(*ConsCell)
		if !ok {
			return nil, fmt.Errorf("car: expected *ConsCell, got %T", args[0])
		}

		return cell.Car(), nil
	},
}

// Cdr returns the second value (cdr) in the ConsCell.
func (cc *ConsCell) Cdr() Value {
	return cc.cdr
}

// CdrFunc is a built-in function that returns the second element (cdr) of a ConsCell.
var CdrFunc = &BuiltInFuncValue{
	Name: "cdr",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("cdr: expected 1 argument, got %d", len(args))
		}

		cell, ok := args[0].(*ConsCell)
		if !ok {
			return nil, fmt.Errorf("cdr: expected *ConsCell, got %T", args[0])
		}

		return cell.Cdr(), nil
	},
}

// Ensure that ConsCell implements the Value interface.
var _ Value = (*ConsCell)(nil)

const (
	// properListEntry indicates that this value is part of a proper list.
	properListEntry = ""

	// improperListEntry indicates that we've hit an improper list.
	improperListEntry = ". "
)

// ErrImproperList is returned when a list is not proper.
var ErrImproperList = fmt.Errorf("improper list")

// ErrMalformedList is returned when a list is malformed.
var ErrMalformedList = fmt.Errorf("malformed list")

// visit calls the provided function for each element in the cons list.
func (cc *ConsCell) visit(visit func(kind string, index int, value Value)) error {
	index := 0
	for cc != nil && cc.car != nil && cc.cdr != nil {
		visit(properListEntry, index, cc.car)
		switch cdr := cc.cdr.(type) {
		case *ConsCell:
			cc = cdr
			index++
		case *UnitValue:
			return nil
		default:
			index++
			visit(improperListEntry, index, cdr)
			return ErrImproperList
		}
	}
	return ErrMalformedList
}

// String returns a string representation of this ConsCell and all subsequent cells.
func (cc *ConsCell) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "(")
	// we don't care about the error since we print it correctly anyway
	_ = cc.visit(func(kind string, index int, value Value) {
		if index > 0 {
			fmt.Fprintf(&builder, " ")
		}
		fmt.Fprintf(&builder, "%s%s", kind, value.String())
	})
	fmt.Fprintf(&builder, ")")
	return builder.String()
}

// NewConsList creates a list of ConsCells from the provided values.
// If no values are provided, it returns a unit value.
func NewConsList(values ...Value) Value {
	if len(values) < 1 {
		return NewUnitValue()
	}
	front := NewConsCell(values[0], NewUnitValue())
	back := front
	for idx := 1; idx < len(values); idx++ {
		back.cdr = NewConsCell(values[idx], NewUnitValue())
		back = back.cdr.(*ConsCell)
	}
	return front
}

// ListFunc is a built-in function that constructs a list from the provided arguments.
// If no arguments are provided, it returns a unit value.
var ListFunc = &BuiltInFuncValue{
	Name: "list",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) < 1 {
			return NewUnitValue(), nil
		}
		return NewConsList(args...), nil
	},
}

// NullFunc is a built-in function that checks if the provided argument is a unit value (null).
var NullFunc = &BuiltInFuncValue{
	Name: "null?",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("null?: expected 1 argument, got %d", len(args))
		}
		switch args[0].(type) {
		case *UnitValue:
			return &BoolValue{Value: true}, nil
		default:
			return &BoolValue{Value: false}, nil
		}
	},
}

// Length returns the length of the list starting from this ConsCell.
func (cc *ConsCell) Length() (int, error) {
	length := 0
	err := cc.visit(func(_ string, _ int, value Value) {
		length++
	})
	return length, err
}

// LengthFunc is a built-in function that returns the length of the list.
var LengthFunc = &BuiltInFuncValue{
	Name: "length",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("length: expected 1 argument, got %d", len(args))
		}
		switch v := args[0].(type) {
		case *UnitValue:
			return &IntValue{Value: 0}, nil
		case *ConsCell:
			value, err := v.Length()
			return &IntValue{Value: value}, err
		default:
			return nil, fmt.Errorf("length: expected *ConsCell, got %T", args[0])
		}
	},
}

// collect gathers all values in the list starting from this ConsCell into a slice.
func (cc *ConsCell) collect() ([]Value, error) {
	var values []Value
	err := cc.visit(func(_ string, _ int, value Value) {
		values = append(values, value)
	})
	return values, err
}

// AppendConsLists appends the provided lists together into a new list.
// If no lists are provided, it returns a unit value.
// If any of the provided values is not a list, it returns an error.
func AppendConsLists(lists ...Value) (Value, error) {
	if len(lists) < 1 {
		return NewUnitValue(), nil
	}
	var values []Value
	for _, entry := range lists {
		switch entry := entry.(type) {
		case *UnitValue:
			continue
		case *ConsCell:
			additionalValues, err := entry.collect()
			if err != nil {
				return nil, err
			}
			values = append(values, additionalValues...)
		default:
			return nil, fmt.Errorf("append: expected a *ConsCell, got %T", entry)
		}
	}
	return NewConsList(values...), nil
}

// AppendFunc is a built-in function that appends multiple lists together.
var AppendFunc = &BuiltInFuncValue{
	Name: "append",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		return AppendConsLists(args...)
	},
}

// Reverse reverses the cons list starting at this ConsCell.
func (cc *ConsCell) Reverse() (Value, error) {
	values, err := cc.collect()
	if err != nil {
		return nil, err
	}
	// note: if len(values) < 1, NewConsList will return a unit value
	slices.Reverse(values)
	return NewConsList(values...), nil
}

// ReverseFunc is a built-in function that reverses a list.
var ReverseFunc = &BuiltInFuncValue{
	Name: "reverse",
	Fx: func(_ context.Context, _ Environment, args ...Value) (Value, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("reverse: expected 1 argument, got %d", len(args))
		}
		switch v := args[0].(type) {
		case *UnitValue:
			return NewUnitValue(), nil
		case *ConsCell:
			return v.Reverse()
		default:
			return nil, fmt.Errorf("reverse: expected *ConsCell, got %T", args[0])
		}
	},
}
