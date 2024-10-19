// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseBlock parses a block form into an AST node.
func (p *parser) parseBlock(tok token.Token) (ast.Node, error) {
	// Syntax: ... <expr>* CLOSE
	var exprs []ast.Node
	for p.currentToken().TokenType != token.CLOSE {
		expr, err := p.parseAtomOrForm()
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
	}
	_, _ = p.consumeTokenWithType(token.CLOSE) // can't fail
	if len(exprs) <= 0 {
		return &ast.UnitExpr{Token: tok}, nil
	}
	return &ast.BlockExpr{Token: tok, Exprs: exprs}, nil
}

// parseCond parses the cond special form into an AST node.
func (p *parser) parseCond(tok token.Token) (ast.Node, error) {
	// Syntax: ... (OPEN <predicate> <expr> CLOSE)* (OPEN "else" <expr> CLOSE)? CLOSE

	// 1. (OPEN <predicate> <expr> CLOSE)* (OPEN "else" <expr> CLOSE)?
	var (
		cases    []ast.CondCase
		elseExpr ast.Node
		err      error
	)
	for p.currentToken().TokenType == token.OPEN {
		p.advance()

		if p.currentToken().TokenType == token.ATOM && p.currentToken().Value == "else" {
			p.advance()
			elseExpr, err = p.parseAtomOrForm()
			if err != nil {
				return nil, err
			}
			if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
				return nil, err
			}
			break // no need to parse more cases
		}

		predicate, err := p.parseAtomOrForm()
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

		cases = append(cases, ast.CondCase{Predicate: predicate, Expr: expr})
	}

	// 2. CLOSE
	if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
		return nil, err
	}

	// 3. reduce to `Unit` in case it's all empty
	if len(cases) <= 0 && elseExpr == nil {
		return &ast.UnitExpr{Token: tok}, nil
	}

	// 4. reduce to `else` if there are no cases
	if len(cases) <= 0 && elseExpr != nil {
		return elseExpr, nil
	}

	// 5. add the else branch if missing
	if elseExpr == nil {
		elseExpr = &ast.UnitExpr{Token: tok}
	}

	rv := &ast.CondExpr{Token: tok, Cases: cases, ElseExpr: elseExpr}
	return rv, nil
}

// parseIf parses an if form into an AST node.
func (p *parser) parseIf(tok token.Token) (ast.Node, error) {
	// Syntax: ... <predicate> <then-expr> <else-expr>? CLOSE

	// 1. collect predicate and then
	predicate, err := p.parseAtomOrForm()
	if err != nil {
		return nil, err
	}
	thenExpr, err := p.parseAtomOrForm()
	if err != nil {
		return nil, err
	}

	// 2. add else branch if present and consume final CLOSE
	var elseExpr ast.Node = &ast.UnitExpr{Token: tok}
	if p.currentToken().TokenType != token.CLOSE {
		elseExpr, err = p.parseAtomOrForm()
		if err != nil {
			return nil, err
		}
	}
	_, _ = p.consumeTokenWithType(token.CLOSE) // cannot fail

	// 3. desugar `if` into a `cond`
	cases := []ast.CondCase{{Predicate: predicate, Expr: thenExpr}}
	return &ast.CondExpr{Token: tok, Cases: cases, ElseExpr: elseExpr}, nil
}

// parseReturn parses a return form into an AST node.
func (p *parser) parseReturn(tok token.Token) (ast.Node, error) {
	// Syntax: ... <expr> CLOSE
	expr, err := p.parseAtomOrForm()
	if err != nil {
		return nil, err
	}
	if _, err := p.consumeTokenWithType(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.ReturnStmt{Token: tok, Expr: expr}, nil
}

// parseWhile parses a while form into an AST node.
func (p *parser) parseWhile(tok token.Token) (ast.Node, error) {
	// Syntax: ... <predicate> <expr> CLOSE
	predicate, err := p.parseAtomOrForm()
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
	return &ast.WhileExpr{Token: tok, Predicate: predicate, Expr: expr}, nil
}
