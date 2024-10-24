// SPDX-License-Identifier: GPL-3.0-or-later

package rtx_test

import (
	"fmt"
	"testing"

	"github.com/bassosimone/buresu/internal/rtx"
)

func TestMust(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but did not panic")
		}
	}()
	rtx.Must(fmt.Errorf("this is an error"))
}

func TestMust_NoError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Did not expect panic, but panicked")
		}
	}()
	rtx.Must(nil)
}

func TestAssert(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but did not panic")
		}
	}()
	rtx.Assert(false, "this should panic")
}

func TestAssert_NoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Did not expect panic, but panicked")
		}
	}()
	rtx.Assert(true, "this should not panic")
}
