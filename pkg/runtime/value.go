// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"context"
	"fmt"
	"strings"
)

// Value is a generic value managed by the runtime.
type Value interface {
	String() string
}

// DisplayFunc implements the `display` built-in function.
var DisplayFunc = &BuiltInFuncValue{
	Name: "display",
	Fx: func(_ context.Context, env Environment, args ...Value) (Value, error) {
		var builder strings.Builder
		for idx, arg := range args {
			stringer, ok := arg.(fmt.Stringer)
			if !ok {
				return nil, fmt.Errorf("display: cannot convert argument to string: %v", arg)
			}
			fmt.Fprintf(&builder, "%s", stringer.String())
			if idx < len(args)-1 {
				builder.WriteString(" ")
			}
		}
		fmt.Fprintln(env.Output(), builder.String())
		return &UnitValue{}, nil
	},
}
