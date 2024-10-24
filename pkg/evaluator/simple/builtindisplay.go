// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewBuiltInDisplay creates a new built-in function that displays its arguments.
func NewBuiltInDisplay(writer io.Writer) *BuiltInFuncValue {
	return &BuiltInFuncValue{
		Name: "display",
		Fx: func(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
			var buffer strings.Builder
			for idx, arg := range args {
				fmt.Printf("%s", arg.String())
				if idx < len(args)-1 {
					fmt.Printf(" ")
				}
			}
			_, err := fmt.Fprintln(writer, buffer.String())
			return &Unit{}, err
		},
	}
}
