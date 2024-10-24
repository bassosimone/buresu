// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
)

// Check evaluates the given node type in the given environment and returns a type.
func Check(ctx context.Context, env Environment, node ast.Node) (Type, error) {
	// make sure we check for context cancellation before evaluating
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// dispatch according to the node type
	switch node := node.(type) {
	case *ast.BlockExpr:
		return checkBlockExpr(ctx, env, node)
	case *ast.CallExpr:
		return checkCallExpr(ctx, env, node)
	case *ast.CondExpr:
		return checkCondExpr(ctx, env, node)
	case *ast.DefineExpr:
		return checkDefineExpr(ctx, env, node)
	case *ast.FalseLiteral:
		return checkFalseLiteral(ctx, env, node)
	case *ast.FloatLiteral:
		return checkFloatLiteral(ctx, env, node)
	case *ast.IntLiteral:
		return checkIntLiteral(ctx, env, node)
	case *ast.LambdaExpr:
		return evalLambdaExpr(ctx, env, node)
	case *ast.QuoteExpr:
		return checkQuoteExpr(ctx, env, node)
	case *ast.ReturnStmt:
		return checkReturnStmt(ctx, env, node)
	case *ast.SetExpr:
		return checkSetExpr(ctx, env, node)
	case *ast.StringLiteral:
		return checkStringLiteral(ctx, env, node)
	case *ast.SymbolName:
		return checkSymbolName(ctx, env, node)
	case *ast.TrueLiteral:
		return checkTrueLiteral(ctx, env, node)
	case *ast.UnitExpr:
		return checkUnitExpr(ctx, env, node)
	case *ast.WhileExpr:
		return checkWhileExpr(ctx, env, node)
	default:
		return nil, fmt.Errorf("unsupported node type: %T", node)
	}
}
