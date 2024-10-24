package simple

import (
	"errors"

	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
)

// MergeReturnTypes implements [visitor.Environment].
func (env *Environment) MergeReturnTypes(exprType visitor.Type) (visitor.Type, error) {
	// make sure we are actually at function scope
	if env.flags&environmentFlagScopeFunc == 0 {
		return nil, errors.New("not at function scope")
	}

	// fetch the current return type and add the exprType
	rvType := env.rts
	rvType.Add(exprType)

	// simplify the return type
	return maybeMergeUnion(rvType), nil
}

func maybeMergeUnion(t visitor.Type) visitor.Type {
	// Base case: if the type is not a union, return it as is
	union, ok := t.(*Union)
	if !ok {
		return t
	}

	// Create a new union to collect flattened types
	flattenedUnion := NewUnion()
	for _, nestedType := range union.Types {
		reducedType := maybeMergeUnion(nestedType)
		nestedUnion, ok := reducedType.(*Union)
		if !ok {
			flattenedUnion.Add(reducedType)
			continue
		}
		// Flatten nested unions
		for _, nestedNestedType := range nestedUnion.Types {
			flattenedUnion.Add(nestedNestedType)
		}
	}

	// If the union is empty, return Unit
	if len(flattenedUnion.Types) <= 0 {
		return &Unit{}
	}

	// If the union contains only one type, return that type
	if len(flattenedUnion.Types) == 1 {
		for _, t := range flattenedUnion.Types {
			return t
		}
	}

	return flattenedUnion
}
