// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalLambdaExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("simple lambda", func(t *testing.T) {
		// Create a lambda expression
		lambda := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x"},
			Docs:   "This is a lambda function",
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		// Evaluate the lambda expression
		result, err := evalLambdaExpr(ctx, env, lambda)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewLambdaValue(lambda).String() {
			t.Errorf("expected %v, got %v", env.NewLambdaValue(lambda), result)
		}
	})

	t.Run("lambda with multiple parameters", func(t *testing.T) {
		// Create a lambda expression with multiple parameters
		lambda := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x", "y"},
			Docs:   "This is a lambda function with multiple parameters",
			Expr: &ast.CallExpr{
				Token: token.Token{TokenType: token.ATOM, Value: "+"},
				Callable: &ast.SymbolName{
					Token: token.Token{TokenType: token.ATOM, Value: "+"},
					Value: "+",
				},
				Args: []ast.Node{
					&ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "x"}, Value: "x"},
					&ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "y"}, Value: "y"},
				},
			},
		}

		// Evaluate the lambda expression
		result, err := evalLambdaExpr(ctx, env, lambda)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Check the result of the evaluation
		if result.String() != env.NewLambdaValue(lambda).String() {
			t.Errorf("expected %v, got %v", env.NewLambdaValue(lambda), result)
		}
	})
}
