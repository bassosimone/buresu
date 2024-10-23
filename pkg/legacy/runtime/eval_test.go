// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/legacy/runtime"
	"github.com/bassosimone/buresu/pkg/legacy/runtimemock"
)

func newMockEnvironmentForEvalFunc() *runtimemock.MockEnvironment {
	return &runtimemock.MockEnvironment{
		MockEval: func(ctx context.Context, node ast.Node) (runtime.Value, error) {
			switch node.(type) {
			case *ast.IntLiteral:
				return &runtime.IntValue{Value: 42}, nil
			default:
				return nil, nil
			}
		},
	}
}

func TestEvalFunc(t *testing.T) {
	ctx := context.Background()
	env := newMockEnvironmentForEvalFunc()

	t.Run("valid argument", func(t *testing.T) {
		quotedExpr := &runtime.QuotedValue{Value: &ast.IntLiteral{Value: "42"}}
		args := []runtime.Value{quotedExpr}

		result, err := runtime.EvalFunc.Fx(ctx, env, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		intResult, ok := result.(*runtime.IntValue)
		if !ok {
			t.Fatalf("expected *runtime.IntValue, got %T", result)
		}

		if intResult.Value != 42 {
			t.Errorf("expected 42, got %d", intResult.Value)
		}
	})

	t.Run("incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{}

		_, err := runtime.EvalFunc.Fx(ctx, env, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-quoted argument", func(t *testing.T) {
		args := []runtime.Value{&runtime.IntValue{Value: 42}}

		_, err := runtime.EvalFunc.Fx(ctx, env, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
