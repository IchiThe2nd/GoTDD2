package main

import (
	"bytes"
	"fmt"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go`

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
	fmt.Print("spySleeper.Calls = ", spySleeper.Calls)
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper want 3 got %d", spySleeper.Calls)
	}
}
