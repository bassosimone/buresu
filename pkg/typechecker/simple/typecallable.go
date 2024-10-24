// SPDX-License-Identifier: GPL-3.0-or-later

package simple

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bassosimone/buresu/internal/rtx"
	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// ErrWrongNumberOfArguments is the error returned when the number of arguments.
var ErrWrongNumberOfArguments = fmt.Errorf("wrong number of arguments")

// ErrWrongArgumentType is the error returned when the argument type is wrong.
var ErrWrongArgumentType = fmt.Errorf("wrong argument type")

// ErrWrongReturnType is the error returned when the return type is wrong.
var ErrWrongReturnType = fmt.Errorf("wrong return type")

// Callable represents a callable object.
type Callable struct {
	// ParamsTypes is a list of types that the callable expects.
	ParamsTypes []visitor.Type

	// ReturnType is the type that the callable returns.
	ReturnType visitor.Type

	// Body is the body of the callable.
	//
	// For built-in functions, there is no need to push a function
	// scope, since we're not running inside the program stack, and
	// no need to reduce the return type, because we control the
	// return type easily.
	//
	// Conversely, lambdas, which run in their own closure, will
	// always need to push a function scope and reduce the return
	// value, which they can to using their closure.
	Body func(ctx context.Context, args ...visitor.Type) (visitor.Type, error)

	// Previous is the optional previous callable overload.
	Previous *Callable
}

// Call calls the given callable with the given arguments.
func (c *Callable) Call(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
	var errs []error
	for cur := c; cur != nil; cur = cur.Previous {
		// attempt overloaded function resolution using the arguments
		if err := checkCallableArguments(cur.ParamsTypes, args); err != nil {
			err = fmt.Errorf("failed to call %s:\n    %w", cur.String(), err)
			errs = append(errs, err)
			continue
		}

		// call the actual callable
		//
		// it's safe to unconditionally cast since this is the package
		// that is calling the shots and defining the environment
		return cur.call(ctx, args...)
	}
	rtx.Assert(len(errs) > 0, "no errors collected")
	return nil, errors.Join(errs...)
}

func (c *Callable) call(ctx context.Context, args ...visitor.Type) (visitor.Type, error) {
	// issue the proper call and get the return type
	//
	// Note: the body will push a function scope for lambdas and there
	// is no need to have a function scope for built-ins.
	//
	// Note: the body should also merge the expression type with the
	// types returned by `(return! <expr>)`.
	rvType, err := c.Body(ctx, args...)
	if err != nil {
		return nil, err
	}

	// make sure the return type is the expected return type
	if !sameType(rvType, c.ReturnType) {
		err := fmt.Errorf("%w: expected %s, got %s", ErrWrongReturnType, c.ReturnType.String(), rvType.String())
		return nil, err
	}

	return rvType, nil
}

func checkCallableArguments(params, args []visitor.Type) error {
	// make a copy of the params because we may need to edit them
	// to account for variadic arguments
	params = append([]visitor.Type{}, params...)

	// make sure variadic parameters are at the end of the list
	var isvariadic bool
	for idx, param := range params {
		if _, ok := param.(*Variadic); ok {
			if idx < len(params)-1 {
				return fmt.Errorf("%w: variadic parameter must be the last one", ErrWrongNumberOfArguments)
			}
			isvariadic = true
		}
	}

	// if we have variadic parameters and the provided arguments
	// are more than the count of parameters, inflate the parameters
	// to match the number of arguments
	if isvariadic && len(args) > len(params)-1 {
		variadic := params[len(params)-1].(*Variadic)
		kind := variadic.Type
		params[len(params)-1] = kind
		for len(params) < len(args) {
			params = append(params, kind)
		}
	}

	// ensure that the number of arguments is correct
	if len(params) != len(args) {
		err := fmt.Errorf("%w: expected %d, got %d",
			ErrWrongNumberOfArguments, len(params), len(args))
		return err
	}

	// ensure that the types of the arguments are correct
	for idx := 0; idx < len(args); idx++ {
		if !sameType(params[idx], args[idx]) {
			err := fmt.Errorf(
				"%w for param #%d expected %s, got %s",
				ErrWrongArgumentType,
				idx+1,
				params[idx].String(),
				args[idx].String(),
			)
			return err
		}
	}

	return nil
}

// String implements visitor.Callable.
func (c *Callable) String() string {
	paramsTypes := make([]string, 0, len(c.ParamsTypes))
	for _, param := range c.ParamsTypes {
		paramsTypes = append(paramsTypes, param.String())
	}
	return fmt.Sprintf("(Callable (%s) %s)", strings.Join(paramsTypes, " "), c.ReturnType.String())
}
