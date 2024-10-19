// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseDefine parses a define form into an AST node.
func (p *parser) parseDefine(tok token.Token) (ast.Node, error) {
	// Syntax: ... <symbol> <expr> CLOSE
	return p.parseDefineOrSet(tok, func(tok token.Token, symbol string, expr ast.Node) ast.Node {
		return &ast.DefineExpr{Token: tok, Symbol: symbol, Expr: expr}
	})
}

// parseSet parses a set form into an AST node.
func (p *parser) parseSet(tok token.Token) (ast.Node, error) {
	// syntax: ... <symbol> <expr> CLOSE
	return p.parseDefineOrSet(tok, func(tok token.Token, symbol string, expr ast.Node) ast.Node {
		return &ast.SetExpr{Token: tok, Symbol: symbol, Expr: expr}
	})
}

func (p *parser) parseDefineOrSet(
	tok token.Token, build func(token.Token, string, ast.Node) ast.Node) (ast.Node, error) {
	// syntax: ... <symbol> <expr> CLOSE
	symbol, err := p.consumeTokenWithType(token.ATOM)
	if err != nil {
		return nil, err
	}
	expr, err := p.parseAtomOrForm()
	if err != nil {
		return nil, err
	}
	if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
		return nil, err
	}
	rv := build(tok, symbol.Value, expr)
	return rv, nil
}
