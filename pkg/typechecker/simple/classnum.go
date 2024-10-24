// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// Num is the numeric type class.
type Num struct {
	Type visitor.Type
}

// Ensure that Num is a type class.
var _ Class = &Num{}

// Instantiate implements Class.
func (n *Num) Instantiate(env *Environment) {
	// define the `+` operator
	rtx.Must(env.DefineType("+", &Callable{
		ParamsTypes: []visitor.Type{
			n.Type,
			n.Type,
		},
		ReturnType: n.Type,
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return n.Type, nil
		},
	}))

	// define the `*` operator
	rtx.Must(env.DefineType("*", &Callable{
		ParamsTypes: []visitor.Type{
			n.Type,
			n.Type,
		},
		ReturnType: n.Type,
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return n.Type, nil
		},
	}))
}
