// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalBlockExpr(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	// Create the block expression
	block := &ast.BlockExpr{
		Token: token.Token{TokenType: token.ATOM, Value: "block"},
		Exprs: []ast.Node{
			&ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "1"},
				Value: "1",
			},
			&ast.IntLiteral{
				Token: token.Token{TokenType: token.NUMBER, Value: "2"},
				Value: "2",
			},
		},
	}

	// Evaluate the block expression
	result, err := evalBlockExpr(ctx, env, block)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check the result of the evaluation
	if result.String() != env.NewIntValue(2).String() {
		t.Errorf("expected %v, got %v", env.NewIntValue(2), result)
	}
}
