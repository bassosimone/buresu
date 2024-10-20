// SPDX-License-Identifier: GPL-3.0-or-later

package scanner_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"

	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestErrorString(t *testing.T) {
	err := &scanner.Error{
		Pos: token.Position{
			FileName:   "testfile",
			LineNumber: 10,
			LineColumn: 5,
		},
		Message: "unexpected character",
	}
	expected := "testfile:10:5: scanner: unexpected character"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}

func TestScanner(t *testing.T) {
	testdataDir := filepath.Join("testdata")
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
				inputSource    string
				expectedOutput []byte
				expectedError  string
			)

			for _, file := range archive.Files {
				switch file.Name {
				case "input.txt":
					inputSource = string(file.Data)
				case "expected_tokens.json":
					expectedOutput = file.Data
				case "expected_error.txt":
					expectedError = string(file.Data)
				}
			}

			// Run the scanner
			tokens, err := scanner.Scan("input.txt", bytes.NewReader([]byte(inputSource)))
			if expectedError != "" {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				if diff := cmp.Diff(strings.TrimSpace(expectedError), err.Error()); diff != "" {
					t.Errorf("mismatch (-expected +got):\n%s", diff)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Serialize the tokens to JSON
			var buf bytes.Buffer
			encoder := json.NewEncoder(&buf)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(tokens); err != nil {
				t.Fatalf("failed to encode tokens: %v", err)
			}

			// Compare the serialized output with the expected output
			if diff := cmp.Diff(string(expectedOutput), buf.String()); diff != "" {
				t.Errorf("mismatch (-expected +got):\n%s", diff)
			}
		})
	}
}

func runScanTest(t *testing.T, input string, expected []token.Token, expectError bool, expectedErrorMsg string) {
	tokens, err := scanner.Scan("test", strings.NewReader(input))
	if expectError {
		if err == nil {
			t.Fatalf("expected error, got none")
		}
		if err.Error() != expectedErrorMsg {
			t.Fatalf("expected error message %q, got %q", expectedErrorMsg, err.Error())
		}
		return
	}
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, token := range tokens {
		if diff := cmp.Diff(expected[i], token); diff != "" {
			t.Errorf(diff)
		}
	}
}

func TestScan_ErrorNonPrintableCharInString(t *testing.T) {
	input := "\"hello\x01world\""
	expectedErrorMsg := "test:1:1: scanner: expected printable character, found: U+0001 '\x01'"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_ErrorEOFInEscapeSequence(t *testing.T) {
	input := `"hello\`
	expectedErrorMsg := "test:1:1: scanner: expected [nrt\"\\\\] character, found: EOF"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_AlphabeticAtomEOF(t *testing.T) {
	expected := []token.Token{
		{
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			TokenType: token.ATOM,
			Value:     "abc",
		},
		{
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 3,
			},
			TokenType: token.EOF,
			Value:     "",
		},
	}
	input := `abc`
	runScanTest(t, input, expected, false, "")
}

func TestScan_NumberEOF(t *testing.T) {
	expected := []token.Token{
		{
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			TokenType: token.NUMBER,
			Value:     "012",
		},
		{
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 3,
			},
			TokenType: token.EOF,
			Value:     "",
		},
	}
	input := `012`
	runScanTest(t, input, expected, false, "")
}

func TestScan_LookAheadIsDigit(t *testing.T) {
	expected := []token.Token{
		{
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			TokenType: token.ATOM,
			Value:     "-",
		},
		{
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			TokenType: token.EOF,
			Value:     "",
		},
	}
	input := `-`
	runScanTest(t, input, expected, false, "")
}
