// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

// SequenceTrait is the trait shared by all sequences.
type SequenceTrait interface {
	// Length returns the length of the sequence.
	Length() int

	// A Sequence is also a Value.
	Value
}
