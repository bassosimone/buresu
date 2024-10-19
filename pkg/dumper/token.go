// SPDX-License-Identifier: GPL-3.0-or-later

package dumper

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bassosimone/buresu/pkg/token"
)

// DumpTokens serializes the list of tokens to JSON and prints it.
func DumpTokens(writer io.Writer, tokens []token.Token) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(tokens); err != nil {
		return fmt.Errorf("failed to dump tokens: %w", err)
	}
	return nil
}
