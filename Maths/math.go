package clockface

import "time"

// a point represent cartesioan co ordinate
type Point struct {
	X float64
	Y float64
}

// second hand is unit vector of the second hand of an analogue clock at the time as a point
func SecondHand(t time.Time) Point {
	return Point{}
}
