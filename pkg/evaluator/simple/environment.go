// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"errors"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
	"github.com/bassosimone/buresu/pkg/token"
)

const (
	// environmentFlagScopeFunc indicates that the scope is a function scope.
	environmentFlagScopeFunc = 1 << iota
)

// Environment is the environment used by the simple evaluator.
//
// Use [NewEnvironment] to construct.
type Environment struct {
	// flags contains flags describing this environment.
	flags int

	// parent is a pointer to the parent environment.
	//
	// The root environment has a nil parent.
	parent *Environment

	// symbols contains the symbols defined in the current environment.
	symbols map[string]visitor.Value
}

// Environment implements [visitor.Environment].
var _ visitor.Environment = (*Environment)(nil)

// NewEnvironment creates a new [*Environment] instance.
func NewEnvironment() *Environment {
	return &Environment{
		flags:   0,
		parent:  nil,
		symbols: make(map[string]visitor.Value),
	}
}

// IsInsideFunc implements [visitor.Environment].
func (env *Environment) IsInsideFunc() bool {
	if env.flags&environmentFlagScopeFunc != 0 {
		return true
	}
	if env.parent != nil {
		return env.parent.IsInsideFunc()
	}
	return false
}

// PushFunctionScope implements [visitor.Environment].
func (env *Environment) PushFunctionScope() visitor.Environment {
	return env.pushScope(environmentFlagScopeFunc)
}

// PushBlockScope implements [visitor.Environment].
func (env *Environment) PushBlockScope() visitor.Environment {
	return env.pushScope(0)
}

// pushScope creates a new child environment with the given flags and returns it.
func (env *Environment) pushScope(flags int) *Environment {
	return &Environment{
		flags:   flags,
		parent:  env,
		symbols: make(map[string]visitor.Value),
	}
}

// ErrSymbolNotFound is the error returned when a symbol is not found.
var ErrSymbolNotFound = errors.New("symbol not found")

// GetValue returns the value associated with the given symbol.
//
// If the symbol is not found in the current environment, the parent
// environments are searched recursively.
func (env *Environment) GetValue(symbol string) (visitor.Value, error) {
	if value, ok := env.symbols[symbol]; ok {
		return value, nil
	}
	if env.parent != nil {
		return env.parent.GetValue(symbol)
	}
	return env.NewUnitValue(), fmt.Errorf("%w: %s", ErrSymbolNotFound, symbol)
}

// ErrSymbolAlreadyDefined is the error returned when a symbol is already defined.
var ErrSymbolAlreadyDefined = errors.New("symbol already defined")

// DefineValue implements [visitor.Environment].
func (env *Environment) DefineValue(symbol string, value visitor.Value) error {
	if _, found := env.symbols[symbol]; found {
		return fmt.Errorf("%w: %s", ErrSymbolAlreadyDefined, symbol)
	}
	env.symbols[symbol] = value
	return nil
}

// SetValue sets the value of an existing symbol in the current environment.
func (env *Environment) SetValue(symbol string, value visitor.Value) error {
	// attempt to set the value in the current environment first
	if _, found := env.symbols[symbol]; found {
		env.symbols[symbol] = value
		return nil
	}

	// otherwise attempt to set the value in the parent environment
	if env.parent != nil {
		return env.parent.SetValue(symbol, value)
	}

	// as a base case, bail
	return fmt.Errorf("%w: %s", ErrSymbolNotFound, symbol)
}

// EvalCallable implements [visitor.Environment].
func (env *Environment) EvalCallable(ctx context.Context, node ast.Node) (visitor.Callable, error) {
	callable, err := visitor.Eval(ctx, env, node)
	if err != nil {
		return nil, err
	}
	if _, ok := callable.(visitor.Callable); !ok {
		return nil, fmt.Errorf("expected a callable, got %T", callable)
	}
	return callable.(visitor.Callable), nil
}

// WrapError implements [visitor.Environment].
func (env *Environment) WrapError(tok token.Token, err error) error {
	return fmt.Errorf("%s: interpreter: %w", tok.TokenPos, err)
}
