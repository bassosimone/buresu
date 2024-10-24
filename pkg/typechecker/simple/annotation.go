package simple

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/bassosimone/buresu/pkg/token"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// ErrNoTypeAnnotationFound is returned when no type annotation is found.
var ErrNoTypeAnnotationFound = errors.New("no type annotation found")

// ParseTypeAnnotationFromDocs parses the lambda docs string searching for a type annotation.
func ParseTypeAnnotationFromDocs(docs string) (*Callable, error) {
	// for each line search for a line starting with `::`
	// if found, parse the callable, then make sure there
	// are no more annotations in the documentation.
	var callable *Callable
	for _, line := range strings.Split(docs, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "::") {
			if callable != nil {
				return nil, errors.New("multiple type annotations found")
			}
			var err error
			callable, err = ParseTypeAnnotationFromString(strings.TrimPrefix(line, "::"))
			if err != nil {
				return nil, err
			}
		}
	}
	if callable == nil {
		return nil, ErrNoTypeAnnotationFound
	}
	return callable, nil
}

// ParseTypeAnnotationString parses a type annotation from a string.
func ParseTypeAnnotationFromString(input string) (*Callable, error) {
	tokens, err := scanAnnotation(strings.NewReader(input))
	if err != nil {
		return nil, err
	}
	return newAnnotationParser(tokens).Parse()
}

// scanAnnotation scans the given input and returns a slice of tokens or an error.
//
// The tokens are as follows:
//
// - ATOM: [A-Za-z][A-Za-z0-9_]+
// - CLOSE: close parenthesis
// - EOF: end of file
// - OPEN: open parenthesis
func scanAnnotation(input io.Reader) ([]token.Token, error) {
	return newAnnotationScanner(input).Scan()
}

// annotationScanner is the internal scanner implementation for annotations.
type annotationScanner struct {
	reader  *bufio.Reader
	current rune
	pos     int
}

// newAnnotationScanner creates a new scanner for annotations.
func newAnnotationScanner(input io.Reader) *annotationScanner {
	scanner := &annotationScanner{
		reader: bufio.NewReader(input),
	}
	scanner.advance()
	return scanner
}

// Scan scans the input and returns a slice of tokens or an error.
func (s *annotationScanner) Scan() ([]token.Token, error) {
	var tokens []token.Token
	for {
		pos := token.Position{FileName: "<annotation>", LineNumber: 1, LineColumn: s.pos}
		switch {
		case s.current == 0:
			tokens = append(tokens, token.Token{TokenPos: pos, TokenType: token.EOF})
			return tokens, nil

		case s.current == '(':
			tokens = append(tokens, token.Token{TokenPos: pos, TokenType: token.OPEN, Value: "("})
			s.advance()

		case s.current == ')':
			tokens = append(tokens, token.Token{TokenPos: pos, TokenType: token.CLOSE, Value: ")"})
			s.advance()

		case unicode.IsSpace(s.current):
			s.advance()

		case unicode.IsLetter(s.current):
			tokens = append(tokens, s.scanAtom(pos))

		default:
			return nil, fmt.Errorf("annotation scanner: unexpected character: %c", s.current)
		}
	}
}

// advance advances the scanner by one rune.
func (s *annotationScanner) advance() {
	r, _, err := s.reader.ReadRune()
	if err != nil {
		s.current = 0
		return
	}
	s.current = r
	s.pos++
}

// scanAtom scans an atom from the input.
func (s *annotationScanner) scanAtom(pos token.Position) token.Token {
	var value strings.Builder
	for unicode.IsLetter(s.current) || unicode.IsDigit(s.current) || s.current == '_' {
		value.WriteRune(s.current)
		s.advance()
	}
	return token.Token{TokenPos: pos, TokenType: token.ATOM, Value: value.String()}
}

// annotationParser is the internal parser implementation for annotations.
type annotationParser struct {
	tokens  []token.Token
	current int
}

// newAnnotationParser creates a new parser for annotations.
func newAnnotationParser(tokens []token.Token) *annotationParser {
	return &annotationParser{tokens: tokens, current: 0}
}

// Parse parses a lambda annotation and returns a callable or an error.
//
// The grammar is as follows:
//
//	<annotation> ::= <callable> EOF
//
//	<callable> ::= OPEN "Callable" OPEN <expr>* CLOSE <expr> CLOSE
//
//	<expr> ::= <atom> | <decorator>
//
//	<atom> ::= "Any"
//	         | "Bool"
//	         | "Float64"
//	         | "Int"
//	         | "String"
//	         | "Unit"
//
//	<union> ::= OPEN "Union" <expr>* CLOSE
//
//	<variadic> ::= OPEN "Variadic" <expr> CLOSE
//
//	<decorator> := <callable> | <union> | <variadic>
func (p *annotationParser) Parse() (*Callable, error) {
	// <annotation> ::= <callable> EOF
	callable, err := p.parseCallable()
	if err != nil {
		return nil, err
	}
	if !p.check(token.EOF) {
		return nil, p.newError("annotation parser: expected EOF")
	}
	return callable, nil
}

// parseLambdaAnnotation parses an expression.
func (p *annotationParser) parseCallable() (*Callable, error) {
	// <callable> ::= OPEN "Callable" OPEN <expr>* CLOSE <expr> CLOSE

	if !p.match(token.OPEN) {
		return nil, p.newError("annotation parser: expected '('")
	}

	if !p.match(token.ATOM) && p.peek().Value != "Callable" {
		return nil, p.newError("annotation parser: expected 'Callable'")
	}

	if !p.match(token.OPEN) {
		return nil, p.newError("annotation parser: expected '('")
	}

	var callable Callable
	for !p.check(token.CLOSE) {
		kind, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		callable.ParamsTypes = append(callable.ParamsTypes, kind)
	}

	if !p.match(token.CLOSE) {
		return nil, p.newError("annotation parser: expected ')'")
	}

	kind, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	callable.ReturnType = kind

	if !p.match(token.CLOSE) {
		return nil, p.newError("annotation parser: expected ')'")
	}

	return &callable, nil
}

// parseExpr parses a type expression.
func (p *annotationParser) parseExpr() (visitor.Type, error) {
	//	<expr> ::= <atom> | <decorator>
	switch {
	case p.check(token.ATOM):
		return p.parseAtom()
	case p.check(token.OPEN):
		return p.parseDecorator()
	default:
		return nil, p.newError("annotation parser: expected '(' or an atom")
	}
}

// parseAtom parses an atom.
func (p *annotationParser) parseAtom() (visitor.Type, error) {
	//	<atom> ::= "Any"
	//	         | "Bool"
	//	         | "Float64"
	//	         | "Int"
	//	         | "String"
	//	         | "Unit"
	switch p.peek().Value {
	case "Any":
		p.advance()
		return &Any{}, nil
	case "Bool":
		p.advance()
		return &Bool{}, nil
	case "Float64":
		p.advance()
		return &Float64{}, nil
	case "Int":
		p.advance()
		return &Int{}, nil
	case "String":
		p.advance()
		return &String{}, nil
	case "Unit":
		p.advance()
		return &Unit{}, nil
	default:
		return nil, p.newError("annotation parser: unknown type: %s", p.peek().Value)
	}
}

// parseDecorator parses a decorator.
func (p *annotationParser) parseDecorator() (visitor.Type, error) {
	// <decorator> := <callable> | <union> | <variadic>
	tok := p.peekNext()
	switch {
	case tok.TokenType == token.ATOM && tok.Value == "Callable":
		return p.parseCallable()
	case tok.TokenType == token.ATOM && tok.Value == "Union":
		return p.parseUnion()
	case tok.TokenType == token.ATOM && tok.Value == "Variadic":
		return p.parseVariadic()
	default:
		return nil, p.newError("annotation parser: expected 'Callable', 'Union' or 'Variadic'")
	}
}

// parseUnion parses a union.
func (p *annotationParser) parseUnion() (visitor.Type, error) {
	// <union> ::= OPEN "Union" <expr>* CLOSE
	if !p.match(token.OPEN) {
		return nil, p.newError("annotation parser: expected '('")
	}
	if !p.match(token.ATOM) && p.peek().Value != "Union" {
		return nil, p.newError("annotation parser: expected 'Union'")
	}
	union := NewUnion()
	for !p.check(token.CLOSE) {
		expr, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		union.Add(expr)
	}
	if !p.match(token.CLOSE) {
		return nil, p.newError("annotation parser: expected ')'")
	}
	return union, nil
}

// parseVariadic parses a variadic.
func (p *annotationParser) parseVariadic() (visitor.Type, error) {
	// <variadic> ::= OPEN "Variadic" <expr> CLOSE
	if !p.match(token.OPEN) {
		return nil, p.newError("annotation parser: expected '('")
	}
	if !p.match(token.ATOM) && p.peek().Value != "Variadic" {
		return nil, p.newError("annotation parser: expected 'Variadic'")
	}
	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	if !p.match(token.CLOSE) {
		return nil, p.newError("annotation parser: expected ')'")
	}
	return &Variadic{Type: expr}, nil
}

// peek returns the next token without advancing the parser.
func (p *annotationParser) peek() token.Token {
	if p.current < len(p.tokens) {
		return p.tokens[p.current]
	}
	return token.Token{}
}

// peekNext returns the next token without advancing the parser.
func (p *annotationParser) peekNext() token.Token {
	if p.current+1 < len(p.tokens) {
		return p.tokens[p.current+1]
	}
	return token.Token{}
}

// advance advances the parser by one token.
func (p *annotationParser) advance() {
	if p.current < len(p.tokens) {
		p.current++
	}
}

// check checks if the current token matches the given token type.
func (p *annotationParser) check(tt token.TokenType) bool {
	return p.peek().TokenType == tt
}

// match checks if the current token matches the given token type.
func (p *annotationParser) match(tt token.TokenType) bool {
	if p.check(tt) {
		p.advance()
		return true
	}
	return false
}

// newError creates an newError message with the current token.
func (p *annotationParser) newError(format string, v ...any) error {
	tok := p.peek()
	return fmt.Errorf(
		"%s:%d:%d: %s",
		tok.TokenPos.FileName,
		tok.TokenPos.LineNumber,
		tok.TokenPos.LineColumn,
		fmt.Sprintf(format, v...),
	)
}
