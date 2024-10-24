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
	"github.com/bassosimone/buresu/pkg/parser"
	"github.com/bassosimone/buresu/pkg/scanner"
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
func (cmd command) Main(_ context.Context, argv ...string) error {
	// Implementation note: we ignore the main context, which
	// is setup to handle ^C because we need a more granular
	// control over its handling. For example, ^C is both used
	// to interrupt long running evaluation and to stop the
	// editing of an incomplete input line in the REPL.

	// 1. intercept and handle -h, --help, help
	if cliutils.HelpRequested(argv...) {
		return cmd.Help()
	}

	// 2. create command-line parser
	clip := pflag.NewFlagSet("buresu repl", pflag.ContinueOnError)

	// 3. parse the command line
	if err := clip.Parse(argv[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "buresu repl: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu repl --help` for usage.\n")
		return err
	}

	// 4. parse positional arguments
	args := clip.Args()
	if len(args) > 0 {
		err := fmt.Errorf("expected no argument, got: %v", shellquote.Join(args...))
		fmt.Fprintf(os.Stderr, "buresu repl: %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Run `buresu repl --help` for usage.\n")
		return err
	}

	// 5. initialize the readline library
	rl, err := readline.New("> ")
	if err != nil {
		return fmt.Errorf("buresu repl: failed to initialize readline: %w", err)
	}
	defer rl.Close()

	// 6. create the runtime environment
	rootScope := evaluator.NewGlobalEnvironment(os.Stdout)

	// 7. arrange for buffer and prompt reset
	buffer := ""
	prompt := ">>> "
	resetBufferAndPrompt := func() {
		buffer = ""
		prompt = ">>> "
	}

	// 8. start the REPL loop
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

		evaluate(rootScope, nodes)
		resetBufferAndPrompt()
	}
}

func evaluate(rootScope *evaluator.Environment, nodes []ast.Node) {
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

	// 3. evaluate all nodes and print them on stdout
	for _, node := range nodes {
		value, err := evaluator.Eval(ctx, rootScope, node)
		if err != nil {
			fmt.Fprintf(os.Stderr, "evaluation error: %s\n", err.Error())
			return
		}
		fmt.Printf("%s\n", value.String())
	}
}
