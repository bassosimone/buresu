package scanner_test

import (
	"strings"
	"testing"

	"github.com/bassosimone/buresu/scanner"
	"github.com/bassosimone/buresu/token"
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
		if token != expected[i] {
			t.Errorf("expected token %v, got %v", expected[i], token)
		}
	}
}

func TestScan_EmptyInput(t *testing.T) {
	input := ""
	expected := []token.Token{
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 0,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_SingleCharacterTokens(t *testing.T) {
	input := "()"
	expected := []token.Token{
		{
			TokenType: token.OPEN,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "(",
		},
		{
			TokenType: token.CLOSE,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 2,
			},
			Value: ")",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 2,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_NumberToken(t *testing.T) {
	input := "123"
	expected := []token.Token{
		{
			TokenType: token.NUMBER,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "123",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 3,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_StringToken(t *testing.T) {
	input := `"hello"`
	expected := []token.Token{
		{
			TokenType: token.STRING,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "hello",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 7,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_AtomToken(t *testing.T) {
	input := "atom"
	expected := []token.Token{
		{
			TokenType: token.ATOM,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "atom",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 4,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_Comment(t *testing.T) {
	input := "; this is a comment\natom"
	expected := []token.Token{
		{
			TokenType: token.ATOM,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 2,
				LineColumn: 1,
			},
			Value: "atom",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 2,
				LineColumn: 4,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_ErrorUnexpectedChar(t *testing.T) {
	input := "@"
	expectedErrorMsg := "test:1:1: scanner: unexpected char at top level: U+0040 '@'"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_ErrorMultipleDotsInNumber(t *testing.T) {
	input := "12.34.56"
	expectedErrorMsg := "test:1:1: scanner: multiple dots in number literal"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_ErrorUnterminatedString(t *testing.T) {
	input := `"hello`
	expectedErrorMsg := "test:1:1: scanner: expected '\"', found: EOF"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_ErrorUnknownEscapeSequence(t *testing.T) {
	input := `"hello\q"`
	expectedErrorMsg := "test:1:1: scanner: unknown escape sequence: \\U+0071 'q'"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_ErrorMidwayAtom(t *testing.T) {
	input := "atom@"
	expectedErrorMsg := "test:1:1: scanner: expected [ ()], found: U+0040 '@'"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_NumberMidwayAtom(t *testing.T) {
	input := "1234@"
	expectedErrorMsg := "test:1:1: scanner: expected [ ()], found: U+0040 '@'"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_NumberFollowedByCloseParen(t *testing.T) {
	input := "123)"
	expected := []token.Token{
		{
			TokenType: token.NUMBER,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "123",
		},
		{
			TokenType: token.CLOSE,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 4,
			},
			Value: ")",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 4,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_ErrorNonPrintableCharInString(t *testing.T) {
	input := "\"hello\x01world\""
	expectedErrorMsg := "test:1:1: scanner: expected printable character, found: U+0001 '\x01'"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}
func TestScan_StringWithEscapeSequences(t *testing.T) {
	input := `"hello\nworld\t\"escaped\"\rnew\\line"`
	expected := []token.Token{
		{
			TokenType: token.STRING,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "hello\nworld\t\"escaped\"\rnew\\line",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 38,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_AtomFollowedByCloseParen(t *testing.T) {
	input := "atom)"
	expected := []token.Token{
		{
			TokenType: token.ATOM,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "atom",
		},
		{
			TokenType: token.CLOSE,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 5,
			},
			Value: ")",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 5,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}

func TestScan_ErrorEOFInEscapeSequence(t *testing.T) {
	input := `"hello\`
	expectedErrorMsg := "test:1:1: scanner: expected [nrt\"\\\\] character, found: EOF"
	runScanTest(t, input, nil, true, expectedErrorMsg)
}

func TestScan_IfTrueExpression(t *testing.T) {
	input := "(if true 1 0)"
	expected := []token.Token{
		{
			TokenType: token.OPEN,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 1,
			},
			Value: "(",
		},
		{
			TokenType: token.ATOM,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 2,
			},
			Value: "if",
		},
		{
			TokenType: token.ATOM,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 5,
			},
			Value: "true",
		},
		{
			TokenType: token.NUMBER,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 10,
			},
			Value: "1",
		},
		{
			TokenType: token.NUMBER,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 12,
			},
			Value: "0",
		},
		{
			TokenType: token.CLOSE,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 13,
			},
			Value: ")",
		},
		{
			TokenType: token.EOF,
			TokenPos: token.Position{
				FileName:   "test",
				LineNumber: 1,
				LineColumn: 13,
			},
			Value: "",
		},
	}
	runScanTest(t, input, expected, false, "")
}
