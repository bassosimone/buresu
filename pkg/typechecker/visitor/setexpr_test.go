// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckSetExpr(t *testing.T) {
	env := &mockEnvironment{}

	tok := token.Token{TokenType: token.ATOM, Value: "set!"}
	symbol := "x"
	expr := &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}
	setExpr := &ast.SetExpr{Token: tok, Symbol: symbol, Expr: expr}

	expectedType := &mockType{name: "Int"}

	tests := []struct {
		name    string
		ctxFunc func() context.Context
		wantErr bool
	}{
		{
			name:    "successful set expression with normal context",
			ctxFunc: normalContext,
			wantErr: false,
		},
		{
			name:    "set expression with canceled context",
			ctxFunc: canceledContext,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.ctxFunc()
			typ, err := checkSetExpr(ctx, env, setExpr)
			if (err != nil) != tt.wantErr {
				t.Fatalf("checkSetExpr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && typ.String() != expectedType.String() {
				t.Errorf("expected %s, got %s", expectedType.String(), typ.String())
			}
		})
	}
}
