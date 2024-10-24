// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// Seq is the sequence type class.
type Seq struct {
	Type visitor.Type
}

// Ensure that Seq is a type class.
var _ Class = &Seq{}

// Instantiate implements Class.
func (n *Seq) Instantiate(env *Environment) {
	// define the `length` operator
	rtx.Must(env.DefineType("length", &Callable{
		ParamsTypes: []visitor.Type{
			n.Type,
		},
		ReturnType: &Int{},
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return &Int{}, nil
		},
	}))
}
