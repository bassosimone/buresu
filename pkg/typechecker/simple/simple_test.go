// SPDX-License-Identifier: GPL-3.0-or-later

package simple_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"

	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/typechecker/simple"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

func TestCheck(t *testing.T) {
	testdataDir := "testdata"
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
				inputCode      []byte
				expectedOutput []byte
				expectedError  []byte
			)

			for _, file := range archive.Files {
				switch file.Name {
				case "input.txt":
					inputCode = file.Data
				case "expected_output.txt":
					expectedOutput = bytes.TrimSpace(file.Data)
				case "expected_error.txt":
					expectedError = bytes.TrimSpace(file.Data)
				}
			}

			// Scan and parse the input code
			tokens, err := scanner.Scan("input.code", bytes.NewReader(inputCode))
			if err != nil {
				t.Fatalf("failed to scan input code: %v", err)
			}
			nodes, err := parser.Parse(tokens)
			if err != nil {
				t.Fatalf("failed to parse input code: %v", err)
			}

			// Evaluate the parsed nodes
			ctx := context.Background()
			env := simple.NewGlobalEnvironment(os.Stdout)
			var (
				results []string
				result  visitor.Type
			)
			for _, node := range nodes {
				result, err = simple.Check(ctx, env, node)
				if err != nil {
					t.Log("err:", err.Error())
					break
				}
				t.Log("result:", result.String())
				results = append(results, result.String())
			}

			if expectedError != nil {
				// If an error is expected, check if the actual error matches the expected error
				if err == nil {
					t.Fatalf("expected an error but got none")
				}
				actualError := err.Error()
				if diff := cmp.Diff(string(expectedError), actualError); diff != "" {
					t.Errorf("error mismatch (-expected +got):\n%s", diff)
				}
				return
			}

			// If no error is expected, check if the result matches the expected output
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actualOutput := strings.Join(results, "\n")
			if diff := cmp.Diff(string(expectedOutput), actualOutput); diff != "" {
				t.Errorf("output mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}
