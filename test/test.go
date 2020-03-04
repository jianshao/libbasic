package test

import "testing"

func AssertEqual(t *testing.T, a interface{}, b interface{})  {
	t.Helper()
	if a != b {
		t.Errorf("not equal")
	}
}

func AssertNotEqual(t *testing.T, a interface{}, b interface{})  {
	t.Helper()
	if a == b {
		t.Errorf("equal")
	}
}

func AssertTrue(t *testing.T, a interface{})  {
	t.Helper()
	if a != true {
		t.Errorf("not true")
	}
}

func AssertFalse(t *testing.T, a interface{})  {
	t.Helper()
	if a != false {
		t.Errorf("not false")
	}
}

func AssertNil(t *testing.T, a interface{})  {
	t.Helper()
	if a != nil {
		t.Errorf("not nil")
	}
}

func AssertNotNil(t *testing.T, a interface{})  {
	t.Helper()
	if a == nil {
		t.Errorf("nil")
	}
}