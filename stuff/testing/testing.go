package testing

import (
	"sus/stuff"
	"testing"
)

func AssertEq(a, b interface{}, t *testing.T) {
	if stuff.IsNilExt(a) && stuff.IsNilExt(b) {
		return
	}

	if a != b {
		t.Errorf("AssertEq failed:\n%#v\n%#v", a, b)
	}
}

func AssertEqF(a, b interface{}, f func(a, b interface{}) bool, t *testing.T) {
	if !f(a, b) {
		t.Errorf("AssertEqF failed:\n%#v\n%#v", a, b)
	}
}

func Assert(b bool, t *testing.T) {
	if !b {
		t.Error("Assertion failed")
	}
}
