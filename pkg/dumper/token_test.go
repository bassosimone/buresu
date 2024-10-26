// SPDX-License-Identifier: GPL-3.0-or-later

package dumper_test

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/bassosimone/buresu/internal/txtartesting"
	"github.com/bassosimone/buresu/pkg/dumper"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestDumpTokens(t *testing.T) {
	testCases, err := txtartesting.LoadTestCases(filepath.Join("testdata", "token"))
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tokens, err := scanner.Scan("input.txt", bytes.NewReader([]byte(tc.Input)))
			if err != nil {
				t.Fatal(err)
			}

			var buf bytes.Buffer
			err = dumper.DumpTokens(&buf, tokens)
			if err != nil {
				t.Fatalf("failed to dump tokens: %v", err)
			}
			if !json.Valid(buf.Bytes()) {
				t.Fatalf("invalid JSON output: %s", buf.String())
			}

			if err := tc.CompareTextOutput(buf.String()); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDumpTokensWriterError(t *testing.T) {
	errWriter := &failingWriter{}
	tokens := []token.Token{
		{TokenType: token.ATOM, Value: "example"},
	}

	err := dumper.DumpTokens(errWriter, tokens)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedErr := "failed to dump tokens: write error"
	if err.Error() != expectedErr {
		t.Errorf("expected %q, got %q", expectedErr, err.Error())
	}
}
