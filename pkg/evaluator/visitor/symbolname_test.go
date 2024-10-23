// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// TestEvalSymbolName tests the evaluation of symbol names.
func TestEvalSymbolName(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()
	env.DefineValue("x", env.NewIntValue(42))

	t.Run("symbol exists in the environment", func(t *testing.T) {
		symbol := &ast.SymbolName{
			Token: token.Token{TokenType: token.ATOM, Value: "x"},
			Value: "x",
		}
		result, err := evalSymbolName(ctx, env, symbol)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}
	})

	t.Run("symbol does not exist in the environment", func(t *testing.T) {
		symbol := &ast.SymbolName{
			Token: token.Token{TokenType: token.ATOM, Value: "y"},
			Value: "y",
		}
		_, err := evalSymbolName(ctx, env, symbol)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
