// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"fmt"
	"strings"
)

// ConsCellValue represents a cons cell value.
//
// Construct using NewConsCellValue.
type ConsCellValue struct {
	Car Value
	Cdr Value
}

var _ Value = (*ConsCellValue)(nil)

// NewConsCellValue creates a new [*ConsCellValue] instance.
func NewConsCellValue(car, cdr Value) *ConsCellValue {
	return &ConsCellValue{Car: car, Cdr: cdr}
}

// String returns a string representation of this ConsCell and all subsequent cells.
func (v *ConsCellValue) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "(")
	// we don't care about the error since we print it correctly anyway
	_ = v.visit(func(kind string, index int, value Value) {
		if index > 0 {
			fmt.Fprintf(&builder, " ")
		}
		fmt.Fprintf(&builder, "%s%s", kind, value.String())
	})
	fmt.Fprintf(&builder, ")")
	return builder.String()
}

// Type implements Value.
func (*ConsCellValue) Type() string {
	return "<cons cell>"
}

var _ SequenceTrait = (*ConsCellValue)(nil)

// Length returns the length of the cons cell.
func (v *ConsCellValue) Length() int {
	length := 0
	_ = v.visit(func(kind string, index int, value Value) {
		length++
	})
	return length
}

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
func (v *ConsCellValue) visit(visit func(kind string, index int, value Value)) error {
	index := 0
	for v != nil && v.Car != nil && v.Cdr != nil {
		visit(properListEntry, index, v.Car)
		switch cdr := v.Cdr.(type) {
		case *ConsCellValue:
			v = cdr
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
