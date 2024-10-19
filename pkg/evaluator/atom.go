// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"strconv"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

// evalFalseLiteral evaluates a FalseLiteral node and returns a BoolValue set to false.
func evalFalseLiteral(_ context.Context,
	_ runtime.Environment, _ *ast.FalseLiteral) (runtime.Value, error) {
	return &runtime.BoolValue{Value: false}, nil
}

// evalIntLiteral evaluates an IntLiteral node and returns an IntValue with the parsed integer.
func evalIntLiteral(_ context.Context,
	_ runtime.Environment, node *ast.IntLiteral) (runtime.Value, error) {
	value, err := strconv.Atoi(node.Value)
	if err != nil {
		return nil, wrapError(node.Token, err)
	}
	return &runtime.IntValue{Value: value}, nil
}

// evalFloatLiteral evaluates a FloatLiteral node and returns a FloatValue with the parsed float.
func evalFloatLiteral(_ context.Context,
	_ runtime.Environment, node *ast.FloatLiteral) (runtime.Value, error) {
	value, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return nil, wrapError(node.Token, err)
	}
	return &runtime.FloatValue{Value: value}, nil
}

// evalStringLiteral evaluates a StringLiteral node and returns a StringValue with the node's value.
func evalStringLiteral(_ context.Context,
	_ runtime.Environment, node *ast.StringLiteral) (runtime.Value, error) {
	return &runtime.StringValue{Value: node.Value}, nil
}

// evalTrueLiteral evaluates a TrueLiteral node and returns a BoolValue set to true.
func evalTrueLiteral(_ context.Context,
	_ runtime.Environment, _ *ast.TrueLiteral) (runtime.Value, error) {
	return &runtime.BoolValue{Value: true}, nil
}

// evalUnitExpr evaluates a UnitExpr node and returns a UnitValue.
func evalUnitExpr(_ context.Context,
	_ runtime.Environment, _ *ast.UnitExpr) (runtime.Value, error) {
	return &runtime.UnitValue{}, nil
}
