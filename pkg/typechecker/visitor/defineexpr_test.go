package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckDefineExpr(t *testing.T) {
	tests := []struct {
		name     string
		node     *ast.DefineExpr
		expected Type
		err      error
	}{
		{
			name: "define int",
			node: &ast.DefineExpr{
				Token:  token.Token{TokenType: token.ATOM, Value: "define"},
				Symbol: "x",
				Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"},
			},
			expected: &mockType{"Int"},
			err:      nil,
		},
		{
			name: "define float",
			node: &ast.DefineExpr{
				Token:  token.Token{TokenType: token.ATOM, Value: "define"},
				Symbol: "pi",
				Expr:   &ast.FloatLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "3.14"}, Value: "3.14"},
			},
			expected: &mockType{"Float64"},
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &mockEnvironment{}
			result, err := checkDefineExpr(context.Background(), env, tt.node)
			if err != tt.err {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
			if result.String() != tt.expected.String() {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}