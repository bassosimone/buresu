// SPDX-License-Identifier: GPL-3.0-or-later

package optional

import "testing"

func TestSome(t *testing.T) {
	opt := Some[int](42)
	if !opt.IsSome() {
		t.Errorf("Expected IsSome to be true")
	}
	if opt.IsNone() {
		t.Errorf("Expected IsNone to be false")
	}
	if opt.Unwrap() != 42 {
		t.Errorf("Expected value to be 42")
	}
}

func TestNone(t *testing.T) {
	opt := None[int]()
	if opt.IsSome() {
		t.Errorf("Expected IsSome to be false")
	}
	if !opt.IsNone() {
		t.Errorf("Expected IsNone to be true")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on Unwrap")
		}
	}()
	opt.Unwrap()
}

func TestUnwrap(t *testing.T) {
	opt := Some("hello")
	if opt.Unwrap() != "hello" {
		t.Errorf("Expected value to be 'hello'")
	}
}
