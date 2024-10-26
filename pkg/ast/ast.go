// SPDX-License-Identifier: GPL-3.0-or-later

// Package ast contains the abstract syntax tree (AST) of the language.
//
// The parser package generates AST nodes from a sequence of tokens.
//
// Packages like the intepreter will then traverse and manipulate the AST.
package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bassosimone/buresu/pkg/token"
)

// Node represents a node in the AST.
type Node interface {
	// String converts the node back to lisp source code.
	String() string
}

// BlockExpr represents a block of expressions executed sequentially.
type BlockExpr struct {
	Token token.Token
	Exprs []Node
}

// String converts the BlockExpr node back to lisp source code.
func (blk *BlockExpr) String() string {
	exprs := make([]string, len(blk.Exprs))
	for idx, expr := range blk.Exprs {
		exprs[idx] = expr.String()
	}
	return fmt.Sprintf("(block %s)", strings.Join(exprs, " "))
}

// CallExpr represents a call to a given callable with a list
// of arguments and a return type.
type CallExpr struct {
	Token    token.Token
	Callable Node
	Args     []Node
}

// String converts the CallExpr node back to lisp source code.
func (call *CallExpr) String() string {
	args := make([]string, len(call.Args))
	for idx, arg := range call.Args {
		args[idx] = arg.String()
	}
	return fmt.Sprintf(
		"(%s %s)",
		call.Callable.String(),
		strings.Join(args, " "),
	)
}

// CondCase represents a single case in a conditional expression.
type CondCase struct {
	Predicate Node
	Expr      Node
}

// CondExpr represents a conditional expression with multiple branches.
type CondExpr struct {
	Token    token.Token
	Cases    []CondCase
	ElseExpr Node
}

// String converts the CondExpr node back to lisp source code.
func (cond *CondExpr) String() string {
	cases := make([]string, len(cond.Cases))
	for idx, condCase := range cond.Cases {
		cases[idx] = fmt.Sprintf("(%s %s)", condCase.Predicate.String(), condCase.Expr.String())
	}
	elseExpr := fmt.Sprintf(" (else %s)", cond.ElseExpr.String())
	return fmt.Sprintf("(cond %s%s)", strings.Join(cases, " "), elseExpr)
}

// DefineExpr saves a value in a variable within the current scope.
type DefineExpr struct {
	Token  token.Token
	Symbol string
	Expr   Node
}

// String converts the DefineExpr node back to lisp source code.
func (def *DefineExpr) String() string {
	return fmt.Sprintf("(define %s %s)", def.Symbol, def.Expr.String())
}

// FalseLiteral represents a boolean false value.
type FalseLiteral struct {
	Token token.Token
}

// String converts the FalseLiteral node back to lisp source code.
func (fal *FalseLiteral) String() string {
	return "false"
}

// FloatLiteral represents a floating-point value.
type FloatLiteral struct {
	Token token.Token
	Value string
}

// String converts the FloatLiteral node back to lisp source code.
func (fltLit *FloatLiteral) String() string {
	return fltLit.Value
}

// IncludeStmt represents an include expression with a file path.
type IncludeStmt struct {
	Token    token.Token
	FilePath string
}

// String converts the IncludeExpr node back to lisp source code.
func (inc *IncludeStmt) String() string {
	return fmt.Sprintf("(include %s)", jsonMarshalWithoutEscaping(inc.FilePath))
}

// IntLiteral represents an integer value.
type IntLiteral struct {
	Token token.Token
	Value string
}

// String converts the IntLiteral node back to lisp source code.
func (intLit *IntLiteral) String() string {
	return intLit.Value
}

// LambdaExpr represents an inline function definition with docs.
type LambdaExpr struct {
	Token  token.Token
	Params []string
	Docs   string
	Expr   Node
}

// String converts the LambdaExpr node back to lisp source code.
func (lam *LambdaExpr) String() string {
	params := make([]string, len(lam.Params))
	for idx, param := range lam.Params {
		params[idx] = param
	}
	docs := jsonMarshalWithoutEscaping(lam.Docs)
	return fmt.Sprintf("(lambda (%s) %s %s)", strings.Join(params, " "), docs, lam.Expr.String())
}

func jsonMarshalWithoutEscaping(v string) string {
	var buff bytes.Buffer
	enc := json.NewEncoder(&buff)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(v)
	return strings.TrimSpace(buff.String())
}

// QuoteExpr represents a quoted expression.
type QuoteExpr struct {
	Token token.Token
	Expr  Node
}

// String converts the QuoteExpr node back to lisp source code.
func (quote *QuoteExpr) String() string {
	return fmt.Sprintf("(quote %s)", quote.Expr.String())
}

// ReturnStmt represents a return statement to interrupt
// the current function and return a value.
type ReturnStmt struct {
	Token token.Token
	Expr  Node
}

// String converts the ReturnStmt node back to lisp source code.
func (ret *ReturnStmt) String() string {
	return fmt.Sprintf("(return! %s)", ret.Expr.String())
}

// SetExpr sets a value in a previously created variable, which may
// be in the current scope or in a parent scope.
type SetExpr struct {
	Token  token.Token
	Symbol string
	Expr   Node
}

// String converts the SetExpr node back to lisp source code.
func (set *SetExpr) String() string {
	return fmt.Sprintf("(set! %s %s)", set.Symbol, set.Expr.String())
}

// StringLiteral represents a string value containing ASCII text.
type StringLiteral struct {
	Token token.Token
	Value string
}

// String converts the StringLiteral node back to lisp source code.
func (strLit *StringLiteral) String() string {
	return jsonMarshalWithoutEscaping(strLit.Value)
}

// SymbolName represents a symbol that may represent a variable or a function.
type SymbolName struct {
	Token token.Token
	Value string
}

// String converts the SymbolName node back to lisp source code.
func (sym *SymbolName) String() string {
	return sym.Value
}

// TrueLiteral represents a boolean true value.
type TrueLiteral struct {
	Token token.Token
}

// String converts the TrueLiteral node back to lisp source code.
func (tru *TrueLiteral) String() string {
	return "true"
}

// UnitExpr represents an expression returning the value of the Unit type.
type UnitExpr struct {
	Token token.Token
}

// String converts the UnitExpr node back to lisp source code.
func (unit *UnitExpr) String() string {
	return "()"
}

// WhileExpr represents a while loop to execute a block of code repeatedly while the condition is true.
type WhileExpr struct {
	Token     token.Token
	Predicate Node
	Expr      Node
}

// String converts the WhileExpr node back to lisp source code.
func (whl *WhileExpr) String() string {
	return fmt.Sprintf("(while %s %s)", whl.Predicate.String(), whl.Expr.String())
}
