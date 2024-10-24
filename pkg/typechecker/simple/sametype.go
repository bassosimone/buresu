// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import "github.com/bassosimone/buresu/pkg/typechecker/visitor"

// sameType returns whether two types have the same type.
func sameType(a, b visitor.Type) bool {
	// If either type is Any, then the types are considered the same
	if _, ok := a.(*Any); ok {
		return true
	}
	if _, ok := b.(*Any); ok {
		return true
	}

	// TODO(bassosimone): consider using a more robust form of comparison
	// than comparing the string representation of the types
	return a.String() == b.String()
}
