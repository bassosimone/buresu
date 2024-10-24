// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseQuote parses the quote special form into an AST node.
func (p *parser) parseQuote(tok token.Token) (ast.Node, error) {
	// Syntax: ... <expr> CLOSE
	expr, err := p.parseAtomOrExpression()
	if err != nil {
		return nil, err
	}
	if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.QuoteExpr{Token: tok, Expr: expr}, nil
}
