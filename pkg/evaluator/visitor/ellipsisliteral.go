// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"
	"errors"

	"github.com/bassosimone/buresu/pkg/ast"
)

func evalEllipsisLiteral(_ context.Context, env Environment, node *ast.EllipsisLiteral) (Value, error) {
	return nil, env.WrapError(node.Token, errors.New("ellipsis cannot be used as a value"))
}
