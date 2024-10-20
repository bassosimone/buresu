// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
)

func TestConsCell(t *testing.T) {
	t.Run("String representation of cons cell", func(t *testing.T) {
		cc := runtime.NewConsCell(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2})
		expected := "(1 . 2)"
		if cc.String() != expected {
			t.Errorf("expected %s, got %s", expected, cc.String())
		}
	})

	t.Run("String representation of cons list", func(t *testing.T) {
		list := runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3})
		expected := "(1 2 3)"
		if list.String() != expected {
			t.Errorf("expected %s, got %s", expected, list.String())
		}
	})

	t.Run("NewConsCell", func(t *testing.T) {
		car := &runtime.IntValue{Value: 1}
		cdr := &runtime.IntValue{Value: 2}
		cell := runtime.NewConsCell(car, cdr)

		if cell.Car() != car {
			t.Errorf("expected car to be %v, got %v", car, cell.Car())
		}

		if cell.Cdr() != cdr {
			t.Errorf("expected cdr to be %v, got %v", cdr, cell.Cdr())
		}
	})

	t.Run("Car method", func(t *testing.T) {
		car := &runtime.IntValue{Value: 1}
		cdr := &runtime.IntValue{Value: 2}
		cell := runtime.NewConsCell(car, cdr)

		if cell.Car() != car {
			t.Errorf("expected car to be %v, got %v", car, cell.Car())
		}
	})

	t.Run("Cdr method", func(t *testing.T) {
		car := &runtime.IntValue{Value: 1}
		cdr := &runtime.IntValue{Value: 2}
		cell := runtime.NewConsCell(car, cdr)

		if cell.Cdr() != cdr {
			t.Errorf("expected cdr to be %v, got %v", cdr, cell.Cdr())
		}
	})

	t.Run("Length method", func(t *testing.T) {
		list := runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3})
		length, err := list.(*runtime.ConsCell).Length()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if length != 3 {
			t.Errorf("expected length 3, got %d", length)
		}
	})

	t.Run("Reverse method", func(t *testing.T) {
		list := runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3})
		reversed, err := list.(*runtime.ConsCell).Reverse()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := "(3 2 1)"
		if reversed.String() != expected {
			t.Errorf("expected %s, got %s", expected, reversed.String())
		}
	})

	t.Run("Zero value ConsCell", func(t *testing.T) {
		var cell runtime.ConsCell

		if cell.Car() != nil {
			t.Errorf("expected car to be nil, got %v", cell.Car())
		}

		if cell.Cdr() != nil {
			t.Errorf("expected cdr to be nil, got %v", cell.Cdr())
		}

		length, err := cell.Length()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if length != 0 {
			t.Errorf("expected length 0, got %d", length)
		}

		reversed, err := cell.Reverse()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if reversed != nil {
			t.Errorf("expected nil, got %v", reversed)
		}
	})
}

func TestCarFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("car of cons cell", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsCell(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
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
			runtime.NewConsCell(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
			runtime.NewConsCell(&runtime.IntValue{Value: 3}, &runtime.IntValue{Value: 4}),
		}

		_, err := runtime.CarFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestCdrFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("cdr of cons cell", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsCell(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
		}

		result, err := runtime.CdrFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		intResult, ok := result.(*runtime.IntValue)
		if !ok {
			t.Fatalf("expected *runtime.IntValue, got %T", result)
		}

		if intResult.Value != 2 {
			t.Errorf("expected 2, got %d", intResult.Value)
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
			runtime.NewConsCell(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
			runtime.NewConsCell(&runtime.IntValue{Value: 3}, &runtime.IntValue{Value: 4}),
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

		listResult, ok := result.(*runtime.ConsCell)
		if !ok {
			t.Fatalf("expected *runtime.ConsCell, got %T", result)
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

func TestNullFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("null? with unit value", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewUnitValue(),
		}

		result, err := runtime.NullFunc.Fx(ctx, nil, args...)
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

	t.Run("null? with non-unit value", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		result, err := runtime.NullFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		boolResult, ok := result.(*runtime.BoolValue)
		if !ok {
			t.Fatalf("expected *runtime.BoolValue, got %T", result)
		}

		if boolResult.Value {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("null? with incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		_, err := runtime.NullFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
func TestLengthFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("length of cons list", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3}),
		}

		result, err := runtime.LengthFunc.Fx(ctx, nil, args...)
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

	t.Run("length of unit value", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewUnitValue(),
		}

		result, err := runtime.LengthFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		intResult, ok := result.(*runtime.IntValue)
		if !ok {
			t.Fatalf("expected *runtime.IntValue, got %T", result)
		}

		if intResult.Value != 0 {
			t.Errorf("expected 0, got %d", intResult.Value)
		}
	})

	t.Run("length with incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
			runtime.NewConsList(&runtime.IntValue{Value: 3}, &runtime.IntValue{Value: 4}),
		}

		_, err := runtime.LengthFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("length with incorrect argument type", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.LengthFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestAppendFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("append multiple cons lists", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
			runtime.NewConsList(&runtime.IntValue{Value: 3}, &runtime.IntValue{Value: 4}),
		}

		result, err := runtime.AppendFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		listResult, ok := result.(*runtime.ConsCell)
		if !ok {
			t.Fatalf("expected *runtime.ConsCell, got %T", result)
		}

		expected := "(1 2 3 4)"
		if listResult.String() != expected {
			t.Errorf("expected %s, got %s", expected, listResult.String())
		}
	})

	t.Run("append with unit value", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewUnitValue(),
			runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
		}

		result, err := runtime.AppendFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		listResult, ok := result.(*runtime.ConsCell)
		if !ok {
			t.Fatalf("expected *runtime.ConsCell, got %T", result)
		}

		expected := "(1 2)"
		if listResult.String() != expected {
			t.Errorf("expected %s, got %s", expected, listResult.String())
		}
	})
}

func TestReverseFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("reverse cons list", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3}),
		}

		result, err := runtime.ReverseFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		listResult, ok := result.(*runtime.ConsCell)
		if !ok {
			t.Fatalf("expected *runtime.ConsCell, got %T", result)
		}

		expected := "(3 2 1)"
		if listResult.String() != expected {
			t.Errorf("expected %s, got %s", expected, listResult.String())
		}
	})

	t.Run("reverse unit value", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewUnitValue(),
		}

		result, err := runtime.ReverseFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.String() != "()" {
			t.Errorf("expected (), got %s", result.String())
		}
	})

	t.Run("reverse with incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}),
			runtime.NewConsList(&runtime.IntValue{Value: 3}, &runtime.IntValue{Value: 4}),
		}

		_, err := runtime.ReverseFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("reverse with incorrect argument type", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.ReverseFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestConsFunc(t *testing.T) {
	ctx := context.Background()

	t.Run("cons with valid arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
			&runtime.IntValue{Value: 2},
		}

		result, err := runtime.ConsFunc.Fx(ctx, nil, args...)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		consResult, ok := result.(*runtime.ConsCell)
		if !ok {
			t.Fatalf("expected *runtime.ConsCell, got %T", result)
		}

		if consResult.Car().String() != "1" {
			t.Errorf("expected car to be 1, got %s", consResult.Car().String())
		}

		if consResult.Cdr().String() != "2" {
			t.Errorf("expected cdr to be 2, got %s", consResult.Cdr().String())
		}
	})

	t.Run("cons with incorrect number of arguments", func(t *testing.T) {
		args := []runtime.Value{
			&runtime.IntValue{Value: 1},
		}

		_, err := runtime.ConsFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("cons with nil arguments", func(t *testing.T) {
		args := []runtime.Value{
			nil,
			&runtime.IntValue{Value: 2},
		}

		_, err := runtime.ConsFunc.Fx(ctx, nil, args...)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestNewConsList(t *testing.T) {
	t.Run("NewConsList with multiple elements", func(t *testing.T) {
		list := runtime.NewConsList(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3})
		expected := "(1 2 3)"
		if list.String() != expected {
			t.Errorf("expected %s, got %s", expected, list.String())
		}
	})

	t.Run("NewConsList with no elements", func(t *testing.T) {
		list := runtime.NewConsList()
		expected := "()"
		if list.String() != expected {
			t.Errorf("expected %s, got %s", expected, list.String())
		}
	})

	t.Run("NewConsList with one element", func(t *testing.T) {
		list := runtime.NewConsList(&runtime.IntValue{Value: 1})
		expected := "(1)"
		if list.String() != expected {
			t.Errorf("expected %s, got %s", expected, list.String())
		}
	})

	t.Run("NewConsList with nested lists", func(t *testing.T) {
		innerList := runtime.NewConsList(&runtime.IntValue{Value: 2}, &runtime.IntValue{Value: 3})
		list := runtime.NewConsList(&runtime.IntValue{Value: 1}, innerList)
		expected := "(1 (2 3))"
		if list.String() != expected {
			t.Errorf("expected %s, got %s", expected, list.String())
		}
	})
}

func TestAppendConsLists(t *testing.T) {
	t.Run("append with too few arguments", func(t *testing.T) {
		_, err := runtime.AppendConsLists()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("append with invalid argument type", func(t *testing.T) {
		_, err := runtime.AppendConsLists(&runtime.IntValue{Value: 1})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("append with improper list", func(t *testing.T) {
		improperList := runtime.NewConsCell(&runtime.IntValue{Value: 1}, &runtime.IntValue{Value: 2})
		_, err := runtime.AppendConsLists(improperList)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
