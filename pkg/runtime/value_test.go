// SPDX-License-Identifier: GPL-3.0-or-later

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
