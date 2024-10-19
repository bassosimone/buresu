// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"fmt"

	"github.com/bassosimone/buresu/token"
)

// wrapError returns an error contextualized with the given token position.
func wrapError(tok token.Token, err error) error {
	return fmt.Errorf("%s: interpreter: %w", tok.TokenPos, err)
}

// newError formats and returns a new interpreter error including the token context.
func newError(tok token.Token, format string, args ...any) error {
	return wrapError(tok, fmt.Errorf(format, args...))
}
