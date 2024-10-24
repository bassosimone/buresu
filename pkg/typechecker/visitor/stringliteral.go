// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkStringLiteral(_ context.Context, env Environment, _ *ast.StringLiteral) (Type, error) {
	return env.NewStringType(), nil
}
