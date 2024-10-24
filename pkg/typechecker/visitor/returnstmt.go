// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/ast"
)

type errReturn struct {
	value Type
}

func (errReturn) Error() string {
	return "return statement"
}

func checkReturnStmt(ctx context.Context, env Environment, node *ast.ReturnStmt) (Type, error) {
	rvType, err := Check(ctx, env, node.Expr)
	if err != nil {
		return nil, err
	}
	// Note that AddReturnType may fail if we're not inside a function. The
	// parser should guarantee this can't happen, hence rtx.Must.
	rtx.Must(env.AddReturnType(rvType))

	// Technically the return! statement interrupts execution but for the
	// purpose of type checking we're considering it as the identity.
	//
	// Note that the parser only allows `return!` to appear as a statement
	// inside a block. Therefore, it's not possible to assign the return
	// value of return to any variable or use it in expresions.
	//
	// In light of this, having return returning its return type here seems
	// a pragrmatic decision to simplify typechecking without drawbacks.
	return rvType, nil
}
