// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalFalseLiteral(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	falseLiteral := &ast.FalseLiteral{
		Token: token.Token{TokenType: token.ATOM, Value: "false"},
	}

	result, err := evalFalseLiteral(ctx, env, falseLiteral)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result.String() != env.NewBoolValue(false).String() {
		t.Errorf("expected %v, got %v", env.NewBoolValue(false), result)
	}
}
