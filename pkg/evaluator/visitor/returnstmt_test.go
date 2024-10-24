// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalReturnStmt(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("return inside function", func(t *testing.T) {
		// Setup the context for testing return
		env.insideFunc = true
		env.PushFunctionScope()
		ret := &ast.ReturnStmt{
			Token: token.Token{TokenType: token.ATOM, Value: "return!"},
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		// Evaluate the return statement
		_, err := evalReturnStmt(ctx, env, ret)

		// Check for expected error
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		// Check if the error is of type errReturn
		if _, ok := err.(*errReturn); !ok {
			t.Errorf("expected errReturn, got %T", err)
		}

		// Check the value of the return statement
		if err.(*errReturn).value.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), err.(*errReturn).value)
		}
	})
}
