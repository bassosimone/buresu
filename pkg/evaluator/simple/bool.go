// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// UnwrapBoolValue unwraps a boolean value.
func (env *Environment) UnwrapBoolValue(value visitor.Value) (bool, error) {
	bv, ok := value.(*Bool)
	if !ok {
		return false, fmt.Errorf("expected a boolean value, got %T", value)
	}
	return bv.Value, nil
}

// NewBoolValue implements [visitor.Environment].
func (env *Environment) NewBoolValue(value bool) visitor.Value {
	return &Bool{value}
}

// Bool represents a boolean value.
type Bool struct {
	Value bool
}

// Ensure Bool implements [visitor.Value].
var _ visitor.Value = (*Bool)(nil)

// String implements [visitor.Value].
func (v *Bool) String() string {
	return fmt.Sprintf("%t", v.Value)
}
