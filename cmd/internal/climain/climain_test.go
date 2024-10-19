// SPDX-License-Identifier: GPL-3.0-or-later

package climain_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/bassosimone/buresu/cmd/internal/climain"
	"github.com/bassosimone/buresu/cmd/internal/cliutils"
)

type fakecommand struct {
	err error
}

var _ cliutils.Command = fakecommand{}

// Help implements cliutils.Command.
func (f fakecommand) Help(argv ...string) error {
	return nil
}

// Main implements cliutils.Command.
func (f fakecommand) Main(ctx context.Context, argv ...string) error {
	return f.err
}

func TestRun(t *testing.T) {
	t.Run("when the command does not fail", func(t *testing.T) {
		cmd := fakecommand{nil}
		climain.Run(cmd, os.Exit)
	})

	t.Run("when the command fails", func(t *testing.T) {
		var exitcode int
		cmd := fakecommand{errors.New("mocked error")}
		climain.Run(cmd, func(code int) {
			exitcode = code
		})
		if exitcode != 1 {
			t.Fatal("did not call exit")
		}
	})
}
