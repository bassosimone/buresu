// SPDX-License-Identifier: GPL-3.0-or-later

package includer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/token"
)

// Mock the os.ReadFile function
var readFile = os.ReadFile

// Include processes the given AST and handles include statements.
func Include(basePath string, nodes []ast.Node) ([]ast.Node, error) {
	inc := newIncluder(basePath)
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
	// basePath is the base path for all included files.
	basePath string

	// cycle is a temporary map to detect whether we are in an inclusion cycle.
	cycle map[string]struct{}

	// visited is a persistent map to detect whether we have already visited a file.
	visited map[string]struct{}
}

// newIncluder creates a new includer instance.
func newIncluder(basePath string) *includer {
	return &includer{
		basePath: basePath,
		cycle:    map[string]struct{}{},
		visited:  map[string]struct{}{},
	}
}

// includeNodes processes the given AST and handles include statements.
func (inc *includer) includeNodes(nodes []ast.Node) ([]ast.Node, error) {
	var result []ast.Node
	for _, node := range nodes {

		// if the node is an include expression, include the file
		includenode, found := node.(*ast.IncludeStmt)
		if found {
			processedNodes, err := inc.includeFileOnce(includenode.Token, includenode.FilePath)
			if err != nil {
				return nil, err
			}
			result = append(result, processedNodes...)
			continue
		}

		// otherwise just append the node to the result
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

	// Make sure we have not already visited this file
	if _, ok := inc.visited[filename]; ok {
		return nil, nil
	}

	// Mark the file as being under processing right now, then
	// uncover the file and mark it as visited
	inc.cycle[filename] = struct{}{}
	nodes, err := inc.includeFile(tok, filename)
	if err != nil {
		return nil, err
	}
	inc.visited[filename] = struct{}{}

	// Recursively include based on the current set of nodes
	result, err := inc.includeNodes(nodes)

	// Stop visiting the file once we have finished
	// recursively including all its nodes
	delete(inc.cycle, filename)
	return result, err
}

// includeFile includes a file and returns all its nodes.
func (inc *includer) includeFile(tok token.Token, filename string) ([]ast.Node, error) {
	// Obtain a file name within the basepath.
	filename = filepath.Join(inc.basePath, filepath.FromSlash(filename))

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
