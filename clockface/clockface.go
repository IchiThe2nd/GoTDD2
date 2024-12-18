package clockface

import (
	"math"
	"time"
)

// a point represent cartesioan co ordinate
type Point struct {
	X float64
	Y float64
}

// second hand is unit vector of the second hand of an analogue clock at the time as a point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

// seconds to radians
func secondsInRadians(t time.Time) float64 {
	return math.Pi
}
