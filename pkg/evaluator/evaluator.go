// SPDX-License-Identifier: GPL-3.0-or-later

// Package evaluator implements an AST evaluator.
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
