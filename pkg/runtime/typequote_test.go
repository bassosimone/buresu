// SPDX-License-Identifier: GPL-3.0-or-later

package runtime_test

import (
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/runtime"
)

func TestQuoteValue(t *testing.T) {
	t.Run("String representation", func(t *testing.T) {
		node := &ast.StringLiteral{Value: "hello"}
		quote := &runtime.QuoteValue{Value: node}

		expected := "\"hello\""
		if quote.String() != expected {
			t.Errorf("expected %s, got %s", expected, quote.String())
		}
	})

	t.Run("Nested quote", func(t *testing.T) {
		innerNode := &ast.StringLiteral{Value: "inner"}
		outerNode := &ast.QuoteExpr{Expr: innerNode}
		quote := &runtime.QuoteValue{Value: outerNode}

		expected := "(quote \"inner\")"
		if quote.String() != expected {
			t.Errorf("expected %s, got %s", expected, quote.String())
		}
	})
}
