// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import "context"

// CallableTrait is the trait shared by all callables.
type CallableTrait interface {
	// Call is the method that will be called when the callable is invoked.
	Call(ctx context.Context, env *Environment, args ...Value) (Value, error)

	// TypeAnnotationPrefix returns the type annotation prefix of the callable.
	TypeAnnotationPrefix() string

	// A callable is also a Value.
	Value
}
