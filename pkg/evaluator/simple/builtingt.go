// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// NewBuiltInLt creates a new built-in function that compares types.
func NewBuiltInGt() *BuiltInFuncValue {
	return &BuiltInFuncValue{
		Name: ">",
		Fx: func(ctx context.Context, args ...visitor.Value) (visitor.Value, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf(">: %w", ErrWrongNumberOfArguments)
			}
			ord, ok := args[0].(Ord)
			if !ok {
				return nil, fmt.Errorf(">: %w", ErrWrongArgumentType)
			}
			return ord.Gt(args[1])
		},
	}
}
