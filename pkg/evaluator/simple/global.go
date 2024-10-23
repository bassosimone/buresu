// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"io"

	"github.com/bassosimone/buresu/internal/rtx"
)

// NewGlobalEnvironment creates a new global environment.
func NewGlobalEnvironment(writer io.Writer) *Environment {
	env := NewEnvironment()

	builtins := []*BuiltInFuncValue{
		NewBuiltInAdd(),
		NewBuiltInDisplay(writer),
	}
	for _, builtin := range builtins {
		rtx.Must(env.DefineValue(builtin.Name, builtin))
	}

	return env
}
