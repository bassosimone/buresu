// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

// NewUnitValue returns a new unit value. The unit value is
// equivalent to an empty list in this language.
func NewUnitValue() Value {
	return &ListValue{Car: nil, Cdr: nil}
}
