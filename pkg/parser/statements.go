// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

type parseFunc func(tok token.Token) (ast.Node, error)

// parseStmtNowAllowed is a wrapper for statement parsing functions
// that ensures we report an error if there's a statement in a context
// in which it is not allowed by the grammar.
func (p *parser) parseStmtNotAllowed(fx parseFunc) parseFunc {
	return func(tok token.Token) (ast.Node, error) {
		if _, err := fx(tok); err != nil {
			return nil, err
		}
		return nil, newError(tok, "statement not allowed in this context")
	}
}

// parseReturn parses a return form into an AST node.
func (p *parser) parseReturn(tok token.Token) (ast.Node, error) {
	// Syntax: ... <expr> CLOSE

	// 1. reject return outside of any lambda
	if p.lambdadepth <= 0 {
		return nil, newError(tok, "return! outside of lambda")
	}

	// 2. <expr>
	expr, err := p.parseAtomOrExpression()
	if err != nil {
		return nil, err
	}

	// 3. CLOSE
	if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.ReturnStmt{Token: tok, Expr: expr}, nil
}
