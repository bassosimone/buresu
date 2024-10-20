// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"fmt"
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
	"github.com/bassosimone/buresu/pkg/runtimemock"
)

func newInitRootScopeMockEnvironment(errOnDefine map[string]bool) *runtimemock.MockEnvironment {
	values := make(map[string]runtime.Value)
	return &runtimemock.MockEnvironment{
		MockDefineValue: func(symbol string, value runtime.Value) error {
			if errOnDefine[symbol] {
				return fmt.Errorf("error defining value for symbol %s", symbol)
			}
			values[symbol] = value
			return nil
		},
		MockGetValue: func(symbol string) (runtime.Value, bool) {
			value, ok := values[symbol]
			return value, ok
		},
	}
}

func TestInitRootScope(t *testing.T) {
	// Test successful initialization
	env := newInitRootScopeMockEnvironment(map[string]bool{})
	err := runtime.InitRootScope(env)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		symbol string
		want   runtime.Value
	}{
		{"false", &runtime.BoolValue{Value: false}},
		{"true", &runtime.BoolValue{Value: true}},
		{"__intSum", runtime.IntSumFunc},
		{"__float64Sum", runtime.Float64SumFunc},
		{"display", runtime.DisplayFunc},
	}

	for _, tt := range tests {
		got, ok := env.GetValue(tt.symbol)
		if !ok {
			t.Errorf("symbol %s not defined", tt.symbol)
		}
		if got.String() != tt.want.String() {
			t.Errorf("symbol %s: expected %v, got %v", tt.symbol, tt.want, got)
		}
	}

	// Test error during initialization for each symbol
	symbols := []string{
		"__intSum",
		"__float64Sum",
		"car",
		"cdr",
		"display",
		"eval",
		"false",
		"list",
		"true",
	}
	for _, symbol := range symbols {
		envWithError := newInitRootScopeMockEnvironment(map[string]bool{symbol: true})
		err = runtime.InitRootScope(envWithError)
		if err == nil {
			t.Fatalf("expected error for symbol %s, got nil", symbol)
		}
	}
}
