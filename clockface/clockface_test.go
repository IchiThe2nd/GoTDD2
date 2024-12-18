package clockface_test

import (
	"math"
	"testing"
	"time"
)

func TestSecondInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if want != got {
		t.Fatalf("wanted %v radians , got %v", want, got)
	}

}
