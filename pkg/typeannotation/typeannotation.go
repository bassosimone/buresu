// SPDX-License-Identifier: GPL-3.0-or-later

// Package typeannotation manages lambda type annotations.
//
// A type annotation is a line inside the documentation of the lambda that
// specifies the type of the lambda using the following syntax:
//
//	:: <type>  ->... <type> => <type>
//
// Each <type> is either an atomic type or a composed type. For example:
//
// - `Int`, `Float`, `Unit`, `String`, `Bool`;
//
// - `(List Int)`, `(Dict String Unit)`.
//
// Overall, a lambda with type annotations looks like this:
//
//	(lambda (x y z) ":: Int -> Int -> Int => Int" ...)
//
// where the above annotation indicates that the three parameters are
// integers and the return value is an integer as well.
//
// A type annotation must always be specified including the return
// type. If the lambda does not take any arguments or does not return
// any value, use the `Unit` type. For example:
//
//	(lambda () ":: Unit => Unit" ...)
//
// where the above annotation indicates a lambda that takes no
// arguments and returns no value.
//
// An arguments annotation prefix is like a type annotation except that it
// omits the return value, which may not be known at call time. We use
// arguments annotations to search for candidate overloaded lambdas
// to call based on matching with the arguments annotation. For example:
//
//	:: Int -> Int =>
//
// this is an arguments annotation prefix that matches a lambda that
// takes two integers as arguments with unspecified return type.
//
// Note that type checking proper is not implemented by this package, since
// this package executes during parsing, where the type information is not
// available. Type checking will be done by the typechecker package.
package typeannotation

import (
	"errors"
	"strings"
)

// Annotation is a parsed type annotation.
type Annotation struct {
	// Params is the list of types of the parameters.
	Params []string

	// ReturnType is the return type.
	ReturnType string
}

// ErrNoTypeAnnotationFound is returned when no type annotations are found.
var ErrNoTypeAnnotationFound = errors.New("no type annotation found")

// ParseDocs parses the documentation string searching for a type annotation.
func ParseDocs(docs string) (*Annotation, error) {
	// for each line search for a line starting with `::`
	// if found, parse the annotation, then make sure there
	// are no more annotations in the documentation.
	var (
		annotation *Annotation
		err        error
	)
	for _, line := range strings.Split(docs, "\n") {
		if !strings.HasPrefix(line, "::") {
			continue
		}
		if annotation != nil {
			return nil, errors.New("multiple type annotations found")
		}
		annotation, err = ParseString(strings.TrimPrefix(line, "::"))
		if err != nil {
			return nil, err
		}
	}
	if annotation == nil {
		return nil, ErrNoTypeAnnotationFound
	}
	return annotation, nil
}

// ParseString parses a type annotation string after the `::` prefix
// has been stripped from the string itself.
func ParseString(annotation string) (*Annotation, error) {
	annotation = strings.TrimSpace(annotation)
	if annotation == "" {
		return nil, errors.New("annotation is empty")
	}

	parts := strings.Split(annotation, "=>")
	if len(parts) != 2 {
		return nil, errors.New("annotation is missing the `=>` separator")
	}

	rawParams := strings.Split(parts[0], "->")
	params := make([]string, 0, len(rawParams))
	for _, param := range rawParams {
		param = strings.TrimSpace(param)
		if param == "" {
			return nil, errors.New("empty parameter type")
		}
		params = append(params, param)
	}

	rt := strings.TrimSpace(parts[1])
	if rt == "" {
		return nil, errors.New("empty return type")
	}
	if strings.Contains(rt, "->") {
		return nil, errors.New("return type contains `->`")
	}

	result := &Annotation{Params: params, ReturnType: rt}
	return result, nil
}

// String returns the string representation of the annotation.
func (ap *Annotation) String() string {
	params := strings.Join(ap.Params, " -> ")
	return params + " => " + ap.ReturnType
}

// ArgumentsAnnotationPrefix returns the arguments annotation prefix.
func (ap *Annotation) ArgumentsAnnotationPrefix() string {
	params := strings.Join(ap.Params, " -> ")
	return params + " => "
}

// MatchesArgumentsAnnotationPrefix returns true if the annotation matches the
// given arguments annotation prefix string.
func (ap *Annotation) MatchesArgumentsAnnotationPrefix(prefix string) bool {
	return strings.HasPrefix(ap.String(), prefix)
}
