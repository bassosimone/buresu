package runtime_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// Mock environment for testing purposes
type builtinMockEnvironment struct {
	output bytes.Buffer
	values map[string]runtime.Value
}

func newBuiltinMockEnvironment() *builtinMockEnvironment {
	return &builtinMockEnvironment{
		values: make(map[string]runtime.Value),
	}
}

func (env *builtinMockEnvironment) DefineValue(symbol string, value runtime.Value) error {
	env.values[symbol] = value
	return nil
}

func (env *builtinMockEnvironment) Eval(ctx context.Context, node ast.Node) (runtime.Value, error) {
	return nil, nil
}

func (env *builtinMockEnvironment) GetValue(symbol string) (runtime.Value, bool) {
	value, ok := env.values[symbol]
	return value, ok
}

func (env *builtinMockEnvironment) IsInsideFunc() bool {
	return false
}

func (env *builtinMockEnvironment) Output() io.Writer {
	return &env.output
}

func (env *builtinMockEnvironment) PushBlockScope() runtime.Environment {
	return env
}

func (env *builtinMockEnvironment) PushFunctionScope() runtime.Environment {
	return env
}

func (env *builtinMockEnvironment) SetValue(symbol string, value runtime.Value) error {
	if _, ok := env.values[symbol]; !ok {
		return fmt.Errorf("symbol %s not defined", symbol)
	}
	env.values[symbol] = value
	return nil
}

func TestDisplayFunc(t *testing.T) {
	env := newBuiltinMockEnvironment()
	ctx := context.Background()

	t.Run("valid arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.StringValue{Value: "hello"},
			&runtime.IntValue{Value: 42},
		}

		_, err := runtime.DisplayFunc.Fx(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedOutput := "\"hello\" 42\n"
		if env.output.String() != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, env.output.String())
		}
	})

	t.Run("argument cannot be converted to string", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 42},
			nil,
		}

		_, err := runtime.DisplayFunc.Fx(ctx, env, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestIntSumFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("valid arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		result, err := runtime.IntSumFunc.Fx(ctx, nil, args...)
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

	t.Run("incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.IntSumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-integer first argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.StringValue{Value: "hello"},
			&runtime.IntValue{Value: 2},
		}

		_, err := runtime.IntSumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-integer second argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.StringValue{Value: "hello"},
		}

		_, err := runtime.IntSumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestFloatSumFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("valid arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.FloatValue{Value: 0.5},
			&runtime.FloatValue{Value: 0.5},
		}

		result, err := runtime.FloatSumFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		floatResult, ok := result.(*runtime.FloatValue)
		if !ok {
			t.Fatalf("expected *runtime.FloatValue, got %T", result)
		}

		if floatResult.Value != 1.0 {
			t.Errorf("expected 1.0, got %f", floatResult.Value)
		}
	})

	t.Run("incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.FloatValue{Value: 0.5},
		}

		_, err := runtime.FloatSumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-float first argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.StringValue{Value: "hello"},
			&runtime.FloatValue{Value: 0.5},
		}

		_, err := runtime.FloatSumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-float second argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.FloatValue{Value: 0.5},
			&runtime.StringValue{Value: "hello"},
		}

		_, err := runtime.FloatSumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestIntLtFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("valid arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		result, err := runtime.IntLtFunc.Fx(ctx, nil, args...)
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

	t.Run("incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.IntLtFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-integer first argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.StringValue{Value: "hello"},
			&runtime.IntValue{Value: 2},
		}

		_, err := runtime.IntLtFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-integer second argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.StringValue{Value: "hello"},
		}

		_, err := runtime.IntLtFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
