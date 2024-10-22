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
	callable, ok := oc.findCallable(prefix, args...)
	if !ok {
		return nil, fmt.Errorf("no callable found for prefix: %q", prefix)
	}

	// call the callable
	return callable.Call(ctx, env, args...)
}

func (oc *overloadedCallable) findCallable(prefix string, args ...Value) (CallableTrait, bool) {
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

		// 1.2. turn the prefix into a full type annotation
		// pending a more exhaustive search for a match
		//
		// TODO(bassosimone): we should probably refactor this
		// code to avoid string manipulation and using better
		// representation of types and annotations.
		prefix += "Value"

		// 1.3. for now, we hand write checks to match with
		// type traits that matter, but it won't scale.
		//
		// TODO(bassosimone): implement more comprehensive
		// approach to type traits inference here.
		if oc.matchWithSequence(ta, prefix, args...) {
			return callable, true
		}
		if oc.matchWithValue(ta, prefix) {
			return callable, true
		}
	}

	// 2. fallback to the default callable without prefix
	// which implies generic arguments
	callable, ok := oc.callables[""]
	return callable, ok
}

// matchWithSequence just replaces the first argument in callable type
// annotation, if any, with Sequence and checks whether we actually implement
// sequence. In this case, we return true. This custom check is what we need
// to implement the `length` built-in function.
func (oc *overloadedCallable) matchWithSequence(
	ta *typeannotation.Annotation, fullproto string, args ...Value) bool {
	// TODO(bassosimone): implement more comprehensive
	// approach to type traits inference here.
	argta, err := typeannotation.ParseString(fullproto)
	if err != nil {
		return false
	}

	// ensure we actually have a sequence
	//
	// TODO(bassosimone): this should somehow be encoded into the type system
	// either in Go style or in Haskell style.
	if len(args) < 1 {
		return false
	}
	if _, ok := args[0].(SequenceTrait); !ok {
		return false
	}

	// we need at least one parameter to rewrite
	if len(argta.Params) < 1 {
		return false
	}

	// rewrite the param and retry matching
	argta.Params[0] = "Sequence"
	return ta.MatchesArgumentsAnnotationPrefix(argta.ArgumentsAnnotationPrefix())
}

// matchWithValue tries to make the fullproto with a type annotation containing
// all values, which is a custom check to implement `cons`.
func (oc *overloadedCallable) matchWithValue(ta *typeannotation.Annotation, fullproto string) bool {
	// TODO(bassosimone): implement more comprehensive
	// approach to type traits inference here.
	argta, err := typeannotation.ParseString(fullproto)
	if err != nil {
		return false
	}
	for idx := 0; idx < len(argta.Params); idx++ {
		argta.Params[idx] = "Value"
	}
	return ta.MatchesArgumentsAnnotationPrefix(argta.ArgumentsAnnotationPrefix())
}

// TypeAnnotationPrefix implements CallableTrait.
func (oc *overloadedCallable) TypeAnnotationPrefix() string {
	return ""
}

// Type implements Value.
func (oc *overloadedCallable) Type() string {
	return "<overloaded callable>"
}
