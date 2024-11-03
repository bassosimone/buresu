// SPDX-License-Identifier: GPL-3.0-or-later

package ast

import (
	"testing"

	"github.com/bassosimone/buresu/pkg/token"
)

func TestBlockExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "block"}
	trueToken := token.Token{TokenType: token.ATOM, Value: "true"}
	falseToken := token.Token{TokenType: token.ATOM, Value: "false"}
	trueLiteral := &TrueLiteral{Token: trueToken}
	falseLiteral := &FalseLiteral{Token: falseToken}
	expr := &BlockExpr{Token: tok, Exprs: []Node{trueLiteral, falseLiteral}}
	expected := "(block true false)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestCallExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "call"}
	argToken := token.Token{TokenType: token.NUMBER, Value: "42"}
	callable := &SymbolName{Token: tok, Value: "myFunction"}
	arg := &IntLiteral{Token: argToken, Value: "42"}
	expr := &CallExpr{Token: tok, Callable: callable, Args: []Node{arg}}
	expected := "(myFunction 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestCondExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "cond"}
	trueToken := token.Token{TokenType: token.ATOM, Value: "true"}
	falseToken := token.Token{TokenType: token.ATOM, Value: "false"}
	trueLiteral := &TrueLiteral{Token: trueToken}
	falseLiteral := &FalseLiteral{Token: falseToken}
	trueCase := CondCase{Predicate: trueLiteral, Expr: &StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"It's true!\""}, Value: "It's true!"}}
	falseCase := CondCase{Predicate: falseLiteral, Expr: &StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"It's false!\""}, Value: "It's false!"}}
	expr := &CondExpr{Token: tok, Cases: []CondCase{trueCase, falseCase}, ElseExpr: &StringLiteral{Token: token.Token{TokenType: token.STRING, Value: "\"Neither true nor false!\""}, Value: "Neither true nor false!"}}
	expected := "(cond (true \"It's true!\") (false \"It's false!\") (else \"Neither true nor false!\"))"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestDeclareExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "declare"}
	expr := &DeclareExpr{Token: tok, Symbol: "x", Expr: &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}}
	expected := "(declare x 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestDefineExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "define"}
	expr := &DefineExpr{Token: tok, Symbol: "x", Expr: &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}}
	expected := "(define x 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestEllipsisLiteral(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "..."}
	expr := &EllipsisLiteral{Token: tok}
	expected := "..."
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestFloatLiteral(t *testing.T) {
	tok := token.Token{TokenType: token.NUMBER, Value: "3.14"}
	expr := &FloatLiteral{Token: tok, Value: "3.14"}
	expected := "3.14"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestIncludeStmt(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "include!"}
	expr := &IncludeStmt{Token: tok, FilePath: "/path/to/file"}
	expected := "(include! \"/path/to/file\")"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestLambdaExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "lambda"}
	param := "x"
	body := &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}
	expr := &LambdaExpr{Token: tok, Params: []string{param}, Docs: "This is a lambda function", Expr: body}
	expected := "(lambda (x) \"This is a lambda function\" 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestQuoteExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "quote"}
	expr := &QuoteExpr{Token: tok, Expr: &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}}
	expected := "(quote 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestReturnStmt(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "return!"}
	expr := &ReturnStmt{Token: tok, Expr: &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}}
	expected := "(return! 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestSetExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "set!"}
	expr := &SetExpr{Token: tok, Symbol: "x", Expr: &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}}
	expected := "(set! x 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestUnitExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "()"}
	expr := &UnitExpr{Token: tok}
	expected := "()"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}

func TestWhileExpr(t *testing.T) {
	tok := token.Token{TokenType: token.ATOM, Value: "while"}
	predicate := &TrueLiteral{Token: token.Token{TokenType: token.ATOM, Value: "true"}}
	body := &IntLiteral{Token: token.Token{TokenType: token.NUMBER, Value: "42"}, Value: "42"}
	expr := &WhileExpr{Token: tok, Predicate: predicate, Expr: body}
	expected := "(while true 42)"
	t.Run("serialization", func(t *testing.T) {
		if expr.String() != expected {
			t.Errorf("expected %s, got %s", expected, expr.String())
		}
	})
}
