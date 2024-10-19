// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"strings"

	"github.com/bassosimone/buresu/pkg/ast"
)

// parseSymbol parses an atom token into an AST node.
func (p *parser) parseSymbol() (ast.Node, error) {
	// Syntax: ATOM
	tok := p.currentToken()
	p.advance()
	var rv ast.Node
	switch {
	case tok.Value == "false":
		rv = &ast.FalseLiteral{Token: tok}
	case tok.Value == "true":
		rv = &ast.TrueLiteral{Token: tok}
	default:
		rv = &ast.SymbolName{Token: tok, Value: tok.Value}
	}
	return rv, nil
}

// parseNumber parses a number token into an AST node.
func (p *parser) parseNumber() (ast.Node, error) {
	// Syntax: NUMBER
	tok := p.currentToken()
	p.advance()
	var rv ast.Node
	if strings.Contains(tok.Value, ".") {
		rv = &ast.FloatLiteral{Token: tok, Value: tok.Value}
	} else {
		rv = &ast.IntLiteral{Token: tok, Value: tok.Value}
	}
	return rv, nil
}

// parseString parses a string token into an AST node.
func (p *parser) parseString() (ast.Node, error) {
	// Syntax: STRING
	tok := p.currentToken()
	p.advance()
	rv := &ast.StringLiteral{Token: tok, Value: tok.Value}
	return rv, nil
}
