// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEval(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	intLiteral := &ast.IntLiteral{
		Token: token.Token{TokenType: token.NUMBER, Value: "42"},
		Value: "42",
	}

	result, err := Eval(ctx, env, intLiteral)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result.String() != env.NewIntValue(42).String() {
		t.Errorf("expected %v, got %v", env.NewIntValue(42), result)
	}
}
