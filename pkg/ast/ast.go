// SPDX-License-Identifier: GPL-3.0-or-later

// Package ast contains the abstract syntax tree (AST) of the language.
//
// The parser package generates AST nodes from a sequence of tokens.
//
// Packages like the intepreter will then traverse and manipulate the AST.
package ast

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bassosimone/buresu/pkg/token"
)

// Node represents a node in the AST.
type Node interface {
	// String converts the node back to lisp source code.
	String() string

	// Clone returns a deep copy of the node.
	Clone() Node
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

// Clone creates a deep copy of the BlockExpr node.
func (blk *BlockExpr) Clone() Node {
	exprs := make([]Node, len(blk.Exprs))
	for idx, expr := range blk.Exprs {
		exprs[idx] = expr.Clone()
	}
	return &BlockExpr{
		Token: blk.Token.Clone(),
		Exprs: exprs,
	}
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

// Clone creates a deep copy of the CallExpr node.
func (call *CallExpr) Clone() Node {
	args := make([]Node, len(call.Args))
	for idx, arg := range call.Args {
		args[idx] = arg.Clone()
	}
	return &CallExpr{
		Token:    call.Token.Clone(),
		Callable: call.Callable.Clone(),
		Args:     args,
	}
}

// CondCase represents a single case in a conditional expression.
type CondCase struct {
	Predicate Node
	Expr      Node
}

// Clone creates a deep copy of the CondCase node.
func (cc *CondCase) Clone() CondCase {
	return CondCase{
		Predicate: cc.Predicate.Clone(),
		Expr:      cc.Expr.Clone(),
	}
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

// Clone creates a deep copy of the CondExpr node.
func (cond *CondExpr) Clone() Node {
	cases := make([]CondCase, len(cond.Cases))
	for idx, condCase := range cond.Cases {
		cases[idx] = condCase.Clone()
	}
	return &CondExpr{
		Token:    cond.Token.Clone(),
		Cases:    cases,
		ElseExpr: cond.ElseExpr.Clone(),
	}
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

// Clone creates a deep copy of the DefineExpr node.
func (def *DefineExpr) Clone() Node {
	return &DefineExpr{
		Token:  def.Token.Clone(),
		Symbol: def.Symbol,
		Expr:   def.Expr.Clone(),
	}
}

// FalseLiteral represents a boolean false value.
type FalseLiteral struct {
	Token token.Token
}

// String converts the FalseLiteral node back to lisp source code.
func (fal *FalseLiteral) String() string {
	return "false"
}

// Clone creates a deep copy of the FalseLiteral node.
func (fal *FalseLiteral) Clone() Node {
	return &FalseLiteral{
		Token: fal.Token.Clone(),
	}
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

// Clone creates a deep copy of the FloatLiteral node.
func (fltLit *FloatLiteral) Clone() Node {
	return &FloatLiteral{
		Token: fltLit.Token.Clone(),
		Value: fltLit.Value,
	}
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

// Clone creates a deep copy of the IntLiteral node.
func (intLit *IntLiteral) Clone() Node {
	return &IntLiteral{
		Token: intLit.Token.Clone(),
		Value: intLit.Value,
	}
}

// LambdaExpr represents an inline function definition with docs.
type LambdaExpr struct {
	Token  token.Token
	Params []string
	Docs   StringLiteral
	Expr   Node
}

// String converts the LambdaExpr node back to lisp source code.
func (lam *LambdaExpr) String() string {
	params := make([]string, len(lam.Params))
	for idx, param := range lam.Params {
		params[idx] = param
	}
	docsBytes, _ := json.Marshal(lam.Docs.Value)
	docs := fmt.Sprintf("%s", string(docsBytes))
	return fmt.Sprintf("(lambda (%s) %s %s)", strings.Join(params, ", "), docs, lam.Expr.String())
}

// Clone creates a deep copy of the LambdaExpr node.
func (lam *LambdaExpr) Clone() Node {
	params := make([]string, len(lam.Params))
	for idx, param := range lam.Params {
		params[idx] = param
	}
	return &LambdaExpr{
		Token:  lam.Token.Clone(),
		Params: params,
		Docs:   *lam.Docs.Clone().(*StringLiteral),
		Expr:   lam.Expr.Clone(),
	}
}

// ReturnStmt represents a return statement to interrupt
// the current function and return a value.
type ReturnStmt struct {
	Token token.Token
	Expr  Node
}

// String converts the ReturnStmt node back to lisp source code.
func (ret *ReturnStmt) String() string {
	return fmt.Sprintf("(return %s)", ret.Expr.String())
}

// Clone creates a deep copy of the ReturnStmt node.
func (ret *ReturnStmt) Clone() Node {
	return &ReturnStmt{
		Token: ret.Token.Clone(),
		Expr:  ret.Expr.Clone(),
	}
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
	return fmt.Sprintf("(set %s %s)", set.Symbol, set.Expr.String())
}

// Clone creates a deep copy of the SetExpr node.
func (set *SetExpr) Clone() Node {
	return &SetExpr{
		Token:  set.Token.Clone(),
		Symbol: set.Symbol,
		Expr:   set.Expr.Clone(),
	}
}

// StringLiteral represents a string value containing ASCII text.
type StringLiteral struct {
	Token token.Token
	Value string
}

// String converts the StringLiteral node back to lisp source code.
func (strLit *StringLiteral) String() string {
	valueBytes, _ := json.Marshal(strLit.Value)
	return string(valueBytes)
}

// Clone creates a deep copy of the StringLiteral node.
func (strLit *StringLiteral) Clone() Node {
	return &StringLiteral{
		Token: strLit.Token.Clone(),
		Value: strLit.Value,
	}
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

// Clone creates a deep copy of the SymbolName node.
func (sym *SymbolName) Clone() Node {
	return &SymbolName{
		Token: sym.Token.Clone(),
		Value: sym.Value,
	}
}

// TrueLiteral represents a boolean true value.
type TrueLiteral struct {
	Token token.Token
}

// String converts the TrueLiteral node back to lisp source code.
func (tru *TrueLiteral) String() string {
	return "true"
}

// Clone creates a deep copy of the TrueLiteral node.
func (tru *TrueLiteral) Clone() Node {
	return &TrueLiteral{
		Token: tru.Token.Clone(),
	}
}

// UnitExpr represents an expression returning the value of the Unit type.
type UnitExpr struct {
	Token token.Token
}

// String converts the UnitExpr node back to lisp source code.
func (unit *UnitExpr) String() string {
	return "()"
}

// Clone creates a deep copy of the UnitExpr node.
func (unit *UnitExpr) Clone() Node {
	return &UnitExpr{
		Token: unit.Token.Clone(),
	}
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

// Clone creates a deep copy of the WhileExpr node.
func (whl *WhileExpr) Clone() Node {
	return &WhileExpr{
		Token:     whl.Token.Clone(),
		Predicate: whl.Predicate.Clone(),
		Expr:      whl.Expr.Clone(),
	}
}
