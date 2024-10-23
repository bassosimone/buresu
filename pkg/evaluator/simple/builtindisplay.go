// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"
	"io"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewBuiltInDisplay creates a new built-in function that displays its arguments.
func NewBuiltInDisplay(writer io.Writer) *BuiltInFuncValue {
	return &BuiltInFuncValue{
		Name: "display",
		Fx: func(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
			_, err := fmt.Fprintln(writer, args)
			return &Unit{}, err
		},
	}
}
