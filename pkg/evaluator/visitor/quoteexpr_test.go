// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalQuoteExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("integer literal", func(t *testing.T) {
		// Test with an integer literal
		intLiteral := &ast.IntLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "42"},
			Value: "42",
		}
		quoteExpr := &ast.QuoteExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "quote"},
			Expr:  intLiteral,
		}
		result, err := evalQuoteExpr(ctx, env, quoteExpr)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != "(quote 42)" {
			t.Errorf("expected %v, got %v", "(quote 42)", result.String())
		}
	})

	t.Run("string literal", func(t *testing.T) {
		// Test with a string literal
		stringLiteral := &ast.StringLiteral{
			Token: token.Token{TokenType: token.STRING, Value: "\"hello\""},
			Value: "hello",
		}
		quoteExpr := &ast.QuoteExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "quote"},
			Expr:  stringLiteral,
		}
		result, err := evalQuoteExpr(ctx, env, quoteExpr)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != "(quote \"hello\")" {
			t.Errorf("expected %v, got %v", "(quote \"hello\")", result.String())
		}
	})

	t.Run("boolean literal", func(t *testing.T) {
		// Test with a boolean literal
		boolLiteral := &ast.TrueLiteral{
			Token: token.Token{TokenType: token.ATOM, Value: "true"},
		}
		quoteExpr := &ast.QuoteExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "quote"},
			Expr:  boolLiteral,
		}
		result, err := evalQuoteExpr(ctx, env, quoteExpr)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != "(quote true)" {
			t.Errorf("expected %v, got %v", "(quote true)", result.String())
		}
	})

	t.Run("float literal", func(t *testing.T) {
		// Test with a float literal
		floatLiteral := &ast.FloatLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "3.14"},
			Value: "3.14",
		}
		quoteExpr := &ast.QuoteExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "quote"},
			Expr:  floatLiteral,
		}
		result, err := evalQuoteExpr(ctx, env, quoteExpr)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != "(quote 3.14)" {
			t.Errorf("expected %v, got %v", "(quote 3.14)", result.String())
		}
	})
}
