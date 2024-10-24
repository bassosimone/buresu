// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"errors"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckReturnStmt(t *testing.T) {
	tests := []struct {
		name      string
		ctxFunc   func() context.Context
		node      *ast.ReturnStmt
		env       *mockEnvironment
		wantType  Type
		wantError error
	}{
		{
			name:    "simple return with normal context",
			ctxFunc: normalContext,
			node: &ast.ReturnStmt{
				Token: token.Token{TokenType: token.ATOM, Value: "return!"},
				Expr:  &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"},
			},
			env: &mockEnvironment{
				returnType: &mockType{"Int"},
				err:        nil,
			},
			wantType:  &mockType{"Int"},
			wantError: nil,
		},
		{
			name:    "simple return with canceled context",
			ctxFunc: canceledContext,
			node: &ast.ReturnStmt{
				Token: token.Token{TokenType: token.ATOM, Value: "return!"},
				Expr:  &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"},
			},
			env: &mockEnvironment{
				returnType: &mockType{"Int"},
				err:        nil,
			},
			wantType:  nil,
			wantError: context.Canceled,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.ctxFunc()
			gotType, err := checkReturnStmt(ctx, tt.env, tt.node)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("checkReturnStmt() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if gotType != nil && gotType.String() != tt.wantType.String() {
				t.Errorf("checkReturnStmt() gotType = %v, want %v", gotType, tt.wantType)
			}
		})
	}
}
