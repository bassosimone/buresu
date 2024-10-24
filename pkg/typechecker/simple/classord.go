// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// Ord is the ordered type class.
type Ord struct {
	Type visitor.Type
}

// Ensure that Ord is a type class.
var _ Class = &Ord{}

// Instantiate implements Class.
func (n *Ord) Instantiate(env *Environment) {
	// define the `<` operator
	rtx.Must(env.DefineType("<", &Callable{
		ParamsTypes: []visitor.Type{
			n.Type,
			n.Type,
		},
		ReturnType: &Bool{},
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return &Bool{}, nil
		},
	}))

	// define the `>` operator
	rtx.Must(env.DefineType(">", &Callable{
		ParamsTypes: []visitor.Type{
			n.Type,
			n.Type,
		},
		ReturnType: &Bool{},
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return &Bool{}, nil
		},
	}))
}
