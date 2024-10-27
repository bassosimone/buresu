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

// wrapNode wraps a single node with its type and value.
func wrapNode(node ast.Node) ast.Node {
	switch nx := node.(type) {
	case *ast.BlockExpr:
		wrappedExprs := make([]ast.Node, len(nx.Exprs))
		for i, expr := range nx.Exprs {
			wrappedExprs[i] = wrapNode(expr)
		}
		return &nodeWrapper{
			Type: "BlockExpr",
			Value: &ast.BlockExpr{
				Token: nx.Token,
				Exprs: wrappedExprs,
			},
		}

	case *ast.CallExpr:
		wrappedArgs := make([]ast.Node, len(nx.Args))
		for i, arg := range nx.Args {
			wrappedArgs[i] = wrapNode(arg)
		}
		return &nodeWrapper{
			Type: "CallExpr",
			Value: &ast.CallExpr{
				Token:    nx.Token,
				Callable: wrapNode(nx.Callable),
				Args:     wrappedArgs,
			},
		}

	case *ast.CondExpr:
		wrappedCases := make([]ast.CondCase, len(nx.Cases))
		for i, c := range nx.Cases {
			wrappedCases[i] = ast.CondCase{
				Predicate: wrapNode(c.Predicate),
				Expr:      wrapNode(c.Expr),
			}
		}
		return &nodeWrapper{
			Type: "CondExpr",
			Value: &ast.CondExpr{
				Token:    nx.Token,
				Cases:    wrappedCases,
				ElseExpr: wrapNode(nx.ElseExpr),
			},
		}

	case *ast.DefineExpr:
		return &nodeWrapper{
			Type: "DefineExpr",
			Value: &ast.DefineExpr{
				Token:  nx.Token,
				Symbol: nx.Symbol,
				Expr:   wrapNode(nx.Expr),
			},
		}

	case *ast.EllipsisLiteral:
		return &nodeWrapper{
			Type: "EllipsisLiteral",
			Value: &ast.EllipsisLiteral{
				Token: nx.Token,
			},
		}

	case *ast.FalseLiteral:
		return &nodeWrapper{
			Type:  "FalseLiteral",
			Value: nx,
		}

	case *ast.FloatLiteral:
		return &nodeWrapper{
			Type:  "FloatLiteral",
			Value: nx,
		}

	case *ast.IntLiteral:
		return &nodeWrapper{
			Type:  "IntLiteral",
			Value: nx,
		}

	case *ast.LambdaExpr:
		return &nodeWrapper{
			Type: "LambdaExpr",
			Value: &ast.LambdaExpr{
				Token:  nx.Token,
				Params: nx.Params,
				Docs:   nx.Docs,
				Expr:   wrapNode(nx.Expr),
			},
		}

	case *ast.QuoteExpr:
		return &nodeWrapper{
			Type: "QuoteExpr",
			Value: &ast.QuoteExpr{
				Token: nx.Token,
				Expr:  wrapNode(nx.Expr),
			},
		}

	case *ast.ReturnStmt:
		return &nodeWrapper{
			Type: "ReturnStmt",
			Value: &ast.ReturnStmt{
				Token: nx.Token,
				Expr:  wrapNode(nx.Expr),
			},
		}

	case *ast.SetExpr:
		return &nodeWrapper{
			Type: "SetExpr",
			Value: &ast.SetExpr{
				Token:  nx.Token,
				Symbol: nx.Symbol,
				Expr:   wrapNode(nx.Expr),
			},
		}

	case *ast.StringLiteral:
		return &nodeWrapper{
			Type:  "StringLiteral",
			Value: nx,
		}

	case *ast.SymbolName:
		return &nodeWrapper{
			Type:  "SymbolName",
			Value: nx,
		}

	case *ast.TrueLiteral:
		return &nodeWrapper{
			Type:  "TrueLiteral",
			Value: nx,
		}

	case *ast.UnitExpr:
		return &nodeWrapper{
			Type:  "UnitExpr",
			Value: nx,
		}

	case *ast.WhileExpr:
		return &nodeWrapper{
			Type: "WhileExpr",
			Value: &ast.WhileExpr{
				Token:     nx.Token,
				Predicate: wrapNode(nx.Predicate),
				Expr:      wrapNode(nx.Expr),
			},
		}

	default:
		return &nodeWrapper{
			Type:  "Unknown",
			Value: nx,
		}
	}
}
