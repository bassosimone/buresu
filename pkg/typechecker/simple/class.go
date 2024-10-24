// SPDX-License-Identifier: GPL-3.0-or-later

package simple

// Class is a class to which types belong.
type Class interface {
	// Instantiate instantiates the type class in the given environment.
	Instantiate(env *Environment)
}
