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
	for p.currentToken().TokenType != token.EOF {
		node, err := p.parseAtomOrExpression()
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
		if tok.TokenType == token.EOF {
			err = ErrIncompleteInput{err}
		}
		return token.Token{}, err
	}
	p.advance()
	return tok, nil
}

// parseAtomOrExpression determines the type of the current token and delegates
// to the appropriate parsing function for parsing an expression or an atom, which
// we consider as a self-evaluating expression for the sake of simplicity.
func (p *parser) parseAtomOrExpression() (ast.Node, error) {
	return p.parseWithFlags(0)
}

func (p *parser) parseAtomOrExpressionOrStatement() (ast.Node, error) {
	return p.parseWithFlags(allowStmts)
}

const (
	// allowStmts allows parseWithFlags to parse statements
	allowStmts = 1 << iota
)

func (p *parser) parseWithFlags(flags int) (ast.Node, error) {
	switch tp := p.currentToken(); tp.TokenType {
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

// parseForm parses a form honouring the given flags.
func (p *parser) parseForm(flags int) (ast.Node, error) {
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
			"block":   p.parseBlock,
			"cond":    p.parseCond,
			"if":      p.parseIf,
			"define":  p.parseDefine,
			"lambda":  p.parseLambda,
			"quote":   p.parseQuote,
			"return!": p.parseStmtNotAllowed(p.parseReturn),
			"set!":    p.parseSet,
			"while":   p.parseWhile,
		}
		if flags&allowStmts != 0 {
			specialForms["return!"] = p.parseReturn
		}
		if parseFunc, found := specialForms[form.Value]; found {
			p.advance() // consume the special form name
			return parseFunc(tok)
		}
	}

	return p.parseCall(tok)
}
