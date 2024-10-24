// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/evaluator/visitor"

// Num is the numeric type class.
type Num interface {
	// Add sums this number with another number.
	Add(visitor.Value) (visitor.Value, error)

	// Mul multiplies this number with another number.
	Mul(visitor.Value) (visitor.Value, error)

	// We also implement the visitor.Value interface.
	visitor.Value
}
