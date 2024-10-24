// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/evaluator/visitor"

// Ord is the ordered type class.
type Ord interface {
	// Gt returns true if the receiver is greater than the argument.
	Gt(visitor.Value) (visitor.Value, error)

	// Lt returns true if the receiver is less than the argument.
	Lt(visitor.Value) (visitor.Value, error)

	// We also implement the visitor.Value interface.
	visitor.Value
}
