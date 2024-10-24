// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckCondExpr(t *testing.T) {
	mockEnv := &mockEnvironment{}

	trueToken := token.Token{TokenType: token.ATOM, Value: "true"}
	falseToken := token.Token{TokenType: token.ATOM, Value: "false"}
	trueLiteral := &ast.TrueLiteral{Token: trueToken}
	falseLiteral := &ast.FalseLiteral{Token: falseToken}
	stringLiteral := &ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"It's true!\""}, Value: "It's true!"}
	elseLiteral := &ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"Neither true nor false!\""}, Value: "Neither true nor false!"}

	condExpr := &ast.CondExpr{
		Token: trueToken,
		Cases: []ast.CondCase{
			{Predicate: trueLiteral, Expr: stringLiteral},
			{Predicate: falseLiteral, Expr: stringLiteral},
		},
		ElseExpr: elseLiteral,
	}

	condExprOnlyElse := &ast.CondExpr{
		Token:    trueToken,
		Cases:    []ast.CondCase{},
		ElseExpr: elseLiteral,
	}

	expectedType := &mockType{name: "Union"}

	mockEnv.returnType = expectedType

	tests := []struct {
		name    string
		ctxFunc func() context.Context
		env     *mockEnvironment
		node    *ast.CondExpr
		want    Type
		wantErr bool
	}{
		{
			name:    "cond expression with normal context",
			ctxFunc: normalContext,
			env:     mockEnv,
			node:    condExpr,
			want:    expectedType,
			wantErr: false,
		},
		{
			name:    "cond expression with canceled context",
			ctxFunc: canceledContext,
			env:     mockEnv,
			node:    condExpr,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "cond expression with only else and canceled context",
			ctxFunc: canceledContext,
			env:     mockEnv,
			node:    condExprOnlyElse,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.ctxFunc()
			result, err := checkCondExpr(ctx, tt.env, tt.node)
			if (err != nil) != tt.wantErr {
				t.Fatalf("checkCondExpr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if result != nil && result.String() != tt.want.String() {
				t.Errorf("expected %s, got %s", tt.want.String(), result.String())
			}
		})
	}
}
