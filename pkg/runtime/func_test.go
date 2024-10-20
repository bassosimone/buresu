package runtime_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
	"github.com/bassosimone/buresu/pkg/runtimemock"
	"github.com/bassosimone/buresu/pkg/token"
)

func newMockEnvironmentForLambdaValue() *runtimemock.MockEnvironment {
	values := map[string]runtime.Value{}

	env := &runtimemock.MockEnvironment{
		MockDefineValue: func(symbol string, value runtime.Value) error {
			if _, ok := values[symbol]; ok {
				return fmt.Errorf("symbol %s already defined", symbol)
			}
			values[symbol] = value
			return nil
		},
		MockEval: func(ctx context.Context, node ast.Node) (runtime.Value, error) {
			switch n := node.(type) {
			case *ast.IntLiteral:
				val, err := strconv.Atoi(n.Value)
				if err != nil {
					return nil, err
				}
				return &runtime.IntValue{Value: int(val)}, nil
			case *ast.FloatLiteral:
				val, err := strconv.ParseFloat(n.Value, 64)
				if err != nil {
					return nil, err
				}
				return &runtime.Float64Value{Value: float64(val)}, nil
			case *ast.StringLiteral:
				return &runtime.StringValue{Value: n.Value}, nil
			case *ast.TrueLiteral:
				return &runtime.BoolValue{Value: true}, nil
			case *ast.FalseLiteral:
				return &runtime.BoolValue{Value: false}, nil
			case *ast.SymbolName:
				value, ok := values[n.Value]
				if !ok {
					return nil, fmt.Errorf("undefined symbol: %s", n.Value)
				}
				return value, nil
			default:
				return nil, fmt.Errorf("unsupported AST node: %T", node)
			}
		},
		MockGetValue: func(symbol string) (runtime.Value, bool) {
			value, ok := values[symbol]
			return value, ok
		},
		MockIsInsideFunc: func() bool {
			return false
		},
		MockOutput: func() io.Writer {
			return &bytes.Buffer{}
		},
		MockPushBlockScope: func() runtime.Environment {
			return nil
		},
		MockPushFunctionScope: nil, // set below
		MockSetValue: func(symbol string, value runtime.Value) error {
			if _, ok := values[symbol]; !ok {
				return fmt.Errorf("undefined symbol: %s", symbol)
			}
			values[symbol] = value
			return nil
		},
	}
	env.MockPushFunctionScope = func() runtime.Environment {
		return env
	}
	return env
}

func TestLambdaValue(t *testing.T) {
	ctx := context.Background()
	env := newMockEnvironmentForLambdaValue()

	t.Run("Test LambdaValue with correct arguments", func(t *testing.T) {
		// Mock AST node for lambda expression
		node := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x", "y"},
			Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.ATOM, Value: "+"}, Value: "3"},
		}

		lambda := &runtime.LambdaValue{
			Closure: env,
			Node:    node,
		}

		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		result, err := lambda.Call(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		intResult, ok := result.(*runtime.IntValue)
		if !ok {
			t.Fatalf("expected *runtime.IntValue, got %T", result)
		}

		if intResult.Value != 3 {
			t.Errorf("expected 3, got %d", intResult.Value)
		}
	})

	t.Run("Test LambdaValue with wrong number of arguments", func(t *testing.T) {
		// Mock AST node for lambda expression
		node := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x", "y"},
			Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.ATOM, Value: "+"}, Value: "3"},
		}

		lambda := &runtime.LambdaValue{
			Closure: env,
			Node:    node,
		}

		wrongArgs := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := lambda.Call(ctx, env, wrongArgs...)
		if err == nil {
			t.Fatalf("expected error when calling lambda with wrong number of arguments, got nil")
		}
	})

	t.Run("Test LambdaValue string representation", func(t *testing.T) {
		// Mock AST node for lambda expression
		node := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x", "y"},
			Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.ATOM, Value: "x"}, Value: "x"},
		}

		lambda := &runtime.LambdaValue{
			Closure: env,
			Node:    node,
		}

		expectedString := "{:0:0 ATOM lambda}: (lambda (x y) \"\" x)"
		if lambda.String() != expectedString {
			t.Errorf("expected %s, got %s", expectedString, lambda.String())
		}
	})

	t.Run("Test LambdaValue with duplicate parameter names", func(t *testing.T) {
		// Mock AST node for lambda expression with duplicate parameter names
		node := &ast.LambdaExpr{
			Token:  token.Token{TokenType: token.ATOM, Value: "lambda"},
			Params: []string{"x", "x"},
			Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.ATOM, Value: "+"}, Value: "3"},
		}

		lambda := &runtime.LambdaValue{
			Closure: env,
			Node:    node,
		}

		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		_, err := lambda.Call(ctx, env, args...)
		if err == nil {
			t.Fatalf("expected error when calling lambda with duplicate parameter names, got nil")
		}
	})
}

func newMockEnvironmentForBuiltInFuncValue() *runtimemock.MockEnvironment {
	var output bytes.Buffer
	return &runtimemock.MockEnvironment{
		MockDefineValue: func(symbol string, value runtime.Value) error {
			return nil
		},
		MockEval: func(ctx context.Context, node ast.Node) (runtime.Value, error) {
			return nil, nil
		},
		MockGetValue: func(symbol string) (runtime.Value, bool) {
			return nil, false
		},
		MockIsInsideFunc: func() bool {
			return false
		},
		MockOutput: func() io.Writer {
			return &output
		},
		MockPushBlockScope: func() runtime.Environment {
			return nil
		},
		MockPushFunctionScope: func() runtime.Environment {
			return nil
		},
		MockSetValue: func(symbol string, value runtime.Value) error {
			return nil
		},
	}
}

func TestBuiltInFuncValue(t *testing.T) {
	ctx := context.Background()
	env := newMockEnvironmentForBuiltInFuncValue()

	t.Run("Test __intSum built-in function", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		result, err := runtime.IntSumFunc.Call(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		intResult, ok := result.(*runtime.IntValue)
		if !ok {
			t.Fatalf("expected *runtime.IntValue, got %T", result)
		}

		if intResult.Value != 3 {
			t.Errorf("expected 3, got %d", intResult.Value)
		}
	})

	t.Run("Test __float64Sum built-in function", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.Float64Value{Value: 0.5},
			&runtime.Float64Value{Value: 0.5},
		}

		result, err := runtime.Float64SumFunc.Call(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		floatResult, ok := result.(*runtime.Float64Value)
		if !ok {
			t.Fatalf("expected *runtime.FloatValue, got %T", result)
		}

		if floatResult.Value != 1.0 {
			t.Errorf("expected 1.0, got %f", floatResult.Value)
		}
	})

	t.Run("Test __intLt built-in function", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		result, err := runtime.IntLtFunc.Call(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		boolResult, ok := result.(*runtime.BoolValue)
		if !ok {
			t.Fatalf("expected *runtime.BoolValue, got %T", result)
		}

		if !boolResult.Value {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("Test display built-in function", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.StringValue{Value: "hello"},
			&runtime.IntValue{Value: 42},
		}

		_, err := runtime.DisplayFunc.Call(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedOutput := "\"hello\" 42\n"
		if env.Output().(*bytes.Buffer).String() != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, env.Output().(*bytes.Buffer).String())
		}
	})
}
