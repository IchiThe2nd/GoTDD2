package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"
	//var xys := "wfw"
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q wanted %q ", got, want)

	}
}
