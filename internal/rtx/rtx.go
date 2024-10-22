// SPDX-License-Identifier: GPL-3.0-or-later

package rtx

import "fmt"

// Must panics if the provided error is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Assert panics if the provided condition is false.
func Assert(condition bool, message string) {
	if !condition {
		panic(fmt.Errorf("assertion failed: %s", message))
	}
}
