// SPDX-License-Identifier: GPL-3.0-or-later

package parser

import (
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// parseBlock parses a block form into an AST node.
func (p *parser) parseBlock(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "block" <expr>* CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("block"); err != nil {
		return nil, err
	}

	var (
		exprs   []ast.Node
		seenret bool
	)
	for p.peek().TokenType != token.CLOSE {
		// make sure nothing follows a return statement
		if seenret {
			return nil, newError(tok, "unreachable code")
		}

		expr, err := p.parseWithFlags(allowReturn)
		if err != nil {
			return nil, err
		}

		// remember if we've seen a return statement
		if _, ok := expr.(*ast.ReturnStmt); ok {
			seenret = true
		}

		exprs = append(exprs, expr)
	}

	_, _ = p.match(token.CLOSE) // can't fail

	if len(exprs) <= 0 {
		return &ast.UnitExpr{Token: tok}, nil
	}

	return &ast.BlockExpr{Token: tok, Exprs: exprs}, nil
}

// parseCond parses the cond special form into an AST node.
func (p *parser) parseCond(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "cond" (OPEN <predicate> <expr> CLOSE)* (OPEN "else" <expr> CLOSE)? CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("cond"); err != nil {
		return nil, err
	}

	// 1. (OPEN <predicate> <expr> CLOSE)* (OPEN "else" <expr> CLOSE)?
	var (
		cases    []ast.CondCase
		elseExpr ast.Node
		err      error
	)
	for p.peek().TokenType == token.OPEN {
		p.advance()

		if p.peek().TokenType == token.ATOM && p.peek().Value == "else" {
			p.advance()
			elseExpr, err = p.parseWithFlags(0)
			if err != nil {
				return nil, err
			}
			if _, err := p.match(token.CLOSE); err != nil {
				return nil, err
			}
			break // no need to parse more cases
		}

		predicate, err := p.parseWithFlags(0)
		if err != nil {
			return nil, err
		}

		expr, err := p.parseWithFlags(0)
		if err != nil {
			return nil, err
		}

		if _, err := p.match(token.CLOSE); err != nil {
			return nil, err
		}

		cases = append(cases, ast.CondCase{Predicate: predicate, Expr: expr})
	}

	// 2. CLOSE
	if _, err := p.match(token.CLOSE); err != nil {
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
	// Syntax: OPEN "if" <predicate> <then-expr> <else-expr>? CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("if"); err != nil {
		return nil, err
	}

	// 1. collect predicate and then
	predicate, err := p.parseWithFlags(0)
	if err != nil {
		return nil, err
	}
	thenExpr, err := p.parseWithFlags(0)
	if err != nil {
		return nil, err
	}

	// 2. add else branch if present and consume final CLOSE
	var elseExpr ast.Node = &ast.UnitExpr{Token: tok}
	if p.peek().TokenType != token.CLOSE {
		elseExpr, err = p.parseWithFlags(0)
		if err != nil {
			return nil, err
		}
	}
	_, _ = p.match(token.CLOSE) // cannot fail

	// 3. desugar `if` into a `cond`
	cases := []ast.CondCase{{Predicate: predicate, Expr: thenExpr}}
	return &ast.CondExpr{Token: tok, Cases: cases, ElseExpr: elseExpr}, nil
}

// parseWhile parses a while form into an AST node.
func (p *parser) parseWhile(tok token.Token) (ast.Node, error) {
	// Syntax: OPEN "while" <predicate> <expr> CLOSE
	if _, err := p.match(token.OPEN); err != nil {
		return nil, err
	}
	if _, err := p.matchAtomWithName("while"); err != nil {
		return nil, err
	}

	predicate, err := p.parseWithFlags(0)
	if err != nil {
		return nil, err
	}
	expr, err := p.parseWithFlags(0)
	if err != nil {
		return nil, err
	}
	if _, err := p.match(token.CLOSE); err != nil {
		return nil, err
	}
	return &ast.WhileExpr{Token: tok, Predicate: predicate, Expr: expr}, nil
}
