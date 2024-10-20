// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
	"github.com/bassosimone/buresu/pkg/runtimemock"
)

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
