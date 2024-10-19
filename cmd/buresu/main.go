// SPDX-License-Identifier: GPL-3.0-or-later

// Command buresu implements the `buresu` command.
package main

import (
	_ "embed"
	"os"

	"github.com/bassosimone/buresu/cmd/buresu/internal/run"
	"github.com/bassosimone/buresu/cmd/internal/climain"
	"github.com/bassosimone/buresu/cmd/internal/cliutils"
)

var mainArgs = os.Args

func main() {
	climain.Run(newCommand(), os.Exit, mainArgs...)
}

//go:embed README.txt
var readme string

// newCommand constructs a new [cliutils.Command] for the `buresu` command.
func newCommand() cliutils.Command {
	return cliutils.NewCommandWithSubCommands("buresu", readme, map[string]cliutils.Command{
		"run": run.NewCommand(),
	})
}
