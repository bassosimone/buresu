// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// TestEvalStringLiteral tests the evaluation of string literals.
func TestEvalStringLiteral(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("simple string literal", func(t *testing.T) {
		strLiteral := &ast.StringLiteral{
			Token: token.Token{TokenType: token.STRING, Value: "\"Hello, World!\""},
			Value: "Hello, World!",
		}
		result, err := evalStringLiteral(ctx, env, strLiteral)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected := env.NewStringValue("Hello, World!")
		if result.String() != expected.String() {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("empty string literal", func(t *testing.T) {
		strLiteral := &ast.StringLiteral{
			Token: token.Token{TokenType: token.STRING, Value: "\"\""},
			Value: "",
		}
		result, err := evalStringLiteral(ctx, env, strLiteral)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected := env.NewStringValue("")
		if result.String() != expected.String() {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("string literal with special characters", func(t *testing.T) {
		strLiteral := &ast.StringLiteral{
			Token: token.Token{TokenType: token.STRING, Value: "\"Hello, \\nWorld!\""},
			Value: "Hello, \nWorld!",
		}
		result, err := evalStringLiteral(ctx, env, strLiteral)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected := env.NewStringValue("Hello, \nWorld!")
		if result.String() != expected.String() {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("string literal with unicode characters", func(t *testing.T) {
		strLiteral := &ast.StringLiteral{
			Token: token.Token{TokenType: token.STRING, Value: "\"Hello, 世界!\""},
			Value: "Hello, 世界!",
		}
		result, err := evalStringLiteral(ctx, env, strLiteral)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected := env.NewStringValue("Hello, 世界!")
		if result.String() != expected.String() {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
