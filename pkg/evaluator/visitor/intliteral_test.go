// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalIntLiteral(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("Valid integer literal", func(t *testing.T) {
		intLiteral := &ast.IntLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "42"},
			Value: "42",
		}
		result, err := evalIntLiteral(ctx, env, intLiteral)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}
	})

	t.Run("Invalid integer literal", func(t *testing.T) {
		invalidIntLiteral := &ast.IntLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "invalid"},
			Value: "invalid",
		}
		_, err := evalIntLiteral(ctx, env, invalidIntLiteral)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
