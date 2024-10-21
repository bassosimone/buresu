// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"
	"strings"
)

// BuiltInDisplay is a built-in function that displays its argument.
func BuiltInDisplay(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	var builder strings.Builder
	for idx, arg := range args {
		fmt.Fprintf(&builder, "%s", arg.String())
		if idx < len(args)-1 {
			builder.WriteString(" ")
		}
	}
	fmt.Fprintln(env.output, builder.String())
	return NewUnitValue(), nil
}
