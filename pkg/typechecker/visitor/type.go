// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import "fmt"

// Type is the generic type returned by the typechecker.
//
// We use an ~empty interface, which allows each engine that implements
// the evaluator to define its own types.
type Type interface {
	fmt.Stringer
}
