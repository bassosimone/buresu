package simple

import (
	"strings"
	"testing"

	"github.com/bassosimone/buresu/pkg/typechecker/visitor"
	"github.com/google/go-cmp/cmp"
)

func TestParseAnnotation(t *testing.T) {
	tests := []struct {
		input         string
		expected      *Callable
		expectedError string
	}{
		{
			input: "(Callable (Int) Int)",
			expected: &Callable{
				ParamsTypes: []visitor.Type{&Int{}},
				ReturnType:  &Int{},
			},
		},
		{
			input: "(Callable () Unit)",
			expected: &Callable{
				ParamsTypes: []visitor.Type{},
				ReturnType:  &Unit{},
			},
		},
		{
			input: "(Callable ((Callable (Int) Int) (Variadic Int)) Unit)",
			expected: &Callable{
				ParamsTypes: []visitor.Type{
					&Callable{
						ParamsTypes: []visitor.Type{&Int{}},
						ReturnType:  &Int{},
					},
					&Variadic{&Int{}},
				},
				ReturnType: &Unit{},
			},
		},
		{
			input: "(Callable (String Bool) Unit)",
			expected: &Callable{
				ParamsTypes: []visitor.Type{&String{}, &Bool{}},
				ReturnType:  &Unit{},
			},
		},
		{
			input: "(Callable () (Union Int Unit))",
			expected: &Callable{
				ParamsTypes: []visitor.Type{},
				ReturnType: func() *Union {
					u := NewUnion()
					u.Add(&Int{})
					u.Add(&Unit{})
					return u
				}(),
			},
		},
		{
			input: "(Callable (Float64) Any)",
			expected: &Callable{
				ParamsTypes: []visitor.Type{&Float64{}},
				ReturnType:  &Any{},
			},
		},
		// Error cases
		{
			input:         "(Callable (Int) )",
			expectedError: "<annotation>:1:17: annotation parser: expected '(' or an atom",
		},
		{
			input:         "(Callable (Int Int)",
			expectedError: "<annotation>:1:19: annotation parser: expected '(' or an atom",
		},
		{
			input:         "(Callable (Int) Callable)",
			expectedError: "<annotation>:1:17: annotation parser: unknown type: Callable",
		},
		{
			input:         "(Callable (Int) (Union Int)",
			expectedError: "<annotation>:1:27: annotation parser: expected ')'",
		},
		{
			input:         "(Callable (Int) (Variadic))",
			expectedError: "<annotation>:1:26: annotation parser: expected '(' or an atom",
		},
		{
			input:         "(Callable (Int) (UnknownType))",
			expectedError: "<annotation>:1:17: annotation parser: expected 'Callable', 'Union' or 'Variadic'",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tokens, err := scanAnnotation(strings.NewReader(test.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			parser := newAnnotationParser(tokens)
			callable, err := parser.Parse()
			if test.expectedError != "" {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if diff := cmp.Diff(test.expectedError, err.Error()); diff != "" {
					t.Errorf(diff)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(test.expected.String(), callable.String()); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
