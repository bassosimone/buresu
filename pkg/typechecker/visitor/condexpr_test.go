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

	expectedType := &mockType{name: "Union"}

	mockEnv.returnType = expectedType

	t.Run("cond expression", func(t *testing.T) {
		result, err := checkCondExpr(context.Background(), mockEnv, condExpr)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.String() != expectedType.String() {
			t.Errorf("expected %s, got %s", expectedType.String(), result.String())
		}
	})
}
