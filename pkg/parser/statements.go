// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

type parseFunc func(tok token.Token) (ast.Node, error)

// parseStmtNotAllowed is a wrapper for statement parsing functions
// that ensures we report an error if there's a statement in a context
// in which it is not allowed by the grammar.
func (p *parser) parseStmtNotAllowed(name string, fx parseFunc) parseFunc {
	return func(tok token.Token) (ast.Node, error) {
		if _, err := fx(tok); err != nil {
			return nil, err
		}
		return nil, newError(tok, "%s statement not allowed in this context", name)
	}
}

// parseReturn parses a return form into an AST node.
func (p *parser) parseReturn(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "return!" <expr> CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("return!"); err != nil {
		return nil, err
	}

	// 1. reject return outside of any lambda
	if p.lambdadepth <= 0 {
		return nil, newError(tok, "return! outside of lambda")
	}

	// 2. <expr>
	expr, err := p.parseWithFlags(0)
	if err != nil {
		return nil, err
	}

	// 3. CLOSE
	if _, err := p.match(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.ReturnStmt{Token: tok, Expr: expr}, nil
}

// parseInclude parses an include form into an AST node.
func (p *parser) parseInclude(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "include!" STRING CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("include!"); err != nil {
		return nil, err
	}

	// 2. STRING
	node, err := p.parseString()
	if err != nil {
		return nil, err
	}
	filepath := node.(*ast.StringLiteral).Value // guaranteed to be a string

	// 3. CLOSE
	if _, err := p.match(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.IncludeStmt{Token: tok, FilePath: filepath}, nil
}
