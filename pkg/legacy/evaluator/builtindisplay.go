// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"
	"strings"

	"github.com/bassosimone/buresu/internal/optional"
	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/legacy/typeannotation"
)

var builtInDisplayTypeAnnotation = optional.None[*typeannotation.Annotation]()

// BuiltInDisplay is a built-in function that displays its argument.
func BuiltInDisplay(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	rtx.Assert(env != nil, "environment must not be nil")
	rtx.Assert(env.output != nil, "environment output must not be nil")

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
