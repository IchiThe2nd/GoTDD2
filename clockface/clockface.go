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
	return (math.Pi / (30 / (float64(t.Second()))))
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

//divide by zero
