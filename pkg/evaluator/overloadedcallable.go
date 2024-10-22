// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"fmt"
	"strings"

	"github.com/bassosimone/buresu/pkg/typeannotation"
)

// overloadedCallable is an environment entry containing
// one or more overloaded callable types.
//
// Construct using the [newOverloadedCallable] func.
type overloadedCallable struct {
	// callables is the list of callables.
	callables map[string]CallableTrait
}

var (
	_ CallableTrait = (*overloadedCallable)(nil)
	_ Value         = (*overloadedCallable)(nil)
)

// newOverloadedCallable creates a new overloaded callable instance.
func newOverloadedCallable() *overloadedCallable {
	return &overloadedCallable{
		callables: make(map[string]CallableTrait),
	}
}

// Add adds a new callable to the overloaded callable.
func (oc *overloadedCallable) Add(callable CallableTrait) error {
	prefix := callable.TypeAnnotationPrefix()
	if _, ok := oc.callables[prefix]; ok {
		return fmt.Errorf("overloaded callable already has a callable with prefix %s", prefix)
	}
	oc.callables[prefix] = callable
	return nil
}

// String implements Value.
func (oc *overloadedCallable) String() string {
	var builder strings.Builder
	for prefix, callable := range oc.callables {
		fmt.Fprintf(&builder, "%s :: %s\n", callable.String(), prefix)
	}
	return strings.TrimSpace(builder.String())
}

// Call implements CallableTrait.
func (oc *overloadedCallable) Call(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	// build the arguments annotation prefix
	annot := &typeannotation.Annotation{
		Params:     []string{},
		ReturnType: "",
	}
	for _, arg := range args {
		annot.Params = append(annot.Params, arg.Type())
	}
	prefix := annot.ArgumentsAnnotationPrefix()

	// find the callable with the given prefix
	callable, ok := oc.findCallable(prefix)
	if !ok {
		return nil, fmt.Errorf("no callable found for prefix: %q", prefix)
	}

	// call the callable
	return callable.Call(ctx, env, args...)
}

func (oc *overloadedCallable) findCallable(prefix string) (CallableTrait, bool) {
	// 1. attempt to match the prefix literally with the full
	// type annotation prefix of the callable.
	//
	// TODO(bassosimone): if we know the desired return type (how?)
	// here we can match with the full prefix as well.
	for _, callable := range oc.callables {

		// 1.1. attempt an exact match between args and params
		ta, err := typeannotation.ParseString(callable.TypeAnnotationPrefix())
		if err != nil {
			continue
		}
		if ta.MatchesArgumentsAnnotationPrefix(prefix) {
			return callable, true
		}

		// TODO(bassosimone): attempt to match via type traits
	}

	// 2. fallback to the default callable without prefix
	// which implies generic arguments
	callable, ok := oc.callables[""]
	return callable, ok
}

// TypeAnnotationPrefix implements CallableTrait.
func (oc *overloadedCallable) TypeAnnotationPrefix() string {
	return ""
}

// Type implements Value.
func (oc *overloadedCallable) Type() string {
	return "<overloaded callable>"
}
