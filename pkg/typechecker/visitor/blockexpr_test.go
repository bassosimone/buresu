// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckBlockExpr(t *testing.T) {
	env := &mockEnvironment{}

	tests := []struct {
		name     string
		ctxFunc  func() context.Context
		node     *ast.BlockExpr
		expected Type
		wantErr  bool
	}{
		{
			name:    "empty block with normal context",
			ctxFunc: normalContext,
			node: &ast.BlockExpr{
				Token: token.Token{TokenType: token.OPEN, Value: "("},
				Exprs: []ast.Node{},
			},
			expected: &mockType{"Unit"},
			wantErr:  false,
		},
		{
			name:    "block with expressions and normal context",
			ctxFunc: normalContext,
			node: &ast.BlockExpr{
				Token: token.Token{TokenType: token.OPEN, Value: "("},
				Exprs: []ast.Node{
					&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "1"}, Value: "1"},
					&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"hello\""}, Value: "hello"},
				},
			},
			expected: &mockType{"String"},
			wantErr:  false,
		},
		{
			name:    "block with expressions and canceled context",
			ctxFunc: canceledContext,
			node: &ast.BlockExpr{
				Token: token.Token{TokenType: token.OPEN, Value: "("},
				Exprs: []ast.Node{
					&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "1"}, Value: "1"},
					&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"hello\""}, Value: "hello"},
				},
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.ctxFunc()
			got, err := checkBlockExpr(ctx, env, tt.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkBlockExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.String() != tt.expected.String() {
				t.Errorf("checkBlockExpr() = %v, want %v", got, tt.expected)
			}
		})
	}
}
