// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalBlockExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("simple block", func(t *testing.T) {
		// Create the block expression
		block := &ast.BlockExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "block"},
			Exprs: []ast.Node{
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "1"},
					Value: "1",
				},
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "2"},
					Value: "2",
				},
			},
		}

		// Evaluate the block expression
		result, err := evalBlockExpr(ctx, env, block)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewIntValue(2).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(2), result)
		}
	})

	t.Run("empty block", func(t *testing.T) {
		// Create an empty block expression
		block := &ast.BlockExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "block"},
			Exprs: []ast.Node{},
		}

		// Evaluate the block expression
		result, err := evalBlockExpr(ctx, env, block)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewUnitValue().String() {
			t.Errorf("expected %v, got %v", env.NewUnitValue(), result)
		}
	})

	t.Run("block with multiple types", func(t *testing.T) {
		// Create a block expression with multiple types
		block := &ast.BlockExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "block"},
			Exprs: []ast.Node{
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "1"},
					Value: "1",
				},
				&ast.StringLiteral{
					Token: token.Token{TokenType: token.STRING, Value: "\"hello\""},
					Value: "hello",
				},
				&ast.TrueLiteral{
					Token: token.Token{TokenType: token.ATOM, Value: "true"},
				},
			},
		}

		// Evaluate the block expression
		result, err := evalBlockExpr(ctx, env, block)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewBoolValue(true).String() {
			t.Errorf("expected %v, got %v", env.NewBoolValue(true), result)
		}
	})
}
