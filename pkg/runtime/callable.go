// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import "context"

// Callable is the common interface for callables types in the runtime.
type Callable interface {
	Call(ctx context.Context, env Environment, args ...Value) (Value, error)
}
