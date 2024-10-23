// SPDX-License-Identifier: GPL-3.0-or-later

package dumper_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/dumper"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestDumpAST(t *testing.T) {
	testdataDir := filepath.Join("testdata", "ast")
	files, err := os.ReadDir(testdataDir)
	if err != nil {
		t.Fatalf("failed to read testdata directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".txtar" {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			archivePath := filepath.Join(testdataDir, file.Name())
			archiveData, err := os.ReadFile(archivePath)
			if err != nil {
				t.Fatalf("failed to read txtar file: %v", err)
			}

			archive := txtar.Parse(archiveData)

			var (
				inputAST       []ast.Node
				expectedOutput []byte
			)

			for _, file := range archive.Files {
				switch file.Name {
				case "input.ast":
					tokens, err := scanner.Scan("input.ast", bytes.NewReader(file.Data))
					if err != nil {
						t.Fatal(err)
					}
					nodes, err := parser.Parse(tokens)
					if err != nil {
						t.Fatal(err)
					}
					inputAST = nodes
				case "expected_output.json":
					expectedOutput = file.Data
				}
			}

			// Serialize the input AST to JSON
			var buf bytes.Buffer
			err = dumper.DumpAST(&buf, inputAST)
			if err != nil {
				t.Fatalf("failed to dump AST: %v", err)
			}

			// Compare the serialized output with the expected output
			if !json.Valid(buf.Bytes()) {
				t.Fatalf("invalid JSON output: %s", buf.String())
			}
			if diff := cmp.Diff(string(expectedOutput), buf.String()); diff != "" {
				t.Errorf("mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func TestDumpASTWriterError(t *testing.T) {
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
}

type failingWriter struct{}

func (fw *failingWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write error")
}
