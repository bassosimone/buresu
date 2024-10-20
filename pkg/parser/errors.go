// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"errors"
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// IsErrIncompleteInput returns true if the error is an ErrIncompleteInput.
func IsErrIncompleteInput(err error) bool {
	var value ErrIncompleteInput
	return errors.As(err, &value)
}

// ErrIncompleteInput is an error decorator indicating that the
// user has provided us with incomplete input.
type ErrIncompleteInput struct {
	Err error
}

// Error returns the error message.
func (err ErrIncompleteInput) Error() string {
	return err.Err.Error()
}

// Unwrap returns the wrapped error.
func (err ErrIncompleteInput) Unwrap() error {
	return err.Err
}

// Parse processes the provided tokens and returns a slice
// of AST nodes or an error if parsing fails.
func Parse(tokens []token.Token) (nodes []ast.Node, err error) {
	return newParser(tokens).Parse()
}

// Error represents a parsing error with position and message.
type Error struct {
	Tok     token.Token
	Message string
}

// Error returns the error message with file position details.
func (e *Error) Error() string {
	return fmt.Sprintf(
		"%s:%d:%d: parser: %s",
		e.Tok.TokenPos.FileName,
		e.Tok.TokenPos.LineNumber,
		e.Tok.TokenPos.LineColumn,
		e.Message,
	)
}

// newError formats and returns a new parser error including the token context.
func newError(tok token.Token, format string, args ...any) error {
	return &Error{Tok: tok, Message: fmt.Sprintf(format, args...)}
}
