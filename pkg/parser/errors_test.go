package parser

import (
	"errors"
	"testing"

	"github.com/bassosimone/buresu/pkg/token"
)

func TestIsErrIncompleteInput(t *testing.T) {
	err := ErrIncompleteInput{Err: errors.New("incomplete input")}
	if !IsErrIncompleteInput(err) {
		t.Errorf("expected true, got false")
	}

	otherErr := errors.New("some other error")
	if IsErrIncompleteInput(otherErr) {
		t.Errorf("expected false, got true")
	}
}

func TestErrIncompleteInput_Error(t *testing.T) {
	err := ErrIncompleteInput{Err: errors.New("incomplete input")}
	expected := "incomplete input"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

func TestErrIncompleteInput_Unwrap(t *testing.T) {
	innerErr := errors.New("incomplete input")
	err := ErrIncompleteInput{Err: innerErr}
	if !errors.Is(err, innerErr) {
		t.Errorf("expected true, got false")
	}
}

func TestError_Error(t *testing.T) {
	tok := token.Token{TokenPos: token.Position{FileName: "test.go", LineNumber: 1, LineColumn: 1}}
	err := &Error{Tok: tok, Message: "test error"}
	expected := "test.go:1:1: parser: test error"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

func TestNewError(t *testing.T) {
	tok := token.Token{TokenPos: token.Position{FileName: "test.go", LineNumber: 1, LineColumn: 1}}
	err := newError(tok, "test error %d", 123)
	expected := "test.go:1:1: parser: test error 123"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}
