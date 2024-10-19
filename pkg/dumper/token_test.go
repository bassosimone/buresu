package dumper_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"

	"github.com/bassosimone/buresu/pkg/dumper"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

func TestDumpTokens(t *testing.T) {
	testdataDir := filepath.Join("testdata", "token")
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
				inputTokens    []token.Token
				expectedOutput []byte
			)

			for _, file := range archive.Files {
				switch file.Name {
				case "input.txt":
					tokens, err := scanner.Scan("input.txt", bytes.NewReader(file.Data))
					if err != nil {
						t.Fatal(err)
					}
					inputTokens = tokens
				case "expected_tokens.json":
					expectedOutput = file.Data
				}
			}

			// Serialize the input tokens to JSON
			var buf bytes.Buffer
			err = dumper.DumpTokens(&buf, inputTokens)
			if err != nil {
				t.Fatalf("failed to dump tokens: %v", err)
			}

			// Compare the serialized output with the expected output
			if diff := cmp.Diff(string(expectedOutput), buf.String()); diff != "" {
				t.Errorf("mismatch (-expected +got):\n%s", diff)
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
