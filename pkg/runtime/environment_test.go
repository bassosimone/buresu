package runtime_test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

type initRootScopeMockEnvironment struct {
	values      map[string]runtime.Value
	errOnDefine map[string]bool
}

func newInitRootScopeMockEnvironment(errOnDefine map[string]bool) *initRootScopeMockEnvironment {
	return &initRootScopeMockEnvironment{
		values:      make(map[string]runtime.Value),
		errOnDefine: errOnDefine,
	}
}

func (env *initRootScopeMockEnvironment) DefineValue(symbol string, value runtime.Value) error {
	if env.errOnDefine[symbol] {
		return fmt.Errorf("error defining value for symbol %s", symbol)
	}
	env.values[symbol] = value
	return nil
}

func (env *initRootScopeMockEnvironment) Eval(ctx context.Context, node ast.Node) (runtime.Value, error) {
	return nil, nil
}

func (env *initRootScopeMockEnvironment) GetValue(symbol string) (runtime.Value, bool) {
	value, ok := env.values[symbol]
	return value, ok
}

func (env *initRootScopeMockEnvironment) IsInsideFunc() bool {
	return false
}

func (env *initRootScopeMockEnvironment) Output() io.Writer {
	return nil
}

func (env *initRootScopeMockEnvironment) PushBlockScope() runtime.Environment {
	return env
}

func (env *initRootScopeMockEnvironment) PushFunctionScope() runtime.Environment {
	return env
}

func (env *initRootScopeMockEnvironment) SetValue(symbol string, value runtime.Value) error {
	if _, ok := env.values[symbol]; !ok {
		return fmt.Errorf("symbol %s not defined", symbol)
	}
	env.values[symbol] = value
	return nil
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
		{"__floatSum", runtime.FloatSumFunc},
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
	symbols := []string{"false", "true", "__intSum", "__floatSum", "display"}
	for _, symbol := range symbols {
		envWithError := newInitRootScopeMockEnvironment(map[string]bool{symbol: true})
		err = runtime.InitRootScope(envWithError)
		if err == nil {
			t.Fatalf("expected error for symbol %s, got nil", symbol)
		}
	}
}
