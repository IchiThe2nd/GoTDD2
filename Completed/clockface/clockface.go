package clockface

import (
	"math"
	"time"
)

const (
	secondsInhalfClock = 30
	secondsInClock     = secondsInhalfClock * 2
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hourInHalfClock    = 6
	hourInClock        = 2 * hourInHalfClock
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// Seconds

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func minutesInRadians(t time.Time) float64 { // radians from seconds + radians from minutes
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 { // radians from seconds + radians from minutes
	return (minutesInRadians(t) / 12) + math.Pi/(6/float64(t.Hour()%12))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
