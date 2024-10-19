// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"
	"io"

	"github.com/bassosimone/buresu/ast"
	"github.com/bassosimone/buresu/runtime"
)

const (
	// globalScopeFlagScopeFunc indicates that the scope is a function scope.
	globalScopeFlagScopeFunc = 1 << iota
)

// GlobalScope is the global scope used by the evaluator.
//
// The zero value is not ready to use; use NewGlobalScope to construct.
//
// The [GlobalScope] implements [runtime.Environment] and behaves
// as the top-level environment for the evaluator.
type GlobalScope struct {
	// flags contains flags describing this environment.
	flags int

	// output is the writer used to write output.
	output io.Writer

	// parent is a pointer to the parent environment.
	//
	// The root environment has a nil parent.
	parent *GlobalScope

	// symbols contains the symbols defined in the current environment.
	symbols map[string]runtime.Value
}

var _ runtime.Environment = (*GlobalScope)(nil)

// NewGlobalScope creates a new [*GlobalScope] instance.
func NewGlobalScope(output io.Writer) *GlobalScope {
	return &GlobalScope{
		flags:   0,
		output:  output,
		parent:  nil,
		symbols: make(map[string]runtime.Value),
	}
}

// Eval implements [runtime.Environment].
func (env *GlobalScope) Eval(ctx context.Context, node ast.Node) (runtime.Value, error) {
	return Eval(ctx, env, node)
}

// Output implements [runtime.Environment].
func (env *GlobalScope) Output() io.Writer {
	return env.output
}

// IsInsideFunc implements [runtime.Environment].
func (env *GlobalScope) IsInsideFunc() bool {
	if env.flags&globalScopeFlagScopeFunc != 0 {
		return true
	}
	if env.parent != nil {
		return env.parent.IsInsideFunc()
	}
	return false
}

// PushFunctionScope implements [runtime.Environment].
func (env *GlobalScope) PushFunctionScope() runtime.Environment {
	return env.pushScope(globalScopeFlagScopeFunc)
}

// PushBlockScope implements [runtime.Environment].
func (env *GlobalScope) PushBlockScope() runtime.Environment {
	return env.pushScope(0)
}

// pushScope creates a new child environment with the given flags and returns it.
func (env *GlobalScope) pushScope(flags int) *GlobalScope {
	return &GlobalScope{
		flags:   flags,
		output:  env.output,
		parent:  env,
		symbols: make(map[string]runtime.Value),
	}
}

// GetValue implements [runtime.Environment].
func (env *GlobalScope) GetValue(symbol string) (runtime.Value, bool) {
	if value, ok := env.symbols[symbol]; ok {
		return value, true
	}
	if env.parent != nil {
		return env.parent.GetValue(symbol)
	}
	return &runtime.UnitValue{}, false
}

// DefineValue implements [runtime.Environment].
func (env *GlobalScope) DefineValue(symbol string, value runtime.Value) error {
	if _, found := env.symbols[symbol]; found {
		return fmt.Errorf("symbol %s already defined", symbol)
	}
	env.symbols[symbol] = value
	return nil
}

// SetValue implements [runtime.Environment].
func (env *GlobalScope) SetValue(symbol string, value runtime.Value) error {
	// attempt to set the value in the current environment first
	if _, found := env.symbols[symbol]; found {
		env.symbols[symbol] = value
		return nil
	}

	// otherwise attempt to set the value in the parent environment
	if env.parent != nil {
		return env.parent.SetValue(symbol, value)
	}

	// as a base case, say that the symbol does not exist
	return fmt.Errorf("symbol %s not defined", symbol)
}
