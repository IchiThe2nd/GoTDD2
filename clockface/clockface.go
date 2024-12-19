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
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   //scale
	p = Point{p.X, -p.Y}            //flip
	p = Point{p.X + 150, p.Y + 150} // translate to zero

	return p
}

// seconds to radians

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
