// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import "context"

// CallableTrait is the trait shared by all callables.
type CallableTrait interface {
	Call(ctx context.Context, env *Environment, args ...Value) (Value, error)
}
