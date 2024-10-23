// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalTrueLiteral(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	// Create a true literal
	trueLiteral := &ast.TrueLiteral{
		Token: token.Token{TokenType: token.ATOM, Value: "true"},
	}

	// Evaluate the true literal
	result, err := evalTrueLiteral(ctx, env, trueLiteral)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check the result
	if result.String() != env.NewBoolValue(true).String() {
		t.Errorf("expected %v, got %v", env.NewBoolValue(true), result)
	}
}
