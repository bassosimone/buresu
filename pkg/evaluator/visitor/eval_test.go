// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEval(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	t.Run("integer literal", func(t *testing.T) {
		intLiteral := &ast.IntLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "42"},
			Value: "42",
		}

		result, err := Eval(ctx, env, intLiteral)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}
	})

	t.Run("string literal", func(t *testing.T) {
		stringLiteral := &ast.StringLiteral{
			Token: token.Token{TokenType: token.STRING, Value: "\"hello\""},
			Value: "hello",
		}

		result, err := Eval(ctx, env, stringLiteral)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewStringValue("hello").String() {
			t.Errorf("expected %v, got %v", env.NewStringValue("hello"), result)
		}
	})

	t.Run("boolean literal", func(t *testing.T) {
		trueLiteral := &ast.TrueLiteral{
			Token: token.Token{TokenType: token.ATOM, Value: "true"},
		}

		result, err := Eval(ctx, env, trueLiteral)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewBoolValue(true).String() {
			t.Errorf("expected %v, got %v", env.NewBoolValue(true), result)
		}
	})

	t.Run("false literal", func(t *testing.T) {
		falseLiteral := &ast.FalseLiteral{
			Token: token.Token{TokenType: token.ATOM, Value: "false"},
		}

		result, err := Eval(ctx, env, falseLiteral)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewBoolValue(false).String() {
			t.Errorf("expected %v, got %v", env.NewBoolValue(false), result)
		}
	})

	t.Run("float literal", func(t *testing.T) {
		floatLiteral := &ast.FloatLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "3.14"},
			Value: "3.14",
		}

		result, err := Eval(ctx, env, floatLiteral)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewFloat64Value(3.14).String() {
			t.Errorf("expected %v, got %v", env.NewFloat64Value(3.14), result)
		}
	})

	t.Run("unit expression", func(t *testing.T) {
		unitExpr := &ast.UnitExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "()"},
		}

		result, err := Eval(ctx, env, unitExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewUnitValue().String() {
			t.Errorf("expected %v, got %v", env.NewUnitValue(), result)
		}
	})

	t.Run("symbol name", func(t *testing.T) {
		env.DefineValue("x", env.NewIntValue(42))
		symbolName := &ast.SymbolName{
			Token: token.Token{TokenType: token.ATOM, Value: "x"},
			Value: "x",
		}

		result, err := Eval(ctx, env, symbolName)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}
	})

	t.Run("quote expression", func(t *testing.T) {
		quoteExpr := &ast.QuoteExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "quote"},
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		result, err := Eval(ctx, env, quoteExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != "(quote 42)" {
			t.Errorf("expected %v, got %v", "(quote 42)", result.String())
		}
	})

	t.Run("define expression", func(t *testing.T) {
		defineExpr := &ast.DefineExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "define"},
			Symbol: "x",
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		result, err := Eval(ctx, env, defineExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}

		val, err := env.GetValue("x")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if val.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), val)
		}
	})

	t.Run("set expression", func(t *testing.T) {
		env.DefineValue("x", env.NewIntValue(0))
		setExpr := &ast.SetExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "set!"},
			Symbol: "x",
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		result, err := Eval(ctx, env, setExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}

		val, err := env.GetValue("x")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if val.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), val)
		}
	})

	t.Run("block expression", func(t *testing.T) {
		blockExpr := &ast.BlockExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "block"},
			Exprs: []ast.Node{
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "1"},
					Value: "1",
				},
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "2"},
					Value: "2",
				},
			},
		}

		result, err := Eval(ctx, env, blockExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(2).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(2), result)
		}
	})

	t.Run("call expression", func(t *testing.T) {
		env.DefineValue("myFunction", NewMockCallable(func(ctx context.Context, args ...Value) (Value, error) {
			return env.NewIntValue(42), nil
		}))

		callExpr := &ast.CallExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "call"},
			Callable: &ast.SymbolName{
				Token: token.Token{TokenType: token.ATOM, Value: "myFunction"},
				Value: "myFunction",
			},
			Args: []ast.Node{
				&ast.IntLiteral{
					Token: token.Token{TokenType: token.NUMBER, Value: "42"},
					Value: "42",
				},
			},
		}

		result, err := Eval(ctx, env, callExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
		}
	})

	t.Run("conditional expression", func(t *testing.T) {
		condExpr := &ast.CondExpr{
			Token: token.Token{TokenType: token.ATOM, Value: "cond"},
			Cases: []ast.CondCase{
				{
					Predicate: &ast.FalseLiteral{
						Token: token.Token{TokenType: token.ATOM, Value: "false"},
					},
					Expr: &ast.IntLiteral{
						Token: token.Token{TokenType: token.NUMBER, Value: "1"},
						Value: "1",
					},
				},
				{
					Predicate: &ast.TrueLiteral{
						Token: token.Token{TokenType: token.ATOM, Value: "true"},
					},
					Expr: &ast.IntLiteral{
						Token: token.Token{TokenType: token.NUMBER, Value: "2"},
						Value: "2",
					},
				},
			},
			ElseExpr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "3"},
				Value: "3",
			},
		}

		result, err := Eval(ctx, env, condExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewIntValue(2).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(2), result)
		}
	})

	t.Run("lambda expression", func(t *testing.T) {
		lambdaExpr := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x"},
			Docs:   "This is a lambda function",
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		result, err := Eval(ctx, env, lambdaExpr)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.String() != env.NewLambdaValue(lambdaExpr).String() {
			t.Errorf("expected %v, got %v", env.NewLambdaValue(lambdaExpr), result)
		}
	})

	t.Run("return statement", func(t *testing.T) {
		env.insideFunc = true
		env.PushFunctionScope()
		returnStmt := &ast.ReturnStmt{
			Token: token.Token{TokenType: token.ATOM, Value: "return!"},
			Expr: &ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "42"},
				Value: "42",
			},
		}

		_, err := Eval(ctx, env, returnStmt)

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if _, ok := err.(*errReturn); !ok {
			t.Errorf("expected errReturn, got %T", err)
		}

		if err.(*errReturn).value.String() != env.NewIntValue(42).String() {
			t.Errorf("expected %v, got %v", env.NewIntValue(42), err.(*errReturn).value)
		}
	})

	t.Run("while expression", func(t *testing.T) {
		counter := 0
		env.DefineValue("counter", env.NewIntValue(counter))

		whileExpr := &ast.WhileExpr{
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

		_, err := Eval(ctx, env, whileExpr)
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
