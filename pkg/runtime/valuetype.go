// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

// Value is a generic value managed by the runtime.
type Value interface {
	String() string
}
