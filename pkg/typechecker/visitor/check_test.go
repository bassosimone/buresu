// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"testing"

	"github.com/bassosimone/buresu/pkg/ast"
)

func TestCheck(t *testing.T) {
	ctx := context.Background()
	env := &mockEnvironment{}

	tests := []struct {
		name string
		node ast.Node
		want Type
	}{
		{
			name: "BlockExpr",
			node: &ast.BlockExpr{},
			want: &mockType{"Unit"},
		},

		{
			name: "CondExpr",
			node: &ast.CondExpr{
				Cases:    []ast.CondCase{{Predicate: &ast.TrueLiteral{}, Expr: &ast.IntLiteral{}}},
				ElseExpr: &ast.IntLiteral{},
			},
			want: &mockType{"Union"},
		},

		{
			name: "DeclareExpr",
			node: &ast.DeclareExpr{Symbol: "x", Expr: &ast.IntLiteral{}},
			want: &mockType{"Int"},
		},

		{
			name: "DefineExpr",
			node: &ast.DefineExpr{Symbol: "x", Expr: &ast.IntLiteral{}},
			want: &mockType{"Int"},
		},

		{
			name: "EllipsisLiteral",
			node: &ast.EllipsisLiteral{},
			want: &mockType{"Ellipsis"},
		},

		{
			name: "FalseLiteral",
			node: &ast.FalseLiteral{},
			want: &mockType{"Bool"},
		},

		{
			name: "FloatLiteral",
			node: &ast.FloatLiteral{},
			want: &mockType{"Float64"},
		},

		{
			name: "IntLiteral",
			node: &ast.IntLiteral{},
			want: &mockType{"Int"},
		},

		{
			name: "ReturnStmt",
			node: &ast.ReturnStmt{Expr: &ast.IntLiteral{}},
			want: &mockType{"Int"},
		},

		{
			name: "SetExpr",
			node: &ast.SetExpr{Symbol: "x", Expr: &ast.IntLiteral{}},
			want: &mockType{"Int"},
		},

		{
			name: "StringLiteral",
			node: &ast.StringLiteral{},
			want: &mockType{"String"},
		},

		{
			name: "TrueLiteral",
			node: &ast.TrueLiteral{},
			want: &mockType{"Bool"},
		},

		{
			name: "UnitExpr",
			node: &ast.UnitExpr{},
			want: &mockType{"Unit"},
		},

		{
			name: "WhileExpr",
			node: &ast.WhileExpr{Predicate: &ast.TrueLiteral{}, Expr: &ast.IntLiteral{}},
			want: &mockType{"Unit"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Check(ctx, env, tt.node)
			if err != nil {
				t.Fatalf("Check() error = %v", err)
			}
			if got.String() != tt.want.String() {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
