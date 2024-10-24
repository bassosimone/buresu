// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
)

func TestEval(t *testing.T) {
	ctx := context.Background()
	env := NewGlobalEnvironment(nil)

	// Test with a simple integer literal
	intLiteral := &ast.IntLiteral{
		Value: "42",
	}
	result, err := Eval(ctx, env, intLiteral)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.String() != "42" {
		t.Fatalf("expected 42, got %v", result.String())
	}

	// Test with a simple boolean literal
	boolLiteral := &ast.TrueLiteral{}
	result, err = Eval(ctx, env, boolLiteral)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.String() != "true" {
		t.Fatalf("expected true, got %v", result.String())
	}
}
