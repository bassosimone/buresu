// SPDX-License-Identifier: GPL-3.0-or-later

package dumper

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestNodeWrapper(t *testing.T) {
	t.Run("known type", func(t *testing.T) {
		// Create a sample node to wrap and wrap it
		sampleNode := &ast.IntLiteral{
			Token: token.Token{TokenType: token.NUMBER, Value: "42"},
			Value: "42",
		}
		wrappedNode := &nodeWrapper{
			Type:  "IntLiteral",
			Value: sampleNode,
		}

		t.Run("String method", func(t *testing.T) {
			expectedString := "(IntLiteral 42)"
			if wrappedNode.String() != expectedString {
				t.Errorf("expected %s, got %s", expectedString, wrappedNode.String())
			}
		})

		t.Run("MarshalJSON method", func(t *testing.T) {
			expectedJSON := `{"Type":"IntLiteral","Value":{"Token":{"TokenPos":{"FileName":"","LineNumber":0,"LineColumn":0},"TokenType":"NUMBER","Value":"42"},"Value":"42"}}`
			jsonBytes, err := json.Marshal(wrappedNode)
			if err != nil {
				t.Fatalf("failed to marshal JSON: %v", err)
			}
			if diff := cmp.Diff(expectedJSON, string(jsonBytes)); diff != "" {
				t.Errorf("mismatch (-expected +got):\n%s", diff)
			}
		})
	})

	t.Run("unknown type", func(t *testing.T) {
		// Create a sample unknown node type (nodeWrapper itself) and wrap it
		unknownNode := &nodeWrapper{
			Type:  "UnknownType",
			Value: &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "99"}, Value: "99"},
		}
		wrappedNode := wrapNode(unknownNode)

		t.Run("Type is Unknown", func(t *testing.T) {
			// Test that the wrapped node is of type "Unknown"
			if wrappedNode.(*nodeWrapper).Type != "Unknown" {
				t.Errorf("expected Unknown, got %s", wrappedNode.(*nodeWrapper).Type)
			}
		})

		t.Run("Value is original unknown node", func(t *testing.T) {
			// Test that the value of the wrapped node is the original unknown node
			if diff := cmp.Diff(unknownNode, wrappedNode.(*nodeWrapper).Value); diff != "" {
				t.Errorf("mismatch (-expected +got):\n%s", diff)
			}
		})
	})
}
