// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/evaluator/visitor"

// Seq is the sequence type class.
type Seq interface {
	// Length returns the length of the sequence.
	Length() (visitor.Value, error)
}
