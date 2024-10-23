// SPDX-License-Identifier: GPL-3.0-or-later

package runtime

import (
	"encoding/json"
	"fmt"
)

// StringValue represents a string value in the runtime.
type StringValue struct {
	Value string
}

// Ensure *StringValue implements Value interface.
var _ Value = (*StringValue)(nil)

// String returns the string representation of the string value.
func (v *StringValue) String() string {
	repr, _ := json.Marshal(v.Value)
	return fmt.Sprintf("%s", repr)
}
