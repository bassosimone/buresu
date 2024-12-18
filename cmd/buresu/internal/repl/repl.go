// SPDX-License-Identifier: GPL-3.0-or-later

package repl

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bassosimone/buresu/cmd/internal/cliutils"
	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/evaluator"
	"github.com/bassosimone/buresu/pkg/includer"
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
	"github.com/bassosimone/buresu/pkg/typechecker"
	"github.com/chzyer/readline"
	"github.com/kballard/go-shellquote"
	"github.com/spf13/pflag"
)

// NewCommand creates the `buresu repl` [cliutils.Command].
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
	// Implementation note: we only use the main context for loading the
	// standard library runtime in the type checker. We do this because
	// the main context is setup to handle ^C but we need a more granular
	// control over signal handling in the REPL. For example, ^C is both
	// used to interrupt long running evaluation and to interrupt the
	// editing of an incomplete input line in the REPL.

	// 1. intercept and handle -h, --help, help
	if cliutils.HelpRequested(argv...) {
		return cmd.Help()
	}

	// 2. create command-line parser
	clip := pflag.NewFlagSet("buresu repl", pflag.ContinueOnError)

	// 3. add options to the parser
	var features []string
	clip.StringArrayVarP(&features, "feature", "X", []string{}, "Enable experimental features (e.g., typechecker)")

	// 4. parse the command line
	if err := clip.Parse(argv[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "buresu repl: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu repl --help` for usage.\n")
		return err
	}

	// 5. parse positional arguments
	args := clip.Args()
	if len(args) > 0 {
		err := fmt.Errorf("expected no argument, got: %v", shellquote.Join(args...))
		fmt.Fprintf(os.Stderr, "buresu repl: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu repl --help` for usage.\n")
		return err
	}

	// 6. create a map of enabled features
	enabledFeatures := make(map[string]struct{})
	for _, feature := range features {
		enabledFeatures[feature] = struct{}{}
	}

	// 7. initialize the readline library
	rl, err := readline.New("> ")
	if err != nil {
		err = fmt.Errorf("failed to initialize readline: %w", err)
		fmt.Fprintf(os.Stderr, "buresu repl: %s\n", err.Error())
		return err
	}
	defer rl.Close()

	// 8. create the runtime environment
	rootScope := evaluator.NewGlobalEnvironment(os.Stdout)
	tcEnv, err := typechecker.NewGlobalEnvironment(ctx, ".")
	if err != nil {
		err = fmt.Errorf("failed to load the standard library runtime: %w", err)
		fmt.Fprintf(os.Stderr, "buresu repl: %s\n", err.Error())
		return err
	}

	// 9. arrange for buffer and prompt reset
	buffer := ""
	prompt := ">>> "
	resetBufferAndPrompt := func() {
		buffer = ""
		prompt = ">>> "
	}

	// 10. start the REPL loop
	for {
		rl.SetPrompt(prompt)
		line, err := rl.Readline()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return err
			}
			if errors.Is(err, readline.ErrInterrupt) {
				resetBufferAndPrompt()
				continue
			}
			fmt.Fprintf(os.Stderr, "error reading input: %s\n", err.Error())
			continue
		}

		buffer += line
		tokens, err := scanner.Scan("<stdin>", strings.NewReader(buffer))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error scanning input: %s\n", err.Error())
			resetBufferAndPrompt()
			continue
		}

		nodes, err := parser.Parse(tokens)
		if err != nil {
			if parser.IsErrIncompleteInput(err) {
				prompt = "... "
				continue
			}
			fmt.Fprintf(os.Stderr, "syntax error: %s\n", err.Error())
			resetBufferAndPrompt()
			continue
		}

		nodes, err = includer.Include(".", nodes)
		if err != nil {
			fmt.Fprintf(os.Stderr, "include error: %s\n", err.Error())
			resetBufferAndPrompt()
			continue
		}

		evaluate(rootScope, tcEnv, nodes, enabledFeatures)
		resetBufferAndPrompt()
	}
}

func evaluate(
	rootScope *evaluator.Environment,
	tcEnv *typechecker.Environment,
	nodes []ast.Node,
	enabledFeatures map[string]struct{}) {
	// 1. create cancellable context for interrupt evaluation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2. make sure we correctly route SIGINT to the context
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGINT)
	defer signal.Stop(sch)
	go func() {
		select {
		case <-ctx.Done():
		case <-sch:
			cancel()
		}
	}()

	// 3. possibly typecheck, then evaluate all nodes and print them on stdout
	for _, node := range nodes {
		if _, typecheckerEnabled := enabledFeatures["typechecker"]; typecheckerEnabled {
			if _, err := typechecker.Check(ctx, tcEnv, node); err != nil {
				fmt.Fprintf(os.Stderr, "typechecking error: %s\n", err.Error())
				return
			}
		}

		value, err := evaluator.Eval(ctx, rootScope, node)
		if err != nil {
			fmt.Fprintf(os.Stderr, "evaluation error: %s\n", err.Error())
			return
		}

		fmt.Printf("%s\n", value.String())
	}
}
