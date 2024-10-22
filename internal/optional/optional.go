// SPDX-License-Identifier: GPL-3.0-or-later

// Package optional implements the optional type.
package optional

import "github.com/bassosimone/buresu/internal/rtx"

// Value is the optional value.
type Value[T any] struct {
	present bool
	value   T
}

// Some constructs a new optional instance containing some value.
func Some[T any](value T) Value[T] {
	return Value[T]{present: true, value: value}
}

// None constructs a new empty optional instance.
func None[T any]() Value[T] {
	return Value[T]{present: false, value: *new(T)}
}

// IsSome returns whether there is some value in the optional.
func (v Value[T]) IsSome() bool {
	return v.present
}

// IsNone returns whether the optional is empty.
func (v Value[T]) IsNone() bool {
	return !v.present
}

// Unwrap returns the value, if present, otherwise it panics.
func (v Value[T]) Unwrap() T {
	rtx.Assert(v.present, "unwrap on empty optional")
	return v.value
}
