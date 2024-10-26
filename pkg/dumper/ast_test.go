// SPDX-License-Identifier: GPL-3.0-or-later

package dumper_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bassosimone/buresu/internal/txtartesting"
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/dumper"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestDumpAST(t *testing.T) {
	t.Run("dumping the AST", func(t *testing.T) {
		testCases, err := txtartesting.LoadTestCases("testdata/ast")
		if err != nil {
			t.Fatal(err)
		}

		for _, tc := range testCases {
			t.Run(tc.Name, func(t *testing.T) {
				tokens, err := scanner.Scan("input.ast", bytes.NewReader([]byte(tc.Input)))
				if err != nil {
					t.Fatal(err)
				}
				nodes, err := parser.Parse(tokens)
				if err != nil {
					t.Fatal(err)
				}

				var buf bytes.Buffer
				err = dumper.DumpAST(&buf, nodes)
				if err != nil {
					t.Fatalf("failed to dump AST: %v", err)
				}

				if !json.Valid(buf.Bytes()) {
					t.Fatalf("invalid JSON output: %s", buf.String())
				}
				if err := tc.CompareTextOutput(buf.String()); err != nil {
					t.Error(err)
				}
			})
		}
	})

	t.Run("writer error", func(t *testing.T) {
		errWriter := &failingWriter{}
		nodes := []ast.Node{
			&ast.BlockExpr{
				Token: token.Token{TokenType: token.ATOM, Value: "block"},
				Exprs: []ast.Node{
					&ast.DefineExpr{
						Token:  token.Token{TokenType: token.ATOM, Value: "define"},
						Symbol: "x",
						Expr:   &ast.IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"},
					},
					&ast.ReturnStmt{
						Token: token.Token{TokenType: token.ATOM, Value: "return"},
						Expr:  &ast.SymbolName{Token: token.Token{TokenType: token.ATOM, Value: "x"}, Value: "x"},
					},
				},
			},
		}

		err := dumper.DumpAST(errWriter, nodes)
		if err == nil {
			t.Fatalf("expected an error, but got nil")
		}
		if err.Error() != "failed to dump AST: write error" {
			t.Fatalf("unexpected error message: %v", err)
		}
	})
}

type failingWriter struct{}

func (fw *failingWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write error")
}
