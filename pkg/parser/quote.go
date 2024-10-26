// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseQuote parses the quote special form into an AST node.
func (p *parser) parseQuote(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "quote" <expr> CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("quote"); err != nil {
		return nil, err
	}

	expr, err := p.parseWithFlags(0)
	if err != nil {
		return nil, err
	}
	if _, err := p.match(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.QuoteExpr{Token: tok, Expr: expr}, nil
}
