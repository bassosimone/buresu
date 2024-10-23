// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/legacy/runtime"
)

func TestIntValue(t *testing.T) {
	v := &runtime.IntValue{Value: 42}
	if v.String() != "42" {
		t.Errorf("expected 42, got %s", v.String())
	}
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
