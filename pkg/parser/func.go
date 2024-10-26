// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func (p *parser) parseCall(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN <callable> <expr> ... CLOSE
	var args []ast.Node

	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}

	// <callable>
	callable, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// <expr> ... CLOSE
	for p.peek().TokenType != token.CLOSE {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		args = append(args, expr)
	}
	_, _ = p.match(token.CLOSE) // cannot fail

	rv := &ast.CallExpr{
		Token:    tok,
		Callable: callable,
		Args:     args,
	}
	return rv, nil
}

func (p *parser) parseLambda(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "lambda" OPEN <param>* CLOSE [STRING] <expr> CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("lambda"); err != nil {
		return nil, err
	}

	// 1. parse OPEN <param>* CLOSE
	var (
		params  []string
		uniqnam = make(map[string]struct{})
	)
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	for p.peek().TokenType != token.CLOSE {
		paramName, err := p.parseSymbol()
		if err != nil {
			return nil, err
		}
		name := paramName.(*ast.SymbolName).Value
		if _, ok := uniqnam[name]; ok {
			return nil, newError(tok, "lambda parameter %q is duplicated", name)
		}
		uniqnam[name] = struct{}{}
		params = append(params, name)
	}
	_, _ = p.match(token.CLOSE) // cannot fail

	// 2. [STRING] including type annotations
	var docs string
	if p.peek().TokenType == token.STRING {
		docs = p.peek().Value
		p.advance()
	}

	// 3. <expr> CLOSE
	//
	// Track the depth inside lambdas so we know when it
	// is legal to accept a return statement.
	p.lambdadepth++
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	p.lambdadepth--
	if _, err := p.match(token.CLOSE); err != nil {
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
