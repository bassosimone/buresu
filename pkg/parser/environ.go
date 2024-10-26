// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseDefine parses a define form into an AST node.
func (p *parser) parseDefine(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "define" <symbol> <expr> CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("define"); err != nil {
		return nil, err
	}
	return p.parseDefineOrSet(tok, func(tok token.Token, symbol string, expr ast.Node) ast.Node {
		return &ast.DefineExpr{Token: tok, Symbol: symbol, Expr: expr}
	})
}

// parseSet parses a set form into an AST node.
func (p *parser) parseSet(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "set!" <symbol> <expr> CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("set!"); err != nil {
		return nil, err
	}
	return p.parseDefineOrSet(tok, func(tok token.Token, symbol string, expr ast.Node) ast.Node {
		return &ast.SetExpr{Token: tok, Symbol: symbol, Expr: expr}
	})
}

func (p *parser) parseDefineOrSet(
	tok token.Token, build func(token.Token, string, ast.Node) ast.Node) (ast.Node, error) {
	// Syntax: ... <symbol> <expr> CLOSE
	symbol, err := p.match(token.ATOM)
	if err != nil {
		return nil, err
	}
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	if _, err := p.match(token.CLOSE); err != nil {
		return nil, err
	}
	rv := build(tok, symbol.Value, expr)
	return rv, nil
}
