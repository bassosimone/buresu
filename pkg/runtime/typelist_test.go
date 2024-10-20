// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
)

func TestListValue(t *testing.T) {
	t.Run("String representation of empty list", func(t *testing.T) {
		lv := &runtime.ListValue{}
		expected := "()"
		if lv.String() != expected {
			t.Errorf("expected %s, got %s", expected, lv.String())
		}
	})

	t.Run("String representation of non-empty list", func(t *testing.T) {
		lv := &runtime.ListValue{
			Car: &runtime.IntValue{Value: 1},
			Cdr: &runtime.ListValue{
				Car: &runtime.IntValue{Value: 2},
				Cdr: &runtime.ListValue{
					Car: &runtime.IntValue{Value: 3},
				},
			},
		}
		expected := "(1 2 3)"
		if lv.String() != expected {
			t.Errorf("expected %s, got %s", expected, lv.String())
		}
	})
}

func TestCarFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("car of non-empty list", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.ListValue{
				Car: &runtime.IntValue{Value: 1},
				Cdr: &runtime.ListValue{
					Car: &runtime.IntValue{Value: 2},
				},
			},
		}

		result, err := runtime.CarFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		intResult, ok := result.(*runtime.IntValue)
		if !ok {
			t.Fatalf("expected *runtime.IntValue, got %T", result)
		}

		if intResult.Value != 1 {
			t.Errorf("expected 1, got %d", intResult.Value)
		}
	})

	t.Run("car of empty list", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.ListValue{},
		}

		result, err := runtime.CarFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.String() != "()" {
			t.Errorf("expected (), got %s", result.String())
		}
	})

	t.Run("car with incorrect argument type", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.CarFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("car with incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.ListValue{},
			&runtime.ListValue{},
		}

		_, err := runtime.CarFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestCdrFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("cdr of non-empty list", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.ListValue{
				Car: &runtime.IntValue{Value: 1},
				Cdr: &runtime.ListValue{
					Car: &runtime.IntValue{Value: 2},
				},
			},
		}

		result, err := runtime.CdrFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		listResult, ok := result.(*runtime.ListValue)
		if !ok {
			t.Fatalf("expected *runtime.ListValue, got %T", result)
		}

		expected := "(2)"
		if listResult.String() != expected {
			t.Errorf("expected %s, got %s", expected, listResult.String())
		}
	})

	t.Run("cdr of empty list", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.ListValue{},
		}

		result, err := runtime.CdrFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.String() != "()" {
			t.Errorf("expected (), got %s", result.String())
		}
	})

	t.Run("cdr with incorrect argument type", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.CdrFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("cdr with incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.ListValue{},
			&runtime.ListValue{},
		}

		_, err := runtime.CdrFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestListFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("list with multiple elements", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
			&runtime.IntValue{Value: 3},
		}

		result, err := runtime.ListFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		listResult, ok := result.(*runtime.ListValue)
		if !ok {
			t.Fatalf("expected *runtime.ListValue, got %T", result)
		}

		expected := "(1 2 3)"
		if listResult.String() != expected {
			t.Errorf("expected %s, got %s", expected, listResult.String())
		}
	})

	t.Run("list with no elements", func(t *testing.T) {
		args := []runtime.Value{}

		result, err := runtime.ListFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.String() != "()" {
			t.Errorf("expected (), got %s", result.String())
		}
	})
}
