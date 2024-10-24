package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckCallExpr(t *testing.T) {
	env := &mockEnvironment{}

	tests := []struct {
		name    string
		ctxFunc func() context.Context
		node    *ast.CallExpr
		wantErr bool
	}{
		{
			name:    "simple call expression with normal context",
			ctxFunc: normalContext,
			node: &ast.CallExpr{
				Token:    token.Token{TokenType: token.ATOM, Value: "call"},
				Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "myFunction"}, Value: "myFunction"},
				Args:     []ast.Node{&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}},
			},
			wantErr: false,
		},
		{
			name:    "simple call expression with canceled context",
			ctxFunc: canceledContext,
			node: &ast.CallExpr{
				Token:    token.Token{TokenType: token.ATOM, Value: "call"},
				Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "myFunction"}, Value: "myFunction"},
				Args:     []ast.Node{&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}},
			},
			wantErr: true,
		},
		{
			name:    "call expression with multiple arguments and normal context",
			ctxFunc: normalContext,
			node: &ast.CallExpr{
				Token:    token.Token{TokenType: token.ATOM, Value: "call"},
				Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "myFunction"}, Value: "myFunction"},
				Args:     []ast.Node{&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}, &ast.FloatLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "3.14"}, Value: "3.14"}},
			},
			wantErr: false,
		},
		{
			name:    "call expression with multiple arguments and canceled context",
			ctxFunc: canceledContext,
			node: &ast.CallExpr{
				Token:    token.Token{TokenType: token.ATOM, Value: "call"},
				Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "myFunction"}, Value: "myFunction"},
				Args:     []ast.Node{&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}, &ast.FloatLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "3.14"}, Value: "3.14"}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.ctxFunc()
			_, err := checkCallExpr(ctx, env, tt.node)
			if (err != nil) != tt.wantErr {
				t.Fatalf("checkCallExpr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
