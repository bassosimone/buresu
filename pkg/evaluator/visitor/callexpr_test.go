// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalCallExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	// Define the callable function
	env.DefineValue("myFunction", NewMockCallable(func(ctx context.Context, args ...Value) (Value, error) {
		return env.NewIntValue(42), nil
	}))

	t.Run("Call with integer argument", func(t *testing.T) {
		call := &ast.CallExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "call"},
			Callable: &ast.SymbolName{
				Token: token.Token{TokenType: token.ATOM, Value: "myFunction"},
				Value: "myFunction",
			},
			Args: []ast.Node{
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "42"},
					Value: "42",
				},
			},
		}

		result, err := evalCallExpr(ctx, env, call)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}
	})
}
