// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
)



func TestStringValue(t *testing.T) {
	v := &runtime.StringValue{Value: "hello"}
	if v.String() != `"hello"` {
		t.Errorf("expected \"hello\", got %s", v.String())
	}
}
