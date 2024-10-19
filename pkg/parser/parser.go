// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"fmt"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

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
func newError(tok token.Token, format string, args ...any) *Error {
	return &Error{Tok: tok, Message: fmt.Sprintf(format, args...)}
}

// parser is a structure that holds the tokens to be parsed and
// the current position in the token list.
type parser struct {
	tokens  []token.Token
	current int
}

// newParser creates a new parser instance with the provided tokens.
func newParser(tokens []token.Token) *parser {
	return &parser{tokens: tokens, current: 0}
}

// Parse processes the tokens and returns a slice of AST nodes.
func (p *parser) Parse() ([]ast.Node, error) {
	var nodes []ast.Node
	for p.currentToken().TokenType != token.EOF {
		node, err := p.parseAtomOrForm()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// advance moves the current position to the next token.
//
// This method is safe to call even if the current token is the last one.
func (p *parser) advance() {
	if p.current < len(p.tokens)-1 {
		p.current++
	}
}

// currentToken returns the current token being processed.
func (p *parser) currentToken() token.Token {
	return p.tokens[p.current]
}

// consumeTokenWithType consumes the current token if it matches the expected type.
func (p *parser) consumeTokenWithType(tt token.TokenType) (token.Token, error) {
	tok := p.currentToken()
	if tok.TokenType != tt {
		err := newError(tok, "expected token %s, found %s", tt, tok.TokenType)
		return token.Token{}, err
	}
	p.advance()
	return tok, nil
}

// parseAtomOrForm determines the type of the current token and delegates to the appropriate parsing function.
func (p *parser) parseAtomOrForm() (ast.Node, error) {
	switch tp := p.currentToken(); tp.TokenType {
	case token.ATOM:
		return p.parseSymbol()
	case token.NUMBER:
		return p.parseNumber()
	case token.STRING:
		return p.parseString()
	case token.OPEN:
		return p.parseForm()
	default:
		return nil, newError(tp, "unexpected token %s", tp.TokenType)
	}
}

// parseForm parses a form token into an AST node.
func (p *parser) parseForm() (ast.Node, error) {
	tok := p.currentToken()
	p.advance()

	if p.currentToken().TokenType == token.CLOSE {
		p.advance()
		rv := &ast.UnitExpr{Token: tok}
		return rv, nil
	}

	form := p.currentToken()
	if form.TokenType == token.ATOM {
		specialForms := map[string]func(token.Token) (ast.Node, error){
			"block":  p.parseBlock,
			"cond":   p.parseCond,
			"if":     p.parseIf,
			"define": p.parseDefine,
			"lambda": p.parseLambda,
			"return": p.parseReturn,
			"set":    p.parseSet,
			"while":  p.parseWhile,
		}
		if parseFunc, found := specialForms[form.Value]; found {
			p.advance() // consume the special form name
			return parseFunc(tok)
		}
	}

	return p.parseCall(tok)
}
