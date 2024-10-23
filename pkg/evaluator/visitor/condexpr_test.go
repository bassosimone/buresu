// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalCondExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("false predicate", func(t *testing.T) {
		cond := &ast.CondExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "cond"},
			Cases: []ast.CondCase{
				{
					Predicate: &ast.FalseLiteral{
						Token: token.Token{TokenType: token.ATOM, Value: "false"},
					},
					Expr: &ast.IntLiteral{
						Token: token.Token{TokenType: token.NUMBER, Value: "1"},
						Value: "1",
					},
				},
				{
					Predicate: &ast.TrueLiteral{
						Token: token.Token{TokenType: token.ATOM, Value: "true"},
					},
					Expr: &ast.IntLiteral{
						Token: token.Token{TokenType: token.NUMBER, Value: "2"},
						Value: "2",
					},
				},
			},
			ElseExpr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "3"},
				Value: "3",
			},
		}
		result, err := evalCondExpr(ctx, env, cond)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(2).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(2), result)
		}
	})

	t.Run("true predicate", func(t *testing.T) {
		cond := &ast.CondExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "cond"},
			Cases: []ast.CondCase{
				{
					Predicate: &ast.TrueLiteral{
						Token: token.Token{TokenType: token.ATOM, Value: "true"},
					},
					Expr: &ast.IntLiteral{
						Token: token.Token{TokenType: token.NUMBER, Value: "1"},
						Value: "1",
					},
				},
			},
			ElseExpr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "3"},
				Value: "3",
			},
		}
		result, err := evalCondExpr(ctx, env, cond)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(1).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(1), result)
		}
	})

	t.Run("else case", func(t *testing.T) {
		cond := &ast.CondExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "cond"},
			Cases: []ast.CondCase{
				{
					Predicate: &ast.FalseLiteral{
						Token: token.Token{TokenType: token.ATOM, Value: "false"},
					},
					Expr: &ast.IntLiteral{
						Token: token.Token{TokenType: token.NUMBER, Value: "1"},
						Value: "1",
					},
				},
			},
			ElseExpr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "3"},
				Value: "3",
			},
		}
		result, err := evalCondExpr(ctx, env, cond)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(3).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(3), result)
		}
	})
}
