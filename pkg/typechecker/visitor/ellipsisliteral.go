// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
)

func checkEllipsisLiteral(_ context.Context, env Environment, node *ast.EllipsisLiteral) (Type, error) {
	return env.NewEllipsisType(), nil
}
