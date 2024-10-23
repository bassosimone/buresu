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
	// Create a sample node to wrap
	sampleNode := &ast.IntLiteral{
		Token: token.Token{TokenType: token.NUMBER, Value: "42"},
		Value: "42",
	}

	// Wrap the sample node
	wrappedNode := &nodeWrapper{
		Type:  "IntLiteral",
		Value: sampleNode,
	}

	// Test String method
	expectedString := "(IntLiteral 42)"
	if wrappedNode.String() != expectedString {
		t.Errorf("expected %s, got %s", expectedString, wrappedNode.String())
	}

	// Test MarshalJSON method
	expectedJSON := `{"Type":"IntLiteral","Value":{"Token":{"TokenPos":{"FileName":"","LineNumber":0,"LineColumn":0},"TokenType":"NUMBER","Value":"42"},"Value":"42"}}`
	jsonBytes, err := json.Marshal(wrappedNode)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	if diff := cmp.Diff(expectedJSON, string(jsonBytes)); diff != "" {
		t.Errorf("mismatch (-expected +got):\n%s", diff)
	}
}

func TestWrapNodeUnknownType(t *testing.T) {
	// Create a sample unknown node type (nodeWrapper itself)
	unknownNode := &nodeWrapper{
		Type:  "UnknownType",
		Value: &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "99"}, Value: "99"},
	}

	// Wrap the unknown node
	wrappedNode := wrapNode(unknownNode)

	// Test that the wrapped node is of type "Unknown"
	if wrappedNode.(*nodeWrapper).Type != "Unknown" {
		t.Errorf("expected Unknown, got %s", wrappedNode.(*nodeWrapper).Type)
	}

	// Test that the value of the wrapped node is the original unknown node
	if diff := cmp.Diff(unknownNode, wrappedNode.(*nodeWrapper).Value); diff != "" {
		t.Errorf("mismatch (-expected +got):\n%s", diff)
	}
}
