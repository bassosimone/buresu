package includer

import (
	"fmt"
	"os"
	"strings"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

// Mock the os.ReadFile function
var readFile = os.ReadFile

// Include processes the given AST and handles include statements.
func Include(nodes []ast.Node) ([]ast.Node, error) {
	inc := newIncluder()
	return inc.includeNodes(nodes)
}

// Error represents a parsing error with position and message.
type Error struct {
	Tok     token.Token
	Message string
}

// Error returns the error message with file position details.
func (e *Error) Error() string {
	return fmt.Sprintf(
		"%s:%d:%d: includer: %s",
		e.Tok.TokenPos.FileName,
		e.Tok.TokenPos.LineNumber,
		e.Tok.TokenPos.LineColumn,
		e.Message,
	)
}

// newError formats and returns a new parser error including the token context.
func newError(tok token.Token, format string, args ...any) *Error {
	return &Error{Tok: tok, Message: fmt.Sprintf(format, args...)}
}

// includer processes the AST and handles include statements.
type includer struct {
	cycle   map[string]struct{}
	visited map[string]struct{}
}

// newIncluder creates a new includer instance.
func newIncluder() *includer {
	return &includer{
		cycle:   map[string]struct{}{},
		visited: map[string]struct{}{},
	}
}

// parseIncludeNode parses an include node and returns the file path if found.
//
// Returns:
//
// - filenode: the node containing the path to include (nil if the node is not
// an include statement or if the node is an invalid include statement);
//
// - found: true if the node is an include statement, false otherwise or
// if the node is an invalid include statement;
//
// - err: an error if the node is an invalid include statement, nil otherwise.
func parseIncludeNode(node ast.Node) (filenode *ast.StringLiteral, found bool, err error) {
	callExpr, ok := node.(*ast.CallExpr)
	if !ok {
		return nil, false, nil
	}
	symbol, ok := callExpr.Callable.(*ast.SymbolName)
	if !ok || symbol.Value != "include" {
		return nil, false, nil
	}
	if len(callExpr.Args) != 1 {
		return nil, false, newError(callExpr.Token, "include expects exactly one argument")
	}
	arg0, ok := callExpr.Args[0].(*ast.StringLiteral)
	if !ok {
		return nil, false, newError(callExpr.Token, "include expects a string argument")
	}
	return arg0, true, nil
}

// includeNodes processes the given AST and handles include statements.
func (inc *includer) includeNodes(nodes []ast.Node) ([]ast.Node, error) {
	var result []ast.Node
	for _, node := range nodes {

		// filter include nodes and bail in case of error
		filenode, found, err := parseIncludeNode(node)
		if err != nil {
			return nil, err
		}

		// if the node is an include expression, include the file
		if found {
			processedNodes, err := inc.includeFileOnce(filenode.Token, filenode.Value)
			if err != nil {
				return nil, err
			}
			result = append(result, processedNodes...)
			continue
		}

		// otherwise just make sure there are no nested includes expressions
		if err := inc.validateNoNestedIncludes(node); err != nil {
			return nil, err
		}
		result = append(result, node)
	}
	return result, nil
}

// includeFileOnce includes a file unless it has already been included and
// returns an error if we detect an inclusion cycle.
func (inc *includer) includeFileOnce(tok token.Token, filename string) ([]ast.Node, error) {
	// Detect inclusion cycles
	if _, ok := inc.cycle[filename]; ok {
		return nil, newError(tok, "inclusion cycle detected for file %s", filename)
	}

	// Make sure we have not already included this file
	if _, ok := inc.visited[filename]; ok {
		return nil, nil
	}

	// Mark the file as being under processing right now
	inc.cycle[filename] = struct{}{}

	// Unconditionally include
	nodes, err := inc.includeFile(tok, filename)
	if err != nil {
		return nil, err
	}

	// Mark the file as visited
	inc.visited[filename] = struct{}{}

	// Recursively include based on the current set of nodes
	result, err := inc.includeNodes(nodes)

	// Stop visiting the file
	delete(inc.cycle, filename)

	// Return the results
	return result, err
}

// includeFile includes a file and returns all its nodes.
func (inc *includer) includeFile(tok token.Token, filename string) ([]ast.Node, error) {
	// Load the file content.
	content, err := readFile(filename)
	if err != nil {
		return nil, newError(tok, "failed to read file %s: %v", filename, err)
	}

	// Scan the file content.
	tokens, err := scanner.Scan(filename, strings.NewReader(string(content)))
	if err != nil {
		return nil, newError(tok, "failed to scan file %s: %v", filename, err)
	}

	// Parse the tokens.
	includedNodes, err := parser.Parse(tokens)
	if err != nil {
		return nil, newError(tok, "failed to parse file %s: %v", filename, err)
	}

	return includedNodes, nil
}

// validateNoNestedIncludes checks for nested include statements and returns an error if found.
func (inc *includer) validateNoNestedIncludes(node ast.Node) error {
	switch n := node.(type) {
	case *ast.BlockExpr:
		for _, expr := range n.Exprs {
			if err := inc.validateNoNestedIncludes(expr); err != nil {
				return err
			}
		}
		return nil

	case *ast.CondExpr:
		for _, c := range n.Cases {
			if err := inc.validateNoNestedIncludes(c.Predicate); err != nil {
				return err
			}
			if err := inc.validateNoNestedIncludes(c.Expr); err != nil {
				return err
			}
		}
		if err := inc.validateNoNestedIncludes(n.ElseExpr); err != nil {
			return err
		}
		return nil

	case *ast.LambdaExpr:
		return inc.validateNoNestedIncludes(n.Expr)

	case *ast.WhileExpr:
		if err := inc.validateNoNestedIncludes(n.Predicate); err != nil {
			return err
		}
		if err := inc.validateNoNestedIncludes(n.Expr); err != nil {
			return err
		}
		return nil

	case *ast.CallExpr:
		if _, found, err := parseIncludeNode(n); found || err != nil {
			if err != nil {
				return err
			}
			return newError(n.Token, "include statement must be at top level")
		}

		if err := inc.validateNoNestedIncludes(n.Callable); err != nil {
			return err
		}
		for _, arg := range n.Args {
			if err := inc.validateNoNestedIncludes(arg); err != nil {
				return err
			}
		}
		return nil

	case *ast.ReturnStmt:
		return inc.validateNoNestedIncludes(n.Expr)

	case *ast.SetExpr:
		return inc.validateNoNestedIncludes(n.Expr)

	case *ast.DefineExpr:
		return inc.validateNoNestedIncludes(n.Expr)

	// Nothing to do for all other node types
	default:
		return nil
	}
}
