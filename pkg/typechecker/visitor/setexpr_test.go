package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckSetExpr(t *testing.T) {
	ctx := context.Background()
	env := &mockEnvironment{}

	tok := token.Token{TokenType: token.ATOM, Value: "set!"}
	symbol := "x"
	expr := &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}
	setExpr := &ast.SetExpr{Token: tok, Symbol: symbol, Expr: expr}

	expectedType := &mockType{name: "Int"}

	t.Run("successful set expression", func(t *testing.T) {
		typ, err := checkSetExpr(ctx, env, setExpr)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.String() != expectedType.String() {
			t.Errorf("expected %s, got %s", expectedType.String(), typ.String())
		}
	})
}
