// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"path/filepath"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/includer"
	"github.com/bassosimone/buresu/pkg/token"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// NewGlobalEnvironment creates a new global environment loading the
// standard library runtime from the given base path.
func NewGlobalEnvironment(ctx context.Context, basePath string) (*Environment, error) {
	env := NewEnvironment()

	// define the `display` built-in function
	env.DefineType("display", &Callable{
		ParamsTypes: []visitor.Type{&Variadic{&Any{}}},
		ReturnType:  &Unit{},
		Body: func(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
			return &Unit{}, nil
		},
		Previous: nil,
	})

	// most of the standard library runtime is defined in the runtime.brs file
	err := loadStdlibRuntime(ctx, basePath, env)
	return env, err
}

// loadStdlibRuntime loads the standard library runtime.
func loadStdlibRuntime(ctx context.Context, basePath string, tcEnv *Environment) error {
	// manually create the AST node for including the runtime
	runtime := &ast.IncludeStmt{
		Token: token.Token{
			TokenPos: token.Position{
				FileName:   "<runtime>",
				LineNumber: 1,
				LineColumn: 1,
			},
			TokenType: token.ATOM,
			Value:     "include",
		},
		FilePath: filepath.Join("stdlib", "runtime", "runtime.brs"),
	}
	nodes := []ast.Node{runtime}

	// use the includer to pull the nodes from the runtime file(s)
	nodes, err := includer.Include(basePath, nodes)
	if err != nil {
		return err
	}

	// run the typechecker on the runtime nodes
	for _, node := range nodes {
		if _, err := Check(ctx, tcEnv, node); err != nil {
			return err
		}
	}
	return nil
}
