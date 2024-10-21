// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import "fmt"

// Value represents a value managed by the evaluator.
type Value interface {
	fmt.Stringer

	// Type returns the type name.
	Type() string
}
