package visitor

import (
	"context"

	"github.com/bassosimone/buresu/pkg/ast"
	"github.com/bassosimone/buresu/pkg/token"
)

func normalContext() context.Context {
	return context.Background()
}

func canceledContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

type mockEnvironment struct {
	returnType Type
	err        error
}

func (m *mockEnvironment) AddReturnType(kind Type) error {
	return nil
}

func (m *mockEnvironment) DefineType(symbol string, value Type) error {
	return nil
}

func (m *mockEnvironment) CheckCondition(ctx context.Context, predicate ast.Node) error {
	return nil
}

func (m *mockEnvironment) Call(ctx context.Context, node ast.Node, args ...Type) (Type, error) {
	return nil, nil
}

func (m *mockEnvironment) GetType(symbol string) (Type, error) {
	return nil, nil
}

func (m *mockEnvironment) MergeReturnTypes(exprType Type) (Type, error) {
	return nil, nil
}

func (m *mockEnvironment) NewBoolType() Type {
	return &mockType{"Bool"}
}

func (m *mockEnvironment) NewLambdaType(node *ast.LambdaExpr) (Type, error) {
	return nil, nil
}

func (m *mockEnvironment) NewFloat64Type() Type {
	return &mockType{"Float64"}
}

func (m *mockEnvironment) NewIntType() Type {
	return &mockType{"Int"}
}

func (m *mockEnvironment) NewQuotedType(node *ast.QuoteExpr) Type {
	return nil
}

func (m *mockEnvironment) NewStringType() Type {
	return &mockType{"String"}
}

func (m *mockEnvironment) NewUnionType(types ...Type) Type {
	return &mockType{"Union"}
}

func (m *mockEnvironment) NewUnitType() Type {
	return &mockType{"Unit"}
}

func (m *mockEnvironment) PushBlockScope() Environment {
	return m
}

func (m *mockEnvironment) PushFunctionScope() Environment {
	return m
}

func (m *mockEnvironment) SetType(symbol string, value Type) error {
	return nil
}

func (m *mockEnvironment) WrapError(tok token.Token, err error) error {
	return err
}

type mockType struct {
	name string
}

func (m *mockType) String() string {
	return m.name
}
