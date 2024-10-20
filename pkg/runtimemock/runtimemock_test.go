// SPDX-License-Identifier: GPL-3.0-or-later

package runtimemock_test

import (
	"context"
	"io"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
	"github.com/bassosimone/buresu/pkg/runtimemock"
)

func TestMockEnvironmentImplementsEnvironment(t *testing.T) {
	var _ runtime.Environment = &runtimemock.MockEnvironment{}
}

func TestMockEnvironment(t *testing.T) {
	t.Run("DefineValue", func(t *testing.T) {
		env := &runtimemock.MockEnvironment{
			MockDefineValue: func(symbol string, value runtime.Value) error {
				if symbol != "foo" {
					t.Fatalf("unexpected symbol: %s", symbol)
				}
				if value.String() != "42" {
					t.Fatalf("unexpected value: %v", value)
				}
				return nil
			},
		}
		err := env.DefineValue("foo", &runtime.IntValue{Value: 42})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("Eval", func(t *testing.T) {
		env := &runtimemock.MockEnvironment{
			MockEval: func(ctx context.Context, node ast.Node) (runtime.Value, error) {
				return &runtime.IntValue{Value: 42}, nil
			},
		}
		val, err := env.Eval(context.Background(), &ast.IntLiteral{Value: "42"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.String() != "42" {
			t.Fatalf("unexpected value: %v", val)
		}
	})

	t.Run("GetValue", func(t *testing.T) {
		env := &runtimemock.MockEnvironment{
			MockGetValue: func(symbol string) (runtime.Value, bool) {
				if symbol == "foo" {
					return &runtime.IntValue{Value: 42}, true
				}
				return nil, false
			},
		}
		val, ok := env.GetValue("foo")
		if !ok || val.String() != "42" {
			t.Fatalf("unexpected value: %v, ok: %v", val, ok)
		}
	})

	t.Run("IsInsideFunc", func(t *testing.T) {
		env := &runtimemock.MockEnvironment{
			MockIsInsideFunc: func() bool {
				return true
			},
		}
		if !env.IsInsideFunc() {
			t.Fatalf("expected true, got false")
		}
	})

	t.Run("Output", func(t *testing.T) {
		mockWriter := &mockWriter{}
		env := &runtimemock.MockEnvironment{
			MockOutput: func() io.Writer {
				return mockWriter
			},
		}
		if env.Output() != mockWriter {
			t.Fatalf("unexpected writer")
		}
	})

	t.Run("PushBlockScope", func(t *testing.T) {
		childEnv := &runtimemock.MockEnvironment{}
		env := &runtimemock.MockEnvironment{
			MockPushBlockScope: func() runtime.Environment {
				return childEnv
			},
		}
		if env.PushBlockScope() != childEnv {
			t.Fatalf("unexpected child environment")
		}
	})

	t.Run("PushFunctionScope", func(t *testing.T) {
		childEnv := &runtimemock.MockEnvironment{}
		env := &runtimemock.MockEnvironment{
			MockPushFunctionScope: func() runtime.Environment {
				return childEnv
			},
		}
		if env.PushFunctionScope() != childEnv {
			t.Fatalf("unexpected child environment")
		}
	})

	t.Run("SetValue", func(t *testing.T) {
		env := &runtimemock.MockEnvironment{
			MockSetValue: func(symbol string, value runtime.Value) error {
				if symbol != "foo" {
					t.Fatalf("unexpected symbol: %s", symbol)
				}
				if value.String() != "42" {
					t.Fatalf("unexpected value: %v", value)
				}
				return nil
			},
		}
		err := env.SetValue("foo", &runtime.IntValue{Value: 42})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

type mockWriter struct{}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
