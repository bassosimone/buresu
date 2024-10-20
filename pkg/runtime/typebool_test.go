// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
)

func TestBoolValue(t *testing.T) {
	v := &runtime.BoolValue{Value: true}
	if v.String() != "true" {
		t.Errorf("expected true, got %s", v.String())
	}
}
