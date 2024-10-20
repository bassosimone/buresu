// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
)

func TestFloat64Value(t *testing.T) {
	v := &runtime.Float64Value{Value: 3.14}
	if v.String() != "3.140000" {
		t.Errorf("expected 3.140000, got %s", v.String())
	}
}

func TestFloat64SumFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("valid arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.Float64Value{Value: 0.5},
			&runtime.Float64Value{Value: 0.5},
		}

		result, err := runtime.Float64SumFunc.Fx(ctx, nil, args...)
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

	t.Run("incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.Float64Value{Value: 0.5},
		}

		_, err := runtime.Float64SumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-float first argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.StringValue{Value: "hello"},
			&runtime.Float64Value{Value: 0.5},
		}

		_, err := runtime.Float64SumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("non-float second argument", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.Float64Value{Value: 0.5},
			&runtime.StringValue{Value: "hello"},
		}

		_, err := runtime.Float64SumFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
