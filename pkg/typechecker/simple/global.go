// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"

	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// NewGlobalEnvironment creates a new global environment loading the
// standard library runtime from the given base path.
func NewGlobalEnvironment(basePath string) (*Environment, error) {
	env := NewEnvironment()

	// define the `Num` type class for `Int` and `Float64`
	(&Num{&Float64{}}).Instantiate(env)
	(&Num{&Int{}}).Instantiate(env)

	// define the `Ord` type class for `Int` and `Float64`
	(&Ord{&Float64{}}).Instantiate(env)
	(&Ord{&Int{}}).Instantiate(env)

	// define the `Seq` type class for `String` and `Unit`
	(&Seq{&String{}}).Instantiate(env)
	(&Seq{&Unit{}}).Instantiate(env)

	// define the `display` built-in function
	env.DefineType("display", &Callable{
		ParamsTypes: []visitor.Type{&Variadic{&Any{}}},
		ReturnType:  &Unit{},
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return &Unit{}, nil
		},
		Previous: nil,
	})

	return env, nil
}
