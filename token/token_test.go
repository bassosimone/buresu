// SPDX-License-Identifier: GPL-3.0-or-later

package token_test

import (
	"testing"

	"github.com/bassosimone/buresu/token"
)

func TestPositionString(t *testing.T) {
	pos := token.Position{
		FileName:   "example.go",
		LineNumber: 10,
		LineColumn: 5,
	}
	expected := "example.go:10:5"
	if pos.String() != expected {
		t.Errorf("expected %s, got %s", expected, pos.String())
	}
}

func TestTokenClone(t *testing.T) {
	original := token.Token{
		TokenPos: token.Position{
			FileName:   "example.go",
			LineNumber: 10,
			LineColumn: 5,
		},
		TokenType: token.ATOM,
		Value:     "example",
	}
	clone := original.Clone()

	if clone.TokenPos != original.TokenPos {
		t.Errorf("expected %v, got %v", original.TokenPos, clone.TokenPos)
	}
	if clone.TokenType != original.TokenType {
		t.Errorf("expected %s, got %s", original.TokenType, clone.TokenType)
	}
	if clone.Value != original.Value {
		t.Errorf("expected %s, got %s", original.Value, clone.Value)
	}
}
