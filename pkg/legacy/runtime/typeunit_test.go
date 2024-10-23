// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"testing"

	"github.com/bassosimone/buresu/pkg/legacy/runtime"
)

func TestUnitValue(t *testing.T) {
	v := runtime.NewUnitValue()
	if v.String() != "()" {
		t.Errorf("expected (), got %s", v.String())
	}
}
