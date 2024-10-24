// SPDX-License-Identifier: GPL-3.0-or-later

// Package evaluator implements an AST evaluator.
//
// The current design of the evaluator is such that we have two
// packages that actually implement it:
//
// 1. evaluator/visitor implements a generic visitor pattern
// for the evaluation of AST nodes;
//
// 2. evaluator/simple is a simple evaluator using visitor.
//
// The reason why this design makes sense is that it allows us
// to experiment and swap the evaluator algorithm without the
// burden of having visitor and evaluation entangled.
//
// This separation reduces the cognitive load required to understand
// and modify the code, making it easier to add new features or fix bugs.
// It also allows for incremental development, where basic functionality
// can be ensured before moving on to more complex features.
//
// Additionally, this design facilitates experimentation with different
// evaluation strategies without affecting the core visitor logic. This
// can be particularly useful for trying out new optimization techniques
// or evaluation algorithms.
//
// By using A/B testing, we can empirically determine which evaluator
// performs better or produces more accurate results. This approach
// ensures that we can make data-driven decisions about which evaluator
// to use in production.
//
// Overall, this design strategy aims to balance flexibility, maintainability,
// and performance, making it easier to manage and evolve the codebase over time.
package evaluator

import (
	"context"
	"io"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/evaluator/simple"
	"github.com/bassosimone/buresu/pkg/evaluator/visitor"
)

// Value is the type of value returned by the evaluator.
type Value = visitor.Value

// Environment is the execution environment used by the evaluator.
type Environment = simple.Environment

// NewGlobalEnvironment creates a new global environment.
func NewGlobalEnvironment(writer io.Writer) *Environment {
	return simple.NewGlobalEnvironment(writer)
}

// Eval evaluates a node in the AST and returns the result.
func Eval(ctx context.Context, env *Environment, node ast.Node) (Value, error) {
	return visitor.Eval(ctx, env, node)
}
