package runtime_test

import (
	"testing"

	"github.com/bassosimone/buresu/pkg/runtime"
)

func TestBoolValue(t *testing.T) {
	v := &runtime.BoolValue{Value: true}
	if v.String() != "true" {
		t.Errorf("expected true, got %s", v.String())
	}
}

func TestIntValue(t *testing.T) {
	v := &runtime.IntValue{Value: 42}
	if v.String() != "42" {
		t.Errorf("expected 42, got %s", v.String())
	}
}

func TestFloatValue(t *testing.T) {
	v := &runtime.FloatValue{Value: 3.14}
	if v.String() != "3.140000" {
		t.Errorf("expected 3.140000, got %s", v.String())
	}
}

func TestStringValue(t *testing.T) {
	v := &runtime.StringValue{Value: "hello"}
	if v.String() != `"hello"` {
		t.Errorf("expected \"hello\", got %s", v.String())
	}
}

func TestUnitValue(t *testing.T) {
	v := &runtime.UnitValue{}
	if v.String() != "()" {
		t.Errorf("expected (), got %s", v.String())
	}
}
