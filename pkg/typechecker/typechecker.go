// SPDX-License-Identifier: GPL-3.0-or-later

// Package typechecker implements an AST type checker.
//
// The design of this package is similar to the one of the evaluator.
package typechecker

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/typechecker/simple"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// Type is the type of value returned by the evaluator.
type Type = visitor.Type

// Environment is the execution environment used by the evaluator.
type Environment = simple.Environment

// NewGlobalEnvironment creates a new global environment.
func NewGlobalEnvironment() *Environment {
	return simple.NewGlobalEnvironment()
}

// Check evaluates the type of a node in the AST and returns the result.
func Check(ctx context.Context, env *Environment, node ast.Node) (Type, error) {
	return visitor.Check(ctx, env, node)
}
