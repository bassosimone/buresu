// SPDX-License-Identifier: GPL-3.0-or-later

package run

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"os"

	"github.com/bassosimone/buresu/cmd/internal/cliutils"
	"github.com/bassosimone/buresu/pkg/dumper"
	"github.com/bassosimone/buresu/pkg/evaluator"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/runtime"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/kballard/go-shellquote"
	"github.com/spf13/pflag"
)

// NewCommand creates the `buresu run` [cliutils.Command].
func NewCommand() cliutils.Command {
	return command{}
}

// command implements [cliutils.command].
type command struct{}

var _ cliutils.Command = command{}

//go:embed README.txt
var readme string

// Help implements [cliutils.Command].
func (cmd command) Help(argv ...string) error {
	fmt.Fprintf(os.Stdout, "%s\n", readme)
	return nil
}

// Main implements [cliutils.Command].
func (cmd command) Main(ctx context.Context, argv ...string) error {
	// 1. intercept and handle -h, --help, help
	if cliutils.HelpRequested(argv...) {
		return cmd.Help()
	}

	// 2. create command line parser
	clip := pflag.NewFlagSet("buresu run", pflag.ContinueOnError)

	// 3. add options to the parser
	var emit string
	clip.StringVarP(&emit, "emit", "E", "", "Emit specific output (tokens, ast)")

	// 4. parse the command line
	if err := clip.Parse(argv[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu run --help` for usage.\n")
		return err
	}

	// 5. parse positional arguments
	args := clip.Args()
	switch {
	case len(args) < 1:
		err := errors.New("no script specified")
		fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu run --help` for usage.\n")
		return err

	case len(args) > 1:
		err := fmt.Errorf("expected single script, got: %v", shellquote.Join(args...))
		fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu run --help` for usage.\n")
		return err
	}
	scriptFile := args[0]

	// 6. scan the script to produce tokens
	filep, err := os.Open(scriptFile)
	if err != nil {
		err := fmt.Errorf("buresu: cannot open script: %s", err.Error())
		fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
		return err
	}
	defer filep.Close()
	tokens, err := scanner.Scan(scriptFile, filep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
		return err // already wrapped
	}
	if emit == "tokens" {
		return dumper.DumpTokens(os.Stdout, tokens)
	}

	// 7. parse the tokens to produce an AST
	nodes, err := parser.Parse(tokens)
	if err != nil {
		fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
		return err // already wrapped
	}
	if emit == "ast" {
		return dumper.DumpAST(os.Stdout, nodes)
	}

	// 8. create the runtime environment
	rootScope := evaluator.NewGlobalScope(os.Stdout)
	runtime.InitRootScope(rootScope)

	// 9. evaluate the AST
	for _, node := range nodes {
		if _, err := evaluator.Eval(ctx, rootScope, node); err != nil {
			fmt.Fprintf(os.Stderr, "buresu run: %s\n", err.Error())
			return err // already wrapped
		}
	}
	return nil
}