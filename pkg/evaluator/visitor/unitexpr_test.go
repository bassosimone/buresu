// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalUnitExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	// Create a UnitExpr
	unit := &ast.UnitExpr{
		Token: token.Token{TokenType: token.ATOM, Value: "()"},
	}

	// Evaluate the UnitExpr
	result, err := evalUnitExpr(ctx, env, unit)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check the result
	if result.String() != env.NewUnitValue().String() {
		t.Errorf("expected %v, got %v", env.NewUnitValue(), result)
	}
}
