// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewBuiltInAdd creates a new built-in function that adds two numbers.
func NewBuiltInAdd() *BuiltInFuncValue {
	return &BuiltInFuncValue{
		Name: "+",
		Fx: func(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("+: %w", ErrWrongNumberOfArguments)
			}
			num, ok := args[0].(Num)
			if !ok {
				return nil, fmt.Errorf("+: %w", ErrWrongArgumentType)
			}
			return num.Add(args[1])
		},
	}
}
