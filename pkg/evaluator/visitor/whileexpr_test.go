// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalWhileExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()
	counter := 0
	env.DefineValue("counter", env.NewIntValue(counter))

	while := &ast.WhileExpr{
		Token: token.Token{TokenType: token.ATOM, Value: "while"},
		Predicate: &ast.CallExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "call"},
			Callable: &ast.SymbolName{
				Token: token.Token{TokenType: token.ATOM, Value: "lessThanTen"},
				Value: "lessThanTen",
			},
		},
		Expr: &ast.CallExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "call"},
			Callable: &ast.SymbolName{
				Token: token.Token{TokenType: token.ATOM, Value: "incrementCounter"},
				Value: "incrementCounter",
			},
		},
	}

	env.DefineValue("lessThanTen", NewMockCallable(func(ctx context.Context, args ...Value) (Value, error) {
		val, _ := env.GetValue("counter")
		intVal := val.(MockValue).value.(int)
		return env.NewBoolValue(intVal < 10), nil
	}))

	env.DefineValue("incrementCounter", NewMockCallable(func(ctx context.Context, args ...Value) (Value, error) {
		val, _ := env.GetValue("counter")
		intVal := val.(MockValue).value.(int)
		env.SetValue("counter", env.NewIntValue(intVal+1))
		return env.NewUnitValue(), nil
	}))

	t.Run("basic while loop", func(t *testing.T) {
		_, err := evalWhileExpr(ctx, env, while)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		val, err := env.GetValue("counter")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if val.String() != env.NewIntValue(10).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(10), val)
		}
	})

	t.Run("while loop with initial counter 5", func(t *testing.T) {
		env.SetValue("counter", env.NewIntValue(5))
		_, err := evalWhileExpr(ctx, env, while)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		val, err := env.GetValue("counter")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if val.String() != env.NewIntValue(10).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(10), val)
		}
	})

	t.Run("while loop with initial counter 10", func(t *testing.T) {
		env.SetValue("counter", env.NewIntValue(10))
		_, err := evalWhileExpr(ctx, env, while)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		val, err := env.GetValue("counter")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if val.String() != env.NewIntValue(10).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(10), val)
		}
	})
}
