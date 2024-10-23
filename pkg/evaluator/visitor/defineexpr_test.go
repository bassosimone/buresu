// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalDefineExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("define integer literal", func(t *testing.T) {
		define := &ast.DefineExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "define"},
			Symbol: "x",
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		// Evaluate the define expression
		result, err := evalDefineExpr(ctx, env, define)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}

		// Retrieve the value from the environment and check it
		val, err := env.GetValue("x")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if val.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), val)
		}
	})

	t.Run("define string literal", func(t *testing.T) {
		define := &ast.DefineExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "define"},
			Symbol: "y",
			Expr: &ast.StringLiteral{
				Token: token.Token{TokenType: token.STRING, Value: "\"hello\""},
				Value: "hello",
			},
		}

		// Evaluate the define expression
		result, err := evalDefineExpr(ctx, env, define)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewStringValue("hello").String() {
			t.Errorf("expected %v, got %v", env.NewStringValue("hello"), result)
		}

		// Retrieve the value from the environment and check it
		val, err := env.GetValue("y")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if val.String() != env.NewStringValue("hello").String() {
			t.Errorf("expected %v, got %v", env.NewStringValue("hello"), val)
		}
	})

	t.Run("define boolean literal", func(t *testing.T) {
		define := &ast.DefineExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "define"},
			Symbol: "z",
			Expr: &ast.TrueLiteral{
				Token: token.Token{TokenType: token.ATOM, Value: "true"},
			},
		}

		// Evaluate the define expression
		result, err := evalDefineExpr(ctx, env, define)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewBoolValue(true).String() {
			t.Errorf("expected %v, got %v", env.NewBoolValue(true), result)
		}

		// Retrieve the value from the environment and check it
		val, err := env.GetValue("z")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if val.String() != env.NewBoolValue(true).String() {
			t.Errorf("expected %v, got %v", env.NewBoolValue(true), val)
		}
	})
}
