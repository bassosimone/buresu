// SPDX-License-Identifier: GPL-3.0-or-later

// Package climain implements a command's main function.
package climain

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bassosimone/buresu/cmd/internal/cliutils"
)

// ExitFunc is the type of the [os.Exit] func.
type ExitFunc func(code int)

// make sure [os.Exit] implements [ExitFunc].
var _ = ExitFunc(os.Exit)

// Run runs the main function for the given command with the given [ExitFunc] and arguments.
func Run(cmd cliutils.Command, exitfn ExitFunc, argv ...string) {
	// 1. create cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 2. handle signals by canceling the context
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGINT)
	go func() {
		defer cancel()
		<-sch
	}()

	// 3. run the selected command.
	if err := cmd.Main(ctx, argv...); err != nil {
		exitfn(1)
	}
}
