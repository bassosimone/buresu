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

// TypeAnnotationPrefix implements CallableTrait.
func (oc *overloadedCallable) TypeAnnotationPrefix() string {
	return ""
}

// Type implements Value.
func (oc *overloadedCallable) Type() string {
	return "<overloaded callable>"
}

// Call implements CallableTrait.
func (oc *overloadedCallable) Call(ctx context.Context, env *Environment, args ...Value) (Value, error) {
	prefix := buildArgsAnnotationPrefix(args)

	callable, err := oc.findCallable(env, prefix)
	if err != nil {
		return nil, fmt.Errorf("no callable found for prefix: %q", prefix)
	}

	return callable.Call(ctx, env, args...)
}

// buildArgsAnnotationPrefix constructs the annotation prefix from the given arguments.
func buildArgsAnnotationPrefix(args []Value) *typeannotation.Annotation {
	annot := &typeannotation.Annotation{
		Params:     []typeannotation.Type{},
		ReturnType: typeannotation.Type{Name: ""},
	}
	for _, arg := range args {
		annot.Params = append(annot.Params, typeannotation.Type{Name: arg.Type()})
	}
	return annot
}

func (oc *overloadedCallable) findCallable(env *Environment, prefix *typeannotation.Annotation) (CallableTrait, error) {
	for _, callable := range oc.callables {
		ta, err := typeannotation.ParseString(callable.TypeAnnotationPrefix())
		if err != nil {
			continue
		}

		// direct match
		if ta.MatchesArgumentsAnnotationPrefix(prefix) {
			return callable, nil
		}

		// attempt using traits
		if oc.matchWithTraits(env, ta, prefix) {
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
	ta *typeannotation.Annotation, prefix *typeannotation.Annotation) bool {
	if len(prefix.Params) != len(ta.Params) {
		return false
	}
	for idx := 0; idx < len(ta.Params); idx++ {
		if !env.GetImplements(prefix.Params[idx].Name, ta.Params[idx].Name) {
			return false
		}
	}
	return true
}
