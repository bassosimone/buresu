// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
)

// BuiltInFuncValue is a function that is built-in in the runtime.
//
// The zero value is not ready for use.
type BuiltInFuncValue struct {
	// Name is the name of the built-in function.
	Name string

	// Fx is the actual function that implements the built-in function.
	Fx func(ctx context.Context, env Environment, args ...Value) (Value, error)
}

var (
	_ Callable = (*BuiltInFuncValue)(nil)
	_ Value    = (*BuiltInFuncValue)(nil)
)

// String returns the `<builtin: %p>` string representation of the built-in function.
func (fx *BuiltInFuncValue) String() string {
	return fmt.Sprintf("<builtin: %s>", fx.Name)
}

// Call calls the built-in function with the given arguments.
func (fx *BuiltInFuncValue) Call(ctx context.Context, env Environment, args ...Value) (Value, error) {
	return fx.Fx(ctx, env, args...)
}
