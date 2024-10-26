package includer

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// mockReadFile is a mock implementation of the os.ReadFile function.
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

	default:
		return nil, errors.New("file not found")
	}
}

func TestInclude(t *testing.T) {
	// Override the os.ReadFile function with the mock and restore
	// the original function after the test
	readFile = mockReadFile
	defer func() { readFile = os.ReadFile }()

	// Allows to temporarily skip all tests such that we can
	// only run the ones marked as overrideSkip
	skipAll := false

	tests := []struct {
		name         string
		input        []ast.Node
		expected     []ast.Node
		err          string
		overrideSkip bool
	}{
		{
			name: "simple include",
			input: []ast.Node{
				&ast.IncludeStmt{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					FilePath: "file1.lisp",
				},
			},
			expected: []ast.Node{
				&ast.DefineExpr{
					Token:  token.Token{TokenType: token.ATOM, Value: "define"},
					Symbol: "x",
					Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"},
				},
			},
			err:          "",
			overrideSkip: false,
		},

		{
			name: "include with cycle",
			input: []ast.Node{
				&ast.IncludeStmt{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					FilePath: "cycle.lisp",
				},
			},
			expected:     nil,
			err:          "inclusion cycle detected for file cycle.lisp",
			overrideSkip: false,
		},

		{
			name: "read file error",
			input: []ast.Node{
				&ast.IncludeStmt{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					FilePath: "nonexistent.lisp",
				},
			},
			expected:     nil,
			err:          "failed to read file nonexistent.lisp",
			overrideSkip: false,
		},

		{
			name: "scanner error",
			input: []ast.Node{
				&ast.IncludeStmt{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					FilePath: "invalid_scan.lisp",
				},
			},
			expected:     nil,
			err:          "failed to scan file invalid_scan.lisp",
			overrideSkip: false,
		},

		{
			name: "parse error",
			input: []ast.Node{
				&ast.IncludeStmt{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					FilePath: "invalid_parse.lisp",
				},
			},
			expected:     nil,
			err:          "failed to parse file invalid_parse.lisp",
			overrideSkip: false,
		},

		{
			name: "nested includes with duplicate",
			input: []ast.Node{
				&ast.IncludeStmt{
					Token:    token.Token{TokenType: token.ATOM, Value: "include"},
					FilePath: "fileX.lisp",
				},
			},
			expected: []ast.Node{
				&ast.DefineExpr{
					Token:  token.Token{TokenType: token.ATOM, Value: "define"},
					Symbol: "z",
					Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "44"}, Value: "44"},
				},
			},
			err:          "",
			overrideSkip: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if skipAll && !tt.overrideSkip {
				t.Skip("skipAll is true and overrideSkip is false")
			}
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

			for idx, node := range result {
				if node.String() != tt.expected[idx].String() {
					t.Errorf("expected node %s, got %s", tt.expected[idx].String(), node.String())
				}
			}
		})
	}
}
