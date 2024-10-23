// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalFloatLiteral(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	floatLiteral := &ast.FloatLiteral{
		Token: token.Token{TokenType: token.NUMBER, Value: "3.14"},
		Value: "3.14",
	}

	result, err := evalFloatLiteral(ctx, env, floatLiteral)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result.String() != env.NewFloat64Value(3.14).String() {
		t.Errorf("expected %v, got %v", env.NewFloat64Value(3.14), result)
	}
}
