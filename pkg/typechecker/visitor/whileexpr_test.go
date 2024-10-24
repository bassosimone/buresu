package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckWhileExpr(t *testing.T) {
	env := &mockEnvironment{}

	tests := []struct {
		name      string
		ctxFunc   func() context.Context
		predicate ast.Node
		body      ast.Node
		wantType  Type
		wantErr   bool
	}{
		{
			name:      "simple while loop with normal context",
			ctxFunc:   normalContext,
			predicate: &ast.TrueLiteral{Token: token.Token{TokenType: token.ATOM, Value: "true"}},
			body:      &ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"foobar\""}, Value: "foobar"},
			wantType:  &mockType{name: "Unit"},
			wantErr:   false,
		},
		{
			name:      "while loop with false predicate and normal context",
			ctxFunc:   normalContext,
			predicate: &ast.FalseLiteral{Token: token.Token{TokenType: token.ATOM, Value: "false"}},
			body:      &ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"foobar\""}, Value: "foobar"},
			wantType:  &mockType{name: "Unit"},
			wantErr:   false,
		},
		{
			name:      "simple while loop with canceled context",
			ctxFunc:   canceledContext,
			predicate: &ast.TrueLiteral{Token: token.Token{TokenType: token.ATOM, Value: "true"}},
			body:      &ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"foobar\""}, Value: "foobar"},
			wantType:  nil,
			wantErr:   true,
		},
		{
			name:      "while loop with false predicate and canceled context",
			ctxFunc:   canceledContext,
			predicate: &ast.FalseLiteral{Token: token.Token{TokenType: token.ATOM, Value: "false"}},
			body:      &ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"foobar\""}, Value: "foobar"},
			wantType:  nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.ctxFunc()
			node := &ast.WhileExpr{
				Token:     token.Token{TokenType: token.ATOM, Value: "while"},
				Predicate: tt.predicate,
				Expr:      tt.body,
			}
			gotType, err := checkWhileExpr(ctx, env, node)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkWhileExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotType != nil && gotType.String() != tt.wantType.String() {
				t.Errorf("checkWhileExpr() gotType = %v, want %v", gotType, tt.wantType)
			}
		})
	}
}
