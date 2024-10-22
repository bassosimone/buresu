// SPDX-License-Identifier: GPL-3.0-or-later

package evaluator

import (
	"context"
	"errors"
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
	callable, err := oc.findCallable(env, prefix)
	if err != nil {
		return nil, fmt.Errorf("no callable found for prefix: %q", prefix)
	}

	// call the callable
	return callable.Call(ctx, env, args...)
}

func (oc *overloadedCallable) findCallable(env *Environment, prefix string) (CallableTrait, error) {
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
			return callable, nil
		}

		// 1.2. turn the prefix into a full type annotation
		// pending a more exhaustive search for a match
		//
		// TODO(bassosimone): we should probably refactor this
		// code to avoid string manipulation and using better
		// representation of types and annotations.
		fullproto := prefix + "Value"

		// 1.3. run algorithm that checks whether each type
		// in the call matches the type in the callable through
		// type traits promotion (e.g., say we have `Sequence`
		// and `String`, we can replace `Sequence` with `String`
		// because `String` implements `Sequence`).
		if oc.matchWithTraits(env, ta, fullproto) {
			return callable, nil
		}
	}

	// 2. fallback to the default callable without prefix
	// which implies generic arguments
	callable, ok := oc.callables[""]
	if !ok {
		var buffer strings.Builder
		fmt.Fprintf(&buffer, "no callable found for prefix: %q\n", prefix)
		fmt.Fprintf(&buffer, "candidate callables:\n")
		for _, callable := range oc.callables {
			fmt.Fprintf(&buffer, "  %s\n", callable.String())
		}
		return nil, errors.New(buffer.String())
	}
	return callable, nil
}

// matchWithTraits checks whether each argument in the callable type
// can become a type trait in the type annotation. If so, we return true.
func (oc *overloadedCallable) matchWithTraits(env *Environment,
	ta *typeannotation.Annotation, fullproto string) bool {
	// map the fullproto to a type annotation and make sure the
	// number of arguments and parameters match
	argta, err := typeannotation.ParseString(fullproto)
	if err != nil {
		return false
	}
	if len(argta.Params) != len(ta.Params) {
		return false
	}
	if len(argta.Params) < 1 {
		return false
	}

	// check whether each argument in the callable type can
	// become the corresponding type annotation trait
	for idx := 0; idx < len(ta.Params); idx++ {
		if !env.GetImplements(argta.Params[idx], ta.Params[idx]) {
			return false
		}
	}
	return true
}

// TypeAnnotationPrefix implements CallableTrait.
func (oc *overloadedCallable) TypeAnnotationPrefix() string {
	return ""
}

// Type implements Value.
func (oc *overloadedCallable) Type() string {
	return "<overloaded callable>"
}
