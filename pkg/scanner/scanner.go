// SPDX-License-Identifier: GPL-3.0-or-later

// Package scanner contains the language scanner.
//
// The scanner takes in input source code and emits a sequence
// of tokens (a type defined in the token package).
package scanner

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/bassosimone/buresu/pkg/token"
)

// Scan scans the given file and returns a slice of tokens or an error.
func Scan(filename string, file io.Reader) ([]token.Token, error) {
	return newScanner(filename, file).Scan()
}

// Error represents a scanning error with position and message.
type Error struct {
	Pos     token.Position
	Message string
}

// Error returns the error message with file position details.
func (e *Error) Error() string {
	return fmt.Sprintf(
		"%s:%d:%d: scanner: %s",
		e.Pos.FileName,
		e.Pos.LineNumber,
		e.Pos.LineColumn,
		e.Message,
	)
}

// scanner is the internal scanner implementation.
type scanner struct {
	filename string
	reader   *bufio.Reader
	lineno   int
	col      int
	current  rune
}

// newScanner creates a new scanner for the given filename and reader.
func newScanner(filename string, reader io.Reader) *scanner {
	scanner := &scanner{
		filename: filename,
		reader:   bufio.NewReader(reader),
		lineno:   1,
		col:      0,
		current:  0,
	}
	scanner.advance() // load the first rune
	return scanner
}

// position returns the current position in the source code.
func (s *scanner) position() token.Position {
	return token.Position{
		FileName:   s.filename,
		LineNumber: s.lineno,
		LineColumn: s.col,
	}
}

// advance reads the next rune from the input and updates the scanner's state.
//
// Note: advance is a no-op if the scanner has reached the end of the input.
func (s *scanner) advance() {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		s.current = 0
		return
	}
	s.current = r
	s.col++
	if s.current == '\n' {
		s.lineno++
		s.col = 0
	}
}

// newError creates a new scanning error with the given position and message.
func newError(pos token.Position, message string) error {
	return &Error{
		Pos:     pos,
		Message: message,
	}
}

// newToken creates a new token with the given type, position, and value.
func (s *scanner) newToken(tokenType token.TokenType, pos token.Position, value string) token.Token {
	return token.Token{TokenType: tokenType, TokenPos: pos, Value: value}
}

// Scan scans the input and returns a slice of tokens or an error.
func (s *scanner) Scan() ([]token.Token, error) {
	var tokens []token.Token
	for {
		pos := s.position()
		chr := s.current

		switch {
		case chr == 0:
			tokens = append(tokens, s.newToken(token.EOF, pos, ""))
			return tokens, nil

		case unicode.IsSpace(chr):
			s.advance()
			continue

		case chr == '(':
			tokens = append(tokens, s.newToken(token.OPEN, pos, string(chr)))
			s.advance()
			continue

		case chr == ')':
			tokens = append(tokens, s.newToken(token.CLOSE, pos, string(chr)))
			s.advance()
			continue

		case chr == ';':
			s.skipComment()
			continue

		case chr == '-':
			if s.lookaheadIsDigit() {
				token, err := s.scanNumber(pos)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, token)
				continue
			}
			token, err := s.scanSymbolicAtom(pos)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue

		case unicode.IsDigit(chr):
			token, err := s.scanNumber(pos)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue

		case chr == '"':
			token, err := s.scanString(pos)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue

		case unicode.IsLetter(chr) || chr == '_':
			token, err := s.scanAlphabeticAtom(pos)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue

		default:
			token, err := s.scanSymbolicAtom(pos)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue
		}
	}
}

// skipComment skips over a comment in the input.
func (s *scanner) skipComment() {
	for s.current != '\n' && s.current != 0 {
		s.advance()
	}
	s.advance()
}

// scanNumber scans a number token from the input.
func (s *scanner) scanNumber(pos token.Position) (token.Token, error) {
	var value strings.Builder
	value.WriteRune(s.current)
	s.advance()
	seendot := false

	for {
		chr := s.current
		if chr == 0 {
			break
		} else if chr == '.' {
			if seendot {
				return token.Token{}, newError(pos, "multiple dots in number literal")
			}
			seendot = true
		} else if !unicode.IsDigit(chr) {
			if !unicode.IsSpace(chr) && !strings.ContainsRune("()", chr) {
				err := newError(pos, fmt.Sprintf("expected [ ()], found: %U '%c'", chr, chr))
				return token.Token{}, err
			}
			break
		}
		value.WriteRune(chr)
		s.advance()
	}

	return s.newToken(token.NUMBER, pos, value.String()), nil
}

// scanString scans a string token from the input.
func (s *scanner) scanString(pos token.Position) (token.Token, error) {
	var value strings.Builder

	for {
		s.advance()
		chr := s.current
		if chr == 0 {
			return token.Token{}, newError(pos, "expected '\"', found: EOF")
		}

		if chr == '\\' {
			esc, err := s.scanEscapeSequence(pos)
			if err != nil {
				return token.Token{}, err
			}
			value.WriteString(esc)
			continue
		}

		if chr == '"' {
			s.advance()
			break
		}

		// Allow all valid Unicode characters
		//
		// We support multiline strings
		if !unicode.IsPrint(chr) && chr != '\n' && chr != '\t' {
			err := newError(pos, fmt.Sprintf("expected printable character, found: %U '%c'", chr, chr))
			return token.Token{}, err
		}

		value.WriteRune(chr)
	}

	return s.newToken(token.STRING, pos, value.String()), nil
}

// scanEscapeSequence scans an escape sequence from the input.
func (s *scanner) scanEscapeSequence(pos token.Position) (string, error) {
	s.advance()
	chr := s.current
	if chr == 0 {
		return "", newError(pos, `expected [nrt"\\] character, found: EOF`)
	}
	switch chr {
	case 'n':
		return "\n", nil
	case 'r':
		return "\r", nil
	case 't':
		return "\t", nil
	case '"':
		return "\"", nil
	case '\\':
		return "\\", nil
	default:
		return "", newError(pos, fmt.Sprintf("unknown escape sequence: \\%U '%c'", chr, chr))
	}
}

// scanAlphabeticAtom scans an alphabetic atom token from the input.
func (s *scanner) scanAlphabeticAtom(pos token.Position) (token.Token, error) {
	var value strings.Builder
	value.WriteRune(s.current)
	s.advance()

	for {
		chr := s.current
		if chr == 0 {
			break
		}
		if !unicode.IsLetter(chr) && !unicode.IsDigit(chr) && chr != '_' && chr != '-' {
			if chr == '!' || chr == '?' {
				value.WriteRune(chr)
				s.advance()
				// Check if the next character is a valid separator
				// after the legitimate `?!` ending rune
				chr = s.current
			}
			if chr != 0 && !unicode.IsSpace(chr) && !strings.ContainsRune("()", chr) {
				err := newError(pos, fmt.Sprintf("expected [ ()], found: %U '%c'", chr, chr))
				return token.Token{}, err
			}
			break
		}
		value.WriteRune(chr)
		s.advance()
	}

	return s.newToken(token.ATOM, pos, value.String()), nil
}

// scanSymbolicAtom scans a symbolic atom token from the input.
func (s *scanner) scanSymbolicAtom(pos token.Position) (token.Token, error) {
	var tok token.Token
	switch s.current {
	case '+':
		s.advance()
		tok = s.newToken(token.ATOM, pos, "+")

	case '-':
		s.advance()
		tok = s.newToken(token.ATOM, pos, "-")

	case '*':
		s.advance()
		tok = s.newToken(token.ATOM, pos, "*")

	case '/':
		s.advance()
		tok = s.newToken(token.ATOM, pos, "/")

	case '.':
		s.advance()
		tok = s.newToken(token.ATOM, pos, ".")

	case '=':
		s.advance()
		if s.current == '=' {
			s.advance()
			tok = s.newToken(token.ATOM, pos, "==")
		} else {
			tok = s.newToken(token.ATOM, pos, "=")
		}

	case '<':
		s.advance()
		if s.current == '=' {
			s.advance()
			if s.current == '>' {
				s.advance()
				tok = s.newToken(token.ATOM, pos, "<=>")
			} else {
				tok = s.newToken(token.ATOM, pos, "<=")
			}
		} else {
			tok = s.newToken(token.ATOM, pos, "<")
		}

	case '>':
		s.advance()
		if s.current == '=' {
			s.advance()
			tok = s.newToken(token.ATOM, pos, ">=")
		} else {
			tok = s.newToken(token.ATOM, pos, ">")
		}

	case ':':
		s.advance()
		if s.current == ':' {
			s.advance()
			tok = s.newToken(token.ATOM, pos, "::")
		} else {
			tok = s.newToken(token.ATOM, pos, ":")
		}

	default:
		return token.Token{}, newError(pos, fmt.Sprintf(
			"unexpected symbolic atom: %U '%c'", s.current, s.current))
	}

	if s.current != 0 && !unicode.IsSpace(s.current) &&
		!strings.ContainsRune("()", s.current) {
		err := newError(pos, fmt.Sprintf(
			"expected [ ()], found: %U '%c'", s.current, s.current))
		return token.Token{}, err
	}

	return tok, nil
}

// lookaheadIsDigit checks if the next character is a digit.
func (s *scanner) lookaheadIsDigit() bool {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		return false
	}
	defer s.reader.UnreadRune()
	return unicode.IsDigit(r)
}
