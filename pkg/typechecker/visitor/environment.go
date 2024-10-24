// SPDX-License-Identifier: GPL-3.0-or-later

package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

// Environment is the generic interface for the environment.
type Environment interface {
	// AddReturnType adds the return type to the set of types that
	// the function can potentially return via `(return! ...)`.
	AddReturnType(kind Type) error

	// DefineType defines a new symbol type in the current environment.
	DefineType(symbol string, value Type) error

	// CheckCondition checks whether the current node evaluates
	// to a boolean value and otherwise returns an error.
	CheckCondition(ctx context.Context, predicate ast.Node) error

	// Call attempts to call a given node and returns the result type.
	Call(ctx context.Context, node ast.Node, args ...Type) (Type, error)

	// GetType returns the type associated with the given symbol.
	//
	// If the symbol is not found in the current environment, the parent
	// environments are searched recursively.
	GetType(symbol string) (Type, error)

	// MergeReturnTypes should merge the return types of a function
	// with the types the lambda body returned using `return!`.
	MergeReturnTypes(exprType Type) (Type, error)

	// NewBoolType returns a new bool type instance.
	NewBoolType() Type

	// NewLambdaType returns a new lambda type instance.
	NewLambdaType(node *ast.LambdaExpr) (Type, error)

	// NewFloat64Type returns a new float64 type instance.
	NewFloat64Type() Type

	// NewIntType returns a new int type instance.
	NewIntType() Type

	// NewQuotedType returns a new quoted type instance.
	NewQuotedType(node *ast.QuoteExpr) Type

	// NewStringType returns a new string type instance.
	NewStringType() Type

	// NewUnionType creates the union of the provided types.
	NewUnionType(types ...Type) Type

	// NewUnitType returns a new unit value instance.
	NewUnitType() Type

	// PushBlockScope creates a new child environment associated with
	// a block execution and returns it. The returned environment will
	// use the current environment as its parent.
	PushBlockScope() Environment

	// PushFunctionScope creates a new child environment associated with
	// a function execution and returns it. The returned environment will
	// use the current environment as its parent.
	PushFunctionScope() Environment

	// SetType sets the type of an existing symbol in the current environment.
	SetType(symbol string, value Type) error

	// WrapError wraps an error adding contextual token information.
	WrapError(tok token.Token, err error) error
}
