// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestEvalEllipsisLiteral(t *testing.T) {
	ctx := context.Background()
	env := NewMockEnvironment()

	ellipsisLiteral := &ast.EllipsisLiteral{
		Token: token.Token{TokenType: token.ELLIPSIS, Value: "..."},
	}
	result, err := evalEllipsisLiteral(ctx, env, ellipsisLiteral)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}
