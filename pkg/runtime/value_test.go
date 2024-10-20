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

func newMockEnvironmentForDisplayFunc() *runtimemock.MockEnvironment {
	output := &bytes.Buffer{}
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
			return output
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

func TestDisplayFunc(t *testing.T) {
	env := newMockEnvironmentForDisplayFunc()
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
		if env.Output().(*bytes.Buffer).String() != expectedOutput {
			t.Errorf("expected %q, got %q", expectedOutput, env.Output().(*bytes.Buffer).String())
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
