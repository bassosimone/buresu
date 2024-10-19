// SPDX-License-Identifier: GPL-3.0-or-later

package dumper

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bassosimone/buresu/pkg/ast"
)

// DumpAST serializes the AST nodes to JSON and prints them.
func DumpAST(writer io.Writer, nodes []ast.Node) error {
	wrappedNodes := wrapNodes(nodes)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(wrappedNodes); err != nil {
		return fmt.Errorf("failed to dump AST: %w", err)
	}
	return nil
}

// wrapNodes wraps each node with its type and value.
func wrapNodes(nodes []ast.Node) []ast.Node {
	var wrappedNodes []ast.Node
	for _, node := range nodes {
		wrappedNodes = append(wrappedNodes, wrapNode(node))
	}
	return wrappedNodes
}

// nodeWrapper wraps a node with its type and value.
type nodeWrapper struct {
	Type  string
	Value ast.Node
}

// Ensure NodeWrapper implements ast.Node
var _ ast.Node = (*nodeWrapper)(nil)

// String converts the NodeWrapper back to lisp source code.
func (nw *nodeWrapper) String() string {
	return fmt.Sprintf("(%s %s)", nw.Type, nw.Value.String())
}

// Clone creates a deep copy of the NodeWrapper.
func (nw *nodeWrapper) Clone() ast.Node {
	return &nodeWrapper{
		Type:  nw.Type,
		Value: nw.Value.Clone(),
	}
}

// wrapNode wraps a single node with its type and value.
func wrapNode(node ast.Node) ast.Node {
	switch n := node.(type) {
	case *ast.BlockExpr:
		wrappedExprs := make([]ast.Node, len(n.Exprs))
		for i, expr := range n.Exprs {
			wrappedExprs[i] = wrapNode(expr)
		}
		return &nodeWrapper{
			Type: "BlockExpr",
			Value: &ast.BlockExpr{
				Token: n.Token,
				Exprs: wrappedExprs,
			},
		}
	case *ast.CallExpr:
		wrappedArgs := make([]ast.Node, len(n.Args))
		for i, arg := range n.Args {
			wrappedArgs[i] = wrapNode(arg)
		}
		return &nodeWrapper{
			Type: "CallExpr",
			Value: &ast.CallExpr{
				Token:    n.Token,
				Callable: wrapNode(n.Callable),
				Args:     wrappedArgs,
			},
		}
	case *ast.CondExpr:
		wrappedCases := make([]ast.CondCase, len(n.Cases))
		for i, c := range n.Cases {
			wrappedCases[i] = ast.CondCase{
				Predicate: wrapNode(c.Predicate),
				Expr:      wrapNode(c.Expr),
			}
		}
		return &nodeWrapper{
			Type: "CondExpr",
			Value: &ast.CondExpr{
				Token:    n.Token,
				Cases:    wrappedCases,
				ElseExpr: wrapNode(n.ElseExpr),
			},
		}
	case *ast.DefineExpr:
		return &nodeWrapper{
			Type: "DefineExpr",
			Value: &ast.DefineExpr{
				Token:  n.Token,
				Symbol: n.Symbol,
				Expr:   wrapNode(n.Expr),
			},
		}
	case *ast.FalseLiteral:
		return &nodeWrapper{
			Type:  "FalseLiteral",
			Value: n,
		}
	case *ast.FloatLiteral:
		return &nodeWrapper{
			Type:  "FloatLiteral",
			Value: n,
		}
	case *ast.IntLiteral:
		return &nodeWrapper{
			Type:  "IntLiteral",
			Value: n,
		}
	case *ast.LambdaExpr:
		return &nodeWrapper{
			Type: "LambdaExpr",
			Value: &ast.LambdaExpr{
				Token:  n.Token,
				Params: n.Params,
				Docs:   n.Docs,
				Expr:   wrapNode(n.Expr),
			},
		}
	case *ast.ReturnStmt:
		return &nodeWrapper{
			Type: "ReturnStmt",
			Value: &ast.ReturnStmt{
				Token: n.Token,
				Expr:  wrapNode(n.Expr),
			},
		}
	case *ast.SetExpr:
		return &nodeWrapper{
			Type: "SetExpr",
			Value: &ast.SetExpr{
				Token:  n.Token,
				Symbol: n.Symbol,
				Expr:   wrapNode(n.Expr),
			},
		}
	case *ast.StringLiteral:
		return &nodeWrapper{
			Type:  "StringLiteral",
			Value: n,
		}
	case *ast.SymbolName:
		return &nodeWrapper{
			Type:  "SymbolName",
			Value: n,
		}
	case *ast.TrueLiteral:
		return &nodeWrapper{
			Type:  "TrueLiteral",
			Value: n,
		}
	case *ast.UnitExpr:
		return &nodeWrapper{
			Type:  "UnitExpr",
			Value: n,
		}
	case *ast.WhileExpr:
		return &nodeWrapper{
			Type: "WhileExpr",
			Value: &ast.WhileExpr{
				Token:     n.Token,
				Predicate: wrapNode(n.Predicate),
				Expr:      wrapNode(n.Expr),
			},
		}
	default:
		return &nodeWrapper{
			Type:  "Unknown",
			Value: n,
		}
	}
}
