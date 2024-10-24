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
		node      *ast.ReturnStmt
		env       *mockEnvironment
		wantType  Type
		wantError error
	}{
		{
			name: "simple return",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, err := checkReturnStmt(context.Background(), tt.env, tt.node)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("checkReturnStmt() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if gotType.String() != tt.wantType.String() {
				t.Errorf("checkReturnStmt() gotType = %v, want %v", gotType, tt.wantType)
			}
		})
	}
}
