// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"fmt"
	"slices"
	"strings"

	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// NewUnionType implements [visitor.Environment].
func (env *Environment) NewUnionType(types ...visitor.Type) visitor.Type {
	ut := NewUnion()
	for _, kind := range types {
		ut.Add(kind)
	}
	return ut
}

// NewUnion creates a ready-to-use union type.
func NewUnion() *Union {
	return &Union{
		Types: make(map[string]visitor.Type),
	}
}

// Union is the union of several types.
type Union struct {
	Types map[string]visitor.Type
}

// Add adds a new type to the union.
func (ut *Union) Add(kind visitor.Type) {
	ut.Types[kind.String()] = kind
}

// TODO(bassosimone): implement union simplification algorithm
// either here on in a separate file for clarity.

// Ensure that UnionType implements [visitor.Type].
var _ visitor.Type = &Union{}

// String implements [visitor.Type].
func (ut *Union) String() string {
	kinds := make([]string, 0, len(ut.Types))
	for _, kind := range ut.Types {
		kinds = append(kinds, kind.String())
	}
	slices.Sort(kinds)
	return fmt.Sprintf("(Union %s)", strings.Join(kinds, " "))
}
