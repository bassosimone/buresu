// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// ErrWrongNumberOfArguments is the error returned when the number of arguments.
var ErrWrongNumberOfArguments = fmt.Errorf("wrong number of arguments")

// ErrWrongArgumentType is the error returned when the argument type is wrong.
var ErrWrongArgumentType = fmt.Errorf("wrong argument type")

// BuiltInFunc is a function that is built-in to the evaluator.
type BuiltInFunc func(ctx context.Context, args ...visitor.Value) (visitor.Value, error)

// BuiltInFuncValue is a built-in function value.
type BuiltInFuncValue struct {
	// Name is the name of the function.
	Name string

	// Fx is the function itself.
	Fx BuiltInFunc
}

// Ensure BuiltInFuncValue implements [visitor.Callable].
var _ visitor.Callable = (*BuiltInFuncValue)(nil)

// Call implements [visitor.Callable].
func (bf *BuiltInFuncValue) Call(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
	return bf.Fx(ctx, args...)
}

// Ensure BuiltInFuncValue implements [visitor.Value].
var _ visitor.Value = (*BuiltInFuncValue)(nil)

// String implements [visitor.Value].
func (bf *BuiltInFuncValue) String() string {
	return fmt.Sprintf("<built-in function %s>", bf.Name)
}
