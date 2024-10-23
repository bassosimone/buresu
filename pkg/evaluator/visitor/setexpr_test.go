// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalSetExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()
	env.DefineValue("x", env.NewIntValue(0))

	// Create a set expression
	set := &ast.SetExpr{
		Token:  token.Token{TokenType: token.ATOM, Value: "set!"},
		Symbol: "x",
		Expr: &ast.IntLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "42"},
			Value: "42",
		},
	}

	// Evaluate the set expression
	result, err := evalSetExpr(ctx, env, set)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check the result of the set expression
	if result.String() != env.NewIntValue(42).String() {
		t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
	}

	// Verify the value in the environment
	val, err := env.GetValue("x")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if val.String() != env.NewIntValue(42).String() {
		t.Errorf("expected %v, got %v", env.NewIntValue(42), val)
	}
}
