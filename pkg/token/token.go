// SPDX-License-Identifier: GPL-3.0-or-later

// Package token contains the scanner token definitions.
//
// The scanner package takes in input source code and emits
// a sequence of tokens then consumed by the parser.
package token

import "fmt"

// TokenType represents the type of a token in the language.
type TokenType string

const (
	// ATOM represents an atomic token, which typically is
	// bound to a keyword or a symbol name.
	ATOM TokenType = "ATOM"

	// CLOSE represents a closing parenthesis token.
	CLOSE TokenType = "CLOSE"

	// EOF represents the end of file token.
	EOF TokenType = "EOF"

	// NUMBER represents a numeric token.
	NUMBER TokenType = "NUMBER"

	// OPEN represents an open parenthesis token.
	OPEN TokenType = "OPEN"

	// STRING represents a string token.
	STRING TokenType = "STRING"

	// ELLIPSIS represents an ellipsis token.
	ELLIPSIS TokenType = "ELLIPSIS"
)

// Position represents the position of a token in the source code.
type Position struct {
	FileName   string
	LineNumber int
	LineColumn int
}

// String returns the string representation of the Position.
func (p Position) String() string {
	return fmt.Sprintf("%s:%d:%d", p.FileName, p.LineNumber, p.LineColumn)
}

// Token represents a token with its type, position, and value.
type Token struct {
	TokenPos  Position
	TokenType TokenType
	Value     string
}

// Clone creates a copy of the Token.
func (t Token) Clone() Token {
	return Token{
		TokenPos:  t.TokenPos,
		TokenType: t.TokenType,
		Value:     t.Value,
	}
}
