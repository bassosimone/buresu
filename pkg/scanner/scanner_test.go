// SPDX-License-Identifier: GPL-3.0-or-later

package scanner_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bassosimone/buresu/internal/txtartesting"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
	"github.com/google/go-cmp/cmp"
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
	testCases, err := txtartesting.LoadTestCases("testdata")
	if err != nil {
		t.Fatalf("failed to load test cases: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Run the scanner
			tokens, err := scanner.Scan("input.txt", bytes.NewReader([]byte(tc.Input)))
			if tc.Error != "" {
				if err := tc.CompareError(err); err != nil {
					t.Fatal(err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Compare the serialized output with the expected output
			if err := tc.CompareJSONOutput(tokens); err != nil {
				t.Fatal(err)
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
