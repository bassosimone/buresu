package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestCheckCallExpr(t *testing.T) {
	ctx := context.Background()
	env := &mockEnvironment{}

	t.Run("simple call expression", func(t *testing.T) {
		callable := &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "myFunction"}, Value: "myFunction"}
		arg := &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}
		node := &ast.CallExpr{Token: token.Token{TokenType: token.ATOM, Value: "call"}, Callable: callable, Args: []ast.Node{arg}}

		_, err := checkCallExpr(ctx, env, node)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("call expression with multiple arguments", func(t *testing.T) {
		callable := &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "myFunction"}, Value: "myFunction"}
		arg1 := &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}
		arg2 := &ast.FloatLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "3.14"}, Value: "3.14"}
		node := &ast.CallExpr{Token: token.Token{TokenType: token.ATOM, Value: "call"}, Callable: callable, Args: []ast.Node{arg1, arg2}}

		_, err := checkCallExpr(ctx, env, node)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
