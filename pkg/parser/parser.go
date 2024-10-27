// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parser is a structure that holds the tokens to be parsed and
// the current position in the token list.
type parser struct {
	// current is the current position in the tokens list.
	current int

	// lambdadepth is the current depth of lambda expressions: if zero
	// we're at top-level, if one we're inside a lambda, if two we're
	// inside a lambda inside a lambda, and so on.
	lambdadepth int

	// tokens contains the tokens to be parsed.
	tokens []token.Token
}

// newParser creates a new parser instance with the provided tokens.
func newParser(tokens []token.Token) *parser {
	return &parser{tokens: tokens, current: 0, lambdadepth: 0}
}

// Parse processes the tokens and returns a slice of AST nodes.
func (p *parser) Parse() ([]ast.Node, error) {
	var nodes []ast.Node
	for p.peek().TokenType != token.EOF {
		node, err := p.parseWithFlags(allowInclude) // only at top-level
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

// peek returns the current token being processed.
func (p *parser) peek() token.Token {
	if p.current < len(p.tokens) {
		return p.tokens[p.current]
	}
	return token.Token{}
}

// peekNext returns the next token to be processed.
func (p *parser) peekNext() token.Token {
	if p.current+1 < len(p.tokens) {
		return p.tokens[p.current+1]
	}
	return token.Token{}
}

// match consumes the current token if it matches the expected type.
func (p *parser) match(tt token.TokenType) (token.Token, error) {
	tok := p.peek()
	if tok.TokenType != tt {
		err := newError(tok, "expected token %s, found %s", tt, tok.TokenType)
		if tok.TokenType == token.EOF {
			err = ErrIncompleteInput{err}
		}
		return token.Token{}, err
	}
	p.advance()
	return tok, nil
}

// check returns true if the current token matches the expected type.
func (p *parser) check(tt token.TokenType) bool {
	return p.peek().TokenType == tt
}

// matchAtomWithName consumes the current token if it matches the expected type and name.
func (p *parser) matchAtomWithName(name string) (token.Token, error) {
	tok := p.peek()
	if tok.TokenType != token.ATOM || tok.Value != name {
		err := newError(tok, "expected atom with name %s, found %s", name, tok.Value)
		if tok.TokenType == token.EOF {
			err = ErrIncompleteInput{err}
		}
		return token.Token{}, err
	}
	p.advance()
	return tok, nil
}

const (
	// allowReturn allows parseWithFlags to parse statements
	allowReturn = 1 << iota

	// allowInclude allows parseWithFlags to parse include
	allowInclude
)

// parseWithFlags parses atoms, numbers, strings, expressions, and
// statements. While some constructs are always allowed, some of them
// are only allowed in specific contexts, like return! which is only
// allowed inside a block expression, or include, which is only allowed
// at the toplevel. The flags parameter controls when it is legal to
// accept these context-dependent constructs.
func (p *parser) parseWithFlags(flags int) (ast.Node, error) {
	switch tp := p.peek(); tp.TokenType {
	case token.ATOM:
		return p.parseSymbol()
	case token.NUMBER:
		return p.parseNumber()
	case token.STRING:
		return p.parseString()
	case token.OPEN:
		return p.parseForm(flags)
	default:
		err := newError(tp, "unexpected token %s", tp.TokenType)
		if tp.TokenType == token.EOF {
			err = ErrIncompleteInput{err}
		}
		return nil, err
	}
}

// parseForm parses a form honoring the given flags.
func (p *parser) parseForm(flags int) (ast.Node, error) {
	tok := p.peek()

	if p.peekNext().TokenType == token.CLOSE {
		p.advance() // consume OPEN
		p.advance() // consume CLOSE
		return &ast.UnitExpr{Token: tok}, nil
	}

	form := p.peekNext()
	if form.TokenType == token.ATOM {
		specialForms := map[string]func(token.Token) (ast.Node, error){
			"block":    p.parseBlock,
			"cond":     p.parseCond,
			"define":   p.parseDefine,
			"if":       p.parseIf,
			"include!": p.parseStmtNotAllowed("include!", p.parseInclude),
			"lambda":   p.parseLambda,
			"quote":    p.parseQuote,
			"return!":  p.parseStmtNotAllowed("return!", p.parseReturn),
			"set!":     p.parseSet,
			"while":    p.parseWhile,
		}
		if flags&allowInclude != 0 {
			specialForms["include!"] = p.parseInclude
		}
		if flags&allowReturn != 0 {
			specialForms["return!"] = p.parseReturn
		}
		if parseFunc, found := specialForms[form.Value]; found {
			return parseFunc(tok)
		}
	}

	return p.parseCall(tok)
}
