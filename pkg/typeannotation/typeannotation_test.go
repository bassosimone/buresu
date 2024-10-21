// SPDX-License-Identifier: GPL-3.0-or-later

package typeannotation

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected *Annotation
		err      string
	}{
		{":: ", nil, "annotation is empty"},
		{":: foo", nil, "annotation is missing the `=>` separator"},
		{":: Int -> Int => Int", &Annotation{Params: []string{"Int", "Int"}, ReturnType: "Int"}, ""},
		{":: Int => Int", &Annotation{Params: []string{"Int"}, ReturnType: "Int"}, ""},
		{":: => Int", nil, "empty parameter type"},
		{":: Int -> => Int", nil, "empty parameter type"},
		{":: Int -> Int =>", nil, "empty return type"},
		{":: Int -> Int => Int -> Int", nil, "return type contains `->`"},
		{":: InvalidType -> Int => Int", &Annotation{Params: []string{"InvalidType", "Int"}, ReturnType: "Int"}, ""},
		{"", nil, "no type annotation found"},
		{":: Int -> Int => Int\n:: Float -> Float => Float", nil, "multiple type annotations found"},
		{":: Int ->  => Int", nil, "empty parameter type"},
		{"::  -> Int => Int", nil, "empty parameter type"},
	}

	for _, test := range tests {
		result, err := Parse(test.input)
		if err != nil {
			if test.err == "" || err.Error() != test.err {
				t.Errorf("Parse(%q) returned error %q, expected %q", test.input, err, test.err)
			}
		} else {
			if test.err != "" {
				t.Errorf("Parse(%q) expected error %q, got nil", test.input, test.err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Parse(%q) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestString(t *testing.T) {
	annotation := &Annotation{Params: []string{"Int", "Int"}, ReturnType: "Int"}
	expected := "Int -> Int => Int"
	result := annotation.String()
	if result != expected {
		t.Errorf("String() = %v, expected %v", result, expected)
	}
}

func TestArgumentsAnnotationPrefix(t *testing.T) {
	annotation := &Annotation{Params: []string{"Int", "Int"}, ReturnType: "Int"}
	expected := "Int -> Int => "
	result := annotation.ArgumentsAnnotationPrefix()
	if result != expected {
		t.Errorf("ArgumentsAnnotationPrefix() = %v, expected %v", result, expected)
	}
}

func TestMatchesArgumentsAnnotationPrefix(t *testing.T) {
	annotation := &Annotation{Params: []string{"Int", "Int"}, ReturnType: "Int"}
	prefix := "Int -> Int => "
	if !annotation.MatchesArgumentsAnnotationPrefix(prefix) {
		t.Errorf("MatchesArgumentsAnnotationPrefix(%v) = false, expected true", prefix)
	}
	prefix = "Int -> Float => "
	if annotation.MatchesArgumentsAnnotationPrefix(prefix) {
		t.Errorf("MatchesArgumentsAnnotationPrefix(%v) = true, expected false", prefix)
	}
}
