// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func (p *parser) parseCall(tok token.Token) (ast.Node, error) {
	// syntax: ... <callable> <expr> [COMMA] ... CLOSE
	var args []ast.Node

	// <callable>
	callable, err := p.parseAtomOrForm()
	if err != nil {
		return nil, err
	}

	// <expr> [COMMA] ... CLOSE
	for p.currentToken().TokenType != token.CLOSE {
		expr, err := p.parseAtomOrForm()
		if err != nil {
			return nil, err
		}
		args = append(args, expr)
	}
	_, _ = p.consumeTokenWithType(token.CLOSE) // cannot fail

	rv := &ast.CallExpr{
		Token:    tok,
		Callable: callable,
		Args:     args,
	}
	return rv, nil
}

func (p *parser) parseLambda(tok token.Token) (ast.Node, error) {
	// Syntax: ... OPEN <param>* CLOSE [STRING] <expr>
	var err error

	// 1. parse OPEN <param>* CLOSE
	var params []string
	if _, err := p.consumeTokenWithType(token.OPEN); err != nil {
		return nil, err
	}
	for p.currentToken().TokenType != token.CLOSE {
		paramName, err := p.parseAtomOrForm()
		if err != nil {
			return nil, err
		}
		if _, ok := paramName.(*ast.SymbolName); !ok {
			return nil, newError(tok, "lambda parameter name must be a symbol")
		}
		params = append(params, paramName.(*ast.SymbolName).Value)
	}
	_, _ = p.consumeTokenWithType(token.CLOSE) // cannot fail

	// 2. [STRING]
	var docs string
	if p.currentToken().TokenType == token.STRING {
		docs = p.currentToken().Value
		p.advance()
	}

	// 3. <expr> CLOSE
	expr, err := p.parseAtomOrForm()
	if err != nil {
		return nil, err
	}
	if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
		return nil, err
	}

	rv := &ast.LambdaExpr{
		Token:  tok,
		Params: params,
		Docs:   docs,
		Expr:   expr,
	}
	return rv, nil
}
