// SPDX-License-Identifier: GPL-3.0-or-later

// The evaluator package is responsible for evaluating the Abstract Syntax Tree (AST) nodes
// and returning the corresponding runtime values. It handles various types of expressions,
// literals, and control flow constructs by dispatching to specific evaluation functions.
package evaluator

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/ast"
	"github.com/bassosimone/buresu/runtime"
)

// Eval evaluates a node in the AST and returns the result.
func Eval(ctx context.Context, env runtime.Environment, node ast.Node) (runtime.Value, error) {
	// make sure we check for context cancellation before
	// evaluating any instruction
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// dispatch according to the node type
	switch node := node.(type) {
	case *ast.BlockExpr:
		return evalBlockExpr(ctx, env, node)
	case *ast.CallExpr:
		return evalCallExpr(ctx, env, node)
	case *ast.CondExpr:
		return evalCondExpr(ctx, env, node)
	case *ast.DefineExpr:
		return evalDefineExpr(ctx, env, node)
	case *ast.FalseLiteral:
		return evalFalseLiteral(ctx, env, node)
	case *ast.FloatLiteral:
		return evalFloatLiteral(ctx, env, node)
	case *ast.IntLiteral:
		return evalIntLiteral(ctx, env, node)
	case *ast.LambdaExpr:
		return evalLambdaExpr(ctx, env, node)
	case *ast.ReturnStmt:
		return evalReturnStmt(ctx, env, node)
	case *ast.SetExpr:
		return evalSetExpr(ctx, env, node)
	case *ast.StringLiteral:
		return evalStringLiteral(ctx, env, node)
	case *ast.SymbolName:
		return evalSymbolName(ctx, env, node)
	case *ast.TrueLiteral:
		return evalTrueLiteral(ctx, env, node)
	case *ast.UnitExpr:
		return evalUnitExpr(ctx, env, node)
	case *ast.WhileExpr:
		return evalWhileExpr(ctx, env, node)
	default:
		return nil, fmt.Errorf("unsupported node type: %T", node)
	}
}
