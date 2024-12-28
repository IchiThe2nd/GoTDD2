package main

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %v got %v", want, got)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didnt want %v", got)
	}

}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v ,want True", got)
	}
}
func AssertNotTrue(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v ,want fucking False", got)
	}
}
