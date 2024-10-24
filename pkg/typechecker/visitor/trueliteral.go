// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkTrueLiteral(_ context.Context, env Environment, _ *ast.TrueLiteral) (Type, error) {
	return env.NewBoolType(), nil
}
