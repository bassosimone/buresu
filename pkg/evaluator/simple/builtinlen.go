// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewBuiltInLength creates a new built-in function that computes the length.
func NewBuiltInLength() *BuiltInFuncValue {
	return &BuiltInFuncValue{
		Name: "length",
		Fx: func(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("length: %w", ErrWrongNumberOfArguments)
			}
			num, ok := args[0].(Seq)
			if !ok {
				return nil, fmt.Errorf("length: %w", ErrWrongArgumentType)
			}
			return num.Length()
		},
	}
}
