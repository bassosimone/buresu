// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"strings"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseSymbol parses an atom token into an AST node.
func (p *parser) parseSymbol() (ast.Node, error) {
	// Syntax: ATOM
	tok, err := p.match(token.ATOM)
	if err != nil {
		return nil, err
	}
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
	tok, err := p.match(token.NUMBER)
	if err != nil {
		return nil, err
	}
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
	tok, err := p.match(token.STRING)
	if err != nil {
		return nil, err
	}
	rv := &ast.StringLiteral{Token: tok, Value: tok.Value}
	return rv, nil
}
