// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import "fmt"

// Value is the generic value returned by the evaluator.
//
// We use an ~empty interface, which allows each engine that implements
// the evaluator to define its own types.
type Value interface {
	fmt.Stringer
}
