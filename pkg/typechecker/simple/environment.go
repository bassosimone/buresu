// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"errors"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
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

	// rts contains the returned types.
	rts *Union

	// symbols contains the symbols defined in the current environment.
	symbols map[string]visitor.Type
}

// Environment implements [visitor.Environment].
var _ visitor.Environment = (*Environment)(nil)

// NewEnvironment creates a new [*Environment] instance.
func NewEnvironment() *Environment {
	return &Environment{
		flags:   0,
		parent:  nil,
		rts:     NewUnion(),
		symbols: make(map[string]visitor.Type),
	}
}

// AddReturnType implements [visitor.Environment].
func (env *Environment) AddReturnType(t visitor.Type) error {
	if env.flags&environmentFlagScopeFunc != 0 {
		env.rts.Add(t)
		return nil
	}
	if env.parent != nil {
		return env.parent.AddReturnType(t)
	}
	return errors.New("no function scope found")
}

// Call implements [visitor.Environment].
func (env *Environment) Call(ctx context.Context,
	node ast.Node, args ...visitor.Type) (visitor.Type, error) {
	candidate, err := visitor.Check(ctx, env, node)
	if err != nil {
		return nil, err
	}

	callable, ok := candidate.(*Callable)
	if !ok {
		return nil, errors.New("node is not callable")
	}

	rvTypes, err := callable.Call(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rvTypes, nil
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
		rts:     NewUnion(),
		symbols: make(map[string]visitor.Type),
	}
}

// ErrSymbolNotFound is the error returned when a symbol is not found.
var ErrSymbolNotFound = errors.New("symbol not found")

// GetType returns the type associated with the given symbol.
//
// If the symbol is not found in the current environment, the parent
// environments are searched recursively.
func (env *Environment) GetType(symbol string) (visitor.Type, error) {
	if value, ok := env.symbols[symbol]; ok {
		return value, nil
	}
	if env.parent != nil {
		return env.parent.GetType(symbol)
	}
	return env.NewUnitType(), fmt.Errorf("%w: %s", ErrSymbolNotFound, symbol)
}

// ErrSymbolAlreadyDefined is the error returned when a symbol is already defined.
var ErrSymbolAlreadyDefined = errors.New("symbol already defined")

// DefineValue implements [visitor.Environment].
func (env *Environment) DefineType(symbol string, value visitor.Type) error {
	if callable, ok := value.(*Callable); ok {
		return env.defineCallable(symbol, callable)
	}
	if _, found := env.symbols[symbol]; found {
		return fmt.Errorf("%w: %s", ErrSymbolAlreadyDefined, symbol)
	}
	env.symbols[symbol] = value
	return nil
}

// defineCallable defines a callable in the current environment.
func (env *Environment) defineCallable(symbol string, callable *Callable) error {
	// if the previous symbol is a callable in the current environment,  owerwrite it
	// with the new callable and add a reference to it in the new callable
	if entry, found := env.symbols[symbol]; found {
		if prevCallable, ok := entry.(*Callable); ok {
			env.symbols[symbol] = callable
			callable.Previous = prevCallable
			return nil
		}
		return fmt.Errorf("%w: %s", ErrSymbolAlreadyDefined, symbol)
	}

	// Otherwise, search for the symbol in previous scopes. If not found
	// we just create a callable in the current scope.
	entry, err := env.GetType(symbol)
	if err != nil {
		env.symbols[symbol] = callable
		return nil
	}

	// If the previous symbol is not a callable, just shadow it
	// inside the current scope.
	prevCallable, ok := entry.(*Callable)
	if !ok {
		env.symbols[symbol] = callable
		return nil
	}

	// Otherwise, make sure there's a reference to it
	callable.Previous = prevCallable
	env.symbols[symbol] = callable
	return nil
}

// SetType sets the value of an existing symbol in the current environment.
func (env *Environment) SetType(symbol string, value visitor.Type) error {
	// attempt to set the value in the current environment first
	if kind, found := env.symbols[symbol]; found {
		if _, ok := kind.(*Callable); ok {
			return fmt.Errorf("cannot reassign function %s", symbol)
		}
		env.symbols[symbol] = value
		return nil
	}

	// otherwise attempt to set the value in the parent environment
	if env.parent != nil {
		return env.parent.SetType(symbol, value)
	}

	// as a base case, bail
	return fmt.Errorf("%w: %s", ErrSymbolNotFound, symbol)
}

// WrapError implements [visitor.Environment].
func (env *Environment) WrapError(tok token.Token, err error) error {
	return fmt.Errorf("%s: typechecker: %w", tok.TokenPos, err)
}

// CheckCondition implements [visitor.Environment].
func (env *Environment) CheckCondition(ctx context.Context, node ast.Node) error {
	kind, err := visitor.Check(ctx, env, node)
	if err != nil {
		return err
	}
	if _, ok := kind.(*Bool); !ok {
		return fmt.Errorf("condition must be *simple.Bool, found %T", kind)
	}
	return nil
}
