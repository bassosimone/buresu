package includer

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func mockReadFile(filename string) ([]byte, error) {
	switch filename {
	case "file1.lisp":
		return []byte(`(define x 42)`), nil
	case "file2.lisp":
		return []byte(`(define y 43)`), nil
	case "cycle.lisp":
		return []byte(`(include "cycle.lisp")`), nil
	case "invalid_scan.lisp":
		return []byte(`@`), nil
	case "invalid_parse.lisp":
		return []byte(`(if`), nil
	case "fileX.lisp":
		return []byte(`(include "fileY.lisp") (include "fileZ.lisp") (include "fileZ.lisp")`), nil
	case "fileY.lisp":
		return []byte(`(include "fileZ.lisp")`), nil
	case "fileZ.lisp":
		return []byte(`(define z 44)`), nil
	case "block_include.lisp":
		return []byte(`(block (include "file1.lisp"))`), nil
	default:
		return nil, errors.New("file not found")
	}
}

func TestInclude(t *testing.T) {
	// Override the os.ReadFile function with the mock
	readFile = mockReadFile
	defer func() { readFile = os.ReadFile }() // Restore the original function after the test

	tests := []struct {
		name     string
		input    []ast.Node
		expected []ast.Node
		err      string
	}{
		{
			name: "simple include",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"file1.lisp\""}, Value: "file1.lisp"}},
				},
			},
			expected: []ast.Node{
				&ast.DefineExpr{
					Token:  token.Token{TokenType: token.ATOM, Value: "define"},
					Symbol: "x",
					Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"},
				},
			},
			err: "",
		},
		{
			name: "include with cycle",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"cycle.lisp\""}, Value: "cycle.lisp"}},
				},
			},
			expected: nil,
			err:      "inclusion cycle detected for file cycle.lisp",
		},
		{
			name: "read file error",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"nonexistent.lisp\""}, Value: "nonexistent.lisp"}},
				},
			},
			expected: nil,
			err:      "failed to read file nonexistent.lisp",
		},
		{
			name: "scanner error",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"invalid_scan.lisp\""}, Value: "invalid_scan.lisp"}},
				},
			},
			expected: nil,
			err:      "failed to scan file invalid_scan.lisp",
		},
		{
			name: "parse error",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"invalid_parse.lisp\""}, Value: "invalid_parse.lisp"}},
				},
			},
			expected: nil,
			err:      "failed to parse file invalid_parse.lisp",
		},
		{
			name: "nested includes with duplicate",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"fileX.lisp\""}, Value: "fileX.lisp"}},
				},
			},
			expected: []ast.Node{
				&ast.DefineExpr{
					Token:  token.Token{TokenType: token.ATOM, Value: "define"},
					Symbol: "z",
					Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "44"}, Value: "44"},
				},
			},
			err: "",
		},
		{
			name: "callable not include",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "notInclude"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "notInclude"}, Value: "notInclude"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"file1.lisp\""}, Value: "file1.lisp"}},
				},
			},
			expected: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "notInclude"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "notInclude"}, Value: "notInclude"},
					Args:     []ast.Node{&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"file1.lisp\""}, Value: "file1.lisp"}},
				},
			},
			err: "",
		},
		{
			name: "include with wrong number of args",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args: []ast.Node{
						&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"file1.lisp\""}, Value: "file1.lisp"},
						&ast.StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"file2.lisp\""}, Value: "file2.lisp"},
					},
				},
			},
			expected: nil,
			err:      "include expects exactly one argument",
		},
		{
			name: "include with non-string argument",
			input: []ast.Node{
				&ast.CallExpr{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					Callable: &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "include"}, Value: "include"},
					Args:     []ast.Node{&ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}},
				},
			},
			expected: nil,
			err:      "include expects a string argument",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Include(tt.input)
			if tt.err != "" {
				if err == nil || !strings.Contains(err.Error(), tt.err) {
					t.Errorf("expected error %s, got %v", tt.err, err)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d nodes, got %d", len(tt.expected), len(result))
				return
			}
			for i, node := range result {
				if node.String() != tt.expected[i].String() {
					t.Errorf("expected node %s, got %s", tt.expected[i].String(), node.String())
				}
			}
		})
	}
}
